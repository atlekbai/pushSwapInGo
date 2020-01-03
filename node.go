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
