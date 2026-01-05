package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
    defer f.Close()
	if err != nil {
		return ""
	}
	
	scanner := bufio.NewScanner(f)
	curentLine := 0
	var result string

	for scanner.Scan(){
		if curentLine == lineNum{
			result =  scanner.Text()
			break
		}
		curentLine ++
	}

	return result
}