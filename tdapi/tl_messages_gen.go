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

// Messages represents TL type `messages#b34c8c60`.
type Messages struct {
	// Approximate total count of messages found
	TotalCount int32
	// List of messages; messages may be null
	Messages []Message
}

// MessagesTypeID is TL type id of Messages.
const MessagesTypeID = 0xb34c8c60

// Ensuring interfaces in compile-time for Messages.
var (
	_ bin.Encoder     = &Messages{}
	_ bin.Decoder     = &Messages{}
	_ bin.BareEncoder = &Messages{}
	_ bin.BareDecoder = &Messages{}
)

func (m *Messages) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.TotalCount == 0) {
		return false
	}
	if !(m.Messages == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *Messages) String() string {
	if m == nil {
		return "Messages(nil)"
	}
	type Alias Messages
	return fmt.Sprintf("Messages%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Messages) TypeID() uint32 {
	return MessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*Messages) TypeName() string {
	return "messages"
}

// TypeInfo returns info about TL type.
func (m *Messages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages",
		ID:   MessagesTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *Messages) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages#b34c8c60 as nil")
	}
	b.PutID(MessagesTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *Messages) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages#b34c8c60 as nil")
	}
	b.PutInt32(m.TotalCount)
	b.PutInt(len(m.Messages))
	for idx, v := range m.Messages {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare messages#b34c8c60: field messages element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *Messages) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages#b34c8c60 to nil")
	}
	if err := b.ConsumeID(MessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages#b34c8c60: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *Messages) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages#b34c8c60 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messages#b34c8c60: field total_count: %w", err)
		}
		m.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages#b34c8c60: field messages: %w", err)
		}

		if headerLen > 0 {
			m.Messages = make([]Message, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Message
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare messages#b34c8c60: field messages: %w", err)
			}
			m.Messages = append(m.Messages, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *Messages) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messages#b34c8c60 as nil")
	}
	b.ObjStart()
	b.PutID("messages")
	b.FieldStart("total_count")
	b.PutInt32(m.TotalCount)
	b.FieldStart("messages")
	b.ArrStart()
	for idx, v := range m.Messages {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode messages#b34c8c60: field messages element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *Messages) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messages#b34c8c60 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messages"); err != nil {
				return fmt.Errorf("unable to decode messages#b34c8c60: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messages#b34c8c60: field total_count: %w", err)
			}
			m.TotalCount = value
		case "messages":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Message
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode messages#b34c8c60: field messages: %w", err)
				}
				m.Messages = append(m.Messages, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode messages#b34c8c60: field messages: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (m *Messages) GetTotalCount() (value int32) {
	return m.TotalCount
}

// GetMessages returns value of Messages field.
func (m *Messages) GetMessages() (value []Message) {
	return m.Messages
}
