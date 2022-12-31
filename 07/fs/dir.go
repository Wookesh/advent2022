package fs

type Dir struct {
	name     string
	children map[string]Entity

	parent     *Dir
	cachedSize int
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:     name,
		parent:   parent,
		children: make(map[string]Entity),
	}
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) Size() int {
	if d.cachedSize != 0 {
		return d.cachedSize
	}
	s := 0
	for _, c := range d.children {
		s += c.Size()
	}
	d.cachedSize = s
	return s
}

func (d *Dir) Add(e Entity) {
	d.children[e.Name()] = e
}

func (d *Dir) Walk(f func(e Entity)) {
	for _, c := range d.children {
		f(c)
		if c.IsDir() {
			d := c.(*Dir)
			d.Walk(f)
		}
	}
}
