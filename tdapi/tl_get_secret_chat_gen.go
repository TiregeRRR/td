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

// GetSecretChatRequest represents TL type `getSecretChat#26b7e81`.
type GetSecretChatRequest struct {
	// Secret chat identifier
	SecretChatID int32
}

// GetSecretChatRequestTypeID is TL type id of GetSecretChatRequest.
const GetSecretChatRequestTypeID = 0x26b7e81

// Ensuring interfaces in compile-time for GetSecretChatRequest.
var (
	_ bin.Encoder     = &GetSecretChatRequest{}
	_ bin.Decoder     = &GetSecretChatRequest{}
	_ bin.BareEncoder = &GetSecretChatRequest{}
	_ bin.BareDecoder = &GetSecretChatRequest{}
)

func (g *GetSecretChatRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SecretChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSecretChatRequest) String() string {
	if g == nil {
		return "GetSecretChatRequest(nil)"
	}
	type Alias GetSecretChatRequest
	return fmt.Sprintf("GetSecretChatRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSecretChatRequest) TypeID() uint32 {
	return GetSecretChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSecretChatRequest) TypeName() string {
	return "getSecretChat"
}

// TypeInfo returns info about TL type.
func (g *GetSecretChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSecretChat",
		ID:   GetSecretChatRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SecretChatID",
			SchemaName: "secret_chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSecretChatRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSecretChat#26b7e81 as nil")
	}
	b.PutID(GetSecretChatRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSecretChatRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSecretChat#26b7e81 as nil")
	}
	b.PutInt32(g.SecretChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSecretChatRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSecretChat#26b7e81 to nil")
	}
	if err := b.ConsumeID(GetSecretChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSecretChat#26b7e81: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSecretChatRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSecretChat#26b7e81 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSecretChat#26b7e81: field secret_chat_id: %w", err)
		}
		g.SecretChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSecretChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSecretChat#26b7e81 as nil")
	}
	b.ObjStart()
	b.PutID("getSecretChat")
	b.FieldStart("secret_chat_id")
	b.PutInt32(g.SecretChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSecretChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSecretChat#26b7e81 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSecretChat"); err != nil {
				return fmt.Errorf("unable to decode getSecretChat#26b7e81: %w", err)
			}
		case "secret_chat_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSecretChat#26b7e81: field secret_chat_id: %w", err)
			}
			g.SecretChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSecretChatID returns value of SecretChatID field.
func (g *GetSecretChatRequest) GetSecretChatID() (value int32) {
	return g.SecretChatID
}

// GetSecretChat invokes method getSecretChat#26b7e81 returning error if any.
func (c *Client) GetSecretChat(ctx context.Context, secretchatid int32) (*SecretChat, error) {
	var result SecretChat

	request := &GetSecretChatRequest{
		SecretChatID: secretchatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
