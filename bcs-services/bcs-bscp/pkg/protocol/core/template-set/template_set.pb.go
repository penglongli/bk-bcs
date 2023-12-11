// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: template_set.proto

package pbtset

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

// TemplateSet source resource reference: pkg/dal/table/template_set.go
type TemplateSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *TemplateSetSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *TemplateSetAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision         `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *TemplateSet) Reset() {
	*x = TemplateSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSet) ProtoMessage() {}

func (x *TemplateSet) ProtoReflect() protoreflect.Message {
	mi := &file_template_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSet.ProtoReflect.Descriptor instead.
func (*TemplateSet) Descriptor() ([]byte, []int) {
	return file_template_set_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateSet) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateSet) GetSpec() *TemplateSetSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TemplateSet) GetAttachment() *TemplateSetAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *TemplateSet) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// TemplateSetSpec source resource reference: pkg/dal/table/template_set.go
type TemplateSetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Memo        string   `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	TemplateIds []uint32 `protobuf:"varint,3,rep,packed,name=template_ids,json=templateIds,proto3" json:"template_ids,omitempty"`
	Public      bool     `protobuf:"varint,4,opt,name=public,proto3" json:"public,omitempty"`
	BoundApps   []uint32 `protobuf:"varint,5,rep,packed,name=bound_apps,json=boundApps,proto3" json:"bound_apps,omitempty"`
}

func (x *TemplateSetSpec) Reset() {
	*x = TemplateSetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSetSpec) ProtoMessage() {}

func (x *TemplateSetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_template_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSetSpec.ProtoReflect.Descriptor instead.
func (*TemplateSetSpec) Descriptor() ([]byte, []int) {
	return file_template_set_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateSetSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateSetSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TemplateSetSpec) GetTemplateIds() []uint32 {
	if x != nil {
		return x.TemplateIds
	}
	return nil
}

func (x *TemplateSetSpec) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *TemplateSetSpec) GetBoundApps() []uint32 {
	if x != nil {
		return x.BoundApps
	}
	return nil
}

// TemplateSetAttachment source resource reference: pkg/dal/table/template_set.go
type TemplateSetAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId           uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	TemplateSpaceId uint32 `protobuf:"varint,2,opt,name=template_space_id,json=templateSpaceId,proto3" json:"template_space_id,omitempty"`
}

func (x *TemplateSetAttachment) Reset() {
	*x = TemplateSetAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSetAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSetAttachment) ProtoMessage() {}

func (x *TemplateSetAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_template_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSetAttachment.ProtoReflect.Descriptor instead.
func (*TemplateSetAttachment) Descriptor() ([]byte, []int) {
	return file_template_set_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateSetAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *TemplateSetAttachment) GetTemplateSpaceId() uint32 {
	if x != nil {
		return x.TemplateSpaceId
	}
	return 0
}

// TemplateSetOfBizDetail is template set of biz detail
type TemplateSetOfBizDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateSpaceId   uint32                                     `protobuf:"varint,1,opt,name=template_space_id,json=templateSpaceId,proto3" json:"template_space_id,omitempty"`
	TemplateSpaceName string                                     `protobuf:"bytes,2,opt,name=template_space_name,json=templateSpaceName,proto3" json:"template_space_name,omitempty"`
	TemplateSets      []*TemplateSetOfBizDetail_TemplateSetOfBiz `protobuf:"bytes,3,rep,name=template_sets,json=templateSets,proto3" json:"template_sets,omitempty"`
}

func (x *TemplateSetOfBizDetail) Reset() {
	*x = TemplateSetOfBizDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSetOfBizDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSetOfBizDetail) ProtoMessage() {}

func (x *TemplateSetOfBizDetail) ProtoReflect() protoreflect.Message {
	mi := &file_template_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSetOfBizDetail.ProtoReflect.Descriptor instead.
func (*TemplateSetOfBizDetail) Descriptor() ([]byte, []int) {
	return file_template_set_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateSetOfBizDetail) GetTemplateSpaceId() uint32 {
	if x != nil {
		return x.TemplateSpaceId
	}
	return 0
}

func (x *TemplateSetOfBizDetail) GetTemplateSpaceName() string {
	if x != nil {
		return x.TemplateSpaceName
	}
	return ""
}

func (x *TemplateSetOfBizDetail) GetTemplateSets() []*TemplateSetOfBizDetail_TemplateSetOfBiz {
	if x != nil {
		return x.TemplateSets
	}
	return nil
}

type TemplateSetBriefInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateSpaceId   uint32 `protobuf:"varint,1,opt,name=template_space_id,json=templateSpaceId,proto3" json:"template_space_id,omitempty"`
	TemplateSpaceName string `protobuf:"bytes,2,opt,name=template_space_name,json=templateSpaceName,proto3" json:"template_space_name,omitempty"`
	TemplateSetId     uint32 `protobuf:"varint,3,opt,name=template_set_id,json=templateSetId,proto3" json:"template_set_id,omitempty"`
	TemplateSetName   string `protobuf:"bytes,4,opt,name=template_set_name,json=templateSetName,proto3" json:"template_set_name,omitempty"`
}

func (x *TemplateSetBriefInfo) Reset() {
	*x = TemplateSetBriefInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSetBriefInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSetBriefInfo) ProtoMessage() {}

