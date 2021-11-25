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

// MessagesReorderPinnedDialogsRequest represents TL type `messages.reorderPinnedDialogs#3b1adf37`.
// Reorder pinned dialogs
//
// See https://core.telegram.org/method/messages.reorderPinnedDialogs for reference.
type MessagesReorderPinnedDialogsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, dialogs pinned server-side but not present in the order field will be unpinned.
	Force bool
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
	// New dialog order
	Order []InputDialogPeerClass
}

// MessagesReorderPinnedDialogsRequestTypeID is TL type id of MessagesReorderPinnedDialogsRequest.
const MessagesReorderPinnedDialogsRequestTypeID = 0x3b1adf37

// Ensuring interfaces in compile-time for MessagesReorderPinnedDialogsRequest.
var (
	_ bin.Encoder     = &MessagesReorderPinnedDialogsRequest{}
	_ bin.Decoder     = &MessagesReorderPinnedDialogsRequest{}
	_ bin.BareEncoder = &MessagesReorderPinnedDialogsRequest{}
	_ bin.BareDecoder = &MessagesReorderPinnedDialogsRequest{}
)

func (r *MessagesReorderPinnedDialogsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Force == false) {
		return false
	}
	if !(r.FolderID == 0) {
		return false
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReorderPinnedDialogsRequest) String() string {
	if r == nil {
		return "MessagesReorderPinnedDialogsRequest(nil)"
	}
	type Alias MessagesReorderPinnedDialogsRequest
	return fmt.Sprintf("MessagesReorderPinnedDialogsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReorderPinnedDialogsRequest from given interface.
func (r *MessagesReorderPinnedDialogsRequest) FillFrom(from interface {
	GetForce() (value bool)
	GetFolderID() (value int)
	GetOrder() (value []InputDialogPeerClass)
}) {
	r.Force = from.GetForce()
	r.FolderID = from.GetFolderID()
	r.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReorderPinnedDialogsRequest) TypeID() uint32 {
	return MessagesReorderPinnedDialogsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReorderPinnedDialogsRequest) TypeName() string {
	return "messages.reorderPinnedDialogs"
}

// TypeInfo returns info about TL type.
func (r *MessagesReorderPinnedDialogsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reorderPinnedDialogs",
		ID:   MessagesReorderPinnedDialogsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Force",
			SchemaName: "force",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReorderPinnedDialogsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderPinnedDialogs#3b1adf37 as nil")
	}
	b.PutID(MessagesReorderPinnedDialogsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReorderPinnedDialogsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderPinnedDialogs#3b1adf37 as nil")
	}
	if !(r.Force == false) {
		r.Flags.Set(0)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reorderPinnedDialogs#3b1adf37: field flags: %w", err)
	}
	b.PutInt(r.FolderID)
	b.PutVectorHeader(len(r.Order))
	for idx, v := range r.Order {
		if v == nil {
			return fmt.Errorf("unable to encode messages.reorderPinnedDialogs#3b1adf37: field order element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.reorderPinnedDialogs#3b1adf37: field order element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReorderPinnedDialogsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderPinnedDialogs#3b1adf37 to nil")
	}
	if err := b.ConsumeID(MessagesReorderPinnedDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReorderPinnedDialogsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderPinnedDialogs#3b1adf37 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field flags: %w", err)
		}
	}
	r.Force = r.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field folder_id: %w", err)
		}
		r.FolderID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field order: %w", err)
		}

		if headerLen > 0 {
			r.Order = make([]InputDialogPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// SetForce sets value of Force conditional field.
func (r *MessagesReorderPinnedDialogsRequest) SetForce(value bool) {
	if value {
		r.Flags.Set(0)
		r.Force = true
	} else {
		r.Flags.Unset(0)
		r.Force = false
	}
}

// GetForce returns value of Force conditional field.
func (r *MessagesReorderPinnedDialogsRequest) GetForce() (value bool) {
	return r.Flags.Has(0)
}

// GetFolderID returns value of FolderID field.
func (r *MessagesReorderPinnedDialogsRequest) GetFolderID() (value int) {
	return r.FolderID
}

// GetOrder returns value of Order field.
func (r *MessagesReorderPinnedDialogsRequest) GetOrder() (value []InputDialogPeerClass) {
	return r.Order
}

// MapOrder returns field Order wrapped in InputDialogPeerClassArray helper.
func (r *MessagesReorderPinnedDialogsRequest) MapOrder() (value InputDialogPeerClassArray) {
	return InputDialogPeerClassArray(r.Order)
}

// MessagesReorderPinnedDialogs invokes method messages.reorderPinnedDialogs#3b1adf37 returning error if any.
// Reorder pinned dialogs
//
// Possible errors:
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.reorderPinnedDialogs for reference.
func (c *Client) MessagesReorderPinnedDialogs(ctx context.Context, request *MessagesReorderPinnedDialogsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
