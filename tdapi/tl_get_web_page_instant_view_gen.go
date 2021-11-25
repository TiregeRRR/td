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

// GetWebPageInstantViewRequest represents TL type `getWebPageInstantView#8b045689`.
type GetWebPageInstantViewRequest struct {
	// The web page URL
	URL string
	// If true, the full instant view for the web page will be returned
	ForceFull bool
}

// GetWebPageInstantViewRequestTypeID is TL type id of GetWebPageInstantViewRequest.
const GetWebPageInstantViewRequestTypeID = 0x8b045689

// Ensuring interfaces in compile-time for GetWebPageInstantViewRequest.
var (
	_ bin.Encoder     = &GetWebPageInstantViewRequest{}
	_ bin.Decoder     = &GetWebPageInstantViewRequest{}
	_ bin.BareEncoder = &GetWebPageInstantViewRequest{}
	_ bin.BareDecoder = &GetWebPageInstantViewRequest{}
)

func (g *GetWebPageInstantViewRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.URL == "") {
		return false
	}
	if !(g.ForceFull == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetWebPageInstantViewRequest) String() string {
	if g == nil {
		return "GetWebPageInstantViewRequest(nil)"
	}
	type Alias GetWebPageInstantViewRequest
	return fmt.Sprintf("GetWebPageInstantViewRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetWebPageInstantViewRequest) TypeID() uint32 {
	return GetWebPageInstantViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetWebPageInstantViewRequest) TypeName() string {
	return "getWebPageInstantView"
}

// TypeInfo returns info about TL type.
func (g *GetWebPageInstantViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getWebPageInstantView",
		ID:   GetWebPageInstantViewRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ForceFull",
			SchemaName: "force_full",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetWebPageInstantViewRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getWebPageInstantView#8b045689 as nil")
	}
	b.PutID(GetWebPageInstantViewRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetWebPageInstantViewRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getWebPageInstantView#8b045689 as nil")
	}
	b.PutString(g.URL)
	b.PutBool(g.ForceFull)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetWebPageInstantViewRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getWebPageInstantView#8b045689 to nil")
	}
	if err := b.ConsumeID(GetWebPageInstantViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getWebPageInstantView#8b045689: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetWebPageInstantViewRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getWebPageInstantView#8b045689 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getWebPageInstantView#8b045689: field url: %w", err)
		}
		g.URL = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getWebPageInstantView#8b045689: field force_full: %w", err)
		}
		g.ForceFull = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetWebPageInstantViewRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getWebPageInstantView#8b045689 as nil")
	}
	b.ObjStart()
	b.PutID("getWebPageInstantView")
	b.FieldStart("url")
	b.PutString(g.URL)
	b.FieldStart("force_full")
	b.PutBool(g.ForceFull)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetWebPageInstantViewRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getWebPageInstantView#8b045689 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getWebPageInstantView"); err != nil {
				return fmt.Errorf("unable to decode getWebPageInstantView#8b045689: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getWebPageInstantView#8b045689: field url: %w", err)
			}
			g.URL = value
		case "force_full":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getWebPageInstantView#8b045689: field force_full: %w", err)
			}
			g.ForceFull = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (g *GetWebPageInstantViewRequest) GetURL() (value string) {
	return g.URL
}

// GetForceFull returns value of ForceFull field.
func (g *GetWebPageInstantViewRequest) GetForceFull() (value bool) {
	return g.ForceFull
}

// GetWebPageInstantView invokes method getWebPageInstantView#8b045689 returning error if any.
func (c *Client) GetWebPageInstantView(ctx context.Context, request *GetWebPageInstantViewRequest) (*WebPageInstantView, error) {
	var result WebPageInstantView

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
