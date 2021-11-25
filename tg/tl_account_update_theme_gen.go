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

// AccountUpdateThemeRequest represents TL type `account.updateTheme#2bf40ccc`.
// Update theme
//
// See https://core.telegram.org/method/account.updateTheme for reference.
type AccountUpdateThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Theme format, a string that identifies the theming engines supported by the client
	Format string
	// Theme to update
	Theme InputThemeClass
	// Unique theme ID
	//
	// Use SetSlug and GetSlug helpers.
	Slug string
	// Theme name
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Theme file
	//
	// Use SetDocument and GetDocument helpers.
	Document InputDocumentClass
	// Theme settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings []InputThemeSettings
}

// AccountUpdateThemeRequestTypeID is TL type id of AccountUpdateThemeRequest.
const AccountUpdateThemeRequestTypeID = 0x2bf40ccc

// Ensuring interfaces in compile-time for AccountUpdateThemeRequest.
var (
	_ bin.Encoder     = &AccountUpdateThemeRequest{}
	_ bin.Decoder     = &AccountUpdateThemeRequest{}
	_ bin.BareEncoder = &AccountUpdateThemeRequest{}
	_ bin.BareDecoder = &AccountUpdateThemeRequest{}
)

func (u *AccountUpdateThemeRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Format == "") {
		return false
	}
	if !(u.Theme == nil) {
		return false
	}
	if !(u.Slug == "") {
		return false
	}
	if !(u.Title == "") {
		return false
	}
	if !(u.Document == nil) {
		return false
	}
	if !(u.Settings == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateThemeRequest) String() string {
	if u == nil {
		return "AccountUpdateThemeRequest(nil)"
	}
	type Alias AccountUpdateThemeRequest
	return fmt.Sprintf("AccountUpdateThemeRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateThemeRequest from given interface.
func (u *AccountUpdateThemeRequest) FillFrom(from interface {
	GetFormat() (value string)
	GetTheme() (value InputThemeClass)
	GetSlug() (value string, ok bool)
	GetTitle() (value string, ok bool)
	GetDocument() (value InputDocumentClass, ok bool)
	GetSettings() (value []InputThemeSettings, ok bool)
}) {
	u.Format = from.GetFormat()
	u.Theme = from.GetTheme()
	if val, ok := from.GetSlug(); ok {
		u.Slug = val
	}

	if val, ok := from.GetTitle(); ok {
		u.Title = val
	}

	if val, ok := from.GetDocument(); ok {
		u.Document = val
	}

	if val, ok := from.GetSettings(); ok {
		u.Settings = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateThemeRequest) TypeID() uint32 {
	return AccountUpdateThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateThemeRequest) TypeName() string {
	return "account.updateTheme"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateTheme",
		ID:   AccountUpdateThemeRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Format",
			SchemaName: "format",
		},
		{
			Name:       "Theme",
			SchemaName: "theme",
		},
		{
			Name:       "Slug",
			SchemaName: "slug",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !u.Flags.Has(3),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateThemeRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateTheme#2bf40ccc as nil")
	}
	b.PutID(AccountUpdateThemeRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateThemeRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateTheme#2bf40ccc as nil")
	}
	if !(u.Slug == "") {
		u.Flags.Set(0)
	}
	if !(u.Title == "") {
		u.Flags.Set(1)
	}
	if !(u.Document == nil) {
		u.Flags.Set(2)
	}
	if !(u.Settings == nil) {
		u.Flags.Set(3)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateTheme#2bf40ccc: field flags: %w", err)
	}
	b.PutString(u.Format)
	if u.Theme == nil {
		return fmt.Errorf("unable to encode account.updateTheme#2bf40ccc: field theme is nil")
	}
	if err := u.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateTheme#2bf40ccc: field theme: %w", err)
	}
	if u.Flags.Has(0) {
		b.PutString(u.Slug)
	}
	if u.Flags.Has(1) {
		b.PutString(u.Title)
	}
	if u.Flags.Has(2) {
		if u.Document == nil {
			return fmt.Errorf("unable to encode account.updateTheme#2bf40ccc: field document is nil")
		}
		if err := u.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.updateTheme#2bf40ccc: field document: %w", err)
		}
	}
	if u.Flags.Has(3) {
		b.PutVectorHeader(len(u.Settings))
		for idx, v := range u.Settings {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode account.updateTheme#2bf40ccc: field settings element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateThemeRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateTheme#2bf40ccc to nil")
	}
	if err := b.ConsumeID(AccountUpdateThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateThemeRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateTheme#2bf40ccc to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field format: %w", err)
		}
		u.Format = value
	}
	{
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field theme: %w", err)
		}
		u.Theme = value
	}
	if u.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field slug: %w", err)
		}
		u.Slug = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field title: %w", err)
		}
		u.Title = value
	}
	if u.Flags.Has(2) {
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field document: %w", err)
		}
		u.Document = value
	}
	if u.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field settings: %w", err)
		}

		if headerLen > 0 {
			u.Settings = make([]InputThemeSettings, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputThemeSettings
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.updateTheme#2bf40ccc: field settings: %w", err)
			}
			u.Settings = append(u.Settings, value)
		}
	}
	return nil
}

// GetFormat returns value of Format field.
func (u *AccountUpdateThemeRequest) GetFormat() (value string) {
	return u.Format
}

// GetTheme returns value of Theme field.
func (u *AccountUpdateThemeRequest) GetTheme() (value InputThemeClass) {
	return u.Theme
}

// SetSlug sets value of Slug conditional field.
func (u *AccountUpdateThemeRequest) SetSlug(value string) {
	u.Flags.Set(0)
	u.Slug = value
}

// GetSlug returns value of Slug conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateThemeRequest) GetSlug() (value string, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.Slug, true
}

// SetTitle sets value of Title conditional field.
func (u *AccountUpdateThemeRequest) SetTitle(value string) {
	u.Flags.Set(1)
	u.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateThemeRequest) GetTitle() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.Title, true
}

// SetDocument sets value of Document conditional field.
func (u *AccountUpdateThemeRequest) SetDocument(value InputDocumentClass) {
	u.Flags.Set(2)
	u.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateThemeRequest) GetDocument() (value InputDocumentClass, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.Document, true
}

// SetSettings sets value of Settings conditional field.
func (u *AccountUpdateThemeRequest) SetSettings(value []InputThemeSettings) {
	u.Flags.Set(3)
	u.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateThemeRequest) GetSettings() (value []InputThemeSettings, ok bool) {
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.Settings, true
}

// GetDocumentAsNotEmpty returns mapped value of Document conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateThemeRequest) GetDocumentAsNotEmpty() (*InputDocument, bool) {
	if value, ok := u.GetDocument(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// AccountUpdateTheme invokes method account.updateTheme#2bf40ccc returning error if any.
// Update theme
//
// See https://core.telegram.org/method/account.updateTheme for reference.
func (c *Client) AccountUpdateTheme(ctx context.Context, request *AccountUpdateThemeRequest) (*Theme, error) {
	var result Theme

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
