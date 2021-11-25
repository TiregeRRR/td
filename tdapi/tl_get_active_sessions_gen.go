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

// GetActiveSessionsRequest represents TL type `getActiveSessions#42bd6d3e`.
type GetActiveSessionsRequest struct {
}

// GetActiveSessionsRequestTypeID is TL type id of GetActiveSessionsRequest.
const GetActiveSessionsRequestTypeID = 0x42bd6d3e

// Ensuring interfaces in compile-time for GetActiveSessionsRequest.
var (
	_ bin.Encoder     = &GetActiveSessionsRequest{}
	_ bin.Decoder     = &GetActiveSessionsRequest{}
	_ bin.BareEncoder = &GetActiveSessionsRequest{}
	_ bin.BareDecoder = &GetActiveSessionsRequest{}
)

func (g *GetActiveSessionsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetActiveSessionsRequest) String() string {
	if g == nil {
		return "GetActiveSessionsRequest(nil)"
	}
	type Alias GetActiveSessionsRequest
	return fmt.Sprintf("GetActiveSessionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetActiveSessionsRequest) TypeID() uint32 {
	return GetActiveSessionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetActiveSessionsRequest) TypeName() string {
	return "getActiveSessions"
}

// TypeInfo returns info about TL type.
func (g *GetActiveSessionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getActiveSessions",
		ID:   GetActiveSessionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetActiveSessionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getActiveSessions#42bd6d3e as nil")
	}
	b.PutID(GetActiveSessionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetActiveSessionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getActiveSessions#42bd6d3e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetActiveSessionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getActiveSessions#42bd6d3e to nil")
	}
	if err := b.ConsumeID(GetActiveSessionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getActiveSessions#42bd6d3e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetActiveSessionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getActiveSessions#42bd6d3e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetActiveSessionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getActiveSessions#42bd6d3e as nil")
	}
	b.ObjStart()
	b.PutID("getActiveSessions")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetActiveSessionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getActiveSessions#42bd6d3e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getActiveSessions"); err != nil {
				return fmt.Errorf("unable to decode getActiveSessions#42bd6d3e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetActiveSessions invokes method getActiveSessions#42bd6d3e returning error if any.
func (c *Client) GetActiveSessions(ctx context.Context) (*Sessions, error) {
	var result Sessions

	request := &GetActiveSessionsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
