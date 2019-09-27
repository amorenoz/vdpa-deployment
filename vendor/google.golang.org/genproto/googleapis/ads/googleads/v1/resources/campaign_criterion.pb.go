// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_criterion.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A campaign criterion.
type CampaignCriterion struct {
	// The resource name of the campaign criterion.
	// Campaign criterion resource names have the form:
	//
	// `customers/{customer_id}/campaignCriteria/{campaign_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored during mutate.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,5,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bids when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. Most targetable criteria types support modifiers.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.FloatValue `protobuf:"bytes,14,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether to target (`false`) or exclude (`true`) the criterion.
	Negative *wrappers.BoolValue `protobuf:"bytes,7,opt,name=negative,proto3" json:"negative,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The campaign criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignCriterion_Keyword
	//	*CampaignCriterion_Placement
	//	*CampaignCriterion_MobileAppCategory
	//	*CampaignCriterion_MobileApplication
	//	*CampaignCriterion_Location
	//	*CampaignCriterion_Device
	//	*CampaignCriterion_AdSchedule
	//	*CampaignCriterion_AgeRange
	//	*CampaignCriterion_Gender
	//	*CampaignCriterion_IncomeRange
	//	*CampaignCriterion_ParentalStatus
	//	*CampaignCriterion_UserList
	//	*CampaignCriterion_YoutubeVideo
	//	*CampaignCriterion_YoutubeChannel
	//	*CampaignCriterion_Proximity
	//	*CampaignCriterion_Topic
	//	*CampaignCriterion_ListingScope
	//	*CampaignCriterion_Language
	//	*CampaignCriterion_IpBlock
	//	*CampaignCriterion_ContentLabel
	//	*CampaignCriterion_Carrier
	//	*CampaignCriterion_UserInterest
	//	*CampaignCriterion_Webpage
	//	*CampaignCriterion_OperatingSystemVersion
	//	*CampaignCriterion_MobileDevice
	Criterion            isCampaignCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignCriterion) Reset()         { *m = CampaignCriterion{} }
func (m *CampaignCriterion) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterion) ProtoMessage()    {}
func (*CampaignCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_criterion_d0371d692ca57c85, []int{0}
}
func (m *CampaignCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterion.Unmarshal(m, b)
}
func (m *CampaignCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterion.Marshal(b, m, deterministic)
}
func (dst *CampaignCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterion.Merge(dst, src)
}
func (m *CampaignCriterion) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterion.Size(m)
}
func (m *CampaignCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterion proto.InternalMessageInfo

func (m *CampaignCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterion) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterion) GetBidModifier() *wrappers.FloatValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *CampaignCriterion) GetNegative() *wrappers.BoolValue {
	if m != nil {
		return m.Negative
	}
	return nil
}

func (m *CampaignCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isCampaignCriterion_Criterion interface {
	isCampaignCriterion_Criterion()
}

type CampaignCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,8,opt,name=keyword,proto3,oneof"`
}

type CampaignCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,9,opt,name=placement,proto3,oneof"`
}

type CampaignCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,10,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CampaignCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,11,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CampaignCriterion_Location struct {
	Location *common.LocationInfo `protobuf:"bytes,12,opt,name=location,proto3,oneof"`
}

type CampaignCriterion_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,13,opt,name=device,proto3,oneof"`
}

