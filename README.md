# Collection

A Go package that provides a generic collection type and various functions to manipulate it. The collection can be used with any type that implements the `Collectable` interface.

## Usage

In the following example, we create a collection of `*Item` and apply various functions to it. Its important to note that the collection is generic and can be used with any type, however the type should implement the `Collectable` interface.

```go
package main

import (
	"fmt"

	"github.com/bad33ndj3/collection"
)

type Item struct {
	value int
}

func main() {
	data := []*Item{
		{value: 1},
		{value: 2},
		{value: 3},
		{value: 4},
	}
	col := collection.New[int, *Item](data)

	avg := col.Avg()

	contains := col.Contains(3)

	col.Each(func(i *Item) {
		fmt.Println(i)
	})

	filtered := col.Filter(func(val *Item) bool {
		return val.value%2 == 0
	})

	first := col.First()

	// and more...
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

Also if you find any bugs or have any suggestions, please open an issue.

## License
[MIT](https://choosealicense.com/licenses/mit/)
