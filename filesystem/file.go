package filesystem

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
