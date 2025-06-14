// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: customerspb/messages.proto

package customerspb

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

type CustomerRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SmsNumber string `protobuf:"bytes,3,opt,name=sms_number,json=smsNumber,proto3" json:"sms_number,omitempty"`
}

func (x *CustomerRegistered) Reset() {
	*x = CustomerRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerspb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerRegistered) ProtoMessage() {}

func (x *CustomerRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_customerspb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerRegistered.ProtoReflect.Descriptor instead.
func (*CustomerRegistered) Descriptor() ([]byte, []int) {
	return file_customerspb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerRegistered) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomerRegistered) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerRegistered) GetSmsNumber() string {
	if x != nil {
		return x.SmsNumber
	}
	return ""
}

type CustomerSmsChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SmsNumber string `protobuf:"bytes,2,opt,name=sms_number,json=smsNumber,proto3" json:"sms_number,omitempty"`
}

func (x *CustomerSmsChanged) Reset() {
	*x = CustomerSmsChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerspb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerSmsChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerSmsChanged) ProtoMessage() {}

func (x *CustomerSmsChanged) ProtoReflect() protoreflect.Message {
	mi := &file_customerspb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerSmsChanged.ProtoReflect.Descriptor instead.
func (*CustomerSmsChanged) Descriptor() ([]byte, []int) {
	return file_customerspb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerSmsChanged) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomerSmsChanged) GetSmsNumber() string {
	if x != nil {
		return x.SmsNumber
	}
	return ""
}

type CustomerEnabled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerEnabled) Reset() {
	*x = CustomerEnabled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerspb_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerEnabled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerEnabled) ProtoMessage() {}

func (x *CustomerEnabled) ProtoReflect() protoreflect.Message {
	mi := &file_customerspb_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerEnabled.ProtoReflect.Descriptor instead.
func (*CustomerEnabled) Descriptor() ([]byte, []int) {
	return file_customerspb_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerEnabled) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CustomerDisabled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerDisabled) Reset() {
	*x = CustomerDisabled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerspb_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDisabled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDisabled) ProtoMessage() {}

func (x *CustomerDisabled) ProtoReflect() protoreflect.Message {
	mi := &file_customerspb_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDisabled.ProtoReflect.Descriptor instead.
func (*CustomerDisabled) Descriptor() ([]byte, []int) {
	return file_customerspb_messages_proto_rawDescGZIP(), []int{3}
}

func (x *CustomerDisabled) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AuthorizeCustomer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorizeCustomer) Reset() {
	*x = AuthorizeCustomer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerspb_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeCustomer) ProtoMessage() {}

func (x *AuthorizeCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_customerspb_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeCustomer.ProtoReflect.Descriptor instead.
func (*AuthorizeCustomer) Descriptor() ([]byte, []int) {
	return file_customerspb_messages_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorizeCustomer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_customerspb_messages_proto protoreflect.FileDescriptor

var file_customerspb_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x70, 0x62, 0x22, 0x57, 0x0a, 0x12, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x43, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6d,
	0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6d,
	0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23,
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x42, 0xb7, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x70, 0x62, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x65, 0x7a, 0x7a, 0x79, 0x2f, 0x73, 0x6f, 0x6b, 0x6f,
	0x2d, 0x62, 0x6f, 0x72, 0x61, 0x2d, 0x6d, 0x6e, 0x67, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x70, 0x62, 0xca, 0x02, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x70, 0x62, 0xe2, 0x02, 0x17, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customerspb_messages_proto_rawDescOnce sync.Once
	file_customerspb_messages_proto_rawDescData = file_customerspb_messages_proto_rawDesc
)

func file_customerspb_messages_proto_rawDescGZIP() []byte {
	file_customerspb_messages_proto_rawDescOnce.Do(func() {
		file_customerspb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_customerspb_messages_proto_rawDescData)
	})
	return file_customerspb_messages_proto_rawDescData
}

var file_customerspb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_customerspb_messages_proto_goTypes = []interface{}{
	(*CustomerRegistered)(nil), // 0: customerspb.CustomerRegistered
	(*CustomerSmsChanged)(nil), // 1: customerspb.CustomerSmsChanged
	(*CustomerEnabled)(nil),    // 2: customerspb.CustomerEnabled
	(*CustomerDisabled)(nil),   // 3: customerspb.CustomerDisabled
	(*AuthorizeCustomer)(nil),  // 4: customerspb.AuthorizeCustomer
}
var file_customerspb_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customerspb_messages_proto_init() }
func file_customerspb_messages_proto_init() {
	if File_customerspb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customerspb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerRegistered); i {
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
		file_customerspb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerSmsChanged); i {
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
		file_customerspb_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerEnabled); i {
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
		file_customerspb_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDisabled); i {
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
		file_customerspb_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeCustomer); i {
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
			RawDescriptor: file_customerspb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customerspb_messages_proto_goTypes,
		DependencyIndexes: file_customerspb_messages_proto_depIdxs,
		MessageInfos:      file_customerspb_messages_proto_msgTypes,
	}.Build()
	File_customerspb_messages_proto = out.File
	file_customerspb_messages_proto_rawDesc = nil
	file_customerspb_messages_proto_goTypes = nil
	file_customerspb_messages_proto_depIdxs = nil
}
