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

// GetChatStatisticsURLRequest represents TL type `getChatStatisticsUrl#426fc4ff`.
type GetChatStatisticsURLRequest struct {
	// Chat identifier
	ChatID int64
	// Parameters for the request
	Parameters string
	// Pass true if a URL with the dark theme must be returned
	IsDark bool
}

// GetChatStatisticsURLRequestTypeID is TL type id of GetChatStatisticsURLRequest.
const GetChatStatisticsURLRequestTypeID = 0x426fc4ff

// Ensuring interfaces in compile-time for GetChatStatisticsURLRequest.
var (
	_ bin.Encoder     = &GetChatStatisticsURLRequest{}
	_ bin.Decoder     = &GetChatStatisticsURLRequest{}
	_ bin.BareEncoder = &GetChatStatisticsURLRequest{}
	_ bin.BareDecoder = &GetChatStatisticsURLRequest{}
)

func (g *GetChatStatisticsURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Parameters == "") {
		return false
	}
	if !(g.IsDark == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatStatisticsURLRequest) String() string {
	if g == nil {
		return "GetChatStatisticsURLRequest(nil)"
	}
	type Alias GetChatStatisticsURLRequest
	return fmt.Sprintf("GetChatStatisticsURLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatStatisticsURLRequest) TypeID() uint32 {
	return GetChatStatisticsURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatStatisticsURLRequest) TypeName() string {
	return "getChatStatisticsUrl"
}

// TypeInfo returns info about TL type.
func (g *GetChatStatisticsURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatStatisticsUrl",
		ID:   GetChatStatisticsURLRequestTypeID,
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
			Name:       "Parameters",
			SchemaName: "parameters",
		},
		{
			Name:       "IsDark",
			SchemaName: "is_dark",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatStatisticsURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatStatisticsUrl#426fc4ff as nil")
	}
	b.PutID(GetChatStatisticsURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatStatisticsURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatStatisticsUrl#426fc4ff as nil")
	}
	b.PutLong(g.ChatID)
	b.PutString(g.Parameters)
	b.PutBool(g.IsDark)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatStatisticsURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatStatisticsUrl#426fc4ff to nil")
	}
	if err := b.ConsumeID(GetChatStatisticsURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatStatisticsURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatStatisticsUrl#426fc4ff to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: field parameters: %w", err)
		}
		g.Parameters = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: field is_dark: %w", err)
		}
		g.IsDark = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatStatisticsURLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatStatisticsUrl#426fc4ff as nil")
	}
	b.ObjStart()
	b.PutID("getChatStatisticsUrl")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("parameters")
	b.PutString(g.Parameters)
	b.FieldStart("is_dark")
	b.PutBool(g.IsDark)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatStatisticsURLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatStatisticsUrl#426fc4ff to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatStatisticsUrl"); err != nil {
				return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: field chat_id: %w", err)
			}
			g.ChatID = value
		case "parameters":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: field parameters: %w", err)
			}
			g.Parameters = value
		case "is_dark":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getChatStatisticsUrl#426fc4ff: field is_dark: %w", err)
			}
			g.IsDark = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatStatisticsURLRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetParameters returns value of Parameters field.
func (g *GetChatStatisticsURLRequest) GetParameters() (value string) {
	return g.Parameters
}

// GetIsDark returns value of IsDark field.
func (g *GetChatStatisticsURLRequest) GetIsDark() (value bool) {
	return g.IsDark
}

// GetChatStatisticsURL invokes method getChatStatisticsUrl#426fc4ff returning error if any.
func (c *Client) GetChatStatisticsURL(ctx context.Context, request *GetChatStatisticsURLRequest) (*HTTPURL, error) {
	var result HTTPURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
