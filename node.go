package main

/*
** Node Used for List Data Structure
*/
type Node struct {
	Number	int
	Next	*node
}

func createNode(int number) *Node {
	return &Node{number, nil}
}

/*
** Stack Data Structure
*/
type Stack struct {
	List	*Node // head of stack
	Length	int
}

func (s *Stack) pushFront(number int) {
	newNode := createNode(number)
	newNode.Next = s.List
	s.List = newNode
	s.Length++
}

func (s *Stack) pushBack(number int) {
	newNode := createNode(number)
	s.Length++
	if s.List == nil {
		s.List = newNode
		return
	}
	for tmp := s.List; tmp.Next != nil; tmp = tmp.Next {
		// Go to Back of Stack
	}
	tmp.Next = newNode
}

/*
** GoLang has Garbage Collector, no need for freeing memory
*/
func (s *Stack) Pop() *int {
	var tmpNode	*node
	var number	*int

	if s.Length == 0 {
		return nil
	}
	tmpNode = s.List.Next
	*number = tmpNode.Number

	s.List = tmpNode
	s.Length--
	return number
}