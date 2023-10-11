package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

type Dir struct {
	Name string
	Cnt  int
}

type File struct {
	Root string
	Dirs []Dir
}

func (f *File) Swap(i int, j int) {
	f.Dirs[i], f.Dirs[j] = f.Dirs[j], f.Dirs[i]
}

func (f File) Len() int {
	return len(f.Dirs)
}

func (f File) Less(i int, j int) bool {
	if f.Dirs[i].Cnt < f.Dirs[j].Cnt {
		return true
	}
	return false
}

func (f File) data(n int) interface{} {
	return interface{}(f.Dirs[:n])
}

/*
func (f *File) collect() error {
	var sb strings.Builder
	r, err := os.Open(f.Root)
	if err != nil {
		return err
	}
	files, err := r.Readdir(-1)
	defer r.Close()
	if err != nil {
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
*/

func walkCurr(f string) ([]Dir, error) {
	r, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	files, err := r.Readdir(-1)
	defer r.Close()
	if err != nil {
		return nil, err
	}

	fd := Dir{Name: f, Cnt: 0}
	dirs := []Dir{}
	for _, file := range files {
		fp := fmt.Sprintf("%s/%s", f, file.Name())
		if file.IsDir() {
			dirs = append(dirs, Dir{Name: fp, Cnt: 0})
		} else {
			dirs = append(dirs, Dir{Name: fp, Cnt: 1})
		}
		fd.Cnt++
	}
	return append(dirs, fd), nil
}

func walkChild(f string) (int, error) {
	cnt := -1
	err := filepath.Walk(f, func(fname string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		cnt++
		return nil
	})
	if err != nil {
		return 0, err
	}

	return cnt, nil
}
