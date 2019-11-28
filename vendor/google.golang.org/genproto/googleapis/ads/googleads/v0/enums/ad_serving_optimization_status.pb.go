// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_serving_optimization_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enum describing possible serving statuses.
type AdServingOptimizationStatusEnum_AdServingOptimizationStatus int32

const (
	// No value has been specified.
	AdServingOptimizationStatusEnum_UNSPECIFIED AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdServingOptimizationStatusEnum_UNKNOWN AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 1
	// Ad serving is optimized based on CTR for the campaign.
	AdServingOptimizationStatusEnum_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 2
	// Ad serving is optimized based on CTR * Conversion for the campaign. If
	// the campaign is not in the conversion optimizer bidding strategy, it will
	// default to OPTIMIZED.
	AdServingOptimizationStatusEnum_CONVERSION_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 3
	// Ads are rotated evenly for 90 days, then optimized for clicks.
	AdServingOptimizationStatusEnum_ROTATE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 4
	// Show lower performing ads more evenly with higher performing ads, and do
	// not optimize.
	AdServingOptimizationStatusEnum_ROTATE_INDEFINITELY AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 5
	// Ad serving optimization status is not available.
	AdServingOptimizationStatusEnum_UNAVAILABLE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 6
)

var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "CONVERSION_OPTIMIZE",
	4: "ROTATE",
	5: "ROTATE_INDEFINITELY",
	6: "UNAVAILABLE",
}
var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"OPTIMIZE":            2,
	"CONVERSION_OPTIMIZE": 3,
	"ROTATE":              4,
	"ROTATE_INDEFINITELY": 5,
	"UNAVAILABLE":         6,
}

func (x AdServingOptimizationStatusEnum_AdServingOptimizationStatus) String() string {
	return proto.EnumName(AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, int32(x))
}
func (AdServingOptimizationStatusEnum_AdServingOptimizationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_serving_optimization_status_28618b0ffcd14ece, []int{0, 0}
}

// Possible ad serving statuses of a campaign.
type AdServingOptimizationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdServingOptimizationStatusEnum) Reset()         { *m = AdServingOptimizationStatusEnum{} }
func (m *AdServingOptimizationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdServingOptimizationStatusEnum) ProtoMessage()    {}
func (*AdServingOptimizationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_serving_optimization_status_28618b0ffcd14ece, []int{0}
}
func (m *AdServingOptimizationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Unmarshal(m, b)
}
func (m *AdServingOptimizationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Marshal(b, m, deterministic)
}
func (dst *AdServingOptimizationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdServingOptimizationStatusEnum.Merge(dst, src)
}
func (m *AdServingOptimizationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Size(m)
}
func (m *AdServingOptimizationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdServingOptimizationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdServingOptimizationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdServingOptimizationStatusEnum)(nil), "google.ads.googleads.v0.enums.AdServingOptimizationStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdServingOptimizationStatusEnum_AdServingOptimizationStatus", AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_serving_optimization_status.proto", fileDescriptor_ad_serving_optimization_status_28618b0ffcd14ece)
}

