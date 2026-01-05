package main

import (
	"io"
	"os"
)

func ReadContent(filename string) string {
	f, err := os.Open(filename)
	if err != nil{
		return ""
	}

	data, errRead := io.ReadAll(f)
	if errRead != nil{
		return ""
	}

	return string(data)
}


