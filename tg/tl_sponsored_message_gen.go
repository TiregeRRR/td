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

// SponsoredMessage represents TL type `sponsoredMessage#d151e19a`.
//
// See https://core.telegram.org/constructor/sponsoredMessage for reference.
type SponsoredMessage struct {
	// Flags field of SponsoredMessage.
	Flags bin.Fields
	// RandomID field of SponsoredMessage.
	RandomID []byte
	// FromID field of SponsoredMessage.
	FromID PeerClass
	// ChannelPost field of SponsoredMessage.
	//
	// Use SetChannelPost and GetChannelPost helpers.
	ChannelPost int
	// StartParam field of SponsoredMessage.
	//
	// Use SetStartParam and GetStartParam helpers.
	StartParam string
	// Message field of SponsoredMessage.
	Message string
	// Entities field of SponsoredMessage.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// SponsoredMessageTypeID is TL type id of SponsoredMessage.
const SponsoredMessageTypeID = 0xd151e19a

// Ensuring interfaces in compile-time for SponsoredMessage.
var (
	_ bin.Encoder     = &SponsoredMessage{}
	_ bin.Decoder     = &SponsoredMessage{}
	_ bin.BareEncoder = &SponsoredMessage{}
	_ bin.BareDecoder = &SponsoredMessage{}
)

func (s *SponsoredMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.RandomID == nil) {
		return false
	}
	if !(s.FromID == nil) {
		return false
	}
	if !(s.ChannelPost == 0) {
		return false
	}
	if !(s.StartParam == "") {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredMessage) String() string {
	if s == nil {
		return "SponsoredMessage(nil)"
	}
	type Alias SponsoredMessage
	return fmt.Sprintf("SponsoredMessage%+v", Alias(*s))
}

// FillFrom fills SponsoredMessage from given interface.
func (s *SponsoredMessage) FillFrom(from interface {
	GetRandomID() (value []byte)
	GetFromID() (value PeerClass)
	GetChannelPost() (value int, ok bool)
	GetStartParam() (value string, ok bool)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
}) {
	s.RandomID = from.GetRandomID()
	s.FromID = from.GetFromID()
	if val, ok := from.GetChannelPost(); ok {
		s.ChannelPost = val
	}

	if val, ok := from.GetStartParam(); ok {
		s.StartParam = val
	}

	s.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		s.Entities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredMessage) TypeID() uint32 {
	return SponsoredMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredMessage) TypeName() string {
	return "sponsoredMessage"
}

// TypeInfo returns info about TL type.
func (s *SponsoredMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredMessage",
		ID:   SponsoredMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "FromID",
			SchemaName: "from_id",
		},
		{
			Name:       "ChannelPost",
			SchemaName: "channel_post",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SponsoredMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#d151e19a as nil")
	}
	b.PutID(SponsoredMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#d151e19a as nil")
	}
	if !(s.ChannelPost == 0) {
		s.Flags.Set(2)
	}
	if !(s.StartParam == "") {
		s.Flags.Set(0)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#d151e19a: field flags: %w", err)
	}
	b.PutBytes(s.RandomID)
	if s.FromID == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#d151e19a: field from_id is nil")
	}
	if err := s.FromID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#d151e19a: field from_id: %w", err)
	}
	if s.Flags.Has(2) {
		b.PutInt(s.ChannelPost)
	}
	if s.Flags.Has(0) {
		b.PutString(s.StartParam)
	}
	b.PutString(s.Message)
	if s.Flags.Has(1) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode sponsoredMessage#d151e19a: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode sponsoredMessage#d151e19a: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#d151e19a to nil")
	}
	if err := b.ConsumeID(SponsoredMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#d151e19a to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field flags: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field from_id: %w", err)
		}
		s.FromID = value
	}
	if s.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field channel_post: %w", err)
		}
		s.ChannelPost = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field start_param: %w", err)
		}
		s.StartParam = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#d151e19a: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	return nil
}

// GetRandomID returns value of RandomID field.
func (s *SponsoredMessage) GetRandomID() (value []byte) {
	return s.RandomID
}

// GetFromID returns value of FromID field.
func (s *SponsoredMessage) GetFromID() (value PeerClass) {
	return s.FromID
}

// SetChannelPost sets value of ChannelPost conditional field.
func (s *SponsoredMessage) SetChannelPost(value int) {
	s.Flags.Set(2)
	s.ChannelPost = value
}

// GetChannelPost returns value of ChannelPost conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetChannelPost() (value int, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ChannelPost, true
}

// SetStartParam sets value of StartParam conditional field.
func (s *SponsoredMessage) SetStartParam(value string) {
	s.Flags.Set(0)
	s.StartParam = value
}

// GetStartParam returns value of StartParam conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetStartParam() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.StartParam, true
}

// GetMessage returns value of Message field.
func (s *SponsoredMessage) GetMessage() (value string) {
	return s.Message
}

// SetEntities sets value of Entities conditional field.
func (s *SponsoredMessage) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(1)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetEntities() (value []MessageEntityClass, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Entities, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (s *SponsoredMessage) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return MessageEntityClassArray(s.Entities), true
}
