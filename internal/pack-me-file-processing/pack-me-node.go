package pack_me_file_processing

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	dir "PackMe/internal/dir-processing"

	flate "compress/flate"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type PackMeNode struct {
	Name     string
	Children *[]PackMeNode
	Data     []byte
}

func NewPackMeNode(ctx context.Context, node *dir.Node) *PackMeNode {
	parts := strings.Split(node.Name, "/")
	packMeNode := &PackMeNode{Name: parts[len(parts)-1]}

	if children, ok := node.Data.([]dir.Node); ok {
		packMeNode.Children = &[]PackMeNode{}
		for _, child := range children {
			*packMeNode.Children = append(*packMeNode.Children, *NewPackMeNode(ctx, &child))
		}
	} else {
		runtime.EventsEmit(ctx, "app:nodeProcessed", node.Name)
		packMeNode.Data, _ = compressToByteArray(node.Data.([]byte))
	}
	return packMeNode
}

func compressToByteArray(data []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer, err := flate.NewWriter(&buffer, flate.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = writer.Write(data)
	if err != nil {
		writer.Close()
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func decompressToByteArray(compressedData []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressedData))
	defer reader.Close()

	uncompressedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return uncompressedData, nil
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

		decompressedData, err := decompressToByteArray(node.Data)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		}
		_, err = file.Write(decompressedData)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		}
	}
}
