// Code generated by gen, DO NOT EDIT.

package photos

import (
	"context"

	"github.com/gotd/td/tg"
)

// No-op definition for keeping imports.
var _ = context.Background()

// Request is a parameter for Query.
type Request struct {
	Offset int
	Limit  int
}

// Query is a abstraction for photos request.
// NB: iterator mutates returned data (sorts, at least).
type Query interface {
	Query(ctx context.Context, req Request) (tg.PhotosPhotosClass, error)
}

// QueryFunc is a function adapter for Query.
type QueryFunc func(ctx context.Context, req Request) (tg.PhotosPhotosClass, error)

// Query implements Query interface.
func (q QueryFunc) Query(ctx context.Context, req Request) (tg.PhotosPhotosClass, error) {
	return q(ctx, req)
}

// QueryBuilder is a helper to create message queries.
type QueryBuilder struct {
	raw *tg.Client
}

// NewQueryBuilder creates new QueryBuilder.
func NewQueryBuilder(raw *tg.Client) *QueryBuilder {
	return &QueryBuilder{raw: raw}
}

// GetUserPhotosQueryBuilder is query builder of PhotosGetUserPhotos.
type GetUserPhotosQueryBuilder struct {
	raw       *tg.Client
	req       tg.PhotosGetUserPhotosRequest
	batchSize int
	offset    int
}

// GetUserPhotos creates query builder of PhotosGetUserPhotos.
func (q *QueryBuilder) GetUserPhotos(paramUserID tg.InputUserClass) *GetUserPhotosQueryBuilder {
	b := &GetUserPhotosQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req:       tg.PhotosGetUserPhotosRequest{},
	}

	b.req.UserID = paramUserID
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *GetUserPhotosQueryBuilder) BatchSize(batchSize int) *GetUserPhotosQueryBuilder {
	b.batchSize = batchSize
	return b
}

// UserID sets UserID field of GetUserPhotos query.
func (b *GetUserPhotosQueryBuilder) UserID(paramUserID tg.InputUserClass) *GetUserPhotosQueryBuilder {
	b.req.UserID = paramUserID
	return b
}

// Query implements Query interface.
func (b *GetUserPhotosQueryBuilder) Query(ctx context.Context, req Request) (tg.PhotosPhotosClass, error) {
	r := &tg.PhotosGetUserPhotosRequest{
		Limit: req.Limit,
	}

	r.UserID = b.req.UserID
	r.Offset = req.Offset
	return b.raw.PhotosGetUserPhotos(ctx, r)
}

// Iter returns iterator using built query.
func (b *GetUserPhotosQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.Offset(b.offset)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *GetUserPhotosQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}
