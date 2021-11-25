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

// ShippingOption represents TL type `shippingOption#731bffce`.
type ShippingOption struct {
	// Shipping option identifier
	ID string
	// Option title
	Title string
	// A list of objects used to calculate the total shipping costs
	PriceParts []LabeledPricePart
}

// ShippingOptionTypeID is TL type id of ShippingOption.
const ShippingOptionTypeID = 0x731bffce

// Ensuring interfaces in compile-time for ShippingOption.
var (
	_ bin.Encoder     = &ShippingOption{}
	_ bin.Decoder     = &ShippingOption{}
	_ bin.BareEncoder = &ShippingOption{}
	_ bin.BareDecoder = &ShippingOption{}
)

func (s *ShippingOption) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.PriceParts == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ShippingOption) String() string {
	if s == nil {
		return "ShippingOption(nil)"
	}
	type Alias ShippingOption
	return fmt.Sprintf("ShippingOption%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ShippingOption) TypeID() uint32 {
	return ShippingOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*ShippingOption) TypeName() string {
	return "shippingOption"
}

// TypeInfo returns info about TL type.
func (s *ShippingOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "shippingOption",
		ID:   ShippingOptionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "PriceParts",
			SchemaName: "price_parts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ShippingOption) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode shippingOption#731bffce as nil")
	}
	b.PutID(ShippingOptionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ShippingOption) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode shippingOption#731bffce as nil")
	}
	b.PutString(s.ID)
	b.PutString(s.Title)
	b.PutInt(len(s.PriceParts))
	for idx, v := range s.PriceParts {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare shippingOption#731bffce: field price_parts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ShippingOption) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode shippingOption#731bffce to nil")
	}
	if err := b.ConsumeID(ShippingOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode shippingOption#731bffce: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ShippingOption) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode shippingOption#731bffce to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#731bffce: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#731bffce: field title: %w", err)
		}
		s.Title = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#731bffce: field price_parts: %w", err)
		}

		if headerLen > 0 {
			s.PriceParts = make([]LabeledPricePart, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LabeledPricePart
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare shippingOption#731bffce: field price_parts: %w", err)
			}
			s.PriceParts = append(s.PriceParts, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *ShippingOption) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode shippingOption#731bffce as nil")
	}
	b.ObjStart()
	b.PutID("shippingOption")
	b.FieldStart("id")
	b.PutString(s.ID)
	b.FieldStart("title")
	b.PutString(s.Title)
	b.FieldStart("price_parts")
	b.ArrStart()
	for idx, v := range s.PriceParts {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode shippingOption#731bffce: field price_parts element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *ShippingOption) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode shippingOption#731bffce to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("shippingOption"); err != nil {
				return fmt.Errorf("unable to decode shippingOption#731bffce: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode shippingOption#731bffce: field id: %w", err)
			}
			s.ID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode shippingOption#731bffce: field title: %w", err)
			}
			s.Title = value
		case "price_parts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value LabeledPricePart
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode shippingOption#731bffce: field price_parts: %w", err)
				}
				s.PriceParts = append(s.PriceParts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode shippingOption#731bffce: field price_parts: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *ShippingOption) GetID() (value string) {
	return s.ID
}

// GetTitle returns value of Title field.
func (s *ShippingOption) GetTitle() (value string) {
	return s.Title
}

// GetPriceParts returns value of PriceParts field.
func (s *ShippingOption) GetPriceParts() (value []LabeledPricePart) {
	return s.PriceParts
}
