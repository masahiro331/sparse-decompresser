package spdc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"golang.org/x/sync/errgroup"
)

const (
	blockSize     = 4096
	maxGoroutines = 64
)

// DecompressSparseGzipReader decompress sparse file
func DecompressSparseReader(src io.Reader, dst os.File) error {
	buf := make([]byte, blockSize)

	var size int
	ctx := context.Background()

	// Use errgroup to manage goroutine and capture the first error returned by goroutines.
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(maxGoroutines)

	for {
		n, err := src.Read(buf)
		if n == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		size += n

		err = dst.Truncate(int64(size))
		if err != nil {
			return err
		}

		// Copy buffer to avoid race conditions.
		b := make([]byte, n)
		copy(b, buf[:n])

		// Copy the current size to avoid race conditions.
		currentSize := size

		g.Go(func() error {

			if !bytes.Equal(b, make([]byte, len(b))) {
				if _, err := dst.WriteAt(b, int64(currentSize-n)); err != nil {
					return err
				}
			}
			return nil
		})
	}

	// If there is an error, return it. The errgroup package ensures that only the first error is returned.
	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to decompress gzip file: %w", err)
	}

	return nil
}