func (x *TemplateSetBriefInfo) ProtoReflect() protoreflect.Message {
	mi := &file_template_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSetBriefInfo.ProtoReflect.Descriptor instead.
func (*TemplateSetBriefInfo) Descriptor() ([]byte, []int) {
	return file_template_set_proto_rawDescGZIP(), []int{4}
}

func (x *TemplateSetBriefInfo) GetTemplateSpaceId() uint32 {
	if x != nil {
		return x.TemplateSpaceId
	}
	return 0
}

func (x *TemplateSetBriefInfo) GetTemplateSpaceName() string {
	if x != nil {
		return x.TemplateSpaceName
	}
	return ""
}

func (x *TemplateSetBriefInfo) GetTemplateSetId() uint32 {
	if x != nil {
		return x.TemplateSetId
	}
	return 0
}

func (x *TemplateSetBriefInfo) GetTemplateSetName() string {
	if x != nil {
		return x.TemplateSetName
	}
	return ""
}

type TemplateSetOfBizDetail_TemplateSetOfBiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateSetId   uint32   `protobuf:"varint,1,opt,name=template_set_id,json=templateSetId,proto3" json:"template_set_id,omitempty"`
	TemplateSetName string   `protobuf:"bytes,2,opt,name=template_set_name,json=templateSetName,proto3" json:"template_set_name,omitempty"`
	TemplateIds     []uint32 `protobuf:"varint,3,rep,packed,name=template_ids,json=templateIds,proto3" json:"template_ids,omitempty"`
}

func (x *TemplateSetOfBizDetail_TemplateSetOfBiz) Reset() {
	*x = TemplateSetOfBizDetail_TemplateSetOfBiz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSetOfBizDetail_TemplateSetOfBiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSetOfBizDetail_TemplateSetOfBiz) ProtoMessage() {}

func (x *TemplateSetOfBizDetail_TemplateSetOfBiz) ProtoReflect() protoreflect.Message {
	mi := &file_template_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSetOfBizDetail_TemplateSetOfBiz.ProtoReflect.Descriptor instead.
func (*TemplateSetOfBizDetail_TemplateSetOfBiz) Descriptor() ([]byte, []int) {
	return file_template_set_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TemplateSetOfBizDetail_TemplateSetOfBiz) GetTemplateSetId() uint32 {
	if x != nil {
		return x.TemplateSetId
	}
	return 0
}

func (x *TemplateSetOfBizDetail_TemplateSetOfBiz) GetTemplateSetName() string {
	if x != nil {
		return x.TemplateSetName
	}
	return ""
}

func (x *TemplateSetOfBizDetail_TemplateSetOfBiz) GetTemplateIds() []uint32 {
	if x != nil {
		return x.TemplateIds
	}
	return nil
}

var File_template_set_proto protoreflect.FileDescriptor

var file_template_set_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x62, 0x74, 0x73, 0x65, 0x74, 0x1a, 0x21, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb7, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x62, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x62, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x73, 0x22,
	0x5a, 0x0a, 0x15, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xd6, 0x02, 0x0a, 0x16,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4f, 0x66, 0x42, 0x69, 0x7a,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x62, 0x74, 0x73,
	0x65, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4f, 0x66,
	0x42, 0x69, 0x7a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x4f, 0x66, 0x42, 0x69, 0x7a, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x73, 0x1a, 0x89, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4f, 0x66, 0x42, 0x69, 0x7a, 0x12, 0x26, 0x0a,
	0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x14, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a,
	0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x2f, 0x5a,
	0x2d, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x74, 0x3b, 0x70, 0x62, 0x74, 0x73, 0x65, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_set_proto_rawDescOnce sync.Once
	file_template_set_proto_rawDescData = file_template_set_proto_rawDesc
)

func file_template_set_proto_rawDescGZIP() []byte {
	file_template_set_proto_rawDescOnce.Do(func() {
		file_template_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_set_proto_rawDescData)
	})
	return file_template_set_proto_rawDescData
}

var file_template_set_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_template_set_proto_goTypes = []interface{}{
	(*TemplateSet)(nil),                             // 0: pbtset.TemplateSet
	(*TemplateSetSpec)(nil),                         // 1: pbtset.TemplateSetSpec
	(*TemplateSetAttachment)(nil),                   // 2: pbtset.TemplateSetAttachment
	(*TemplateSetOfBizDetail)(nil),                  // 3: pbtset.TemplateSetOfBizDetail
	(*TemplateSetBriefInfo)(nil),                    // 4: pbtset.TemplateSetBriefInfo
	(*TemplateSetOfBizDetail_TemplateSetOfBiz)(nil), // 5: pbtset.TemplateSetOfBizDetail.TemplateSetOfBiz
	(*base.Revision)(nil),                           // 6: pbbase.Revision
}
var file_template_set_proto_depIdxs = []int32{
	1, // 0: pbtset.TemplateSet.spec:type_name -> pbtset.TemplateSetSpec
	2, // 1: pbtset.TemplateSet.attachment:type_name -> pbtset.TemplateSetAttachment
	6, // 2: pbtset.TemplateSet.revision:type_name -> pbbase.Revision
	5, // 3: pbtset.TemplateSetOfBizDetail.template_sets:type_name -> pbtset.TemplateSetOfBizDetail.TemplateSetOfBiz
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_template_set_proto_init() }
func file_template_set_proto_init() {
	if File_template_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSet); i {
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
		file_template_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSetSpec); i {
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
		file_template_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSetAttachment); i {
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
		file_template_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSetOfBizDetail); i {
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
		file_template_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSetBriefInfo); i {
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
		file_template_set_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSetOfBizDetail_TemplateSetOfBiz); i {
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
			RawDescriptor: file_template_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_set_proto_goTypes,
		DependencyIndexes: file_template_set_proto_depIdxs,
		MessageInfos:      file_template_set_proto_msgTypes,
	}.Build()
	File_template_set_proto = out.File
	file_template_set_proto_rawDesc = nil
	file_template_set_proto_goTypes = nil
	file_template_set_proto_depIdxs = nil
}
