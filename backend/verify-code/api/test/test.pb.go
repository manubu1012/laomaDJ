// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/test/test.proto

package test

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

type CreateTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTestRequest) Reset() {
	*x = CreateTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestRequest) ProtoMessage() {}

func (x *CreateTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestRequest.ProtoReflect.Descriptor instead.
func (*CreateTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{0}
}

type CreateTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTestReply) Reset() {
	*x = CreateTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestReply) ProtoMessage() {}

func (x *CreateTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestReply.ProtoReflect.Descriptor instead.
func (*CreateTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{1}
}

type UpdateTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTestRequest) Reset() {
	*x = UpdateTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestRequest) ProtoMessage() {}

func (x *UpdateTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestRequest.ProtoReflect.Descriptor instead.
func (*UpdateTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{2}
}

type UpdateTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTestReply) Reset() {
	*x = UpdateTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestReply) ProtoMessage() {}

func (x *UpdateTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestReply.ProtoReflect.Descriptor instead.
func (*UpdateTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{3}
}

type DeleteTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTestRequest) Reset() {
	*x = DeleteTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestRequest) ProtoMessage() {}

func (x *DeleteTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestRequest.ProtoReflect.Descriptor instead.
func (*DeleteTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{4}
}

type DeleteTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTestReply) Reset() {
	*x = DeleteTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestReply) ProtoMessage() {}

func (x *DeleteTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestReply.ProtoReflect.Descriptor instead.
func (*DeleteTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{5}
}

type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{6}
}

type GetTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTestReply) Reset() {
	*x = GetTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestReply) ProtoMessage() {}

func (x *GetTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestReply.ProtoReflect.Descriptor instead.
func (*GetTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{7}
}

type ListTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTestRequest) Reset() {
	*x = ListTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestRequest) ProtoMessage() {}

func (x *ListTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestRequest.ProtoReflect.Descriptor instead.
func (*ListTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{8}
}

type ListTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTestReply) Reset() {
	*x = ListTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestReply) ProtoMessage() {}

func (x *ListTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestReply.ProtoReflect.Descriptor instead.
func (*ListTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_test_proto_rawDescGZIP(), []int{9}
}

var File_api_test_test_proto protoreflect.FileDescriptor

var file_api_test_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22,
	0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xd5, 0x02,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x27, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x50, 0x01, 0x5a, 0x19, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2d, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_test_test_proto_rawDescOnce sync.Once
	file_api_test_test_proto_rawDescData = file_api_test_test_proto_rawDesc
)

func file_api_test_test_proto_rawDescGZIP() []byte {
	file_api_test_test_proto_rawDescOnce.Do(func() {
		file_api_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_test_test_proto_rawDescData)
	})
	return file_api_test_test_proto_rawDescData
}

var file_api_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_test_test_proto_goTypes = []interface{}{
	(*CreateTestRequest)(nil), // 0: api.test.CreateTestRequest
	(*CreateTestReply)(nil),   // 1: api.test.CreateTestReply
	(*UpdateTestRequest)(nil), // 2: api.test.UpdateTestRequest
	(*UpdateTestReply)(nil),   // 3: api.test.UpdateTestReply
	(*DeleteTestRequest)(nil), // 4: api.test.DeleteTestRequest
	(*DeleteTestReply)(nil),   // 5: api.test.DeleteTestReply
	(*GetTestRequest)(nil),    // 6: api.test.GetTestRequest
	(*GetTestReply)(nil),      // 7: api.test.GetTestReply
	(*ListTestRequest)(nil),   // 8: api.test.ListTestRequest
	(*ListTestReply)(nil),     // 9: api.test.ListTestReply
}
var file_api_test_test_proto_depIdxs = []int32{
	0, // 0: api.test.Test.CreateTest:input_type -> api.test.CreateTestRequest
	2, // 1: api.test.Test.UpdateTest:input_type -> api.test.UpdateTestRequest
	4, // 2: api.test.Test.DeleteTest:input_type -> api.test.DeleteTestRequest
	6, // 3: api.test.Test.GetTest:input_type -> api.test.GetTestRequest
	8, // 4: api.test.Test.ListTest:input_type -> api.test.ListTestRequest
	1, // 5: api.test.Test.CreateTest:output_type -> api.test.CreateTestReply
	3, // 6: api.test.Test.UpdateTest:output_type -> api.test.UpdateTestReply
	5, // 7: api.test.Test.DeleteTest:output_type -> api.test.DeleteTestReply
	7, // 8: api.test.Test.GetTest:output_type -> api.test.GetTestReply
	9, // 9: api.test.Test.ListTest:output_type -> api.test.ListTestReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_test_test_proto_init() }
func file_api_test_test_proto_init() {
	if File_api_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestRequest); i {
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
		file_api_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestReply); i {
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
		file_api_test_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestRequest); i {
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
		file_api_test_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestReply); i {
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
		file_api_test_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestRequest); i {
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
		file_api_test_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestReply); i {
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
		file_api_test_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
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
		file_api_test_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestReply); i {
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
		file_api_test_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestRequest); i {
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
		file_api_test_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestReply); i {
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
			RawDescriptor: file_api_test_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_test_test_proto_goTypes,
		DependencyIndexes: file_api_test_test_proto_depIdxs,
		MessageInfos:      file_api_test_test_proto_msgTypes,
	}.Build()
	File_api_test_test_proto = out.File
	file_api_test_test_proto_rawDesc = nil
	file_api_test_test_proto_goTypes = nil
	file_api_test_test_proto_depIdxs = nil
}