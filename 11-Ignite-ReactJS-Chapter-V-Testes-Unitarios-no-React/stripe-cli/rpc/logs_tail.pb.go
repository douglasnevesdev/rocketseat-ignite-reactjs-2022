// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: logs_tail.proto

package rpc

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

type LogsTailRequest_Account int32

const (
	LogsTailRequest_ACCOUNT_UNSPECIFIED LogsTailRequest_Account = 0
	LogsTailRequest_ACCOUNT_CONNECT_IN  LogsTailRequest_Account = 1
	LogsTailRequest_ACCOUNT_CONNECT_OUT LogsTailRequest_Account = 2
	LogsTailRequest_ACCOUNT_SELF        LogsTailRequest_Account = 3
)

// Enum value maps for LogsTailRequest_Account.
var (
	LogsTailRequest_Account_name = map[int32]string{
		0: "ACCOUNT_UNSPECIFIED",
		1: "ACCOUNT_CONNECT_IN",
		2: "ACCOUNT_CONNECT_OUT",
		3: "ACCOUNT_SELF",
	}
	LogsTailRequest_Account_value = map[string]int32{
		"ACCOUNT_UNSPECIFIED": 0,
		"ACCOUNT_CONNECT_IN":  1,
		"ACCOUNT_CONNECT_OUT": 2,
		"ACCOUNT_SELF":        3,
	}
)

func (x LogsTailRequest_Account) Enum() *LogsTailRequest_Account {
	p := new(LogsTailRequest_Account)
	*p = x
	return p
}

func (x LogsTailRequest_Account) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTailRequest_Account) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_tail_proto_enumTypes[0].Descriptor()
}

func (LogsTailRequest_Account) Type() protoreflect.EnumType {
	return &file_logs_tail_proto_enumTypes[0]
}

func (x LogsTailRequest_Account) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTailRequest_Account.Descriptor instead.
func (LogsTailRequest_Account) EnumDescriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{0, 0}
}

type LogsTailRequest_HttpMethod int32

const (
	LogsTailRequest_HTTP_METHOD_UNSPECIFIED LogsTailRequest_HttpMethod = 0
	LogsTailRequest_HTTP_METHOD_GET         LogsTailRequest_HttpMethod = 1
	LogsTailRequest_HTTP_METHOD_POST        LogsTailRequest_HttpMethod = 2
	LogsTailRequest_HTTP_METHOD_DELETE      LogsTailRequest_HttpMethod = 3
)

// Enum value maps for LogsTailRequest_HttpMethod.
var (
	LogsTailRequest_HttpMethod_name = map[int32]string{
		0: "HTTP_METHOD_UNSPECIFIED",
		1: "HTTP_METHOD_GET",
		2: "HTTP_METHOD_POST",
		3: "HTTP_METHOD_DELETE",
	}
	LogsTailRequest_HttpMethod_value = map[string]int32{
		"HTTP_METHOD_UNSPECIFIED": 0,
		"HTTP_METHOD_GET":         1,
		"HTTP_METHOD_POST":        2,
		"HTTP_METHOD_DELETE":      3,
	}
)

func (x LogsTailRequest_HttpMethod) Enum() *LogsTailRequest_HttpMethod {
	p := new(LogsTailRequest_HttpMethod)
	*p = x
	return p
}

func (x LogsTailRequest_HttpMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTailRequest_HttpMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_tail_proto_enumTypes[1].Descriptor()
}

func (LogsTailRequest_HttpMethod) Type() protoreflect.EnumType {
	return &file_logs_tail_proto_enumTypes[1]
}

func (x LogsTailRequest_HttpMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTailRequest_HttpMethod.Descriptor instead.
func (LogsTailRequest_HttpMethod) EnumDescriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{0, 1}
}

type LogsTailRequest_RequestStatus int32

const (
	LogsTailRequest_REQUEST_STATUS_UNSPECIFIED LogsTailRequest_RequestStatus = 0
	LogsTailRequest_REQUEST_STATUS_SUCCEEDED   LogsTailRequest_RequestStatus = 1
	LogsTailRequest_REQUEST_STATUS_FAILED      LogsTailRequest_RequestStatus = 2
)

// Enum value maps for LogsTailRequest_RequestStatus.
var (
	LogsTailRequest_RequestStatus_name = map[int32]string{
		0: "REQUEST_STATUS_UNSPECIFIED",
		1: "REQUEST_STATUS_SUCCEEDED",
		2: "REQUEST_STATUS_FAILED",
	}
	LogsTailRequest_RequestStatus_value = map[string]int32{
		"REQUEST_STATUS_UNSPECIFIED": 0,
		"REQUEST_STATUS_SUCCEEDED":   1,
		"REQUEST_STATUS_FAILED":      2,
	}
)

