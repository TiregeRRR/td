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

// MessageCopyOptions represents TL type `messageCopyOptions#48076039`.
type MessageCopyOptions struct {
	// True, if content of the message needs to be copied without a link to the original
	// message. Always true if the message is forwarded to a secret chat
	SendCopy bool
	// True, if media caption of the message copy needs to be replaced. Ignored if send_copy
	// is false
	ReplaceCaption bool
	// New message caption. Ignored if replace_caption is false
	NewCaption FormattedText
}

// MessageCopyOptionsTypeID is TL type id of MessageCopyOptions.
const MessageCopyOptionsTypeID = 0x48076039

// Ensuring interfaces in compile-time for MessageCopyOptions.
var (
	_ bin.Encoder     = &MessageCopyOptions{}
	_ bin.Decoder     = &MessageCopyOptions{}
	_ bin.BareEncoder = &MessageCopyOptions{}
	_ bin.BareDecoder = &MessageCopyOptions{}
)

func (m *MessageCopyOptions) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SendCopy == false) {
		return false
	}
	if !(m.ReplaceCaption == false) {
		return false
	}
	if !(m.NewCaption.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageCopyOptions) String() string {
	if m == nil {
		return "MessageCopyOptions(nil)"
	}
	type Alias MessageCopyOptions
	return fmt.Sprintf("MessageCopyOptions%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageCopyOptions) TypeID() uint32 {
	return MessageCopyOptionsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageCopyOptions) TypeName() string {
	return "messageCopyOptions"
}

// TypeInfo returns info about TL type.
func (m *MessageCopyOptions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageCopyOptions",
		ID:   MessageCopyOptionsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SendCopy",
			SchemaName: "send_copy",
		},
		{
			Name:       "ReplaceCaption",
			SchemaName: "replace_caption",
		},
		{
			Name:       "NewCaption",
			SchemaName: "new_caption",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageCopyOptions) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageCopyOptions#48076039 as nil")
	}
	b.PutID(MessageCopyOptionsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageCopyOptions) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageCopyOptions#48076039 as nil")
	}
	b.PutBool(m.SendCopy)
	b.PutBool(m.ReplaceCaption)
	if err := m.NewCaption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageCopyOptions#48076039: field new_caption: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageCopyOptions) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageCopyOptions#48076039 to nil")
	}
	if err := b.ConsumeID(MessageCopyOptionsTypeID); err != nil {
		return fmt.Errorf("unable to decode messageCopyOptions#48076039: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageCopyOptions) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageCopyOptions#48076039 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageCopyOptions#48076039: field send_copy: %w", err)
		}
		m.SendCopy = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageCopyOptions#48076039: field replace_caption: %w", err)
		}
		m.ReplaceCaption = value
	}
	{
		if err := m.NewCaption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageCopyOptions#48076039: field new_caption: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageCopyOptions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageCopyOptions#48076039 as nil")
	}
	b.ObjStart()
	b.PutID("messageCopyOptions")
	b.FieldStart("send_copy")
	b.PutBool(m.SendCopy)
	b.FieldStart("replace_caption")
	b.PutBool(m.ReplaceCaption)
	b.FieldStart("new_caption")
	if err := m.NewCaption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageCopyOptions#48076039: field new_caption: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageCopyOptions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageCopyOptions#48076039 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageCopyOptions"); err != nil {
				return fmt.Errorf("unable to decode messageCopyOptions#48076039: %w", err)
			}
		case "send_copy":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageCopyOptions#48076039: field send_copy: %w", err)
			}
			m.SendCopy = value
		case "replace_caption":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageCopyOptions#48076039: field replace_caption: %w", err)
			}
			m.ReplaceCaption = value
		case "new_caption":
			if err := m.NewCaption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageCopyOptions#48076039: field new_caption: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSendCopy returns value of SendCopy field.
func (m *MessageCopyOptions) GetSendCopy() (value bool) {
	return m.SendCopy
}

// GetReplaceCaption returns value of ReplaceCaption field.
func (m *MessageCopyOptions) GetReplaceCaption() (value bool) {
	return m.ReplaceCaption
}

// GetNewCaption returns value of NewCaption field.
func (m *MessageCopyOptions) GetNewCaption() (value FormattedText) {
	return m.NewCaption
}
