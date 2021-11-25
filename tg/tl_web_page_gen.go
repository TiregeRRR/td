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

// WebPageEmpty represents TL type `webPageEmpty#eb1477e8`.
// No preview is available for the webpage
//
// See https://core.telegram.org/constructor/webPageEmpty for reference.
type WebPageEmpty struct {
	// Preview ID
	ID int64
}

// WebPageEmptyTypeID is TL type id of WebPageEmpty.
const WebPageEmptyTypeID = 0xeb1477e8

// construct implements constructor of WebPageClass.
func (w WebPageEmpty) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPageEmpty.
var (
	_ bin.Encoder     = &WebPageEmpty{}
	_ bin.Decoder     = &WebPageEmpty{}
	_ bin.BareEncoder = &WebPageEmpty{}
	_ bin.BareDecoder = &WebPageEmpty{}

	_ WebPageClass = &WebPageEmpty{}
)

func (w *WebPageEmpty) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageEmpty) String() string {
	if w == nil {
		return "WebPageEmpty(nil)"
	}
	type Alias WebPageEmpty
	return fmt.Sprintf("WebPageEmpty%+v", Alias(*w))
}

// FillFrom fills WebPageEmpty from given interface.
func (w *WebPageEmpty) FillFrom(from interface {
	GetID() (value int64)
}) {
	w.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageEmpty) TypeID() uint32 {
	return WebPageEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageEmpty) TypeName() string {
	return "webPageEmpty"
}

// TypeInfo returns info about TL type.
func (w *WebPageEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageEmpty",
		ID:   WebPageEmptyTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPageEmpty) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageEmpty#eb1477e8 as nil")
	}
	b.PutID(WebPageEmptyTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageEmpty) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageEmpty#eb1477e8 as nil")
	}
	b.PutLong(w.ID)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageEmpty) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageEmpty#eb1477e8 to nil")
	}
	if err := b.ConsumeID(WebPageEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageEmpty#eb1477e8: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageEmpty) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageEmpty#eb1477e8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webPageEmpty#eb1477e8: field id: %w", err)
		}
		w.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (w *WebPageEmpty) GetID() (value int64) {
	return w.ID
}

// WebPagePending represents TL type `webPagePending#c586da1c`.
// A preview of the webpage is currently being generated
//
// See https://core.telegram.org/constructor/webPagePending for reference.
type WebPagePending struct {
	// ID of preview
	ID int64
	// When was the processing started
	Date int
}

// WebPagePendingTypeID is TL type id of WebPagePending.
const WebPagePendingTypeID = 0xc586da1c

// construct implements constructor of WebPageClass.
func (w WebPagePending) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPagePending.
var (
	_ bin.Encoder     = &WebPagePending{}
	_ bin.Decoder     = &WebPagePending{}
	_ bin.BareEncoder = &WebPagePending{}
	_ bin.BareDecoder = &WebPagePending{}

	_ WebPageClass = &WebPagePending{}
)

func (w *WebPagePending) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.ID == 0) {
		return false
	}
	if !(w.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPagePending) String() string {
	if w == nil {
		return "WebPagePending(nil)"
	}
	type Alias WebPagePending
	return fmt.Sprintf("WebPagePending%+v", Alias(*w))
}

// FillFrom fills WebPagePending from given interface.
func (w *WebPagePending) FillFrom(from interface {
	GetID() (value int64)
	GetDate() (value int)
}) {
	w.ID = from.GetID()
	w.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPagePending) TypeID() uint32 {
	return WebPagePendingTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPagePending) TypeName() string {
	return "webPagePending"
}

