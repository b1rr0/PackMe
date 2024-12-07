package dir_processing

func BuildDirTree(path string) (*Node, int) {
	fileCount := 0
	root := NewNode(path, &fileCount, nil)

	return root, fileCount
}
