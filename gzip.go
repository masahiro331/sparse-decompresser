package spdc

import (
	"compress/gzip"
	"os"
)

func DecompressSparseGzip(src, dst string) error {
	w, err := os.Create(dst)

	if err != nil {
		return err
	}
	defer w.Close()

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	return DecompressSparseReader(gr, *w)
}