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

// ToggleSupergroupIsBroadcastGroupRequest represents TL type `toggleSupergroupIsBroadcastGroup#9c3ec48a`.
type ToggleSupergroupIsBroadcastGroupRequest struct {
	// Identifier of the supergroup
	SupergroupID int32
}

// ToggleSupergroupIsBroadcastGroupRequestTypeID is TL type id of ToggleSupergroupIsBroadcastGroupRequest.
const ToggleSupergroupIsBroadcastGroupRequestTypeID = 0x9c3ec48a

// Ensuring interfaces in compile-time for ToggleSupergroupIsBroadcastGroupRequest.
var (
	_ bin.Encoder     = &ToggleSupergroupIsBroadcastGroupRequest{}
	_ bin.Decoder     = &ToggleSupergroupIsBroadcastGroupRequest{}
	_ bin.BareEncoder = &ToggleSupergroupIsBroadcastGroupRequest{}
	_ bin.BareDecoder = &ToggleSupergroupIsBroadcastGroupRequest{}
)

func (t *ToggleSupergroupIsBroadcastGroupRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSupergroupIsBroadcastGroupRequest) String() string {
	if t == nil {
		return "ToggleSupergroupIsBroadcastGroupRequest(nil)"
	}
	type Alias ToggleSupergroupIsBroadcastGroupRequest
	return fmt.Sprintf("ToggleSupergroupIsBroadcastGroupRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSupergroupIsBroadcastGroupRequest) TypeID() uint32 {
	return ToggleSupergroupIsBroadcastGroupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSupergroupIsBroadcastGroupRequest) TypeName() string {
	return "toggleSupergroupIsBroadcastGroup"
}

// TypeInfo returns info about TL type.
func (t *ToggleSupergroupIsBroadcastGroupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSupergroupIsBroadcastGroup",
		ID:   ToggleSupergroupIsBroadcastGroupRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSupergroupIsBroadcastGroupRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupIsBroadcastGroup#9c3ec48a as nil")
	}
	b.PutID(ToggleSupergroupIsBroadcastGroupRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSupergroupIsBroadcastGroupRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupIsBroadcastGroup#9c3ec48a as nil")
	}
	b.PutInt32(t.SupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSupergroupIsBroadcastGroupRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupIsBroadcastGroup#9c3ec48a to nil")
	}
	if err := b.ConsumeID(ToggleSupergroupIsBroadcastGroupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSupergroupIsBroadcastGroup#9c3ec48a: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSupergroupIsBroadcastGroupRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupIsBroadcastGroup#9c3ec48a to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupIsBroadcastGroup#9c3ec48a: field supergroup_id: %w", err)
		}
		t.SupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSupergroupIsBroadcastGroupRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupIsBroadcastGroup#9c3ec48a as nil")
	}
	b.ObjStart()
	b.PutID("toggleSupergroupIsBroadcastGroup")
	b.FieldStart("supergroup_id")
	b.PutInt32(t.SupergroupID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSupergroupIsBroadcastGroupRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupIsBroadcastGroup#9c3ec48a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSupergroupIsBroadcastGroup"); err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupIsBroadcastGroup#9c3ec48a: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupIsBroadcastGroup#9c3ec48a: field supergroup_id: %w", err)
			}
			t.SupergroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (t *ToggleSupergroupIsBroadcastGroupRequest) GetSupergroupID() (value int32) {
	return t.SupergroupID
}

// ToggleSupergroupIsBroadcastGroup invokes method toggleSupergroupIsBroadcastGroup#9c3ec48a returning error if any.
func (c *Client) ToggleSupergroupIsBroadcastGroup(ctx context.Context, supergroupid int32) error {
	var ok Ok

	request := &ToggleSupergroupIsBroadcastGroupRequest{
		SupergroupID: supergroupid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
