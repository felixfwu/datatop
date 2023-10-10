package filesystem

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	testCase := struct {
		iDir []Dir
		eDir []Dir
	}{
		iDir: []Dir{{Name: "/root", Cnt: 10}, {Name: "/usr", Cnt: 20}},
		eDir: []Dir{{Name: "/usr", Cnt: 20}, {Name: "/root", Cnt: 10}},
	}
	f := File{Dirs: testCase.iDir}
	i, j := 0, 1
	f.Swap(i, j)
	if f.Dirs[i] != testCase.eDir[i] || f.Dirs[j] != testCase.eDir[j] {
		t.Errorf("filesystem.file Swap() error: expect %v, acture %v", testCase.eDir, f.Dirs)
	}
}

func TestLen(t *testing.T) {
	testCase := []File{
		{Root: "/home"},
		{Root: "/usr", Dirs: []Dir{{Name: "/tmp", Cnt: 10}}},
	}
	for _, tc := range testCase {
		if tc.Len() != len(tc.Dirs) {
			t.Errorf("filesystem.file Len() error: expect %d, acture %d", len(tc.Dirs), tc.Len())
		}
	}
}

func TestLess(t *testing.T) {
	mockFile := File{Root: "/usr", Dirs: []Dir{{Name: "/tmp", Cnt: 10}, {Name: "root", Cnt: 20}}}
	testCase := []struct {
		ii      int
		ij      int
		eResult bool
	}{
		{0, 1, true},
		{1, 0, false},
	}
	for _, tc := range testCase {
		if mockFile.Less(tc.ii, tc.ij) != tc.eResult {
			t.Errorf("filesystem.file Less(%d, %d) error: mockFile=%v", tc.ii, tc.ij, mockFile)
		}

	}
}

func TestData(t *testing.T) {
	testCase := File{Root: "/usr", Dirs: []Dir{{Name: "/tmp", Cnt: 10}}}

	r := testCase.data(1)
	if nrt, ok := r.([]Dir); !ok {
		t.Errorf("type error %t", reflect.TypeOf(nrt))
	}
}

func TestIsDir(t *testing.T) {
	testCase := []struct {
		file   string
		expect bool
	}{
		{".", true},
		{"./", true},
		{"../filesystem", true},
		{"./noteexists", false},
		{"file.go", false},
		{"./file.go", false},
	}
	for _, tc := range testCase {
		if isDir(tc.file) != tc.expect {
			t.Errorf("isDir error: %s expect %t", tc.file, tc.expect)
		}
	}
}
