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

// UpdatesDifferenceEmpty represents TL type `updates.differenceEmpty#5d75a138`.
// No events.
//
// See https://core.telegram.org/constructor/updates.differenceEmpty for reference.
type UpdatesDifferenceEmpty struct {
	// Current date
	Date int
	// Number of sent updates
	Seq int
}

// UpdatesDifferenceEmptyTypeID is TL type id of UpdatesDifferenceEmpty.
const UpdatesDifferenceEmptyTypeID = 0x5d75a138

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifferenceEmpty) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifferenceEmpty.
var (
	_ bin.Encoder     = &UpdatesDifferenceEmpty{}
	_ bin.Decoder     = &UpdatesDifferenceEmpty{}
	_ bin.BareEncoder = &UpdatesDifferenceEmpty{}
	_ bin.BareDecoder = &UpdatesDifferenceEmpty{}

	_ UpdatesDifferenceClass = &UpdatesDifferenceEmpty{}
)

func (d *UpdatesDifferenceEmpty) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Date == 0) {
		return false
	}
	if !(d.Seq == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifferenceEmpty) String() string {
	if d == nil {
		return "UpdatesDifferenceEmpty(nil)"
	}
	type Alias UpdatesDifferenceEmpty
	return fmt.Sprintf("UpdatesDifferenceEmpty%+v", Alias(*d))
}

// FillFrom fills UpdatesDifferenceEmpty from given interface.
func (d *UpdatesDifferenceEmpty) FillFrom(from interface {
	GetDate() (value int)
	GetSeq() (value int)
}) {
	d.Date = from.GetDate()
	d.Seq = from.GetSeq()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesDifferenceEmpty) TypeID() uint32 {
	return UpdatesDifferenceEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesDifferenceEmpty) TypeName() string {
	return "updates.differenceEmpty"
}

// TypeInfo returns info about TL type.
func (d *UpdatesDifferenceEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.differenceEmpty",
		ID:   UpdatesDifferenceEmptyTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Seq",
			SchemaName: "seq",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *UpdatesDifferenceEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceEmpty#5d75a138 as nil")
	}
	b.PutID(UpdatesDifferenceEmptyTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *UpdatesDifferenceEmpty) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceEmpty#5d75a138 as nil")
	}
	b.PutInt(d.Date)
	b.PutInt(d.Seq)
	return nil
}

// Decode implements bin.Decoder.
func (d *UpdatesDifferenceEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceEmpty#5d75a138 to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.differenceEmpty#5d75a138: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *UpdatesDifferenceEmpty) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceEmpty#5d75a138 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceEmpty#5d75a138: field date: %w", err)
		}
		d.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceEmpty#5d75a138: field seq: %w", err)
		}
		d.Seq = value
	}
	return nil
}

// GetDate returns value of Date field.
func (d *UpdatesDifferenceEmpty) GetDate() (value int) {
	return d.Date
}

// GetSeq returns value of Seq field.
func (d *UpdatesDifferenceEmpty) GetSeq() (value int) {
	return d.Seq
}

// UpdatesDifference represents TL type `updates.difference#f49ca0`.
// Full list of occurred events.
//
// See https://core.telegram.org/constructor/updates.difference for reference.
type UpdatesDifference struct {
	// List of new messages
	NewMessages []MessageClass
	// List of new encrypted secret chat messages
	NewEncryptedMessages []EncryptedMessageClass
	// List of updates
	OtherUpdates []UpdateClass
	// List of chats mentioned in events
	Chats []ChatClass
	// List of users mentioned in events
	Users []UserClass
	// Current state
	State UpdatesState
}

// UpdatesDifferenceTypeID is TL type id of UpdatesDifference.
const UpdatesDifferenceTypeID = 0xf49ca0

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifference) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifference.
var (
	_ bin.Encoder     = &UpdatesDifference{}
	_ bin.Decoder     = &UpdatesDifference{}
	_ bin.BareEncoder = &UpdatesDifference{}
	_ bin.BareDecoder = &UpdatesDifference{}

	_ UpdatesDifferenceClass = &UpdatesDifference{}
)

