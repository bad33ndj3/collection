# collection

A generic collection type and various functions to work with it.

## Installation

` $ go get github.com/bad33ndj3/collection`


## Usage

In the following example, we create a collection of `*Item` and apply various functions to it. Its important to note that the collection is generic and can be used with any type, however the type should implement the `Collectable` interface.

```go
import "github.com/bad33ndj3/collection"

// Create a new collection
data := []*Item{
    {value: 1, uid: "a"},
    {value: 2, uid: "b"},
    {value: 3, uid: "c"},
    {value: 4, uid: "d"},
}
col := New[int, *Item](data)

// Average of the elements
avg := col.Avg()

// Check if a value is present in the collection
contains := col.Contains(3)

// Iterate over the elements and apply a function
col.Each(func(val int) {
    fmt.Println(val)
})

// Check if a function returns true for all elements
every := col.Every(func(val int) bool {
    return val > 0
})

// Filter elements that pass a function
filtered := col.Filter(func(val int) bool {
    return val%2 == 0
})

// Get the first element in the collection
first := col.First()

// FlatMap applies a function to each element and flattens the resulting slices
flatMapped := col.FlatMap(func(val int) []int {
    return []int{val, val + 1}
})

// GroupBy groups the elements by the key returned by a function
grouped := col.GroupBy(func(val int) int {
    return val % 2
})

// Max element in the collection
max := col.Max()

// Min element in the collection
min := col.Min()

// Partition elements that pass or fail a function
pass, fail := col.Partition(func(val int) bool {
    return val > 2
})

// Reject elements that pass a function
rejected := col.Reject(func(val int) bool {
    return val > 2
})

// Select elements that pass a function
selected := col.Select(func(val int) bool {
    return val > 2
})

// Sum of the elements
sum := col.Sum()

// TakeWhile takes elements while a function is true
taken := col.TakeWhile(func(val int) bool {
    return val < 3
})
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
