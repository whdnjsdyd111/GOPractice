package doublylinkedlist

import "strconv"

type DoublyLinkedList struct {
	Header *Node
	Tail   *Node
	Count  int
}

type Node struct {
	Prev  *Node
	Value int
	Next  *Node
}

func (link DoublyLinkedList) String() string {
	str := "["
	if n := link.Header; n == nil {
		str += "]"
		return str
	} else {
		for n != nil {
			str += strconv.Itoa(n.Value) + ", "
			n = n.Next
		}
		str += "]"
		return str
	}
}

func (link DoublyLinkedList) Reverse() string {
	str := "["
	if n := link.Tail; n == nil {
		str += "]"
		return str
	} else {
		for n != nil {
			str += strconv.Itoa(n.Value) + ", "
			n = n.Prev
		}
		str += "]"
		return str
	}
}

func NewLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

func (link *DoublyLinkedList) FindNode(val int) *Node {
	for n := link.Header; n != nil; {
		if n.Value == val {
			return n
		} else {
			n = n.Next
		}
	}
	return nil
}

func (link *DoublyLinkedList) InsertFirstNode(node *Node) {
	if link.Header == nil {
		link.Header = node
		link.Tail = node
	} else {
		node.Next = link.Header
		link.Header = node
	}
	link.Count++
}

func (link *DoublyLinkedList) InsertLastNode(newNode *Node) {
	if n := link.Tail; n != nil {
		n.Next = newNode
		newNode.Prev = n
		link.Tail = newNode
	} else {
		link.Header = newNode
		link.Tail = newNode
	}
	link.Count++
}

func (link *DoublyLinkedList) InsertNextNode(n *Node, v int) bool {
	old := link.FindNode(v)
	if old == nil || n == nil {
		return false
	}

	if old.Next == nil {
		old.Next = n
		n.Prev = old
	} else {
		n.Next = old.Next
		old.Next = n
		n.Next.Prev = n
		n.Prev = old
	}
	link.Count++
	return true
}

func (link *DoublyLinkedList) InsertPrevNode(n *Node, v int) bool {
	old := link.FindNode(v)
	if old == nil || n == nil {
		return false
	}

	if old.Prev == nil {
		old.Prev = n
		n.Next = old
	} else {
		n.Prev = old.Prev
		old.Prev = n
		n.Prev.Next = n
		n.Next = old
	}
	link.Count++
	return true
}

func (link *DoublyLinkedList) RemoveLastNode() {
	if link.Tail == nil {
		return
	}

	if n := link.Tail; n == link.Header {
		link.Tail = nil
		link.Header = nil
	} else {
		n.Prev.Next = nil
	}
	link.Count--
}

func (link *DoublyLinkedList) RemoveFirstNode() {
	if n := link.Header; n == nil {
		return
	} else {
		n.Next.Prev = nil
		link.Header = n.Next
	}
	link.Count--
}

func (link *DoublyLinkedList) RemoveNode(val int) bool {
	if n := link.FindNode(val); n == nil {
		return false
	} else {
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
		link.Count--
		return true
	}
}

func (link *DoublyLinkedList) RemovePrevNode(val int) bool {
	if n := link.FindNode(val); n == nil {
		return false
	} else {
		n.Prev.Prev.Next = n
		n.Prev = n.Prev.Prev
		link.Count--
		return true
	}
}

func (link *DoublyLinkedList) RemoveNextNode(val int) bool {
	if n := link.FindNode(val); n == nil {
		return false
	} else {
		n.Next.Next.Prev = n
		n.Next = n.Next.Next
		link.Count--
		return true
	}
}

func (link *DoublyLinkedList) IndexOf(i int) *Node {
	node := link.Header
	if node == nil {
		return nil
	}
	for n := 0; n < i; n++ {
		if node == nil {
			return nil
		} else {
			node = node.Next
		}
	}
	return node
}

func NewNode(val int) *Node {
	return &Node{nil, val, nil}
}
