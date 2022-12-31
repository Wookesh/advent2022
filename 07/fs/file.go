package fs

type File struct {
	size int
	name string
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) IsDir() bool {
	return false
}

func (f *File) Size() int {
	return f.size
}
