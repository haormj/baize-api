// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: idl/dogsvscats/dogs_vs_cats.proto

package dogsvscats

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DogsVsCatsService service

func NewDogsVsCatsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DogsVsCatsService service

type DogsVsCatsService interface {
	GetUploadParam(ctx context.Context, in *GetUploadParamInput, opts ...client.CallOption) (*GetUploadParamOutput, error)
	Classify(ctx context.Context, in *ClassifyInput, opts ...client.CallOption) (*ClassifyOutput, error)
}

type dogsVsCatsService struct {
	c    client.Client
	name string
}

func NewDogsVsCatsService(name string, c client.Client) DogsVsCatsService {
	return &dogsVsCatsService{
		c:    c,
		name: name,
	}
}

func (c *dogsVsCatsService) GetUploadParam(ctx context.Context, in *GetUploadParamInput, opts ...client.CallOption) (*GetUploadParamOutput, error) {
	req := c.c.NewRequest(c.name, "DogsVsCatsService.GetUploadParam", in)
	out := new(GetUploadParamOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dogsVsCatsService) Classify(ctx context.Context, in *ClassifyInput, opts ...client.CallOption) (*ClassifyOutput, error) {
	req := c.c.NewRequest(c.name, "DogsVsCatsService.Classify", in)
	out := new(ClassifyOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DogsVsCatsService service

type DogsVsCatsServiceHandler interface {
	GetUploadParam(context.Context, *GetUploadParamInput, *GetUploadParamOutput) error
	Classify(context.Context, *ClassifyInput, *ClassifyOutput) error
}

func RegisterDogsVsCatsServiceHandler(s server.Server, hdlr DogsVsCatsServiceHandler, opts ...server.HandlerOption) error {
	type dogsVsCatsService interface {
		GetUploadParam(ctx context.Context, in *GetUploadParamInput, out *GetUploadParamOutput) error
		Classify(ctx context.Context, in *ClassifyInput, out *ClassifyOutput) error
	}
	type DogsVsCatsService struct {
		dogsVsCatsService
	}
	h := &dogsVsCatsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DogsVsCatsService{h}, opts...))
}

type dogsVsCatsServiceHandler struct {
	DogsVsCatsServiceHandler
}

func (h *dogsVsCatsServiceHandler) GetUploadParam(ctx context.Context, in *GetUploadParamInput, out *GetUploadParamOutput) error {
	return h.DogsVsCatsServiceHandler.GetUploadParam(ctx, in, out)
}

func (h *dogsVsCatsServiceHandler) Classify(ctx context.Context, in *ClassifyInput, out *ClassifyOutput) error {
	return h.DogsVsCatsServiceHandler.Classify(ctx, in, out)
}
