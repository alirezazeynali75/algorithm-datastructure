package datastructure

type Node struct {
	Data string
	NextNodePointer *Node
}

func (n *Node) GetNext() (Node, bool) {
	if n.NextNodePointer == nil {
		return Node{}, false
	}
	return *n.NextNodePointer, true
}

func (n *Node) SetNext(next *Node) {
	n.NextNodePointer = next
}

func (n *Node) GetData() string {
	return n.Data
}

func (n *Node) SetData(data string) {
	n.Data = data
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddAtFront(n *Node) {
	n.SetNext(l.head)
	l.head = n
}

func (l *LinkedList) IsEmpty() bool

func (l *LinkedList) AddAtEnd(n *Node) {
	if l.head == nil {
		l.head = n
		n.SetNext(nil)
	} else {
		n2 := l.getLastNode()
		n2.SetNext(n)
	}
}

func (l *LinkedList) getLastNode() *Node {
	ptr := l.head
	for ptr.NextNodePointer != nil {
		ptr = ptr.NextNodePointer
	}
	return ptr
}
