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

// GetMessageRequest represents TL type `getMessage#9372c080`.
type GetMessageRequest struct {
	// Identifier of the chat the message belongs to
	ChatID int64
	// Identifier of the message to get
	MessageID int64
}

// GetMessageRequestTypeID is TL type id of GetMessageRequest.
const GetMessageRequestTypeID = 0x9372c080

// Ensuring interfaces in compile-time for GetMessageRequest.
var (
	_ bin.Encoder     = &GetMessageRequest{}
	_ bin.Decoder     = &GetMessageRequest{}
	_ bin.BareEncoder = &GetMessageRequest{}
	_ bin.BareDecoder = &GetMessageRequest{}
)

func (g *GetMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageRequest) String() string {
	if g == nil {
		return "GetMessageRequest(nil)"
	}
	type Alias GetMessageRequest
	return fmt.Sprintf("GetMessageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageRequest) TypeID() uint32 {
	return GetMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageRequest) TypeName() string {
	return "getMessage"
}

// TypeInfo returns info about TL type.
func (g *GetMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessage",
		ID:   GetMessageRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessage#9372c080 as nil")
	}
	b.PutID(GetMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessage#9372c080 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutLong(g.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessage#9372c080 to nil")
	}
	if err := b.ConsumeID(GetMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessage#9372c080: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessage#9372c080 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessage#9372c080: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessage#9372c080: field message_id: %w", err)
		}
		g.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessage#9372c080 as nil")
	}
	b.ObjStart()
	b.PutID("getMessage")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_id")
	b.PutLong(g.MessageID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessage#9372c080 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessage"); err != nil {
				return fmt.Errorf("unable to decode getMessage#9372c080: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessage#9372c080: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessage#9372c080: field message_id: %w", err)
			}
			g.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetMessage invokes method getMessage#9372c080 returning error if any.
func (c *Client) GetMessage(ctx context.Context, request *GetMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
