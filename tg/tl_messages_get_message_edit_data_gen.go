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

// MessagesGetMessageEditDataRequest represents TL type `messages.getMessageEditData#fda68d36`.
// Find out if a media message's caption can be edited
//
// See https://core.telegram.org/method/messages.getMessageEditData for reference.
type MessagesGetMessageEditDataRequest struct {
	// Peer where the media was sent
	Peer InputPeerClass
	// ID of message
	ID int
}

// MessagesGetMessageEditDataRequestTypeID is TL type id of MessagesGetMessageEditDataRequest.
const MessagesGetMessageEditDataRequestTypeID = 0xfda68d36

// Ensuring interfaces in compile-time for MessagesGetMessageEditDataRequest.
var (
	_ bin.Encoder     = &MessagesGetMessageEditDataRequest{}
	_ bin.Decoder     = &MessagesGetMessageEditDataRequest{}
	_ bin.BareEncoder = &MessagesGetMessageEditDataRequest{}
	_ bin.BareDecoder = &MessagesGetMessageEditDataRequest{}
)

func (g *MessagesGetMessageEditDataRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetMessageEditDataRequest) String() string {
	if g == nil {
		return "MessagesGetMessageEditDataRequest(nil)"
	}
	type Alias MessagesGetMessageEditDataRequest
	return fmt.Sprintf("MessagesGetMessageEditDataRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetMessageEditDataRequest from given interface.
func (g *MessagesGetMessageEditDataRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetMessageEditDataRequest) TypeID() uint32 {
	return MessagesGetMessageEditDataRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetMessageEditDataRequest) TypeName() string {
	return "messages.getMessageEditData"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetMessageEditDataRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getMessageEditData",
		ID:   MessagesGetMessageEditDataRequestTypeID,
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
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessageEditDataRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageEditData#fda68d36 as nil")
	}
	b.PutID(MessagesGetMessageEditDataRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetMessageEditDataRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageEditData#fda68d36 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessageEditData#fda68d36: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessageEditData#fda68d36: field peer: %w", err)
	}
	b.PutInt(g.ID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessageEditDataRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageEditData#fda68d36 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessageEditDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetMessageEditDataRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageEditData#fda68d36 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageEditData#fda68d36: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetMessageEditDataRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetMessageEditDataRequest) GetID() (value int) {
	return g.ID
}

// MessagesGetMessageEditData invokes method messages.getMessageEditData#fda68d36 returning error if any.
// Find out if a media message's caption can be edited
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 MESSAGE_AUTHOR_REQUIRED: Message author required
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getMessageEditData for reference.
func (c *Client) MessagesGetMessageEditData(ctx context.Context, request *MessagesGetMessageEditDataRequest) (*MessagesMessageEditData, error) {
	var result MessagesMessageEditData

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
