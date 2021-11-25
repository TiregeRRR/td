// Code generated by gotdgen, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// Config represents TL type `config#330b4067`.
type Config struct {
	// Flags field of Config.
	Flags bin.Fields
	// PhonecallsEnabled field of Config.
	PhonecallsEnabled bool
	// DefaultP2PContacts field of Config.
	DefaultP2PContacts bool
	// PreloadFeaturedStickers field of Config.
	PreloadFeaturedStickers bool
	// IgnorePhoneEntities field of Config.
	IgnorePhoneEntities bool
	// RevokePmInbox field of Config.
	RevokePmInbox bool
	// BlockedMode field of Config.
	BlockedMode bool
	// PFSEnabled field of Config.
	PFSEnabled bool
	// Date field of Config.
	Date int
	// Expires field of Config.
	Expires int
	// TestMode field of Config.
	TestMode bool
	// ThisDC field of Config.
	ThisDC int
	// DCOptions field of Config.
	DCOptions []DCOption
	// DCTxtDomainName field of Config.
	DCTxtDomainName string
	// ChatSizeMax field of Config.
	ChatSizeMax int
	// MegagroupSizeMax field of Config.
	MegagroupSizeMax int
	// ForwardedCountMax field of Config.
	ForwardedCountMax int
	// OnlineUpdatePeriodMs field of Config.
	OnlineUpdatePeriodMs int
	// OfflineBlurTimeoutMs field of Config.
	OfflineBlurTimeoutMs int
	// OfflineIdleTimeoutMs field of Config.
	OfflineIdleTimeoutMs int
	// OnlineCloudTimeoutMs field of Config.
	OnlineCloudTimeoutMs int
	// NotifyCloudDelayMs field of Config.
	NotifyCloudDelayMs int
	// NotifyDefaultDelayMs field of Config.
	NotifyDefaultDelayMs int
	// PushChatPeriodMs field of Config.
	PushChatPeriodMs int
	// PushChatLimit field of Config.
	PushChatLimit int
	// SavedGifsLimit field of Config.
	SavedGifsLimit int
	// EditTimeLimit field of Config.
	EditTimeLimit int
	// RevokeTimeLimit field of Config.
	RevokeTimeLimit int
	// RevokePmTimeLimit field of Config.
	RevokePmTimeLimit int
	// RatingEDecay field of Config.
	RatingEDecay int
	// StickersRecentLimit field of Config.
	StickersRecentLimit int
	// StickersFavedLimit field of Config.
	StickersFavedLimit int
	// ChannelsReadMediaPeriod field of Config.
	ChannelsReadMediaPeriod int
	// TmpSessions field of Config.
	//
	// Use SetTmpSessions and GetTmpSessions helpers.
	TmpSessions int
	// PinnedDialogsCountMax field of Config.
	PinnedDialogsCountMax int
	// PinnedInfolderCountMax field of Config.
	PinnedInfolderCountMax int
	// CallReceiveTimeoutMs field of Config.
	CallReceiveTimeoutMs int
	// CallRingTimeoutMs field of Config.
	CallRingTimeoutMs int
	// CallConnectTimeoutMs field of Config.
	CallConnectTimeoutMs int
	// CallPacketTimeoutMs field of Config.
	CallPacketTimeoutMs int
	// MeURLPrefix field of Config.
	MeURLPrefix string
	// AutoupdateURLPrefix field of Config.
	//
	// Use SetAutoupdateURLPrefix and GetAutoupdateURLPrefix helpers.
	AutoupdateURLPrefix string
	// GifSearchUsername field of Config.
	//
	// Use SetGifSearchUsername and GetGifSearchUsername helpers.
	GifSearchUsername string
	// VenueSearchUsername field of Config.
	//
	// Use SetVenueSearchUsername and GetVenueSearchUsername helpers.
	VenueSearchUsername string
	// ImgSearchUsername field of Config.
	//
	// Use SetImgSearchUsername and GetImgSearchUsername helpers.
	ImgSearchUsername string
	// StaticMapsProvider field of Config.
	//
	// Use SetStaticMapsProvider and GetStaticMapsProvider helpers.
	StaticMapsProvider string
	// CaptionLengthMax field of Config.
	CaptionLengthMax int
	// MessageLengthMax field of Config.
	MessageLengthMax int
	// WebfileDCID field of Config.
	WebfileDCID int
	// SuggestedLangCode field of Config.
	//
	// Use SetSuggestedLangCode and GetSuggestedLangCode helpers.
	SuggestedLangCode string
	// LangPackVersion field of Config.
	//
	// Use SetLangPackVersion and GetLangPackVersion helpers.
	LangPackVersion int
	// BaseLangPackVersion field of Config.
	//
	// Use SetBaseLangPackVersion and GetBaseLangPackVersion helpers.
	BaseLangPackVersion int
}

