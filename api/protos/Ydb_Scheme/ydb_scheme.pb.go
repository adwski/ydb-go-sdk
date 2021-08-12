// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_scheme.proto

package Ydb_Scheme

import (
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/v2/api/protos/Ydb_Operations"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Entry_Type int32

const (
	Entry_TYPE_UNSPECIFIED   Entry_Type = 0
	Entry_DIRECTORY          Entry_Type = 1
	Entry_TABLE              Entry_Type = 2
	Entry_PERS_QUEUE_GROUP   Entry_Type = 3
	Entry_DATABASE           Entry_Type = 4
	Entry_RTMR_VOLUME        Entry_Type = 5
	Entry_BLOCK_STORE_VOLUME Entry_Type = 6
	Entry_COORDINATION_NODE  Entry_Type = 7
	Entry_SEQUENCE           Entry_Type = 15
)

var Entry_Type_name = map[int32]string{
	0:  "TYPE_UNSPECIFIED",
	1:  "DIRECTORY",
	2:  "TABLE",
	3:  "PERS_QUEUE_GROUP",
	4:  "DATABASE",
	5:  "RTMR_VOLUME",
	6:  "BLOCK_STORE_VOLUME",
	7:  "COORDINATION_NODE",
	15: "SEQUENCE",
}

var Entry_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED":   0,
	"DIRECTORY":          1,
	"TABLE":              2,
	"PERS_QUEUE_GROUP":   3,
	"DATABASE":           4,
	"RTMR_VOLUME":        5,
	"BLOCK_STORE_VOLUME": 6,
	"COORDINATION_NODE":  7,
	"SEQUENCE":           15,
}

func (x Entry_Type) String() string {
	return proto.EnumName(Entry_Type_name, int32(x))
}

func (Entry_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{7, 0}
}

// Create directory.
// All intermediate directories must be created
type MakeDirectoryRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Path                 string                          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MakeDirectoryRequest) Reset()         { *m = MakeDirectoryRequest{} }
func (m *MakeDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*MakeDirectoryRequest) ProtoMessage()    {}
func (*MakeDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{0}
}

