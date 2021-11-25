// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesEditChatPhotoRequest represents TL type `messages.editChatPhoto#35ddd674`.
// Changes chat photo and sends a service message on it
//
// See https://core.telegram.org/method/messages.editChatPhoto for reference.
type MessagesEditChatPhotoRequest struct {
	// Chat ID
	ChatID int64
	// Photo to be set
	Photo InputChatPhotoClass
}

// MessagesEditChatPhotoRequestTypeID is TL type id of MessagesEditChatPhotoRequest.
const MessagesEditChatPhotoRequestTypeID = 0x35ddd674

// Ensuring interfaces in compile-time for MessagesEditChatPhotoRequest.
var (
	_ bin.Encoder     = &MessagesEditChatPhotoRequest{}
	_ bin.Decoder     = &MessagesEditChatPhotoRequest{}
	_ bin.BareEncoder = &MessagesEditChatPhotoRequest{}
	_ bin.BareDecoder = &MessagesEditChatPhotoRequest{}
)

func (e *MessagesEditChatPhotoRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatPhotoRequest) String() string {
	if e == nil {
		return "MessagesEditChatPhotoRequest(nil)"
	}
	type Alias MessagesEditChatPhotoRequest
	return fmt.Sprintf("MessagesEditChatPhotoRequest%+v", Alias(*e))
}

// FillFrom fills MessagesEditChatPhotoRequest from given interface.
func (e *MessagesEditChatPhotoRequest) FillFrom(from interface {
	GetChatID() (value int64)
	GetPhoto() (value InputChatPhotoClass)
}) {
	e.ChatID = from.GetChatID()
	e.Photo = from.GetPhoto()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditChatPhotoRequest) TypeID() uint32 {
	return MessagesEditChatPhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditChatPhotoRequest) TypeName() string {
	return "messages.editChatPhoto"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditChatPhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editChatPhoto",
		ID:   MessagesEditChatPhotoRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatPhotoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatPhoto#35ddd674 as nil")
	}
	b.PutID(MessagesEditChatPhotoRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditChatPhotoRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatPhoto#35ddd674 as nil")
	}
	b.PutLong(e.ChatID)
	if e.Photo == nil {
		return fmt.Errorf("unable to encode messages.editChatPhoto#35ddd674: field photo is nil")
	}
	if err := e.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatPhoto#35ddd674: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatPhotoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatPhoto#35ddd674 to nil")
	}
	if err := b.ConsumeID(MessagesEditChatPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatPhoto#35ddd674: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditChatPhotoRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatPhoto#35ddd674 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatPhoto#35ddd674: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatPhoto#35ddd674: field photo: %w", err)
		}
		e.Photo = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (e *MessagesEditChatPhotoRequest) GetChatID() (value int64) {
	return e.ChatID
}

// GetPhoto returns value of Photo field.
func (e *MessagesEditChatPhotoRequest) GetPhoto() (value InputChatPhotoClass) {
	return e.Photo
}

// MessagesEditChatPhoto invokes method messages.editChatPhoto#35ddd674 returning error if any.
// Changes chat photo and sends a service message on it
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  400 INPUT_CONSTRUCTOR_INVALID: The provided constructor is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 PHOTO_CROP_SIZE_SMALL: Photo is too small
//  400 PHOTO_EXT_INVALID: The extension of the photo is invalid
//  400 PHOTO_INVALID: Photo invalid
//
// See https://core.telegram.org/method/messages.editChatPhoto for reference.
// Can be used by bots.
func (c *Client) MessagesEditChatPhoto(ctx context.Context, request *MessagesEditChatPhotoRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
