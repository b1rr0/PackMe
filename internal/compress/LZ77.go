package compress

import (
	lz77 "github.com/urkx/LZ77"
)

func Lz77Compress(data []byte, windowSize int) []lz77.Result {
	compressedResults := lz77.Compress(string(data), windowSize)
	return compressedResults
}
func Lz77Decompress(data []lz77.Result) []byte {
	decompressed := lz77.Decompress(data)
	return []byte(decompressed)
}
