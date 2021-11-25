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

// AccountGetPasswordSettingsRequest represents TL type `account.getPasswordSettings#9cd4eaf9`.
// Get private info associated to the password info (recovery email, telegram passport¹
// info & so on)
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.getPasswordSettings for reference.
type AccountGetPasswordSettingsRequest struct {
	// The password (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Password InputCheckPasswordSRPClass
}

// AccountGetPasswordSettingsRequestTypeID is TL type id of AccountGetPasswordSettingsRequest.
const AccountGetPasswordSettingsRequestTypeID = 0x9cd4eaf9

// Ensuring interfaces in compile-time for AccountGetPasswordSettingsRequest.
var (
	_ bin.Encoder     = &AccountGetPasswordSettingsRequest{}
	_ bin.Decoder     = &AccountGetPasswordSettingsRequest{}
	_ bin.BareEncoder = &AccountGetPasswordSettingsRequest{}
	_ bin.BareDecoder = &AccountGetPasswordSettingsRequest{}
)

func (g *AccountGetPasswordSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetPasswordSettingsRequest) String() string {
	if g == nil {
		return "AccountGetPasswordSettingsRequest(nil)"
	}
	type Alias AccountGetPasswordSettingsRequest
	return fmt.Sprintf("AccountGetPasswordSettingsRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetPasswordSettingsRequest from given interface.
func (g *AccountGetPasswordSettingsRequest) FillFrom(from interface {
	GetPassword() (value InputCheckPasswordSRPClass)
}) {
	g.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetPasswordSettingsRequest) TypeID() uint32 {
	return AccountGetPasswordSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetPasswordSettingsRequest) TypeName() string {
	return "account.getPasswordSettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetPasswordSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getPasswordSettings",
		ID:   AccountGetPasswordSettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetPasswordSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPasswordSettings#9cd4eaf9 as nil")
	}
	b.PutID(AccountGetPasswordSettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetPasswordSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPasswordSettings#9cd4eaf9 as nil")
	}
	if g.Password == nil {
		return fmt.Errorf("unable to encode account.getPasswordSettings#9cd4eaf9: field password is nil")
	}
	if err := g.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getPasswordSettings#9cd4eaf9: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetPasswordSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPasswordSettings#9cd4eaf9 to nil")
	}
	if err := b.ConsumeID(AccountGetPasswordSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPasswordSettings#9cd4eaf9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetPasswordSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPasswordSettings#9cd4eaf9 to nil")
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getPasswordSettings#9cd4eaf9: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// GetPassword returns value of Password field.
func (g *AccountGetPasswordSettingsRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	return g.Password
}

// GetPasswordAsNotEmpty returns mapped value of Password field.
func (g *AccountGetPasswordSettingsRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return g.Password.AsNotEmpty()
}

// AccountGetPasswordSettings invokes method account.getPasswordSettings#9cd4eaf9 returning error if any.
// Get private info associated to the password info (recovery email, telegram passport¹
// info & so on)
//
// Links:
//  1) https://core.telegram.org/passport
//
// Possible errors:
//  400 PASSWORD_HASH_INVALID: The provided password hash is invalid
//
// See https://core.telegram.org/method/account.getPasswordSettings for reference.
func (c *Client) AccountGetPasswordSettings(ctx context.Context, password InputCheckPasswordSRPClass) (*AccountPasswordSettings, error) {
	var result AccountPasswordSettings

	request := &AccountGetPasswordSettingsRequest{
		Password: password,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
