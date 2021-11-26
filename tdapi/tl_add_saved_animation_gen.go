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

// AddSavedAnimationRequest represents TL type `addSavedAnimation#a44bf860`.
type AddSavedAnimationRequest struct {
	// The animation file to be added. Only animations known to the server (i.e.,
	// successfully sent via a message) can be added to the list
	Animation InputFileClass
}

// AddSavedAnimationRequestTypeID is TL type id of AddSavedAnimationRequest.
const AddSavedAnimationRequestTypeID = 0xa44bf860

// Ensuring interfaces in compile-time for AddSavedAnimationRequest.
var (
	_ bin.Encoder     = &AddSavedAnimationRequest{}
	_ bin.Decoder     = &AddSavedAnimationRequest{}
	_ bin.BareEncoder = &AddSavedAnimationRequest{}
	_ bin.BareDecoder = &AddSavedAnimationRequest{}
)

func (a *AddSavedAnimationRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Animation == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddSavedAnimationRequest) String() string {
	if a == nil {
		return "AddSavedAnimationRequest(nil)"
	}
	type Alias AddSavedAnimationRequest
	return fmt.Sprintf("AddSavedAnimationRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddSavedAnimationRequest) TypeID() uint32 {
	return AddSavedAnimationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddSavedAnimationRequest) TypeName() string {
	return "addSavedAnimation"
}

// TypeInfo returns info about TL type.
func (a *AddSavedAnimationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addSavedAnimation",
		ID:   AddSavedAnimationRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddSavedAnimationRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addSavedAnimation#a44bf860 as nil")
	}
	b.PutID(AddSavedAnimationRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddSavedAnimationRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addSavedAnimation#a44bf860 as nil")
	}
	if a.Animation == nil {
		return fmt.Errorf("unable to encode addSavedAnimation#a44bf860: field animation is nil")
	}
	if err := a.Animation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addSavedAnimation#a44bf860: field animation: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddSavedAnimationRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addSavedAnimation#a44bf860 to nil")
	}
	if err := b.ConsumeID(AddSavedAnimationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addSavedAnimation#a44bf860: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddSavedAnimationRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addSavedAnimation#a44bf860 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode addSavedAnimation#a44bf860: field animation: %w", err)
		}
		a.Animation = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddSavedAnimationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addSavedAnimation#a44bf860 as nil")
	}
	b.ObjStart()
	b.PutID("addSavedAnimation")
	b.FieldStart("animation")
	if a.Animation == nil {
		return fmt.Errorf("unable to encode addSavedAnimation#a44bf860: field animation is nil")
	}
	if err := a.Animation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addSavedAnimation#a44bf860: field animation: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddSavedAnimationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addSavedAnimation#a44bf860 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addSavedAnimation"); err != nil {
				return fmt.Errorf("unable to decode addSavedAnimation#a44bf860: %w", err)
			}
		case "animation":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode addSavedAnimation#a44bf860: field animation: %w", err)
			}
			a.Animation = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAnimation returns value of Animation field.
func (a *AddSavedAnimationRequest) GetAnimation() (value InputFileClass) {
	return a.Animation
}

// AddSavedAnimation invokes method addSavedAnimation#a44bf860 returning error if any.
func (c *Client) AddSavedAnimation(ctx context.Context, animation InputFileClass) error {
	var ok Ok

	request := &AddSavedAnimationRequest{
		Animation: animation,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}