// TypeInfo returns info about TL type.
func (w *WebPagePending) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPagePending",
		ID:   WebPagePendingTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPagePending) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPagePending#c586da1c as nil")
	}
	b.PutID(WebPagePendingTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPagePending) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPagePending#c586da1c as nil")
	}
	b.PutLong(w.ID)
	b.PutInt(w.Date)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPagePending) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPagePending#c586da1c to nil")
	}
	if err := b.ConsumeID(WebPagePendingTypeID); err != nil {
		return fmt.Errorf("unable to decode webPagePending#c586da1c: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPagePending) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPagePending#c586da1c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webPagePending#c586da1c: field id: %w", err)
		}
		w.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPagePending#c586da1c: field date: %w", err)
		}
		w.Date = value
	}
	return nil
}

// GetID returns value of ID field.
func (w *WebPagePending) GetID() (value int64) {
	return w.ID
}

// GetDate returns value of Date field.
func (w *WebPagePending) GetDate() (value int) {
	return w.Date
}

// WebPage represents TL type `webPage#e89c45b2`.
// Webpage preview
//
// See https://core.telegram.org/constructor/webPage for reference.
type WebPage struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Preview ID
	ID int64
	// URL of previewed webpage
	URL string
	// Webpage URL to be displayed to the user
	DisplayURL string
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Type of the web page. Can be: article, photo, audio, video, document, profile, app, or
	// something else
	//
	// Use SetType and GetType helpers.
	Type string
	// Short name of the site (e.g., Google Docs, App Store)
	//
	// Use SetSiteName and GetSiteName helpers.
	SiteName string
	// Title of the content
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Content description
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// Image representing the content
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo PhotoClass
	// URL to show in the embedded preview
	//
	// Use SetEmbedURL and GetEmbedURL helpers.
	EmbedURL string
	// MIME type of the embedded preview, (e.g., text/html or video/mp4)
	//
	// Use SetEmbedType and GetEmbedType helpers.
	EmbedType string
	// Width of the embedded preview
	//
	// Use SetEmbedWidth and GetEmbedWidth helpers.
	EmbedWidth int
	// Height of the embedded preview
	//
	// Use SetEmbedHeight and GetEmbedHeight helpers.
	EmbedHeight int
	// Duration of the content, in seconds
	//
	// Use SetDuration and GetDuration helpers.
	Duration int
	// Author of the content
	//
	// Use SetAuthor and GetAuthor helpers.
	Author string
	// Preview of the content as a media file
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Page contents in instant view¹ format
	//
	// Links:
	//  1) https://instantview.telegram.org
	//
	// Use SetCachedPage and GetCachedPage helpers.
	CachedPage Page
	// Webpage attributes
	//
	// Use SetAttributes and GetAttributes helpers.
	Attributes []WebPageAttributeTheme
}

// WebPageTypeID is TL type id of WebPage.
const WebPageTypeID = 0xe89c45b2

// construct implements constructor of WebPageClass.
func (w WebPage) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPage.
var (
	_ bin.Encoder     = &WebPage{}
	_ bin.Decoder     = &WebPage{}
	_ bin.BareEncoder = &WebPage{}
	_ bin.BareDecoder = &WebPage{}

	_ WebPageClass = &WebPage{}
)

func (w *WebPage) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.ID == 0) {
		return false
	}
	if !(w.URL == "") {
		return false
	}
	if !(w.DisplayURL == "") {
		return false
	}
	if !(w.Hash == 0) {
		return false
	}
	if !(w.Type == "") {
		return false
	}
	if !(w.SiteName == "") {
		return false
	}
	if !(w.Title == "") {
		return false
	}
	if !(w.Description == "") {
		return false
	}
	if !(w.Photo == nil) {
		return false
	}
	if !(w.EmbedURL == "") {
		return false
	}
	if !(w.EmbedType == "") {
		return false
	}
	if !(w.EmbedWidth == 0) {
		return false
	}
	if !(w.EmbedHeight == 0) {
		return false
	}
	if !(w.Duration == 0) {
		return false
	}
	if !(w.Author == "") {
		return false
	}
	if !(w.Document == nil) {
		return false
	}
	if !(w.CachedPage.Zero()) {
		return false
	}
	if !(w.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPage) String() string {
	if w == nil {
		return "WebPage(nil)"
	}
	type Alias WebPage
	return fmt.Sprintf("WebPage%+v", Alias(*w))
}