// ConfigTypeID is TL type id of Config.
const ConfigTypeID = 0x330b4067

// Ensuring interfaces in compile-time for Config.
var (
	_ bin.Encoder     = &Config{}
	_ bin.Decoder     = &Config{}
	_ bin.BareEncoder = &Config{}
	_ bin.BareDecoder = &Config{}
)

func (c *Config) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.PhonecallsEnabled == false) {
		return false
	}
	if !(c.DefaultP2PContacts == false) {
		return false
	}
	if !(c.PreloadFeaturedStickers == false) {
		return false
	}
	if !(c.IgnorePhoneEntities == false) {
		return false
	}
	if !(c.RevokePmInbox == false) {
		return false
	}
	if !(c.BlockedMode == false) {
		return false
	}
	if !(c.PFSEnabled == false) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.Expires == 0) {
		return false
	}
	if !(c.TestMode == false) {
		return false
	}
	if !(c.ThisDC == 0) {
		return false
	}
	if !(c.DCOptions == nil) {
		return false
	}
	if !(c.DCTxtDomainName == "") {
		return false
	}
	if !(c.ChatSizeMax == 0) {
		return false
	}
	if !(c.MegagroupSizeMax == 0) {
		return false
	}
	if !(c.ForwardedCountMax == 0) {
		return false
	}
	if !(c.OnlineUpdatePeriodMs == 0) {
		return false
	}
	if !(c.OfflineBlurTimeoutMs == 0) {
		return false
	}
	if !(c.OfflineIdleTimeoutMs == 0) {
		return false
	}
	if !(c.OnlineCloudTimeoutMs == 0) {
		return false
	}
	if !(c.NotifyCloudDelayMs == 0) {
		return false
	}
	if !(c.NotifyDefaultDelayMs == 0) {
		return false
	}
	if !(c.PushChatPeriodMs == 0) {
		return false
	}
	if !(c.PushChatLimit == 0) {
		return false
	}
	if !(c.SavedGifsLimit == 0) {
		return false
	}
	if !(c.EditTimeLimit == 0) {
		return false
	}
	if !(c.RevokeTimeLimit == 0) {
		return false
	}
	if !(c.RevokePmTimeLimit == 0) {
		return false
	}
	if !(c.RatingEDecay == 0) {
		return false
	}
	if !(c.StickersRecentLimit == 0) {
		return false
	}
	if !(c.StickersFavedLimit == 0) {
		return false
	}
	if !(c.ChannelsReadMediaPeriod == 0) {
		return false
	}
	if !(c.TmpSessions == 0) {
		return false
	}
	if !(c.PinnedDialogsCountMax == 0) {
		return false
	}
	if !(c.PinnedInfolderCountMax == 0) {
		return false
	}
	if !(c.CallReceiveTimeoutMs == 0) {
		return false
	}
	if !(c.CallRingTimeoutMs == 0) {
		return false
	}
	if !(c.CallConnectTimeoutMs == 0) {
		return false
	}
	if !(c.CallPacketTimeoutMs == 0) {
		return false
	}
	if !(c.MeURLPrefix == "") {
		return false
	}
	if !(c.AutoupdateURLPrefix == "") {
		return false
	}
	if !(c.GifSearchUsername == "") {
		return false
	}
	if !(c.VenueSearchUsername == "") {
		return false
	}
	if !(c.ImgSearchUsername == "") {
		return false
	}
	if !(c.StaticMapsProvider == "") {
		return false
	}
	if !(c.CaptionLengthMax == 0) {
		return false
	}
	if !(c.MessageLengthMax == 0) {
		return false
	}
	if !(c.WebfileDCID == 0) {
		return false
	}
	if !(c.SuggestedLangCode == "") {
		return false
	}
	if !(c.LangPackVersion == 0) {
		return false
	}
	if !(c.BaseLangPackVersion == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Config) String() string {
	if c == nil {
		return "Config(nil)"
	}
	type Alias Config
	return fmt.Sprintf("Config%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Config) TypeID() uint32 {
	return ConfigTypeID
}

// TypeName returns name of type in TL schema.
func (*Config) TypeName() string {
	return "config"
}

// TypeInfo returns info about TL type.
func (c *Config) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "config",
		ID:   ConfigTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhonecallsEnabled",
			SchemaName: "phonecalls_enabled",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "DefaultP2PContacts",
			SchemaName: "default_p2p_contacts",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "PreloadFeaturedStickers",
			SchemaName: "preload_featured_stickers",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "IgnorePhoneEntities",
			SchemaName: "ignore_phone_entities",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "RevokePmInbox",
			SchemaName: "revoke_pm_inbox",
			Null:       !c.Flags.Has(6),
		},
		{
			Name:       "BlockedMode",
			SchemaName: "blocked_mode",
			Null:       !c.Flags.Has(8),
		},
		{
			Name:       "PFSEnabled",
			SchemaName: "pfs_enabled",
			Null:       !c.Flags.Has(13),
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
		{
			Name:       "TestMode",
			SchemaName: "test_mode",
		},
		{
			Name:       "ThisDC",
			SchemaName: "this_dc",
		},
		{
			Name:       "DCOptions",
			SchemaName: "dc_options",
		},
		{
			Name:       "DCTxtDomainName",
			SchemaName: "dc_txt_domain_name",
		},
		{
			Name:       "ChatSizeMax",
			SchemaName: "chat_size_max",
		},
		{
			Name:       "MegagroupSizeMax",
			SchemaName: "megagroup_size_max",
		},
		{
			Name:       "ForwardedCountMax",
			SchemaName: "forwarded_count_max",
		},
		{
			Name:       "OnlineUpdatePeriodMs",
			SchemaName: "online_update_period_ms",
		},
		{
			Name:       "OfflineBlurTimeoutMs",
			SchemaName: "offline_blur_timeout_ms",
		},
		{
			Name:       "OfflineIdleTimeoutMs",
			SchemaName: "offline_idle_timeout_ms",
		},
		{
			Name:       "OnlineCloudTimeoutMs",
			SchemaName: "online_cloud_timeout_ms",
		},
		{
			Name:       "NotifyCloudDelayMs",
			SchemaName: "notify_cloud_delay_ms",
		},
		{
			Name:       "NotifyDefaultDelayMs",
			SchemaName: "notify_default_delay_ms",
		},
		{
			Name:       "PushChatPeriodMs",
			SchemaName: "push_chat_period_ms",
		},
		{
			Name:       "PushChatLimit",
			SchemaName: "push_chat_limit",
		},
		{
			Name:       "SavedGifsLimit",
			SchemaName: "saved_gifs_limit",
		},
		{
			Name:       "EditTimeLimit",
			SchemaName: "edit_time_limit",
		},
		{
			Name:       "RevokeTimeLimit",
			SchemaName: "revoke_time_limit",
		},
		{
			Name:       "RevokePmTimeLimit",
			SchemaName: "revoke_pm_time_limit",
		},
		{
			Name:       "RatingEDecay",
			SchemaName: "rating_e_decay",
		},
		{
			Name:       "StickersRecentLimit",
			SchemaName: "stickers_recent_limit",
		},
		{
			Name:       "StickersFavedLimit",
			SchemaName: "stickers_faved_limit",
		},
		{
			Name:       "ChannelsReadMediaPeriod",
			SchemaName: "channels_read_media_period",
		},
		{
			Name:       "TmpSessions",
			SchemaName: "tmp_sessions",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "PinnedDialogsCountMax",
			SchemaName: "pinned_dialogs_count_max",
		},
		{
			Name:       "PinnedInfolderCountMax",
			SchemaName: "pinned_infolder_count_max",
		},
		{
			Name:       "CallReceiveTimeoutMs",
			SchemaName: "call_receive_timeout_ms",
		},
		{
			Name:       "CallRingTimeoutMs",
			SchemaName: "call_ring_timeout_ms",
		},
		{
			Name:       "CallConnectTimeoutMs",
			SchemaName: "call_connect_timeout_ms",
		},
		{
			Name:       "CallPacketTimeoutMs",
			SchemaName: "call_packet_timeout_ms",
		},
		{
			Name:       "MeURLPrefix",
			SchemaName: "me_url_prefix",
		},
		{
			Name:       "AutoupdateURLPrefix",
			SchemaName: "autoupdate_url_prefix",
			Null:       !c.Flags.Has(7),
		},
		{
			Name:       "GifSearchUsername",
			SchemaName: "gif_search_username",
			Null:       !c.Flags.Has(9),
		},
		{
			Name:       "VenueSearchUsername",
			SchemaName: "venue_search_username",
			Null:       !c.Flags.Has(10),
		},
		{
			Name:       "ImgSearchUsername",
			SchemaName: "img_search_username",
			Null:       !c.Flags.Has(11),
		},
		{
			Name:       "StaticMapsProvider",
			SchemaName: "static_maps_provider",
			Null:       !c.Flags.Has(12),
		},
		{
			Name:       "CaptionLengthMax",
			SchemaName: "caption_length_max",
		},
		{
			Name:       "MessageLengthMax",
			SchemaName: "message_length_max",
		},
		{
			Name:       "WebfileDCID",
			SchemaName: "webfile_dc_id",
		},
		{
			Name:       "SuggestedLangCode",
			SchemaName: "suggested_lang_code",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "LangPackVersion",
			SchemaName: "lang_pack_version",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "BaseLangPackVersion",
			SchemaName: "base_lang_pack_version",
			Null:       !c.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *Config) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode config#330b4067 as nil")
	}
	b.PutID(ConfigTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *Config) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode config#330b4067 as nil")
	}
	if !(c.PhonecallsEnabled == false) {
		c.Flags.Set(1)
	}
	if !(c.DefaultP2PContacts == false) {
		c.Flags.Set(3)
	}
	if !(c.PreloadFeaturedStickers == false) {
		c.Flags.Set(4)
	}
	if !(c.IgnorePhoneEntities == false) {
		c.Flags.Set(5)
	}
	if !(c.RevokePmInbox == false) {
		c.Flags.Set(6)
	}
	if !(c.BlockedMode == false) {
		c.Flags.Set(8)
	}
	if !(c.PFSEnabled == false) {
		c.Flags.Set(13)
	}
	if !(c.TmpSessions == 0) {
		c.Flags.Set(0)
	}
	if !(c.AutoupdateURLPrefix == "") {
		c.Flags.Set(7)
	}
	if !(c.GifSearchUsername == "") {
		c.Flags.Set(9)
	}
	if !(c.VenueSearchUsername == "") {
		c.Flags.Set(10)
	}
	if !(c.ImgSearchUsername == "") {
		c.Flags.Set(11)
	}
	if !(c.StaticMapsProvider == "") {
		c.Flags.Set(12)
	}
	if !(c.SuggestedLangCode == "") {
		c.Flags.Set(2)
	}
	if !(c.LangPackVersion == 0) {
		c.Flags.Set(2)
	}
	if !(c.BaseLangPackVersion == 0) {
		c.Flags.Set(2)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode config#330b4067: field flags: %w", err)
	}
	b.PutInt(c.Date)
	b.PutInt(c.Expires)
	b.PutBool(c.TestMode)
	b.PutInt(c.ThisDC)
	b.PutVectorHeader(len(c.DCOptions))
	for idx, v := range c.DCOptions {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode config#330b4067: field dc_options element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.DCTxtDomainName)
	b.PutInt(c.ChatSizeMax)
	b.PutInt(c.MegagroupSizeMax)
	b.PutInt(c.ForwardedCountMax)
	b.PutInt(c.OnlineUpdatePeriodMs)
	b.PutInt(c.OfflineBlurTimeoutMs)
	b.PutInt(c.OfflineIdleTimeoutMs)
	b.PutInt(c.OnlineCloudTimeoutMs)
	b.PutInt(c.NotifyCloudDelayMs)
	b.PutInt(c.NotifyDefaultDelayMs)
	b.PutInt(c.PushChatPeriodMs)
	b.PutInt(c.PushChatLimit)
	b.PutInt(c.SavedGifsLimit)
	b.PutInt(c.EditTimeLimit)
	b.PutInt(c.RevokeTimeLimit)
	b.PutInt(c.RevokePmTimeLimit)
	b.PutInt(c.RatingEDecay)
	b.PutInt(c.StickersRecentLimit)
	b.PutInt(c.StickersFavedLimit)
	b.PutInt(c.ChannelsReadMediaPeriod)
	if c.Flags.Has(0) {
		b.PutInt(c.TmpSessions)
	}
	b.PutInt(c.PinnedDialogsCountMax)
	b.PutInt(c.PinnedInfolderCountMax)
	b.PutInt(c.CallReceiveTimeoutMs)
	b.PutInt(c.CallRingTimeoutMs)
	b.PutInt(c.CallConnectTimeoutMs)
	b.PutInt(c.CallPacketTimeoutMs)
	b.PutString(c.MeURLPrefix)
	if c.Flags.Has(7) {
		b.PutString(c.AutoupdateURLPrefix)
	}
	if c.Flags.Has(9) {
		b.PutString(c.GifSearchUsername)
	}
	if c.Flags.Has(10) {
		b.PutString(c.VenueSearchUsername)
	}
	if c.Flags.Has(11) {
		b.PutString(c.ImgSearchUsername)
	}
	if c.Flags.Has(12) {
		b.PutString(c.StaticMapsProvider)
	}
	b.PutInt(c.CaptionLengthMax)
	b.PutInt(c.MessageLengthMax)
	b.PutInt(c.WebfileDCID)
	if c.Flags.Has(2) {
		b.PutString(c.SuggestedLangCode)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.LangPackVersion)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.BaseLangPackVersion)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *Config) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode config#330b4067 to nil")
	}
	if err := b.ConsumeID(ConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode config#330b4067: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *Config) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode config#330b4067 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field flags: %w", err)
		}
	}
	c.PhonecallsEnabled = c.Flags.Has(1)
	c.DefaultP2PContacts = c.Flags.Has(3)
	c.PreloadFeaturedStickers = c.Flags.Has(4)
	c.IgnorePhoneEntities = c.Flags.Has(5)
	c.RevokePmInbox = c.Flags.Has(6)
	c.BlockedMode = c.Flags.Has(8)
	c.PFSEnabled = c.Flags.Has(13)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field expires: %w", err)
		}
		c.Expires = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field test_mode: %w", err)
		}
		c.TestMode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field this_dc: %w", err)
		}
		c.ThisDC = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field dc_options: %w", err)
		}

		if headerLen > 0 {
			c.DCOptions = make([]DCOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DCOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode config#330b4067: field dc_options: %w", err)
			}
			c.DCOptions = append(c.DCOptions, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field dc_txt_domain_name: %w", err)
		}
		c.DCTxtDomainName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field chat_size_max: %w", err)
		}
		c.ChatSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field megagroup_size_max: %w", err)
		}
		c.MegagroupSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field forwarded_count_max: %w", err)
		}
		c.ForwardedCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field online_update_period_ms: %w", err)
		}
		c.OnlineUpdatePeriodMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field offline_blur_timeout_ms: %w", err)
		}
		c.OfflineBlurTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field offline_idle_timeout_ms: %w", err)
		}
		c.OfflineIdleTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field online_cloud_timeout_ms: %w", err)
		}
		c.OnlineCloudTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field notify_cloud_delay_ms: %w", err)
		}
		c.NotifyCloudDelayMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field notify_default_delay_ms: %w", err)
		}
		c.NotifyDefaultDelayMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field push_chat_period_ms: %w", err)
		}
		c.PushChatPeriodMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field push_chat_limit: %w", err)
		}
		c.PushChatLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field saved_gifs_limit: %w", err)
		}
		c.SavedGifsLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field edit_time_limit: %w", err)
		}
		c.EditTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field revoke_time_limit: %w", err)
		}
		c.RevokeTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field revoke_pm_time_limit: %w", err)
		}
		c.RevokePmTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field rating_e_decay: %w", err)
		}
		c.RatingEDecay = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field stickers_recent_limit: %w", err)
		}
		c.StickersRecentLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field stickers_faved_limit: %w", err)
		}
		c.StickersFavedLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field channels_read_media_period: %w", err)
		}
		c.ChannelsReadMediaPeriod = value
	}
	if c.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field tmp_sessions: %w", err)
		}
		c.TmpSessions = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field pinned_dialogs_count_max: %w", err)
		}
		c.PinnedDialogsCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field pinned_infolder_count_max: %w", err)
		}
		c.PinnedInfolderCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_receive_timeout_ms: %w", err)
		}
		c.CallReceiveTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_ring_timeout_ms: %w", err)
		}
		c.CallRingTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_connect_timeout_ms: %w", err)
		}
		c.CallConnectTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_packet_timeout_ms: %w", err)
		}
		c.CallPacketTimeoutMs = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field me_url_prefix: %w", err)
		}
		c.MeURLPrefix = value
	}
	if c.Flags.Has(7) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field autoupdate_url_prefix: %w", err)
		}
		c.AutoupdateURLPrefix = value
	}
	if c.Flags.Has(9) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field gif_search_username: %w", err)
		}
		c.GifSearchUsername = value
	}
	if c.Flags.Has(10) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field venue_search_username: %w", err)
		}
		c.VenueSearchUsername = value
	}
	if c.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field img_search_username: %w", err)
		}
		c.ImgSearchUsername = value
	}
	if c.Flags.Has(12) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field static_maps_provider: %w", err)
		}
		c.StaticMapsProvider = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field caption_length_max: %w", err)
		}
		c.CaptionLengthMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field message_length_max: %w", err)
		}
		c.MessageLengthMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field webfile_dc_id: %w", err)
		}
		c.WebfileDCID = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field suggested_lang_code: %w", err)
		}
		c.SuggestedLangCode = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field lang_pack_version: %w", err)
		}
		c.LangPackVersion = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field base_lang_pack_version: %w", err)
		}
		c.BaseLangPackVersion = value
	}
	return nil
}

