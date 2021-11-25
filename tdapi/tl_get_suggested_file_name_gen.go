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

// GetSuggestedFileNameRequest represents TL type `getSuggestedFileName#85d8a486`.
type GetSuggestedFileNameRequest struct {
	// Identifier of the file
	FileID int32
	// Directory in which the file is supposed to be saved
	Directory string
}

// GetSuggestedFileNameRequestTypeID is TL type id of GetSuggestedFileNameRequest.
const GetSuggestedFileNameRequestTypeID = 0x85d8a486

// Ensuring interfaces in compile-time for GetSuggestedFileNameRequest.
var (
	_ bin.Encoder     = &GetSuggestedFileNameRequest{}
	_ bin.Decoder     = &GetSuggestedFileNameRequest{}
	_ bin.BareEncoder = &GetSuggestedFileNameRequest{}
	_ bin.BareDecoder = &GetSuggestedFileNameRequest{}
)

func (g *GetSuggestedFileNameRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FileID == 0) {
		return false
	}
	if !(g.Directory == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSuggestedFileNameRequest) String() string {
	if g == nil {
		return "GetSuggestedFileNameRequest(nil)"
	}
	type Alias GetSuggestedFileNameRequest
	return fmt.Sprintf("GetSuggestedFileNameRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSuggestedFileNameRequest) TypeID() uint32 {
	return GetSuggestedFileNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSuggestedFileNameRequest) TypeName() string {
	return "getSuggestedFileName"
}

// TypeInfo returns info about TL type.
func (g *GetSuggestedFileNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSuggestedFileName",
		ID:   GetSuggestedFileNameRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "Directory",
			SchemaName: "directory",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSuggestedFileNameRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSuggestedFileName#85d8a486 as nil")
	}
	b.PutID(GetSuggestedFileNameRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSuggestedFileNameRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSuggestedFileName#85d8a486 as nil")
	}
	b.PutInt32(g.FileID)
	b.PutString(g.Directory)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSuggestedFileNameRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSuggestedFileName#85d8a486 to nil")
	}
	if err := b.ConsumeID(GetSuggestedFileNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSuggestedFileName#85d8a486: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSuggestedFileNameRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSuggestedFileName#85d8a486 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSuggestedFileName#85d8a486: field file_id: %w", err)
		}
		g.FileID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getSuggestedFileName#85d8a486: field directory: %w", err)
		}
		g.Directory = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSuggestedFileNameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSuggestedFileName#85d8a486 as nil")
	}
	b.ObjStart()
	b.PutID("getSuggestedFileName")
	b.FieldStart("file_id")
	b.PutInt32(g.FileID)
	b.FieldStart("directory")
	b.PutString(g.Directory)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSuggestedFileNameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSuggestedFileName#85d8a486 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSuggestedFileName"); err != nil {
				return fmt.Errorf("unable to decode getSuggestedFileName#85d8a486: %w", err)
			}
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSuggestedFileName#85d8a486: field file_id: %w", err)
			}
			g.FileID = value
		case "directory":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getSuggestedFileName#85d8a486: field directory: %w", err)
			}
			g.Directory = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileID returns value of FileID field.
func (g *GetSuggestedFileNameRequest) GetFileID() (value int32) {
	return g.FileID
}

// GetDirectory returns value of Directory field.
func (g *GetSuggestedFileNameRequest) GetDirectory() (value string) {
	return g.Directory
}

// GetSuggestedFileName invokes method getSuggestedFileName#85d8a486 returning error if any.
func (c *Client) GetSuggestedFileName(ctx context.Context, request *GetSuggestedFileNameRequest) (*Text, error) {
	var result Text

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
