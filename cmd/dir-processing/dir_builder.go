package dir_processing

import (
	"fmt"
)

func BuildDirTree(path string) (*Node, int) {
	fileCount := 0
	root := NewNode(path, &fileCount, nil)

	//root.PrintNode()
	//	fmt.Println(root.Data)
	fmt.Println("FILE COUNT = ", fileCount)
	fmt.Println("FILE COUNT = ", fileCount)

	fmt.Println("FILE COUNT = ", fileCount)

	fmt.Println("FILE COUNT = ", fileCount)
	fmt.Println("FILE COUNT = ", fileCount)
	fmt.Println("FILE COUNT = ", fileCount)
	fmt.Println("FILE COUNT = ", fileCount)

	return root, fileCount
}
