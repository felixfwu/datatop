package filesystem

import (
	"os"
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

func isDir(f string) bool {
	s, err := os.Stat(f)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/*
func walkCurr(f string) error {

	return nil
}
*/
