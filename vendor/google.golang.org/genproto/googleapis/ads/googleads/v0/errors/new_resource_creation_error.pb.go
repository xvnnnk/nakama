// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/new_resource_creation_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum_NewResourceCreationError int32

const (
	// Enum unspecified.
	NewResourceCreationErrorEnum_UNSPECIFIED NewResourceCreationErrorEnum_NewResourceCreationError = 0
	// The received error code is not known in this version.
	NewResourceCreationErrorEnum_UNKNOWN NewResourceCreationErrorEnum_NewResourceCreationError = 1
	// Do not set the id field while creating new resources.
	NewResourceCreationErrorEnum_CANNOT_SET_ID_FOR_CREATE NewResourceCreationErrorEnum_NewResourceCreationError = 2
	// Creating more than one resource with the same temp ID is not allowed.
	NewResourceCreationErrorEnum_DUPLICATE_TEMP_IDS NewResourceCreationErrorEnum_NewResourceCreationError = 3
	// Parent resource with specified temp ID failed validation, so no
	// validation will be done for this child resource.
	NewResourceCreationErrorEnum_TEMP_ID_RESOURCE_HAD_ERRORS NewResourceCreationErrorEnum_NewResourceCreationError = 4
)

var NewResourceCreationErrorEnum_NewResourceCreationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_SET_ID_FOR_CREATE",
	3: "DUPLICATE_TEMP_IDS",
	4: "TEMP_ID_RESOURCE_HAD_ERRORS",
}
var NewResourceCreationErrorEnum_NewResourceCreationError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"CANNOT_SET_ID_FOR_CREATE":    2,
	"DUPLICATE_TEMP_IDS":          3,
	"TEMP_ID_RESOURCE_HAD_ERRORS": 4,
}

func (x NewResourceCreationErrorEnum_NewResourceCreationError) String() string {
	return proto.EnumName(NewResourceCreationErrorEnum_NewResourceCreationError_name, int32(x))
}
func (NewResourceCreationErrorEnum_NewResourceCreationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_new_resource_creation_error_1adcb44311ee3769, []int{0, 0}
}

// Container for enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResourceCreationErrorEnum) Reset()         { *m = NewResourceCreationErrorEnum{} }
func (m *NewResourceCreationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NewResourceCreationErrorEnum) ProtoMessage()    {}
func (*NewResourceCreationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_new_resource_creation_error_1adcb44311ee3769, []int{0}
}
func (m *NewResourceCreationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Unmarshal(m, b)
}
func (m *NewResourceCreationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *NewResourceCreationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResourceCreationErrorEnum.Merge(dst, src)
}
func (m *NewResourceCreationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Size(m)
}
func (m *NewResourceCreationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResourceCreationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NewResourceCreationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewResourceCreationErrorEnum)(nil), "google.ads.googleads.v0.errors.NewResourceCreationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.NewResourceCreationErrorEnum_NewResourceCreationError", NewResourceCreationErrorEnum_NewResourceCreationError_name, NewResourceCreationErrorEnum_NewResourceCreationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/new_resource_creation_error.proto", fileDescriptor_new_resource_creation_error_1adcb44311ee3769)
}

var fileDescriptor_new_resource_creation_error_1adcb44311ee3769 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0x7f, 0xdd, 0x7e, 0x28, 0x64, 0x07, 0x4b, 0x0e, 0x32, 0x70, 0x4e, 0xd8, 0x0b, 0x48,
	0x0b, 0x1e, 0xbd, 0x98, 0xb5, 0xd9, 0x2c, 0x6a, 0x5a, 0xd2, 0x76, 0x82, 0x14, 0x42, 0x5d, 0x43,
	0x19, 0x6c, 0xcd, 0x48, 0xf6, 0xe7, 0x65, 0xf8, 0x1e, 0x3c, 0x7a, 0xf2, 0x75, 0x78, 0xf5, 0x0d,
	0x49, 0x9b, 0x6e, 0xb7, 0x79, 0xca, 0x13, 0xbe, 0x4f, 0x3e, 0x4f, 0xbe, 0x0f, 0xb8, 0x2f, 0xa5,
	0x2c, 0x97, 0xc2, 0xc9, 0x0b, 0xed, 0x18, 0x59, 0xab, 0x9d, 0xeb, 0x08, 0xa5, 0xa4, 0xd2, 0x4e,
	0x25, 0xf6, 0x5c, 0x09, 0x2d, 0xb7, 0x6a, 0x2e, 0xf8, 0x5c, 0x89, 0x7c, 0xb3, 0x90, 0x15, 0x6f,
	0x86, 0x68, 0xad, 0xe4, 0x46, 0xc2, 0xa1, 0x79, 0x86, 0xf2, 0x42, 0xa3, 0x23, 0x01, 0xed, 0x5c,
	0x64, 0x08, 0xa3, 0x2f, 0x0b, 0x0c, 0xa8, 0xd8, 0xb3, 0x16, 0xe2, 0xb5, 0x0c, 0x52, 0x4f, 0x49,
	0xb5, 0x5d, 0x8d, 0xde, 0x2d, 0xd0, 0x3f, 0x65, 0x80, 0x17, 0xa0, 0x97, 0xd2, 0x38, 0x22, 0x5e,
	0x30, 0x09, 0x88, 0x6f, 0xff, 0x83, 0x3d, 0x70, 0x9e, 0xd2, 0x47, 0x1a, 0xbe, 0x50, 0xdb, 0x82,
	0x03, 0xd0, 0xf7, 0x30, 0xa5, 0x61, 0xc2, 0x63, 0x92, 0xf0, 0xc0, 0xe7, 0x93, 0x90, 0x71, 0x8f,
	0x11, 0x9c, 0x10, 0xbb, 0x03, 0x2f, 0x01, 0xf4, 0xd3, 0xe8, 0x29, 0xf0, 0x70, 0x42, 0x78, 0x42,
	0x9e, 0x23, 0x1e, 0xf8, 0xb1, 0xdd, 0x85, 0x37, 0xe0, 0xaa, 0xbd, 0x71, 0x46, 0xe2, 0x30, 0x65,
	0x1e, 0xe1, 0x0f, 0xd8, 0xe7, 0x84, 0xb1, 0x90, 0xc5, 0xf6, 0xff, 0xf1, 0x8f, 0x05, 0x46, 0x73,
	0xb9, 0x42, 0x7f, 0x6f, 0x36, 0xbe, 0x3e, 0xf5, 0xeb, 0xa8, 0x2e, 0x26, 0xb2, 0x5e, 0xfd, 0x16,
	0x50, 0xca, 0x65, 0x5e, 0x95, 0x48, 0xaa, 0xd2, 0x29, 0x45, 0xd5, 0xd4, 0x76, 0x28, 0x7b, 0xbd,
	0xd0, 0xa7, 0xba, 0xbf, 0x33, 0xc7, 0x47, 0xa7, 0x3b, 0xc5, 0xf8, 0xb3, 0x33, 0x9c, 0x1a, 0x18,
	0x2e, 0x34, 0x32, 0xb2, 0x56, 0x33, 0x17, 0x35, 0x91, 0xfa, 0xfb, 0x60, 0xc8, 0x70, 0xa1, 0xb3,
	0xa3, 0x21, 0x9b, 0xb9, 0x99, 0x31, 0xbc, 0x9d, 0x35, 0xc1, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x60, 0xb1, 0xef, 0x4e, 0xf3, 0x01, 0x00, 0x00,
}