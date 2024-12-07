package pack_me_file_processing

import (
	dir "PackMe/internal/dir-processing"
	"context"
)

type PackMeFormat struct {
	Node *PackMeNode
	Size int
}

func NewPackMeFormatFromDirNode(ctx context.Context, node *dir.Node, size int) *PackMeFormat {
	packMeNode := NewPackMeNode(ctx, node)

	return &PackMeFormat{Node: packMeNode, Size: size}
}