// FillFrom fills WebPage from given interface.
func (w *WebPage) FillFrom(from interface {
	GetID() (value int64)
	GetURL() (value string)
	GetDisplayURL() (value string)
	GetHash() (value int)
	GetType() (value string, ok bool)
	GetSiteName() (value string, ok bool)
	GetTitle() (value string, ok bool)
	GetDescription() (value string, ok bool)
	GetPhoto() (value PhotoClass, ok bool)
	GetEmbedURL() (value string, ok bool)
	GetEmbedType() (value string, ok bool)
	GetEmbedWidth() (value int, ok bool)
	GetEmbedHeight() (value int, ok bool)
	GetDuration() (value int, ok bool)
	GetAuthor() (value string, ok bool)
	GetDocument() (value DocumentClass, ok bool)
	GetCachedPage() (value Page, ok bool)
	GetAttributes() (value []WebPageAttributeTheme, ok bool)
}) {
	w.ID = from.GetID()
	w.URL = from.GetURL()
	w.DisplayURL = from.GetDisplayURL()
	w.Hash = from.GetHash()
	if val, ok := from.GetType(); ok {
		w.Type = val
	}

	if val, ok := from.GetSiteName(); ok {
		w.SiteName = val
	}

	if val, ok := from.GetTitle(); ok {
		w.Title = val
	}

	if val, ok := from.GetDescription(); ok {
		w.Description = val
	}

	if val, ok := from.GetPhoto(); ok {
		w.Photo = val
	}

	if val, ok := from.GetEmbedURL(); ok {
		w.EmbedURL = val
	}

	if val, ok := from.GetEmbedType(); ok {
		w.EmbedType = val
	}

	if val, ok := from.GetEmbedWidth(); ok {
		w.EmbedWidth = val
	}

	if val, ok := from.GetEmbedHeight(); ok {
		w.EmbedHeight = val
	}

	if val, ok := from.GetDuration(); ok {
		w.Duration = val
	}

	if val, ok := from.GetAuthor(); ok {
		w.Author = val
	}

	if val, ok := from.GetDocument(); ok {
		w.Document = val
	}

	if val, ok := from.GetCachedPage(); ok {
		w.CachedPage = val
	}

	if val, ok := from.GetAttributes(); ok {
		w.Attributes = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPage) TypeID() uint32 {
	return WebPageTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPage) TypeName() string {
	return "webPage"
}

// TypeInfo returns info about TL type.
func (w *WebPage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPage",
		ID:   WebPageTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "DisplayURL",
			SchemaName: "display_url",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Type",
			SchemaName: "type",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "SiteName",
			SchemaName: "site_name",
			Null:       !w.Flags.Has(1),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !w.Flags.Has(2),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !w.Flags.Has(3),
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
			Null:       !w.Flags.Has(4),
		},
		{
			Name:       "EmbedURL",
			SchemaName: "embed_url",
			Null:       !w.Flags.Has(5),
		},
		{
			Name:       "EmbedType",
			SchemaName: "embed_type",
			Null:       !w.Flags.Has(5),
		},
		{
			Name:       "EmbedWidth",
			SchemaName: "embed_width",
			Null:       !w.Flags.Has(6),
		},
		{
			Name:       "EmbedHeight",
			SchemaName: "embed_height",
			Null:       !w.Flags.Has(6),
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
			Null:       !w.Flags.Has(7),
		},
		{
			Name:       "Author",
			SchemaName: "author",
			Null:       !w.Flags.Has(8),
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !w.Flags.Has(9),
		},
		{
			Name:       "CachedPage",
			SchemaName: "cached_page",
			Null:       !w.Flags.Has(10),
		},
		{
			Name:       "Attributes",
			SchemaName: "attributes",
			Null:       !w.Flags.Has(12),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPage) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPage#e89c45b2 as nil")
	}
	b.PutID(WebPageTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPage) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPage#e89c45b2 as nil")
	}
	if !(w.Type == "") {
		w.Flags.Set(0)
	}
	if !(w.SiteName == "") {
		w.Flags.Set(1)
	}
	if !(w.Title == "") {
		w.Flags.Set(2)
	}
	if !(w.Description == "") {
		w.Flags.Set(3)
	}
	if !(w.Photo == nil) {
		w.Flags.Set(4)
	}
	if !(w.EmbedURL == "") {
		w.Flags.Set(5)
	}
	if !(w.EmbedType == "") {
		w.Flags.Set(5)
	}
	if !(w.EmbedWidth == 0) {
		w.Flags.Set(6)
	}
	if !(w.EmbedHeight == 0) {
		w.Flags.Set(6)
	}
	if !(w.Duration == 0) {
		w.Flags.Set(7)
	}
	if !(w.Author == "") {
		w.Flags.Set(8)
	}
	if !(w.Document == nil) {
		w.Flags.Set(9)
	}
	if !(w.CachedPage.Zero()) {
		w.Flags.Set(10)
	}
	if !(w.Attributes == nil) {
		w.Flags.Set(12)
	}
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#e89c45b2: field flags: %w", err)
	}
	b.PutLong(w.ID)
	b.PutString(w.URL)
	b.PutString(w.DisplayURL)
	b.PutInt(w.Hash)
	if w.Flags.Has(0) {
		b.PutString(w.Type)
	}
	if w.Flags.Has(1) {
		b.PutString(w.SiteName)
	}
	if w.Flags.Has(2) {
		b.PutString(w.Title)
	}
	if w.Flags.Has(3) {
		b.PutString(w.Description)
	}
	if w.Flags.Has(4) {
		if w.Photo == nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field photo is nil")
		}
		if err := w.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field photo: %w", err)
		}
	}
	if w.Flags.Has(5) {
		b.PutString(w.EmbedURL)
	}
	if w.Flags.Has(5) {
		b.PutString(w.EmbedType)
	}
	if w.Flags.Has(6) {
		b.PutInt(w.EmbedWidth)
	}
	if w.Flags.Has(6) {
		b.PutInt(w.EmbedHeight)
	}
	if w.Flags.Has(7) {
		b.PutInt(w.Duration)
	}
	if w.Flags.Has(8) {
		b.PutString(w.Author)
	}
	if w.Flags.Has(9) {
		if w.Document == nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field document is nil")
		}
		if err := w.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field document: %w", err)
		}
	}
	if w.Flags.Has(10) {
		if err := w.CachedPage.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field cached_page: %w", err)
		}
	}
	if w.Flags.Has(12) {
		b.PutVectorHeader(len(w.Attributes))
		for idx, v := range w.Attributes {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode webPage#e89c45b2: field attributes element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPage) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPage#e89c45b2 to nil")
	}
	if err := b.ConsumeID(WebPageTypeID); err != nil {
		return fmt.Errorf("unable to decode webPage#e89c45b2: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPage) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPage#e89c45b2 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field id: %w", err)
		}
		w.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field display_url: %w", err)
		}
		w.DisplayURL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field hash: %w", err)
		}
		w.Hash = value
	}
	if w.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field type: %w", err)
		}
		w.Type = value
	}
	if w.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field site_name: %w", err)
		}
		w.SiteName = value
	}
	if w.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field title: %w", err)
		}
		w.Title = value
	}
	if w.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field description: %w", err)
		}
		w.Description = value
	}
	if w.Flags.Has(4) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field photo: %w", err)
		}
		w.Photo = value
	}
	if w.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_url: %w", err)
		}
		w.EmbedURL = value
	}
	if w.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_type: %w", err)
		}
		w.EmbedType = value
	}
	if w.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_width: %w", err)
		}
		w.EmbedWidth = value
	}
	if w.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_height: %w", err)
		}
		w.EmbedHeight = value
	}
	if w.Flags.Has(7) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field duration: %w", err)
		}
		w.Duration = value
	}
	if w.Flags.Has(8) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field author: %w", err)
		}
		w.Author = value
	}
	if w.Flags.Has(9) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field document: %w", err)
		}
		w.Document = value
	}
	if w.Flags.Has(10) {
		if err := w.CachedPage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field cached_page: %w", err)
		}
	}
	if w.Flags.Has(12) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field attributes: %w", err)
		}

		if headerLen > 0 {
			w.Attributes = make([]WebPageAttributeTheme, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value WebPageAttributeTheme
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode webPage#e89c45b2: field attributes: %w", err)
			}
			w.Attributes = append(w.Attributes, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (w *WebPage) GetID() (value int64) {
	return w.ID
}

// GetURL returns value of URL field.
func (w *WebPage) GetURL() (value string) {
	return w.URL
}

// GetDisplayURL returns value of DisplayURL field.
func (w *WebPage) GetDisplayURL() (value string) {
	return w.DisplayURL
}

// GetHash returns value of Hash field.
func (w *WebPage) GetHash() (value int) {
	return w.Hash
}

// SetType sets value of Type conditional field.
func (w *WebPage) SetType(value string) {
	w.Flags.Set(0)
	w.Type = value
}

// GetType returns value of Type conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetType() (value string, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.Type, true
}

// SetSiteName sets value of SiteName conditional field.
func (w *WebPage) SetSiteName(value string) {
	w.Flags.Set(1)
	w.SiteName = value
}

// GetSiteName returns value of SiteName conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetSiteName() (value string, ok bool) {
	if !w.Flags.Has(1) {
		return value, false
	}
	return w.SiteName, true
}

