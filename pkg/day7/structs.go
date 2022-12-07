package day7

type FileInfo struct {
	Name  string
	IsDir bool
	Size  int
}

type Directory struct {
	Name     string
	Parent   *Directory `json:"-"`
	Children map[string]*Directory
	Files    map[string]int
	size     *int
}

func NewDirectory(name string) *Directory {
	return &Directory{
		Name:     name,
		Children: map[string]*Directory{},
		Files:    map[string]int{},
	}
}

func (d *Directory) AddFileInfo(fileInfo *FileInfo) {
	if fileInfo.IsDir {
		newDir := NewDirectory(fileInfo.Name)
		newDir.Parent = d
		d.Children[fileInfo.Name] = newDir
		return
	}
	d.Files[fileInfo.Name] = fileInfo.Size
}

func (d *Directory) Root() *Directory {
	if d.Name == "/" {
		return d
	}
	return d.Parent.Root()
}

func (d *Directory) Size() int {
	if d.size != nil {
		return *d.size
	}
	var size int
	for _, s := range d.Files {
		size += s
	}
	for _, c := range d.Children {
		size += c.Size()
	}
	d.size = &size
	return size
}

// func (d *Directory) SizeGreaterThan(gt int) int {

// }
