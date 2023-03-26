package collection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CollectionTestSuite struct {
	suite.Suite
	col *Collection[int, *Item]
}

func TestCollection(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}

func (s *CollectionTestSuite) SetupTest() {
	// setup test data
	data := []*Item{
		{value: 1, uid: "a"},
		{value: 2, uid: "b"},
		{value: 3, uid: "c"},
		{value: 4, uid: "d"},
	}
	s.col = New[int, *Item](data)
}

func (s *CollectionTestSuite) TestNew() {
	col := New[int, *Item]([]*Item{})
	s.NotNil(col)
	s.Len(col.list, 0)
}

func (s *CollectionTestSuite) TestAverage() {
	avg := s.col.Average()
	s.Equal(2, avg)
}

func (s *CollectionTestSuite) TestAvg() {
	avg := s.col.Avg()
	s.Equal(2, avg)
}

func (s *CollectionTestSuite) TestContains() {
	val := &Item{value: 3, uid: "c"}
	contains := s.col.Contains(val)
	s.True(contains)
	val = &Item{value: 5, uid: "e"}
	contains = s.col.Contains(val)
	s.False(contains)
}

func (s *CollectionTestSuite) TestEach() {
	var sum int
	s.col.Each(func(val *Item) {
		sum += val.value
	})
	s.Equal(10, sum)
}

func (s *CollectionTestSuite) TestEvery() {
	allEven := s.col.Every(func(val *Item) bool {
		return val.value%2 == 0
	})
	s.False(allEven)
	allLessThan5 := s.col.Every(func(val *Item) bool {
		return val.value < 5
	})
	s.True(allLessThan5)
}

func (s *CollectionTestSuite) TestFilter() {
	filtered := s.col.Filter(func(val *Item) bool {
		return val.value%2 == 0
	})
	s.Len(filtered, 2)
	for _, val := range filtered {
		s.Equal(0, val.value%2)
	}
}

func (s *CollectionTestSuite) TestFirst() {
	first := s.col.First()
	s.NotNil(first)
	s.NotNil(*first)
	a := *first
	s.Equal(1, a.value)
	s.col = New[int, *Item]([]*Item{})
	first = s.col.First()
	s.Nil(first)
}

func (s *CollectionTestSuite) TestFlatMap() {
	flatMapped := s.col.FlatMap(func(val *Item) []*Item {
		return []*Item{val, val}
	})
	s.Len(flatMapped, 8)
}

func (s *CollectionTestSuite) TestGroupBy() {
	grouped := s.col.GroupBy(func(val *Item) string {
		return val.uid
	})
	s.Len(grouped, 4)
	for _, v := range grouped {
		s.Len(v, 1)
	}
}

func (s *CollectionTestSuite) TestKeyBy() {
	keyed := s.col.KeyBy(func(val *Item) string {
		return val.uid
	})
	s.Len(keyed, 4)
	for key, val := range keyed {
		s.Equal(val.uid, key)
	}
}

func (s *CollectionTestSuite) TestMap() {
	mapped := s.col.Map(func(val *Item) *Item {
		val.value *= 2

		return val
	})
	s.Len(mapped, 4)
	for _, item := range mapped {
		s.Equal(0, item.value%2)
	}
}

func (s *CollectionTestSuite) TestMax() {
	maxVal := s.col.Max()
	s.Equal(4, maxVal.value)

	col := New[int, *Item]([]*Item{})
	maxVal = col.Max()
	s.Nil(maxVal)
}

func (s *CollectionTestSuite) TestMin() {
	minVal := s.col.Min()
	s.Equal(1, minVal.value)

	col := New[int, *Item]([]*Item{})
	minVal = col.Min()
	s.Nil(minVal)
}

func (s *CollectionTestSuite) TestPartition() {
	partition1, partition2 := s.col.Partition(func(val *Item) bool {
		return val.value%2 == 0
	})
	s.Len(partition1, 2)
	s.Len(partition2, 2)
	for _, val := range partition1 {
		s.Equal(0, val.value%2)
	}
	for _, val := range partition2 {
		s.Equal(1, val.value%2)
	}
}

func (s *CollectionTestSuite) TestReject() {
	rejected := s.col.Reject(func(val *Item) bool {
		return val.value%2 == 0
	})
	s.Len(rejected, 2)
	for _, val := range rejected {
		s.Equal(1, val.value%2)
	}
}

func (s *CollectionTestSuite) TestSkipUntil() {
	skipped := s.col.SkipUntil(func(val *Item) bool {
		return val.uid == "c"
	})
	s.Len(skipped, 2)
	for _, val := range skipped {
		s.True(val.value > 2)
	}
}

func (s *CollectionTestSuite) TestSkipWhile() {
	skipped := s.col.SkipWhile(func(val *Item) bool {
		return val.value < 3
	})
	s.Len(skipped, 2)
	for _, val := range skipped {
		s.True(val.value > 2)
	}
}

func (s *CollectionTestSuite) TestSome() {
	hasEven := s.col.Some(func(val *Item) bool {
		return val.value%2 == 0
	})
	s.True(hasEven)
	hasOver4 := s.col.Some(func(val *Item) bool {
		return val.value > 4
	})
	s.False(hasOver4)
}

func (s *CollectionTestSuite) TestSort() {
	sorted := s.col.Sort(func(val *Item) int {
		return val.value
	}, false)
	s.Equal(4, sorted.Len())
	for i := 1; i < sorted.Len(); i++ {
		s.True(sorted.Nth(i-1).value <= sorted.Nth(i).value)
	}
}

func (s *CollectionTestSuite) TestSortByDesc() {
	sorted := s.col.Sort(func(val *Item) int {
		return val.value
	}, true)
	s.Equal(sorted.Len(), 4)
	for i := 1; i < sorted.Len(); i++ {
		s.True(sorted.Nth(i-1).value >= sorted.Nth(i).value)
	}
}

func (s *CollectionTestSuite) TestSum() {
	sum := s.col.Sum()
	s.Equal(10, sum)
}

func (s *CollectionTestSuite) TestTakeUntil() {
	taken := s.col.TakeUntil(func(val *Item) bool {
		return val.uid == "c"
	})
	s.Len(taken, 2)
	for _, val := range taken {
		s.True(val.value < 3)
	}
}

func (s *CollectionTestSuite) TestNth() {
	item := s.col.Nth(0)
	s.Equal(1, item.value)
}

func (s *CollectionTestSuite) TestLen() {
	s.Equal(4, s.col.Len())
}

func (s *CollectionTestSuite) TestTakeWhile() {
	data := []*Item{
		{value: 1, uid: "a"},
		{value: 2, uid: "b"},
		{value: 3, uid: "c"},
		{value: 4, uid: "d"},
		{value: 5, uid: "e"},
	}
	col := New[int, *Item](data)
	takeWhile := col.TakeWhile(func(val *Item) bool {
		return val.value < 4
	})
	s.Len(takeWhile, 3)
	for _, val := range takeWhile {
		s.True(val.value < 4)
	}
}

func (s *CollectionTestSuite) TestUnique() {
	data := []*Item{
		{value: 1, uid: "a"},
		{value: 2, uid: "b"},
		{value: 3, uid: "c"},
		{value: 2, uid: "d"},
		{value: 1, uid: "e"},
	}
	col := New[int, *Item](data)
	unique := col.Unique()
	s.Require().Len(unique, 1)
	s.Equal(3, unique[0].value)
}

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