func (x LogsTailRequest_RequestStatus) Enum() *LogsTailRequest_RequestStatus {
	p := new(LogsTailRequest_RequestStatus)
	*p = x
	return p
}

func (x LogsTailRequest_RequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTailRequest_RequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_tail_proto_enumTypes[2].Descriptor()
}

func (LogsTailRequest_RequestStatus) Type() protoreflect.EnumType {
	return &file_logs_tail_proto_enumTypes[2]
}

func (x LogsTailRequest_RequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTailRequest_RequestStatus.Descriptor instead.
func (LogsTailRequest_RequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{0, 2}
}

type LogsTailRequest_Source int32

const (
	LogsTailRequest_SOURCE_UNSPECIFIED LogsTailRequest_Source = 0
	LogsTailRequest_SOURCE_API         LogsTailRequest_Source = 1
	LogsTailRequest_SOURCE_DASHBOARD   LogsTailRequest_Source = 2
)

// Enum value maps for LogsTailRequest_Source.
var (
	LogsTailRequest_Source_name = map[int32]string{
		0: "SOURCE_UNSPECIFIED",
		1: "SOURCE_API",
		2: "SOURCE_DASHBOARD",
	}
	LogsTailRequest_Source_value = map[string]int32{
		"SOURCE_UNSPECIFIED": 0,
		"SOURCE_API":         1,
		"SOURCE_DASHBOARD":   2,
	}
)

func (x LogsTailRequest_Source) Enum() *LogsTailRequest_Source {
	p := new(LogsTailRequest_Source)
	*p = x
	return p
}

func (x LogsTailRequest_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTailRequest_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_tail_proto_enumTypes[3].Descriptor()
}

func (LogsTailRequest_Source) Type() protoreflect.EnumType {
	return &file_logs_tail_proto_enumTypes[3]
}

func (x LogsTailRequest_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTailRequest_Source.Descriptor instead.
func (LogsTailRequest_Source) EnumDescriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{0, 3}
}

type LogsTailRequest_StatusCodeType int32

const (
	LogsTailRequest_STATUS_CODE_TYPE_UNSPECIFIED LogsTailRequest_StatusCodeType = 0
	LogsTailRequest_STATUS_CODE_TYPE_2XX         LogsTailRequest_StatusCodeType = 1
	LogsTailRequest_STATUS_CODE_TYPE_4XX         LogsTailRequest_StatusCodeType = 2
	LogsTailRequest_STATUS_CODE_TYPE_5XX         LogsTailRequest_StatusCodeType = 3
)

// Enum value maps for LogsTailRequest_StatusCodeType.
var (
	LogsTailRequest_StatusCodeType_name = map[int32]string{
		0: "STATUS_CODE_TYPE_UNSPECIFIED",
		1: "STATUS_CODE_TYPE_2XX",
		2: "STATUS_CODE_TYPE_4XX",
		3: "STATUS_CODE_TYPE_5XX",
	}
	LogsTailRequest_StatusCodeType_value = map[string]int32{
		"STATUS_CODE_TYPE_UNSPECIFIED": 0,
		"STATUS_CODE_TYPE_2XX":         1,
		"STATUS_CODE_TYPE_4XX":         2,
		"STATUS_CODE_TYPE_5XX":         3,
	}
)

func (x LogsTailRequest_StatusCodeType) Enum() *LogsTailRequest_StatusCodeType {
	p := new(LogsTailRequest_StatusCodeType)
	*p = x
	return p
}

func (x LogsTailRequest_StatusCodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTailRequest_StatusCodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_tail_proto_enumTypes[4].Descriptor()
}

func (LogsTailRequest_StatusCodeType) Type() protoreflect.EnumType {
	return &file_logs_tail_proto_enumTypes[4]
}

func (x LogsTailRequest_StatusCodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTailRequest_StatusCodeType.Descriptor instead.
func (LogsTailRequest_StatusCodeType) EnumDescriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{0, 4}
}

type LogsTailResponse_State int32

const (
	LogsTailResponse_STATE_UNSPECIFIED  LogsTailResponse_State = 0
	LogsTailResponse_STATE_LOADING      LogsTailResponse_State = 1
	LogsTailResponse_STATE_RECONNECTING LogsTailResponse_State = 2
	LogsTailResponse_STATE_READY        LogsTailResponse_State = 3
	LogsTailResponse_STATE_DONE         LogsTailResponse_State = 4
)

// Enum value maps for LogsTailResponse_State.
var (
	LogsTailResponse_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_LOADING",
		2: "STATE_RECONNECTING",
		3: "STATE_READY",
		4: "STATE_DONE",
	}
	LogsTailResponse_State_value = map[string]int32{
		"STATE_UNSPECIFIED":  0,
		"STATE_LOADING":      1,
		"STATE_RECONNECTING": 2,
		"STATE_READY":        3,
		"STATE_DONE":         4,
	}
)

