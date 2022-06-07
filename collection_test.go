package collection

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type CollectionTestSuite struct {
	suite.Suite
	collection    Collection[string]
	newItem       func() *string
	originalSlice []*string
}

func (s *CollectionTestSuite) SetupTest() {
	s.newItem = func() *string {
		b := make([]byte, 10)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		a := string(b)
		return &a
	}
	s.originalSlice = []*string{s.newItem(), s.newItem(), s.newItem()}
	s.collection = New(s.originalSlice)
}

func (s *CollectionTestSuite) TestNew() {
	// don't use s.collection here, the test is not generic.
	a := "a"
	b := "b"
	c := "c"
	col := New([]*string{&a, &b, &c})
	s.NotNil(col)
}

func (s *CollectionTestSuite) TestCollection_Len() {
	s.Equal(3, s.collection.Len())
}

func (s *CollectionTestSuite) TestCollection_Get() {
	res, err := s.collection.Get(0)
	s.Require().NoError(err)
	s.Equal(s.originalSlice[0], res)
}

func (s *CollectionTestSuite) TestCollection_Get_Error() {
	_, err := s.collection.Get(5)
	s.Require().Error(err)
}

func (s *CollectionTestSuite) TestCollection_Set() {
	want := s.newItem()
	s.collection.Set(0, want)
	got, err := s.collection.Get(0)
	s.Require().NoError(err)
	s.Equal(want, got)
	s.Equal(3, s.collection.Len())
}

func (s *CollectionTestSuite) TestCollection_Exists() {
	s.True(s.collection.Exists(0))
	s.False(s.collection.Exists(5))
}

func (s *CollectionTestSuite) TestCollection_Remove() {
	_, err := s.collection.Remove(0)
	s.Require().NoError(err)
	s.Equal(2, s.collection.Len())
}

func (s *CollectionTestSuite) TestCollection_Remove_Error() {
	_, err := s.collection.Remove(5)
	s.Require().Error(err)
}
func (s *CollectionTestSuite) TestCollection_Append() {
	want := s.newItem()
	s.collection.Append(want)
	s.Equal(4, s.collection.Len())
	got, err := s.collection.Get(3)
	s.Require().NoError(err)
	s.Equal(want, got)
}

func (s *CollectionTestSuite) TestCollection_Prepend() {
	want := s.newItem()
	s.collection.Prepend(want)
	s.Equal(4, s.collection.Len())
	got, err := s.collection.Get(0)
	s.Require().NoError(err)
	s.Equal(want, got)
}
func (s *CollectionTestSuite) TestCollection_Insert() {
	want := s.newItem()
	s.collection.Insert(1, want)
	s.Equal(4, s.collection.Len())
	got, err := s.collection.Get(1)
	s.Require().NoError(err)
	s.Equal(want, got)
}

func TestCollectionTestSuiteString(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}
