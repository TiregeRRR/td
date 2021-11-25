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

// InputDialogPeer represents TL type `inputDialogPeer#fcaafeb7`.
// A peer
//
// See https://core.telegram.org/constructor/inputDialogPeer for reference.
type InputDialogPeer struct {
	// Peer
	Peer InputPeerClass
}

// InputDialogPeerTypeID is TL type id of InputDialogPeer.
const InputDialogPeerTypeID = 0xfcaafeb7

// construct implements constructor of InputDialogPeerClass.
func (i InputDialogPeer) construct() InputDialogPeerClass { return &i }

// Ensuring interfaces in compile-time for InputDialogPeer.
var (
	_ bin.Encoder     = &InputDialogPeer{}
	_ bin.Decoder     = &InputDialogPeer{}
	_ bin.BareEncoder = &InputDialogPeer{}
	_ bin.BareDecoder = &InputDialogPeer{}

	_ InputDialogPeerClass = &InputDialogPeer{}
)

func (i *InputDialogPeer) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputDialogPeer) String() string {
	if i == nil {
		return "InputDialogPeer(nil)"
	}
	type Alias InputDialogPeer
	return fmt.Sprintf("InputDialogPeer%+v", Alias(*i))
}

// FillFrom fills InputDialogPeer from given interface.
func (i *InputDialogPeer) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	i.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputDialogPeer) TypeID() uint32 {
	return InputDialogPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*InputDialogPeer) TypeName() string {
	return "inputDialogPeer"
}

// TypeInfo returns info about TL type.
func (i *InputDialogPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputDialogPeer",
		ID:   InputDialogPeerTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputDialogPeer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDialogPeer#fcaafeb7 as nil")
	}
	b.PutID(InputDialogPeerTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputDialogPeer) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDialogPeer#fcaafeb7 as nil")
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputDialogPeer#fcaafeb7: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputDialogPeer#fcaafeb7: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDialogPeer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDialogPeer#fcaafeb7 to nil")
	}
	if err := b.ConsumeID(InputDialogPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDialogPeer#fcaafeb7: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputDialogPeer) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDialogPeer#fcaafeb7 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputDialogPeer#fcaafeb7: field peer: %w", err)
		}
		i.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputDialogPeer) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// InputDialogPeerFolder represents TL type `inputDialogPeerFolder#64600527`.
// All peers in a peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/constructor/inputDialogPeerFolder for reference.
type InputDialogPeerFolder struct {
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// InputDialogPeerFolderTypeID is TL type id of InputDialogPeerFolder.
const InputDialogPeerFolderTypeID = 0x64600527

// construct implements constructor of InputDialogPeerClass.
func (i InputDialogPeerFolder) construct() InputDialogPeerClass { return &i }

// Ensuring interfaces in compile-time for InputDialogPeerFolder.
var (
	_ bin.Encoder     = &InputDialogPeerFolder{}
	_ bin.Decoder     = &InputDialogPeerFolder{}
	_ bin.BareEncoder = &InputDialogPeerFolder{}
	_ bin.BareDecoder = &InputDialogPeerFolder{}

	_ InputDialogPeerClass = &InputDialogPeerFolder{}
)

func (i *InputDialogPeerFolder) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputDialogPeerFolder) String() string {
	if i == nil {
		return "InputDialogPeerFolder(nil)"
	}
	type Alias InputDialogPeerFolder
	return fmt.Sprintf("InputDialogPeerFolder%+v", Alias(*i))
}

// FillFrom fills InputDialogPeerFolder from given interface.
func (i *InputDialogPeerFolder) FillFrom(from interface {
	GetFolderID() (value int)
}) {
	i.FolderID = from.GetFolderID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputDialogPeerFolder) TypeID() uint32 {
	return InputDialogPeerFolderTypeID
}

// TypeName returns name of type in TL schema.
func (*InputDialogPeerFolder) TypeName() string {
	return "inputDialogPeerFolder"
}

// TypeInfo returns info about TL type.
func (i *InputDialogPeerFolder) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputDialogPeerFolder",
		ID:   InputDialogPeerFolderTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputDialogPeerFolder) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDialogPeerFolder#64600527 as nil")
	}
	b.PutID(InputDialogPeerFolderTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputDialogPeerFolder) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDialogPeerFolder#64600527 as nil")
	}
	b.PutInt(i.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDialogPeerFolder) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDialogPeerFolder#64600527 to nil")
	}
	if err := b.ConsumeID(InputDialogPeerFolderTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDialogPeerFolder#64600527: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputDialogPeerFolder) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDialogPeerFolder#64600527 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputDialogPeerFolder#64600527: field folder_id: %w", err)
		}
		i.FolderID = value
	}
	return nil
}

// GetFolderID returns value of FolderID field.
func (i *InputDialogPeerFolder) GetFolderID() (value int) {
	return i.FolderID
}

// InputDialogPeerClass represents InputDialogPeer generic type.
//
// See https://core.telegram.org/type/InputDialogPeer for reference.
//
// Example:
//  g, err := tg.DecodeInputDialogPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputDialogPeer: // inputDialogPeer#fcaafeb7
//  case *tg.InputDialogPeerFolder: // inputDialogPeerFolder#64600527
//  default: panic(v)
//  }
type InputDialogPeerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputDialogPeerClass

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

// DecodeInputDialogPeer implements binary de-serialization for InputDialogPeerClass.
func DecodeInputDialogPeer(buf *bin.Buffer) (InputDialogPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputDialogPeerTypeID:
		// Decoding inputDialogPeer#fcaafeb7.
		v := InputDialogPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDialogPeerClass: %w", err)
		}
		return &v, nil
	case InputDialogPeerFolderTypeID:
		// Decoding inputDialogPeerFolder#64600527.
		v := InputDialogPeerFolder{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDialogPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputDialogPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputDialogPeer boxes the InputDialogPeerClass providing a helper.
type InputDialogPeerBox struct {
	InputDialogPeer InputDialogPeerClass
}

// Decode implements bin.Decoder for InputDialogPeerBox.
func (b *InputDialogPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputDialogPeerBox to nil")
	}
	v, err := DecodeInputDialogPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputDialogPeer = v
	return nil
}

// Encode implements bin.Encode for InputDialogPeerBox.
func (b *InputDialogPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputDialogPeer == nil {
		return fmt.Errorf("unable to encode InputDialogPeerClass as nil")
	}
	return b.InputDialogPeer.Encode(buf)
}
