// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: client-errors.proto

package clienterrorspb

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

type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg              string                         `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Params           map[string]string              `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ValidationErrors *ErrorDetails_ValidationErrors `protobuf:"bytes,4,opt,name=validationErrors,proto3" json:"validationErrors,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_client_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_client_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ErrorDetails) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ErrorDetails) GetValidationErrors() *ErrorDetails_ValidationErrors {
	if x != nil {
		return x.ValidationErrors
	}
	return nil
}

type ErrorDetails_ValidationErrors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ErrorDetails_ValidationErrorItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ErrorDetails_ValidationErrors) Reset() {
	*x = ErrorDetails_ValidationErrors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_errors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails_ValidationErrors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails_ValidationErrors) ProtoMessage() {}

func (x *ErrorDetails_ValidationErrors) ProtoReflect() protoreflect.Message {
	mi := &file_client_errors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails_ValidationErrors.ProtoReflect.Descriptor instead.
func (*ErrorDetails_ValidationErrors) Descriptor() ([]byte, []int) {
	return file_client_errors_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ErrorDetails_ValidationErrors) GetItems() []*ErrorDetails_ValidationErrorItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ErrorDetails_ValidationErrorItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ErrorDetails_ValidationErrorItem) Reset() {
	*x = ErrorDetails_ValidationErrorItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_errors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails_ValidationErrorItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails_ValidationErrorItem) ProtoMessage() {}

func (x *ErrorDetails_ValidationErrorItem) ProtoReflect() protoreflect.Message {
	mi := &file_client_errors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails_ValidationErrorItem.ProtoReflect.Descriptor instead.
func (*ErrorDetails_ValidationErrorItem) Descriptor() ([]byte, []int) {
	return file_client_errors_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ErrorDetails_ValidationErrorItem) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ErrorDetails_ValidationErrorItem) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_client_errors_proto protoreflect.FileDescriptor

var file_client_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x90, 0x03, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x58, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x59, 0x0a,
	0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x45, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x3d, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x65, 0x6d, 0x6f, 0x66, 0x74,
	0x68, 0x65, 0x73, 0x6f, 0x75, 0x6c, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_errors_proto_rawDescOnce sync.Once
	file_client_errors_proto_rawDescData = file_client_errors_proto_rawDesc
)

func file_client_errors_proto_rawDescGZIP() []byte {
	file_client_errors_proto_rawDescOnce.Do(func() {
		file_client_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_errors_proto_rawDescData)
	})
	return file_client_errors_proto_rawDescData
}

var file_client_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_client_errors_proto_goTypes = []interface{}{
	(*ErrorDetails)(nil),                     // 0: client_errors.ErrorDetails
	nil,                                      // 1: client_errors.ErrorDetails.ParamsEntry
	(*ErrorDetails_ValidationErrors)(nil),    // 2: client_errors.ErrorDetails.ValidationErrors
	(*ErrorDetails_ValidationErrorItem)(nil), // 3: client_errors.ErrorDetails.ValidationErrorItem
}
var file_client_errors_proto_depIdxs = []int32{
	1, // 0: client_errors.ErrorDetails.params:type_name -> client_errors.ErrorDetails.ParamsEntry
	2, // 1: client_errors.ErrorDetails.validationErrors:type_name -> client_errors.ErrorDetails.ValidationErrors
	3, // 2: client_errors.ErrorDetails.ValidationErrors.items:type_name -> client_errors.ErrorDetails.ValidationErrorItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_errors_proto_init() }
func file_client_errors_proto_init() {
	if File_client_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
		file_client_errors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails_ValidationErrors); i {
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
		file_client_errors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails_ValidationErrorItem); i {
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
			RawDescriptor: file_client_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_errors_proto_goTypes,
		DependencyIndexes: file_client_errors_proto_depIdxs,
		MessageInfos:      file_client_errors_proto_msgTypes,
	}.Build()
	File_client_errors_proto = out.File
	file_client_errors_proto_rawDesc = nil
	file_client_errors_proto_goTypes = nil
	file_client_errors_proto_depIdxs = nil
}