// SetTitle sets value of Title conditional field.
func (w *WebPage) SetTitle(value string) {
	w.Flags.Set(2)
	w.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetTitle() (value string, ok bool) {
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Title, true
}

// SetDescription sets value of Description conditional field.
func (w *WebPage) SetDescription(value string) {
	w.Flags.Set(3)
	w.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetDescription() (value string, ok bool) {
	if !w.Flags.Has(3) {
		return value, false
	}
	return w.Description, true
}

// SetPhoto sets value of Photo conditional field.
func (w *WebPage) SetPhoto(value PhotoClass) {
	w.Flags.Set(4)
	w.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetPhoto() (value PhotoClass, ok bool) {
	if !w.Flags.Has(4) {
		return value, false
	}
	return w.Photo, true
}

// SetEmbedURL sets value of EmbedURL conditional field.
func (w *WebPage) SetEmbedURL(value string) {
	w.Flags.Set(5)
	w.EmbedURL = value
}

// GetEmbedURL returns value of EmbedURL conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedURL() (value string, ok bool) {
	if !w.Flags.Has(5) {
		return value, false
	}
	return w.EmbedURL, true
}

// SetEmbedType sets value of EmbedType conditional field.
func (w *WebPage) SetEmbedType(value string) {
	w.Flags.Set(5)
	w.EmbedType = value
}

// GetEmbedType returns value of EmbedType conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedType() (value string, ok bool) {
	if !w.Flags.Has(5) {
		return value, false
	}
	return w.EmbedType, true
}

// SetEmbedWidth sets value of EmbedWidth conditional field.
func (w *WebPage) SetEmbedWidth(value int) {
	w.Flags.Set(6)
	w.EmbedWidth = value
}

// GetEmbedWidth returns value of EmbedWidth conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedWidth() (value int, ok bool) {
	if !w.Flags.Has(6) {
		return value, false
	}
	return w.EmbedWidth, true
}

