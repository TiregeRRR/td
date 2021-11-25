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

// ForwardMessagesRequest represents TL type `forwardMessages#1842041f`.
type ForwardMessagesRequest struct {
	// Identifier of the chat to which to forward messages
	ChatID int64
	// Identifier of the chat from which to forward messages
	FromChatID int64
	// Identifiers of the messages to forward. Message identifiers must be in a strictly
	// increasing order. At most 100 messages can be forwarded simultaneously
	MessageIDs []int64
	// Options to be used to send the messages
	Options MessageSendOptions
	// True, if content of the messages needs to be copied without links to the original
	// messages. Always true if the messages are forwarded to a secret chat
	SendCopy bool
	// True, if media caption of message copies needs to be removed. Ignored if send_copy is
	// false
	RemoveCaption bool
}

// ForwardMessagesRequestTypeID is TL type id of ForwardMessagesRequest.
const ForwardMessagesRequestTypeID = 0x1842041f

// Ensuring interfaces in compile-time for ForwardMessagesRequest.
var (
	_ bin.Encoder     = &ForwardMessagesRequest{}
	_ bin.Decoder     = &ForwardMessagesRequest{}
	_ bin.BareEncoder = &ForwardMessagesRequest{}
	_ bin.BareDecoder = &ForwardMessagesRequest{}
)

func (f *ForwardMessagesRequest) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.ChatID == 0) {
		return false
	}
	if !(f.FromChatID == 0) {
		return false
	}
	if !(f.MessageIDs == nil) {
		return false
	}
	if !(f.Options.Zero()) {
		return false
	}
	if !(f.SendCopy == false) {
		return false
	}
	if !(f.RemoveCaption == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *ForwardMessagesRequest) String() string {
	if f == nil {
		return "ForwardMessagesRequest(nil)"
	}
	type Alias ForwardMessagesRequest
	return fmt.Sprintf("ForwardMessagesRequest%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ForwardMessagesRequest) TypeID() uint32 {
	return ForwardMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ForwardMessagesRequest) TypeName() string {
	return "forwardMessages"
}

// TypeInfo returns info about TL type.
func (f *ForwardMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "forwardMessages",
		ID:   ForwardMessagesRequestTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "FromChatID",
			SchemaName: "from_chat_id",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "SendCopy",
			SchemaName: "send_copy",
		},
		{
			Name:       "RemoveCaption",
			SchemaName: "remove_caption",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *ForwardMessagesRequest) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode forwardMessages#1842041f as nil")
	}
	b.PutID(ForwardMessagesRequestTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *ForwardMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode forwardMessages#1842041f as nil")
	}
	b.PutLong(f.ChatID)
	b.PutLong(f.FromChatID)
	b.PutInt(len(f.MessageIDs))
	for _, v := range f.MessageIDs {
		b.PutLong(v)
	}
	if err := f.Options.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forwardMessages#1842041f: field options: %w", err)
	}
	b.PutBool(f.SendCopy)
	b.PutBool(f.RemoveCaption)
	return nil
}

// Decode implements bin.Decoder.
func (f *ForwardMessagesRequest) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode forwardMessages#1842041f to nil")
	}
	if err := b.ConsumeID(ForwardMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode forwardMessages#1842041f: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *ForwardMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode forwardMessages#1842041f to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode forwardMessages#1842041f: field chat_id: %w", err)
		}
		f.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode forwardMessages#1842041f: field from_chat_id: %w", err)
		}
		f.FromChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode forwardMessages#1842041f: field message_ids: %w", err)
		}

		if headerLen > 0 {
			f.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field message_ids: %w", err)
			}
			f.MessageIDs = append(f.MessageIDs, value)
		}
	}
	{
		if err := f.Options.Decode(b); err != nil {
			return fmt.Errorf("unable to decode forwardMessages#1842041f: field options: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forwardMessages#1842041f: field send_copy: %w", err)
		}
		f.SendCopy = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forwardMessages#1842041f: field remove_caption: %w", err)
		}
		f.RemoveCaption = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *ForwardMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode forwardMessages#1842041f as nil")
	}
	b.ObjStart()
	b.PutID("forwardMessages")
	b.FieldStart("chat_id")
	b.PutLong(f.ChatID)
	b.FieldStart("from_chat_id")
	b.PutLong(f.FromChatID)
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range f.MessageIDs {
		b.PutLong(v)
	}
	b.ArrEnd()
	b.FieldStart("options")
	if err := f.Options.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forwardMessages#1842041f: field options: %w", err)
	}
	b.FieldStart("send_copy")
	b.PutBool(f.SendCopy)
	b.FieldStart("remove_caption")
	b.PutBool(f.RemoveCaption)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *ForwardMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode forwardMessages#1842041f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("forwardMessages"); err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field chat_id: %w", err)
			}
			f.ChatID = value
		case "from_chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field from_chat_id: %w", err)
			}
			f.FromChatID = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Long()
				if err != nil {
					return fmt.Errorf("unable to decode forwardMessages#1842041f: field message_ids: %w", err)
				}
				f.MessageIDs = append(f.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field message_ids: %w", err)
			}
		case "options":
			if err := f.Options.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field options: %w", err)
			}
		case "send_copy":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field send_copy: %w", err)
			}
			f.SendCopy = value
		case "remove_caption":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forwardMessages#1842041f: field remove_caption: %w", err)
			}
			f.RemoveCaption = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (f *ForwardMessagesRequest) GetChatID() (value int64) {
	return f.ChatID
}

// GetFromChatID returns value of FromChatID field.
func (f *ForwardMessagesRequest) GetFromChatID() (value int64) {
	return f.FromChatID
}

// GetMessageIDs returns value of MessageIDs field.
func (f *ForwardMessagesRequest) GetMessageIDs() (value []int64) {
	return f.MessageIDs
}

// GetOptions returns value of Options field.
func (f *ForwardMessagesRequest) GetOptions() (value MessageSendOptions) {
	return f.Options
}

// GetSendCopy returns value of SendCopy field.
func (f *ForwardMessagesRequest) GetSendCopy() (value bool) {
	return f.SendCopy
}

// GetRemoveCaption returns value of RemoveCaption field.
func (f *ForwardMessagesRequest) GetRemoveCaption() (value bool) {
	return f.RemoveCaption
}

// ForwardMessages invokes method forwardMessages#1842041f returning error if any.
func (c *Client) ForwardMessages(ctx context.Context, request *ForwardMessagesRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
