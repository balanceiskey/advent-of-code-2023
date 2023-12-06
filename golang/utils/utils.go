package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAndSplit(filename string) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("../../puzzle-inputs/%s", filename))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, err
}