// SetPhonecallsEnabled sets value of PhonecallsEnabled conditional field.
func (c *Config) SetPhonecallsEnabled(value bool) {
	if value {
		c.Flags.Set(1)
		c.PhonecallsEnabled = true
	} else {
		c.Flags.Unset(1)
		c.PhonecallsEnabled = false
	}
}

// GetPhonecallsEnabled returns value of PhonecallsEnabled conditional field.
func (c *Config) GetPhonecallsEnabled() (value bool) {
	return c.Flags.Has(1)
}

// SetDefaultP2PContacts sets value of DefaultP2PContacts conditional field.
func (c *Config) SetDefaultP2PContacts(value bool) {
	if value {
		c.Flags.Set(3)
		c.DefaultP2PContacts = true
	} else {
		c.Flags.Unset(3)
		c.DefaultP2PContacts = false
	}
}

// GetDefaultP2PContacts returns value of DefaultP2PContacts conditional field.
func (c *Config) GetDefaultP2PContacts() (value bool) {
	return c.Flags.Has(3)
}

// SetPreloadFeaturedStickers sets value of PreloadFeaturedStickers conditional field.
func (c *Config) SetPreloadFeaturedStickers(value bool) {
	if value {
		c.Flags.Set(4)
		c.PreloadFeaturedStickers = true
	} else {
		c.Flags.Unset(4)
		c.PreloadFeaturedStickers = false
	}
}

