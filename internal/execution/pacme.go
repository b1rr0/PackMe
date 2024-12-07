package execution

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	dir "PackMe/internal/dir-processing"
	processing "PackMe/internal/pack-me-file-processing"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Pack(ctx context.Context, pathsToDir string) string {
	root, localSize := dir.BuildDirTree(pathsToDir)
	runtime.EventsEmit(ctx, "app:nodeToProcess", localSize)

	packMeFormat := processing.NewPackMeFormatFromDirNode(ctx, root, localSize)

	marshaled, err := json.Marshal(packMeFormat)
	if err != nil {
		fmt.Println("Error marshaling pack me format:", err.Error())
		return ""
	}
	outputFileName := pathsToDir + ".PackMe"

	err = os.WriteFile(outputFileName, marshaled, 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err.Error())
	}
	return outputFileName
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
