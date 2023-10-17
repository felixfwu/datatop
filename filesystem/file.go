package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type File struct {
	Path      string
	FileCount int
}

type FileToper struct {
	Root string
	Dirs []File
}

func (fs FileToper) Swap(i int, j int) {
	fs.Dirs[i], fs.Dirs[j] = fs.Dirs[j], fs.Dirs[i]
}

func (fs FileToper) Len() int {
	return len(fs.Dirs)
}

func (fs FileToper) Less(i int, j int) bool {
	return fs.Dirs[i].FileCount < fs.Dirs[j].FileCount
}

func (fs FileToper) Data(n int) interface{} {
	return interface{}(fs.Dirs[:n])
}

func (fs *FileToper) Collect() error {
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

func walkCurr(path string) ([]File, error) {
	osf, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("walkCurr open path=%s error", path), err)
	}
	fs, err := osf.ReadDir(-1)
	defer osf.Close()
	if err != nil {
		return nil, errors.Join(errors.New("walkCurr read dir error"), err)
	}

	ds := []File{}
	d := File{Path: path, FileCount: 0}
	for _, f := range fs {
		if f.IsDir() {
			ds = append(ds, File{Path: f.Name(), FileCount: 0})
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
