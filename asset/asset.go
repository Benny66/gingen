package asset

import "embed"

//go:embed tml
var f embed.FS

func ReadFile(name string) ([]byte, error) {
	return f.ReadFile(name)
}
