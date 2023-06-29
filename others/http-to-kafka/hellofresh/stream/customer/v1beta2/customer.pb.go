// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.1
// source: hellofresh/stream/customer/v1beta2/customer.proto

package v1beta2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "github.com/kanapuli/http-to-kafka/hellofresh/tags/data_protection/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CustomerValue contains details about the Customer
// including PII such as name.
type CustomerValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Customer full name.
	//
	// [first_name + " " + last_name] from Intfood DB.
	//
	// first_name = Max 255 Characters UTF-8
	// last_name = Max 255 Characters UTF-8
	// Total Max 511 Characters
	//
	// This field is usually empty at the time of registration,
	// customers will provide their name during later steps of the checkout.
	// When customers register using social login, most of the times
	// the name will be present.
	//
	// Attention: Personal Identifiable Information (PII)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Customer email.
	//
	// Attention: Personal Identifiable Information (PII)
	// Alias: hellofresh.alias.EmailAddress
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// The Customer's favourite food
	// Some field to demonstrate non-pii data
	FavouriteFood string `protobuf:"bytes,3,opt,name=favourite_food,json=favouriteFood,proto3" json:"favourite_food,omitempty"`
	// The Customer's height in meters.
	// Some field to demonstrate non-pii data
	Height float32 `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *CustomerValue) Reset() {
	*x = CustomerValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellofresh_stream_customer_v1beta2_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerValue) ProtoMessage() {}

func (x *CustomerValue) ProtoReflect() protoreflect.Message {
	mi := &file_hellofresh_stream_customer_v1beta2_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerValue.ProtoReflect.Descriptor instead.
func (*CustomerValue) Descriptor() ([]byte, []int) {
	return file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerValue) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerValue) GetFavouriteFood() string {
	if x != nil {
		return x.FavouriteFood
	}
	return ""
}

func (x *CustomerValue) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_hellofresh_stream_customer_v1beta2_customer_proto protoreflect.FileDescriptor

var file_hellofresh_stream_customer_v1beta2_customer_proto_rawDesc = []byte{
	0x0a, 0x31, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x3d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0x80, 0xb5, 0x18, 0x01, 0x90, 0xb5, 0x18, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x66, 0x6f, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x06,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12,
	0x1c, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x04, 0x80, 0xb5, 0x18, 0x06, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0xd6, 0x01,
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x0d,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x30, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x74, 0x6f, 0x2d, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x32, 0x90, 0x01, 0x00, 0xd8, 0x01, 0x01, 0xca, 0x02, 0x28, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x46,
	0x72, 0x65, 0x73, 0x68, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0xe2, 0x02, 0x31, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x46, 0x72, 0x65, 0x73, 0x68, 0x5c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescOnce sync.Once
	file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescData = file_hellofresh_stream_customer_v1beta2_customer_proto_rawDesc
)

func file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescGZIP() []byte {
	file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescOnce.Do(func() {
		file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescData)
	})
	return file_hellofresh_stream_customer_v1beta2_customer_proto_rawDescData
}

var file_hellofresh_stream_customer_v1beta2_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hellofresh_stream_customer_v1beta2_customer_proto_goTypes = []interface{}{
	(*CustomerValue)(nil), // 0: hellofresh.stream.customer.v1beta2.CustomerValue
}
var file_hellofresh_stream_customer_v1beta2_customer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hellofresh_stream_customer_v1beta2_customer_proto_init() }
func file_hellofresh_stream_customer_v1beta2_customer_proto_init() {
	if File_hellofresh_stream_customer_v1beta2_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hellofresh_stream_customer_v1beta2_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerValue); i {
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
			RawDescriptor: file_hellofresh_stream_customer_v1beta2_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hellofresh_stream_customer_v1beta2_customer_proto_goTypes,
		DependencyIndexes: file_hellofresh_stream_customer_v1beta2_customer_proto_depIdxs,
		MessageInfos:      file_hellofresh_stream_customer_v1beta2_customer_proto_msgTypes,
	}.Build()
	File_hellofresh_stream_customer_v1beta2_customer_proto = out.File
	file_hellofresh_stream_customer_v1beta2_customer_proto_rawDesc = nil
	file_hellofresh_stream_customer_v1beta2_customer_proto_goTypes = nil
	file_hellofresh_stream_customer_v1beta2_customer_proto_depIdxs = nil
}
