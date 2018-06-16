package main

import (
	"io"
	"os"
)

// CopyFile will receive a source and a destination and will copy
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	defer src.Close()
	if err != nil {
		return
	}

	dst, err := os.Open(dstName)
	defer dst.Close()
	if err != nil {
		return
	}

	return io.Copy(dst, src)
}
