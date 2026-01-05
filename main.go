package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	data, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var result []string
	scanner := bufio.NewScanner(data)
	layout := "02.01.2006"

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		datas := strings.Fields(line)

		currentTime, err := time.Parse(layout, datas[0])
		if err != nil {
			return nil, err
		}

		if currentTime.Equal(end) {
			result = append(result, line)
			break
		}

		if currentTime.After(start) || currentTime.Equal(start) {
			result = append(result, line)
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("No found logs")
	}

	return result, nil
}

func main() {
	exaple := "02.01.2006"
	start, _ := time.Parse(exaple, "19.12.2022")
	end, _ := time.Parse(exaple, "20.12.2022")

	fmt.Println(ExtractLog("test.txt", start, end))
}
