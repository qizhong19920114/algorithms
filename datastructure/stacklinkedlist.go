package datastructure

type stackLinkedList struct {
	Current *node
	Lenght  int
}

type node struct {
	Item string
	Next *node
}

func NewStackLinkedList() Stack {
	return &stackLinkedList{}
}

func (s *stackLinkedList) Push(obj string) {
	newNode := &node{}
	newNode.Item = obj
	newNode.Next = s.Current
	s.Current = newNode
	s.Lenght++
}

func (s *stackLinkedList) Pop() string {
	if !s.IsEmpty() {
		item := s.Current.Item
		s.Current = s.Current.Next
		s.Lenght--
		return item
	}
	return "" //todo: manage error
}

func (s *stackLinkedList) IsEmpty() bool {
	return s.Lenght == 0
}

func (s *stackLinkedList) Size() int {
	return s.Lenght
}
