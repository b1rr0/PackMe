package pack_me_file_processing

func ReadPackMeFile(pathToFile string, format *PackMeFormat) {
	for _, node := range format.Nodes {
		node.Unpack(pathToFile)
	}
}
