// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: operate.proto

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

type OperatorModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" form:"id"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @inject_tag: json:"title" form:"title"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" form:"title"`
	// @inject_tag: json:"content" form:"content"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" form:"content"`
	// @inject_tag: json:"url" form:"url"
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url" form:"url"`
}

func (x *OperatorModel) Reset() {
	*x = OperatorModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorModel) ProtoMessage() {}

func (x *OperatorModel) ProtoReflect() protoreflect.Message {
	mi := &file_operate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorModel.ProtoReflect.Descriptor instead.
func (*OperatorModel) Descriptor() ([]byte, []int) {
	return file_operate_proto_rawDescGZIP(), []int{0}
}

func (x *OperatorModel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OperatorModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OperatorModel) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *OperatorModel) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" form:"id"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @inject_tag: json:"title" form:"title"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" form:"title"`
	// @inject_tag: json:"content" form:"content"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" form:"content"`
	// @inject_tag: json:"status" form:"status"
	Status int64 `protobuf:"varint,4,opt,name=status,proto3" json:"status" form:"status"`
	// @inject_tag: json:"url" form:"url"
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url" form:"url"`
}

func (x *TaskReq) Reset() {
	*x = TaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReq) ProtoMessage() {}

func (x *TaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReq.ProtoReflect.Descriptor instead.
func (*TaskReq) Descriptor() ([]byte, []int) {
	return file_operate_proto_rawDescGZIP(), []int{1}
}

func (x *TaskReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TaskOperatorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"task_operatoe"
	TaskOperatoe *OperatorModel `protobuf:"bytes,1,opt,name=task_operatoe,json=taskOperatoe,proto3" json:"task_operatoe"`
	// @inject_tag: json:"code"
	Code uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"taskid"
	Taskid uint32 `protobuf:"varint,3,opt,name=taskid,proto3" json:"taskid"`
}

func (x *TaskOperatorInfo) Reset() {
	*x = TaskOperatorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOperatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOperatorInfo) ProtoMessage() {}

func (x *TaskOperatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_operate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOperatorInfo.ProtoReflect.Descriptor instead.
func (*TaskOperatorInfo) Descriptor() ([]byte, []int) {
	return file_operate_proto_rawDescGZIP(), []int{2}
}

func (x *TaskOperatorInfo) GetTaskOperatoe() *OperatorModel {
	if x != nil {
		return x.TaskOperatoe
	}
	return nil
}

func (x *TaskOperatorInfo) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TaskOperatorInfo) GetTaskid() uint32 {
	if x != nil {
		return x.Taskid
	}
	return 0
}

var File_operate_proto protoreflect.FileDescriptor

var file_operate_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x0d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x73, 0x0a, 0x07,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x7c, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x32,
	0x56, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operate_proto_rawDescOnce sync.Once
	file_operate_proto_rawDescData = file_operate_proto_rawDesc
)

func file_operate_proto_rawDescGZIP() []byte {
	file_operate_proto_rawDescOnce.Do(func() {
		file_operate_proto_rawDescData = protoimpl.X.CompressGZIP(file_operate_proto_rawDescData)
	})
	return file_operate_proto_rawDescData
}

var file_operate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_operate_proto_goTypes = []any{
	(*OperatorModel)(nil),    // 0: services.OperatorModel
	(*TaskReq)(nil),          // 1: services.TaskReq
	(*TaskOperatorInfo)(nil), // 2: services.TaskOperatorInfo
}
var file_operate_proto_depIdxs = []int32{
	0, // 0: services.TaskOperatorInfo.task_operatoe:type_name -> services.OperatorModel
	1, // 1: services.OperatorService.CreateTaskOperator:input_type -> services.TaskReq
	2, // 2: services.OperatorService.CreateTaskOperator:output_type -> services.TaskOperatorInfo
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_operate_proto_init() }
func file_operate_proto_init() {
	if File_operate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OperatorModel); i {
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
		file_operate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TaskReq); i {
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
		file_operate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TaskOperatorInfo); i {
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
			RawDescriptor: file_operate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operate_proto_goTypes,
		DependencyIndexes: file_operate_proto_depIdxs,
		MessageInfos:      file_operate_proto_msgTypes,
	}.Build()
	File_operate_proto = out.File
	file_operate_proto_rawDesc = nil
	file_operate_proto_goTypes = nil
	file_operate_proto_depIdxs = nil
}