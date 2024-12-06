package pack_me_file_processing

import (
	dir "github.com/b1rr0/PackMe/dir-processing"
)

type PackMeFormat struct {
	Nodes []*PackMeNode
	Size  int
}

func NewPackMeFormatFromDirNode(nodes []*dir.Node, size int) *PackMeFormat {
	packMeNodes := []*PackMeNode{}
	for _, node := range nodes {
		packMeNode := NewPackMeNode(node)
		packMeNodes = append(packMeNodes, packMeNode)
	}

	return &PackMeFormat{Nodes: packMeNodes, Size: size}
}
