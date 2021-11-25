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

// ResendEmailAddressVerificationCodeRequest represents TL type `resendEmailAddressVerificationCode#90653024`.
type ResendEmailAddressVerificationCodeRequest struct {
}

// ResendEmailAddressVerificationCodeRequestTypeID is TL type id of ResendEmailAddressVerificationCodeRequest.
const ResendEmailAddressVerificationCodeRequestTypeID = 0x90653024

// Ensuring interfaces in compile-time for ResendEmailAddressVerificationCodeRequest.
var (
	_ bin.Encoder     = &ResendEmailAddressVerificationCodeRequest{}
	_ bin.Decoder     = &ResendEmailAddressVerificationCodeRequest{}
	_ bin.BareEncoder = &ResendEmailAddressVerificationCodeRequest{}
	_ bin.BareDecoder = &ResendEmailAddressVerificationCodeRequest{}
)

func (r *ResendEmailAddressVerificationCodeRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResendEmailAddressVerificationCodeRequest) String() string {
	if r == nil {
		return "ResendEmailAddressVerificationCodeRequest(nil)"
	}
	type Alias ResendEmailAddressVerificationCodeRequest
	return fmt.Sprintf("ResendEmailAddressVerificationCodeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResendEmailAddressVerificationCodeRequest) TypeID() uint32 {
	return ResendEmailAddressVerificationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResendEmailAddressVerificationCodeRequest) TypeName() string {
	return "resendEmailAddressVerificationCode"
}

// TypeInfo returns info about TL type.
func (r *ResendEmailAddressVerificationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resendEmailAddressVerificationCode",
		ID:   ResendEmailAddressVerificationCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResendEmailAddressVerificationCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendEmailAddressVerificationCode#90653024 as nil")
	}
	b.PutID(ResendEmailAddressVerificationCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResendEmailAddressVerificationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendEmailAddressVerificationCode#90653024 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResendEmailAddressVerificationCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendEmailAddressVerificationCode#90653024 to nil")
	}
	if err := b.ConsumeID(ResendEmailAddressVerificationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resendEmailAddressVerificationCode#90653024: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResendEmailAddressVerificationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendEmailAddressVerificationCode#90653024 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResendEmailAddressVerificationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resendEmailAddressVerificationCode#90653024 as nil")
	}
	b.ObjStart()
	b.PutID("resendEmailAddressVerificationCode")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResendEmailAddressVerificationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resendEmailAddressVerificationCode#90653024 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resendEmailAddressVerificationCode"); err != nil {
				return fmt.Errorf("unable to decode resendEmailAddressVerificationCode#90653024: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ResendEmailAddressVerificationCode invokes method resendEmailAddressVerificationCode#90653024 returning error if any.
func (c *Client) ResendEmailAddressVerificationCode(ctx context.Context) (*EmailAddressAuthenticationCodeInfo, error) {
	var result EmailAddressAuthenticationCodeInfo

	request := &ResendEmailAddressVerificationCodeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