type CampaignCriterion_AdSchedule struct {
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,15,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

type CampaignCriterion_AgeRange struct {
	AgeRange *common.AgeRangeInfo `protobuf:"bytes,16,opt,name=age_range,json=ageRange,proto3,oneof"`
}

type CampaignCriterion_Gender struct {
	Gender *common.GenderInfo `protobuf:"bytes,17,opt,name=gender,proto3,oneof"`
}

type CampaignCriterion_IncomeRange struct {
	IncomeRange *common.IncomeRangeInfo `protobuf:"bytes,18,opt,name=income_range,json=incomeRange,proto3,oneof"`
}

type CampaignCriterion_ParentalStatus struct {
	ParentalStatus *common.ParentalStatusInfo `protobuf:"bytes,19,opt,name=parental_status,json=parentalStatus,proto3,oneof"`
}

type CampaignCriterion_UserList struct {
	UserList *common.UserListInfo `protobuf:"bytes,22,opt,name=user_list,json=userList,proto3,oneof"`
}

type CampaignCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,20,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CampaignCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,21,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type CampaignCriterion_Proximity struct {
	Proximity *common.ProximityInfo `protobuf:"bytes,23,opt,name=proximity,proto3,oneof"`
}

type CampaignCriterion_Topic struct {
	Topic *common.TopicInfo `protobuf:"bytes,24,opt,name=topic,proto3,oneof"`
}

type CampaignCriterion_ListingScope struct {
	ListingScope *common.ListingScopeInfo `protobuf:"bytes,25,opt,name=listing_scope,json=listingScope,proto3,oneof"`
}

type CampaignCriterion_Language struct {
	Language *common.LanguageInfo `protobuf:"bytes,26,opt,name=language,proto3,oneof"`
}

type CampaignCriterion_IpBlock struct {
	IpBlock *common.IpBlockInfo `protobuf:"bytes,27,opt,name=ip_block,json=ipBlock,proto3,oneof"`
}

type CampaignCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,28,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CampaignCriterion_Carrier struct {
	Carrier *common.CarrierInfo `protobuf:"bytes,29,opt,name=carrier,proto3,oneof"`
}

type CampaignCriterion_UserInterest struct {
	UserInterest *common.UserInterestInfo `protobuf:"bytes,30,opt,name=user_interest,json=userInterest,proto3,oneof"`
}

type CampaignCriterion_Webpage struct {
	Webpage *common.WebpageInfo `protobuf:"bytes,31,opt,name=webpage,proto3,oneof"`
}

type CampaignCriterion_OperatingSystemVersion struct {
	OperatingSystemVersion *common.OperatingSystemVersionInfo `protobuf:"bytes,32,opt,name=operating_system_version,json=operatingSystemVersion,proto3,oneof"`
}

type CampaignCriterion_MobileDevice struct {
	MobileDevice *common.MobileDeviceInfo `protobuf:"bytes,33,opt,name=mobile_device,json=mobileDevice,proto3,oneof"`
}

func (*CampaignCriterion_Keyword) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Placement) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileAppCategory) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileApplication) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Location) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Device) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AdSchedule) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AgeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Gender) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IncomeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ParentalStatus) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserList) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeVideo) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeChannel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Proximity) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Topic) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ListingScope) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Language) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IpBlock) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ContentLabel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Carrier) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserInterest) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Webpage) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_OperatingSystemVersion) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileDevice) isCampaignCriterion_Criterion() {}

func (m *CampaignCriterion) GetCriterion() isCampaignCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *CampaignCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CampaignCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CampaignCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (m *CampaignCriterion) GetLocation() *common.LocationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Location); ok {
		return x.Location
	}
	return nil
}

func (m *CampaignCriterion) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Device); ok {
		return x.Device
	}
	return nil
}

func (m *CampaignCriterion) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

func (m *CampaignCriterion) GetAgeRange() *common.AgeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AgeRange); ok {
		return x.AgeRange
	}
	return nil
}

func (m *CampaignCriterion) GetGender() *common.GenderInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Gender); ok {
		return x.Gender
	}
	return nil
}

func (m *CampaignCriterion) GetIncomeRange() *common.IncomeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IncomeRange); ok {
		return x.IncomeRange
	}
	return nil
}

func (m *CampaignCriterion) GetParentalStatus() *common.ParentalStatusInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ParentalStatus); ok {
		return x.ParentalStatus
	}
	return nil
}

func (m *CampaignCriterion) GetUserList() *common.UserListInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserList); ok {
		return x.UserList
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *CampaignCriterion) GetProximity() *common.ProximityInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Proximity); ok {
		return x.Proximity
	}
	return nil
}

func (m *CampaignCriterion) GetTopic() *common.TopicInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Topic); ok {
		return x.Topic
	}
	return nil
}

func (m *CampaignCriterion) GetListingScope() *common.ListingScopeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ListingScope); ok {
		return x.ListingScope
	}
	return nil
}

