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

// TMeURL represents TL type `tMeUrl#bc00fa42`.
type TMeURL struct {
	// URL
	URL string
	// Type of the URL
	Type TMeURLTypeClass
}

// TMeURLTypeID is TL type id of TMeURL.
const TMeURLTypeID = 0xbc00fa42

// Ensuring interfaces in compile-time for TMeURL.
var (
	_ bin.Encoder     = &TMeURL{}
	_ bin.Decoder     = &TMeURL{}
	_ bin.BareEncoder = &TMeURL{}
	_ bin.BareDecoder = &TMeURL{}
)

func (t *TMeURL) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.URL == "") {
		return false
	}
	if !(t.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TMeURL) String() string {
	if t == nil {
		return "TMeURL(nil)"
	}
	type Alias TMeURL
	return fmt.Sprintf("TMeURL%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TMeURL) TypeID() uint32 {
	return TMeURLTypeID
}

// TypeName returns name of type in TL schema.
func (*TMeURL) TypeName() string {
	return "tMeUrl"
}

// TypeInfo returns info about TL type.
func (t *TMeURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "tMeUrl",
		ID:   TMeURLTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TMeURL) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode tMeUrl#bc00fa42 as nil")
	}
	b.PutID(TMeURLTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TMeURL) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode tMeUrl#bc00fa42 as nil")
	}
	b.PutString(t.URL)
	if t.Type == nil {
		return fmt.Errorf("unable to encode tMeUrl#bc00fa42: field type is nil")
	}
	if err := t.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode tMeUrl#bc00fa42: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TMeURL) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode tMeUrl#bc00fa42 to nil")
	}
	if err := b.ConsumeID(TMeURLTypeID); err != nil {
		return fmt.Errorf("unable to decode tMeUrl#bc00fa42: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TMeURL) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode tMeUrl#bc00fa42 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode tMeUrl#bc00fa42: field url: %w", err)
		}
		t.URL = value
	}
	{
		value, err := DecodeTMeURLType(b)
		if err != nil {
			return fmt.Errorf("unable to decode tMeUrl#bc00fa42: field type: %w", err)
		}
		t.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TMeURL) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode tMeUrl#bc00fa42 as nil")
	}
	b.ObjStart()
	b.PutID("tMeUrl")
	b.FieldStart("url")
	b.PutString(t.URL)
	b.FieldStart("type")
	if t.Type == nil {
		return fmt.Errorf("unable to encode tMeUrl#bc00fa42: field type is nil")
	}
	if err := t.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode tMeUrl#bc00fa42: field type: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TMeURL) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode tMeUrl#bc00fa42 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("tMeUrl"); err != nil {
				return fmt.Errorf("unable to decode tMeUrl#bc00fa42: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode tMeUrl#bc00fa42: field url: %w", err)
			}
			t.URL = value
		case "type":
			value, err := DecodeTDLibJSONTMeURLType(b)
			if err != nil {
				return fmt.Errorf("unable to decode tMeUrl#bc00fa42: field type: %w", err)
			}
			t.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (t *TMeURL) GetURL() (value string) {
	return t.URL
}

// GetType returns value of Type field.
func (t *TMeURL) GetType() (value TMeURLTypeClass) {
	return t.Type
}
