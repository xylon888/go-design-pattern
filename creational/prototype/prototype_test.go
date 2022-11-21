package prototype

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name:     "Folder1",
		children: []Inode{file1},
	}

	folder2 := &Folder{
		name:     "Folder2",
		children: []Inode{folder1, file2, file3},
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("   ")

	cloneFolder := folder2.clone()

	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("   ")

}

func TestNewPrototypeManager(t *testing.T) {
	tests := []struct {
		name string
		want *PrototypeMgr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrototypeManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrototypeManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototypeMgr_Get(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		p    *PrototypeMgr
		args args
		want Cloneable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrototypeMgr.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototypeMgr_Set(t *testing.T) {
	type args struct {
		name      string
		prototype Cloneable
	}
	tests := []struct {
		name string
		p    *PrototypeMgr
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Set(tt.args.name, tt.args.prototype)
		})
	}
}

var mgr *PrototypeMgr

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func init() {
	mgr = NewPrototypeManager()
	t1 := &Type1{
		name: "type1",
	}
	mgr.Set("t1", t1)

	t2 := &Type2{
		name: "type2",
	}
	mgr.Set("t2", t2)
}

func TestClone(t *testing.T) {
	t1 := mgr.Get("t1")
	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromMgr(t *testing.T) {
	c := mgr.Get("t1").Clone()
	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}
}
