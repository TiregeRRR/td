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

// StoryViews represents TL type `storyViews#8d595cd6`.
// Aggregated view and reaction information of a story¹.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/constructor/storyViews for reference.
type StoryViews struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// HasViewers field of StoryViews.
	HasViewers bool
	// View counter of the story
	ViewsCount int
	// Forward counter of the story
	//
	// Use SetForwardsCount and GetForwardsCount helpers.
	ForwardsCount int
	// All reactions sent to this story
	//
	// Use SetReactions and GetReactions helpers.
	Reactions []ReactionCount
	// Number of reactions added to the story
	//
	// Use SetReactionsCount and GetReactionsCount helpers.
	ReactionsCount int
	// User IDs of some recent viewers of the story
	//
	// Use SetRecentViewers and GetRecentViewers helpers.
	RecentViewers []int64
}

// StoryViewsTypeID is TL type id of StoryViews.
const StoryViewsTypeID = 0x8d595cd6

// Ensuring interfaces in compile-time for StoryViews.
var (
	_ bin.Encoder     = &StoryViews{}
	_ bin.Decoder     = &StoryViews{}
	_ bin.BareEncoder = &StoryViews{}
	_ bin.BareDecoder = &StoryViews{}
)

func (s *StoryViews) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.HasViewers == false) {
		return false
	}
	if !(s.ViewsCount == 0) {
		return false
	}
	if !(s.ForwardsCount == 0) {
		return false
	}
	if !(s.Reactions == nil) {
		return false
	}
	if !(s.ReactionsCount == 0) {
		return false
	}
	if !(s.RecentViewers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryViews) String() string {
	if s == nil {
		return "StoryViews(nil)"
	}
	type Alias StoryViews
	return fmt.Sprintf("StoryViews%+v", Alias(*s))
}

// FillFrom fills StoryViews from given interface.
func (s *StoryViews) FillFrom(from interface {
	GetHasViewers() (value bool)
	GetViewsCount() (value int)
	GetForwardsCount() (value int, ok bool)
	GetReactions() (value []ReactionCount, ok bool)
	GetReactionsCount() (value int, ok bool)
	GetRecentViewers() (value []int64, ok bool)
}) {
	s.HasViewers = from.GetHasViewers()
	s.ViewsCount = from.GetViewsCount()
	if val, ok := from.GetForwardsCount(); ok {
		s.ForwardsCount = val
	}

	if val, ok := from.GetReactions(); ok {
		s.Reactions = val
	}

	if val, ok := from.GetReactionsCount(); ok {
		s.ReactionsCount = val
	}

	if val, ok := from.GetRecentViewers(); ok {
		s.RecentViewers = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryViews) TypeID() uint32 {
	return StoryViewsTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryViews) TypeName() string {
	return "storyViews"
}

// TypeInfo returns info about TL type.
func (s *StoryViews) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyViews",
		ID:   StoryViewsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasViewers",
			SchemaName: "has_viewers",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "ViewsCount",
			SchemaName: "views_count",
		},
		{
			Name:       "ForwardsCount",
			SchemaName: "forwards_count",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Reactions",
			SchemaName: "reactions",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "ReactionsCount",
			SchemaName: "reactions_count",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "RecentViewers",
			SchemaName: "recent_viewers",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryViews) SetFlags() {
	if !(s.HasViewers == false) {
		s.Flags.Set(1)
	}
	if !(s.ForwardsCount == 0) {
		s.Flags.Set(2)
	}
	if !(s.Reactions == nil) {
		s.Flags.Set(3)
	}
	if !(s.ReactionsCount == 0) {
		s.Flags.Set(4)
	}
	if !(s.RecentViewers == nil) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *StoryViews) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViews#8d595cd6 as nil")
	}
	b.PutID(StoryViewsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryViews) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViews#8d595cd6 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViews#8d595cd6: field flags: %w", err)
	}
	b.PutInt(s.ViewsCount)
	if s.Flags.Has(2) {
		b.PutInt(s.ForwardsCount)
	}
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Reactions))
		for idx, v := range s.Reactions {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode storyViews#8d595cd6: field reactions element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(4) {
		b.PutInt(s.ReactionsCount)
	}
	if s.Flags.Has(0) {
		b.PutVectorHeader(len(s.RecentViewers))
		for _, v := range s.RecentViewers {
			b.PutLong(v)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryViews) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViews#8d595cd6 to nil")
	}
	if err := b.ConsumeID(StoryViewsTypeID); err != nil {
		return fmt.Errorf("unable to decode storyViews#8d595cd6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryViews) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViews#8d595cd6 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyViews#8d595cd6: field flags: %w", err)
		}
	}
	s.HasViewers = s.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#8d595cd6: field views_count: %w", err)
		}
		s.ViewsCount = value
	}
	if s.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#8d595cd6: field forwards_count: %w", err)
		}
		s.ForwardsCount = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#8d595cd6: field reactions: %w", err)
		}

		if headerLen > 0 {
			s.Reactions = make([]ReactionCount, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ReactionCount
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode storyViews#8d595cd6: field reactions: %w", err)
			}
			s.Reactions = append(s.Reactions, value)
		}
	}
	if s.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#8d595cd6: field reactions_count: %w", err)
		}
		s.ReactionsCount = value
	}
	if s.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode storyViews#8d595cd6: field recent_viewers: %w", err)
		}

		if headerLen > 0 {
			s.RecentViewers = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode storyViews#8d595cd6: field recent_viewers: %w", err)
			}
			s.RecentViewers = append(s.RecentViewers, value)
		}
	}
	return nil
}

// SetHasViewers sets value of HasViewers conditional field.
func (s *StoryViews) SetHasViewers(value bool) {
	if value {
		s.Flags.Set(1)
		s.HasViewers = true
	} else {
		s.Flags.Unset(1)
		s.HasViewers = false
	}
}

// GetHasViewers returns value of HasViewers conditional field.
func (s *StoryViews) GetHasViewers() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// GetViewsCount returns value of ViewsCount field.
func (s *StoryViews) GetViewsCount() (value int) {
	if s == nil {
		return
	}
	return s.ViewsCount
}

// SetForwardsCount sets value of ForwardsCount conditional field.
func (s *StoryViews) SetForwardsCount(value int) {
	s.Flags.Set(2)
	s.ForwardsCount = value
}

// GetForwardsCount returns value of ForwardsCount conditional field and
// boolean which is true if field was set.
func (s *StoryViews) GetForwardsCount() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ForwardsCount, true
}

// SetReactions sets value of Reactions conditional field.
func (s *StoryViews) SetReactions(value []ReactionCount) {
	s.Flags.Set(3)
	s.Reactions = value
}

// GetReactions returns value of Reactions conditional field and
// boolean which is true if field was set.
func (s *StoryViews) GetReactions() (value []ReactionCount, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Reactions, true
}

// SetReactionsCount sets value of ReactionsCount conditional field.
func (s *StoryViews) SetReactionsCount(value int) {
	s.Flags.Set(4)
	s.ReactionsCount = value
}

// GetReactionsCount returns value of ReactionsCount conditional field and
// boolean which is true if field was set.
func (s *StoryViews) GetReactionsCount() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.ReactionsCount, true
}

// SetRecentViewers sets value of RecentViewers conditional field.
func (s *StoryViews) SetRecentViewers(value []int64) {
	s.Flags.Set(0)
	s.RecentViewers = value
}

// GetRecentViewers returns value of RecentViewers conditional field and
// boolean which is true if field was set.
func (s *StoryViews) GetRecentViewers() (value []int64, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.RecentViewers, true
}
