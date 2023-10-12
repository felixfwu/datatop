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

func (f *File) collect() error {
	ds, err := walkCurr(f.Root)
	if err != nil {
		return err
	}
	for _, d := range ds {
		c, err := walkChild(d.Name)
		if err != nil {
			return err
		}
		d.Cnt = c
	}
	return nil
}

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

	dirs := []Dir{}
	for _, file := range files {
		fp := fmt.Sprintf("%s/%s", f, file.Name())
		if file.IsDir() {
			dirs = append(dirs, Dir{Name: fp, Cnt: 0})
		} else {
			dirs = append(dirs, Dir{Name: fp, Cnt: 1})
		}

	}
	return dirs, nil
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
