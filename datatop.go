package datatop

import (
	"errors"
	"sort"
)

type Toper interface {
	Collect() error
	Swap(i int, j int)
	Len() int
	Less(i int, j int) bool
	Data(n int) interface{}
}

func Top(n int, t Toper) (interface{}, error) {
	if n <= 0 {
		return nil, errors.New("invalid number")
	}

	if err := t.Collect(); err != nil {
		return nil, err
	}
	sort.Sort(sort.Reverse(t))
	l := t.Len()
	if l < n {
		return t.Data(l), nil
	}
	return t.Data(n), nil
}