func (d *UpdatesDifference) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.NewMessages == nil) {
		return false
	}
	if !(d.NewEncryptedMessages == nil) {
		return false
	}
	if !(d.OtherUpdates == nil) {
		return false
	}
	if !(d.Chats == nil) {
		return false
	}
	if !(d.Users == nil) {
		return false
	}
	if !(d.State.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifference) String() string {
	if d == nil {
		return "UpdatesDifference(nil)"
	}
	type Alias UpdatesDifference
	return fmt.Sprintf("UpdatesDifference%+v", Alias(*d))
}

// FillFrom fills UpdatesDifference from given interface.
func (d *UpdatesDifference) FillFrom(from interface {
	GetNewMessages() (value []MessageClass)
	GetNewEncryptedMessages() (value []EncryptedMessageClass)
	GetOtherUpdates() (value []UpdateClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetState() (value UpdatesState)
}) {
	d.NewMessages = from.GetNewMessages()
	d.NewEncryptedMessages = from.GetNewEncryptedMessages()
	d.OtherUpdates = from.GetOtherUpdates()
	d.Chats = from.GetChats()
	d.Users = from.GetUsers()
	d.State = from.GetState()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesDifference) TypeID() uint32 {
	return UpdatesDifferenceTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesDifference) TypeName() string {
	return "updates.difference"
}

// TypeInfo returns info about TL type.
func (d *UpdatesDifference) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.difference",
		ID:   UpdatesDifferenceTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NewMessages",
			SchemaName: "new_messages",
		},
		{
			Name:       "NewEncryptedMessages",
			SchemaName: "new_encrypted_messages",
		},
		{
			Name:       "OtherUpdates",
			SchemaName: "other_updates",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *UpdatesDifference) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.difference#f49ca0 as nil")
	}
	b.PutID(UpdatesDifferenceTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *UpdatesDifference) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.difference#f49ca0 as nil")
	}
	b.PutVectorHeader(len(d.NewMessages))
	for idx, v := range d.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.NewEncryptedMessages))
	for idx, v := range d.NewEncryptedMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_encrypted_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_encrypted_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.OtherUpdates))
	for idx, v := range d.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field users element with index %d: %w", idx, err)
		}
	}
	if err := d.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.difference#f49ca0: field state: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *UpdatesDifference) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.difference#f49ca0 to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.difference#f49ca0: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *UpdatesDifference) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.difference#f49ca0 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_messages: %w", err)
		}

		if headerLen > 0 {
			d.NewMessages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_messages: %w", err)
			}
			d.NewMessages = append(d.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_encrypted_messages: %w", err)
		}

		if headerLen > 0 {
			d.NewEncryptedMessages = make([]EncryptedMessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEncryptedMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_encrypted_messages: %w", err)
			}
			d.NewEncryptedMessages = append(d.NewEncryptedMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field other_updates: %w", err)
		}

		if headerLen > 0 {
			d.OtherUpdates = make([]UpdateClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field other_updates: %w", err)
			}
			d.OtherUpdates = append(d.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field chats: %w", err)
		}

		if headerLen > 0 {
			d.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field users: %w", err)
		}

		if headerLen > 0 {
			d.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	{
		if err := d.State.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field state: %w", err)
		}
	}
	return nil
}

// GetNewMessages returns value of NewMessages field.
func (d *UpdatesDifference) GetNewMessages() (value []MessageClass) {
	return d.NewMessages
}

// GetNewEncryptedMessages returns value of NewEncryptedMessages field.
func (d *UpdatesDifference) GetNewEncryptedMessages() (value []EncryptedMessageClass) {
	return d.NewEncryptedMessages
}

// GetOtherUpdates returns value of OtherUpdates field.
func (d *UpdatesDifference) GetOtherUpdates() (value []UpdateClass) {
	return d.OtherUpdates
}

// GetChats returns value of Chats field.
func (d *UpdatesDifference) GetChats() (value []ChatClass) {
	return d.Chats
}

// GetUsers returns value of Users field.
func (d *UpdatesDifference) GetUsers() (value []UserClass) {
	return d.Users
}

// GetState returns value of State field.
func (d *UpdatesDifference) GetState() (value UpdatesState) {
	return d.State
}

// MapNewMessages returns field NewMessages wrapped in MessageClassArray helper.
func (d *UpdatesDifference) MapNewMessages() (value MessageClassArray) {
	return MessageClassArray(d.NewMessages)
}

// MapNewEncryptedMessages returns field NewEncryptedMessages wrapped in EncryptedMessageClassArray helper.
func (d *UpdatesDifference) MapNewEncryptedMessages() (value EncryptedMessageClassArray) {
	return EncryptedMessageClassArray(d.NewEncryptedMessages)
}

// MapOtherUpdates returns field OtherUpdates wrapped in UpdateClassArray helper.
func (d *UpdatesDifference) MapOtherUpdates() (value UpdateClassArray) {
	return UpdateClassArray(d.OtherUpdates)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (d *UpdatesDifference) MapChats() (value ChatClassArray) {
	return ChatClassArray(d.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (d *UpdatesDifference) MapUsers() (value UserClassArray) {
	return UserClassArray(d.Users)
}

// UpdatesDifferenceSlice represents TL type `updates.differenceSlice#a8fb1981`.
// Incomplete list of occurred events.
//
// See https://core.telegram.org/constructor/updates.differenceSlice for reference.
type UpdatesDifferenceSlice struct {
	// List of new messgaes
	NewMessages []MessageClass
	// New messages from the encrypted event sequence¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	NewEncryptedMessages []EncryptedMessageClass
	// List of updates
	OtherUpdates []UpdateClass
	// List of chats mentioned in events
	Chats []ChatClass
	// List of users mentioned in events
	Users []UserClass
	// Intermediary state
	IntermediateState UpdatesState
}

// UpdatesDifferenceSliceTypeID is TL type id of UpdatesDifferenceSlice.
const UpdatesDifferenceSliceTypeID = 0xa8fb1981

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifferenceSlice) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifferenceSlice.
var (
	_ bin.Encoder     = &UpdatesDifferenceSlice{}
	_ bin.Decoder     = &UpdatesDifferenceSlice{}
	_ bin.BareEncoder = &UpdatesDifferenceSlice{}
	_ bin.BareDecoder = &UpdatesDifferenceSlice{}

	_ UpdatesDifferenceClass = &UpdatesDifferenceSlice{}
)

func (d *UpdatesDifferenceSlice) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.NewMessages == nil) {
		return false
	}
	if !(d.NewEncryptedMessages == nil) {
		return false
	}
	if !(d.OtherUpdates == nil) {
		return false
	}
	if !(d.Chats == nil) {
		return false
	}
	if !(d.Users == nil) {
		return false
	}
	if !(d.IntermediateState.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifferenceSlice) String() string {
	if d == nil {
		return "UpdatesDifferenceSlice(nil)"
	}
	type Alias UpdatesDifferenceSlice
	return fmt.Sprintf("UpdatesDifferenceSlice%+v", Alias(*d))
}

// FillFrom fills UpdatesDifferenceSlice from given interface.
func (d *UpdatesDifferenceSlice) FillFrom(from interface {
	GetNewMessages() (value []MessageClass)
	GetNewEncryptedMessages() (value []EncryptedMessageClass)
	GetOtherUpdates() (value []UpdateClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetIntermediateState() (value UpdatesState)
}) {
	d.NewMessages = from.GetNewMessages()
	d.NewEncryptedMessages = from.GetNewEncryptedMessages()
	d.OtherUpdates = from.GetOtherUpdates()
	d.Chats = from.GetChats()
	d.Users = from.GetUsers()
	d.IntermediateState = from.GetIntermediateState()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesDifferenceSlice) TypeID() uint32 {
	return UpdatesDifferenceSliceTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesDifferenceSlice) TypeName() string {
	return "updates.differenceSlice"
}

// TypeInfo returns info about TL type.
func (d *UpdatesDifferenceSlice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.differenceSlice",
		ID:   UpdatesDifferenceSliceTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NewMessages",
			SchemaName: "new_messages",
		},
		{
			Name:       "NewEncryptedMessages",
			SchemaName: "new_encrypted_messages",
		},
		{
			Name:       "OtherUpdates",
			SchemaName: "other_updates",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "IntermediateState",
			SchemaName: "intermediate_state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *UpdatesDifferenceSlice) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceSlice#a8fb1981 as nil")
	}
	b.PutID(UpdatesDifferenceSliceTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *UpdatesDifferenceSlice) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceSlice#a8fb1981 as nil")
	}
	b.PutVectorHeader(len(d.NewMessages))
	for idx, v := range d.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.NewEncryptedMessages))
	for idx, v := range d.NewEncryptedMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_encrypted_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_encrypted_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.OtherUpdates))
	for idx, v := range d.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field users element with index %d: %w", idx, err)
		}
	}
	if err := d.IntermediateState.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field intermediate_state: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *UpdatesDifferenceSlice) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceSlice#a8fb1981 to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *UpdatesDifferenceSlice) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceSlice#a8fb1981 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_messages: %w", err)
		}

		if headerLen > 0 {
			d.NewMessages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_messages: %w", err)
			}
			d.NewMessages = append(d.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_encrypted_messages: %w", err)
		}

		if headerLen > 0 {
			d.NewEncryptedMessages = make([]EncryptedMessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEncryptedMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_encrypted_messages: %w", err)
			}
			d.NewEncryptedMessages = append(d.NewEncryptedMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field other_updates: %w", err)
		}

		if headerLen > 0 {
			d.OtherUpdates = make([]UpdateClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field other_updates: %w", err)
			}
			d.OtherUpdates = append(d.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field chats: %w", err)
		}

		if headerLen > 0 {
			d.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field users: %w", err)
		}

		if headerLen > 0 {
			d.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	{
		if err := d.IntermediateState.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field intermediate_state: %w", err)
		}
	}
	return nil
}

// GetNewMessages returns value of NewMessages field.
func (d *UpdatesDifferenceSlice) GetNewMessages() (value []MessageClass) {
	return d.NewMessages
}

// GetNewEncryptedMessages returns value of NewEncryptedMessages field.
func (d *UpdatesDifferenceSlice) GetNewEncryptedMessages() (value []EncryptedMessageClass) {
	return d.NewEncryptedMessages
}

// GetOtherUpdates returns value of OtherUpdates field.
func (d *UpdatesDifferenceSlice) GetOtherUpdates() (value []UpdateClass) {
	return d.OtherUpdates
}

// GetChats returns value of Chats field.
func (d *UpdatesDifferenceSlice) GetChats() (value []ChatClass) {
	return d.Chats
}

// GetUsers returns value of Users field.
func (d *UpdatesDifferenceSlice) GetUsers() (value []UserClass) {
	return d.Users
}

// GetIntermediateState returns value of IntermediateState field.
func (d *UpdatesDifferenceSlice) GetIntermediateState() (value UpdatesState) {
	return d.IntermediateState
}

// MapNewMessages returns field NewMessages wrapped in MessageClassArray helper.
func (d *UpdatesDifferenceSlice) MapNewMessages() (value MessageClassArray) {
	return MessageClassArray(d.NewMessages)
}

// MapNewEncryptedMessages returns field NewEncryptedMessages wrapped in EncryptedMessageClassArray helper.
func (d *UpdatesDifferenceSlice) MapNewEncryptedMessages() (value EncryptedMessageClassArray) {
	return EncryptedMessageClassArray(d.NewEncryptedMessages)
}

// MapOtherUpdates returns field OtherUpdates wrapped in UpdateClassArray helper.
func (d *UpdatesDifferenceSlice) MapOtherUpdates() (value UpdateClassArray) {
	return UpdateClassArray(d.OtherUpdates)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (d *UpdatesDifferenceSlice) MapChats() (value ChatClassArray) {
	return ChatClassArray(d.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (d *UpdatesDifferenceSlice) MapUsers() (value UserClassArray) {
	return UserClassArray(d.Users)
}

// UpdatesDifferenceTooLong represents TL type `updates.differenceTooLong#4afe8f6d`.
// The difference is too long¹, and the specified state must be used to refetch updates.
//
// Links:
//  1) https://core.telegram.org/api/updates#recovering-gaps
//
// See https://core.telegram.org/constructor/updates.differenceTooLong for reference.
type UpdatesDifferenceTooLong struct {
	// The new state to use.
	Pts int
}

// UpdatesDifferenceTooLongTypeID is TL type id of UpdatesDifferenceTooLong.
const UpdatesDifferenceTooLongTypeID = 0x4afe8f6d

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifferenceTooLong) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifferenceTooLong.
var (
	_ bin.Encoder     = &UpdatesDifferenceTooLong{}
	_ bin.Decoder     = &UpdatesDifferenceTooLong{}
	_ bin.BareEncoder = &UpdatesDifferenceTooLong{}
	_ bin.BareDecoder = &UpdatesDifferenceTooLong{}

	_ UpdatesDifferenceClass = &UpdatesDifferenceTooLong{}
)

func (d *UpdatesDifferenceTooLong) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Pts == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifferenceTooLong) String() string {
	if d == nil {
		return "UpdatesDifferenceTooLong(nil)"
	}
	type Alias UpdatesDifferenceTooLong
	return fmt.Sprintf("UpdatesDifferenceTooLong%+v", Alias(*d))
}

// FillFrom fills UpdatesDifferenceTooLong from given interface.
func (d *UpdatesDifferenceTooLong) FillFrom(from interface {
	GetPts() (value int)
}) {
	d.Pts = from.GetPts()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesDifferenceTooLong) TypeID() uint32 {
	return UpdatesDifferenceTooLongTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesDifferenceTooLong) TypeName() string {
	return "updates.differenceTooLong"
}

// TypeInfo returns info about TL type.
func (d *UpdatesDifferenceTooLong) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.differenceTooLong",
		ID:   UpdatesDifferenceTooLongTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pts",
			SchemaName: "pts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *UpdatesDifferenceTooLong) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceTooLong#4afe8f6d as nil")
	}
	b.PutID(UpdatesDifferenceTooLongTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *UpdatesDifferenceTooLong) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceTooLong#4afe8f6d as nil")
	}
	b.PutInt(d.Pts)
	return nil
}

