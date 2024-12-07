package pack_me_file_processing

func ReadPackMeFile(pathToFile string, format *PackMeFormat) {
	format.Node.Unpack(pathToFile)
}
