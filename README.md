# collection

A generic collection type and various functions to work with it.

## Installation

` $ go get github.com/bad33ndj3/collection`


## Usage

In the following example, we create a collection of `*Item` and apply various functions to it. Its important to note that the collection is generic and can be used with any type, however the type should implement the `Collectable` interface.

```go
package main

import (
	"fmt"

	"github.com/bad33ndj3/collection"
)

// Item a struct that implements the Comparable[T] interface
type Item struct {
	value int
	uid   string
}

func (i *Item) Compare(other *Item) int {
	if i.value < other.value {
		return -1
	}
	if i.value > other.value {
		return 1
	}
	return 0
}

func (i *Item) Number() int {
	return i.value
}

func (i *Item) UID() string {
	return i.uid
}
func (i *Item) IsNil() bool {
	return i == nil
}

func main() {
	// Create a new collection
	data := []*Item{
		{value: 1, uid: "a"},
		{value: 2, uid: "b"},
		{value: 3, uid: "c"},
		{value: 4, uid: "d"},
	}
	col := collection.New[int, *Item](data)

	// Average of the elements
	avg := col.Avg()

	// Check if a value is present in the collection
	contains := col.Contains(3)

	// Iterate over the elements and apply a function
	col.Each(func(i *Item) {
		fmt.Println(i)
	})

	// Check if a function returns true for all elements
	every := col.Every(func(val *Item) bool {
		return val.value > 0
	})

	// Filter elements that pass a function
	filtered := col.Filter(func(val *Item) bool {
		return val.value%2 == 0
	})

	// Get the first element in the collection
	first := col.First()

	// and more...
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
