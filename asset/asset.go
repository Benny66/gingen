package asset

import (
	"bufio"
	"embed"
)

//go:embed tml
var f embed.FS

func ReadFile(name string) ([]byte, error) {
	return f.ReadFile(name)
}

func ReadDir(name string) ([]string, error) {
	fils, err := f.ReadDir(name)
	if err != nil {
		return []string{}, err
	}
	var res []string
	for _, f := range fils {
		res = append(res, f.Name())
	}
	return res, nil
}

func ReadLinesFromFile(path string, startLine int) (string, error) {
	file, err := f.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < startLine-1 && scanner.Scan(); i++ {
		// Skip lines until the desired starting line is reached
	}

	var lines string
	for scanner.Scan() {
		lines += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return lines, nil
}
