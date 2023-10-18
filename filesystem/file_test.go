package filesystem

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	testCase := struct {
		iDir []File
		eDir []File
	}{
		iDir: []File{{Path: "/root", FileCount: 10}, {Path: "/usr", FileCount: 20}},
		eDir: []File{{Path: "/usr", FileCount: 20}, {Path: "/root", FileCount: 10}},
	}
	f := FileToper{Dirs: testCase.iDir}
	i, j := 0, 1
	f.Swap(i, j)
	if f.Dirs[i] != testCase.eDir[i] || f.Dirs[j] != testCase.eDir[j] {
		t.Errorf("filesystem.file Swap() error: expect %v, acture %v", testCase.eDir, f.Dirs)
	}
}

func TestLen(t *testing.T) {
	testCase := []FileToper{
		{Root: "/home"},
		{Root: "/usr", Dirs: []File{{Path: "/tmp", FileCount: 10}}},
	}
	for _, tc := range testCase {
		if tc.Len() != len(tc.Dirs) {
			t.Errorf("filesystem.file Len() error: expect %d, acture %d", len(tc.Dirs), tc.Len())
		}
	}
}

func TestLess(t *testing.T) {
	mockFile := FileToper{Root: "/usr", Dirs: []File{{Path: "/tmp", FileCount: 10}, {Path: "root", FileCount: 20}}}
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
	testCase := []struct {
		n      int
		expect int
	}{
		{n: 10, expect: 10},
		{n: 1, expect: 1},
	}
	mockToper := FileToper{Dirs: make([]File, 10)}

	for _, tc := range testCase {
		l := len(mockToper.Data(tc.n).([]File))
		if l != tc.expect {
			t.Errorf("testcase error: %v", tc)
		}
	}
}

func TestCollect(t *testing.T) {
	testCase := []struct {
		path   string // input path
		eDirs  []File //expect return []Dir
		eIsErr bool   // should return error or not
	}{
		{path: "", eDirs: []File{}, eIsErr: true},
		{path: "./notexists", eDirs: []File{}, eIsErr: true},
		{path: ".", eDirs: append([]File{{Path: ".", FileCount: 2}}, File{}), eIsErr: false},
		{path: "./", eDirs: append([]File{{Path: "./", FileCount: 2}}, File{}), eIsErr: false},
		{path: "./.", eDirs: append([]File{{Path: "./.", FileCount: 2}}, File{}), eIsErr: false},
		{path: "../filesystem/", eDirs: append([]File{{Path: "../filesystem/", FileCount: 2}}, File{}), eIsErr: false},
		{path: "../cmd", eDirs: append([]File{{Path: "../cmd", FileCount: 4}}, File{}), eIsErr: false},
	}
	for _, tc := range testCase {
		f := FileToper{Root: tc.path}
		err := f.Collect()
		if reflect.DeepEqual(f.Dirs, tc.eDirs) || (err != nil) != tc.eIsErr {
			t.Errorf("testcase=%v failed: %v=Collect() f=%v", tc, err, f)
		}
	}
	f := FileToper{Root: "."}
	err := f.Collect()
	if err != nil {
		t.Errorf("Collect error: %s", err)
	}
	if len(f.Dirs) != 1 {
		t.Errorf("Collect len error: expect=1 acture=%v", f.Dirs)
	}
}

// use the code directory to test.
//
// ps. actually i don't know how to test is better when there is os.open, filepath.walk etc
func TestWalkCurr(t *testing.T) {
	testCase := []struct {
		path   string // input path
		eDirs  []File //expect return []Dir
		eIsErr bool   // should return error or not
	}{
		{path: "", eDirs: []File{}, eIsErr: true},
		{path: "./notexists", eDirs: []File{}, eIsErr: true},
		{path: ".", eDirs: append([]File{{Path: ".", FileCount: 2}}, File{}), eIsErr: false},
		{path: "./", eDirs: append([]File{{Path: "./", FileCount: 2}}, File{}), eIsErr: false},
		{path: "./.", eDirs: append([]File{{Path: "./.", FileCount: 2}}, File{}), eIsErr: false},
		{path: "../filesystem/", eDirs: append([]File{{Path: "../filesystem/", FileCount: 2}}, File{}), eIsErr: false},
		{path: "../cmd", eDirs: append([]File{{Path: "../cmd", FileCount: 4}}, File{}), eIsErr: false},
	}
	for _, tc := range testCase {
		ds, err := walkCurr(tc.path)
		if reflect.DeepEqual(ds, tc.eDirs) || (err != nil) != tc.eIsErr {
			t.Errorf("testcase=%v failed: %v, %v=walkCurr(%s)", tc, ds, err, tc.path)
		}
	}
}

// use the code directory to test.
//
// ps. actually i don't know how to test is better when there is os.open, filepath.walk etc
func TestWalkChild(t *testing.T) {
	testCase := []struct {
		path   string // input path
		eCnt   int    //expect count
		eIsErr bool   // should return error or not
	}{
		{path: "", eCnt: 0, eIsErr: true},
		{path: "./notexists", eCnt: 0, eIsErr: true},
		{path: ".", eCnt: 2, eIsErr: false},
		{path: "./", eCnt: 2, eIsErr: false},
		{path: "./.", eCnt: 2, eIsErr: false},
		{path: "../filesystem/", eCnt: 2, eIsErr: false},
	}
	for _, tc := range testCase {
		cnt, err := walkChild(tc.path)
		if cnt != tc.eCnt || (err != nil) != tc.eIsErr {
			t.Errorf("testcase=%v failed: %d, %v=walkChild(%s)", tc, cnt, err, tc.path)
		}
	}
}
