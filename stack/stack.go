package stack

import (
	node "../node"
	tools "../tools"
)

type Node = node.Node

/*
** Stack Data Structure
 */
type Stack struct {
	List   *Node
	Length int
}

func (s *Stack) PushFront(number int) {
	newNode := node.CreateNode(number)
	newNode.Next = s.List
	s.List = newNode
	s.Length++
}

func (s *Stack) PushBack(number int) {
	var tmpNode *Node
	newNode := node.CreateNode(number)
	s.Length++
	if s.List == nil {
		s.List = newNode
		return
	}
	for tmpNode := s.List; tmpNode.Next != nil; tmpNode = tmpNode.Next {
		// Go to Back of Stack
	}
	tmpNode.Next = newNode
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

/*
** PlaceIn: Don't Remember
 */
func (s Stack) PlaceIn(number int) int {
	var tmpNode *Node

	for tmpNode = s.List; tmpNode.Next != nil; tmpNode = tmpNode.Next {
		// Move tmpNode to the last node
	}
	if (number < tmpNode.Number && tmpNode.Number < s.List.Number) {
		return 0
	}
	if (number < s.List.Number) {
		return 1
	}
	return 0
}


/*
** Function returns:
** 0 - if best place for `number` is in first half of stack
** 1 - if best place for `number` is in second half of stack
 */
func (s Stack) GetPlace(number int) int {
	var prev, ibest, i int
	var tmpNode *Node

	for tmpNode = s.List; tmpNode.Next != nil; tmpNode = tmpNode.Next {
		// Move tmpNode to the Last node in stack
	}
	prev = tmpNode.Number
	tmpNode = s.List
	if number < prev && prev < tmpNode.Number {
		return 1
	}
	i = 0
	for tmpNode != nil {
		if (tmpNode.Number > number && (number > prev || prev == -1)) {
			ibest = i
		}
		prev = tmpNode.Number
		tmpNode = tmpNode.Next
		i++
	}
	if ibest < s.Length / 2 {
		return 0
	} else {
		return 1
	}
}

/*
** Function returns min number in stack
 */
func (s Stack) Min() int {
	min := s.List.Number
	for tmpNode := s.List; tmpNode != nil; tmpNode = tmpNode.Next {
		if tmpNode.Number < min {
			min = tmpNode.Number
		}
	}
	return min
}

/*
** Function returns maximum number from start position til start+length position
** start can be negative.
** Example:
** 8 4 5 1 3 5 3 - stack
** Max(2, 3) will return 5.
** 0 1  2 3 4  5 6 - indexes
** -------------
** 8 4 |5 1 3| 5 3
** Max number will be returned from window |5 1 3|
 */
func (s Stack) Max(start, length int) int {
	var tmpNode *Node
	var i, max int

	tmpNode = s.List
	if start < 0 {
		i = 0
		for ; tmpNode != nil && i < s.Length + start; tmpNode = tmpNode.Next {
			// Move tmpNode pointer to start position
		} 
	}
	max = tmpNode.Number
	i = 0
	for tmpNode != nil && i < length {
		if tmpNode.Number > max {
			max = tmpNode.Number
		}
		if tmpNode.Next == nil {
			tmpNode = s.List
		} else {
			tmpNode = tmpNode.Next
		}
	}
	return max
}

/*
** Function returns:
** 0 - if biggest number in stack is in the first half of the stack
** 1 - if max number is in second half of stack
 */
func (s Stack) GetBig() int {
	var max, imax, i int

	for tmpNode := s.List; tmpNode != nil; tmpNode = tmpNode.Next {
		if tmpNode.Number > max {
			max = tmpNode.Number
			imax = i
		}
		i++
	}
	if imax < s.Length / 2 {
		return 0
	} else {
		return 1
	}
}

/*
** Old Name: smaller_in
** Function checks whethers stack of size `size` has elements less than median
 */
func (s Stack) IsLess(size, median int) bool {
	i := 0
	for tmpNode := s.List; tmpNode != nil && i < size; tmpNode, i = tmpNode.Next, i + 1 {
		if tmpNode.Number < median {
			return true
		}
	}
	return false
}

/*
** Function returns median number in stack of size `size`
 */
func (s Stack) Median(size int) int {
	var array []int
	i := 0

	for tmpNode := s.List; tmpNode != nil && i < size; tmpNode, i = tmpNode.Next, i + 1 {
		array = append(array,tmpNode.Number)
	}
	tools.Qsort(array, 0, s.Length)
	return array[(i-1)/2]
}
