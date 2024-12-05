package immutableslice

/*/ Reference /*/
// https://www.prequel.co/blog/go-generics-writing-an-immutable-slice-in-go

import (
	"encoding/json"

	"github.com/barkimedes/go-deepcopy"
)

type ImmutableSlice[T any] struct {
	data []T
}

func NewImmutableSlice[T any](values []T) ImmutableSlice[T] {
	// Copy data to ensure immutability
	dataCopy, _ := deepcopy.Anything(values)
	return ImmutableSlice[T]{
		data: dataCopy.([]T),
	}
}

func (s ImmutableSlice[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(s.data) {
		var emptyVal T
		return emptyVal, false
	}

	val, _ := deepcopy.Anything(s.data[index])
	return val.(T), true
}

func (s ImmutableSlice[T]) Len() int {
	return len(s.data)
}

func (s ImmutableSlice[T]) Items() []T {
	dataCopy, _ := deepcopy.Anything(s.data)
	return dataCopy.([]T)
}

func (s *ImmutableSlice[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.data)
}

func (s ImmutableSlice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.data)
}