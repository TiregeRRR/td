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

// SetCustomLanguagePackStringRequest represents TL type `setCustomLanguagePackString#4e762518`.
type SetCustomLanguagePackStringRequest struct {
	// Identifier of a previously added custom local language pack in the current
	// localization target
	LanguagePackID string
	// New language pack string
	NewString LanguagePackString
}

// SetCustomLanguagePackStringRequestTypeID is TL type id of SetCustomLanguagePackStringRequest.
const SetCustomLanguagePackStringRequestTypeID = 0x4e762518

// Ensuring interfaces in compile-time for SetCustomLanguagePackStringRequest.
var (
	_ bin.Encoder     = &SetCustomLanguagePackStringRequest{}
	_ bin.Decoder     = &SetCustomLanguagePackStringRequest{}
	_ bin.BareEncoder = &SetCustomLanguagePackStringRequest{}
	_ bin.BareDecoder = &SetCustomLanguagePackStringRequest{}
)

func (s *SetCustomLanguagePackStringRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.LanguagePackID == "") {
		return false
	}
	if !(s.NewString.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetCustomLanguagePackStringRequest) String() string {
	if s == nil {
		return "SetCustomLanguagePackStringRequest(nil)"
	}
	type Alias SetCustomLanguagePackStringRequest
	return fmt.Sprintf("SetCustomLanguagePackStringRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetCustomLanguagePackStringRequest) TypeID() uint32 {
	return SetCustomLanguagePackStringRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetCustomLanguagePackStringRequest) TypeName() string {
	return "setCustomLanguagePackString"
}

// TypeInfo returns info about TL type.
func (s *SetCustomLanguagePackStringRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setCustomLanguagePackString",
		ID:   SetCustomLanguagePackStringRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LanguagePackID",
			SchemaName: "language_pack_id",
		},
		{
			Name:       "NewString",
			SchemaName: "new_string",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetCustomLanguagePackStringRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setCustomLanguagePackString#4e762518 as nil")
	}
	b.PutID(SetCustomLanguagePackStringRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetCustomLanguagePackStringRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setCustomLanguagePackString#4e762518 as nil")
	}
	b.PutString(s.LanguagePackID)
	if err := s.NewString.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setCustomLanguagePackString#4e762518: field new_string: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetCustomLanguagePackStringRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setCustomLanguagePackString#4e762518 to nil")
	}
	if err := b.ConsumeID(SetCustomLanguagePackStringRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setCustomLanguagePackString#4e762518: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetCustomLanguagePackStringRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setCustomLanguagePackString#4e762518 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setCustomLanguagePackString#4e762518: field language_pack_id: %w", err)
		}
		s.LanguagePackID = value
	}
	{
		if err := s.NewString.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setCustomLanguagePackString#4e762518: field new_string: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetCustomLanguagePackStringRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setCustomLanguagePackString#4e762518 as nil")
	}
	b.ObjStart()
	b.PutID("setCustomLanguagePackString")
	b.FieldStart("language_pack_id")
	b.PutString(s.LanguagePackID)
	b.FieldStart("new_string")
	if err := s.NewString.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setCustomLanguagePackString#4e762518: field new_string: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetCustomLanguagePackStringRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setCustomLanguagePackString#4e762518 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setCustomLanguagePackString"); err != nil {
				return fmt.Errorf("unable to decode setCustomLanguagePackString#4e762518: %w", err)
			}
		case "language_pack_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setCustomLanguagePackString#4e762518: field language_pack_id: %w", err)
			}
			s.LanguagePackID = value
		case "new_string":
			if err := s.NewString.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setCustomLanguagePackString#4e762518: field new_string: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLanguagePackID returns value of LanguagePackID field.
func (s *SetCustomLanguagePackStringRequest) GetLanguagePackID() (value string) {
	return s.LanguagePackID
}

// GetNewString returns value of NewString field.
func (s *SetCustomLanguagePackStringRequest) GetNewString() (value LanguagePackString) {
	return s.NewString
}

// SetCustomLanguagePackString invokes method setCustomLanguagePackString#4e762518 returning error if any.
func (c *Client) SetCustomLanguagePackString(ctx context.Context, request *SetCustomLanguagePackStringRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
