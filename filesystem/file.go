package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Dir struct {
	Path      string
	FileCount int
}

type FileSystem struct {
	Root string
	Dirs []Dir
}

func (fs FileSystem) Swap(i int, j int) {
	fs.Dirs[i], fs.Dirs[j] = fs.Dirs[j], fs.Dirs[i]
}

func (fs FileSystem) Len() int {
	return len(fs.Dirs)
}

func (fs FileSystem) Less(i int, j int) bool {
	if fs.Dirs[i].FileCount < fs.Dirs[j].FileCount {
		return true
	}
	return false
}

func (fs FileSystem) Data(n int) interface{} {
	return interface{}(fs.Dirs[:n])
}

func (fs *FileSystem) Collect() error {
	ds, err := walkCurr(fs.Root)
	if err != nil {
		return errors.Join(errors.New("filesystem collect error"), err)
	}
	for i, d := range ds {
		if d.Path != fs.Root {
			c, err := walkChild(filepath.Join(fs.Root, d.Path))
			if err != nil {
				return errors.Join(errors.New("filesystem collect error"), err)
			}
			ds[i].FileCount = c
		}
	}
	fs.Dirs = ds
	return nil
}

func walkCurr(path string) ([]Dir, error) {
	osf, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("walkCurr open path=%s error", path), err)
	}
	fs, err := osf.Readdir(-1)
	defer osf.Close()
	if err != nil {
		return nil, errors.Join(errors.New("walkCurr read dir error"), err)
	}

	ds := []Dir{}
	d := Dir{Path: path, FileCount: 0}
	for _, f := range fs {
		if f.IsDir() {
			ds = append(ds, Dir{Path: f.Name(), FileCount: 0})
		}
		d.FileCount++
	}
	return append(ds, d), nil
}

func walkChild(path string) (int, error) {
	cnt := -1
	err := filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error {
		return walkFunc(&cnt, path, entry, err)
	})
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func walkFunc(cnt *int, path string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	*cnt++
	return nil
}
