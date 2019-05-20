// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/detail_placement_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	// The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3" json:"group_placement_target_url,omitempty"`
	// URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v1.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetailPlacementView) Reset()         { *m = DetailPlacementView{} }
func (m *DetailPlacementView) String() string { return proto.CompactTextString(m) }
func (*DetailPlacementView) ProtoMessage()    {}
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9703f91bccf33f03, []int{0}
}

func (m *DetailPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailPlacementView.Unmarshal(m, b)
}
func (m *DetailPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailPlacementView.Marshal(b, m, deterministic)
}
func (m *DetailPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailPlacementView.Merge(m, src)
}
func (m *DetailPlacementView) XXX_Size() int {
	return xxx_messageInfo_DetailPlacementView.Size(m)
}
func (m *DetailPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_DetailPlacementView proto.InternalMessageInfo

func (m *DetailPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DetailPlacementView) GetPlacement() *wrappers.StringValue {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *DetailPlacementView) GetDisplayName() *wrappers.StringValue {
	if m != nil {
		return m.DisplayName
	}
	return nil
}

func (m *DetailPlacementView) GetGroupPlacementTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.GroupPlacementTargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.TargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if m != nil {
		return m.PlacementType
	}
	return enums.PlacementTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*DetailPlacementView)(nil), "google.ads.googleads.v1.resources.DetailPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/detail_placement_view.proto", fileDescriptor_9703f91bccf33f03)
}

var fileDescriptor_9703f91bccf33f03 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x66, 0xa6, 0xba, 0xb0, 0xd9, 0x9f, 0x8b, 0xf1, 0xc2, 0xa1, 0x2c, 0xd2, 0x55, 0x16, 0x7a,
	0x95, 0x61, 0xea, 0x5d, 0x16, 0x95, 0x59, 0x94, 0x05, 0x2f, 0xa4, 0x54, 0x1d, 0x50, 0x0a, 0x43,
	0xb6, 0x73, 0x0c, 0x03, 0x33, 0x49, 0x48, 0x32, 0x2d, 0x7d, 0x00, 0x5f, 0xc4, 0x4b, 0x1f, 0xc0,
	0x87, 0xf0, 0x51, 0x7c, 0x0a, 0x69, 0x66, 0x92, 0x76, 0xc1, 0xba, 0xbd, 0x3b, 0xc9, 0xf9, 0xbe,
	0x2f, 0xdf, 0x97, 0x73, 0xd0, 0x2b, 0x26, 0x04, 0xab, 0x21, 0xa1, 0xa5, 0x4e, 0xba, 0x72, 0x53,
	0x2d, 0xd3, 0x44, 0x81, 0x16, 0xad, 0x5a, 0x80, 0x4e, 0x4a, 0x30, 0xb4, 0xaa, 0x0b, 0x59, 0xd3,
	0x05, 0x34, 0xc0, 0x4d, 0xb1, 0xac, 0x60, 0x85, 0xa5, 0x12, 0x46, 0x44, 0x97, 0x1d, 0x07, 0xd3,
	0x52, 0x63, 0x4f, 0xc7, 0xcb, 0x14, 0x7b, 0xfa, 0x70, 0xb2, 0xef, 0x05, 0xe0, 0x6d, 0xa3, 0x93,
	0xad, 0xac, 0x59, 0x4b, 0xe8, 0x64, 0x87, 0xcf, 0x7a, 0x8e, 0x3d, 0xdd, 0xb5, 0xdf, 0x92, 0x95,
	0xa2, 0x52, 0x82, 0xd2, 0x7d, 0xff, 0xc2, 0x69, 0xca, 0x2a, 0xa1, 0x9c, 0x0b, 0x43, 0x4d, 0x25,
	0x78, 0xdf, 0x7d, 0xfe, 0x6b, 0x80, 0x9e, 0xbc, 0xb5, 0xa6, 0xa7, 0x4e, 0x3c, 0xaf, 0x60, 0x15,
	0xbd, 0x40, 0x67, 0xce, 0x56, 0xc1, 0x69, 0x03, 0x71, 0x30, 0x0a, 0xc6, 0xc7, 0xb3, 0x53, 0x77,
	0xf9, 0x81, 0x36, 0x10, 0x11, 0x74, 0xec, 0x2d, 0xc5, 0xe1, 0x28, 0x18, 0x9f, 0x4c, 0x2e, 0xfa,
	0x68, 0xd8, 0xd9, 0xc1, 0x1f, 0x8d, 0xaa, 0x38, 0xcb, 0x69, 0xdd, 0xc2, 0x6c, 0x0b, 0x8f, 0xde,
	0xa0, 0xd3, 0xb2, 0xd2, 0xb2, 0xa6, 0xeb, 0x4e, 0x7f, 0x70, 0x00, 0xfd, 0xa4, 0x67, 0xd8, 0xc7,
	0xbf, 0xa0, 0x21, 0x53, 0xa2, 0x95, 0x3b, 0x9f, 0x6d, 0xa8, 0x62, 0x60, 0x8a, 0x56, 0xd5, 0xf1,
	0xa3, 0x03, 0xe4, 0x9e, 0x5a, 0xbe, 0xcf, 0xfd, 0xc9, 0xb2, 0x3f, 0xab, 0x3a, 0xba, 0x46, 0x68,
	0x47, 0xea, 0xf1, 0x21, 0xc1, 0x8c, 0x27, 0x03, 0x3a, 0xbf, 0x3f, 0xa7, 0xf8, 0x68, 0x14, 0x8c,
	0xcf, 0x27, 0xaf, 0xf1, 0xbe, 0xf9, 0xdb, 0xe1, 0xe2, 0xad, 0x8f, 0xb5, 0x84, 0x77, 0xbc, 0x6d,
	0xee, 0xdf, 0xcc, 0xce, 0xe4, 0xee, 0xf1, 0xe6, 0x7b, 0x88, 0xae, 0x16, 0xa2, 0xc1, 0x0f, 0x2e,
	0xd5, 0x4d, 0xfc, 0x8f, 0xf9, 0x4e, 0x37, 0x21, 0xa6, 0xc1, 0xd7, 0xf7, 0x3d, 0x9d, 0x89, 0x9a,
	0x72, 0x86, 0x85, 0x62, 0x09, 0x03, 0x6e, 0x23, 0xba, 0x05, 0x94, 0x95, 0xfe, 0xcf, 0xc6, 0x5f,
	0xfb, 0xea, 0x47, 0x38, 0xb8, 0xcd, 0xb2, 0x9f, 0xe1, 0xe5, 0x6d, 0x27, 0x99, 0x95, 0x1a, 0x77,
	0xe5, 0xa6, 0xca, 0x53, 0x3c, 0x73, 0xc8, 0xdf, 0x0e, 0x33, 0xcf, 0x4a, 0x3d, 0xf7, 0x98, 0x79,
	0x9e, 0xce, 0x3d, 0xe6, 0x4f, 0x78, 0xd5, 0x35, 0x08, 0xc9, 0x4a, 0x4d, 0x88, 0x47, 0x11, 0x92,
	0xa7, 0x84, 0x78, 0xdc, 0xdd, 0x91, 0x35, 0xfb, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x02, 0xfe, 0xbc, 0x9d, 0x03, 0x00, 0x00,
}
