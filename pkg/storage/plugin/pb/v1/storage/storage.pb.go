// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/storage/storage.proto

package stpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ListRequest gives a module name to get list of versions in storage.
type ListRequest struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

// ListResponse gives a list of versions for the given module.
type ListResponse struct {
	Versions             []string `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetVersions() []string {
	if m != nil {
		return m.Versions
	}
	return nil
}

// GetModuleRequest defines a module and a version to fetch from storage.
type GetModuleRequest struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetModuleRequest) Reset()         { *m = GetModuleRequest{} }
func (m *GetModuleRequest) String() string { return proto.CompactTextString(m) }
func (*GetModuleRequest) ProtoMessage()    {}
func (*GetModuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{2}
}

func (m *GetModuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModuleRequest.Unmarshal(m, b)
}
func (m *GetModuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModuleRequest.Marshal(b, m, deterministic)
}
func (m *GetModuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModuleRequest.Merge(m, src)
}
func (m *GetModuleRequest) XXX_Size() int {
	return xxx_messageInfo_GetModuleRequest.Size(m)
}
func (m *GetModuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModuleRequest proto.InternalMessageInfo

func (m *GetModuleRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *GetModuleRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// GetModuleResponse hold all or part of the requested file.
type GetModuleResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetModuleResponse) Reset()         { *m = GetModuleResponse{} }
func (m *GetModuleResponse) String() string { return proto.CompactTextString(m) }
func (*GetModuleResponse) ProtoMessage()    {}
func (*GetModuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{3}
}

func (m *GetModuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModuleResponse.Unmarshal(m, b)
}
func (m *GetModuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModuleResponse.Marshal(b, m, deterministic)
}
func (m *GetModuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModuleResponse.Merge(m, src)
}
func (m *GetModuleResponse) XXX_Size() int {
	return xxx_messageInfo_GetModuleResponse.Size(m)
}
func (m *GetModuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetModuleResponse proto.InternalMessageInfo

func (m *GetModuleResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// ExistsRequest gives a module to check for in storage.
type ExistsRequest struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistsRequest) Reset()         { *m = ExistsRequest{} }
func (m *ExistsRequest) String() string { return proto.CompactTextString(m) }
func (*ExistsRequest) ProtoMessage()    {}
func (*ExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{4}
}

func (m *ExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistsRequest.Unmarshal(m, b)
}
func (m *ExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistsRequest.Marshal(b, m, deterministic)
}
func (m *ExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistsRequest.Merge(m, src)
}
func (m *ExistsRequest) XXX_Size() int {
	return xxx_messageInfo_ExistsRequest.Size(m)
}
func (m *ExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExistsRequest proto.InternalMessageInfo

func (m *ExistsRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ExistsRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// ExistsResponse reports if the module version exists in storage.
type ExistsResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistsResponse) Reset()         { *m = ExistsResponse{} }
func (m *ExistsResponse) String() string { return proto.CompactTextString(m) }
func (*ExistsResponse) ProtoMessage()    {}
func (*ExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{5}
}

func (m *ExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistsResponse.Unmarshal(m, b)
}
func (m *ExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistsResponse.Marshal(b, m, deterministic)
}
func (m *ExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistsResponse.Merge(m, src)
}
func (m *ExistsResponse) XXX_Size() int {
	return xxx_messageInfo_ExistsResponse.Size(m)
}
func (m *ExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExistsResponse proto.InternalMessageInfo

func (m *ExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

// SaveRequest gives a module definition and resources for saving a file to storage.
type SaveRequest struct {
	ModDefinition        *SaveRequest_Module `protobuf:"bytes,1,opt,name=mod_definition,json=modDefinition,proto3" json:"mod_definition,omitempty"`
	Zip                  []byte              `protobuf:"bytes,2,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{6}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetModDefinition() *SaveRequest_Module {
	if m != nil {
		return m.ModDefinition
	}
	return nil
}

func (m *SaveRequest) GetZip() []byte {
	if m != nil {
		return m.Zip
	}
	return nil
}

type SaveRequest_Module struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Mod                  []byte   `protobuf:"bytes,3,opt,name=mod,proto3" json:"mod,omitempty"`
	Info                 []byte   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveRequest_Module) Reset()         { *m = SaveRequest_Module{} }
func (m *SaveRequest_Module) String() string { return proto.CompactTextString(m) }
func (*SaveRequest_Module) ProtoMessage()    {}
func (*SaveRequest_Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{6, 0}
}

func (m *SaveRequest_Module) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest_Module.Unmarshal(m, b)
}
func (m *SaveRequest_Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest_Module.Marshal(b, m, deterministic)
}
func (m *SaveRequest_Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest_Module.Merge(m, src)
}
func (m *SaveRequest_Module) XXX_Size() int {
	return xxx_messageInfo_SaveRequest_Module.Size(m)
}
func (m *SaveRequest_Module) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest_Module.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest_Module proto.InternalMessageInfo