func (x LogsTailResponse_State) Enum() *LogsTailResponse_State {
	p := new(LogsTailResponse_State)
	*p = x
	return p
}

func (x LogsTailResponse_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTailResponse_State) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_tail_proto_enumTypes[5].Descriptor()
}

func (LogsTailResponse_State) Type() protoreflect.EnumType {
	return &file_logs_tail_proto_enumTypes[5]
}

func (x LogsTailResponse_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTailResponse_State.Descriptor instead.
func (LogsTailResponse_State) EnumDescriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{1, 0}
}

type LogsTailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *CONNECT ONLY* Filter request logs by source and destination account
	FilterAccounts []LogsTailRequest_Account `protobuf:"varint,1,rep,packed,name=filter_accounts,json=filterAccounts,proto3,enum=rpc.LogsTailRequest_Account" json:"filter_accounts,omitempty"`
	// Filter request logs by http method
	FilterHttpMethods []LogsTailRequest_HttpMethod `protobuf:"varint,2,rep,packed,name=filter_http_methods,json=filterHttpMethods,proto3,enum=rpc.LogsTailRequest_HttpMethod" json:"filter_http_methods,omitempty"`
	// Filter request logs by ip address
	FilterIpAddresses []string `protobuf:"bytes,3,rep,name=filter_ip_addresses,json=filterIpAddresses,proto3" json:"filter_ip_addresses,omitempty"`
	// Filter request logs by request path
	FilterRequestPaths []string `protobuf:"bytes,4,rep,name=filter_request_paths,json=filterRequestPaths,proto3" json:"filter_request_paths,omitempty"`
	// Filter request logs by request status
	FilterRequestStatuses []LogsTailRequest_RequestStatus `protobuf:"varint,5,rep,packed,name=filter_request_statuses,json=filterRequestStatuses,proto3,enum=rpc.LogsTailRequest_RequestStatus" json:"filter_request_statuses,omitempty"`
	// Filter request logs by source
	FilterSources []LogsTailRequest_Source `protobuf:"varint,6,rep,packed,name=filter_sources,json=filterSources,proto3,enum=rpc.LogsTailRequest_Source" json:"filter_sources,omitempty"`
	// Filter request logs by status code
	FilterStatusCodes []string `protobuf:"bytes,7,rep,name=filter_status_codes,json=filterStatusCodes,proto3" json:"filter_status_codes,omitempty"`
	// Filter request logs by status code type
	FilterStatusCodeTypes []LogsTailRequest_StatusCodeType `protobuf:"varint,8,rep,packed,name=filter_status_code_types,json=filterStatusCodeTypes,proto3,enum=rpc.LogsTailRequest_StatusCodeType" json:"filter_status_code_types,omitempty"`
}

func (x *LogsTailRequest) Reset() {
	*x = LogsTailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_tail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTailRequest) ProtoMessage() {}

func (x *LogsTailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logs_tail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTailRequest.ProtoReflect.Descriptor instead.
func (*LogsTailRequest) Descriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{0}
}

func (x *LogsTailRequest) GetFilterAccounts() []LogsTailRequest_Account {
	if x != nil {
		return x.FilterAccounts
	}
	return nil
}

func (x *LogsTailRequest) GetFilterHttpMethods() []LogsTailRequest_HttpMethod {
	if x != nil {
		return x.FilterHttpMethods
	}
	return nil
}

func (x *LogsTailRequest) GetFilterIpAddresses() []string {
	if x != nil {
		return x.FilterIpAddresses
	}
	return nil
}

func (x *LogsTailRequest) GetFilterRequestPaths() []string {
	if x != nil {
		return x.FilterRequestPaths
	}
	return nil
}

func (x *LogsTailRequest) GetFilterRequestStatuses() []LogsTailRequest_RequestStatus {
	if x != nil {
		return x.FilterRequestStatuses
	}
	return nil
}

func (x *LogsTailRequest) GetFilterSources() []LogsTailRequest_Source {
	if x != nil {
		return x.FilterSources
	}
	return nil
}