// SetEmbedHeight sets value of EmbedHeight conditional field.
func (w *WebPage) SetEmbedHeight(value int) {
	w.Flags.Set(6)
	w.EmbedHeight = value
}

// GetEmbedHeight returns value of EmbedHeight conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedHeight() (value int, ok bool) {
	if !w.Flags.Has(6) {
		return value, false
	}
	return w.EmbedHeight, true
}

// SetDuration sets value of Duration conditional field.
func (w *WebPage) SetDuration(value int) {
	w.Flags.Set(7)
	w.Duration = value
}

// GetDuration returns value of Duration conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetDuration() (value int, ok bool) {
	if !w.Flags.Has(7) {
		return value, false
	}
	return w.Duration, true
}

// SetAuthor sets value of Author conditional field.
func (w *WebPage) SetAuthor(value string) {
	w.Flags.Set(8)
	w.Author = value
}

// GetAuthor returns value of Author conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetAuthor() (value string, ok bool) {
	if !w.Flags.Has(8) {
		return value, false
	}
	return w.Author, true
}

// SetDocument sets value of Document conditional field.
func (w *WebPage) SetDocument(value DocumentClass) {
	w.Flags.Set(9)
	w.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetDocument() (value DocumentClass, ok bool) {
	if !w.Flags.Has(9) {
		return value, false
	}
	return w.Document, true
}

// SetCachedPage sets value of CachedPage conditional field.
func (w *WebPage) SetCachedPage(value Page) {
	w.Flags.Set(10)
	w.CachedPage = value
}

// GetCachedPage returns value of CachedPage conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetCachedPage() (value Page, ok bool) {
	if !w.Flags.Has(10) {
		return value, false
	}
	return w.CachedPage, true
}