func (m *SaveRequest_Module) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *SaveRequest_Module) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SaveRequest_Module) GetMod() []byte {
	if m != nil {
		return m.Mod
	}
	return nil
}

func (m *SaveRequest_Module) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

// SaveResponse reports the success of saving to storage.
type SaveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{7}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

// DeleteRequest gives a module definition to be removed from storage.
type DeleteRequest struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *DeleteRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// DeleteResponse reports the success of deleting from storage.
type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e551bf102a896df2, []int{9}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListRequest)(nil), "athens.v1.storage.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "athens.v1.storage.ListResponse")
	proto.RegisterType((*GetModuleRequest)(nil), "athens.v1.storage.GetModuleRequest")
	proto.RegisterType((*GetModuleResponse)(nil), "athens.v1.storage.GetModuleResponse")
	proto.RegisterType((*ExistsRequest)(nil), "athens.v1.storage.ExistsRequest")
	proto.RegisterType((*ExistsResponse)(nil), "athens.v1.storage.ExistsResponse")
	proto.RegisterType((*SaveRequest)(nil), "athens.v1.storage.SaveRequest")
	proto.RegisterType((*SaveRequest_Module)(nil), "athens.v1.storage.SaveRequest.Module")
	proto.RegisterType((*SaveResponse)(nil), "athens.v1.storage.SaveResponse")
	proto.RegisterType((*DeleteRequest)(nil), "athens.v1.storage.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "athens.v1.storage.DeleteResponse")
}

func init() { proto.RegisterFile("v1/storage/storage.proto", fileDescriptor_e551bf102a896df2) }

