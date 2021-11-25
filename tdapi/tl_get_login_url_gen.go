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

// GetLoginURLRequest represents TL type `getLoginUrl#296c75e5`.
type GetLoginURLRequest struct {
	// Chat identifier of the message with the button
	ChatID int64
	// Message identifier of the message with the button
	MessageID int64
	// Button identifier
	ButtonID int32
	// True, if the user allowed the bot to send them messages
	AllowWriteAccess bool
}

// GetLoginURLRequestTypeID is TL type id of GetLoginURLRequest.
const GetLoginURLRequestTypeID = 0x296c75e5

// Ensuring interfaces in compile-time for GetLoginURLRequest.
var (
	_ bin.Encoder     = &GetLoginURLRequest{}
	_ bin.Decoder     = &GetLoginURLRequest{}
	_ bin.BareEncoder = &GetLoginURLRequest{}
	_ bin.BareDecoder = &GetLoginURLRequest{}
)

func (g *GetLoginURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.ButtonID == 0) {
		return false
	}
	if !(g.AllowWriteAccess == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLoginURLRequest) String() string {
	if g == nil {
		return "GetLoginURLRequest(nil)"
	}
	type Alias GetLoginURLRequest
	return fmt.Sprintf("GetLoginURLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLoginURLRequest) TypeID() uint32 {
	return GetLoginURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLoginURLRequest) TypeName() string {
	return "getLoginUrl"
}

// TypeInfo returns info about TL type.
func (g *GetLoginURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLoginUrl",
		ID:   GetLoginURLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ButtonID",
			SchemaName: "button_id",
		},
		{
			Name:       "AllowWriteAccess",
			SchemaName: "allow_write_access",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLoginURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLoginUrl#296c75e5 as nil")
	}
	b.PutID(GetLoginURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLoginURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLoginUrl#296c75e5 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutLong(g.MessageID)
	b.PutInt32(g.ButtonID)
	b.PutBool(g.AllowWriteAccess)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLoginURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLoginUrl#296c75e5 to nil")
	}
	if err := b.ConsumeID(GetLoginURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLoginUrl#296c75e5: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLoginURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLoginUrl#296c75e5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field button_id: %w", err)
		}
		g.ButtonID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field allow_write_access: %w", err)
		}
		g.AllowWriteAccess = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLoginURLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLoginUrl#296c75e5 as nil")
	}
	b.ObjStart()
	b.PutID("getLoginUrl")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_id")
	b.PutLong(g.MessageID)
	b.FieldStart("button_id")
	b.PutInt32(g.ButtonID)
	b.FieldStart("allow_write_access")
	b.PutBool(g.AllowWriteAccess)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLoginURLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLoginUrl#296c75e5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLoginUrl"); err != nil {
				return fmt.Errorf("unable to decode getLoginUrl#296c75e5: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field message_id: %w", err)
			}
			g.MessageID = value
		case "button_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field button_id: %w", err)
			}
			g.ButtonID = value
		case "allow_write_access":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrl#296c75e5: field allow_write_access: %w", err)
			}
			g.AllowWriteAccess = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetLoginURLRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetLoginURLRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetButtonID returns value of ButtonID field.
func (g *GetLoginURLRequest) GetButtonID() (value int32) {
	return g.ButtonID
}

// GetAllowWriteAccess returns value of AllowWriteAccess field.
func (g *GetLoginURLRequest) GetAllowWriteAccess() (value bool) {
	return g.AllowWriteAccess
}

// GetLoginURL invokes method getLoginUrl#296c75e5 returning error if any.
func (c *Client) GetLoginURL(ctx context.Context, request *GetLoginURLRequest) (*HTTPURL, error) {
	var result HTTPURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
