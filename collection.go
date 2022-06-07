// Package collection implements helper functions for working with a collection of anything.
package collection

import "fmt"

var (
	// ErrIndexOutOfRange is returned when an index is out of range.
	ErrIndexOutOfRange = fmt.Errorf("index out of range")
)

// Collection is a collection of items.
type Collection[T any] interface {
	// Len returns the number of items in the collection.
	Len() int
	// Get returns the item at the given index.
	Get(int) (*T, error)
	// Set sets the item at the given index to the given value.
	Set(int, *T) Collection[T]
	// Exists returns true if the index exists in the collection.
	Exists(int) bool
	// Remove deletes the item at the given index.
	Remove(int) (Collection[T], error)
	// Append adds the given item to the end of the collection.
	Append(*T) Collection[T]
	// Prepend adds the given item to the beginning of the collection.
	Prepend(*T) Collection[T]
	// Insert adds the given item at the given index.
	Insert(int, *T) Collection[T]
}

// New creates a new collection from the given slice.
func New[T any](slice []*T) Collection[T] {
	return &collection[T]{slice: slice}
}

type collection[T any] struct {
	slice []*T
}

// Len returns the length of the collection.
func (c *collection[T]) Len() int {
	return len(c.slice)
}

// Get returns the item at the given index.
func (c *collection[T]) Get(i int) (*T, error) {
	if i < 0 || i >= c.Len() {
		return nil, ErrIndexOutOfRange
	}

	return c.slice[i], nil
}

// Set sets the item at the given index to the given value.
func (c *collection[T]) Set(i int, v *T) Collection[T] {
	c.slice[i] = v

	return c
}

// Exists returns true if the index exists in the collection.
func (c *collection[T]) Exists(i int) bool {
	return i >= 0 && i < c.Len()
}

// Remove deletes the item at the given index.
func (c *collection[T]) Remove(i int) (Collection[T], error) {
	if i < 0 || i >= c.Len() {
		return nil, ErrIndexOutOfRange
	}
	c.slice = append(c.slice[:i], c.slice[i+1:]...)

	return c, nil
}

// Append adds the given item to the end of the collection.
func (c *collection[T]) Append(v *T) Collection[T] {
	c.slice = append(c.slice, v)

	return c
}

// Prepend adds the given item to the beginning of the collection.
func (c *collection[T]) Prepend(v *T) Collection[T] {
	c.slice = append([]*T{v}, c.slice...)

	return c
}

// Insert adds the given item at the given index.
func (c *collection[T]) Insert(i int, v *T) Collection[T] {
	c.slice = append(c.slice[:i], append([]*T{v}, c.slice[i:]...)...)

	return c
}
