package oracle

import (
	"database/sql"
	"errors"
	"os"
)

type Tablespace struct {
	Name        string  // tablespace name
	SegmentSize int     // bytes
	FileSize    int     // size in bytes
	TotalSize   int     // tablespace total size in bytes, include extensible size
	TPCT        float32 // SegmentSize / TotalSize
}

type TablespaceToper struct {
	DB *sql.DB
	TS []Tablespace
}

func (tt TablespaceToper) Swap(i, j int) {
	tt.TS[i], tt.TS[j] = tt.TS[j], tt.TS[i]
}

func (tt TablespaceToper) Less(i, j int) bool {
	return tt.TS[i].TPCT < tt.TS[j].TPCT
}

func (tt TablespaceToper) Len() int {
	return len(tt.TS)
}

func (tt TablespaceToper) Data(n int) interface{} {
	return interface{}(tt.TS[:n])
}

func (tt *TablespaceToper) Collect() error {
	sql, err := getScript()
	if err != nil {
		return err
	}
	rows, err := tt.DB.Query(sql)
	if err != nil {
		return errors.Join(errors.New("collect error"), err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Tablespace
		rows.Scan(&t.Name, &t.TotalSize, &t.FileSize, &t.SegmentSize, &t.TPCT)
		tt.TS = append(tt.TS, t)
	}
	return nil
}

func getScript() (string, error) {
	b, err := os.ReadFile("script.sql")
	if err != nil {
		return "", errors.Join(errors.New("getScript error"), err)
	}
	return string(b), nil
}