func (m *CampaignCriterion) GetLanguage() *common.LanguageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Language); ok {
		return x.Language
	}
	return nil
}

func (m *CampaignCriterion) GetIpBlock() *common.IpBlockInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IpBlock); ok {
		return x.IpBlock
	}
	return nil
}

func (m *CampaignCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CampaignCriterion) GetCarrier() *common.CarrierInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Carrier); ok {
		return x.Carrier
	}
	return nil
}

func (m *CampaignCriterion) GetUserInterest() *common.UserInterestInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserInterest); ok {
		return x.UserInterest
	}
	return nil
}

func (m *CampaignCriterion) GetWebpage() *common.WebpageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Webpage); ok {
		return x.Webpage
	}
	return nil
}

func (m *CampaignCriterion) GetOperatingSystemVersion() *common.OperatingSystemVersionInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_OperatingSystemVersion); ok {
		return x.OperatingSystemVersion
	}
	return nil
}

func (m *CampaignCriterion) GetMobileDevice() *common.MobileDeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileDevice); ok {
		return x.MobileDevice
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CampaignCriterion) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CampaignCriterion_OneofMarshaler, _CampaignCriterion_OneofUnmarshaler, _CampaignCriterion_OneofSizer, []interface{}{
		(*CampaignCriterion_Keyword)(nil),
		(*CampaignCriterion_Placement)(nil),
		(*CampaignCriterion_MobileAppCategory)(nil),
		(*CampaignCriterion_MobileApplication)(nil),
		(*CampaignCriterion_Location)(nil),
		(*CampaignCriterion_Device)(nil),
		(*CampaignCriterion_AdSchedule)(nil),
		(*CampaignCriterion_AgeRange)(nil),
		(*CampaignCriterion_Gender)(nil),
		(*CampaignCriterion_IncomeRange)(nil),
		(*CampaignCriterion_ParentalStatus)(nil),
		(*CampaignCriterion_UserList)(nil),
		(*CampaignCriterion_YoutubeVideo)(nil),
		(*CampaignCriterion_YoutubeChannel)(nil),
		(*CampaignCriterion_Proximity)(nil),
		(*CampaignCriterion_Topic)(nil),
		(*CampaignCriterion_ListingScope)(nil),
		(*CampaignCriterion_Language)(nil),
		(*CampaignCriterion_IpBlock)(nil),
		(*CampaignCriterion_ContentLabel)(nil),
		(*CampaignCriterion_Carrier)(nil),
		(*CampaignCriterion_UserInterest)(nil),
		(*CampaignCriterion_Webpage)(nil),
		(*CampaignCriterion_OperatingSystemVersion)(nil),
		(*CampaignCriterion_MobileDevice)(nil),
	}
}

func _CampaignCriterion_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CampaignCriterion)
	// criterion
	switch x := m.Criterion.(type) {
	case *CampaignCriterion_Keyword:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Keyword); err != nil {
			return err
		}
	case *CampaignCriterion_Placement:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Placement); err != nil {
			return err
		}
	case *CampaignCriterion_MobileAppCategory:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MobileAppCategory); err != nil {
			return err
		}
	case *CampaignCriterion_MobileApplication:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MobileApplication); err != nil {
			return err
		}
	case *CampaignCriterion_Location:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Location); err != nil {
			return err
		}
	case *CampaignCriterion_Device:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Device); err != nil {
			return err
		}
	case *CampaignCriterion_AdSchedule:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AdSchedule); err != nil {
			return err
		}
	case *CampaignCriterion_AgeRange:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AgeRange); err != nil {
			return err
		}
	case *CampaignCriterion_Gender:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gender); err != nil {
			return err
		}
	case *CampaignCriterion_IncomeRange:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IncomeRange); err != nil {
			return err
		}
	case *CampaignCriterion_ParentalStatus:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ParentalStatus); err != nil {
			return err
		}
	case *CampaignCriterion_UserList:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserList); err != nil {
			return err
		}
	case *CampaignCriterion_YoutubeVideo:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.YoutubeVideo); err != nil {
			return err
		}
	case *CampaignCriterion_YoutubeChannel:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.YoutubeChannel); err != nil {
			return err
		}
	case *CampaignCriterion_Proximity:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Proximity); err != nil {
			return err
		}
	case *CampaignCriterion_Topic:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Topic); err != nil {
			return err
		}
	case *CampaignCriterion_ListingScope:
		b.EncodeVarint(25<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListingScope); err != nil {
			return err
		}
	case *CampaignCriterion_Language:
		b.EncodeVarint(26<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Language); err != nil {
			return err
		}
	case *CampaignCriterion_IpBlock:
		b.EncodeVarint(27<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IpBlock); err != nil {
			return err
		}
	case *CampaignCriterion_ContentLabel:
		b.EncodeVarint(28<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContentLabel); err != nil {
			return err
		}
	case *CampaignCriterion_Carrier:
		b.EncodeVarint(29<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Carrier); err != nil {
			return err
		}
	case *CampaignCriterion_UserInterest:
		b.EncodeVarint(30<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserInterest); err != nil {
			return err
		}
	case *CampaignCriterion_Webpage:
		b.EncodeVarint(31<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Webpage); err != nil {
			return err
		}
	case *CampaignCriterion_OperatingSystemVersion:
		b.EncodeVarint(32<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OperatingSystemVersion); err != nil {
			return err
		}
	case *CampaignCriterion_MobileDevice:
		b.EncodeVarint(33<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MobileDevice); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CampaignCriterion.Criterion has unexpected type %T", x)
	}
	return nil
}

