// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// Double represents TL type `double#2210c154`.
//
// See https://core.telegram.org/constructor/double for reference.
type Double struct {
}

// DoubleTypeID is TL type id of Double.
const DoubleTypeID = 0x2210c154

// Ensuring interfaces in compile-time for Double.
var (
	_ bin.Encoder     = &Double{}
	_ bin.Decoder     = &Double{}
	_ bin.BareEncoder = &Double{}
	_ bin.BareDecoder = &Double{}
)

func (d *Double) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *Double) String() string {
	if d == nil {
		return "Double(nil)"
	}
	type Alias Double
	return fmt.Sprintf("Double%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Double) TypeID() uint32 {
	return DoubleTypeID
}

// TypeName returns name of type in TL schema.
func (*Double) TypeName() string {
	return "double"
}

// TypeInfo returns info about TL type.
func (d *Double) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "double",
		ID:   DoubleTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *Double) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode double#2210c154 as nil")
	}
	b.PutID(DoubleTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *Double) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode double#2210c154 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *Double) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode double#2210c154 to nil")
	}
	if err := b.ConsumeID(DoubleTypeID); err != nil {
		return fmt.Errorf("unable to decode double#2210c154: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *Double) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode double#2210c154 to nil")
	}
	return nil
}