// SetAttributes sets value of Attributes conditional field.
func (w *WebPage) SetAttributes(value []WebPageAttributeTheme) {
	w.Flags.Set(12)
	w.Attributes = value
}

// GetAttributes returns value of Attributes conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetAttributes() (value []WebPageAttributeTheme, ok bool) {
	if !w.Flags.Has(12) {
		return value, false
	}
	return w.Attributes, true
}

// WebPageNotModified represents TL type `webPageNotModified#7311ca11`.
// The preview of the webpage hasn't changed
//
// See https://core.telegram.org/constructor/webPageNotModified for reference.
type WebPageNotModified struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Page view count
	//
	// Use SetCachedPageViews and GetCachedPageViews helpers.
	CachedPageViews int
}

// WebPageNotModifiedTypeID is TL type id of WebPageNotModified.
const WebPageNotModifiedTypeID = 0x7311ca11

// construct implements constructor of WebPageClass.
func (w WebPageNotModified) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPageNotModified.
var (
	_ bin.Encoder     = &WebPageNotModified{}
	_ bin.Decoder     = &WebPageNotModified{}
	_ bin.BareEncoder = &WebPageNotModified{}
	_ bin.BareDecoder = &WebPageNotModified{}

	_ WebPageClass = &WebPageNotModified{}
)

func (w *WebPageNotModified) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.CachedPageViews == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageNotModified) String() string {
	if w == nil {
		return "WebPageNotModified(nil)"
	}
	type Alias WebPageNotModified
	return fmt.Sprintf("WebPageNotModified%+v", Alias(*w))
}

// FillFrom fills WebPageNotModified from given interface.
func (w *WebPageNotModified) FillFrom(from interface {
	GetCachedPageViews() (value int, ok bool)
}) {
	if val, ok := from.GetCachedPageViews(); ok {
		w.CachedPageViews = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageNotModified) TypeID() uint32 {
	return WebPageNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageNotModified) TypeName() string {
	return "webPageNotModified"
}

// TypeInfo returns info about TL type.
func (w *WebPageNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageNotModified",
		ID:   WebPageNotModifiedTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CachedPageViews",
			SchemaName: "cached_page_views",
			Null:       !w.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPageNotModified) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageNotModified#7311ca11 as nil")
	}
	b.PutID(WebPageNotModifiedTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageNotModified) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageNotModified#7311ca11 as nil")
	}
	if !(w.CachedPageViews == 0) {
		w.Flags.Set(0)
	}
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageNotModified#7311ca11: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutInt(w.CachedPageViews)
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageNotModified) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageNotModified#7311ca11 to nil")
	}
	if err := b.ConsumeID(WebPageNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageNotModified#7311ca11: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageNotModified) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageNotModified#7311ca11 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageNotModified#7311ca11: field flags: %w", err)
		}
	}
	if w.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPageNotModified#7311ca11: field cached_page_views: %w", err)
		}
		w.CachedPageViews = value
	}
	return nil
}

