package main

/*
** Stack Data Structure
 */
type Stack struct {
	List   *Node
	Length int
}

func (s *Stack) PushFront(number int) {
	newNode := createNode(number)
	newNode.Next = s.List
	s.List = newNode
	s.Length++
}

func (s *Stack) PushBack(number int) {
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
	var tmpNode *Node
	var number *int

	if s.Length == 0 {
		return nil
	}
	tmpNode = s.List.Next
	*number = tmpNode.Number

	s.List = tmpNode
	s.Length--
	return number
}

/*
** Function checks whether stack is sorted in given window size
** 0 1 2 3 4 5 6 - stack
** 5 - size
** |0 1 2 3 4| - function checks only first 5 numbers in stack
*/
func (s Stack) IsSorted(size int) bool {
	i := 0
	if s.Length == 0 {
		return false
	}
	for tmpNode := s.List; tmpNode.Next != nil && i < size-1; tmpNode = tmpNode.Next {
		if tmpNode.Number > tmpNode.Next.Number {
			return false
		}
		i++
	}
	return true
}

/*
** Function is the same as Stack.IsSorted, but checks for reverse sorted
*/
func (s Stack) IsRevSorted(size int) bool {
	i := 0
	if s.Length == 0 {
		return false
	}
	for tmpNode := s.List; tmpNode.Next != nil && i < size-1; tmpNode = tmpNode.Next {
		if tmpNode.Number < tmpNode.Next.Number {
			return false
		}
		i++
	}
	return true
}
