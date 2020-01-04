package node

/*
** Node Used for List Data Structure
 */
type Node struct {
	Number int
	Next   *Node
}

func CreateNode(number int) *Node {
	return &Node{number, nil}
}
