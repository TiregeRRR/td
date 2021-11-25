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

// AccountResetWebAuthorizationsRequest represents TL type `account.resetWebAuthorizations#682d2594`.
// Reset all active web telegram login¹ sessions
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorizations for reference.
type AccountResetWebAuthorizationsRequest struct {
}

// AccountResetWebAuthorizationsRequestTypeID is TL type id of AccountResetWebAuthorizationsRequest.
const AccountResetWebAuthorizationsRequestTypeID = 0x682d2594

// Ensuring interfaces in compile-time for AccountResetWebAuthorizationsRequest.
var (
	_ bin.Encoder     = &AccountResetWebAuthorizationsRequest{}
	_ bin.Decoder     = &AccountResetWebAuthorizationsRequest{}
	_ bin.BareEncoder = &AccountResetWebAuthorizationsRequest{}
	_ bin.BareDecoder = &AccountResetWebAuthorizationsRequest{}
)

func (r *AccountResetWebAuthorizationsRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetWebAuthorizationsRequest) String() string {
	if r == nil {
		return "AccountResetWebAuthorizationsRequest(nil)"
	}
	type Alias AccountResetWebAuthorizationsRequest
	return fmt.Sprintf("AccountResetWebAuthorizationsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetWebAuthorizationsRequest) TypeID() uint32 {
	return AccountResetWebAuthorizationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetWebAuthorizationsRequest) TypeName() string {
	return "account.resetWebAuthorizations"
}

// TypeInfo returns info about TL type.
func (r *AccountResetWebAuthorizationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetWebAuthorizations",
		ID:   AccountResetWebAuthorizationsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetWebAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorizations#682d2594 as nil")
	}
	b.PutID(AccountResetWebAuthorizationsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetWebAuthorizationsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorizations#682d2594 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetWebAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorizations#682d2594 to nil")
	}
	if err := b.ConsumeID(AccountResetWebAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWebAuthorizations#682d2594: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetWebAuthorizationsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorizations#682d2594 to nil")
	}
	return nil
}

// AccountResetWebAuthorizations invokes method account.resetWebAuthorizations#682d2594 returning error if any.
// Reset all active web telegram login¹ sessions
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorizations for reference.
func (c *Client) AccountResetWebAuthorizations(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountResetWebAuthorizationsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
