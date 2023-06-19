package main

import (
	sparse_decompresser "github.com/masahiro331/sparse-decompresser"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	GZIP = ".gz"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: sparse-decompresser [filename]")
	}
	ext := filepath.Ext(os.Args[1])
	switch ext {
	case GZIP:
		outputFile := strings.TrimSuffix(os.Args[1], ".gz")
		err := sparse_decompresser.DecompressSparseGzip(os.Args[1], outputFile)
		if err != nil {
			os.Remove(outputFile)
			log.Fatal(err)
		}
	default:
		log.Fatal("invalid file extension supported '.gz'")
	}
}
