package oracle

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSwap(t *testing.T) {
	testCase := struct {
		it TablespaceToper
		et TablespaceToper
	}{
		it: TablespaceToper{
			TS: []Tablespace{{Name: "SYSTEM"}, {Name: "SYSAUX"}},
		},
		et: TablespaceToper{
			TS: []Tablespace{{Name: "SYSAUX"}, {Name: "SYSTEM"}},
		},
	}

	testCase.it.Swap(0, 1)
	if testCase.it.TS[0] != testCase.et.TS[0] || testCase.it.TS[1] != testCase.et.TS[1] {
		t.Errorf("testcase=%v error", testCase)
	}
}

func TestLess(t *testing.T) {
	testCase := []struct {
		it     TablespaceToper
		expect bool
	}{
		{it: TablespaceToper{TS: []Tablespace{{TPCT: 0.9}, {TPCT: 0.99}}}, expect: true},
		{it: TablespaceToper{TS: []Tablespace{{TPCT: 0.9}, {TPCT: 0.9}}}, expect: false},
		{it: TablespaceToper{TS: []Tablespace{{TPCT: 0.9}, {TPCT: 0.89}}}, expect: false},
	}

	for _, tc := range testCase {
		if tc.it.Less(0, 1) != tc.expect {
			t.Errorf("testcase error: %v", tc)
		}
	}
}

func TestLen(t *testing.T) {
	testCase := []struct {
		it     TablespaceToper
		expect int
	}{
		{it: TablespaceToper{TS: make([]Tablespace, 10)}, expect: 10},
		{it: TablespaceToper{TS: make([]Tablespace, 3)}, expect: 3},
		{it: TablespaceToper{TS: make([]Tablespace, 2)}, expect: 2},
	}

	for _, tc := range testCase {
		if tc.it.Len() != tc.expect {
			t.Errorf("testcase error: %v", tc)
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
	mockToper := TablespaceToper{TS: make([]Tablespace, 10)}

	for _, tc := range testCase {
		l := len(mockToper.Data(tc.n).([]Tablespace))
		if l != tc.expect {
			t.Errorf("testcase error: %v", tc)
		}
	}
}

func TestGetScript(t *testing.T) {
	s, err := getScript()
	if err != nil {
		t.Errorf("getScript return error: %v", err)
	}
	if s == "" {
		t.Error("getScript error: s is empty")
	}
}

func TestCollect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s, _ := getScript()
	rows := mock.NewRows([]string{"TABLESPACE_NAME", "TOTAL_BYTES", "FILE_BYTES", "USE_BYTES", "USE_PCT"})
	rows.AddRow("SYSTEM", 10000, 1000, 100, 10)
	rows.AddRow("SYSAUX", 10000, 1000, 100, 10)
	mock.ExpectQuery(regexp.QuoteMeta(s)).WillReturnRows(rows)

	tt := TablespaceToper{DB: db}
	if err = tt.Collect(); err != nil {
		t.Error(err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(s)).WillReturnError(errors.New("mock error"))
	if err = tt.Collect(); err == nil {
		t.Error("should return error")
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("unmet expectation error: %s", err)
	}
}
