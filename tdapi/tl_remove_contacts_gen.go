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

// RemoveContactsRequest represents TL type `removeContacts#d47cddcc`.
type RemoveContactsRequest struct {
	// Identifiers of users to be deleted
	UserIDs []int32
}

// RemoveContactsRequestTypeID is TL type id of RemoveContactsRequest.
const RemoveContactsRequestTypeID = 0xd47cddcc

// Ensuring interfaces in compile-time for RemoveContactsRequest.
var (
	_ bin.Encoder     = &RemoveContactsRequest{}
	_ bin.Decoder     = &RemoveContactsRequest{}
	_ bin.BareEncoder = &RemoveContactsRequest{}
	_ bin.BareDecoder = &RemoveContactsRequest{}
)

func (r *RemoveContactsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveContactsRequest) String() string {
	if r == nil {
		return "RemoveContactsRequest(nil)"
	}
	type Alias RemoveContactsRequest
	return fmt.Sprintf("RemoveContactsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveContactsRequest) TypeID() uint32 {
	return RemoveContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveContactsRequest) TypeName() string {
	return "removeContacts"
}

// TypeInfo returns info about TL type.
func (r *RemoveContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeContacts",
		ID:   RemoveContactsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveContactsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeContacts#d47cddcc as nil")
	}
	b.PutID(RemoveContactsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveContactsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeContacts#d47cddcc as nil")
	}
	b.PutInt(len(r.UserIDs))
	for _, v := range r.UserIDs {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveContactsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeContacts#d47cddcc to nil")
	}
	if err := b.ConsumeID(RemoveContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeContacts#d47cddcc: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveContactsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeContacts#d47cddcc to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode removeContacts#d47cddcc: field user_ids: %w", err)
		}

		if headerLen > 0 {
			r.UserIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode removeContacts#d47cddcc: field user_ids: %w", err)
			}
			r.UserIDs = append(r.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveContactsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeContacts#d47cddcc as nil")
	}
	b.ObjStart()
	b.PutID("removeContacts")
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range r.UserIDs {
		b.PutInt32(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveContactsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeContacts#d47cddcc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeContacts"); err != nil {
				return fmt.Errorf("unable to decode removeContacts#d47cddcc: %w", err)
			}
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode removeContacts#d47cddcc: field user_ids: %w", err)
				}
				r.UserIDs = append(r.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode removeContacts#d47cddcc: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserIDs returns value of UserIDs field.
func (r *RemoveContactsRequest) GetUserIDs() (value []int32) {
	return r.UserIDs
}

// RemoveContacts invokes method removeContacts#d47cddcc returning error if any.
func (c *Client) RemoveContacts(ctx context.Context, userids []int32) error {
	var ok Ok

	request := &RemoveContactsRequest{
		UserIDs: userids,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