func _CampaignCriterion_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CampaignCriterion)
	switch tag {
	case 8: // criterion.keyword
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.KeywordInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Keyword{msg}
		return true, err
	case 9: // criterion.placement
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.PlacementInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Placement{msg}
		return true, err
	case 10: // criterion.mobile_app_category
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.MobileAppCategoryInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_MobileAppCategory{msg}
		return true, err
	case 11: // criterion.mobile_application
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.MobileApplicationInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_MobileApplication{msg}
		return true, err
	case 12: // criterion.location
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.LocationInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Location{msg}
		return true, err
	case 13: // criterion.device
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.DeviceInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Device{msg}
		return true, err
	case 15: // criterion.ad_schedule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.AdScheduleInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_AdSchedule{msg}
		return true, err
	case 16: // criterion.age_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.AgeRangeInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_AgeRange{msg}
		return true, err
	case 17: // criterion.gender
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.GenderInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Gender{msg}
		return true, err
	case 18: // criterion.income_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.IncomeRangeInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_IncomeRange{msg}
		return true, err
	case 19: // criterion.parental_status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ParentalStatusInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_ParentalStatus{msg}
		return true, err
	case 22: // criterion.user_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.UserListInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_UserList{msg}
		return true, err
	case 20: // criterion.youtube_video
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.YouTubeVideoInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_YoutubeVideo{msg}
		return true, err
	case 21: // criterion.youtube_channel
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.YouTubeChannelInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_YoutubeChannel{msg}
		return true, err
	case 23: // criterion.proximity
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ProximityInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Proximity{msg}
		return true, err
	case 24: // criterion.topic
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TopicInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Topic{msg}
		return true, err
	case 25: // criterion.listing_scope
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ListingScopeInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_ListingScope{msg}
		return true, err
	case 26: // criterion.language
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.LanguageInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Language{msg}
		return true, err
	case 27: // criterion.ip_block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.IpBlockInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_IpBlock{msg}
		return true, err
	case 28: // criterion.content_label
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ContentLabelInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_ContentLabel{msg}
		return true, err
	case 29: // criterion.carrier
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CarrierInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Carrier{msg}
		return true, err
	case 30: // criterion.user_interest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.UserInterestInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_UserInterest{msg}
		return true, err
	case 31: // criterion.webpage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.WebpageInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Webpage{msg}
		return true, err
	case 32: // criterion.operating_system_version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.OperatingSystemVersionInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_OperatingSystemVersion{msg}
		return true, err
	case 33: // criterion.mobile_device
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.MobileDeviceInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_MobileDevice{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CampaignCriterion_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CampaignCriterion)
	// criterion
	switch x := m.Criterion.(type) {
	case *CampaignCriterion_Keyword:
		s := proto.Size(x.Keyword)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Placement:
		s := proto.Size(x.Placement)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_MobileAppCategory:
		s := proto.Size(x.MobileAppCategory)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_MobileApplication:
		s := proto.Size(x.MobileApplication)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Location:
		s := proto.Size(x.Location)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Device:
		s := proto.Size(x.Device)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_AdSchedule:
		s := proto.Size(x.AdSchedule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_AgeRange:
		s := proto.Size(x.AgeRange)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Gender:
		s := proto.Size(x.Gender)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_IncomeRange:
		s := proto.Size(x.IncomeRange)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_ParentalStatus:
		s := proto.Size(x.ParentalStatus)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_UserList:
		s := proto.Size(x.UserList)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_YoutubeVideo:
		s := proto.Size(x.YoutubeVideo)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_YoutubeChannel:
		s := proto.Size(x.YoutubeChannel)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Proximity:
		s := proto.Size(x.Proximity)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Topic:
		s := proto.Size(x.Topic)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_ListingScope:
		s := proto.Size(x.ListingScope)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Language:
		s := proto.Size(x.Language)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_IpBlock:
		s := proto.Size(x.IpBlock)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_ContentLabel:
		s := proto.Size(x.ContentLabel)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Carrier:
		s := proto.Size(x.Carrier)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_UserInterest:
		s := proto.Size(x.UserInterest)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Webpage:
		s := proto.Size(x.Webpage)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_OperatingSystemVersion:
		s := proto.Size(x.OperatingSystemVersion)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_MobileDevice:
		s := proto.Size(x.MobileDevice)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CampaignCriterion)(nil), "google.ads.googleads.v1.resources.CampaignCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_criterion.proto", fileDescriptor_campaign_criterion_d0371d692ca57c85)
}

