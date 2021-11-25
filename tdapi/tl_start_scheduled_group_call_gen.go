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

// StartScheduledGroupCallRequest represents TL type `startScheduledGroupCall#5a986d16`.
type StartScheduledGroupCallRequest struct {
	// Group call identifier
	GroupCallID int32
}

// StartScheduledGroupCallRequestTypeID is TL type id of StartScheduledGroupCallRequest.
const StartScheduledGroupCallRequestTypeID = 0x5a986d16

// Ensuring interfaces in compile-time for StartScheduledGroupCallRequest.
var (
	_ bin.Encoder     = &StartScheduledGroupCallRequest{}
	_ bin.Decoder     = &StartScheduledGroupCallRequest{}
	_ bin.BareEncoder = &StartScheduledGroupCallRequest{}
	_ bin.BareDecoder = &StartScheduledGroupCallRequest{}
)

func (s *StartScheduledGroupCallRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.GroupCallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StartScheduledGroupCallRequest) String() string {
	if s == nil {
		return "StartScheduledGroupCallRequest(nil)"
	}
	type Alias StartScheduledGroupCallRequest
	return fmt.Sprintf("StartScheduledGroupCallRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StartScheduledGroupCallRequest) TypeID() uint32 {
	return StartScheduledGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StartScheduledGroupCallRequest) TypeName() string {
	return "startScheduledGroupCall"
}

// TypeInfo returns info about TL type.
func (s *StartScheduledGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "startScheduledGroupCall",
		ID:   StartScheduledGroupCallRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StartScheduledGroupCallRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode startScheduledGroupCall#5a986d16 as nil")
	}
	b.PutID(StartScheduledGroupCallRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StartScheduledGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode startScheduledGroupCall#5a986d16 as nil")
	}
	b.PutInt32(s.GroupCallID)
	return nil
}

// Decode implements bin.Decoder.
func (s *StartScheduledGroupCallRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode startScheduledGroupCall#5a986d16 to nil")
	}
	if err := b.ConsumeID(StartScheduledGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode startScheduledGroupCall#5a986d16: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StartScheduledGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode startScheduledGroupCall#5a986d16 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode startScheduledGroupCall#5a986d16: field group_call_id: %w", err)
		}
		s.GroupCallID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StartScheduledGroupCallRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode startScheduledGroupCall#5a986d16 as nil")
	}
	b.ObjStart()
	b.PutID("startScheduledGroupCall")
	b.FieldStart("group_call_id")
	b.PutInt32(s.GroupCallID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StartScheduledGroupCallRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode startScheduledGroupCall#5a986d16 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("startScheduledGroupCall"); err != nil {
				return fmt.Errorf("unable to decode startScheduledGroupCall#5a986d16: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode startScheduledGroupCall#5a986d16: field group_call_id: %w", err)
			}
			s.GroupCallID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (s *StartScheduledGroupCallRequest) GetGroupCallID() (value int32) {
	return s.GroupCallID
}

// StartScheduledGroupCall invokes method startScheduledGroupCall#5a986d16 returning error if any.
func (c *Client) StartScheduledGroupCall(ctx context.Context, groupcallid int32) error {
	var ok Ok

	request := &StartScheduledGroupCallRequest{
		GroupCallID: groupcallid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
