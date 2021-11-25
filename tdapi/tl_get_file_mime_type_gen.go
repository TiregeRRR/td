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

// GetFileMimeTypeRequest represents TL type `getFileMimeType#84631b89`.
type GetFileMimeTypeRequest struct {
	// The name of the file or path to the file
	FileName string
}

// GetFileMimeTypeRequestTypeID is TL type id of GetFileMimeTypeRequest.
const GetFileMimeTypeRequestTypeID = 0x84631b89

// Ensuring interfaces in compile-time for GetFileMimeTypeRequest.
var (
	_ bin.Encoder     = &GetFileMimeTypeRequest{}
	_ bin.Decoder     = &GetFileMimeTypeRequest{}
	_ bin.BareEncoder = &GetFileMimeTypeRequest{}
	_ bin.BareDecoder = &GetFileMimeTypeRequest{}
)

func (g *GetFileMimeTypeRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FileName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetFileMimeTypeRequest) String() string {
	if g == nil {
		return "GetFileMimeTypeRequest(nil)"
	}
	type Alias GetFileMimeTypeRequest
	return fmt.Sprintf("GetFileMimeTypeRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetFileMimeTypeRequest) TypeID() uint32 {
	return GetFileMimeTypeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetFileMimeTypeRequest) TypeName() string {
	return "getFileMimeType"
}

// TypeInfo returns info about TL type.
func (g *GetFileMimeTypeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getFileMimeType",
		ID:   GetFileMimeTypeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileName",
			SchemaName: "file_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetFileMimeTypeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getFileMimeType#84631b89 as nil")
	}
	b.PutID(GetFileMimeTypeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetFileMimeTypeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getFileMimeType#84631b89 as nil")
	}
	b.PutString(g.FileName)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetFileMimeTypeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getFileMimeType#84631b89 to nil")
	}
	if err := b.ConsumeID(GetFileMimeTypeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getFileMimeType#84631b89: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetFileMimeTypeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getFileMimeType#84631b89 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getFileMimeType#84631b89: field file_name: %w", err)
		}
		g.FileName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetFileMimeTypeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getFileMimeType#84631b89 as nil")
	}
	b.ObjStart()
	b.PutID("getFileMimeType")
	b.FieldStart("file_name")
	b.PutString(g.FileName)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetFileMimeTypeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getFileMimeType#84631b89 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getFileMimeType"); err != nil {
				return fmt.Errorf("unable to decode getFileMimeType#84631b89: %w", err)
			}
		case "file_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getFileMimeType#84631b89: field file_name: %w", err)
			}
			g.FileName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileName returns value of FileName field.
func (g *GetFileMimeTypeRequest) GetFileName() (value string) {
	return g.FileName
}

// GetFileMimeType invokes method getFileMimeType#84631b89 returning error if any.
func (c *Client) GetFileMimeType(ctx context.Context, filename string) (*Text, error) {
	var result Text

	request := &GetFileMimeTypeRequest{
		FileName: filename,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
