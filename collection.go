// Package collection provides a generic collection type and various functions to work with it.
package collection

import (
	"sort"
)

// Number is a generic type that holds all available number types.
// This is used for methods such as Sum, Max and Avg.
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Collectable is an interface that combines the Comparable interface with the necessary methods for a collection.
type Collectable[T any, N Number] interface {
	// Compare compares the current value to another value of the same type.
	// It returns a negative number if the current value is less than the other value.
	// It returns a positive number if the current value is greater than the other value.
	// It returns zero if the current value is equal to the other value.
	Compare(other T) int
	// Number returns a number representation of the value.
	Number() N
	// IsNil returns true if the value is nil.
	IsNil() bool
}

// Collection is a generic type that holds a slice of values of type T.
type Collection[N Number, T Collectable[T, N]] struct {
	// list holds the slice of values
	list []T
}

// New returns a new Collection of type T.
func New[N Number, T Collectable[T, N]](list []T) *Collection[N, T] {
	return &Collection[N, T]{
		list: list,
	}
}

// Average calculates the average of the elements in the collection.
func (c *Collection[N, T]) Average() N {
	var sum N
	for _, val := range c.list {
		sum += val.Number()
	}
	return sum / N(len(c.list))
}

// Avg is an alias for Average.
func (c *Collection[N, T]) Avg() N {
	return c.Average()
}

// Contains checks if the given value is present in the collection.
func (c *Collection[N, T]) Contains(val T) bool {
	for _, item := range c.list {
		if item.Compare(val) == 0 {
			return true
		}
	}
	return false
}

// Each iterates over the elements in the collection and applies the given function to each element.
func (c *Collection[N, T]) Each(fn func(T)) *Collection[N, T] {
	for _, val := range c.list {
		fn(val)
	}

	return c
}

// Every checks if the given function returns true for all elements in the collection.
func (c *Collection[N, T]) Every(fn func(T) bool) bool {
	for _, val := range c.list {
		if !fn(val) {
			return false
		}
	}
	return true
}

// Filter returns a new slice with elements that pass the given function.
func (c *Collection[N, T]) Filter(fn func(T) bool) []T {
	var filteredList []T
	expectedLength := 0
	for _, val := range c.list {
		if fn(val) {
			expectedLength++
		}
	}
	filteredList = make([]T, 0, expectedLength)
	for _, val := range c.list {
		if fn(val) {
			filteredList = append(filteredList, val)
		}
	}
	return filteredList
}

// First returns the first element in the collection.
func (c *Collection[N, T]) First() T {
	if len(c.list) > 0 {
		return c.list[0]
	}
	return *new(T)
}

// FlatMap applies the given function to each element in the collection and flattens the resulting slices.
func (c *Collection[N, T]) FlatMap(fn func(T) []T) []T {
	var flatList []T
	for _, val := range c.list {
		flatList = append(flatList, fn(val)...)
	}
	return flatList
}

// GroupBy groups the elements in the collection by the key returned by the given function.
func (c *Collection[N, T]) GroupBy(fn func(T) string) map[string][]T {
	grouped := make(map[string][]T)
	for _, val := range c.list {
		key := fn(val)
		grouped[key] = append(grouped[key], val)
	}
	return grouped
}

// KeyBy creates a map of elements in the collection with keys returned by the given function.
func (c *Collection[N, T]) KeyBy(fn func(T) string) map[string]T {
	keyed := make(map[string]T)
	for _, val := range c.list {
		key := fn(val)
		keyed[key] = val
	}
	return keyed
}

// Map applies the given function to each element in the collection and returns a new slice.
func (c *Collection[N, T]) Map(fn func(T) T) []T {
	var mappedList []T
	for _, val := range c.list {
		mappedList = append(mappedList, fn(val))
	}
	return mappedList
}

// Max returns the maximum element in the collection.
func (c *Collection[N, T]) Max() T {
	if len(c.list) == 0 {
		return *new(T)
	}
	var max T
	for _, val := range c.list {
		if max.IsNil() || val.Number() > max.Number() {
			max = val
		}
	}
	return max
}

