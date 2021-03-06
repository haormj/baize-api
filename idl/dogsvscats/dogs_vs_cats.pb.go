// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.1
// source: idl/dogsvscats/dogs_vs_cats.proto

package dogsvscats

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetUploadParamInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUploadParamInput) Reset() {
	*x = GetUploadParamInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUploadParamInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUploadParamInput) ProtoMessage() {}

func (x *GetUploadParamInput) ProtoReflect() protoreflect.Message {
	mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUploadParamInput.ProtoReflect.Descriptor instead.
func (*GetUploadParamInput) Descriptor() ([]byte, []int) {
	return file_idl_dogsvscats_dogs_vs_cats_proto_rawDescGZIP(), []int{0}
}

type GetUploadParamOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId        string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Bucket          string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	AccessKeyId     string `protobuf:"bytes,3,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	AccessKeySecret string `protobuf:"bytes,4,opt,name=access_key_secret,json=accessKeySecret,proto3" json:"access_key_secret,omitempty"`
	SecurityToken   string `protobuf:"bytes,5,opt,name=security_token,json=securityToken,proto3" json:"security_token,omitempty"`
	Expiration      string `protobuf:"bytes,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	ObjectKey       string `protobuf:"bytes,7,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
}

func (x *GetUploadParamOutput) Reset() {
	*x = GetUploadParamOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUploadParamOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUploadParamOutput) ProtoMessage() {}

func (x *GetUploadParamOutput) ProtoReflect() protoreflect.Message {
	mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUploadParamOutput.ProtoReflect.Descriptor instead.
func (*GetUploadParamOutput) Descriptor() ([]byte, []int) {
	return file_idl_dogsvscats_dogs_vs_cats_proto_rawDescGZIP(), []int{1}
}

func (x *GetUploadParamOutput) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *GetUploadParamOutput) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GetUploadParamOutput) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *GetUploadParamOutput) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *GetUploadParamOutput) GetSecurityToken() string {
	if x != nil {
		return x.SecurityToken
	}
	return ""
}

func (x *GetUploadParamOutput) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *GetUploadParamOutput) GetObjectKey() string {
	if x != nil {
		return x.ObjectKey
	}
	return ""
}

type ClassifyInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectKey string `protobuf:"bytes,1,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
}

func (x *ClassifyInput) Reset() {
	*x = ClassifyInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyInput) ProtoMessage() {}

func (x *ClassifyInput) ProtoReflect() protoreflect.Message {
	mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyInput.ProtoReflect.Descriptor instead.
func (*ClassifyInput) Descriptor() ([]byte, []int) {
	return file_idl_dogsvscats_dogs_vs_cats_proto_rawDescGZIP(), []int{2}
}

func (x *ClassifyInput) GetObjectKey() string {
	if x != nil {
		return x.ObjectKey
	}
	return ""
}

type ClassifyOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *ClassifyOutput) Reset() {
	*x = ClassifyOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyOutput) ProtoMessage() {}

func (x *ClassifyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyOutput.ProtoReflect.Descriptor instead.
func (*ClassifyOutput) Descriptor() ([]byte, []int) {
	return file_idl_dogsvscats_dogs_vs_cats_proto_rawDescGZIP(), []int{3}
}

func (x *ClassifyOutput) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ClassifyOutput) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_idl_dogsvscats_dogs_vs_cats_proto protoreflect.FileDescriptor

var file_idl_dogsvscats_dogs_vs_cats_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x64, 0x6c, 0x2f, 0x64, 0x6f, 0x67, 0x73, 0x76, 0x73, 0x63, 0x61, 0x74, 0x73,
	0x2f, 0x64, 0x6f, 0x67, 0x73, 0x5f, 0x76, 0x73, 0x5f, 0x63, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x6f, 0x67, 0x73, 0x76, 0x73, 0x63, 0x61, 0x74, 0x73, 0x22,
	0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x81, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x2e, 0x0a, 0x0d, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x45, 0x0a, 0x0e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x32, 0xaf, 0x01, 0x0a, 0x11, 0x44, 0x6f, 0x67, 0x73, 0x56, 0x73, 0x43, 0x61, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x64, 0x6f, 0x67, 0x73,
	0x76, 0x73, 0x63, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x6f, 0x67,
	0x73, 0x76, 0x73, 0x63, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x67,
	0x73, 0x76, 0x73, 0x63, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x6f, 0x67, 0x73, 0x76, 0x73, 0x63, 0x61,
	0x74, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x64, 0x6f, 0x67, 0x73, 0x76,
	0x73, 0x63, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_dogsvscats_dogs_vs_cats_proto_rawDescOnce sync.Once
	file_idl_dogsvscats_dogs_vs_cats_proto_rawDescData = file_idl_dogsvscats_dogs_vs_cats_proto_rawDesc
)

func file_idl_dogsvscats_dogs_vs_cats_proto_rawDescGZIP() []byte {
	file_idl_dogsvscats_dogs_vs_cats_proto_rawDescOnce.Do(func() {
		file_idl_dogsvscats_dogs_vs_cats_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_dogsvscats_dogs_vs_cats_proto_rawDescData)
	})
	return file_idl_dogsvscats_dogs_vs_cats_proto_rawDescData
}

var file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_dogsvscats_dogs_vs_cats_proto_goTypes = []interface{}{
	(*GetUploadParamInput)(nil),  // 0: dogsvscats.GetUploadParamInput
	(*GetUploadParamOutput)(nil), // 1: dogsvscats.GetUploadParamOutput
	(*ClassifyInput)(nil),        // 2: dogsvscats.ClassifyInput
	(*ClassifyOutput)(nil),       // 3: dogsvscats.ClassifyOutput
}
var file_idl_dogsvscats_dogs_vs_cats_proto_depIdxs = []int32{
	0, // 0: dogsvscats.DogsVsCatsService.GetUploadParam:input_type -> dogsvscats.GetUploadParamInput
	2, // 1: dogsvscats.DogsVsCatsService.Classify:input_type -> dogsvscats.ClassifyInput
	1, // 2: dogsvscats.DogsVsCatsService.GetUploadParam:output_type -> dogsvscats.GetUploadParamOutput
	3, // 3: dogsvscats.DogsVsCatsService.Classify:output_type -> dogsvscats.ClassifyOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_dogsvscats_dogs_vs_cats_proto_init() }
func file_idl_dogsvscats_dogs_vs_cats_proto_init() {
	if File_idl_dogsvscats_dogs_vs_cats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUploadParamInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUploadParamOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_dogsvscats_dogs_vs_cats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_dogsvscats_dogs_vs_cats_proto_goTypes,
		DependencyIndexes: file_idl_dogsvscats_dogs_vs_cats_proto_depIdxs,
		MessageInfos:      file_idl_dogsvscats_dogs_vs_cats_proto_msgTypes,
	}.Build()
	File_idl_dogsvscats_dogs_vs_cats_proto = out.File
	file_idl_dogsvscats_dogs_vs_cats_proto_rawDesc = nil
	file_idl_dogsvscats_dogs_vs_cats_proto_goTypes = nil
	file_idl_dogsvscats_dogs_vs_cats_proto_depIdxs = nil
}
