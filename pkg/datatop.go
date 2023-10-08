package pkg

import (
	"errors"
	"sort"
)

type Toper interface {
	collect() error
	Swap(i int, j int)
	Len() int
	Less(i, j int) bool
	data() []interface{}
}

func Top(n int, t Toper) ([]interface{}, error) {
	if n <= 0 {
		return nil, errors.New("invalid number")
	}

	if err := t.collect(); err != nil {
		return nil, err
	}
	sort.Sort(sort.Reverse(t))
	l := t.Len()
	if l < n {
		return t.data()[:l], nil
	}
	return t.data()[:l], nil
}