// GetPreloadFeaturedStickers returns value of PreloadFeaturedStickers conditional field.
func (c *Config) GetPreloadFeaturedStickers() (value bool) {
	return c.Flags.Has(4)
}

// SetIgnorePhoneEntities sets value of IgnorePhoneEntities conditional field.
func (c *Config) SetIgnorePhoneEntities(value bool) {
	if value {
		c.Flags.Set(5)
		c.IgnorePhoneEntities = true
	} else {
		c.Flags.Unset(5)
		c.IgnorePhoneEntities = false
	}
}

// GetIgnorePhoneEntities returns value of IgnorePhoneEntities conditional field.
func (c *Config) GetIgnorePhoneEntities() (value bool) {
	return c.Flags.Has(5)
}

// SetRevokePmInbox sets value of RevokePmInbox conditional field.
func (c *Config) SetRevokePmInbox(value bool) {
	if value {
		c.Flags.Set(6)
		c.RevokePmInbox = true
	} else {
		c.Flags.Unset(6)
		c.RevokePmInbox = false
	}
}

// GetRevokePmInbox returns value of RevokePmInbox conditional field.
func (c *Config) GetRevokePmInbox() (value bool) {
	return c.Flags.Has(6)
}

// SetBlockedMode sets value of BlockedMode conditional field.
func (c *Config) SetBlockedMode(value bool) {
	if value {
		c.Flags.Set(8)
		c.BlockedMode = true
	} else {
		c.Flags.Unset(8)
		c.BlockedMode = false
	}
}