// SetCachedPageViews sets value of CachedPageViews conditional field.
func (w *WebPageNotModified) SetCachedPageViews(value int) {
	w.Flags.Set(0)
	w.CachedPageViews = value
}

// GetCachedPageViews returns value of CachedPageViews conditional field and
// boolean which is true if field was set.
func (w *WebPageNotModified) GetCachedPageViews() (value int, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.CachedPageViews, true
}

// WebPageClass represents WebPage generic type.
//
// See https://core.telegram.org/type/WebPage for reference.
//
// Example:
//  g, err := tg.DecodeWebPage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.WebPageEmpty: // webPageEmpty#eb1477e8
//  case *tg.WebPagePending: // webPagePending#c586da1c
//  case *tg.WebPage: // webPage#e89c45b2
//  case *tg.WebPageNotModified: // webPageNotModified#7311ca11
//  default: panic(v)
//  }
type WebPageClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() WebPageClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map WebPageClass to ModifiedWebPage.
	AsModified() (ModifiedWebPage, bool)
}

// ModifiedWebPage represents Modified subset of WebPageClass.
type ModifiedWebPage interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() WebPageClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Preview ID
	GetID() (value int64)
}

// AsModified tries to map WebPageEmpty to ModifiedWebPage.
func (w *WebPageEmpty) AsModified() (ModifiedWebPage, bool) {
	value, ok := (WebPageClass(w)).(ModifiedWebPage)
	return value, ok
}

// AsModified tries to map WebPagePending to ModifiedWebPage.
func (w *WebPagePending) AsModified() (ModifiedWebPage, bool) {
	value, ok := (WebPageClass(w)).(ModifiedWebPage)
	return value, ok
}

// AsModified tries to map WebPage to ModifiedWebPage.
func (w *WebPage) AsModified() (ModifiedWebPage, bool) {
	value, ok := (WebPageClass(w)).(ModifiedWebPage)
	return value, ok
}

// AsModified tries to map WebPageNotModified to ModifiedWebPage.
func (w *WebPageNotModified) AsModified() (ModifiedWebPage, bool) {
	value, ok := (WebPageClass(w)).(ModifiedWebPage)
	return value, ok
}

// DecodeWebPage implements binary de-serialization for WebPageClass.
func DecodeWebPage(buf *bin.Buffer) (WebPageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WebPageEmptyTypeID:
		// Decoding webPageEmpty#eb1477e8.
		v := WebPageEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	case WebPagePendingTypeID:
		// Decoding webPagePending#c586da1c.
		v := WebPagePending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	case WebPageTypeID:
		// Decoding webPage#e89c45b2.
		v := WebPage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	case WebPageNotModifiedTypeID:
		// Decoding webPageNotModified#7311ca11.
		v := WebPageNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WebPageClass: %w", bin.NewUnexpectedID(id))
	}
}

// WebPage boxes the WebPageClass providing a helper.
type WebPageBox struct {
	WebPage WebPageClass
}

// Decode implements bin.Decoder for WebPageBox.
func (b *WebPageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WebPageBox to nil")
	}
	v, err := DecodeWebPage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WebPage = v
	return nil
}

// Encode implements bin.Encode for WebPageBox.
func (b *WebPageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WebPage == nil {
		return fmt.Errorf("unable to encode WebPageClass as nil")
	}
	return b.WebPage.Encode(buf)
}
