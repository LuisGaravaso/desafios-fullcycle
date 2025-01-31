// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: internal/infra/grpc/protofiles/get_order.proto

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

type GetOrderByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderByIdRequest) Reset() {
	*x = GetOrderByIdRequest{}
	mi := &file_internal_infra_grpc_protofiles_get_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdRequest) ProtoMessage() {}

func (x *GetOrderByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_get_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdRequest.ProtoReflect.Descriptor instead.
func (*GetOrderByIdRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_get_order_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrderByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_internal_infra_grpc_protofiles_get_order_proto protoreflect.FileDescriptor

var file_internal_infra_grpc_protofiles_get_order_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x77, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x18, 0x5a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_infra_grpc_protofiles_get_order_proto_rawDescOnce sync.Once
	file_internal_infra_grpc_protofiles_get_order_proto_rawDescData = file_internal_infra_grpc_protofiles_get_order_proto_rawDesc
)

func file_internal_infra_grpc_protofiles_get_order_proto_rawDescGZIP() []byte {
	file_internal_infra_grpc_protofiles_get_order_proto_rawDescOnce.Do(func() {
		file_internal_infra_grpc_protofiles_get_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_infra_grpc_protofiles_get_order_proto_rawDescData)
	})
	return file_internal_infra_grpc_protofiles_get_order_proto_rawDescData
}

var file_internal_infra_grpc_protofiles_get_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_infra_grpc_protofiles_get_order_proto_goTypes = []any{
	(*GetOrderByIdRequest)(nil), // 0: pb.GetOrderByIdRequest
	(*Blank)(nil),               // 1: pb.Blank
	(*Order)(nil),               // 2: pb.Order
	(*GetOrdersResponse)(nil),   // 3: pb.GetOrdersResponse
}
var file_internal_infra_grpc_protofiles_get_order_proto_depIdxs = []int32{
	0, // 0: pb.GetOrderService.GetOrderById:input_type -> pb.GetOrderByIdRequest
	1, // 1: pb.GetOrderService.GetAllOrders:input_type -> pb.Blank
	2, // 2: pb.GetOrderService.GetOrderById:output_type -> pb.Order
	3, // 3: pb.GetOrderService.GetAllOrders:output_type -> pb.GetOrdersResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_infra_grpc_protofiles_get_order_proto_init() }
func file_internal_infra_grpc_protofiles_get_order_proto_init() {
	if File_internal_infra_grpc_protofiles_get_order_proto != nil {
		return
	}
	file_internal_infra_grpc_protofiles_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_infra_grpc_protofiles_get_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_infra_grpc_protofiles_get_order_proto_goTypes,
		DependencyIndexes: file_internal_infra_grpc_protofiles_get_order_proto_depIdxs,
		MessageInfos:      file_internal_infra_grpc_protofiles_get_order_proto_msgTypes,
	}.Build()
	File_internal_infra_grpc_protofiles_get_order_proto = out.File
	file_internal_infra_grpc_protofiles_get_order_proto_rawDesc = nil
	file_internal_infra_grpc_protofiles_get_order_proto_goTypes = nil
	file_internal_infra_grpc_protofiles_get_order_proto_depIdxs = nil
}
