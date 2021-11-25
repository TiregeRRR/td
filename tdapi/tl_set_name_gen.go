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

// SetNameRequest represents TL type `setName#66065f10`.
type SetNameRequest struct {
	// The new value of the first name for the user; 1-64 characters
	FirstName string
	// The new value of the optional last name for the user; 0-64 characters
	LastName string
}

// SetNameRequestTypeID is TL type id of SetNameRequest.
const SetNameRequestTypeID = 0x66065f10

// Ensuring interfaces in compile-time for SetNameRequest.
var (
	_ bin.Encoder     = &SetNameRequest{}
	_ bin.Decoder     = &SetNameRequest{}
	_ bin.BareEncoder = &SetNameRequest{}
	_ bin.BareDecoder = &SetNameRequest{}
)

func (s *SetNameRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FirstName == "") {
		return false
	}
	if !(s.LastName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetNameRequest) String() string {
	if s == nil {
		return "SetNameRequest(nil)"
	}
	type Alias SetNameRequest
	return fmt.Sprintf("SetNameRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetNameRequest) TypeID() uint32 {
	return SetNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetNameRequest) TypeName() string {
	return "setName"
}

// TypeInfo returns info about TL type.
func (s *SetNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setName",
		ID:   SetNameRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetNameRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setName#66065f10 as nil")
	}
	b.PutID(SetNameRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetNameRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setName#66065f10 as nil")
	}
	b.PutString(s.FirstName)
	b.PutString(s.LastName)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetNameRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setName#66065f10 to nil")
	}
	if err := b.ConsumeID(SetNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setName#66065f10: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetNameRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setName#66065f10 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setName#66065f10: field first_name: %w", err)
		}
		s.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setName#66065f10: field last_name: %w", err)
		}
		s.LastName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetNameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setName#66065f10 as nil")
	}
	b.ObjStart()
	b.PutID("setName")
	b.FieldStart("first_name")
	b.PutString(s.FirstName)
	b.FieldStart("last_name")
	b.PutString(s.LastName)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetNameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setName#66065f10 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setName"); err != nil {
				return fmt.Errorf("unable to decode setName#66065f10: %w", err)
			}
		case "first_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setName#66065f10: field first_name: %w", err)
			}
			s.FirstName = value
		case "last_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setName#66065f10: field last_name: %w", err)
			}
			s.LastName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFirstName returns value of FirstName field.
func (s *SetNameRequest) GetFirstName() (value string) {
	return s.FirstName
}

// GetLastName returns value of LastName field.
func (s *SetNameRequest) GetLastName() (value string) {
	return s.LastName
}

// SetName invokes method setName#66065f10 returning error if any.
func (c *Client) SetName(ctx context.Context, request *SetNameRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
