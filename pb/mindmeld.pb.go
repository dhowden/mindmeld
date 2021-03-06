// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: mindmeld.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to use to access the service.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListServicesRequest) Reset() {
	*x = ListServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesRequest) ProtoMessage() {}

func (x *ListServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesRequest.ProtoReflect.Descriptor instead.
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{1}
}

type ListServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListServicesResponse) Reset() {
	*x = ListServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesResponse) ProtoMessage() {}

func (x *ListServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesResponse.ProtoReflect.Descriptor instead.
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{2}
}

func (x *ListServicesResponse) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Time the service was created.  Output only.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{3}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// Create a service hosted by this member.
type CreateServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateServiceRequest) Reset() {
	*x = CreateServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceRequest) ProtoMessage() {}

func (x *CreateServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{4}
}

func (x *CreateServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to use when using this service.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Dial address for the proxy.
	DialAddr string `protobuf:"bytes,2,opt,name=dial_addr,json=dialAddr,proto3" json:"dial_addr,omitempty"`
}

func (x *CreateServiceResponse) Reset() {
	*x = CreateServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceResponse) ProtoMessage() {}

func (x *CreateServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceResponse.ProtoReflect.Descriptor instead.
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{5}
}

func (x *CreateServiceResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateServiceResponse) GetDialAddr() string {
	if x != nil {
		return x.DialAddr
	}
	return ""
}

// Forward a service from this member to another.
type ForwardToServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ForwardToServiceRequest) Reset() {
	*x = ForwardToServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardToServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardToServiceRequest) ProtoMessage() {}

func (x *ForwardToServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardToServiceRequest.ProtoReflect.Descriptor instead.
func (*ForwardToServiceRequest) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{6}
}

func (x *ForwardToServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ForwardToServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to use when using this service.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Dial address for the proxy.
	DialAddr string `protobuf:"bytes,2,opt,name=dial_addr,json=dialAddr,proto3" json:"dial_addr,omitempty"`
}

func (x *ForwardToServiceResponse) Reset() {
	*x = ForwardToServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardToServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardToServiceResponse) ProtoMessage() {}

func (x *ForwardToServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardToServiceResponse.ProtoReflect.Descriptor instead.
func (*ForwardToServiceResponse) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{7}
}

func (x *ForwardToServiceResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ForwardToServiceResponse) GetDialAddr() string {
	if x != nil {
		return x.DialAddr
	}
	return ""
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*Payload_Header
	//	*Payload_Data
	Payload isPayload_Payload `protobuf_oneof:"payload"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mindmeld_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_mindmeld_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_mindmeld_proto_rawDescGZIP(), []int{8}
}

func (m *Payload) GetPayload() isPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Payload) GetHeader() *Header {
	if x, ok := x.GetPayload().(*Payload_Header); ok {
		return x.Header
	}
	return nil
}

func (x *Payload) GetData() []byte {
	if x, ok := x.GetPayload().(*Payload_Data); ok {
		return x.Data
	}
	return nil
}

type isPayload_Payload interface {
	isPayload_Payload()
}

type Payload_Header struct {
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type Payload_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*Payload_Header) isPayload_Payload() {}

func (*Payload_Data) isPayload_Payload() {}

var File_mindmeld_proto protoreflect.FileDescriptor

var file_mindmeld_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x45, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d,
	0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x22, 0x2d, 0x0a,
	0x17, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x18,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x22, 0x56, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c,
	0x64, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x32, 0x8e, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65,
	0x6c, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65,
	0x6c, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x10, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64, 0x2e, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4b, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65,
	0x6c, 0x64, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x11, 0x2e, 0x6d, 0x69, 0x6e,
	0x64, 0x6d, 0x65, 0x6c, 0x64, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x68, 0x6f, 0x77, 0x64, 0x65, 0x6e, 0x2f, 0x6d, 0x69, 0x6e, 0x64, 0x6d, 0x65, 0x6c, 0x64,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mindmeld_proto_rawDescOnce sync.Once
	file_mindmeld_proto_rawDescData = file_mindmeld_proto_rawDesc
)

func file_mindmeld_proto_rawDescGZIP() []byte {
	file_mindmeld_proto_rawDescOnce.Do(func() {
		file_mindmeld_proto_rawDescData = protoimpl.X.CompressGZIP(file_mindmeld_proto_rawDescData)
	})
	return file_mindmeld_proto_rawDescData
}

var file_mindmeld_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mindmeld_proto_goTypes = []interface{}{
	(*Header)(nil),                   // 0: mindmeld.Header
	(*ListServicesRequest)(nil),      // 1: mindmeld.ListServicesRequest
	(*ListServicesResponse)(nil),     // 2: mindmeld.ListServicesResponse
	(*Service)(nil),                  // 3: mindmeld.Service
	(*CreateServiceRequest)(nil),     // 4: mindmeld.CreateServiceRequest
	(*CreateServiceResponse)(nil),    // 5: mindmeld.CreateServiceResponse
	(*ForwardToServiceRequest)(nil),  // 6: mindmeld.ForwardToServiceRequest
	(*ForwardToServiceResponse)(nil), // 7: mindmeld.ForwardToServiceResponse
	(*Payload)(nil),                  // 8: mindmeld.Payload
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
}
var file_mindmeld_proto_depIdxs = []int32{
	3, // 0: mindmeld.ListServicesResponse.services:type_name -> mindmeld.Service
	9, // 1: mindmeld.Service.create_time:type_name -> google.protobuf.Timestamp
	0, // 2: mindmeld.Payload.header:type_name -> mindmeld.Header
	4, // 3: mindmeld.ControlService.CreateService:input_type -> mindmeld.CreateServiceRequest
	6, // 4: mindmeld.ControlService.ForwardToService:input_type -> mindmeld.ForwardToServiceRequest
	1, // 5: mindmeld.ControlService.ListServices:input_type -> mindmeld.ListServicesRequest
	8, // 6: mindmeld.ProxyService.ProxyConnection:input_type -> mindmeld.Payload
	5, // 7: mindmeld.ControlService.CreateService:output_type -> mindmeld.CreateServiceResponse
	7, // 8: mindmeld.ControlService.ForwardToService:output_type -> mindmeld.ForwardToServiceResponse
	2, // 9: mindmeld.ControlService.ListServices:output_type -> mindmeld.ListServicesResponse
	8, // 10: mindmeld.ProxyService.ProxyConnection:output_type -> mindmeld.Payload
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mindmeld_proto_init() }
func file_mindmeld_proto_init() {
	if File_mindmeld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mindmeld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_mindmeld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesRequest); i {
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
		file_mindmeld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesResponse); i {
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
		file_mindmeld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_mindmeld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceRequest); i {
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
		file_mindmeld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceResponse); i {
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
		file_mindmeld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardToServiceRequest); i {
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
		file_mindmeld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardToServiceResponse); i {
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
		file_mindmeld_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
	file_mindmeld_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Payload_Header)(nil),
		(*Payload_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mindmeld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mindmeld_proto_goTypes,
		DependencyIndexes: file_mindmeld_proto_depIdxs,
		MessageInfos:      file_mindmeld_proto_msgTypes,
	}.Build()
	File_mindmeld_proto = out.File
	file_mindmeld_proto_rawDesc = nil
	file_mindmeld_proto_goTypes = nil
	file_mindmeld_proto_depIdxs = nil
}
