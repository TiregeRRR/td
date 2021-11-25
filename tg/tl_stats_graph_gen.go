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

// StatsGraphAsync represents TL type `statsGraphAsync#4a27eb2d`.
// This channel statistics graph¹ must be generated asynchronously using stats
// loadAsyncGraph² to reduce server load
//
// Links:
//  1) https://core.telegram.org/api/stats
//  2) https://core.telegram.org/method/stats.loadAsyncGraph
//
// See https://core.telegram.org/constructor/statsGraphAsync for reference.
type StatsGraphAsync struct {
	// Token to use for fetching the async graph
	Token string
}

// StatsGraphAsyncTypeID is TL type id of StatsGraphAsync.
const StatsGraphAsyncTypeID = 0x4a27eb2d

// construct implements constructor of StatsGraphClass.
func (s StatsGraphAsync) construct() StatsGraphClass { return &s }

// Ensuring interfaces in compile-time for StatsGraphAsync.
var (
	_ bin.Encoder     = &StatsGraphAsync{}
	_ bin.Decoder     = &StatsGraphAsync{}
	_ bin.BareEncoder = &StatsGraphAsync{}
	_ bin.BareDecoder = &StatsGraphAsync{}

	_ StatsGraphClass = &StatsGraphAsync{}
)

func (s *StatsGraphAsync) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Token == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGraphAsync) String() string {
	if s == nil {
		return "StatsGraphAsync(nil)"
	}
	type Alias StatsGraphAsync
	return fmt.Sprintf("StatsGraphAsync%+v", Alias(*s))
}

// FillFrom fills StatsGraphAsync from given interface.
func (s *StatsGraphAsync) FillFrom(from interface {
	GetToken() (value string)
}) {
	s.Token = from.GetToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGraphAsync) TypeID() uint32 {
	return StatsGraphAsyncTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGraphAsync) TypeName() string {
	return "statsGraphAsync"
}

// TypeInfo returns info about TL type.
func (s *StatsGraphAsync) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsGraphAsync",
		ID:   StatsGraphAsyncTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsGraphAsync) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGraphAsync#4a27eb2d as nil")
	}
	b.PutID(StatsGraphAsyncTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsGraphAsync) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGraphAsync#4a27eb2d as nil")
	}
	b.PutString(s.Token)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGraphAsync) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGraphAsync#4a27eb2d to nil")
	}
	if err := b.ConsumeID(StatsGraphAsyncTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGraphAsync#4a27eb2d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsGraphAsync) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGraphAsync#4a27eb2d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statsGraphAsync#4a27eb2d: field token: %w", err)
		}
		s.Token = value
	}
	return nil
}

// GetToken returns value of Token field.
func (s *StatsGraphAsync) GetToken() (value string) {
	return s.Token
}

// StatsGraphError represents TL type `statsGraphError#bedc9822`.
// An error occurred while generating the statistics graph¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/statsGraphError for reference.
type StatsGraphError struct {
	// The error
	Error string
}

// StatsGraphErrorTypeID is TL type id of StatsGraphError.
const StatsGraphErrorTypeID = 0xbedc9822

// construct implements constructor of StatsGraphClass.
func (s StatsGraphError) construct() StatsGraphClass { return &s }

// Ensuring interfaces in compile-time for StatsGraphError.
var (
	_ bin.Encoder     = &StatsGraphError{}
	_ bin.Decoder     = &StatsGraphError{}
	_ bin.BareEncoder = &StatsGraphError{}
	_ bin.BareDecoder = &StatsGraphError{}

	_ StatsGraphClass = &StatsGraphError{}
)

func (s *StatsGraphError) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Error == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGraphError) String() string {
	if s == nil {
		return "StatsGraphError(nil)"
	}
	type Alias StatsGraphError
	return fmt.Sprintf("StatsGraphError%+v", Alias(*s))
}

// FillFrom fills StatsGraphError from given interface.
func (s *StatsGraphError) FillFrom(from interface {
	GetError() (value string)
}) {
	s.Error = from.GetError()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGraphError) TypeID() uint32 {
	return StatsGraphErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGraphError) TypeName() string {
	return "statsGraphError"
}

// TypeInfo returns info about TL type.
func (s *StatsGraphError) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsGraphError",
		ID:   StatsGraphErrorTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Error",
			SchemaName: "error",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsGraphError) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGraphError#bedc9822 as nil")
	}
	b.PutID(StatsGraphErrorTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsGraphError) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGraphError#bedc9822 as nil")
	}
	b.PutString(s.Error)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGraphError) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGraphError#bedc9822 to nil")
	}
	if err := b.ConsumeID(StatsGraphErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGraphError#bedc9822: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsGraphError) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGraphError#bedc9822 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statsGraphError#bedc9822: field error: %w", err)
		}
		s.Error = value
	}
	return nil
}

// GetError returns value of Error field.
func (s *StatsGraphError) GetError() (value string) {
	return s.Error
}

// StatsGraph represents TL type `statsGraph#8ea464b6`.
// Channel statistics graph¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/statsGraph for reference.
type StatsGraph struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Statistics data
	JSON DataJSON
	// Zoom token
	//
	// Use SetZoomToken and GetZoomToken helpers.
	ZoomToken string
}