// GetBlockedMode returns value of BlockedMode conditional field.
func (c *Config) GetBlockedMode() (value bool) {
	return c.Flags.Has(8)
}

// SetPFSEnabled sets value of PFSEnabled conditional field.
func (c *Config) SetPFSEnabled(value bool) {
	if value {
		c.Flags.Set(13)
		c.PFSEnabled = true
	} else {
		c.Flags.Unset(13)
		c.PFSEnabled = false
	}
}

// GetPFSEnabled returns value of PFSEnabled conditional field.
func (c *Config) GetPFSEnabled() (value bool) {
	return c.Flags.Has(13)
}

// GetDate returns value of Date field.
func (c *Config) GetDate() (value int) {
	return c.Date
}

// GetExpires returns value of Expires field.
func (c *Config) GetExpires() (value int) {
	return c.Expires
}

// GetTestMode returns value of TestMode field.
func (c *Config) GetTestMode() (value bool) {
	return c.TestMode
}

// GetThisDC returns value of ThisDC field.
func (c *Config) GetThisDC() (value int) {
	return c.ThisDC
}

// GetDCOptions returns value of DCOptions field.
func (c *Config) GetDCOptions() (value []DCOption) {
	return c.DCOptions
}

// GetDCTxtDomainName returns value of DCTxtDomainName field.
func (c *Config) GetDCTxtDomainName() (value string) {
	return c.DCTxtDomainName
}

