package stack

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	nodes []*Node
	top   int
}

type Node struct {
	Value int
}

func NewStack(size int) *Stack {
	return &Stack{make([]*Node, size), -1}
}

func NewNode(v int) *Node {
	return &Node{v}
}

func (s Stack) String() string {

	var str string = "[ "
	if s.top == -1 {
		str += "Stack is empty!!!"
		return str + "]"
	}

	for i := 0; i <= s.top; i++ {
		str += strconv.Itoa(s.nodes[i].Value) + " "
	}

	str += "]"
	return str
}

func (s Stack) isEmpty() bool {
	return s.top == -1
}

func (s Stack) isFull() bool {
	return len(s.nodes) == s.top+1
}

func (s *Stack) push(node *Node) {
	if s.isFull() {
		fmt.Println("Stack is full")
		return
	}
	s.top++
	s.nodes[s.top] = node
}

func (s *Stack) pop() (*Node, error) {
	if s.isEmpty() {
		return nil, errors.New("error : Stack is empty")
	}

	node := s.nodes[s.top]
	s.top--
	return node, nil
}

func testStack() {
	var size int
	fmt.Print("스택 사이즈 지정 : ")
	_, err := fmt.Scanln(&size)
	if err != nil {
		return
	}
	s := NewStack(size)

LOOP:
	for {
		var opt int
		fmt.Print("----------\n1. push\n2. pop\n3. print\n4. end\n----------\n옵션 선택: ")
		_, err := fmt.Scanln(&opt)
		if err != nil {
			return
		}

		switch opt {
		case 1:
			fmt.Print("넣을 숫자 입력 : ")
			var v int
			_, err := fmt.Scanln(&v)
			if err != nil {
				fmt.Println(err)
			}
			s.push(NewNode(v))
		case 2:
			n, err := s.pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("꺼낸 숫자 :", n.Value)
			}
		case 3:
			fmt.Println(s)
		case 4:
			break LOOP
		default:
			fmt.Println("옵션만 선택해 주십시오.")
		}
	}
}
