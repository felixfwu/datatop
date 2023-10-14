package datatop

import (
	"errors"
	"testing"
)

type mockToper struct {
	isErr    bool
	mockData []int
}

func (m mockToper) Collect() error {
	if m.isErr {
		return errors.New("mock error")
	}
	return nil
}

func (m mockToper) Swap(i int, j int) {

}

func (m mockToper) Len() int {
	return len(m.mockData)
}

func (m mockToper) Less(i int, j int) bool {
	return true
}

func (m mockToper) Data(n int) interface{} {
	return m.mockData[:n]
}

func TestTop(t *testing.T) {
	testCase1 := []int{-1, 0}
	for _, tc := range testCase1 {
		if result, _ := Top(tc, &mockToper{}); result != nil {
			t.Error("n<= should return nil result and a error")
		}
	}

	testCase2 := mockToper{isErr: true}
	if result, _ := Top(2, &testCase2); result != nil {
		t.Error("collect error should return nil result and a error")
	}

	m := mockToper{isErr: false, mockData: []int{1, 2, 3, 4}}
	testCase3 := []struct {
		n    int
		eLen int
	}{
		{2, 2},
		{4, 4},
		{5, 4},
	}
	for _, tc := range testCase3 {
		result, _ := Top(tc.n, &m)
		r, _ := result.([]int)
		if len(r) != tc.eLen {
			t.Errorf("result len error: len(result)=%d, len(data)=%d", len(r), tc.eLen)
		}
	}
}
