// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: orbital.proto

package orbital2

import (
	context "context"
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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbital_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_orbital_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_orbital_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbital_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_orbital_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_orbital_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbital_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_orbital_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_orbital_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EncodedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // This will contain the marshalled protobuf message
}

func (x *EncodedMessage) Reset() {
	*x = EncodedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbital_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodedMessage) ProtoMessage() {}

func (x *EncodedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_orbital_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodedMessage.ProtoReflect.Descriptor instead.
func (*EncodedMessage) Descriptor() ([]byte, []int) {
	return file_orbital_proto_rawDescGZIP(), []int{3}
}

func (x *EncodedMessage) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *EncodedMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_orbital_proto protoreflect.FileDescriptor

var file_orbital_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x32, 0x22, 0x2e, 0x0a, 0x06, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x75, 0x0a, 0x0e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x64, 0x69, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x32, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x32, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f,
	0x12, 0x11, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x32, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x62, 0x61, 0x79, 0x69, 0x70,
	0x70, 0x79, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x61, 0x6c, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orbital_proto_rawDescOnce sync.Once
	file_orbital_proto_rawDescData = file_orbital_proto_rawDesc
)

func file_orbital_proto_rawDescGZIP() []byte {
	file_orbital_proto_rawDescOnce.Do(func() {
		file_orbital_proto_rawDescData = protoimpl.X.CompressGZIP(file_orbital_proto_rawDescData)
	})
	return file_orbital_proto_rawDescData
}

var file_orbital_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orbital_proto_goTypes = []interface{}{
	(*Person)(nil),         // 0: orbital2.Person
	(*Request)(nil),        // 1: orbital2.Request
	(*Response)(nil),       // 2: orbital2.Response
	(*EncodedMessage)(nil), // 3: orbital2.EncodedMessage
}
var file_orbital_proto_depIdxs = []int32{
	0, // 0: orbital2.PeoplesService.editPerson:input_type -> orbital2.Person
	1, // 1: orbital2.PeoplesService.echo:input_type -> orbital2.Request
	0, // 2: orbital2.PeoplesService.editPerson:output_type -> orbital2.Person
	2, // 3: orbital2.PeoplesService.echo:output_type -> orbital2.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orbital_proto_init() }
func file_orbital_proto_init() {
	if File_orbital_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orbital_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_orbital_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_orbital_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_orbital_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodedMessage); i {
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
			RawDescriptor: file_orbital_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orbital_proto_goTypes,
		DependencyIndexes: file_orbital_proto_depIdxs,
		MessageInfos:      file_orbital_proto_msgTypes,
	}.Build()
	File_orbital_proto = out.File
	file_orbital_proto_rawDesc = nil
	file_orbital_proto_goTypes = nil
	file_orbital_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.5.2. DO NOT EDIT.

type PeoplesService interface {
	EditPerson(ctx context.Context, req *Person) (res *Person, err error)
	Echo(ctx context.Context, req *Request) (res *Response, err error)
}