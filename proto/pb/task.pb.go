// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: task.proto

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

type TaskModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" form:"id"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @inject_tag: json:"title" form:"title"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" form:"title"`
	// @inject_tag: json:"content" form:"content"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" form:"content"`
	// @inject_tag: json:"start_time" form:"start_time"
	StartTime int64 `protobuf:"varint,4,opt,name=start_time,json=startTime,proto3" json:"start_time" form:"start_time"`
	// @inject_tag: json:"end_time" form:"end_time"
	EndTime int64 `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time" form:"end_time"`
	// @inject_tag: json:"status" form:"status"
	Status int64 `protobuf:"varint,6,opt,name=status,proto3" json:"status" form:"status"`
	// @inject_tag: json:"create_time" form:"create_time"
	CreateTime int64 `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time" form:"create_time"`
	// @inject_tag: json:"update_time" form:"update_time"
	UpdateTime int64 `protobuf:"varint,8,opt,name=update_time,json=updateTime,proto3" json:"update_time" form:"update_time"`
	// @inject_tag: json:"url" form:"url"
	Url string `protobuf:"bytes,9,opt,name=url,proto3" json:"url" form:"url"`
}

func (x *TaskModel) Reset() {
	*x = TaskModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskModel) ProtoMessage() {}

func (x *TaskModel) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskModel.ProtoReflect.Descriptor instead.
func (*TaskModel) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskModel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskModel) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskModel) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TaskModel) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *TaskModel) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskModel) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *TaskModel) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *TaskModel) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" form:"id"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @inject_tag: json:"title" form:"title"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" form:"title"`
	// @inject_tag: json:"content" form:"content"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" form:"content"`
	// @inject_tag: json:"start_time" form:"start_time"
	StartTime int64 `protobuf:"varint,4,opt,name=start_time,json=startTime,proto3" json:"start_time" form:"start_time"`
	// @inject_tag: json:"end_time" form:"end_time"
	EndTime int64 `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time" form:"end_time"`
	// @inject_tag: json:"status" form:"status"
	Status int64 `protobuf:"varint,6,opt,name=status,proto3" json:"status" form:"status"`
	// @inject_tag: json:"start" form:"start"
	Start uint32 `protobuf:"varint,7,opt,name=start,proto3" json:"start" form:"start"`
	// @inject_tag: json:"limit" form:"limit"
	Limit uint32 `protobuf:"varint,8,opt,name=limit,proto3" json:"limit" form:"limit"`
	// @inject_tag: json:"url" form:"url"
	Url string `protobuf:"bytes,9,opt,name=url,proto3" json:"url" form:"url"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TaskRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *TaskRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TaskRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"task_list"
	TaskList []*TaskModel `protobuf:"bytes,1,rep,name=task_list,json=taskList,proto3" json:"task_list"`
	// @inject_tag: json:"count"
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
	// @inject_tag: json:"code"
	Code uint32 `protobuf:"varint,3,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"uip"
	Uip string `protobuf:"bytes,4,opt,name=uip,proto3" json:"uip"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskListResponse) GetTaskList() []*TaskModel {
	if x != nil {
		return x.TaskList
	}
	return nil
}

func (x *TaskListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *TaskListResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TaskListResponse) GetUip() string {
	if x != nil {
		return x.Uip
	}
	return ""
}

type TaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"task_detail"
	TaskDetail *TaskModel `protobuf:"bytes,1,opt,name=task_detail,json=taskDetail,proto3" json:"task_detail"`
	// @inject_tag: json:"code"
	Code uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code"`
}

func (x *TaskDetailResponse) Reset() {
	*x = TaskDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetailResponse) ProtoMessage() {}

func (x *TaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetailResponse.ProtoReflect.Descriptor instead.
func (*TaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskDetailResponse) GetTaskDetail() *TaskModel {
	if x != nil {
		return x.TaskDetail
	}
	return nil
}

func (x *TaskDetailResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x70, 0x22, 0x5e, 0x0a,
	0x12, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0a, 0x74,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xd9, 0x02,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_task_proto_goTypes = []any{
	(*TaskModel)(nil),          // 0: services.TaskModel
	(*TaskRequest)(nil),        // 1: services.TaskRequest
	(*TaskListResponse)(nil),   // 2: services.TaskListResponse
	(*TaskDetailResponse)(nil), // 3: services.TaskDetailResponse
}
var file_task_proto_depIdxs = []int32{
	0, // 0: services.TaskListResponse.task_list:type_name -> services.TaskModel
	0, // 1: services.TaskDetailResponse.task_detail:type_name -> services.TaskModel
	1, // 2: services.TaskService.CreateTask:input_type -> services.TaskRequest
	1, // 3: services.TaskService.GetTasksList:input_type -> services.TaskRequest
	1, // 4: services.TaskService.GetTask:input_type -> services.TaskRequest
	1, // 5: services.TaskService.UpdateTask:input_type -> services.TaskRequest
	1, // 6: services.TaskService.DeleteTask:input_type -> services.TaskRequest
	3, // 7: services.TaskService.CreateTask:output_type -> services.TaskDetailResponse
	2, // 8: services.TaskService.GetTasksList:output_type -> services.TaskListResponse
	3, // 9: services.TaskService.GetTask:output_type -> services.TaskDetailResponse
	3, // 10: services.TaskService.UpdateTask:output_type -> services.TaskDetailResponse
	3, // 11: services.TaskService.DeleteTask:output_type -> services.TaskDetailResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TaskModel); i {
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
		file_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRequest); i {
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
		file_task_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TaskListResponse); i {
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
		file_task_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TaskDetailResponse); i {
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
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
