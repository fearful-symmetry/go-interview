package trees

import (
	"fmt"
)

//tree data structures are cool.
//A somewhat hasty Binary Search Tree.

//Tree is a binary search tree type
type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

//NewTree returns a new tree
func NewTree(v int) *Tree {
	return &Tree{Data: v}
}

//Add a value to the given node
func (t *Tree) Add(new int) {

	if new == t.Data {
		return
	}

	if new < t.Data {
		if t.Left == nil {
			t.Left = &Tree{Data: new}
		}
		t.Left.Add(new)
	}

	if new > t.Data {
		if t.Right == nil {
			t.Right = &Tree{Data: new}
		}
		t.Right.Add(new)
	}

}

//AddList adds multiple nodes to the tree
func (t *Tree) AddList(new []int) {

	for _, toAdd := range new {
		t.Add(toAdd)
	}
}

//Print Prints the binary tree
func (t *Tree) Print() {

	if t.Left == nil && t.Right == nil {
		fmt.Printf("Leaf: %d\n", t.Data)
		return
	}
	fmt.Printf("Data: %d\n", t.Data)

	if t.Left != nil {
		fmt.Printf("Left Branch of %d:\n", t.Data)
		t.Left.Print()
	}

	if t.Right != nil {
		fmt.Printf("Right Branch of %d:\n", t.Data)
		t.Right.Print()
	}

}

//BFS implements a....BFS using queues.
func (t *Tree) BFS() []int {

	tq := Newq()
	tq.Push(*t)
	visted := make(map[*Tree]bool)
	var sortedList []int
	for {

		if tq.Size() == 0 {
			break
		}

		currentTree := tq.Pop()

		visted[currentTree] = true

		if currentTree.Left != nil && !visted[currentTree.Left] {
			tq.Push(*currentTree.Left)
			visted[currentTree.Left] = true
		}

		if currentTree.Right != nil && !visted[currentTree.Right] {
			tq.Push(*currentTree.Right)
			visted[currentTree.Right] = true
		}
		sortedList = append(sortedList, currentTree.Data)
		//fmt.Println(currentTree.Data)
	}

	return sortedList
}

//DFS implements a DFS using iteration
func (t *Tree) DFS() []int {

	if t.Left == nil && t.Right == nil {
		return []int{t.Data}
	}

	dat := []int{t.Data}

	if t.Left != nil {
		dat = append(dat, t.Left.DFS()...)
	}

	if t.Right != nil {
		dat = append(dat, t.Right.DFS()...)
	}

	return dat
}
