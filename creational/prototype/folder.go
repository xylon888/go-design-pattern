package prototype

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, item := range f.children {
		item.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tmpChildren []Inode
	for _, item := range f.children {
		copy := item.clone()
		tmpChildren = append(tmpChildren, copy)
	}
	cloneFolder.children = tmpChildren
	return cloneFolder
}
