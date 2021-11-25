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

// Minithumbnail represents TL type `minithumbnail#ec6addaa`.
type Minithumbnail struct {
	// Thumbnail width, usually doesn't exceed 40
	Width int32
	// Thumbnail height, usually doesn't exceed 40
	Height int32
	// The thumbnail in JPEG format
	Data []byte
}

// MinithumbnailTypeID is TL type id of Minithumbnail.
const MinithumbnailTypeID = 0xec6addaa

// Ensuring interfaces in compile-time for Minithumbnail.
var (
	_ bin.Encoder     = &Minithumbnail{}
	_ bin.Decoder     = &Minithumbnail{}
	_ bin.BareEncoder = &Minithumbnail{}
	_ bin.BareDecoder = &Minithumbnail{}
)

func (m *Minithumbnail) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Width == 0) {
		return false
	}
	if !(m.Height == 0) {
		return false
	}
	if !(m.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *Minithumbnail) String() string {
	if m == nil {
		return "Minithumbnail(nil)"
	}
	type Alias Minithumbnail
	return fmt.Sprintf("Minithumbnail%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Minithumbnail) TypeID() uint32 {
	return MinithumbnailTypeID
}

// TypeName returns name of type in TL schema.
func (*Minithumbnail) TypeName() string {
	return "minithumbnail"
}

// TypeInfo returns info about TL type.
func (m *Minithumbnail) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "minithumbnail",
		ID:   MinithumbnailTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Width",
			SchemaName: "width",
		},
		{
			Name:       "Height",
			SchemaName: "height",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *Minithumbnail) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode minithumbnail#ec6addaa as nil")
	}
	b.PutID(MinithumbnailTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *Minithumbnail) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode minithumbnail#ec6addaa as nil")
	}
	b.PutInt32(m.Width)
	b.PutInt32(m.Height)
	b.PutBytes(m.Data)
	return nil
}

// Decode implements bin.Decoder.
func (m *Minithumbnail) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode minithumbnail#ec6addaa to nil")
	}
	if err := b.ConsumeID(MinithumbnailTypeID); err != nil {
		return fmt.Errorf("unable to decode minithumbnail#ec6addaa: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *Minithumbnail) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode minithumbnail#ec6addaa to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode minithumbnail#ec6addaa: field width: %w", err)
		}
		m.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode minithumbnail#ec6addaa: field height: %w", err)
		}
		m.Height = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode minithumbnail#ec6addaa: field data: %w", err)
		}
		m.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *Minithumbnail) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode minithumbnail#ec6addaa as nil")
	}
	b.ObjStart()
	b.PutID("minithumbnail")
	b.FieldStart("width")
	b.PutInt32(m.Width)
	b.FieldStart("height")
	b.PutInt32(m.Height)
	b.FieldStart("data")
	b.PutBytes(m.Data)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *Minithumbnail) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode minithumbnail#ec6addaa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("minithumbnail"); err != nil {
				return fmt.Errorf("unable to decode minithumbnail#ec6addaa: %w", err)
			}
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode minithumbnail#ec6addaa: field width: %w", err)
			}
			m.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode minithumbnail#ec6addaa: field height: %w", err)
			}
			m.Height = value
		case "data":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode minithumbnail#ec6addaa: field data: %w", err)
			}
			m.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetWidth returns value of Width field.
func (m *Minithumbnail) GetWidth() (value int32) {
	return m.Width
}

// GetHeight returns value of Height field.
func (m *Minithumbnail) GetHeight() (value int32) {
	return m.Height
}

// GetData returns value of Data field.
func (m *Minithumbnail) GetData() (value []byte) {
	return m.Data
}
