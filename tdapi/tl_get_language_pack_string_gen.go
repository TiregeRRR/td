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

// GetLanguagePackStringRequest represents TL type `getLanguagePackString#8fcde73`.
type GetLanguagePackStringRequest struct {
	// Path to the language pack database in which strings are stored
	LanguagePackDatabasePath string
	// Localization target to which the language pack belongs
	LocalizationTarget string
	// Language pack identifier
	LanguagePackID string
	// Language pack key of the string to be returned
	Key string
}

// GetLanguagePackStringRequestTypeID is TL type id of GetLanguagePackStringRequest.
const GetLanguagePackStringRequestTypeID = 0x8fcde73

// Ensuring interfaces in compile-time for GetLanguagePackStringRequest.
var (
	_ bin.Encoder     = &GetLanguagePackStringRequest{}
	_ bin.Decoder     = &GetLanguagePackStringRequest{}
	_ bin.BareEncoder = &GetLanguagePackStringRequest{}
	_ bin.BareDecoder = &GetLanguagePackStringRequest{}
)

func (g *GetLanguagePackStringRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LanguagePackDatabasePath == "") {
		return false
	}
	if !(g.LocalizationTarget == "") {
		return false
	}
	if !(g.LanguagePackID == "") {
		return false
	}
	if !(g.Key == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLanguagePackStringRequest) String() string {
	if g == nil {
		return "GetLanguagePackStringRequest(nil)"
	}
	type Alias GetLanguagePackStringRequest
	return fmt.Sprintf("GetLanguagePackStringRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLanguagePackStringRequest) TypeID() uint32 {
	return GetLanguagePackStringRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLanguagePackStringRequest) TypeName() string {
	return "getLanguagePackString"
}

// TypeInfo returns info about TL type.
func (g *GetLanguagePackStringRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLanguagePackString",
		ID:   GetLanguagePackStringRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LanguagePackDatabasePath",
			SchemaName: "language_pack_database_path",
		},
		{
			Name:       "LocalizationTarget",
			SchemaName: "localization_target",
		},
		{
			Name:       "LanguagePackID",
			SchemaName: "language_pack_id",
		},
		{
			Name:       "Key",
			SchemaName: "key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLanguagePackStringRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLanguagePackString#8fcde73 as nil")
	}
	b.PutID(GetLanguagePackStringRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLanguagePackStringRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLanguagePackString#8fcde73 as nil")
	}
	b.PutString(g.LanguagePackDatabasePath)
	b.PutString(g.LocalizationTarget)
	b.PutString(g.LanguagePackID)
	b.PutString(g.Key)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLanguagePackStringRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLanguagePackString#8fcde73 to nil")
	}
	if err := b.ConsumeID(GetLanguagePackStringRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLanguagePackStringRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLanguagePackString#8fcde73 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field language_pack_database_path: %w", err)
		}
		g.LanguagePackDatabasePath = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field localization_target: %w", err)
		}
		g.LocalizationTarget = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field language_pack_id: %w", err)
		}
		g.LanguagePackID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field key: %w", err)
		}
		g.Key = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLanguagePackStringRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLanguagePackString#8fcde73 as nil")
	}
	b.ObjStart()
	b.PutID("getLanguagePackString")
	b.FieldStart("language_pack_database_path")
	b.PutString(g.LanguagePackDatabasePath)
	b.FieldStart("localization_target")
	b.PutString(g.LocalizationTarget)
	b.FieldStart("language_pack_id")
	b.PutString(g.LanguagePackID)
	b.FieldStart("key")
	b.PutString(g.Key)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLanguagePackStringRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLanguagePackString#8fcde73 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLanguagePackString"); err != nil {
				return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: %w", err)
			}
		case "language_pack_database_path":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field language_pack_database_path: %w", err)
			}
			g.LanguagePackDatabasePath = value
		case "localization_target":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field localization_target: %w", err)
			}
			g.LocalizationTarget = value
		case "language_pack_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field language_pack_id: %w", err)
			}
			g.LanguagePackID = value
		case "key":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getLanguagePackString#8fcde73: field key: %w", err)
			}
			g.Key = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLanguagePackDatabasePath returns value of LanguagePackDatabasePath field.
func (g *GetLanguagePackStringRequest) GetLanguagePackDatabasePath() (value string) {
	return g.LanguagePackDatabasePath
}

// GetLocalizationTarget returns value of LocalizationTarget field.
func (g *GetLanguagePackStringRequest) GetLocalizationTarget() (value string) {
	return g.LocalizationTarget
}

// GetLanguagePackID returns value of LanguagePackID field.
func (g *GetLanguagePackStringRequest) GetLanguagePackID() (value string) {
	return g.LanguagePackID
}

// GetKey returns value of Key field.
func (g *GetLanguagePackStringRequest) GetKey() (value string) {
	return g.Key
}

// GetLanguagePackString invokes method getLanguagePackString#8fcde73 returning error if any.
func (c *Client) GetLanguagePackString(ctx context.Context, request *GetLanguagePackStringRequest) (LanguagePackStringValueClass, error) {
	var result LanguagePackStringValueBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LanguagePackStringValue, nil
}
