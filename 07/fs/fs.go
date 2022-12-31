package fs

type FS struct {
	root *Dir

	currentDir *Dir
}

func NewFS() *FS {
	root := NewDir("", nil)

	return &FS{
		root:       root,
		currentDir: root,
	}
}

func (f *FS) Root() *Dir {
	return f.root
}

func (f *FS) CurrentDir() *Dir {
	return f.currentDir
}

func (f *FS) CD(p string) {
	if p == ".." {
		f.currentDir = f.currentDir.parent
	} else if p == "/" {
		f.currentDir = f.root
	} else {
		e, ok := f.currentDir.children[p]
		if !ok || !e.IsDir() {
			panic("invalid path")
		}
		f.currentDir = e.(*Dir)
	}
}

func (f *FS) Walk(fn func(e Entity)) {
	f.root.Walk(fn)
}

type Entity interface {
	Name() string
	IsDir() bool
	Size() int
}
