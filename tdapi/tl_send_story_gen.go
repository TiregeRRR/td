// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// SendStoryRequest represents TL type `sendStory#736858bb`.
type SendStoryRequest struct {
	// Identifier of the chat that will post the story
	ChatID int64
	// Content of the story
	Content InputStoryContentClass
	// Clickable rectangle areas to be shown on the story media; pass null if none
	Areas InputStoryAreas
	// Story caption; pass null to use an empty caption;
	// 0-getOption("story_caption_length_max") characters
	Caption FormattedText
	// The privacy settings for the story
	PrivacySettings StoryPrivacySettingsClass
	// Period after which the story is moved to archive, in seconds; must be one of 6 * 3600,
	// 12 * 3600, 86400, or 2 * 86400 for Telegram Premium users, and 86400 otherwise
	ActivePeriod int32
	// Full identifier of the original story, which content was used to create the story
	FromStoryFullID StoryFullID
	// Pass true to keep the story accessible after expiration
	IsPinned bool
	// Pass true if the content of the story must be protected from forwarding and
	// screenshotting
	ProtectContent bool
}

// SendStoryRequestTypeID is TL type id of SendStoryRequest.
const SendStoryRequestTypeID = 0x736858bb

// Ensuring interfaces in compile-time for SendStoryRequest.
var (
	_ bin.Encoder     = &SendStoryRequest{}
	_ bin.Decoder     = &SendStoryRequest{}
	_ bin.BareEncoder = &SendStoryRequest{}
	_ bin.BareDecoder = &SendStoryRequest{}
)

func (s *SendStoryRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Content == nil) {
		return false
	}
	if !(s.Areas.Zero()) {
		return false
	}
	if !(s.Caption.Zero()) {
		return false
	}
	if !(s.PrivacySettings == nil) {
		return false
	}
	if !(s.ActivePeriod == 0) {
		return false
	}
	if !(s.FromStoryFullID.Zero()) {
		return false
	}
	if !(s.IsPinned == false) {
		return false
	}
	if !(s.ProtectContent == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendStoryRequest) String() string {
	if s == nil {
		return "SendStoryRequest(nil)"
	}
	type Alias SendStoryRequest
	return fmt.Sprintf("SendStoryRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendStoryRequest) TypeID() uint32 {
	return SendStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendStoryRequest) TypeName() string {
	return "sendStory"
}

// TypeInfo returns info about TL type.
func (s *SendStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendStory",
		ID:   SendStoryRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
		{
			Name:       "Areas",
			SchemaName: "areas",
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
		},
		{
			Name:       "PrivacySettings",
			SchemaName: "privacy_settings",
		},
		{
			Name:       "ActivePeriod",
			SchemaName: "active_period",
		},
		{
			Name:       "FromStoryFullID",
			SchemaName: "from_story_full_id",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
		{
			Name:       "ProtectContent",
			SchemaName: "protect_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendStoryRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendStory#736858bb as nil")
	}
	b.PutID(SendStoryRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendStoryRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendStory#736858bb as nil")
	}
	b.PutInt53(s.ChatID)
	if s.Content == nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field content is nil")
	}
	if err := s.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field content: %w", err)
	}
	if err := s.Areas.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field areas: %w", err)
	}
	if err := s.Caption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field caption: %w", err)
	}
	if s.PrivacySettings == nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field privacy_settings is nil")
	}
	if err := s.PrivacySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field privacy_settings: %w", err)
	}
	b.PutInt32(s.ActivePeriod)
	if err := s.FromStoryFullID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field from_story_full_id: %w", err)
	}
	b.PutBool(s.IsPinned)
	b.PutBool(s.ProtectContent)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendStoryRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendStory#736858bb to nil")
	}
	if err := b.ConsumeID(SendStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendStory#736858bb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendStoryRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendStory#736858bb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := DecodeInputStoryContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field content: %w", err)
		}
		s.Content = value
	}
	{
		if err := s.Areas.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field areas: %w", err)
		}
	}
	{
		if err := s.Caption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field caption: %w", err)
		}
	}
	{
		value, err := DecodeStoryPrivacySettings(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field privacy_settings: %w", err)
		}
		s.PrivacySettings = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field active_period: %w", err)
		}
		s.ActivePeriod = value
	}
	{
		if err := s.FromStoryFullID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field from_story_full_id: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field is_pinned: %w", err)
		}
		s.IsPinned = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sendStory#736858bb: field protect_content: %w", err)
		}
		s.ProtectContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendStoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendStory#736858bb as nil")
	}
	b.ObjStart()
	b.PutID("sendStory")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("content")
	if s.Content == nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field content is nil")
	}
	if err := s.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field content: %w", err)
	}
	b.Comma()
	b.FieldStart("areas")
	if err := s.Areas.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field areas: %w", err)
	}
	b.Comma()
	b.FieldStart("caption")
	if err := s.Caption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field caption: %w", err)
	}
	b.Comma()
	b.FieldStart("privacy_settings")
	if s.PrivacySettings == nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field privacy_settings is nil")
	}
	if err := s.PrivacySettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field privacy_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("active_period")
	b.PutInt32(s.ActivePeriod)
	b.Comma()
	b.FieldStart("from_story_full_id")
	if err := s.FromStoryFullID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendStory#736858bb: field from_story_full_id: %w", err)
	}
	b.Comma()
	b.FieldStart("is_pinned")
	b.PutBool(s.IsPinned)
	b.Comma()
	b.FieldStart("protect_content")
	b.PutBool(s.ProtectContent)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendStoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendStory#736858bb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendStory"); err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field chat_id: %w", err)
			}
			s.ChatID = value
		case "content":
			value, err := DecodeTDLibJSONInputStoryContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field content: %w", err)
			}
			s.Content = value
		case "areas":
			if err := s.Areas.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field areas: %w", err)
			}
		case "caption":
			if err := s.Caption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field caption: %w", err)
			}
		case "privacy_settings":
			value, err := DecodeTDLibJSONStoryPrivacySettings(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field privacy_settings: %w", err)
			}
			s.PrivacySettings = value
		case "active_period":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field active_period: %w", err)
			}
			s.ActivePeriod = value
		case "from_story_full_id":
			if err := s.FromStoryFullID.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field from_story_full_id: %w", err)
			}
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field is_pinned: %w", err)
			}
			s.IsPinned = value
		case "protect_content":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sendStory#736858bb: field protect_content: %w", err)
			}
			s.ProtectContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SendStoryRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetContent returns value of Content field.