// GetChatSizeMax returns value of ChatSizeMax field.
func (c *Config) GetChatSizeMax() (value int) {
	return c.ChatSizeMax
}

// GetMegagroupSizeMax returns value of MegagroupSizeMax field.
func (c *Config) GetMegagroupSizeMax() (value int) {
	return c.MegagroupSizeMax
}

// GetForwardedCountMax returns value of ForwardedCountMax field.
func (c *Config) GetForwardedCountMax() (value int) {
	return c.ForwardedCountMax
}

// GetOnlineUpdatePeriodMs returns value of OnlineUpdatePeriodMs field.
func (c *Config) GetOnlineUpdatePeriodMs() (value int) {
	return c.OnlineUpdatePeriodMs
}

// GetOfflineBlurTimeoutMs returns value of OfflineBlurTimeoutMs field.
func (c *Config) GetOfflineBlurTimeoutMs() (value int) {
	return c.OfflineBlurTimeoutMs
}

// GetOfflineIdleTimeoutMs returns value of OfflineIdleTimeoutMs field.
func (c *Config) GetOfflineIdleTimeoutMs() (value int) {
	return c.OfflineIdleTimeoutMs
}

// GetOnlineCloudTimeoutMs returns value of OnlineCloudTimeoutMs field.
func (c *Config) GetOnlineCloudTimeoutMs() (value int) {
	return c.OnlineCloudTimeoutMs
}

