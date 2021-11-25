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

// VideoNote represents TL type `videoNote#fbb96a3a`.
type VideoNote struct {
	// Duration of the video, in seconds; as defined by the sender
	Duration int32
	// Video width and height; as defined by the sender
	Length int32
	// Video minithumbnail; may be null
	Minithumbnail Minithumbnail
	// Video thumbnail in JPEG format; as defined by the sender; may be null
	Thumbnail Thumbnail
	// File containing the video
	Video File
}

// VideoNoteTypeID is TL type id of VideoNote.
const VideoNoteTypeID = 0xfbb96a3a

// Ensuring interfaces in compile-time for VideoNote.
var (
	_ bin.Encoder     = &VideoNote{}
	_ bin.Decoder     = &VideoNote{}
	_ bin.BareEncoder = &VideoNote{}
	_ bin.BareDecoder = &VideoNote{}
)

func (v *VideoNote) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Duration == 0) {
		return false
	}
	if !(v.Length == 0) {
		return false
	}
	if !(v.Minithumbnail.Zero()) {
		return false
	}
	if !(v.Thumbnail.Zero()) {
		return false
	}
	if !(v.Video.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VideoNote) String() string {
	if v == nil {
		return "VideoNote(nil)"
	}
	type Alias VideoNote
	return fmt.Sprintf("VideoNote%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VideoNote) TypeID() uint32 {
	return VideoNoteTypeID
}

// TypeName returns name of type in TL schema.
func (*VideoNote) TypeName() string {
	return "videoNote"
}

// TypeInfo returns info about TL type.
func (v *VideoNote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "videoNote",
		ID:   VideoNoteTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Length",
			SchemaName: "length",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "Video",
			SchemaName: "video",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VideoNote) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoNote#fbb96a3a as nil")
	}
	b.PutID(VideoNoteTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VideoNote) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoNote#fbb96a3a as nil")
	}
	b.PutInt32(v.Duration)
	b.PutInt32(v.Length)
	if err := v.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoNote#fbb96a3a: field minithumbnail: %w", err)
	}
	if err := v.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoNote#fbb96a3a: field thumbnail: %w", err)
	}
	if err := v.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoNote#fbb96a3a: field video: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VideoNote) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoNote#fbb96a3a to nil")
	}
	if err := b.ConsumeID(VideoNoteTypeID); err != nil {
		return fmt.Errorf("unable to decode videoNote#fbb96a3a: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VideoNote) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoNote#fbb96a3a to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode videoNote#fbb96a3a: field duration: %w", err)
		}
		v.Duration = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode videoNote#fbb96a3a: field length: %w", err)
		}
		v.Length = value
	}
	{
		if err := v.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoNote#fbb96a3a: field minithumbnail: %w", err)
		}
	}
	{
		if err := v.Thumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoNote#fbb96a3a: field thumbnail: %w", err)
		}
	}
	{
		if err := v.Video.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoNote#fbb96a3a: field video: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *VideoNote) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode videoNote#fbb96a3a as nil")
	}
	b.ObjStart()
	b.PutID("videoNote")
	b.FieldStart("duration")
	b.PutInt32(v.Duration)
	b.FieldStart("length")
	b.PutInt32(v.Length)
	b.FieldStart("minithumbnail")
	if err := v.Minithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode videoNote#fbb96a3a: field minithumbnail: %w", err)
	}
	b.FieldStart("thumbnail")
	if err := v.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode videoNote#fbb96a3a: field thumbnail: %w", err)
	}
	b.FieldStart("video")
	if err := v.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode videoNote#fbb96a3a: field video: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *VideoNote) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode videoNote#fbb96a3a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("videoNote"); err != nil {
				return fmt.Errorf("unable to decode videoNote#fbb96a3a: %w", err)
			}
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode videoNote#fbb96a3a: field duration: %w", err)
			}
			v.Duration = value
		case "length":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode videoNote#fbb96a3a: field length: %w", err)
			}
			v.Length = value
		case "minithumbnail":
			if err := v.Minithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode videoNote#fbb96a3a: field minithumbnail: %w", err)
			}
		case "thumbnail":
			if err := v.Thumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode videoNote#fbb96a3a: field thumbnail: %w", err)
			}
		case "video":
			if err := v.Video.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode videoNote#fbb96a3a: field video: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDuration returns value of Duration field.
func (v *VideoNote) GetDuration() (value int32) {
	return v.Duration
}

// GetLength returns value of Length field.
func (v *VideoNote) GetLength() (value int32) {
	return v.Length
}

// GetMinithumbnail returns value of Minithumbnail field.
func (v *VideoNote) GetMinithumbnail() (value Minithumbnail) {
	return v.Minithumbnail
}

// GetThumbnail returns value of Thumbnail field.
func (v *VideoNote) GetThumbnail() (value Thumbnail) {
	return v.Thumbnail
}

// GetVideo returns value of Video field.
func (v *VideoNote) GetVideo() (value File) {
	return v.Video
}
