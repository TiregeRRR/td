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

// AccountGetPrivacyRequest represents TL type `account.getPrivacy#dadbc950`.
// Get privacy settings of current account
//
// See https://core.telegram.org/method/account.getPrivacy for reference.
type AccountGetPrivacyRequest struct {
	// Peer category whose privacy settings should be fetched
	Key InputPrivacyKeyClass
}

// AccountGetPrivacyRequestTypeID is TL type id of AccountGetPrivacyRequest.
const AccountGetPrivacyRequestTypeID = 0xdadbc950

// Ensuring interfaces in compile-time for AccountGetPrivacyRequest.
var (
	_ bin.Encoder     = &AccountGetPrivacyRequest{}
	_ bin.Decoder     = &AccountGetPrivacyRequest{}
	_ bin.BareEncoder = &AccountGetPrivacyRequest{}
	_ bin.BareDecoder = &AccountGetPrivacyRequest{}
)

func (g *AccountGetPrivacyRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Key == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetPrivacyRequest) String() string {
	if g == nil {
		return "AccountGetPrivacyRequest(nil)"
	}
	type Alias AccountGetPrivacyRequest
	return fmt.Sprintf("AccountGetPrivacyRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetPrivacyRequest from given interface.
func (g *AccountGetPrivacyRequest) FillFrom(from interface {
	GetKey() (value InputPrivacyKeyClass)
}) {
	g.Key = from.GetKey()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetPrivacyRequest) TypeID() uint32 {
	return AccountGetPrivacyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetPrivacyRequest) TypeName() string {
	return "account.getPrivacy"
}

// TypeInfo returns info about TL type.
func (g *AccountGetPrivacyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getPrivacy",
		ID:   AccountGetPrivacyRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Key",
			SchemaName: "key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetPrivacyRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPrivacy#dadbc950 as nil")
	}
	b.PutID(AccountGetPrivacyRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetPrivacyRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPrivacy#dadbc950 as nil")
	}
	if g.Key == nil {
		return fmt.Errorf("unable to encode account.getPrivacy#dadbc950: field key is nil")
	}
	if err := g.Key.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getPrivacy#dadbc950: field key: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetPrivacyRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPrivacy#dadbc950 to nil")
	}
	if err := b.ConsumeID(AccountGetPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPrivacy#dadbc950: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetPrivacyRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPrivacy#dadbc950 to nil")
	}
	{
		value, err := DecodeInputPrivacyKey(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getPrivacy#dadbc950: field key: %w", err)
		}
		g.Key = value
	}
	return nil
}

// GetKey returns value of Key field.
func (g *AccountGetPrivacyRequest) GetKey() (value InputPrivacyKeyClass) {
	return g.Key
}

// AccountGetPrivacy invokes method account.getPrivacy#dadbc950 returning error if any.
// Get privacy settings of current account
//
// Possible errors:
//  400 PRIVACY_KEY_INVALID: The privacy key is invalid
//
// See https://core.telegram.org/method/account.getPrivacy for reference.
func (c *Client) AccountGetPrivacy(ctx context.Context, key InputPrivacyKeyClass) (*AccountPrivacyRules, error) {
	var result AccountPrivacyRules

	request := &AccountGetPrivacyRequest{
		Key: key,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