// GetNotifyCloudDelayMs returns value of NotifyCloudDelayMs field.
func (c *Config) GetNotifyCloudDelayMs() (value int) {
	return c.NotifyCloudDelayMs
}

// GetNotifyDefaultDelayMs returns value of NotifyDefaultDelayMs field.
func (c *Config) GetNotifyDefaultDelayMs() (value int) {
	return c.NotifyDefaultDelayMs
}

// GetPushChatPeriodMs returns value of PushChatPeriodMs field.
func (c *Config) GetPushChatPeriodMs() (value int) {
	return c.PushChatPeriodMs
}

// GetPushChatLimit returns value of PushChatLimit field.
func (c *Config) GetPushChatLimit() (value int) {
	return c.PushChatLimit
}

// GetSavedGifsLimit returns value of SavedGifsLimit field.
func (c *Config) GetSavedGifsLimit() (value int) {
	return c.SavedGifsLimit
}

// GetEditTimeLimit returns value of EditTimeLimit field.
func (c *Config) GetEditTimeLimit() (value int) {
	return c.EditTimeLimit
}

// GetRevokeTimeLimit returns value of RevokeTimeLimit field.
func (c *Config) GetRevokeTimeLimit() (value int) {
	return c.RevokeTimeLimit
}

// GetRevokePmTimeLimit returns value of RevokePmTimeLimit field.
func (c *Config) GetRevokePmTimeLimit() (value int) {
	return c.RevokePmTimeLimit
}

// GetRatingEDecay returns value of RatingEDecay field.
func (c *Config) GetRatingEDecay() (value int) {
	return c.RatingEDecay
}

// GetStickersRecentLimit returns value of StickersRecentLimit field.
func (c *Config) GetStickersRecentLimit() (value int) {
	return c.StickersRecentLimit
}

// GetStickersFavedLimit returns value of StickersFavedLimit field.
func (c *Config) GetStickersFavedLimit() (value int) {
	return c.StickersFavedLimit
}

// GetChannelsReadMediaPeriod returns value of ChannelsReadMediaPeriod field.
func (c *Config) GetChannelsReadMediaPeriod() (value int) {
	return c.ChannelsReadMediaPeriod
}

// SetTmpSessions sets value of TmpSessions conditional field.
func (c *Config) SetTmpSessions(value int) {
	c.Flags.Set(0)
	c.TmpSessions = value
}

// GetTmpSessions returns value of TmpSessions conditional field and
// boolean which is true if field was set.
func (c *Config) GetTmpSessions() (value int, ok bool) {
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.TmpSessions, true
}

// GetPinnedDialogsCountMax returns value of PinnedDialogsCountMax field.
func (c *Config) GetPinnedDialogsCountMax() (value int) {
	return c.PinnedDialogsCountMax
}

// GetPinnedInfolderCountMax returns value of PinnedInfolderCountMax field.
func (c *Config) GetPinnedInfolderCountMax() (value int) {
	return c.PinnedInfolderCountMax
}

// GetCallReceiveTimeoutMs returns value of CallReceiveTimeoutMs field.
func (c *Config) GetCallReceiveTimeoutMs() (value int) {
	return c.CallReceiveTimeoutMs
}