func (x *LogsTailRequest) GetFilterStatusCodes() []string {
	if x != nil {
		return x.FilterStatusCodes
	}
	return nil
}

func (x *LogsTailRequest) GetFilterStatusCodeTypes() []LogsTailRequest_StatusCodeType {
	if x != nil {
		return x.FilterStatusCodeTypes
	}
	return nil
}

type LogsTailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*LogsTailResponse_State_
	//	*LogsTailResponse_Log_
	Content isLogsTailResponse_Content `protobuf_oneof:"content"`
}

func (x *LogsTailResponse) Reset() {
	*x = LogsTailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_tail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTailResponse) ProtoMessage() {}

func (x *LogsTailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logs_tail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTailResponse.ProtoReflect.Descriptor instead.
func (*LogsTailResponse) Descriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{1}
}

func (m *LogsTailResponse) GetContent() isLogsTailResponse_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *LogsTailResponse) GetState() LogsTailResponse_State {
	if x, ok := x.GetContent().(*LogsTailResponse_State_); ok {
		return x.State
	}
	return LogsTailResponse_STATE_UNSPECIFIED
}

func (x *LogsTailResponse) GetLog() *LogsTailResponse_Log {
	if x, ok := x.GetContent().(*LogsTailResponse_Log_); ok {
		return x.Log
	}
	return nil
}

type isLogsTailResponse_Content interface {
	isLogsTailResponse_Content()
}

type LogsTailResponse_State_ struct {
	// Check if the stream ready
	State LogsTailResponse_State `protobuf:"varint,1,opt,name=state,proto3,enum=rpc.LogsTailResponse_State,oneof"`
}

type LogsTailResponse_Log_ struct {
	// A Stripe API log
	Log *LogsTailResponse_Log `protobuf:"bytes,2,opt,name=log,proto3,oneof"`
}

func (*LogsTailResponse_State_) isLogsTailResponse_Content() {}

func (*LogsTailResponse_Log_) isLogsTailResponse_Content() {}

type LogsTailResponse_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Livemode  bool                        `protobuf:"varint,1,opt,name=livemode,proto3" json:"livemode,omitempty"`
	Method    string                      `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Url       string                      `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Status    int64                       `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	RequestId string                      `protobuf:"bytes,5,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	CreatedAt int64                       `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Error     *LogsTailResponse_Log_Error `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LogsTailResponse_Log) Reset() {
	*x = LogsTailResponse_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_tail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTailResponse_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTailResponse_Log) ProtoMessage() {}

func (x *LogsTailResponse_Log) ProtoReflect() protoreflect.Message {
	mi := &file_logs_tail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTailResponse_Log.ProtoReflect.Descriptor instead.
func (*LogsTailResponse_Log) Descriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LogsTailResponse_Log) GetLivemode() bool {
	if x != nil {
		return x.Livemode
	}
	return false
}

func (x *LogsTailResponse_Log) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *LogsTailResponse_Log) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LogsTailResponse_Log) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LogsTailResponse_Log) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *LogsTailResponse_Log) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *LogsTailResponse_Log) GetError() *LogsTailResponse_Log_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type LogsTailResponse_Log_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Charge      string `protobuf:"bytes,2,opt,name=charge,proto3" json:"charge,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	DeclineCode string `protobuf:"bytes,4,opt,name=decline_code,json=declineCode,proto3" json:"decline_code,omitempty"`
	Message     string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Param       string `protobuf:"bytes,6,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *LogsTailResponse_Log_Error) Reset() {
	*x = LogsTailResponse_Log_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_tail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTailResponse_Log_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTailResponse_Log_Error) ProtoMessage() {}

