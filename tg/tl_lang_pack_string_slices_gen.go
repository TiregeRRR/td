//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// LangPackStringClassArray is adapter for slice of LangPackStringClass.
type LangPackStringClassArray []LangPackStringClass

// Sort sorts slice of LangPackStringClass.
func (s LangPackStringClassArray) Sort(less func(a, b LangPackStringClass) bool) LangPackStringClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackStringClass.
func (s LangPackStringClassArray) SortStable(less func(a, b LangPackStringClass) bool) LangPackStringClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackStringClass.
func (s LangPackStringClassArray) Retain(keep func(x LangPackStringClass) bool) LangPackStringClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s LangPackStringClassArray) First() (v LangPackStringClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringClassArray) Last() (v LangPackStringClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringClassArray) PopFirst() (v LangPackStringClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackStringClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringClassArray) Pop() (v LangPackStringClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsLangPackString returns copy with only LangPackString constructors.
func (s LangPackStringClassArray) AsLangPackString() (to LangPackStringArray) {
	for _, elem := range s {
		value, ok := elem.(*LangPackString)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsLangPackStringPluralized returns copy with only LangPackStringPluralized constructors.
func (s LangPackStringClassArray) AsLangPackStringPluralized() (to LangPackStringPluralizedArray) {
	for _, elem := range s {
		value, ok := elem.(*LangPackStringPluralized)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsLangPackStringDeleted returns copy with only LangPackStringDeleted constructors.
func (s LangPackStringClassArray) AsLangPackStringDeleted() (to LangPackStringDeletedArray) {
	for _, elem := range s {
		value, ok := elem.(*LangPackStringDeleted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// LangPackStringArray is adapter for slice of LangPackString.
type LangPackStringArray []LangPackString

// Sort sorts slice of LangPackString.
func (s LangPackStringArray) Sort(less func(a, b LangPackString) bool) LangPackStringArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackString.
func (s LangPackStringArray) SortStable(less func(a, b LangPackString) bool) LangPackStringArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackString.
func (s LangPackStringArray) Retain(keep func(x LangPackString) bool) LangPackStringArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s LangPackStringArray) First() (v LangPackString, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringArray) Last() (v LangPackString, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringArray) PopFirst() (v LangPackString, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackString
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringArray) Pop() (v LangPackString, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// LangPackStringPluralizedArray is adapter for slice of LangPackStringPluralized.
type LangPackStringPluralizedArray []LangPackStringPluralized

// Sort sorts slice of LangPackStringPluralized.
func (s LangPackStringPluralizedArray) Sort(less func(a, b LangPackStringPluralized) bool) LangPackStringPluralizedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackStringPluralized.
func (s LangPackStringPluralizedArray) SortStable(less func(a, b LangPackStringPluralized) bool) LangPackStringPluralizedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackStringPluralized.
func (s LangPackStringPluralizedArray) Retain(keep func(x LangPackStringPluralized) bool) LangPackStringPluralizedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s LangPackStringPluralizedArray) First() (v LangPackStringPluralized, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringPluralizedArray) Last() (v LangPackStringPluralized, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringPluralizedArray) PopFirst() (v LangPackStringPluralized, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackStringPluralized
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringPluralizedArray) Pop() (v LangPackStringPluralized, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// LangPackStringDeletedArray is adapter for slice of LangPackStringDeleted.
type LangPackStringDeletedArray []LangPackStringDeleted

// Sort sorts slice of LangPackStringDeleted.
func (s LangPackStringDeletedArray) Sort(less func(a, b LangPackStringDeleted) bool) LangPackStringDeletedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackStringDeleted.
func (s LangPackStringDeletedArray) SortStable(less func(a, b LangPackStringDeleted) bool) LangPackStringDeletedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackStringDeleted.
func (s LangPackStringDeletedArray) Retain(keep func(x LangPackStringDeleted) bool) LangPackStringDeletedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s LangPackStringDeletedArray) First() (v LangPackStringDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringDeletedArray) Last() (v LangPackStringDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringDeletedArray) PopFirst() (v LangPackStringDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackStringDeleted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringDeletedArray) Pop() (v LangPackStringDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
