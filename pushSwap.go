/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 12:44:51
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 13:34:31
*/

package main

/*
** Push Swap Data Structure for subject
*/
type PushSwap struct {
	A, B	*Stack
	Array	[]int
	Length	int
	Ops		int
	Print	int
}

/*
** Initialization of PushSwap Data Structure
** Main Job: to push back numbers from PushSwap.Array to Stack PushSwap.A
*/
func (p *PushSwap) Init(length int) {
	p.Length = length
	p.Ops = 0
	p.Print = 0
	for i := 0; i < p.Length; i++ {
		p.A.PushBack(p.Array[i])
	}
}

/*
** Prints Push Swap's Stack A, B and ops
*/
func (p PushSwap) Print() {
	fmt.Print("Stack A: ")
	for tmpNode := p.A.List; tmpNode != nil; tmpNode = tmpNode.Next {
		fmt.Printf("%d ", tmpNode.Number)
	}
	fmt.Printf("\nStack B: ")
	for tmpNode := p.B.List; tmpNode != nil; tmpNode = tmpNode.Next {
		fmt.Printf("%d ", tmpNode.Number)
	}
	fmt.Printf("\nops: %d\n", p.Ops)
}

/*               Operations               */

/*             Swap Operations            */

/*
** Function swaps first two numbers of given a stack
** If swap cannot be done returns false
*/
func SwapStack(stack *Stack) bool {
	var tmpNode *Node

	if stack.Length < 2 {
		return false
	}
	tmpNode = stack.List
	stack.List = tmpNode.Next

	tmpNode.Next = stack.List.Next
	stack.List.Next = tmpNode
	return true
}

/*
** Swap A Stack
*/
func (p *PushSwap) SwapA() {
	if SwapStack(p.A) == false {
		return
	}
	if p.Print == 1 {
		fmt.Println("sa")
	}
	p.Ops++
}

/*
** Swap B Stack
*/
func (p *PushSwap) SwapB() {
	if SwapStack(p.B) == false {
		return
	}
	if p.Print == 1 {
		fmt.Println("sb")
	}
	p.Ops++
}

/*
** Swap both A and B Stacks
*/
func (p *PushSwap) SwapAB() {
	SwapStack(p.A)
	SwapStack(p.B)
	if p.Print == 1 {
		fmt.Println("ss")
	}
	p.Ops++
}

/*             Push Operations            */

/*
** Push A
** Will not push if Stack B is empty
*/
func (p *PushSwap) PushA() {
	if p.Print == 1 {
		fmt.Println("pa")
	}
	if p.B.Length > 0 {
		numberPtr := p.B.Pop()
		p.A.PushFront(*numberPtr)
	}
	p.Ops++
}

/*
** Push B
** Will not push if Stack A is empty
*/
func (p *PushSwap) PushB() {
	if p.Print == 1 {
		fmt.Println("pb")
	}
	if p.A.Length > 0 {
		numberPtr := p.A.Pop()
		p.B.PushFront(*numberPtr)
	}
	p.Ops++
}

/*             Rotate Operations            */

/*
** Function rotates stack one node next
** Example:
** 0 1 2 3 4 5 6 - stack
** After RotateStack:
** 1 2 3 4 5 6 0
*/
func RotateStack(stack *Stack) bool {
	var tmpNode, head *Node

	if stack.Length < 2 {
		return false
	}
	head = stack.List
	stack.List = stack.List.Next
	for tmpNode = stack.List; tmpNode.Next != nil; tmpNode = tmpNode.Next {
		// Move pointer tmpNode to the last Node of Stack
	}
	tmpNode.Next = head
	head.Next = nil
	return true
}

/*
** Rotate A
*/
func (p *PushSwap) RotateA() {
	if RotateStack(p.A) == false {
		return
	}
	if p.Print == 1 {
		fmt.Println("ra")
	}
	p.Ops++
}

/*
** Rotate A
*/
func (p *PushSwap) RotateB() {
	if RotateStack(p.B) == false {
		return
	}
	if p.Print == 1 {
		fmt.Println("rb")
	}
	p.Ops++
}

/*
** Rotate A and B
*/
func (p *PushSwap) RotateAB() {
	RotateStack(p.A)
	RotateStack(p.B)
	if p.Print == 1 {
		fmt.Println("rr")
	}
	p.Ops++
}

/*             Rev Rotate Operations            */

/*
** Function reverse rotates stack one node next
** Example:
** 0 1 2 3 4 5 6 - stack
** After RevRotateStack:
** 6 0 1 2 3 4 5
*/
func RevRotateStack(stack *Stack) bool {
	var tmpNode, penultimate, tail *Node

	if stack.Length < 2 {
		return false
	}
	for tmpNode = stack.List; tmpNode.Next.Next != nil; tmpNode = tmpNode.Next {
		// Move pointer tmpNode to the penultimate Node of Stack
	}
	penultimate = tmpNode
	tail = penultimate.Next

	penultimate.Next = nil
	tail.Next = stack.List
	stack.List = tail

	return true
}

/*
** Rev Rotate A
*/
func (p *PushSwap) RevRotateA() {
	if RevRotateStack(p.A) == false {
		return
	}
	if p.Print == 1 {
		fmt.Println("rra")
	}
	p.Ops++
}

/*
** Rev Rotate A
*/
func (p *PushSwap) RevRotateB() {
	if RevRotateStack(p.B) == false {
		return
	}
	if p.Print == 1 {
		fmt.Println("rrb")
	}
	p.Ops++
}

/*
** Rev Rotate A and B
*/
func (p *PushSwap) RevRotateAB() {
	RevRotateStack(p.A)
	RevRotateStack(p.B)
	if p.Print == 1 {
		fmt.Println("rrr")
	}
	p.Ops++
}