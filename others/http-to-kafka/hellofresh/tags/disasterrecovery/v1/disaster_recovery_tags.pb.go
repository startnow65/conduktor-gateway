// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.1
// source: hellofresh/tags/disasterrecovery/v1/disaster_recovery_tags.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enumerates backup strategies for a Kafka topic
type BackupStrategy int32

const (
	BackupStrategy_BACKUP_STRATEGY_UNSPECIFIED BackupStrategy = 0
	// No backups required
	BackupStrategy_BACKUP_STRATEGY_NONE BackupStrategy = 1
	// Backup offsite
	BackupStrategy_BACKUP_STRATEGY_OFFSITE BackupStrategy = 2
)

// Enum value maps for BackupStrategy.
var (
	BackupStrategy_name = map[int32]string{
		0: "BACKUP_STRATEGY_UNSPECIFIED",
		1: "BACKUP_STRATEGY_NONE",
		2: "BACKUP_STRATEGY_OFFSITE",
	}
	BackupStrategy_value = map[string]int32{
		"BACKUP_STRATEGY_UNSPECIFIED": 0,
		"BACKUP_STRATEGY_NONE":        1,
		"BACKUP_STRATEGY_OFFSITE":     2,
	}
)

func (x BackupStrategy) Enum() *BackupStrategy {
	p := new(BackupStrategy)
	*p = x
	return p
}

func (x BackupStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackupStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_enumTypes[0].Descriptor()
}

func (BackupStrategy) Type() protoreflect.EnumType {
	return &file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_enumTypes[0]
}

func (x BackupStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackupStrategy.Descriptor instead.
func (BackupStrategy) EnumDescriptor() ([]byte, []int) {
	return file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescGZIP(), []int{0}
}

var file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*BackupStrategy)(nil),
		Field:         50000,
		Name:          "hellofresh.tags.disaster_recovery.v1.backup_strategy",
		Tag:           "varint,50000,opt,name=backup_strategy,enum=hellofresh.tags.disaster_recovery.v1.BackupStrategy",
		Filename:      "hellofresh/tags/disasterrecovery/v1/disaster_recovery_tags.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional hellofresh.tags.disaster_recovery.v1.BackupStrategy backup_strategy = 50000;
	E_BackupStrategy = &file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_extTypes[0]
)

var File_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto protoreflect.FileDescriptor

var file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDesc = []byte{
	0x0a, 0x40, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2f, 0x74, 0x61, 0x67,
	0x73, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x24, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2e, 0x74,
	0x61, 0x67, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x68, 0x0a, 0x0e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1f, 0x0a, 0x1b,
	0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x41, 0x43, 0x4b, 0x55,
	0x50, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4f, 0x46, 0x46, 0x53, 0x49,
	0x54, 0x45, 0x10, 0x02, 0x3a, 0x83, 0x01, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x34, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2e, 0x74,
	0x61, 0x67, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x88, 0x01, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x6f, 0x75, 0x74, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2f, 0x74,
	0x61, 0x67, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescOnce sync.Once
	file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescData = file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDesc
)

func file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescGZIP() []byte {
	file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescOnce.Do(func() {
		file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescData)
	})
	return file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDescData
}

var file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_goTypes = []interface{}{
	(BackupStrategy)(0),                 // 0: hellofresh.tags.disaster_recovery.v1.BackupStrategy
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_depIdxs = []int32{
	1, // 0: hellofresh.tags.disaster_recovery.v1.backup_strategy:extendee -> google.protobuf.MessageOptions
	0, // 1: hellofresh.tags.disaster_recovery.v1.backup_strategy:type_name -> hellofresh.tags.disaster_recovery.v1.BackupStrategy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_init() }
func file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_init() {
	if File_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_goTypes,
		DependencyIndexes: file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_depIdxs,
		EnumInfos:         file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_enumTypes,
		ExtensionInfos:    file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_extTypes,
	}.Build()
	File_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto = out.File
	file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_rawDesc = nil
	file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_goTypes = nil
	file_hellofresh_tags_disasterrecovery_v1_disaster_recovery_tags_proto_depIdxs = nil
}
