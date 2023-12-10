// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// PremiumGetUserBoostsRequest represents TL type `premium.getUserBoosts#39854d1f`.
//
// See https://core.telegram.org/method/premium.getUserBoosts for reference.
type PremiumGetUserBoostsRequest struct {
	// Peer field of PremiumGetUserBoostsRequest.
	Peer InputPeerClass
	// UserID field of PremiumGetUserBoostsRequest.
	UserID InputUserClass
}

// PremiumGetUserBoostsRequestTypeID is TL type id of PremiumGetUserBoostsRequest.
const PremiumGetUserBoostsRequestTypeID = 0x39854d1f

// Ensuring interfaces in compile-time for PremiumGetUserBoostsRequest.
var (
	_ bin.Encoder     = &PremiumGetUserBoostsRequest{}
	_ bin.Decoder     = &PremiumGetUserBoostsRequest{}
	_ bin.BareEncoder = &PremiumGetUserBoostsRequest{}
	_ bin.BareDecoder = &PremiumGetUserBoostsRequest{}
)

func (g *PremiumGetUserBoostsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PremiumGetUserBoostsRequest) String() string {
	if g == nil {
		return "PremiumGetUserBoostsRequest(nil)"
	}
	type Alias PremiumGetUserBoostsRequest
	return fmt.Sprintf("PremiumGetUserBoostsRequest%+v", Alias(*g))
}

// FillFrom fills PremiumGetUserBoostsRequest from given interface.
func (g *PremiumGetUserBoostsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetUserID() (value InputUserClass)
}) {
	g.Peer = from.GetPeer()
	g.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumGetUserBoostsRequest) TypeID() uint32 {
	return PremiumGetUserBoostsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumGetUserBoostsRequest) TypeName() string {
	return "premium.getUserBoosts"
}

// TypeInfo returns info about TL type.
func (g *PremiumGetUserBoostsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premium.getUserBoosts",
		ID:   PremiumGetUserBoostsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PremiumGetUserBoostsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode premium.getUserBoosts#39854d1f as nil")
	}
	b.PutID(PremiumGetUserBoostsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PremiumGetUserBoostsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode premium.getUserBoosts#39854d1f as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode premium.getUserBoosts#39854d1f: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premium.getUserBoosts#39854d1f: field peer: %w", err)
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode premium.getUserBoosts#39854d1f: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premium.getUserBoosts#39854d1f: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PremiumGetUserBoostsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode premium.getUserBoosts#39854d1f to nil")
	}
	if err := b.ConsumeID(PremiumGetUserBoostsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode premium.getUserBoosts#39854d1f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PremiumGetUserBoostsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode premium.getUserBoosts#39854d1f to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode premium.getUserBoosts#39854d1f: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode premium.getUserBoosts#39854d1f: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *PremiumGetUserBoostsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetUserID returns value of UserID field.
func (g *PremiumGetUserBoostsRequest) GetUserID() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.UserID
}

// PremiumGetUserBoosts invokes method premium.getUserBoosts#39854d1f returning error if any.
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/premium.getUserBoosts for reference.
// Can be used by bots.
func (c *Client) PremiumGetUserBoosts(ctx context.Context, request *PremiumGetUserBoostsRequest) (*PremiumBoostsList, error) {
	var result PremiumBoostsList

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