var fileDescriptor_e551bf102a896df2 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x15, 0x1a, 0x65, 0xdb, 0x6b, 0x57, 0x75, 0x96, 0x98, 0xa2, 0x1c, 0xa0, 0x04, 0x26,
	0x2a, 0x0e, 0x09, 0x1b, 0x47, 0x24, 0x04, 0x53, 0x51, 0x35, 0x69, 0xbb, 0xa4, 0x70, 0xd9, 0x05,
	0xd2, 0xfa, 0x35, 0xb3, 0xd6, 0xc4, 0xa1, 0x76, 0x23, 0xc4, 0xbf, 0xc8, 0x81, 0x7f, 0x09, 0xf9,
	0x47, 0x43, 0x0a, 0x61, 0x08, 0xe8, 0x29, 0x7e, 0xf6, 0x7b, 0x1f, 0xbf, 0x1f, 0xdf, 0x18, 0xfc,
	0xea, 0x34, 0x16, 0x92, 0xaf, 0xd2, 0x0c, 0x37, 0xdf, 0xa8, 0x5c, 0x71, 0xc9, 0xc9, 0x51, 0x2a,
	0x6f, 0xb0, 0x10, 0x51, 0x75, 0x1a, 0xd9, 0x83, 0xf0, 0x04, 0xba, 0x97, 0x4c, 0xc8, 0x04, 0x3f,
	0xad, 0x51, 0x48, 0x72, 0x0c, 0x5e, 0xce, 0xe9, 0x7a, 0x89, 0xbe, 0x33, 0x74, 0x46, 0x07, 0x89,
	0xb5, 0xc2, 0x67, 0xd0, 0x33, 0x6e, 0xa2, 0xe4, 0x85, 0x40, 0x12, 0xc0, 0x7e, 0x85, 0x2b, 0xc1,
	0x78, 0x21, 0x7c, 0x67, 0xd8, 0x19, 0x1d, 0x24, 0xb5, 0x1d, 0x8e, 0x61, 0x30, 0x41, 0x79, 0xa5,
	0x03, 0xff, 0xc0, 0x25, 0x3e, 0xec, 0xd9, 0x38, 0xff, 0x9e, 0x3e, 0xd8, 0x98, 0xe1, 0x53, 0x38,
	0x6a, 0x50, 0xec, 0xb5, 0x04, 0x5c, 0x9a, 0xca, 0x54, 0x43, 0x7a, 0x89, 0x5e, 0x87, 0x6f, 0xe0,
	0xf0, 0xed, 0x67, 0x26, 0xa4, 0xf8, 0xf7, 0xbb, 0x46, 0xd0, 0xdf, 0x20, 0xec, 0x45, 0xc7, 0xe0,
	0xa1, 0xde, 0xd1, 0x8c, 0xfd, 0xc4, 0x5a, 0xe1, 0x37, 0x07, 0xba, 0xd3, 0xb4, 0xaa, 0xeb, 0xba,
	0x84, 0x7e, 0xce, 0xe9, 0x07, 0x8a, 0x0b, 0x56, 0x30, 0xa9, 0xd0, 0xca, 0xbf, 0x7b, 0x76, 0x12,
	0xfd, 0xd2, 0xea, 0xa8, 0x11, 0x17, 0xd9, 0xba, 0x0e, 0x73, 0x4e, 0xc7, 0x75, 0x2c, 0x19, 0x40,
	0xe7, 0x0b, 0x2b, 0x75, 0x76, 0xbd, 0x44, 0x2d, 0x83, 0x8f, 0xe0, 0x19, 0xd7, 0xbf, 0xaf, 0x4a,
	0xd1, 0x72, 0x4e, 0xfd, 0x8e, 0xa1, 0xe5, 0x9c, 0xaa, 0xf6, 0xb1, 0x62, 0xc1, 0x7d, 0xd7, 0xb4,
	0x4f, 0xad, 0xc3, 0x3e, 0xf4, 0x4c, 0x62, 0xa6, 0x72, 0xd5, 0xce, 0x31, 0x2e, 0x51, 0xfe, 0xc7,
	0xe8, 0x06, 0xd0, 0xdf, 0x20, 0x0c, 0xf4, 0xec, 0xab, 0x0b, 0xf7, 0xa7, 0xa6, 0x0d, 0xe7, 0xe9,
	0xfc, 0x16, 0x0b, 0x3a, 0xc5, 0x55, 0xc5, 0xe6, 0x48, 0x26, 0xe0, 0x2a, 0x61, 0x91, 0x07, 0x2d,
	0x0d, 0x6b, 0x08, 0x33, 0x78, 0xf8, 0xdb, 0x73, 0x3b, 0xb1, 0x77, 0xb0, 0x37, 0x41, 0x79, 0x51,
	0x2c, 0x38, 0x79, 0xdc, 0xe2, 0xfb, 0xb3, 0x22, 0x83, 0x27, 0x77, 0x3b, 0x59, 0xea, 0x14, 0x3c,
	0xb3, 0xb9, 0x4b, 0xe8, 0x7b, 0x0d, 0xbd, 0x66, 0xe5, 0x0e, 0xa1, 0xcf, 0x1d, 0x72, 0x05, 0x9e,
	0x51, 0x31, 0x19, 0xb6, 0x44, 0x6c, 0xfd, 0x23, 0xc1, 0xa3, 0x3b, 0x3c, 0x6c, 0x96, 0x17, 0xe0,
	0x2a, 0x61, 0xb4, 0x4e, 0xa6, 0x21, 0xe5, 0xd6, 0xc9, 0x34, 0x15, 0x35, 0xd2, 0x99, 0x19, 0x41,
	0xb4, 0x66, 0xb6, 0x25, 0xb7, 0xd6, 0xcc, 0xb6, 0xd5, 0x74, 0xfe, 0xfa, 0xfa, 0x55, 0xc6, 0xe4,
	0xcd, 0x7a, 0x16, 0xcd, 0x79, 0x1e, 0x67, 0x3c, 0xe7, 0x54, 0xc4, 0x26, 0x2a, 0x2e, 0x6f, 0xb3,
	0xfa, 0xf1, 0x2b, 0x97, 0xeb, 0x8c, 0x15, 0x71, 0x39, 0x8b, 0x7f, 0xbc, 0x88, 0x2f, 0x85, 0x2c,
	0x67, 0x33, 0x4f, 0xbf, 0x87, 0x2f, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x5b, 0x0d, 0xe4,
	0x2b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageBackendServiceClient is the client API for StorageBackendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageBackendServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	GetInfo(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error)
	GetMod(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error)
	GetZip(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (StorageBackendService_GetZipClient, error)
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Save(ctx context.Context, opts ...grpc.CallOption) (StorageBackendService_SaveClient, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type storageBackendServiceClient struct {
	cc *grpc.ClientConn
}

func NewStorageBackendServiceClient(cc *grpc.ClientConn) StorageBackendServiceClient {
	return &storageBackendServiceClient{cc}
}

func (c *storageBackendServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/athens.v1.storage.StorageBackendService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBackendServiceClient) GetInfo(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error) {
	out := new(GetModuleResponse)
	err := c.cc.Invoke(ctx, "/athens.v1.storage.StorageBackendService/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBackendServiceClient) GetMod(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error) {
	out := new(GetModuleResponse)
	err := c.cc.Invoke(ctx, "/athens.v1.storage.StorageBackendService/GetMod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBackendServiceClient) GetZip(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (StorageBackendService_GetZipClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StorageBackendService_serviceDesc.Streams[0], "/athens.v1.storage.StorageBackendService/GetZip", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageBackendServiceGetZipClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StorageBackendService_GetZipClient interface {
	Recv() (*GetModuleResponse, error)
	grpc.ClientStream
}

type storageBackendServiceGetZipClient struct {
	grpc.ClientStream
}

func (x *storageBackendServiceGetZipClient) Recv() (*GetModuleResponse, error) {
	m := new(GetModuleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageBackendServiceClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	out := new(ExistsResponse)
	err := c.cc.Invoke(ctx, "/athens.v1.storage.StorageBackendService/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBackendServiceClient) Save(ctx context.Context, opts ...grpc.CallOption) (StorageBackendService_SaveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StorageBackendService_serviceDesc.Streams[1], "/athens.v1.storage.StorageBackendService/Save", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageBackendServiceSaveClient{stream}
	return x, nil
}

type StorageBackendService_SaveClient interface {
	Send(*SaveRequest) error
	CloseAndRecv() (*SaveResponse, error)
	grpc.ClientStream
}

type storageBackendServiceSaveClient struct {
	grpc.ClientStream
}

func (x *storageBackendServiceSaveClient) Send(m *SaveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageBackendServiceSaveClient) CloseAndRecv() (*SaveResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SaveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageBackendServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/athens.v1.storage.StorageBackendService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageBackendServiceServer is the server API for StorageBackendService service.
type StorageBackendServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	GetInfo(context.Context, *GetModuleRequest) (*GetModuleResponse, error)
	GetMod(context.Context, *GetModuleRequest) (*GetModuleResponse, error)
	GetZip(*GetModuleRequest, StorageBackendService_GetZipServer) error
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Save(StorageBackendService_SaveServer) error
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedStorageBackendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStorageBackendServiceServer struct {
}

func (*UnimplementedStorageBackendServiceServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStorageBackendServiceServer) GetInfo(ctx context.Context, req *GetModuleRequest) (*GetModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedStorageBackendServiceServer) GetMod(ctx context.Context, req *GetModuleRequest) (*GetModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMod not implemented")
}
func (*UnimplementedStorageBackendServiceServer) GetZip(req *GetModuleRequest, srv StorageBackendService_GetZipServer) error {
	return status.Errorf(codes.Unimplemented, "method GetZip not implemented")
}
func (*UnimplementedStorageBackendServiceServer) Exists(ctx context.Context, req *ExistsRequest) (*ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (*UnimplementedStorageBackendServiceServer) Save(srv StorageBackendService_SaveServer) error {
	return status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (*UnimplementedStorageBackendServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterStorageBackendServiceServer(s *grpc.Server, srv StorageBackendServiceServer) {
	s.RegisterService(&_StorageBackendService_serviceDesc, srv)
}

func _StorageBackendService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBackendServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athens.v1.storage.StorageBackendService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBackendServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBackendService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBackendServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athens.v1.storage.StorageBackendService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBackendServiceServer).GetInfo(ctx, req.(*GetModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBackendService_GetMod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBackendServiceServer).GetMod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athens.v1.storage.StorageBackendService/GetMod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBackendServiceServer).GetMod(ctx, req.(*GetModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBackendService_GetZip_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetModuleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageBackendServiceServer).GetZip(m, &storageBackendServiceGetZipServer{stream})
}

type StorageBackendService_GetZipServer interface {
	Send(*GetModuleResponse) error
	grpc.ServerStream
}

type storageBackendServiceGetZipServer struct {
	grpc.ServerStream
}

func (x *storageBackendServiceGetZipServer) Send(m *GetModuleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StorageBackendService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBackendServiceServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athens.v1.storage.StorageBackendService/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBackendServiceServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBackendService_Save_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageBackendServiceServer).Save(&storageBackendServiceSaveServer{stream})
}

type StorageBackendService_SaveServer interface {
	SendAndClose(*SaveResponse) error
	Recv() (*SaveRequest, error)
	grpc.ServerStream
}

type storageBackendServiceSaveServer struct {
	grpc.ServerStream
}

func (x *storageBackendServiceSaveServer) SendAndClose(m *SaveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageBackendServiceSaveServer) Recv() (*SaveRequest, error) {
	m := new(SaveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StorageBackendService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBackendServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athens.v1.storage.StorageBackendService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBackendServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageBackendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "athens.v1.storage.StorageBackendService",
	HandlerType: (*StorageBackendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _StorageBackendService_List_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _StorageBackendService_GetInfo_Handler,
		},
		{
			MethodName: "GetMod",
			Handler:    _StorageBackendService_GetMod_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _StorageBackendService_Exists_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StorageBackendService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetZip",
			Handler:       _StorageBackendService_GetZip_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Save",
			Handler:       _StorageBackendService_Save_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "v1/storage/storage.proto",
}
