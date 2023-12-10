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

// StatsStoryStats represents TL type `stats.storyStats#50cd067c`.
//
// See https://core.telegram.org/constructor/stats.storyStats for reference.
type StatsStoryStats struct {
	// ViewsGraph field of StatsStoryStats.
	ViewsGraph StatsGraphClass
	// ReactionsByEmotionGraph field of StatsStoryStats.
	ReactionsByEmotionGraph StatsGraphClass
}

// StatsStoryStatsTypeID is TL type id of StatsStoryStats.
const StatsStoryStatsTypeID = 0x50cd067c

// Ensuring interfaces in compile-time for StatsStoryStats.
var (
	_ bin.Encoder     = &StatsStoryStats{}
	_ bin.Decoder     = &StatsStoryStats{}
	_ bin.BareEncoder = &StatsStoryStats{}
	_ bin.BareDecoder = &StatsStoryStats{}
)

func (s *StatsStoryStats) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ViewsGraph == nil) {
		return false
	}
	if !(s.ReactionsByEmotionGraph == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsStoryStats) String() string {
	if s == nil {
		return "StatsStoryStats(nil)"
	}
	type Alias StatsStoryStats
	return fmt.Sprintf("StatsStoryStats%+v", Alias(*s))
}

// FillFrom fills StatsStoryStats from given interface.
func (s *StatsStoryStats) FillFrom(from interface {
	GetViewsGraph() (value StatsGraphClass)
	GetReactionsByEmotionGraph() (value StatsGraphClass)
}) {
	s.ViewsGraph = from.GetViewsGraph()
	s.ReactionsByEmotionGraph = from.GetReactionsByEmotionGraph()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsStoryStats) TypeID() uint32 {
	return StatsStoryStatsTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsStoryStats) TypeName() string {
	return "stats.storyStats"
}

// TypeInfo returns info about TL type.
func (s *StatsStoryStats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.storyStats",
		ID:   StatsStoryStatsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ViewsGraph",
			SchemaName: "views_graph",
		},
		{
			Name:       "ReactionsByEmotionGraph",
			SchemaName: "reactions_by_emotion_graph",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsStoryStats) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stats.storyStats#50cd067c as nil")
	}
	b.PutID(StatsStoryStatsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsStoryStats) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stats.storyStats#50cd067c as nil")
	}
	if s.ViewsGraph == nil {
		return fmt.Errorf("unable to encode stats.storyStats#50cd067c: field views_graph is nil")
	}
	if err := s.ViewsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.storyStats#50cd067c: field views_graph: %w", err)
	}
	if s.ReactionsByEmotionGraph == nil {
		return fmt.Errorf("unable to encode stats.storyStats#50cd067c: field reactions_by_emotion_graph is nil")
	}
	if err := s.ReactionsByEmotionGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.storyStats#50cd067c: field reactions_by_emotion_graph: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsStoryStats) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stats.storyStats#50cd067c to nil")
	}
	if err := b.ConsumeID(StatsStoryStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.storyStats#50cd067c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsStoryStats) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stats.storyStats#50cd067c to nil")
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.storyStats#50cd067c: field views_graph: %w", err)
		}
		s.ViewsGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.storyStats#50cd067c: field reactions_by_emotion_graph: %w", err)
		}
		s.ReactionsByEmotionGraph = value
	}
	return nil
}

// GetViewsGraph returns value of ViewsGraph field.
func (s *StatsStoryStats) GetViewsGraph() (value StatsGraphClass) {
	if s == nil {
		return
	}
	return s.ViewsGraph
}

// GetReactionsByEmotionGraph returns value of ReactionsByEmotionGraph field.
func (s *StatsStoryStats) GetReactionsByEmotionGraph() (value StatsGraphClass) {
	if s == nil {
		return
	}
	return s.ReactionsByEmotionGraph
}
