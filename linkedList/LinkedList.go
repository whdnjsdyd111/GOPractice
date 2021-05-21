package linkedList

import "strconv"

type LinkedList struct {
	Header *Node
	Tail   *Node
	Count  int
}

type Node struct {
	Next  *Node
	Value int
}

func (link LinkedList) String() string {
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

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (link *LinkedList) FindByValue(val int) *Node {
	for non := link.Header; non != nil; {
		if non.Value == val {
			return non
		} else {
			non = non.Next
		}
	}
	return nil
}

func (link *LinkedList) InsertFirstNode(node *Node) {
	if link.Header == nil {
		link.Header = node
		link.Tail = node
	} else {
		node.Next = link.Header
		link.Header = node
	}
	link.Count++
}

func (link *LinkedList) InsertLastNode(node *Node) {
	if n := link.Tail; n != nil {
		n.Next = node
		link.Tail = node
	} else {
		link.Header = node
		link.Tail = node
	}
	link.Count++
}

func (link *LinkedList) InsertMiddleNode(newNode *Node, val int) bool {
	prevNode := link.FindByValue(val)
	if prevNode == nil {
		return false
	}
	for n := link.Header; n != nil; {
		if n == prevNode {
			newNode.Next = prevNode.Next
			prevNode.Next = newNode
			if link.Tail == prevNode {
				link.Tail = newNode
			}
			link.Count++
			return true
		} else {
			n = n.Next
		}
	}
	return false
}

func (link *LinkedList) RemoveLastNode() {
	if link.Tail == nil {
		return
	}

	if link.Tail == link.Header {
		link.Count--
		link.Tail = nil
		link.Header = nil
	}

	for n := link.Header; n != nil; {
		if n.Next == link.Tail {
			n.Next = nil
			link.Tail = n.Next
			link.Count--
		}
		n = n.Next
	}
}

func (link *LinkedList) RemoveNode(r *Node) bool {
	if link.Header == r {
		link.Header = r.Next
		link.Count--
		return true
	}
	for n := link.Header; n != nil; {
		if n.Next == r {
			n.Next = r.Next
			link.Count--
		}
		n = n.Next
	}
	return false
}

func (link *LinkedList) IndexOf(i int) *Node {
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
	return &Node{nil, val}
}