// GetCallRingTimeoutMs returns value of CallRingTimeoutMs field.
func (c *Config) GetCallRingTimeoutMs() (value int) {
	return c.CallRingTimeoutMs
}

// GetCallConnectTimeoutMs returns value of CallConnectTimeoutMs field.
func (c *Config) GetCallConnectTimeoutMs() (value int) {
	return c.CallConnectTimeoutMs
}

// GetCallPacketTimeoutMs returns value of CallPacketTimeoutMs field.
func (c *Config) GetCallPacketTimeoutMs() (value int) {
	return c.CallPacketTimeoutMs
}

// GetMeURLPrefix returns value of MeURLPrefix field.
func (c *Config) GetMeURLPrefix() (value string) {
	return c.MeURLPrefix
}

// SetAutoupdateURLPrefix sets value of AutoupdateURLPrefix conditional field.
func (c *Config) SetAutoupdateURLPrefix(value string) {
	c.Flags.Set(7)
	c.AutoupdateURLPrefix = value
}

// GetAutoupdateURLPrefix returns value of AutoupdateURLPrefix conditional field and
// boolean which is true if field was set.
func (c *Config) GetAutoupdateURLPrefix() (value string, ok bool) {
	if !c.Flags.Has(7) {
		return value, false
	}
	return c.AutoupdateURLPrefix, true
}

// SetGifSearchUsername sets value of GifSearchUsername conditional field.
func (c *Config) SetGifSearchUsername(value string) {
	c.Flags.Set(9)
	c.GifSearchUsername = value
}

// GetGifSearchUsername returns value of GifSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetGifSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(9) {
		return value, false
	}
	return c.GifSearchUsername, true
}

// SetVenueSearchUsername sets value of VenueSearchUsername conditional field.
func (c *Config) SetVenueSearchUsername(value string) {
	c.Flags.Set(10)
	c.VenueSearchUsername = value
}

// GetVenueSearchUsername returns value of VenueSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetVenueSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(10) {
		return value, false
	}
	return c.VenueSearchUsername, true
}

// SetImgSearchUsername sets value of ImgSearchUsername conditional field.
func (c *Config) SetImgSearchUsername(value string) {
	c.Flags.Set(11)
	c.ImgSearchUsername = value
}

// GetImgSearchUsername returns value of ImgSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetImgSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(11) {
		return value, false
	}
	return c.ImgSearchUsername, true
}

// SetStaticMapsProvider sets value of StaticMapsProvider conditional field.
func (c *Config) SetStaticMapsProvider(value string) {
	c.Flags.Set(12)
	c.StaticMapsProvider = value
}

// GetStaticMapsProvider returns value of StaticMapsProvider conditional field and
// boolean which is true if field was set.
func (c *Config) GetStaticMapsProvider() (value string, ok bool) {
	if !c.Flags.Has(12) {
		return value, false
	}
	return c.StaticMapsProvider, true
}

// GetCaptionLengthMax returns value of CaptionLengthMax field.
func (c *Config) GetCaptionLengthMax() (value int) {
	return c.CaptionLengthMax
}

// GetMessageLengthMax returns value of MessageLengthMax field.
func (c *Config) GetMessageLengthMax() (value int) {
	return c.MessageLengthMax
}

// GetWebfileDCID returns value of WebfileDCID field.
func (c *Config) GetWebfileDCID() (value int) {
	return c.WebfileDCID
}

// SetSuggestedLangCode sets value of SuggestedLangCode conditional field.
func (c *Config) SetSuggestedLangCode(value string) {
	c.Flags.Set(2)
	c.SuggestedLangCode = value
}

// GetSuggestedLangCode returns value of SuggestedLangCode conditional field and
// boolean which is true if field was set.
func (c *Config) GetSuggestedLangCode() (value string, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.SuggestedLangCode, true
}

// SetLangPackVersion sets value of LangPackVersion conditional field.
func (c *Config) SetLangPackVersion(value int) {
	c.Flags.Set(2)
	c.LangPackVersion = value
}

// GetLangPackVersion returns value of LangPackVersion conditional field and
// boolean which is true if field was set.
func (c *Config) GetLangPackVersion() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.LangPackVersion, true
}

// SetBaseLangPackVersion sets value of BaseLangPackVersion conditional field.
func (c *Config) SetBaseLangPackVersion(value int) {
	c.Flags.Set(2)
	c.BaseLangPackVersion = value
}

// GetBaseLangPackVersion returns value of BaseLangPackVersion conditional field and
// boolean which is true if field was set.
func (c *Config) GetBaseLangPackVersion() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.BaseLangPackVersion, true
}
