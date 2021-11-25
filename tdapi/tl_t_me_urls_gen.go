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

// TMeURLs represents TL type `tMeUrls#655b1f52`.
type TMeURLs struct {
	// List of URLs
	URLs []TMeURL
}

// TMeURLsTypeID is TL type id of TMeURLs.
const TMeURLsTypeID = 0x655b1f52

// Ensuring interfaces in compile-time for TMeURLs.
var (
	_ bin.Encoder     = &TMeURLs{}
	_ bin.Decoder     = &TMeURLs{}
	_ bin.BareEncoder = &TMeURLs{}
	_ bin.BareDecoder = &TMeURLs{}
)

func (t *TMeURLs) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.URLs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TMeURLs) String() string {
	if t == nil {
		return "TMeURLs(nil)"
	}
	type Alias TMeURLs
	return fmt.Sprintf("TMeURLs%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TMeURLs) TypeID() uint32 {
	return TMeURLsTypeID
}

// TypeName returns name of type in TL schema.
func (*TMeURLs) TypeName() string {
	return "tMeUrls"
}

// TypeInfo returns info about TL type.
func (t *TMeURLs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "tMeUrls",
		ID:   TMeURLsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URLs",
			SchemaName: "urls",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TMeURLs) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode tMeUrls#655b1f52 as nil")
	}
	b.PutID(TMeURLsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TMeURLs) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode tMeUrls#655b1f52 as nil")
	}
	b.PutInt(len(t.URLs))
	for idx, v := range t.URLs {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare tMeUrls#655b1f52: field urls element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TMeURLs) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode tMeUrls#655b1f52 to nil")
	}
	if err := b.ConsumeID(TMeURLsTypeID); err != nil {
		return fmt.Errorf("unable to decode tMeUrls#655b1f52: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TMeURLs) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode tMeUrls#655b1f52 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode tMeUrls#655b1f52: field urls: %w", err)
		}

		if headerLen > 0 {
			t.URLs = make([]TMeURL, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TMeURL
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare tMeUrls#655b1f52: field urls: %w", err)
			}
			t.URLs = append(t.URLs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TMeURLs) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode tMeUrls#655b1f52 as nil")
	}
	b.ObjStart()
	b.PutID("tMeUrls")
	b.FieldStart("urls")
	b.ArrStart()
	for idx, v := range t.URLs {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode tMeUrls#655b1f52: field urls element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TMeURLs) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode tMeUrls#655b1f52 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("tMeUrls"); err != nil {
				return fmt.Errorf("unable to decode tMeUrls#655b1f52: %w", err)
			}
		case "urls":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value TMeURL
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode tMeUrls#655b1f52: field urls: %w", err)
				}
				t.URLs = append(t.URLs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode tMeUrls#655b1f52: field urls: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURLs returns value of URLs field.
func (t *TMeURLs) GetURLs() (value []TMeURL) {
	return t.URLs
}
