// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/ad_group_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum listing the possible types of an ad group.
type AdGroupTypeEnum_AdGroupType int32

const (
	// The type has not been specified.
	AdGroupTypeEnum_UNSPECIFIED AdGroupTypeEnum_AdGroupType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupTypeEnum_UNKNOWN AdGroupTypeEnum_AdGroupType = 1
	// The default ad group type for Search campaigns.
	AdGroupTypeEnum_SEARCH_STANDARD AdGroupTypeEnum_AdGroupType = 2
	// The default ad group type for Display campaigns.
	AdGroupTypeEnum_DISPLAY_STANDARD AdGroupTypeEnum_AdGroupType = 3
	// The ad group type for Shopping campaigns serving standard product ads.
	AdGroupTypeEnum_SHOPPING_PRODUCT_ADS AdGroupTypeEnum_AdGroupType = 4
	// The default ad group type for Hotel campaigns.
	AdGroupTypeEnum_HOTEL_ADS AdGroupTypeEnum_AdGroupType = 6
	// The type for ad groups in Smart Shopping campaigns.
	AdGroupTypeEnum_SHOPPING_SMART_ADS AdGroupTypeEnum_AdGroupType = 7
	// Short unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_BUMPER AdGroupTypeEnum_AdGroupType = 8
	// TrueView (skippable) in-stream video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_STREAM AdGroupTypeEnum_AdGroupType = 9
	// TrueView in-display video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_DISPLAY AdGroupTypeEnum_AdGroupType = 10
	// Unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_NON_SKIPPABLE_IN_STREAM AdGroupTypeEnum_AdGroupType = 11
	// Outstream video ads.
	AdGroupTypeEnum_VIDEO_OUTSTREAM AdGroupTypeEnum_AdGroupType = 12
	// Ad group type for Dynamic Search Ads ad groups.
	AdGroupTypeEnum_SEARCH_DYNAMIC_ADS AdGroupTypeEnum_AdGroupType = 13
)

var AdGroupTypeEnum_AdGroupType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_STANDARD",
	3:  "DISPLAY_STANDARD",
	4:  "SHOPPING_PRODUCT_ADS",
	6:  "HOTEL_ADS",
	7:  "SHOPPING_SMART_ADS",
	8:  "VIDEO_BUMPER",
	9:  "VIDEO_TRUE_VIEW_IN_STREAM",
	10: "VIDEO_TRUE_VIEW_IN_DISPLAY",
	11: "VIDEO_NON_SKIPPABLE_IN_STREAM",
	12: "VIDEO_OUTSTREAM",
	13: "SEARCH_DYNAMIC_ADS",
}
var AdGroupTypeEnum_AdGroupType_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"SEARCH_STANDARD":               2,
	"DISPLAY_STANDARD":              3,
	"SHOPPING_PRODUCT_ADS":          4,
	"HOTEL_ADS":                     6,
	"SHOPPING_SMART_ADS":            7,
	"VIDEO_BUMPER":                  8,
	"VIDEO_TRUE_VIEW_IN_STREAM":     9,
	"VIDEO_TRUE_VIEW_IN_DISPLAY":    10,
	"VIDEO_NON_SKIPPABLE_IN_STREAM": 11,
	"VIDEO_OUTSTREAM":               12,
	"SEARCH_DYNAMIC_ADS":            13,
}

func (x AdGroupTypeEnum_AdGroupType) String() string {
	return proto.EnumName(AdGroupTypeEnum_AdGroupType_name, int32(x))
}
func (AdGroupTypeEnum_AdGroupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_type_7f0b2073c53dd4dc, []int{0, 0}
}

// Defines types of an ad group, specific to a particular campaign channel
// type. This type drives validations that restrict which entities can be
// added to the ad group.
type AdGroupTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupTypeEnum) Reset()         { *m = AdGroupTypeEnum{} }
func (m *AdGroupTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupTypeEnum) ProtoMessage()    {}
func (*AdGroupTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_type_7f0b2073c53dd4dc, []int{0}
}
func (m *AdGroupTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupTypeEnum.Unmarshal(m, b)
}
func (m *AdGroupTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupTypeEnum.Merge(dst, src)
}
func (m *AdGroupTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupTypeEnum.Size(m)
}
func (m *AdGroupTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupTypeEnum)(nil), "google.ads.googleads.v1.enums.AdGroupTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdGroupTypeEnum_AdGroupType", AdGroupTypeEnum_AdGroupType_name, AdGroupTypeEnum_AdGroupType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/ad_group_type.proto", fileDescriptor_ad_group_type_7f0b2073c53dd4dc)
}

