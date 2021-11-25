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

// SetChatNotificationSettingsRequest represents TL type `setChatNotificationSettings#2e531ffe`.
type SetChatNotificationSettingsRequest struct {
	// Chat identifier
	ChatID int64
	// New notification settings for the chat. If the chat is muted for more than 1 week, it
	// is considered to be muted forever
	NotificationSettings ChatNotificationSettings
}

// SetChatNotificationSettingsRequestTypeID is TL type id of SetChatNotificationSettingsRequest.
const SetChatNotificationSettingsRequestTypeID = 0x2e531ffe

// Ensuring interfaces in compile-time for SetChatNotificationSettingsRequest.
var (
	_ bin.Encoder     = &SetChatNotificationSettingsRequest{}
	_ bin.Decoder     = &SetChatNotificationSettingsRequest{}
	_ bin.BareEncoder = &SetChatNotificationSettingsRequest{}
	_ bin.BareDecoder = &SetChatNotificationSettingsRequest{}
)

func (s *SetChatNotificationSettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.NotificationSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatNotificationSettingsRequest) String() string {
	if s == nil {
		return "SetChatNotificationSettingsRequest(nil)"
	}
	type Alias SetChatNotificationSettingsRequest
	return fmt.Sprintf("SetChatNotificationSettingsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatNotificationSettingsRequest) TypeID() uint32 {
	return SetChatNotificationSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatNotificationSettingsRequest) TypeName() string {
	return "setChatNotificationSettings"
}

// TypeInfo returns info about TL type.
func (s *SetChatNotificationSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatNotificationSettings",
		ID:   SetChatNotificationSettingsRequestTypeID,
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
			Name:       "NotificationSettings",
			SchemaName: "notification_settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatNotificationSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatNotificationSettings#2e531ffe as nil")
	}
	b.PutID(SetChatNotificationSettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatNotificationSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatNotificationSettings#2e531ffe as nil")
	}
	b.PutLong(s.ChatID)
	if err := s.NotificationSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatNotificationSettings#2e531ffe: field notification_settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatNotificationSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatNotificationSettings#2e531ffe to nil")
	}
	if err := b.ConsumeID(SetChatNotificationSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatNotificationSettings#2e531ffe: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatNotificationSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatNotificationSettings#2e531ffe to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setChatNotificationSettings#2e531ffe: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		if err := s.NotificationSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setChatNotificationSettings#2e531ffe: field notification_settings: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatNotificationSettingsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatNotificationSettings#2e531ffe as nil")
	}
	b.ObjStart()
	b.PutID("setChatNotificationSettings")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("notification_settings")
	if err := s.NotificationSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatNotificationSettings#2e531ffe: field notification_settings: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatNotificationSettingsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatNotificationSettings#2e531ffe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatNotificationSettings"); err != nil {
				return fmt.Errorf("unable to decode setChatNotificationSettings#2e531ffe: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode setChatNotificationSettings#2e531ffe: field chat_id: %w", err)
			}
			s.ChatID = value
		case "notification_settings":
			if err := s.NotificationSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setChatNotificationSettings#2e531ffe: field notification_settings: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatNotificationSettingsRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetNotificationSettings returns value of NotificationSettings field.
func (s *SetChatNotificationSettingsRequest) GetNotificationSettings() (value ChatNotificationSettings) {
	return s.NotificationSettings
}

// SetChatNotificationSettings invokes method setChatNotificationSettings#2e531ffe returning error if any.
func (c *Client) SetChatNotificationSettings(ctx context.Context, request *SetChatNotificationSettingsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
