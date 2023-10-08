package fs

type Dir struct {
	Name string
	Cnt  int
}

type File struct {
	root string
	Dirs []dir
}
