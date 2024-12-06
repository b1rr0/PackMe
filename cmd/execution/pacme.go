package execution

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	dir "github.com/b1rr0/PackMe/dir-processing"
	processing "github.com/b1rr0/PackMe/pack-me-file-processing"
)

func Pack(pathsToDir []string, outputFileName string) {
	roots := []*dir.Node{}
	totalSize := 0
	for _, path := range pathsToDir {
		root, localSize := dir.BuildDirTree(path)
		roots = append(roots, root)
		totalSize += localSize
	}

	packMeFormat := processing.NewPackMeFormatFromDirNode(roots, totalSize)

	marshaled, err := json.Marshal(packMeFormat)
	if err != nil {
		fmt.Println("Error marshaling pack me format:", err.Error())
		return
	}

	if outputFileName == "" {
		outputFileName = time.Now().Format("2006-01-02") + ".PackMe"
	} else {
		outputFileName += ".PackMe"
	}

	err = os.WriteFile(outputFileName, marshaled, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err.Error())
	}
}

func Unpack(pathToArchive string) {
	content, err := os.ReadFile(pathToArchive)
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return
	}

	var packMeFormat processing.PackMeFormat
	err = json.Unmarshal(content, &packMeFormat)
	if err != nil {
		fmt.Println("Error unmarshaling pack me format:", err.Error())
		return
	}

	outputPath := strings.Join(strings.Split(pathToArchive, "/")[:len(strings.Split(pathToArchive, "/"))-1], "/")
	fmt.Println(outputPath)
	processing.ReadPackMeFile(outputPath, &packMeFormat)
}
