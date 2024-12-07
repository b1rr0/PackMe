package compress

import (
	"bytes"
	"io/ioutil"
	"log"

	huffman "github.com/icza/huffman/hufio"
)

func Compress(data []byte) []byte {
	buf := &bytes.Buffer{}
	w := huffman.NewWriter(buf)
	if _, err := w.Write(data); err != nil {
		log.Panicln("Failed to write:", err)
	}
	if err := w.Close(); err != nil {
		log.Panicln("Failed to close:", err)
	}
	return buf.Bytes()
}

func Decompress(data []byte) []byte {
	r := huffman.NewReader(bytes.NewReader(data))
	if data, err := ioutil.ReadAll(r); err != nil {
		log.Panicln("Failed to read:", err)
	} else {
		log.Println("Read:", data)
	}
	return data
}