func (x *LogsTailResponse_Log_Error) ProtoReflect() protoreflect.Message {
	mi := &file_logs_tail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTailResponse_Log_Error.ProtoReflect.Descriptor instead.
func (*LogsTailResponse_Log_Error) Descriptor() ([]byte, []int) {
	return file_logs_tail_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *LogsTailResponse_Log_Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LogsTailResponse_Log_Error) GetCharge() string {
	if x != nil {
		return x.Charge
	}
	return ""
}

func (x *LogsTailResponse_Log_Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LogsTailResponse_Log_Error) GetDeclineCode() string {
	if x != nil {
		return x.DeclineCode
	}
	return ""
}

func (x *LogsTailResponse_Log_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogsTailResponse_Log_Error) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

var File_logs_tail_proto protoreflect.FileDescriptor

var file_logs_tail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0xc3, 0x08, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x73, 0x54,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x4f, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x11, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x5a, 0x0a, 0x17, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x15, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x12, 0x42, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x11, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x18, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67,
	0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x15, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x22, 0x65, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x13, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4c, 0x46, 0x10, 0x03, 0x22, 0x6c, 0x0a, 0x0a, 0x48, 0x74, 0x74,
	0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x54, 0x54, 0x50, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x54, 0x54,
	0x50, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x02, 0x12,
	0x16, 0x0a, 0x12, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22, 0x68, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x02, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x41, 0x50,
	0x49, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x41,
	0x53, 0x48, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x02, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x32, 0x58, 0x58, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x34, 0x58, 0x58,
	0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x35, 0x58, 0x58, 0x10, 0x03, 0x22, 0xe5, 0x04, 0x0a,
	0x10, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x48, 0x00,
	0x52, 0x03, 0x6c, 0x6f, 0x67, 0x1a, 0xf5, 0x02, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4c, 0x6f, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x1a, 0x9a, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65,
	0x63, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x6a, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x04, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logs_tail_proto_rawDescOnce sync.Once
	file_logs_tail_proto_rawDescData = file_logs_tail_proto_rawDesc
)

func file_logs_tail_proto_rawDescGZIP() []byte {
	file_logs_tail_proto_rawDescOnce.Do(func() {
		file_logs_tail_proto_rawDescData = protoimpl.X.CompressGZIP(file_logs_tail_proto_rawDescData)
	})
	return file_logs_tail_proto_rawDescData
}

var file_logs_tail_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_logs_tail_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logs_tail_proto_goTypes = []interface{}{
	(LogsTailRequest_Account)(0),        // 0: rpc.LogsTailRequest.Account
	(LogsTailRequest_HttpMethod)(0),     // 1: rpc.LogsTailRequest.HttpMethod
	(LogsTailRequest_RequestStatus)(0),  // 2: rpc.LogsTailRequest.RequestStatus
	(LogsTailRequest_Source)(0),         // 3: rpc.LogsTailRequest.Source
	(LogsTailRequest_StatusCodeType)(0), // 4: rpc.LogsTailRequest.StatusCodeType
	(LogsTailResponse_State)(0),         // 5: rpc.LogsTailResponse.State
	(*LogsTailRequest)(nil),             // 6: rpc.LogsTailRequest
	(*LogsTailResponse)(nil),            // 7: rpc.LogsTailResponse
	(*LogsTailResponse_Log)(nil),        // 8: rpc.LogsTailResponse.Log
	(*LogsTailResponse_Log_Error)(nil),  // 9: rpc.LogsTailResponse.Log.Error
}
var file_logs_tail_proto_depIdxs = []int32{
	0, // 0: rpc.LogsTailRequest.filter_accounts:type_name -> rpc.LogsTailRequest.Account
	1, // 1: rpc.LogsTailRequest.filter_http_methods:type_name -> rpc.LogsTailRequest.HttpMethod
	2, // 2: rpc.LogsTailRequest.filter_request_statuses:type_name -> rpc.LogsTailRequest.RequestStatus
	3, // 3: rpc.LogsTailRequest.filter_sources:type_name -> rpc.LogsTailRequest.Source
	4, // 4: rpc.LogsTailRequest.filter_status_code_types:type_name -> rpc.LogsTailRequest.StatusCodeType
	5, // 5: rpc.LogsTailResponse.state:type_name -> rpc.LogsTailResponse.State
	8, // 6: rpc.LogsTailResponse.log:type_name -> rpc.LogsTailResponse.Log
	9, // 7: rpc.LogsTailResponse.Log.error:type_name -> rpc.LogsTailResponse.Log.Error
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_logs_tail_proto_init() }
func file_logs_tail_proto_init() {
	if File_logs_tail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logs_tail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsTailRequest); i {
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
		file_logs_tail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsTailResponse); i {
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
		file_logs_tail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsTailResponse_Log); i {
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
		file_logs_tail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsTailResponse_Log_Error); i {
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
	file_logs_tail_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LogsTailResponse_State_)(nil),
		(*LogsTailResponse_Log_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logs_tail_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logs_tail_proto_goTypes,
		DependencyIndexes: file_logs_tail_proto_depIdxs,
		EnumInfos:         file_logs_tail_proto_enumTypes,
		MessageInfos:      file_logs_tail_proto_msgTypes,
	}.Build()
	File_logs_tail_proto = out.File
	file_logs_tail_proto_rawDesc = nil
	file_logs_tail_proto_goTypes = nil
	file_logs_tail_proto_depIdxs = nil
}