// StatsGraphTypeID is TL type id of StatsGraph.
const StatsGraphTypeID = 0x8ea464b6

// construct implements constructor of StatsGraphClass.
func (s StatsGraph) construct() StatsGraphClass { return &s }

// Ensuring interfaces in compile-time for StatsGraph.
var (
	_ bin.Encoder     = &StatsGraph{}
	_ bin.Decoder     = &StatsGraph{}
	_ bin.BareEncoder = &StatsGraph{}
	_ bin.BareDecoder = &StatsGraph{}

	_ StatsGraphClass = &StatsGraph{}
)

func (s *StatsGraph) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.JSON.Zero()) {
		return false
	}
	if !(s.ZoomToken == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGraph) String() string {
	if s == nil {
		return "StatsGraph(nil)"
	}
	type Alias StatsGraph
	return fmt.Sprintf("StatsGraph%+v", Alias(*s))
}

// FillFrom fills StatsGraph from given interface.
func (s *StatsGraph) FillFrom(from interface {
	GetJSON() (value DataJSON)
	GetZoomToken() (value string, ok bool)
}) {
	s.JSON = from.GetJSON()
	if val, ok := from.GetZoomToken(); ok {
		s.ZoomToken = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGraph) TypeID() uint32 {
	return StatsGraphTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGraph) TypeName() string {
	return "statsGraph"
}

// TypeInfo returns info about TL type.
func (s *StatsGraph) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsGraph",
		ID:   StatsGraphTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JSON",
			SchemaName: "json",
		},
		{
			Name:       "ZoomToken",
			SchemaName: "zoom_token",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsGraph) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGraph#8ea464b6 as nil")
	}
	b.PutID(StatsGraphTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsGraph) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGraph#8ea464b6 as nil")
	}
	if !(s.ZoomToken == "") {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode statsGraph#8ea464b6: field flags: %w", err)
	}
	if err := s.JSON.Encode(b); err != nil {
		return fmt.Errorf("unable to encode statsGraph#8ea464b6: field json: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutString(s.ZoomToken)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGraph) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGraph#8ea464b6 to nil")
	}
	if err := b.ConsumeID(StatsGraphTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGraph#8ea464b6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsGraph) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGraph#8ea464b6 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode statsGraph#8ea464b6: field flags: %w", err)
		}
	}
	{
		if err := s.JSON.Decode(b); err != nil {
			return fmt.Errorf("unable to decode statsGraph#8ea464b6: field json: %w", err)
		}
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statsGraph#8ea464b6: field zoom_token: %w", err)
		}
		s.ZoomToken = value
	}
	return nil
}

// GetJSON returns value of JSON field.
func (s *StatsGraph) GetJSON() (value DataJSON) {
	return s.JSON
}

// SetZoomToken sets value of ZoomToken conditional field.
func (s *StatsGraph) SetZoomToken(value string) {
	s.Flags.Set(0)
	s.ZoomToken = value
}

// GetZoomToken returns value of ZoomToken conditional field and
// boolean which is true if field was set.
func (s *StatsGraph) GetZoomToken() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ZoomToken, true
}

// StatsGraphClass represents StatsGraph generic type.
//
// See https://core.telegram.org/type/StatsGraph for reference.
//
// Example:
//  g, err := tg.DecodeStatsGraph(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.StatsGraphAsync: // statsGraphAsync#4a27eb2d
//  case *tg.StatsGraphError: // statsGraphError#bedc9822
//  case *tg.StatsGraph: // statsGraph#8ea464b6
//  default: panic(v)
//  }
type StatsGraphClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StatsGraphClass

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

// DecodeStatsGraph implements binary de-serialization for StatsGraphClass.
func DecodeStatsGraph(buf *bin.Buffer) (StatsGraphClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StatsGraphAsyncTypeID:
		// Decoding statsGraphAsync#4a27eb2d.
		v := StatsGraphAsync{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatsGraphClass: %w", err)
		}
		return &v, nil
	case StatsGraphErrorTypeID:
		// Decoding statsGraphError#bedc9822.
		v := StatsGraphError{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatsGraphClass: %w", err)
		}
		return &v, nil
	case StatsGraphTypeID:
		// Decoding statsGraph#8ea464b6.
		v := StatsGraph{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StatsGraphClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StatsGraphClass: %w", bin.NewUnexpectedID(id))
	}
}

// StatsGraph boxes the StatsGraphClass providing a helper.
type StatsGraphBox struct {
	StatsGraph StatsGraphClass
}

// Decode implements bin.Decoder for StatsGraphBox.
func (b *StatsGraphBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StatsGraphBox to nil")
	}
	v, err := DecodeStatsGraph(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StatsGraph = v
	return nil
}

// Encode implements bin.Encode for StatsGraphBox.
func (b *StatsGraphBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StatsGraph == nil {
		return fmt.Errorf("unable to encode StatsGraphClass as nil")
	}
	return b.StatsGraph.Encode(buf)
}