var fileDescriptor_campaign_criterion_d0371d692ca57c85 = []byte{
	// 1058 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x6f, 0x6f, 0xdc, 0x34,
	0x18, 0xc0, 0x69, 0xd9, 0xba, 0xab, 0xef, 0xda, 0x51, 0x17, 0x8a, 0xe9, 0xca, 0x68, 0x41, 0x93,
	0xca, 0x9f, 0xe5, 0x68, 0x81, 0x09, 0x1d, 0xd2, 0xa4, 0xeb, 0x0d, 0xba, 0xdb, 0x5a, 0x28, 0xd7,
	0x72, 0x13, 0xa8, 0x28, 0xf2, 0x25, 0x4f, 0x33, 0x6b, 0x89, 0x1d, 0xd9, 0xce, 0x95, 0x7b, 0xcd,
	0x37, 0xe1, 0x25, 0x9f, 0x81, 0x4f, 0xc0, 0x47, 0xe1, 0x53, 0x20, 0x3b, 0x76, 0xda, 0x52, 0x4a,
	0xc2, 0xbb, 0xe4, 0xf1, 0xf3, 0xfb, 0xe5, 0xf1, 0x93, 0xc4, 0x36, 0xea, 0x25, 0x42, 0x24, 0x29,
	0x74, 0x69, 0xac, 0xba, 0xe5, 0xa5, 0xb9, 0x9a, 0xee, 0x74, 0x25, 0x28, 0x51, 0xc8, 0x08, 0x54,
	0x37, 0xa2, 0x59, 0x4e, 0x59, 0xc2, 0xc3, 0x48, 0x32, 0x0d, 0x92, 0x09, 0x1e, 0xe4, 0x52, 0x68,
	0x81, 0xb7, 0x4a, 0x20, 0xa0, 0xb1, 0x0a, 0x2a, 0x36, 0x98, 0xee, 0x04, 0x15, 0xbb, 0xfe, 0xf0,
	0x26, 0x7d, 0x24, 0xb2, 0x4c, 0xf0, 0xae, 0x53, 0xd2, 0xd2, 0xb8, 0xbe, 0x7b, 0x53, 0x3a, 0xf0,
	0x22, 0x53, 0xdd, 0xaa, 0x80, 0x50, 0xcf, 0x72, 0x70, 0xcc, 0x7d, 0xc7, 0xd8, 0xbb, 0x49, 0x71,
	0xd6, 0x3d, 0x97, 0x34, 0xcf, 0x41, 0x2a, 0x37, 0xbe, 0xe1, 0x9d, 0x39, 0xeb, 0x52, 0xce, 0x85,
	0xa6, 0x9a, 0x09, 0xee, 0x46, 0xdf, 0xff, 0x63, 0x15, 0xad, 0x0c, 0xdc, 0x04, 0x07, 0x5e, 0x8f,
	0x3f, 0x40, 0x4b, 0x7e, 0x0e, 0x21, 0xa7, 0x19, 0x90, 0xb9, 0xcd, 0xb9, 0xed, 0xc5, 0x51, 0xc7,
	0x07, 0xbf, 0xa5, 0x19, 0xe0, 0x2f, 0x51, 0xcb, 0xb7, 0x86, 0xdc, 0xda, 0x9c, 0xdb, 0x6e, 0xef,
	0x6e, 0xb8, 0x36, 0x04, 0xbe, 0x96, 0xe0, 0x58, 0x4b, 0xc6, 0x93, 0x31, 0x4d, 0x0b, 0x18, 0x55,
	0xd9, 0xf8, 0x31, 0xea, 0x5c, 0x4c, 0x85, 0xc5, 0xe4, 0xb6, 0xa5, 0xef, 0x5d, 0xa3, 0x87, 0x5c,
	0x3f, 0xfa, 0xbc, 0x84, 0xdb, 0x15, 0x30, 0x8c, 0x0d, 0x3f, 0x61, 0x71, 0x98, 0x89, 0x98, 0x9d,
	0x31, 0x90, 0x64, 0xf9, 0x06, 0xfe, 0x9b, 0x54, 0x50, 0xed, 0xf8, 0x09, 0x8b, 0x0f, 0x5d, 0x3e,
	0x7e, 0x84, 0x5a, 0x1c, 0x12, 0xaa, 0xd9, 0x14, 0xc8, 0x1d, 0xcb, 0xae, 0x5f, 0x63, 0xf7, 0x84,
	0x48, 0x5d, 0xdd, 0x3e, 0x17, 0x8f, 0xd0, 0x2d, 0xd3, 0x78, 0xb2, 0xb0, 0x39, 0xb7, 0xbd, 0xbc,
	0xfb, 0x38, 0xb8, 0xe9, 0xfd, 0xdb, 0xb7, 0x15, 0x54, 0xed, 0x3c, 0x99, 0xe5, 0xf0, 0x35, 0x2f,
	0xb2, 0xab, 0x91, 0x91, 0x75, 0xe1, 0x7d, 0x74, 0xe7, 0x15, 0xcc, 0xce, 0x85, 0x8c, 0x49, 0xcb,
	0x96, 0xf2, 0xf1, 0x8d, 0xda, 0xf2, 0x9b, 0x09, 0x9e, 0x97, 0xe9, 0x43, 0x7e, 0x26, 0x9e, 0xbe,
	0x36, 0xf2, 0x34, 0x3e, 0x44, 0x8b, 0x79, 0x4a, 0x23, 0xc8, 0x80, 0x6b, 0xb2, 0x68, 0x55, 0x0f,
	0xeb, 0x54, 0x47, 0x1e, 0x70, 0xb2, 0x0b, 0x03, 0x4e, 0xd0, 0x6a, 0x26, 0x26, 0x2c, 0x85, 0x90,
	0xe6, 0x79, 0x18, 0x51, 0x0d, 0x89, 0x90, 0x33, 0x82, 0xac, 0xf8, 0x8b, 0x3a, 0xf1, 0xa1, 0x45,
	0xfb, 0x79, 0x3e, 0x70, 0xa0, 0x7b, 0xc0, 0x4a, 0xf6, 0xcf, 0x01, 0x7c, 0x86, 0xf0, 0xc5, 0x83,
	0x52, 0x16, 0xd9, 0xcf, 0x93, 0xb4, 0xff, 0xe7, 0x73, 0x3c, 0x78, 0xed, 0x39, 0x7e, 0x00, 0x3f,
	0x43, 0xad, 0x54, 0x38, 0x7b, 0xc7, 0xda, 0x3f, 0xa9, 0xb3, 0x1f, 0x88, 0x2b, 0xd2, 0x8a, 0xc7,
	0x4f, 0xd0, 0x42, 0x0c, 0x53, 0x16, 0x01, 0x59, 0xb2, 0xa6, 0x8f, 0xea, 0x4c, 0x4f, 0x6c, 0xb6,
	0xf3, 0x38, 0x16, 0x7f, 0x8f, 0xda, 0x34, 0x0e, 0x55, 0xf4, 0x12, 0xe2, 0x22, 0x05, 0x72, 0xd7,
	0xaa, 0x82, 0x3a, 0x55, 0x3f, 0x3e, 0x76, 0x84, 0xd3, 0x21, 0x5a, 0x45, 0xf0, 0x73, 0xb4, 0x48,
	0x13, 0x08, 0x25, 0xe5, 0x09, 0x90, 0x37, 0x9a, 0xcd, 0xb2, 0x9f, 0xc0, 0xc8, 0xe4, 0xfb, 0x59,
	0x52, 0x77, 0x6f, 0x66, 0x99, 0x00, 0x8f, 0x41, 0x92, 0x95, 0x66, 0xb3, 0xdc, 0xb7, 0xd9, 0x7e,
	0x96, 0x25, 0x8b, 0x4f, 0x50, 0x87, 0xf1, 0x48, 0x64, 0xbe, 0x2a, 0x6c, 0x5d, 0xdd, 0x3a, 0xd7,
	0xd0, 0x32, 0x97, 0x0b, 0x6b, 0xb3, 0x8b, 0x10, 0xfe, 0x19, 0xdd, 0xcd, 0xa9, 0x04, 0xae, 0x69,
	0x1a, 0x2a, 0x4d, 0x75, 0xa1, 0xc8, 0xaa, 0x15, 0xef, 0xd6, 0x7e, 0xf3, 0x0e, 0x3b, 0xb6, 0x94,
	0x73, 0x2f, 0xe7, 0x57, 0xa2, 0xa6, 0x8f, 0x85, 0x02, 0x19, 0xa6, 0x4c, 0x69, 0xb2, 0xd6, 0xac,
	0x8f, 0x3f, 0x28, 0x90, 0x07, 0x4c, 0xf9, 0x7f, 0xa9, 0x55, 0xb8, 0x7b, 0xfc, 0x02, 0x2d, 0xcd,
	0x44, 0xa1, 0x8b, 0x09, 0x84, 0x53, 0x16, 0x83, 0x20, 0x6f, 0x5a, 0xe1, 0xa7, 0x75, 0xc2, 0x1f,
	0x45, 0x71, 0x52, 0x4c, 0x60, 0x6c, 0x18, 0x27, 0xed, 0x38, 0x91, 0x8d, 0x99, 0x26, 0x78, 0x71,
	0xf4, 0x92, 0x72, 0x0e, 0x29, 0x79, 0xab, 0x59, 0x13, 0x9c, 0x7a, 0x50, 0x52, 0xbe, 0x09, 0x4e,
	0xe6, 0xa2, 0x76, 0x45, 0x91, 0xe2, 0x17, 0x96, 0x31, 0x3d, 0x23, 0x6f, 0x37, 0x5c, 0x51, 0x3c,
	0x50, 0xad, 0x28, 0x3e, 0x80, 0xfb, 0xe8, 0xb6, 0x16, 0x39, 0x8b, 0x08, 0xb1, 0xaa, 0x0f, 0xeb,
	0x54, 0x27, 0x26, 0xd9, 0x69, 0x4a, 0xd2, 0x74, 0xd2, 0xbc, 0x11, 0xc6, 0x93, 0x50, 0x45, 0x22,
	0x07, 0xf2, 0x4e, 0xb3, 0x4e, 0x1e, 0x94, 0xd0, 0xb1, 0x61, 0x7c, 0x27, 0xd3, 0x4b, 0x31, 0xbb,
	0x38, 0x50, 0x9e, 0x14, 0x34, 0x01, 0xb2, 0xde, 0x70, 0x71, 0x70, 0xf9, 0xd5, 0xe2, 0xe0, 0xee,
	0xf1, 0x53, 0xd4, 0x62, 0x79, 0x38, 0x49, 0x45, 0xf4, 0x8a, 0xdc, 0x6b, 0xb6, 0xa4, 0x0f, 0xf3,
	0x3d, 0x93, 0xee, 0x97, 0x74, 0x56, 0xde, 0x9a, 0xe9, 0x46, 0x82, 0x6b, 0xe0, 0x3a, 0x4c, 0xe9,
	0x04, 0x52, 0xb2, 0xd1, 0x6c, 0xba, 0x83, 0x12, 0x3a, 0x30, 0x8c, 0x9f, 0x6e, 0x74, 0x29, 0x66,
	0x36, 0x9d, 0x88, 0x4a, 0x69, 0xf6, 0xce, 0x77, 0x9b, 0x55, 0x38, 0x28, 0xd3, 0x7d, 0x85, 0x8e,
	0x36, 0x15, 0xda, 0xff, 0x84, 0x71, 0x0d, 0x12, 0x94, 0x26, 0xf7, 0x9b, 0x55, 0x68, 0xfe, 0x95,
	0xa1, 0x63, 0x7c, 0x85, 0xc5, 0xa5, 0x98, 0xa9, 0xf0, 0x1c, 0x26, 0xb9, 0x79, 0x1f, 0xef, 0x35,
	0xab, 0xf0, 0x45, 0x99, 0xee, 0x2b, 0x74, 0x34, 0x9e, 0x22, 0x22, 0x72, 0x90, 0xb4, 0xfc, 0x68,
	0x66, 0x4a, 0x43, 0x16, 0x4e, 0x41, 0x2a, 0xb3, 0x0d, 0x6c, 0x5a, 0x73, 0xaf, 0xce, 0xfc, 0x9d,
	0xe7, 0x8f, 0x2d, 0x3e, 0x2e, 0x69, 0xf7, 0xa0, 0x35, 0xf1, 0xaf, 0xa3, 0xa6, 0x33, 0x6e, 0x5b,
	0x73, 0x3b, 0xc5, 0x56, 0xb3, 0xce, 0x94, 0x3b, 0xda, 0x95, 0xfd, 0xa2, 0x93, 0x5d, 0x8a, 0xed,
	0xb5, 0xd1, 0x62, 0x75, 0x16, 0xda, 0xfb, 0x75, 0x1e, 0x3d, 0x88, 0x44, 0x16, 0xd4, 0x9e, 0x44,
	0xf7, 0xd6, 0xae, 0x9d, 0xf2, 0x8e, 0xcc, 0x51, 0xe7, 0x68, 0xee, 0xa7, 0x67, 0x0e, 0x4e, 0x84,
	0xf9, 0x84, 0x03, 0x21, 0x93, 0x6e, 0x02, 0xdc, 0x1e, 0x84, 0xfc, 0x21, 0x34, 0x67, 0xea, 0x3f,
	0x4e, 0xc8, 0x5f, 0x55, 0x57, 0xbf, 0xcd, 0xbf, 0xbe, 0xdf, 0xef, 0xff, 0x3e, 0xbf, 0xb5, 0x5f,
	0x2a, 0xfb, 0xb1, 0x0a, 0xca, 0x4b, 0x73, 0x35, 0xde, 0x09, 0x46, 0x3e, 0xf3, 0x4f, 0x9f, 0x73,
	0xda, 0x8f, 0xd5, 0x69, 0x95, 0x73, 0x3a, 0xde, 0x39, 0xad, 0x72, 0xfe, 0x9a, 0x7f, 0x50, 0x0e,
	0xf4, 0x7a, 0xfd, 0x58, 0xf5, 0x7a, 0x55, 0x56, 0xaf, 0x37, 0xde, 0xe9, 0xf5, 0xaa, 0xbc, 0xc9,
	0x82, 0x2d, 0xf6, 0xb3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xa3, 0xc9, 0x99, 0xcd, 0x0b,
	0x00, 0x00,
}
