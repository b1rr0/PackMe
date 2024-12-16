package dir_processing

import "fmt"

type Node struct {
	Name string

	Data interface{}
}

func NewNode(path string, fileCount *int) *Node {
	*fileCount++
	isDir, err := CheckIsDir(path)
	if err != nil {
		return nil
	}
	node := &Node{Name: path}

	if isDir {
		node.Data = []Node{}
		dirEntries := ReadDir(path)
		for _, dirEntry := range dirEntries {
			child := NewNode(dirEntry, fileCount)
			node.Data = append(node.Data.([]Node), *child)
		}
	} else {
		node.Data = ReadFile(path)
	}

	return node
}

func (node *Node) PrintNode() {
	fmt.Println(node.Name)
	if children, ok := node.Data.([]Node); ok {
		for _, child := range children {
			child.PrintNode()
		}
		fmt.Println("___________________")
	} else {
		fmt.Println(node.Data)
	}
}