func (s *SendStoryRequest) GetContent() (value InputStoryContentClass) {
	if s == nil {
		return
	}
	return s.Content
}

// GetAreas returns value of Areas field.
func (s *SendStoryRequest) GetAreas() (value InputStoryAreas) {
	if s == nil {
		return
	}
	return s.Areas
}

// GetCaption returns value of Caption field.
func (s *SendStoryRequest) GetCaption() (value FormattedText) {
	if s == nil {
		return
	}
	return s.Caption
}

// GetPrivacySettings returns value of PrivacySettings field.
func (s *SendStoryRequest) GetPrivacySettings() (value StoryPrivacySettingsClass) {
	if s == nil {
		return
	}
	return s.PrivacySettings
}

// GetActivePeriod returns value of ActivePeriod field.
func (s *SendStoryRequest) GetActivePeriod() (value int32) {
	if s == nil {
		return
	}
	return s.ActivePeriod
}

// GetFromStoryFullID returns value of FromStoryFullID field.
func (s *SendStoryRequest) GetFromStoryFullID() (value StoryFullID) {
	if s == nil {
		return
	}
	return s.FromStoryFullID
}

// GetIsPinned returns value of IsPinned field.
func (s *SendStoryRequest) GetIsPinned() (value bool) {
	if s == nil {
		return
	}
	return s.IsPinned
}

// GetProtectContent returns value of ProtectContent field.
func (s *SendStoryRequest) GetProtectContent() (value bool) {
	if s == nil {
		return
	}
	return s.ProtectContent
}

// SendStory invokes method sendStory#736858bb returning error if any.
func (c *Client) SendStory(ctx context.Context, request *SendStoryRequest) (*Story, error) {
	var result Story

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
