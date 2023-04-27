package asset

import "embed"

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
