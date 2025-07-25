// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: paymentspb/messages.proto

package paymentspb

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

type InvoicePaid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *InvoicePaid) Reset() {
	*x = InvoicePaid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoicePaid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoicePaid) ProtoMessage() {}

func (x *InvoicePaid) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoicePaid.ProtoReflect.Descriptor instead.
func (*InvoicePaid) Descriptor() ([]byte, []int) {
	return file_paymentspb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *InvoicePaid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvoicePaid) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type ConfirmPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ConfirmPayment) Reset() {
	*x = ConfirmPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmPayment) ProtoMessage() {}

func (x *ConfirmPayment) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmPayment.ProtoReflect.Descriptor instead.
func (*ConfirmPayment) Descriptor() ([]byte, []int) {
	return file_paymentspb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmPayment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConfirmPayment) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_paymentspb_messages_proto protoreflect.FileDescriptor

var file_paymentspb_messages_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x22, 0x38, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xaf, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x42, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x65, 0x7a,
	0x7a, 0x79, 0x2f, 0x73, 0x6f, 0x6b, 0x6f, 0x2d, 0x62, 0x6f, 0x72, 0x61, 0x2d, 0x6d, 0x6e, 0x67,
	0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paymentspb_messages_proto_rawDescOnce sync.Once
	file_paymentspb_messages_proto_rawDescData = file_paymentspb_messages_proto_rawDesc
)

func file_paymentspb_messages_proto_rawDescGZIP() []byte {
	file_paymentspb_messages_proto_rawDescOnce.Do(func() {
		file_paymentspb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentspb_messages_proto_rawDescData)
	})
	return file_paymentspb_messages_proto_rawDescData
}

var file_paymentspb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_paymentspb_messages_proto_goTypes = []interface{}{
	(*InvoicePaid)(nil),    // 0: paymentspb.InvoicePaid
	(*ConfirmPayment)(nil), // 1: paymentspb.ConfirmPayment
}
var file_paymentspb_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paymentspb_messages_proto_init() }
func file_paymentspb_messages_proto_init() {
	if File_paymentspb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paymentspb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoicePaid); i {
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
		file_paymentspb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmPayment); i {
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
			RawDescriptor: file_paymentspb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paymentspb_messages_proto_goTypes,
		DependencyIndexes: file_paymentspb_messages_proto_depIdxs,
		MessageInfos:      file_paymentspb_messages_proto_msgTypes,
	}.Build()
	File_paymentspb_messages_proto = out.File
	file_paymentspb_messages_proto_rawDesc = nil
	file_paymentspb_messages_proto_goTypes = nil
	file_paymentspb_messages_proto_depIdxs = nil
}
