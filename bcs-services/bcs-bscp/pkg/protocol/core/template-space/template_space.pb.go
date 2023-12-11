// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: template_space.proto

package pbts

import (
	base "bscp.io/pkg/protocol/core/base"
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

// TemplateSpace source resource reference: pkg/dal/table/template_space.go
type TemplateSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *TemplateSpaceSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *TemplateSpaceAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision           `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *TemplateSpace) Reset() {
	*x = TemplateSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSpace) ProtoMessage() {}

func (x *TemplateSpace) ProtoReflect() protoreflect.Message {
	mi := &file_template_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSpace.ProtoReflect.Descriptor instead.
func (*TemplateSpace) Descriptor() ([]byte, []int) {
	return file_template_space_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateSpace) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateSpace) GetSpec() *TemplateSpaceSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TemplateSpace) GetAttachment() *TemplateSpaceAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *TemplateSpace) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// TemplateSpaceSpec source resource reference: pkg/dal/table/template_space.go
type TemplateSpaceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *TemplateSpaceSpec) Reset() {
	*x = TemplateSpaceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSpaceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSpaceSpec) ProtoMessage() {}

func (x *TemplateSpaceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_template_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSpaceSpec.ProtoReflect.Descriptor instead.
func (*TemplateSpaceSpec) Descriptor() ([]byte, []int) {
	return file_template_space_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateSpaceSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateSpaceSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// TemplateSpaceAttachment source resource reference: pkg/dal/table/template_space.go
type TemplateSpaceAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
}

func (x *TemplateSpaceAttachment) Reset() {
	*x = TemplateSpaceAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_space_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSpaceAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSpaceAttachment) ProtoMessage() {}

func (x *TemplateSpaceAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_template_space_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSpaceAttachment.ProtoReflect.Descriptor instead.
func (*TemplateSpaceAttachment) Descriptor() ([]byte, []int) {
	return file_template_space_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateSpaceAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

var File_template_space_proto protoreflect.FileDescriptor

var file_template_space_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x74, 0x73, 0x1a, 0x21, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb9, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x62, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3d,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x30, 0x0a, 0x17, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x62, 0x73,
	0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3b, 0x70, 0x62, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_template_space_proto_rawDescOnce sync.Once
	file_template_space_proto_rawDescData = file_template_space_proto_rawDesc
)

func file_template_space_proto_rawDescGZIP() []byte {
	file_template_space_proto_rawDescOnce.Do(func() {
		file_template_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_space_proto_rawDescData)
	})
	return file_template_space_proto_rawDescData
}

var file_template_space_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_template_space_proto_goTypes = []interface{}{
	(*TemplateSpace)(nil),           // 0: pbts.TemplateSpace
	(*TemplateSpaceSpec)(nil),       // 1: pbts.TemplateSpaceSpec
	(*TemplateSpaceAttachment)(nil), // 2: pbts.TemplateSpaceAttachment
	(*base.Revision)(nil),           // 3: pbbase.Revision
}
var file_template_space_proto_depIdxs = []int32{
	1, // 0: pbts.TemplateSpace.spec:type_name -> pbts.TemplateSpaceSpec
	2, // 1: pbts.TemplateSpace.attachment:type_name -> pbts.TemplateSpaceAttachment
	3, // 2: pbts.TemplateSpace.revision:type_name -> pbbase.Revision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_template_space_proto_init() }
func file_template_space_proto_init() {
	if File_template_space_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSpace); i {
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
		file_template_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSpaceSpec); i {
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
		file_template_space_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSpaceAttachment); i {
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
			RawDescriptor: file_template_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_space_proto_goTypes,
		DependencyIndexes: file_template_space_proto_depIdxs,
		MessageInfos:      file_template_space_proto_msgTypes,
	}.Build()
	File_template_space_proto = out.File
	file_template_space_proto_rawDesc = nil
	file_template_space_proto_goTypes = nil
	file_template_space_proto_depIdxs = nil
}
