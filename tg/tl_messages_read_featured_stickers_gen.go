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

// MessagesReadFeaturedStickersRequest represents TL type `messages.readFeaturedStickers#5b118126`.
// Mark new featured stickers as read
//
// See https://core.telegram.org/method/messages.readFeaturedStickers for reference.
type MessagesReadFeaturedStickersRequest struct {
	// IDs of stickersets to mark as read
	ID []int64
}

// MessagesReadFeaturedStickersRequestTypeID is TL type id of MessagesReadFeaturedStickersRequest.
const MessagesReadFeaturedStickersRequestTypeID = 0x5b118126

// Ensuring interfaces in compile-time for MessagesReadFeaturedStickersRequest.
var (
	_ bin.Encoder     = &MessagesReadFeaturedStickersRequest{}
	_ bin.Decoder     = &MessagesReadFeaturedStickersRequest{}
	_ bin.BareEncoder = &MessagesReadFeaturedStickersRequest{}
	_ bin.BareDecoder = &MessagesReadFeaturedStickersRequest{}
)

func (r *MessagesReadFeaturedStickersRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadFeaturedStickersRequest) String() string {
	if r == nil {
		return "MessagesReadFeaturedStickersRequest(nil)"
	}
	type Alias MessagesReadFeaturedStickersRequest
	return fmt.Sprintf("MessagesReadFeaturedStickersRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadFeaturedStickersRequest from given interface.
func (r *MessagesReadFeaturedStickersRequest) FillFrom(from interface {
	GetID() (value []int64)
}) {
	r.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReadFeaturedStickersRequest) TypeID() uint32 {
	return MessagesReadFeaturedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReadFeaturedStickersRequest) TypeName() string {
	return "messages.readFeaturedStickers"
}

// TypeInfo returns info about TL type.
func (r *MessagesReadFeaturedStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.readFeaturedStickers",
		ID:   MessagesReadFeaturedStickersRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReadFeaturedStickersRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readFeaturedStickers#5b118126 as nil")
	}
	b.PutID(MessagesReadFeaturedStickersRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReadFeaturedStickersRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readFeaturedStickers#5b118126 as nil")
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadFeaturedStickersRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readFeaturedStickers#5b118126 to nil")
	}
	if err := b.ConsumeID(MessagesReadFeaturedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readFeaturedStickers#5b118126: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReadFeaturedStickersRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readFeaturedStickers#5b118126 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readFeaturedStickers#5b118126: field id: %w", err)
		}

		if headerLen > 0 {
			r.ID = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.readFeaturedStickers#5b118126: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (r *MessagesReadFeaturedStickersRequest) GetID() (value []int64) {
	return r.ID
}

// MessagesReadFeaturedStickers invokes method messages.readFeaturedStickers#5b118126 returning error if any.
// Mark new featured stickers as read
//
// See https://core.telegram.org/method/messages.readFeaturedStickers for reference.
func (c *Client) MessagesReadFeaturedStickers(ctx context.Context, id []int64) (bool, error) {
	var result BoolBox

	request := &MessagesReadFeaturedStickersRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
