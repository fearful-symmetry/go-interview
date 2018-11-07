package trees

import (
	"reflect"
	"testing"
)

func TestTreeBFS(t *testing.T) {

	tt := NewTree(10)

	tt.AddList([]int{5, 11, 1, 6, 12})
	st := tt.BFS()

	if !reflect.DeepEqual(st, []int{10, 5, 11, 1, 6, 12}) {
		t.Fatalf("Bad BFS sort: %#v", t)
	}
}

func TestTreeDFS(t *testing.T) {
	tt := NewTree(10)

	tt.AddList([]int{5, 13, 1, 6, 12, 15})
	st := tt.DFS()

	if !reflect.DeepEqual(st, []int{10, 5, 1, 6, 13, 12, 15}) {
		t.Fatalf("Bad DFS sort: %#v", t)
	}
}

func TestTreeIsBal(t *testing.T) {
	tt := NewTree(10)

	tt.AddList([]int{5, 13, 1, 6, 12, 15})
	//ib := tt.IsBal()

	b1 := tt.IsBal()
	if !b1 {
		t.Fatalf("Tree at  should be true")
	}
	//Should still be true
	tt.Add(4)
	b2 := tt.IsBal()
	if !b2 {
		t.Fatalf("Tree at  should be true")
	}
	tt.Add(3)
	b3 := tt.IsBal()
	if b3 {
		t.Fatalf("Tree at  should be false")
	}
}