// Min returns the minimum element in the collection.
func (c *Collection[N, T]) Min() T {
	if len(c.list) == 0 {
		return *new(T)
	}
	var min T
	for _, val := range c.list {
		if min.IsNil() || val.Number() < min.Number() {
			min = val
		}
	}
	return min
}

// Nth returns the element at the given index.
func (c *Collection[N, T]) Nth(index int) T {
	return c.list[index]
}

// Partition splits the elements in the collection into two slices based on the given function.
func (c *Collection[N, T]) Partition(fn func(T) bool) (passed []T, failed []T) {
	for _, val := range c.list {
		if fn(val) {
			passed = append(passed, val)
		} else {
			failed = append(failed, val)
		}
	}
	return
}

// Reject returns a new slice with elements that do not pass the given function.
func (c *Collection[N, T]) Reject(fn func(T) bool) []T {
	var rejectedList []T
	for _, val := range c.list {
		if !fn(val) {
			rejectedList = append(rejectedList, val)
		}
	}
	return rejectedList
}

// SkipUntil returns elements from the collection until the given function returns true.
func (c *Collection[N, T]) SkipUntil(fn func(T) bool) []T {
	var skippedList []T
	skip := true
	for _, val := range c.list {
		if skip && !fn(val) {
			continue
		}
		skip = false
		skippedList = append(skippedList, val)
	}
	return skippedList
}

// SkipWhile returns elements from the collection while the given function returns true.
func (c *Collection[N, T]) SkipWhile(fn func(T) bool) []T {
	var skippedList []T
	skip := true
	for _, val := range c.list {
		if skip && fn(val) {
			continue
		}
		skip = false
		skippedList = append(skippedList, val)
	}
	return skippedList
}

// Some checks if the given function returns true for any element in the collection.
func (c *Collection[N, T]) Some(fn func(T) bool) bool {
	for _, val := range c.list {
		if fn(val) {
			return true
		}
	}
	return false
}

// Sort sorts the elements in the collection by the value returned by the given function.
// The order can be specified with the asc or desc parameter.
func (c *Collection[N, T]) Sort(fn func(T) int, desc bool) *Collection[N, T] {
	if desc {
		sort.Slice(c.list, func(i, j int) bool {
			return fn(c.list[i]) > fn(c.list[j])
		})
	} else {
		sort.Slice(c.list, func(i, j int) bool {
			return fn(c.list[i]) < fn(c.list[j])
		})
	}

	return c
}

// Sum returns the sum of the elements in the collection.
func (c *Collection[N, T]) Sum() N {
	var sum N
	for _, val := range c.list {
		sum += val.Number()
	}
	return sum
}

// TakeUntil returns elements from the collection until the given function returns true.
func (c *Collection[N, T]) TakeUntil(fn func(T) bool) []T {
	var takenList []T
	for _, val := range c.list {
		if fn(val) {
			break
		}
		takenList = append(takenList, val)
	}
	return takenList
}

// TakeWhile returns elements from the collection while the given function returns true.
func (c *Collection[N, T]) TakeWhile(fn func(T) bool) []T {
	var takenList []T
	for _, val := range c.list {
		if !fn(val) {
			break
		}
		takenList = append(takenList, val)
	}
	return takenList
}

// Unique returns a new slice with unique elements in the collection by filtering out elements that are not unique.
// Uniqueness is determined by using the Compare method of the element type.
func (c *Collection[N, T]) Unique() []T {
	var uniqueList []T
	for _, item := range c.list {
		unique := true
		for _, otherItem := range uniqueList {
			if item.Compare(otherItem) == 0 {
				unique = false
				break
			}
		}
		if unique {
			uniqueList = append(uniqueList, item)
		}
	}
	return uniqueList
}

// Len returns the length of the collection.
func (c *Collection[N, T]) Len() int {
	return len(c.list)
}