var fileDescriptor_ad_group_type_7f0b2073c53dd4dc = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x6a, 0xdb, 0x3c,
	0x18, 0xfd, 0xe3, 0xfe, 0xb4, 0xab, 0xd2, 0x12, 0xa1, 0x95, 0xb1, 0x95, 0x65, 0x90, 0x3e, 0x80,
	0x8d, 0xd9, 0x9d, 0x77, 0x25, 0xc7, 0x5a, 0x22, 0x9a, 0xc8, 0xc2, 0xb2, 0x5d, 0x3a, 0x02, 0xc2,
	0x9b, 0x8d, 0x09, 0x34, 0x96, 0x89, 0x92, 0x42, 0x5f, 0x67, 0x97, 0x7b, 0x82, 0x3d, 0xc3, 0xd8,
	0x6b, 0xec, 0x66, 0x4f, 0x31, 0x2c, 0x25, 0x59, 0x2e, 0xb6, 0xdd, 0x98, 0xcf, 0xe7, 0x9c, 0xef,
	0x70, 0xa4, 0x23, 0xe0, 0xd7, 0x4a, 0xd5, 0x0f, 0x95, 0x57, 0x94, 0xda, 0xb3, 0x63, 0x37, 0x3d,
	0xfa, 0x5e, 0xd5, 0x6c, 0x57, 0xda, 0x2b, 0x4a, 0x59, 0xaf, 0xd5, 0xb6, 0x95, 0x9b, 0xa7, 0xb6,
	0x72, 0xdb, 0xb5, 0xda, 0x28, 0x34, 0xb4, 0x3a, 0xb7, 0x28, 0xb5, 0x7b, 0x58, 0x71, 0x1f, 0x7d,
	0xd7, 0xac, 0x5c, 0xbf, 0xde, 0x3b, 0xb6, 0x4b, 0xaf, 0x68, 0x1a, 0xb5, 0x29, 0x36, 0x4b, 0xd5,
	0x68, 0xbb, 0x7c, 0xf3, 0xdd, 0x01, 0x03, 0x5c, 0x4e, 0x3a, 0xcf, 0xf4, 0xa9, 0xad, 0x48, 0xb3,
	0x5d, 0xdd, 0x7c, 0x75, 0x40, 0xff, 0x08, 0x43, 0x03, 0xd0, 0xcf, 0x98, 0xe0, 0x64, 0x4c, 0xdf,
	0x53, 0x12, 0xc1, 0xff, 0x50, 0x1f, 0x9c, 0x65, 0xec, 0x96, 0xc5, 0x77, 0x0c, 0xf6, 0xd0, 0x73,
	0x30, 0x10, 0x04, 0x27, 0xe3, 0xa9, 0x14, 0x29, 0x66, 0x11, 0x4e, 0x22, 0xe8, 0xa0, 0x2b, 0x00,
	0x23, 0x2a, 0xf8, 0x0c, 0xdf, 0xff, 0x46, 0x4f, 0xd0, 0x4b, 0x70, 0x25, 0xa6, 0x31, 0xe7, 0x94,
	0x4d, 0x24, 0x4f, 0xe2, 0x28, 0x1b, 0xa7, 0x12, 0x47, 0x02, 0xfe, 0x8f, 0x2e, 0xc1, 0xf9, 0x34,
	0x4e, 0xc9, 0xcc, 0xfc, 0x9e, 0xa2, 0x17, 0x00, 0x1d, 0x84, 0x62, 0x8e, 0x13, 0x2b, 0x3b, 0x43,
	0x10, 0x5c, 0xe4, 0x34, 0x22, 0xb1, 0x0c, 0xb3, 0x39, 0x27, 0x09, 0x7c, 0x86, 0x86, 0xe0, 0x95,
	0x45, 0xd2, 0x24, 0x23, 0x32, 0xa7, 0xe4, 0x4e, 0x52, 0x26, 0x45, 0x9a, 0x10, 0x3c, 0x87, 0xe7,
	0xe8, 0x0d, 0xb8, 0xfe, 0x03, 0xbd, 0x8b, 0x06, 0x01, 0x1a, 0x81, 0xa1, 0xe5, 0x59, 0xcc, 0xa4,
	0xb8, 0xa5, 0x9c, 0xe3, 0x70, 0x46, 0x8e, 0x2c, 0xfa, 0xdd, 0xf9, 0xac, 0x24, 0xce, 0xd2, 0x1d,
	0x78, 0x61, 0x02, 0xda, 0x43, 0x47, 0xf7, 0x0c, 0xcf, 0xe9, 0xd8, 0x04, 0xbc, 0x0c, 0x7f, 0xf4,
	0xc0, 0xe8, 0x93, 0x5a, 0xb9, 0xff, 0xac, 0x24, 0x84, 0x47, 0xb7, 0xcb, 0xbb, 0x1a, 0x78, 0xef,
	0x43, 0xb8, 0x5b, 0xa9, 0xd5, 0x43, 0xd1, 0xd4, 0xae, 0x5a, 0xd7, 0x5e, 0x5d, 0x35, 0xa6, 0xa4,
	0xfd, 0x43, 0x68, 0x97, 0xfa, 0x2f, 0xef, 0xe2, 0x9d, 0xf9, 0x7e, 0x76, 0x4e, 0x26, 0x18, 0x7f,
	0x71, 0x86, 0x13, 0x6b, 0x85, 0x4b, 0xed, 0xda, 0xb1, 0x9b, 0x72, 0xdf, 0xed, 0xda, 0xd5, 0xdf,
	0xf6, 0xfc, 0x02, 0x97, 0x7a, 0x71, 0xe0, 0x17, 0xb9, 0xbf, 0x30, 0xfc, 0x4f, 0x67, 0x64, 0xc1,
	0x20, 0xc0, 0xa5, 0x0e, 0x82, 0x83, 0x22, 0x08, 0x72, 0x3f, 0x08, 0x8c, 0xe6, 0xe3, 0xa9, 0x09,
	0xf6, 0xf6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x7b, 0x50, 0x3f, 0xaf, 0x02, 0x00, 0x00,
}
