package main

/*
** Disclaimer:
** Student should use own implemented string.Split(...) and strconv.Atoi(...)
 */

import (
	"os"
	"fmt"
	"strings" // NOT ALLOWED
	"strconv" // NOT ALLOWED
	tools "../tools"
	stack "../stack"
	pushswap "../pushswap"
)

type Stack = stack.Stack
type PushSwap = pushswap.PushSwap

func sortSecondStack(PS *PushSwap, size int) int {

}

func bestMove(stack *Stack) int {

}

func insertionSort(PS *PushSwap, size int) {

}

func sort(PS *PushSwap, size int) {
	if size == 2 {
		
	}
}

func main() {
	var PS PushSwap
	var numbers []string

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [args]\n", os.Args[0])
		fmt.Println("\targs - space separated numbers")
		return
	}
	numbers := strings.Split(os.Args[1], " ")
	for _, num := range numbers {
		if tools.CheckNumber(num) == false {
			tools.Exit()
		}
		tmpNum, _ := strconv.Atoi(num)
		PS.Array = append(PS.Array, tmpNum)
	}
	PS.Init(len(numbers))
	PS.Print = 1
	tools.Qsort(PS.Array, 0, len(numbers)-1)
	if tools.CheckDuplicates(PS.Array) == false {
		tools.Exit()
	}
	sort(&PS, len(numbers))
}