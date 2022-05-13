// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: cpfd-proto.proto

package pb

import (
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

type ParticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticlePath string `protobuf:"bytes,1,opt,name=particle_path,proto3" json:"particle_path,omitempty"`
	ActivityPath string `protobuf:"bytes,2,opt,name=activity_path,proto3" json:"activity_path,omitempty"`
	Start        string `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End          string `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ParticleReq) Reset() {
	*x = ParticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpfd_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticleReq) ProtoMessage() {}

func (x *ParticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_cpfd_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticleReq.ProtoReflect.Descriptor instead.
func (*ParticleReq) Descriptor() ([]byte, []int) {
	return file_cpfd_proto_proto_rawDescGZIP(), []int{0}
}

func (x *ParticleReq) GetParticlePath() string {
	if x != nil {
		return x.ParticlePath
	}
	return ""
}

func (x *ParticleReq) GetActivityPath() string {
	if x != nil {
		return x.ActivityPath
	}
	return ""
}

func (x *ParticleReq) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *ParticleReq) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type ParticleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,proto3" json:"file_path,omitempty"`
}

func (x *ParticleRes) Reset() {
	*x = ParticleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpfd_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticleRes) ProtoMessage() {}

func (x *ParticleRes) ProtoReflect() protoreflect.Message {
	mi := &file_cpfd_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticleRes.ProtoReflect.Descriptor instead.
func (*ParticleRes) Descriptor() ([]byte, []int) {
	return file_cpfd_proto_proto_rawDescGZIP(), []int{1}
}

func (x *ParticleRes) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_cpfd_proto_proto protoreflect.FileDescriptor

var file_cpfd_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x70, 0x66, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x2b, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x32, 0x48, 0x0a, 0x0d, 0x44, 0x61,
	0x74, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0b, 0x47,
	0x65, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cpfd_proto_proto_rawDescOnce sync.Once
	file_cpfd_proto_proto_rawDescData = file_cpfd_proto_proto_rawDesc
)

func file_cpfd_proto_proto_rawDescGZIP() []byte {
	file_cpfd_proto_proto_rawDescOnce.Do(func() {
		file_cpfd_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_cpfd_proto_proto_rawDescData)
	})
	return file_cpfd_proto_proto_rawDescData
}

var file_cpfd_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cpfd_proto_proto_goTypes = []interface{}{
	(*ParticleReq)(nil), // 0: proto.ParticleReq
	(*ParticleRes)(nil), // 1: proto.ParticleRes
}
var file_cpfd_proto_proto_depIdxs = []int32{
	0, // 0: proto.DataGenerator.GenParticle:input_type -> proto.ParticleReq
	1, // 1: proto.DataGenerator.GenParticle:output_type -> proto.ParticleRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cpfd_proto_proto_init() }
func file_cpfd_proto_proto_init() {
	if File_cpfd_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cpfd_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticleReq); i {
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
		file_cpfd_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticleRes); i {
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
			RawDescriptor: file_cpfd_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cpfd_proto_proto_goTypes,
		DependencyIndexes: file_cpfd_proto_proto_depIdxs,
		MessageInfos:      file_cpfd_proto_proto_msgTypes,
	}.Build()
	File_cpfd_proto_proto = out.File
	file_cpfd_proto_proto_rawDesc = nil
	file_cpfd_proto_proto_goTypes = nil
	file_cpfd_proto_proto_depIdxs = nil
}
