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

// PaymentFormTheme represents TL type `paymentFormTheme#97180f8f`.
type PaymentFormTheme struct {
	// A color of the payment form background in the RGB24 format
	BackgroundColor int32
	// A color of text in the RGB24 format
	TextColor int32
	// A color of hints in the RGB24 format
	HintColor int32
	// A color of links in the RGB24 format
	LinkColor int32
	// A color of thebuttons in the RGB24 format
	ButtonColor int32
	// A color of text on the buttons in the RGB24 format
	ButtonTextColor int32
}

// PaymentFormThemeTypeID is TL type id of PaymentFormTheme.
const PaymentFormThemeTypeID = 0x97180f8f

// Ensuring interfaces in compile-time for PaymentFormTheme.
var (
	_ bin.Encoder     = &PaymentFormTheme{}
	_ bin.Decoder     = &PaymentFormTheme{}
	_ bin.BareEncoder = &PaymentFormTheme{}
	_ bin.BareDecoder = &PaymentFormTheme{}
)

func (p *PaymentFormTheme) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.BackgroundColor == 0) {
		return false
	}
	if !(p.TextColor == 0) {
		return false
	}
	if !(p.HintColor == 0) {
		return false
	}
	if !(p.LinkColor == 0) {
		return false
	}
	if !(p.ButtonColor == 0) {
		return false
	}
	if !(p.ButtonTextColor == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentFormTheme) String() string {
	if p == nil {
		return "PaymentFormTheme(nil)"
	}
	type Alias PaymentFormTheme
	return fmt.Sprintf("PaymentFormTheme%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentFormTheme) TypeID() uint32 {
	return PaymentFormThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentFormTheme) TypeName() string {
	return "paymentFormTheme"
}

// TypeInfo returns info about TL type.
func (p *PaymentFormTheme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentFormTheme",
		ID:   PaymentFormThemeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BackgroundColor",
			SchemaName: "background_color",
		},
		{
			Name:       "TextColor",
			SchemaName: "text_color",
		},
		{
			Name:       "HintColor",
			SchemaName: "hint_color",
		},
		{
			Name:       "LinkColor",
			SchemaName: "link_color",
		},
		{
			Name:       "ButtonColor",
			SchemaName: "button_color",
		},
		{
			Name:       "ButtonTextColor",
			SchemaName: "button_text_color",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentFormTheme) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentFormTheme#97180f8f as nil")
	}
	b.PutID(PaymentFormThemeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentFormTheme) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentFormTheme#97180f8f as nil")
	}
	b.PutInt32(p.BackgroundColor)
	b.PutInt32(p.TextColor)
	b.PutInt32(p.HintColor)
	b.PutInt32(p.LinkColor)
	b.PutInt32(p.ButtonColor)
	b.PutInt32(p.ButtonTextColor)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentFormTheme) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentFormTheme#97180f8f to nil")
	}
	if err := b.ConsumeID(PaymentFormThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentFormTheme) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentFormTheme#97180f8f to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field background_color: %w", err)
		}
		p.BackgroundColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field text_color: %w", err)
		}
		p.TextColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field hint_color: %w", err)
		}
		p.HintColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field link_color: %w", err)
		}
		p.LinkColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field button_color: %w", err)
		}
		p.ButtonColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field button_text_color: %w", err)
		}
		p.ButtonTextColor = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentFormTheme) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentFormTheme#97180f8f as nil")
	}
	b.ObjStart()
	b.PutID("paymentFormTheme")
	b.FieldStart("background_color")
	b.PutInt32(p.BackgroundColor)
	b.FieldStart("text_color")
	b.PutInt32(p.TextColor)
	b.FieldStart("hint_color")
	b.PutInt32(p.HintColor)
	b.FieldStart("link_color")
	b.PutInt32(p.LinkColor)
	b.FieldStart("button_color")
	b.PutInt32(p.ButtonColor)
	b.FieldStart("button_text_color")
	b.PutInt32(p.ButtonTextColor)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentFormTheme) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentFormTheme#97180f8f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentFormTheme"); err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: %w", err)
			}
		case "background_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field background_color: %w", err)
			}
			p.BackgroundColor = value
		case "text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field text_color: %w", err)
			}
			p.TextColor = value
		case "hint_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field hint_color: %w", err)
			}
			p.HintColor = value
		case "link_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field link_color: %w", err)
			}
			p.LinkColor = value
		case "button_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field button_color: %w", err)
			}
			p.ButtonColor = value
		case "button_text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentFormTheme#97180f8f: field button_text_color: %w", err)
			}
			p.ButtonTextColor = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBackgroundColor returns value of BackgroundColor field.
func (p *PaymentFormTheme) GetBackgroundColor() (value int32) {
	return p.BackgroundColor
}

// GetTextColor returns value of TextColor field.
func (p *PaymentFormTheme) GetTextColor() (value int32) {
	return p.TextColor
}

// GetHintColor returns value of HintColor field.
func (p *PaymentFormTheme) GetHintColor() (value int32) {
	return p.HintColor
}

// GetLinkColor returns value of LinkColor field.
func (p *PaymentFormTheme) GetLinkColor() (value int32) {
	return p.LinkColor
}

// GetButtonColor returns value of ButtonColor field.
func (p *PaymentFormTheme) GetButtonColor() (value int32) {
	return p.ButtonColor
}

// GetButtonTextColor returns value of ButtonTextColor field.
func (p *PaymentFormTheme) GetButtonTextColor() (value int32) {
	return p.ButtonTextColor
}
