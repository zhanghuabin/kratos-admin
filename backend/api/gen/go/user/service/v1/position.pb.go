// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: user/service/v1/position.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 职位
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint32                `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`                                    // 职位ID
	Name       *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`                                 // 职位名称
	OrderNo    *int32                 `protobuf:"varint,3,opt,name=order_no,json=orderNo,proto3,oneof" json:"order_no,omitempty"`           // 排序号
	Code       *string                `protobuf:"bytes,4,opt,name=code,proto3,oneof" json:"code,omitempty"`                                 // 职位值
	Status     *string                `protobuf:"bytes,5,opt,name=status,proto3,oneof" json:"status,omitempty"`                             // 状态
	Remark     *string                `protobuf:"bytes,6,opt,name=remark,proto3,oneof" json:"remark,omitempty"`                             // 备注
	ParentId   *uint32                `protobuf:"varint,100,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`      // 父节点ID
	Children   []*Position            `protobuf:"bytes,101,rep,name=children,proto3" json:"children,omitempty"`                             // 子节点树
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,200,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"` // 创建时间
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,201,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"` // 更新时间
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,202,opt,name=delete_time,json=deleteTime,proto3,oneof" json:"delete_time,omitempty"` // 删除时间
}

func (x *Position) Reset() {
	*x = Position{}
	mi := &file_user_service_v1_position_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_position_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_user_service_v1_position_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Position) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Position) GetOrderNo() int32 {
	if x != nil && x.OrderNo != nil {
		return *x.OrderNo
	}
	return 0
}

func (x *Position) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *Position) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Position) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

func (x *Position) GetParentId() uint32 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Position) GetChildren() []*Position {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Position) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Position) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Position) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

// 获取职位列表 - 答复
type ListPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Position `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListPositionResponse) Reset() {
	*x = ListPositionResponse{}
	mi := &file_user_service_v1_position_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPositionResponse) ProtoMessage() {}

func (x *ListPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_position_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPositionResponse.ProtoReflect.Descriptor instead.
func (*ListPositionResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_position_proto_rawDescGZIP(), []int{1}
}

func (x *ListPositionResponse) GetItems() []*Position {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPositionResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 获取职位数据 - 请求
type GetPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPositionRequest) Reset() {
	*x = GetPositionRequest{}
	mi := &file_user_service_v1_position_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionRequest) ProtoMessage() {}

func (x *GetPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_position_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionRequest.ProtoReflect.Descriptor instead.
func (*GetPositionRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_position_proto_rawDescGZIP(), []int{2}
}

func (x *GetPositionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 创建职位 - 请求
type CreatePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId *uint32   `protobuf:"varint,1,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"` // 操作用户ID
	Position   *Position `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *CreatePositionRequest) Reset() {
	*x = CreatePositionRequest{}
	mi := &file_user_service_v1_position_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePositionRequest) ProtoMessage() {}

func (x *CreatePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_position_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePositionRequest.ProtoReflect.Descriptor instead.
func (*CreatePositionRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_position_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePositionRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

func (x *CreatePositionRequest) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

// 更新职位 - 请求
type UpdatePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId   *uint32                `protobuf:"varint,1,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"` // 操作用户ID
	Position     *Position              `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	UpdateMask   *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`              // 要更新的字段列表
	AllowMissing *bool                  `protobuf:"varint,4,opt,name=allow_missing,json=allowMissing,proto3,oneof" json:"allow_missing,omitempty"` // 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。
}

func (x *UpdatePositionRequest) Reset() {
	*x = UpdatePositionRequest{}
	mi := &file_user_service_v1_position_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePositionRequest) ProtoMessage() {}

func (x *UpdatePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_position_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePositionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePositionRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_position_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePositionRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

func (x *UpdatePositionRequest) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *UpdatePositionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdatePositionRequest) GetAllowMissing() bool {
	if x != nil && x.AllowMissing != nil {
		return *x.AllowMissing
	}
	return false
}

// 删除职位 - 请求
type DeletePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId *uint32 `protobuf:"varint,1,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"` // 操作用户ID
	Id         uint32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePositionRequest) Reset() {
	*x = DeletePositionRequest{}
	mi := &file_user_service_v1_position_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePositionRequest) ProtoMessage() {}

