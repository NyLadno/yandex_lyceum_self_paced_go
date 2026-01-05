package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	in, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer in.Close()


	_, err = in.Seek(int64(startpos), 0)
	if err != nil {
		return err
	}

	out, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
