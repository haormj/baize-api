package service

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/haormj/baize-api/idl/dogsvscats"
	"github.com/haormj/baize-api/option"
	"github.com/haormj/ccutil/module/goapi/vision"
	"github.com/haormj/log"
	"github.com/haormj/util"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/micro/go-micro/v2/errors"
)

type DogsVsCatsService struct {
	opt       option.DogsVsCatsServiceOpt
	ossClient *oss.Client
	stsClient *sts.Client
	dvcModel  vision.DogsVsCats
}

func NewDogsVsCatsService(opt option.DogsVsCatsServiceOpt) (*DogsVsCatsService, error) {
	ossClient, err := oss.New(opt.Endpoint, opt.AccessKeyID, opt.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	stsClient, err := sts.NewClientWithAccessKey(opt.RegionID, opt.AccessKeyID, opt.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	dvcModel := vision.NewDogsVsCats()
	dvcModel.Load(opt.ModelPath)
	s := DogsVsCatsService{
		opt:       opt,
		ossClient: ossClient,
		stsClient: stsClient,
		dvcModel:  dvcModel,
	}
	return &s, nil
}

func (s *DogsVsCatsService) GetUploadParam(ctx context.Context, input *dogsvscats.GetUploadParamInput,
	output *dogsvscats.GetUploadParamOutput) error {
	policyTemplate := `{
    "Statement":[
        {
            "Action":[
                "*"
            ],
            "Effect":"Allow",
            "Resource":[
                "acs:oss:*:*:%s"
            ]
        }
    ],
    "Version":"1"
}
`
	objectKey := s.opt.UploadPathPrefix + util.GetUUIDV1WithoutLine()
	p := s.opt.Bucket + "/" + objectKey
	policy := fmt.Sprintf(policyTemplate, p)
	req := sts.CreateAssumeRoleRequest()
	req.SetScheme(requests.HTTPS)
	req.RoleSessionName = s.opt.RoleSessionName
	req.Policy = policy
	req.RoleArn = s.opt.RoleArn
	req.DurationSeconds = requests.NewInteger(15 * 60)
	rsp, err := s.stsClient.AssumeRole(req)
	if err != nil {
		log.Logger.Errorw("sts AssumeRole error", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	output.RegionId = "oss-" + s.opt.RegionID
	output.Bucket = s.opt.Bucket
	output.AccessKeyId = rsp.Credentials.AccessKeyId
	output.AccessKeySecret = rsp.Credentials.AccessKeySecret
	output.SecurityToken = rsp.Credentials.SecurityToken
	output.Expiration = rsp.Credentials.Expiration
	output.ObjectKey = objectKey
	return nil
}

func (s *DogsVsCatsService) Classify(ctx context.Context, input *dogsvscats.ClassifyInput,
	output *dogsvscats.ClassifyOutput) error {
	objectKey := input.ObjectKey
	bucket, err := s.ossClient.Bucket(s.opt.Bucket)
	if err != nil {
		log.Logger.Errorw("oss client bucket error", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	exist, err := bucket.IsObjectExist(objectKey)
	if err != nil {
		log.Logger.Errorw("bucket is object exist", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	if !exist {
		log.Logger.Debugw("object not exist", "objectKey", objectKey)
		return errors.NotFound("", "NotFound")
	}
	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 5*60)
	if err != nil {
		log.Logger.Errorw("sign url error", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	log.Logger.Debug("signedURL", signedURL)
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Errorw("panic", "err", r)
		}
	}()
	result := s.dvcModel.PredictByUrl(signedURL)
	_, file := path.Split(objectKey)
	dstObjectKey := s.opt.ClassifyPathPrefix + strings.ToLower(result) + "/" + file
	if _, err := bucket.CopyObject(objectKey, dstObjectKey); err != nil {
		log.Logger.Errorw("copy object error", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	if err := bucket.DeleteObject(objectKey); err != nil {
		log.Logger.Errorw("delete object error", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	dstObjectKeySignedURL, err := bucket.SignURL(dstObjectKey, oss.HTTPGet, 5*60)
	if err != nil {
		log.Logger.Errorw("sign url error", "err", err)
		return errors.InternalServerError("", "InternalServerError")
	}
	output.Result = result
	output.ImageUrl = dstObjectKeySignedURL
	return nil
}
