package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/masahiro331/spdc/pkg/spdc"
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
		err := spdc.DecompressSparseGzip(os.Args[1], outputFile)
		if err != nil {
			os.Remove(outputFile)
			log.Fatal(err)
		}
	default:
		log.Fatal("invalid file extension supported '.gz'")
	}
}