func (x *DeletePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_position_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePositionRequest.ProtoReflect.Descriptor instead.
func (*DeletePositionRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_position_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePositionRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

func (x *DeletePositionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_user_service_v1_position_proto protoreflect.FileDescriptor

var file_user_service_v1_position_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x06, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x0e, 0xba, 0x47, 0x0b, 0x92, 0x02, 0x08, 0xe8, 0x81, 0x8c, 0xe4, 0xbd, 0x8d, 0x49, 0x44,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe8,
	0x81, 0x8c, 0xe4, 0xbd, 0x8d, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0x48, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xba, 0x47, 0x0c, 0x92, 0x02, 0x09,
	0xe6, 0x8e, 0x92, 0xe5, 0xba, 0x8f, 0xe5, 0x8f, 0xb7, 0x48, 0x02, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xba, 0x47, 0x0c, 0x92, 0x02, 0x09, 0xe8, 0x81, 0x8c,
	0xe4, 0xbd, 0x8d, 0xe5, 0x80, 0xbc, 0x48, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x22, 0xba, 0x47, 0x1f, 0xc2, 0x01, 0x04, 0x12, 0x02, 0x4f, 0x4e, 0xc2, 0x01, 0x05,
	0x12, 0x03, 0x4f, 0x46, 0x46, 0x8a, 0x02, 0x04, 0x1a, 0x02, 0x4f, 0x4e, 0x92, 0x02, 0x06, 0xe7,
	0x8a, 0xb6, 0xe6, 0x80, 0x81, 0x48, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xba, 0x47, 0x09, 0x92, 0x02, 0x06, 0xe5, 0xa4, 0x87, 0xe6, 0xb3, 0xa8,
	0x48, 0x05, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x11, 0xba, 0x47, 0x0e, 0x92, 0x02, 0x0b, 0xe7, 0x88, 0xb6, 0xe8, 0x8a, 0x82, 0xe7, 0x82,
	0xb9, 0x49, 0x44, 0x48, 0x06, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x49, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x65,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe5, 0xad, 0x90, 0xe8, 0x8a, 0x82, 0xe7, 0x82, 0xb9,
	0xe6, 0xa0, 0x91, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x55, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xc8, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe6, 0x97, 0xb6,
	0xe9, 0x97, 0xb4, 0x48, 0x07, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe6, 0x9b,
	0xb4, 0xe6, 0x96, 0xb0, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4, 0x48, 0x08, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xba,
	0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97,
	0xb4, 0x48, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x9c, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x16, 0xba, 0x47, 0x13, 0x18, 0x01, 0x92, 0x02, 0x0e, 0xe6, 0x93, 0x8d, 0xe4, 0xbd, 0x9c,
	0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x22, 0xda, 0x03, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x16, 0xba, 0x47, 0x13, 0x18, 0x01, 0x92, 0x02, 0x0e, 0xe6, 0x93, 0x8d, 0xe4, 0xbd, 0x9c, 0xe7,
	0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x73, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x36, 0xba, 0x47, 0x33, 0x3a, 0x16, 0x12, 0x14, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x92, 0x02, 0x18,
	0xe8, 0xa6, 0x81, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe7, 0x9a, 0x84, 0xe5, 0xad, 0x97, 0xe6,
	0xae, 0xb5, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x12, 0xb4, 0x01, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x89, 0x01, 0xba,
	0x47, 0x85, 0x01, 0x92, 0x02, 0x81, 0x01, 0xe5, 0xa6, 0x82, 0xe6, 0x9e, 0x9c, 0xe8, 0xae, 0xbe,
	0xe7, 0xbd, 0xae, 0xe4, 0xb8, 0xba, 0x74, 0x72, 0x75, 0x65, 0xe7, 0x9a, 0x84, 0xe6, 0x97, 0xb6,
	0xe5, 0x80, 0x99, 0xef, 0xbc, 0x8c, 0xe8, 0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe4, 0xb8, 0x8d, 0xe5,
	0xad, 0x98, 0xe5, 0x9c, 0xa8, 0xe5, 0x88, 0x99, 0xe4, 0xbc, 0x9a, 0xe6, 0x96, 0xb0, 0xe5, 0xa2,
	0x9e, 0x28, 0xe6, 0x8f, 0x92, 0xe5, 0x85, 0xa5, 0x29, 0xef, 0xbc, 0x8c, 0xe5, 0xb9, 0xb6, 0xe4,
	0xb8, 0x94, 0xe5, 0x9c, 0xa8, 0xe8, 0xbf, 0x99, 0xe7, 0xa7, 0x8d, 0xe6, 0x83, 0x85, 0xe5, 0x86,
	0xb5, 0xe4, 0xb8, 0x8b, 0x60, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x60,
	0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe5, 0xb0, 0x86, 0xe4, 0xbc, 0x9a, 0xe8, 0xa2, 0xab, 0xe5,
	0xbf, 0xbd, 0xe7, 0x95, 0xa5, 0xe3, 0x80, 0x82, 0x48, 0x01, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x75, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x16, 0xba, 0x47, 0x13,
	0x18, 0x01, 0x92, 0x02, 0x0e, 0xe6, 0x93, 0x8d, 0xe4, 0xbd, 0x9c, 0xe7, 0x94, 0xa8, 0xe6, 0x88,
	0xb7, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x32, 0xb2, 0x03, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xb5, 0x01, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x31, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_v1_position_proto_rawDescOnce sync.Once
	file_user_service_v1_position_proto_rawDescData = file_user_service_v1_position_proto_rawDesc
)

func file_user_service_v1_position_proto_rawDescGZIP() []byte {
	file_user_service_v1_position_proto_rawDescOnce.Do(func() {
		file_user_service_v1_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_v1_position_proto_rawDescData)
	})
	return file_user_service_v1_position_proto_rawDescData
}

var file_user_service_v1_position_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_service_v1_position_proto_goTypes = []any{
	(*Position)(nil),              // 0: user.service.v1.Position
	(*ListPositionResponse)(nil),  // 1: user.service.v1.ListPositionResponse
	(*GetPositionRequest)(nil),    // 2: user.service.v1.GetPositionRequest
	(*CreatePositionRequest)(nil), // 3: user.service.v1.CreatePositionRequest
	(*UpdatePositionRequest)(nil), // 4: user.service.v1.UpdatePositionRequest
	(*DeletePositionRequest)(nil), // 5: user.service.v1.DeletePositionRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*v1.PagingRequest)(nil),      // 8: pagination.PagingRequest
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_user_service_v1_position_proto_depIdxs = []int32{
	0,  // 0: user.service.v1.Position.children:type_name -> user.service.v1.Position
	6,  // 1: user.service.v1.Position.create_time:type_name -> google.protobuf.Timestamp
	6,  // 2: user.service.v1.Position.update_time:type_name -> google.protobuf.Timestamp
	6,  // 3: user.service.v1.Position.delete_time:type_name -> google.protobuf.Timestamp
	0,  // 4: user.service.v1.ListPositionResponse.items:type_name -> user.service.v1.Position
	0,  // 5: user.service.v1.CreatePositionRequest.position:type_name -> user.service.v1.Position
	0,  // 6: user.service.v1.UpdatePositionRequest.position:type_name -> user.service.v1.Position
	7,  // 7: user.service.v1.UpdatePositionRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 8: user.service.v1.PositionService.ListPosition:input_type -> pagination.PagingRequest
	2,  // 9: user.service.v1.PositionService.GetPosition:input_type -> user.service.v1.GetPositionRequest
	3,  // 10: user.service.v1.PositionService.CreatePosition:input_type -> user.service.v1.CreatePositionRequest
	4,  // 11: user.service.v1.PositionService.UpdatePosition:input_type -> user.service.v1.UpdatePositionRequest
	5,  // 12: user.service.v1.PositionService.DeletePosition:input_type -> user.service.v1.DeletePositionRequest
	1,  // 13: user.service.v1.PositionService.ListPosition:output_type -> user.service.v1.ListPositionResponse
	0,  // 14: user.service.v1.PositionService.GetPosition:output_type -> user.service.v1.Position
	9,  // 15: user.service.v1.PositionService.CreatePosition:output_type -> google.protobuf.Empty
	9,  // 16: user.service.v1.PositionService.UpdatePosition:output_type -> google.protobuf.Empty
	9,  // 17: user.service.v1.PositionService.DeletePosition:output_type -> google.protobuf.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_user_service_v1_position_proto_init() }
func file_user_service_v1_position_proto_init() {
	if File_user_service_v1_position_proto != nil {
		return
	}
	file_user_service_v1_position_proto_msgTypes[0].OneofWrappers = []any{}
	file_user_service_v1_position_proto_msgTypes[3].OneofWrappers = []any{}
	file_user_service_v1_position_proto_msgTypes[4].OneofWrappers = []any{}
	file_user_service_v1_position_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_v1_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_v1_position_proto_goTypes,
		DependencyIndexes: file_user_service_v1_position_proto_depIdxs,
		MessageInfos:      file_user_service_v1_position_proto_msgTypes,
	}.Build()
	File_user_service_v1_position_proto = out.File
	file_user_service_v1_position_proto_rawDesc = nil
	file_user_service_v1_position_proto_goTypes = nil
	file_user_service_v1_position_proto_depIdxs = nil
}
