package trees

//Queue implements a queue of Tree nodes
type Queue struct {
	items []Tree
}

//Newq returns a new queue
func Newq() *Queue {
	return &Queue{items: []Tree{}}
}

//Push pushes a new item onto the queue
func (q *Queue) Push(t Tree) {
	q.items = append(q.items, t)
}

//Pop returns the next item in the queue
func (q *Queue) Pop() *Tree {
	t := q.items[0]
	q.items = q.items[1:len(q.items)]

	return &t
}

//Size returns the size of the queue
func (q *Queue) Size() int {
	return len(q.items)
}