func (m *MakeDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeDirectoryRequest.Unmarshal(m, b)
}
func (m *MakeDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *MakeDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeDirectoryRequest.Merge(m, src)
}
func (m *MakeDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_MakeDirectoryRequest.Size(m)
}
func (m *MakeDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakeDirectoryRequest proto.InternalMessageInfo

func (m *MakeDirectoryRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *MakeDirectoryRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type MakeDirectoryResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MakeDirectoryResponse) Reset()         { *m = MakeDirectoryResponse{} }
func (m *MakeDirectoryResponse) String() string { return proto.CompactTextString(m) }
func (*MakeDirectoryResponse) ProtoMessage()    {}
func (*MakeDirectoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{1}
}

func (m *MakeDirectoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeDirectoryResponse.Unmarshal(m, b)
}
func (m *MakeDirectoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeDirectoryResponse.Marshal(b, m, deterministic)
}
func (m *MakeDirectoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeDirectoryResponse.Merge(m, src)
}
func (m *MakeDirectoryResponse) XXX_Size() int {
	return xxx_messageInfo_MakeDirectoryResponse.Size(m)
}
func (m *MakeDirectoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeDirectoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MakeDirectoryResponse proto.InternalMessageInfo

func (m *MakeDirectoryResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// Remove directory
type RemoveDirectoryRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Path                 string                          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RemoveDirectoryRequest) Reset()         { *m = RemoveDirectoryRequest{} }
func (m *RemoveDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveDirectoryRequest) ProtoMessage()    {}
func (*RemoveDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{2}
}

func (m *RemoveDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDirectoryRequest.Unmarshal(m, b)
}
func (m *RemoveDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *RemoveDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDirectoryRequest.Merge(m, src)
}
func (m *RemoveDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveDirectoryRequest.Size(m)
}
func (m *RemoveDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDirectoryRequest proto.InternalMessageInfo

func (m *RemoveDirectoryRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *RemoveDirectoryRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type RemoveDirectoryResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RemoveDirectoryResponse) Reset()         { *m = RemoveDirectoryResponse{} }
func (m *RemoveDirectoryResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveDirectoryResponse) ProtoMessage()    {}
func (*RemoveDirectoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{3}
}

func (m *RemoveDirectoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDirectoryResponse.Unmarshal(m, b)
}
func (m *RemoveDirectoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDirectoryResponse.Marshal(b, m, deterministic)
}
func (m *RemoveDirectoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDirectoryResponse.Merge(m, src)
}
func (m *RemoveDirectoryResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveDirectoryResponse.Size(m)
}
func (m *RemoveDirectoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDirectoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDirectoryResponse proto.InternalMessageInfo

func (m *RemoveDirectoryResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// List directory
type ListDirectoryRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Path                 string                          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ListDirectoryRequest) Reset()         { *m = ListDirectoryRequest{} }
func (m *ListDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*ListDirectoryRequest) ProtoMessage()    {}
func (*ListDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{4}
}

func (m *ListDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDirectoryRequest.Unmarshal(m, b)
}
func (m *ListDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *ListDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDirectoryRequest.Merge(m, src)
}
func (m *ListDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_ListDirectoryRequest.Size(m)
}
func (m *ListDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDirectoryRequest proto.InternalMessageInfo

func (m *ListDirectoryRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *ListDirectoryRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ListDirectoryResponse struct {
	// Holds ListDirectoryResult in case of successful call
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListDirectoryResponse) Reset()         { *m = ListDirectoryResponse{} }
func (m *ListDirectoryResponse) String() string { return proto.CompactTextString(m) }
func (*ListDirectoryResponse) ProtoMessage()    {}
func (*ListDirectoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{5}
}

func (m *ListDirectoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDirectoryResponse.Unmarshal(m, b)
}
func (m *ListDirectoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDirectoryResponse.Marshal(b, m, deterministic)
}
func (m *ListDirectoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDirectoryResponse.Merge(m, src)
}
func (m *ListDirectoryResponse) XXX_Size() int {
	return xxx_messageInfo_ListDirectoryResponse.Size(m)
}
func (m *ListDirectoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDirectoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDirectoryResponse proto.InternalMessageInfo

func (m *ListDirectoryResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type Permissions struct {
	// SID (Security ID) of user or group
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	PermissionNames      []string `protobuf:"bytes,2,rep,name=permission_names,json=permissionNames,proto3" json:"permission_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{6}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Permissions) GetPermissionNames() []string {
	if m != nil {
		return m.PermissionNames
	}
	return nil
}

type Entry struct {
	// Name of scheme entry (dir2 of /dir1/dir2)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SID (Security ID) of user or group
	Owner                string         `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 Entry_Type     `protobuf:"varint,5,opt,name=type,proto3,enum=Ydb.Scheme.Entry_Type" json:"type,omitempty"`
	EffectivePermissions []*Permissions `protobuf:"bytes,6,rep,name=effective_permissions,json=effectivePermissions,proto3" json:"effective_permissions,omitempty"`
	Permissions          []*Permissions `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{7}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Entry) GetType() Entry_Type {
	if m != nil {
		return m.Type
	}
	return Entry_TYPE_UNSPECIFIED
}

func (m *Entry) GetEffectivePermissions() []*Permissions {
	if m != nil {
		return m.EffectivePermissions
	}
	return nil
}

func (m *Entry) GetPermissions() []*Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type ListDirectoryResult struct {
	Self                 *Entry   `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Children             []*Entry `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDirectoryResult) Reset()         { *m = ListDirectoryResult{} }
func (m *ListDirectoryResult) String() string { return proto.CompactTextString(m) }
func (*ListDirectoryResult) ProtoMessage()    {}
func (*ListDirectoryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{8}
}

func (m *ListDirectoryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDirectoryResult.Unmarshal(m, b)
}
func (m *ListDirectoryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDirectoryResult.Marshal(b, m, deterministic)
}
func (m *ListDirectoryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDirectoryResult.Merge(m, src)
}
func (m *ListDirectoryResult) XXX_Size() int {
	return xxx_messageInfo_ListDirectoryResult.Size(m)
}
func (m *ListDirectoryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDirectoryResult.DiscardUnknown(m)
}

var xxx_messageInfo_ListDirectoryResult proto.InternalMessageInfo

func (m *ListDirectoryResult) GetSelf() *Entry {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *ListDirectoryResult) GetChildren() []*Entry {
	if m != nil {
		return m.Children
	}
	return nil
}

// Returns information about object with given path
type DescribePathRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Path                 string                          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DescribePathRequest) Reset()         { *m = DescribePathRequest{} }
func (m *DescribePathRequest) String() string { return proto.CompactTextString(m) }
func (*DescribePathRequest) ProtoMessage()    {}
func (*DescribePathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{9}
}

func (m *DescribePathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePathRequest.Unmarshal(m, b)
}
func (m *DescribePathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePathRequest.Marshal(b, m, deterministic)
}
func (m *DescribePathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePathRequest.Merge(m, src)
}
func (m *DescribePathRequest) XXX_Size() int {
	return xxx_messageInfo_DescribePathRequest.Size(m)
}
func (m *DescribePathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePathRequest proto.InternalMessageInfo

func (m *DescribePathRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *DescribePathRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type DescribePathResponse struct {
	// Holds DescribePathResult in case of DescribePathResult
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DescribePathResponse) Reset()         { *m = DescribePathResponse{} }
func (m *DescribePathResponse) String() string { return proto.CompactTextString(m) }
func (*DescribePathResponse) ProtoMessage()    {}
func (*DescribePathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{10}
}

func (m *DescribePathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePathResponse.Unmarshal(m, b)
}
func (m *DescribePathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePathResponse.Marshal(b, m, deterministic)
}
func (m *DescribePathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePathResponse.Merge(m, src)
}
func (m *DescribePathResponse) XXX_Size() int {
	return xxx_messageInfo_DescribePathResponse.Size(m)
}
func (m *DescribePathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePathResponse proto.InternalMessageInfo

func (m *DescribePathResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type DescribePathResult struct {
	Self                 *Entry   `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribePathResult) Reset()         { *m = DescribePathResult{} }
func (m *DescribePathResult) String() string { return proto.CompactTextString(m) }
func (*DescribePathResult) ProtoMessage()    {}
func (*DescribePathResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{11}
}

func (m *DescribePathResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePathResult.Unmarshal(m, b)
}
func (m *DescribePathResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePathResult.Marshal(b, m, deterministic)
}
func (m *DescribePathResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePathResult.Merge(m, src)
}
func (m *DescribePathResult) XXX_Size() int {
	return xxx_messageInfo_DescribePathResult.Size(m)
}
func (m *DescribePathResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePathResult.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePathResult proto.InternalMessageInfo

func (m *DescribePathResult) GetSelf() *Entry {
	if m != nil {
		return m.Self
	}
	return nil
}

type PermissionsAction struct {
	// Types that are valid to be assigned to Action:
	//	*PermissionsAction_Grant
	//	*PermissionsAction_Revoke
	//	*PermissionsAction_Set
	//	*PermissionsAction_ChangeOwner
	Action               isPermissionsAction_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PermissionsAction) Reset()         { *m = PermissionsAction{} }
func (m *PermissionsAction) String() string { return proto.CompactTextString(m) }
func (*PermissionsAction) ProtoMessage()    {}
func (*PermissionsAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{12}
}

func (m *PermissionsAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionsAction.Unmarshal(m, b)
}
func (m *PermissionsAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionsAction.Marshal(b, m, deterministic)
}
func (m *PermissionsAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionsAction.Merge(m, src)
}
func (m *PermissionsAction) XXX_Size() int {
	return xxx_messageInfo_PermissionsAction.Size(m)
}
func (m *PermissionsAction) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionsAction.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionsAction proto.InternalMessageInfo

type isPermissionsAction_Action interface {
	isPermissionsAction_Action()
}

type PermissionsAction_Grant struct {
	Grant *Permissions `protobuf:"bytes,1,opt,name=grant,proto3,oneof"`
}

type PermissionsAction_Revoke struct {
	Revoke *Permissions `protobuf:"bytes,2,opt,name=revoke,proto3,oneof"`
}

type PermissionsAction_Set struct {
	Set *Permissions `protobuf:"bytes,3,opt,name=set,proto3,oneof"`
}

type PermissionsAction_ChangeOwner struct {
	ChangeOwner string `protobuf:"bytes,4,opt,name=change_owner,json=changeOwner,proto3,oneof"`
}

func (*PermissionsAction_Grant) isPermissionsAction_Action() {}

func (*PermissionsAction_Revoke) isPermissionsAction_Action() {}

func (*PermissionsAction_Set) isPermissionsAction_Action() {}

func (*PermissionsAction_ChangeOwner) isPermissionsAction_Action() {}

func (m *PermissionsAction) GetAction() isPermissionsAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *PermissionsAction) GetGrant() *Permissions {
	if x, ok := m.GetAction().(*PermissionsAction_Grant); ok {
		return x.Grant
	}
	return nil
}

func (m *PermissionsAction) GetRevoke() *Permissions {
	if x, ok := m.GetAction().(*PermissionsAction_Revoke); ok {
		return x.Revoke
	}
	return nil
}

func (m *PermissionsAction) GetSet() *Permissions {
	if x, ok := m.GetAction().(*PermissionsAction_Set); ok {
		return x.Set
	}
	return nil
}

func (m *PermissionsAction) GetChangeOwner() string {
	if x, ok := m.GetAction().(*PermissionsAction_ChangeOwner); ok {
		return x.ChangeOwner
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PermissionsAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PermissionsAction_Grant)(nil),
		(*PermissionsAction_Revoke)(nil),
		(*PermissionsAction_Set)(nil),
		(*PermissionsAction_ChangeOwner)(nil),
	}
}

// Modify permissions of given object
type ModifyPermissionsRequest struct {
	OperationParams *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Path            string                          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Actions         []*PermissionsAction            `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	// Clear all permissions on the object for all subjects
	ClearPermissions     bool     `protobuf:"varint,4,opt,name=clear_permissions,json=clearPermissions,proto3" json:"clear_permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyPermissionsRequest) Reset()         { *m = ModifyPermissionsRequest{} }
func (m *ModifyPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyPermissionsRequest) ProtoMessage()    {}
func (*ModifyPermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{13}
}

func (m *ModifyPermissionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyPermissionsRequest.Unmarshal(m, b)
}
func (m *ModifyPermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyPermissionsRequest.Marshal(b, m, deterministic)
}
func (m *ModifyPermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyPermissionsRequest.Merge(m, src)
}
func (m *ModifyPermissionsRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyPermissionsRequest.Size(m)
}
func (m *ModifyPermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyPermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyPermissionsRequest proto.InternalMessageInfo

func (m *ModifyPermissionsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *ModifyPermissionsRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ModifyPermissionsRequest) GetActions() []*PermissionsAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *ModifyPermissionsRequest) GetClearPermissions() bool {
	if m != nil {
		return m.ClearPermissions
	}
	return false
}

type ModifyPermissionsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ModifyPermissionsResponse) Reset()         { *m = ModifyPermissionsResponse{} }
func (m *ModifyPermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyPermissionsResponse) ProtoMessage()    {}
func (*ModifyPermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf0ffb27b25a32, []int{14}
}

func (m *ModifyPermissionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyPermissionsResponse.Unmarshal(m, b)
}
func (m *ModifyPermissionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyPermissionsResponse.Marshal(b, m, deterministic)
}
func (m *ModifyPermissionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyPermissionsResponse.Merge(m, src)
}
func (m *ModifyPermissionsResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyPermissionsResponse.Size(m)
}
func (m *ModifyPermissionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyPermissionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyPermissionsResponse proto.InternalMessageInfo

func (m *ModifyPermissionsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func init() {
	proto.RegisterEnum("Ydb.Scheme.Entry_Type", Entry_Type_name, Entry_Type_value)
	proto.RegisterType((*MakeDirectoryRequest)(nil), "Ydb.Scheme.MakeDirectoryRequest")
	proto.RegisterType((*MakeDirectoryResponse)(nil), "Ydb.Scheme.MakeDirectoryResponse")
	proto.RegisterType((*RemoveDirectoryRequest)(nil), "Ydb.Scheme.RemoveDirectoryRequest")
	proto.RegisterType((*RemoveDirectoryResponse)(nil), "Ydb.Scheme.RemoveDirectoryResponse")
	proto.RegisterType((*ListDirectoryRequest)(nil), "Ydb.Scheme.ListDirectoryRequest")
	proto.RegisterType((*ListDirectoryResponse)(nil), "Ydb.Scheme.ListDirectoryResponse")
	proto.RegisterType((*Permissions)(nil), "Ydb.Scheme.Permissions")
	proto.RegisterType((*Entry)(nil), "Ydb.Scheme.Entry")
	proto.RegisterType((*ListDirectoryResult)(nil), "Ydb.Scheme.ListDirectoryResult")
	proto.RegisterType((*DescribePathRequest)(nil), "Ydb.Scheme.DescribePathRequest")
	proto.RegisterType((*DescribePathResponse)(nil), "Ydb.Scheme.DescribePathResponse")
	proto.RegisterType((*DescribePathResult)(nil), "Ydb.Scheme.DescribePathResult")
	proto.RegisterType((*PermissionsAction)(nil), "Ydb.Scheme.PermissionsAction")
	proto.RegisterType((*ModifyPermissionsRequest)(nil), "Ydb.Scheme.ModifyPermissionsRequest")
	proto.RegisterType((*ModifyPermissionsResponse)(nil), "Ydb.Scheme.ModifyPermissionsResponse")
}

func init() { proto.RegisterFile("ydb_scheme.proto", fileDescriptor_eeaf0ffb27b25a32) }

var fileDescriptor_eeaf0ffb27b25a32 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0xdd, 0x6e, 0xeb, 0x44,
	0x10, 0xc7, 0xeb, 0xe6, 0x7b, 0x72, 0x20, 0xce, 0x9e, 0xa4, 0xc7, 0x07, 0x09, 0x11, 0x19, 0x21,
	0x05, 0x0a, 0x8e, 0x5a, 0x2e, 0x2a, 0xc4, 0x55, 0x3e, 0x16, 0x1a, 0x48, 0x62, 0x77, 0xe3, 0x20,
	0xf5, 0xca, 0x72, 0x9c, 0x4d, 0x63, 0x92, 0xd8, 0xc6, 0xeb, 0x84, 0xfa, 0x31, 0x78, 0x0b, 0xde,
	0x87, 0x17, 0xe0, 0x31, 0xb8, 0x44, 0x5e, 0xe7, 0xc3, 0x4d, 0x3f, 0x40, 0x8a, 0xd4, 0x73, 0xb7,
	0x9e, 0xfd, 0xcd, 0xcc, 0xfe, 0x67, 0x67, 0xac, 0x05, 0x31, 0x9c, 0x8c, 0x0d, 0x66, 0xcd, 0xe8,
	0x92, 0x2a, 0x9e, 0xef, 0x06, 0x2e, 0x82, 0xdb, 0xc9, 0x58, 0x19, 0x72, 0xcb, 0x27, 0x5f, 0xcf,
	0xed, 0xb9, 0xbd, 0xf4, 0x1b, 0xde, 0x6a, 0xbc, 0xb0, 0xad, 0x86, 0xe9, 0xd9, 0x0d, 0x0e, 0xb1,
	0x46, 0xe4, 0xe6, 0x7a, 0xd4, 0x37, 0x03, 0xdb, 0x75, 0x62, 0x4f, 0x79, 0x0d, 0x95, 0xbe, 0x39,
	0xa7, 0x1d, 0xdb, 0xa7, 0x56, 0xe0, 0xfa, 0x21, 0xa1, 0xbf, 0xad, 0x28, 0x0b, 0xd0, 0x4f, 0x20,
	0xee, 0x50, 0xc3, 0x33, 0x7d, 0x73, 0xc9, 0x24, 0xa1, 0x26, 0xd4, 0x8b, 0x97, 0x9f, 0x29, 0x51,
	0x32, 0x75, 0xbb, 0xc9, 0xf6, 0x4b, 0x8d, 0x63, 0xa4, 0xe4, 0x3e, 0x34, 0x20, 0x04, 0x69, 0xcf,
	0x0c, 0x66, 0xd2, 0x69, 0x4d, 0xa8, 0x17, 0x08, 0x5f, 0xcb, 0x1a, 0x54, 0x0f, 0xf2, 0x32, 0xcf,
	0x75, 0x18, 0x45, 0x57, 0x50, 0xd8, 0xf9, 0x6f, 0x32, 0xbe, 0x7f, 0x36, 0x23, 0xd9, 0xb3, 0xf2,
	0x3d, 0x9c, 0x11, 0xba, 0x74, 0xd7, 0xaf, 0xaf, 0x85, 0xc0, 0xbb, 0x47, 0x99, 0x8f, 0x55, 0xb3,
	0x86, 0x4a, 0xcf, 0x66, 0xc1, 0x87, 0xb8, 0x97, 0x83, 0xbc, 0xc7, 0x2a, 0x21, 0x50, 0xd4, 0xa8,
	0xbf, 0xb4, 0x19, 0x8b, 0x18, 0x24, 0x41, 0x8e, 0xad, 0xc6, 0xbf, 0x52, 0x2b, 0xe0, 0x51, 0x0a,
	0x64, 0xfb, 0x89, 0xbe, 0x04, 0xd1, 0xdb, 0x81, 0x86, 0x63, 0x2e, 0x29, 0x93, 0x4e, 0x6b, 0xa9,
	0x7a, 0x81, 0x94, 0xf6, 0xf6, 0x41, 0x64, 0x96, 0xff, 0x48, 0x41, 0x06, 0x3b, 0x81, 0x1f, 0x46,
	0x1a, 0x22, 0x72, 0x13, 0x8b, 0xaf, 0x51, 0x05, 0x32, 0xee, 0xef, 0x0e, 0xf5, 0x37, 0xc2, 0xe2,
	0x0f, 0xf4, 0x15, 0xa4, 0x83, 0xd0, 0xa3, 0x52, 0xa6, 0x26, 0xd4, 0x3f, 0xbe, 0x3c, 0x53, 0xf6,
	0x23, 0xa3, 0xf0, 0x50, 0x8a, 0x1e, 0x7a, 0x94, 0x70, 0x06, 0xf5, 0xa0, 0x4a, 0xa7, 0x53, 0x6a,
	0x05, 0xf6, 0x9a, 0x1a, 0xfb, 0xe4, 0x4c, 0xca, 0xd6, 0x52, 0xf5, 0xe2, 0xe5, 0xbb, 0xa4, 0x73,
	0x42, 0x1c, 0xa9, 0xec, 0xbc, 0x92, 0x92, 0xbf, 0x83, 0x62, 0x32, 0x46, 0xee, 0xe5, 0x18, 0x49,
	0x56, 0xfe, 0x53, 0x80, 0x74, 0x74, 0x2e, 0x54, 0x01, 0x51, 0xbf, 0xd5, 0xb0, 0x31, 0x1a, 0x0c,
	0x35, 0xdc, 0xee, 0xfe, 0xd0, 0xc5, 0x1d, 0xf1, 0x04, 0x7d, 0x04, 0x85, 0x4e, 0x97, 0xe0, 0xb6,
	0xae, 0x92, 0x5b, 0x51, 0x40, 0x05, 0xc8, 0xe8, 0xcd, 0x56, 0x0f, 0x8b, 0xa7, 0x11, 0xaf, 0x61,
	0x32, 0x34, 0x6e, 0x46, 0x78, 0x84, 0x8d, 0x1f, 0x89, 0x3a, 0xd2, 0xc4, 0x14, 0x7a, 0x03, 0xf9,
	0x4e, 0x53, 0x6f, 0xb6, 0x9a, 0x43, 0x2c, 0xa6, 0x51, 0x09, 0x8a, 0x44, 0xef, 0x13, 0xe3, 0x17,
	0xb5, 0x37, 0xea, 0x63, 0x31, 0x83, 0xce, 0x00, 0xb5, 0x7a, 0x6a, 0xfb, 0x67, 0x63, 0xa8, 0xab,
	0x04, 0x6f, 0xed, 0x59, 0x54, 0x85, 0x72, 0x5b, 0x55, 0x49, 0xa7, 0x3b, 0x68, 0xea, 0x5d, 0x75,
	0x60, 0x0c, 0xd4, 0x0e, 0x16, 0x73, 0x51, 0xb4, 0x21, 0xbe, 0x19, 0xe1, 0x41, 0x1b, 0x8b, 0x25,
	0x79, 0x0e, 0x6f, 0x0f, 0x3b, 0x67, 0xb5, 0x08, 0xd0, 0x17, 0x90, 0x66, 0x74, 0x31, 0xdd, 0xb4,
	0x4c, 0xf9, 0x51, 0xd9, 0x09, 0xdf, 0x46, 0xdf, 0x40, 0xde, 0x9a, 0xd9, 0x8b, 0x89, 0x4f, 0x1d,
	0x7e, 0xe9, 0x4f, 0xa2, 0x3b, 0x44, 0x5e, 0xc1, 0xdb, 0x0e, 0x65, 0x96, 0x6f, 0x8f, 0xa9, 0x66,
	0x06, 0xb3, 0xd7, 0x9a, 0x0e, 0x15, 0x2a, 0x0f, 0xd3, 0x1e, 0x3b, 0x1c, 0xdf, 0x03, 0x3a, 0x08,
	0xf8, 0xff, 0x6b, 0x26, 0xff, 0x25, 0x40, 0x39, 0xd1, 0x39, 0x4d, 0x2b, 0x0a, 0x89, 0x1a, 0x90,
	0xb9, 0xf3, 0x4d, 0x27, 0xd8, 0x78, 0x3f, 0xd7, 0x67, 0xd7, 0x27, 0x24, 0xe6, 0xd0, 0x05, 0x64,
	0x7d, 0xba, 0x76, 0xe7, 0x94, 0x4b, 0x7d, 0xd1, 0x63, 0x03, 0xa2, 0x73, 0x48, 0x31, 0x1a, 0x48,
	0xa9, 0xff, 0xe2, 0x23, 0x0a, 0x7d, 0x0e, 0x6f, 0xac, 0x99, 0xe9, 0xdc, 0x51, 0x23, 0x9e, 0xca,
	0x74, 0x54, 0xd0, 0xeb, 0x13, 0x52, 0x8c, 0xad, 0x6a, 0x64, 0x6c, 0xe5, 0x21, 0x6b, 0xf2, 0xf3,
	0xcb, 0x7f, 0x0b, 0x20, 0xf5, 0xdd, 0x89, 0x3d, 0x0d, 0x93, 0x53, 0xf1, 0x3a, 0x17, 0x8c, 0xae,
	0x20, 0x17, 0x1f, 0x83, 0x49, 0x29, 0xde, 0x85, 0x9f, 0x3e, 0x23, 0x2e, 0x2e, 0x36, 0xd9, 0xd2,
	0xe8, 0x1c, 0xca, 0xd6, 0x82, 0x9a, 0xfe, 0x83, 0xbf, 0x45, 0xa4, 0x34, 0x4f, 0x44, 0xbe, 0x91,
	0xf0, 0x95, 0x75, 0x78, 0xff, 0x84, 0xc2, 0x23, 0x7b, 0xa9, 0x75, 0x01, 0x55, 0xcb, 0x5d, 0x2a,
	0xa1, 0xe9, 0x4c, 0xe8, 0xbd, 0x12, 0x4e, 0xc6, 0x4a, 0xfc, 0x46, 0x68, 0x55, 0xe3, 0xe3, 0xef,
	0x0b, 0xc2, 0x9f, 0x03, 0xff, 0x08, 0xc2, 0x38, 0xcb, 0x1f, 0x01, 0xdf, 0xfe, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xbf, 0x90, 0xe3, 0x05, 0x52, 0x08, 0x00, 0x00,
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *MakeDirectoryRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *RemoveDirectoryRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *ListDirectoryRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *DescribePathRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *ModifyPermissionsRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}
