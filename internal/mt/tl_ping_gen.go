// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// PingRequest represents TL type `ping#7abe77ec`.
type PingRequest struct {
	// PingID field of PingRequest.
	PingID int64
}

// PingRequestTypeID is TL type id of PingRequest.
const PingRequestTypeID = 0x7abe77ec

// Ensuring interfaces in compile-time for PingRequest.
var (
	_ bin.Encoder     = &PingRequest{}
	_ bin.Decoder     = &PingRequest{}
	_ bin.BareEncoder = &PingRequest{}
	_ bin.BareDecoder = &PingRequest{}
)

func (p *PingRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PingID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PingRequest) String() string {
	if p == nil {
		return "PingRequest(nil)"
	}
	type Alias PingRequest
	return fmt.Sprintf("PingRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PingRequest) TypeID() uint32 {
	return PingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PingRequest) TypeName() string {
	return "ping"
}

// TypeInfo returns info about TL type.
func (p *PingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ping",
		ID:   PingRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PingID",
			SchemaName: "ping_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PingRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping#7abe77ec as nil")
	}
	b.PutID(PingRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PingRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping#7abe77ec as nil")
	}
	b.PutLong(p.PingID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PingRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping#7abe77ec to nil")
	}
	if err := b.ConsumeID(PingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode ping#7abe77ec: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PingRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping#7abe77ec to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode ping#7abe77ec: field ping_id: %w", err)
		}
		p.PingID = value
	}
	return nil
}

// GetPingID returns value of PingID field.
func (p *PingRequest) GetPingID() (value int64) {
	return p.PingID
}