// Decode implements bin.Decoder.
func (d *UpdatesDifferenceTooLong) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceTooLong#4afe8f6d to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceTooLongTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.differenceTooLong#4afe8f6d: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *UpdatesDifferenceTooLong) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceTooLong#4afe8f6d to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceTooLong#4afe8f6d: field pts: %w", err)
		}
		d.Pts = value
	}
	return nil
}

// GetPts returns value of Pts field.
func (d *UpdatesDifferenceTooLong) GetPts() (value int) {
	return d.Pts
}

// UpdatesDifferenceClass represents updates.Difference generic type.
//
// See https://core.telegram.org/type/updates.Difference for reference.
//
// Example:
//  g, err := tg.DecodeUpdatesDifference(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UpdatesDifferenceEmpty: // updates.differenceEmpty#5d75a138
//  case *tg.UpdatesDifference: // updates.difference#f49ca0
//  case *tg.UpdatesDifferenceSlice: // updates.differenceSlice#a8fb1981
//  case *tg.UpdatesDifferenceTooLong: // updates.differenceTooLong#4afe8f6d
//  default: panic(v)
//  }
type UpdatesDifferenceClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UpdatesDifferenceClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeUpdatesDifference implements binary de-serialization for UpdatesDifferenceClass.
func DecodeUpdatesDifference(buf *bin.Buffer) (UpdatesDifferenceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UpdatesDifferenceEmptyTypeID:
		// Decoding updates.differenceEmpty#5d75a138.
		v := UpdatesDifferenceEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesDifferenceTypeID:
		// Decoding updates.difference#f49ca0.
		v := UpdatesDifference{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesDifferenceSliceTypeID:
		// Decoding updates.differenceSlice#a8fb1981.
		v := UpdatesDifferenceSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesDifferenceTooLongTypeID:
		// Decoding updates.differenceTooLong#4afe8f6d.
		v := UpdatesDifferenceTooLong{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", bin.NewUnexpectedID(id))
	}
}

// UpdatesDifference boxes the UpdatesDifferenceClass providing a helper.
type UpdatesDifferenceBox struct {
	Difference UpdatesDifferenceClass
}

// Decode implements bin.Decoder for UpdatesDifferenceBox.
func (b *UpdatesDifferenceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UpdatesDifferenceBox to nil")
	}
	v, err := DecodeUpdatesDifference(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Difference = v
	return nil
}

// Encode implements bin.Encode for UpdatesDifferenceBox.
func (b *UpdatesDifferenceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Difference == nil {
		return fmt.Errorf("unable to encode UpdatesDifferenceClass as nil")
	}
	return b.Difference.Encode(buf)
}