var fileDescriptor_ad_serving_optimization_status_28618b0ffcd14ece = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x18, 0x85, 0x6f, 0xd2, 0x7b, 0x7b, 0x91, 0x8b, 0x44, 0x14, 0x06, 0x06, 0x54, 0x41, 0xfb, 0x00,
	0x4e, 0x24, 0x36, 0x33, 0x39, 0xad, 0x5b, 0x59, 0x14, 0x27, 0x6a, 0xd2, 0x20, 0xaa, 0x48, 0x51,
	0x20, 0x51, 0x14, 0xa9, 0x89, 0xab, 0x3a, 0xed, 0xc0, 0x93, 0x30, 0x33, 0xf2, 0x0c, 0x3c, 0x01,
	0x8f, 0xc2, 0xce, 0x8e, 0x62, 0xd3, 0x88, 0x85, 0x2e, 0xd6, 0x91, 0xcf, 0xf1, 0xe7, 0xdf, 0xc7,
	0xc0, 0xc9, 0x39, 0xcf, 0x57, 0x99, 0x95, 0xa4, 0xc2, 0x52, 0xb2, 0x51, 0x3b, 0xdb, 0xca, 0xaa,
	0x6d, 0x29, 0xac, 0x24, 0x8d, 0x45, 0xb6, 0xd9, 0x15, 0x55, 0x1e, 0xf3, 0x75, 0x5d, 0x94, 0xc5,
	0x53, 0x52, 0x17, 0xbc, 0x8a, 0x45, 0x9d, 0xd4, 0x5b, 0x01, 0xd7, 0x1b, 0x5e, 0x73, 0xb3, 0xaf,
	0x0e, 0xc2, 0x24, 0x15, 0xb0, 0x65, 0xc0, 0x9d, 0x0d, 0x25, 0x63, 0xf8, 0xa6, 0x81, 0x0b, 0x9c,
	0xfa, 0x0a, 0xe3, 0xfe, 0xa0, 0xf8, 0x12, 0x42, 0xaa, 0x6d, 0x39, 0x7c, 0xd6, 0xc0, 0xf9, 0x81,
	0x8c, 0x79, 0x02, 0x7a, 0x0b, 0xe6, 0x7b, 0x64, 0x44, 0x27, 0x94, 0x8c, 0x8d, 0x3f, 0x66, 0x0f,
	0xfc, 0x5f, 0xb0, 0x1b, 0xe6, 0xde, 0x31, 0x43, 0x33, 0x8f, 0xc1, 0x91, 0xeb, 0x05, 0xf4, 0x96,
	0x2e, 0x89, 0xa1, 0x9b, 0x67, 0xe0, 0x74, 0xe4, 0xb2, 0x90, 0xcc, 0x7d, 0xea, 0xb2, 0xb8, 0x35,
	0x3a, 0x26, 0x00, 0xdd, 0xb9, 0x1b, 0xe0, 0x80, 0x18, 0x7f, 0x9b, 0x90, 0xd2, 0x31, 0x65, 0x63,
	0x32, 0xa1, 0x8c, 0x06, 0x64, 0x76, 0x6f, 0xfc, 0x53, 0x37, 0xe1, 0x10, 0xd3, 0x19, 0x76, 0x66,
	0xc4, 0xe8, 0x3a, 0x9f, 0x1a, 0x18, 0x3c, 0xf2, 0x12, 0x1e, 0x7c, 0xa4, 0x73, 0x79, 0x60, 0x7a,
	0xaf, 0x69, 0xc9, 0xd3, 0x96, 0xdf, 0x5d, 0xc3, 0x9c, 0xaf, 0x92, 0x2a, 0x87, 0x7c, 0x93, 0x5b,
	0x79, 0x56, 0xc9, 0x0e, 0xf7, 0xdd, 0xaf, 0x0b, 0xf1, 0xcb, 0x57, 0x5c, 0xcb, 0xf5, 0x45, 0xef,
	0x4c, 0x31, 0x7e, 0xd5, 0xfb, 0x53, 0x85, 0xc2, 0xa9, 0x80, 0x4a, 0x36, 0x2a, 0xb4, 0x61, 0xd3,
	0xa6, 0x78, 0xdf, 0xfb, 0x11, 0x4e, 0x45, 0xd4, 0xfa, 0x51, 0x68, 0x47, 0xd2, 0xff, 0xd0, 0x07,
	0x6a, 0x13, 0x21, 0x9c, 0x0a, 0x84, 0xda, 0x04, 0x42, 0xa1, 0x8d, 0x90, 0xcc, 0x3c, 0x74, 0xe5,
	0x60, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0x32, 0x58, 0xc6, 0x22, 0x02, 0x00, 0x00,
}
