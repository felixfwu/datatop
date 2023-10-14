package pkg

import (
	"errors"
	"testing"
)

type mockToper struct {
	isErr    bool
	mockData []interface{}
}

func (m *mockToper) collect() error {
	if m.isErr {
		return errors.New("mock error")
	}
	return nil
}

func (m *mockToper) Swap(i int, j int) {

}

func (m mockToper) Len() int {
	return len(m.mockData)
}

func (m mockToper) Less(i int, j int) bool {
	return true
}

func (m mockToper) data() []interface{} {
	return m.mockData
}

func TestTop(t *testing.T) {
	testCase := []struct {
		iMock mockToper
		iN    int
	}{
		{iMock: mockToper{isErr: true, mockData: []interface{}{"1", "2", "3"}}, iN: 2},
		{iMock: mockToper{isErr: false, mockData: []interface{}{"1", "2", "3"}}, iN: 2},
		{iMock: mockToper{isErr: false, mockData: []interface{}{"1", "2", "3"}}, iN: 2},
		{iMock: mockToper{isErr: false, mockData: []interface{}{"1", "2", "3"}}, iN: 0},
		{iMock: mockToper{isErr: false, mockData: []interface{}{"1", "2", "3"}}, iN: -1},
		{iMock: mockToper{isErr: false, mockData: []interface{}{"1"}}, iN: 2},
		{iMock: mockToper{isErr: false, mockData: []interface{}{"1"}}, iN: 2},
	}

	for _, tc := range testCase {
		result, err := Top(tc.iN, &tc.iMock)
		if err != nil && tc.iN > 0 && !tc.iMock.isErr {
			t.Errorf("when n>0 and collect return error, Top() should return error: tc=%v", tc)
		}
		if err != nil && tc.iN <= 0 && tc.iMock.isErr {
			t.Errorf("when n<=0 and collect dosen't return error, Top() should return error: tc=%v", tc)
		}
		dl := len(tc.iMock.data())
		rl := len(result)
		if dl <= tc.iN && dl != rl && result != nil {
			t.Errorf("return len error, tc=%v, result=%v", tc, result)
		}
		if dl > tc.iN && rl != tc.iN && result != nil {
			t.Errorf("return len error, tc=%v, result=%v", tc, result)
		}
	}
}
