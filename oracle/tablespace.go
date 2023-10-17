package oracle

import (
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
)

type Tablespace struct {
	Name        string
	SegmentSize int // bytes
	FileSize    int //bytes
	TotalSize   int //bytes
	PCT         int
}

type TablespaceToper struct {
	TS []Tablespace
}

func testdb() error {
	connStr := go_ora.BuildUrl("192.168.0.31", 1521, "orcl", "test", "oracle", nil)
	conn, err := sql.Open("oracle", connStr)
	if err != nil {
		return err
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn.Stats())
	return nil
}
