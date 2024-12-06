package pack_me_file_processing

import (
	"fmt"
	"os"
	"strings"

	compress "github.com/b1rr0/PackMe/compress"
	dir "github.com/b1rr0/PackMe/dir-processing"
)

type PackMeNode struct {
	Name     string
	Children *[]PackMeNode
	Data     []byte
}

func NewPackMeNode(node *dir.Node) *PackMeNode {
	parts := strings.Split(node.Name, "/")
	packMeNode := &PackMeNode{Name: parts[len(parts)-1]}

	if children, ok := node.Data.([]dir.Node); ok {
		packMeNode.Children = &[]PackMeNode{}
		for _, child := range children {
			*packMeNode.Children = append(*packMeNode.Children, *NewPackMeNode(&child))
		}
	} else {
		packMeNode.Data = compress.Compress(node.Data.([]byte))
	}
	return packMeNode
}

func (node *PackMeNode) PrintNode() {
	fmt.Println(node.Name)
	if children := node.Children; children != nil {
		for _, child := range *children {
			child.PrintNode()
		}
	} else {
		fmt.Println(node.Data)
	}
}

func (node *PackMeNode) Unpack(pathToFile string) {

	//	runtime.EventsEmit(context.TODO(), "nodeProcessed", map[string]interface{}{
	//		"name": node.Name,
	//	})

	if node.Children != nil {
		os.MkdirAll(pathToFile+"/"+node.Name, 0755)
		for _, child := range *node.Children {
			fmt.Println("child")
			child.Unpack(pathToFile + "/" + node.Name)

		}

	} else {
		filePath := pathToFile + "/" + node.Name
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()
		_, err = file.Write(node.Data)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		}
	}
}
