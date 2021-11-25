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

// DatedFile represents TL type `datedFile#9247b09d`.
type DatedFile struct {
	// The file
	File File
	// Point in time (Unix timestamp) when the file was uploaded
	Date int32
}

// DatedFileTypeID is TL type id of DatedFile.
const DatedFileTypeID = 0x9247b09d

// Ensuring interfaces in compile-time for DatedFile.
var (
	_ bin.Encoder     = &DatedFile{}
	_ bin.Decoder     = &DatedFile{}
	_ bin.BareEncoder = &DatedFile{}
	_ bin.BareDecoder = &DatedFile{}
)

func (d *DatedFile) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.File.Zero()) {
		return false
	}
	if !(d.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DatedFile) String() string {
	if d == nil {
		return "DatedFile(nil)"
	}
	type Alias DatedFile
	return fmt.Sprintf("DatedFile%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DatedFile) TypeID() uint32 {
	return DatedFileTypeID
}

// TypeName returns name of type in TL schema.
func (*DatedFile) TypeName() string {
	return "datedFile"
}

// TypeInfo returns info about TL type.
func (d *DatedFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "datedFile",
		ID:   DatedFileTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DatedFile) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode datedFile#9247b09d as nil")
	}
	b.PutID(DatedFileTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DatedFile) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode datedFile#9247b09d as nil")
	}
	if err := d.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode datedFile#9247b09d: field file: %w", err)
	}
	b.PutInt32(d.Date)
	return nil
}

// Decode implements bin.Decoder.
func (d *DatedFile) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode datedFile#9247b09d to nil")
	}
	if err := b.ConsumeID(DatedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode datedFile#9247b09d: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DatedFile) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode datedFile#9247b09d to nil")
	}
	{
		if err := d.File.Decode(b); err != nil {
			return fmt.Errorf("unable to decode datedFile#9247b09d: field file: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode datedFile#9247b09d: field date: %w", err)
		}
		d.Date = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DatedFile) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode datedFile#9247b09d as nil")
	}
	b.ObjStart()
	b.PutID("datedFile")
	b.FieldStart("file")
	if err := d.File.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode datedFile#9247b09d: field file: %w", err)
	}
	b.FieldStart("date")
	b.PutInt32(d.Date)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DatedFile) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode datedFile#9247b09d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("datedFile"); err != nil {
				return fmt.Errorf("unable to decode datedFile#9247b09d: %w", err)
			}
		case "file":
			if err := d.File.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode datedFile#9247b09d: field file: %w", err)
			}
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode datedFile#9247b09d: field date: %w", err)
			}
			d.Date = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFile returns value of File field.
func (d *DatedFile) GetFile() (value File) {
	return d.File
}

// GetDate returns value of Date field.
func (d *DatedFile) GetDate() (value int32) {
	return d.Date
}
