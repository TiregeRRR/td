// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// ReceivedNotifyMessageVector is a box for Vector<ReceivedNotifyMessage>
type ReceivedNotifyMessageVector struct {
	// Elements of Vector<ReceivedNotifyMessage>
	Elems []ReceivedNotifyMessage
}

// Encode implements bin.Encoder.
func (vec *ReceivedNotifyMessageVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<ReceivedNotifyMessage> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<ReceivedNotifyMessage>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *ReceivedNotifyMessageVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<ReceivedNotifyMessage> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<ReceivedNotifyMessage>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ReceivedNotifyMessage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<ReceivedNotifyMessage>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ReceivedNotifyMessageVector.
var (
	_ bin.Encoder = &ReceivedNotifyMessageVector{}
	_ bin.Decoder = &ReceivedNotifyMessageVector{}
)