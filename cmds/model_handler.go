package cmds

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func modelHandler(ctx context.Context, args []string, opts *modelOptions) error {
	fmt.Println(opts)
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	modName, err := readCurrentModName(dir)
	if err != nil {
		return err
	}
	fileNotExistsAndCreate("./models")
	dir = dir + "/models"
	p = &project{
		ModName:     modName,
		ModuleName:  opts.Name,
		UModuleName: strings.ToUpper(string(opts.Name[0])) + opts.Name[1:],
	}
	err = p.createTemplateFile(fmt.Sprintf("%s/%s.go", dir, opts.Name), "tml/models/dao.tml")
	if err != nil {
		return err
	}
	err = p.createTemplateFile(fmt.Sprintf("%s/%s_model.go", dir, opts.Name), "tml/models/model.tml")
	if err != nil {
		return err
	}
	return nil
}

// 读取当前go.mod的文件mod名称
func readCurrentModName(dir string) (string, error) {
	modFile, err := os.Open(dir + "/go.mod")
	if err != nil {
		return "", err
	}
	defer modFile.Close()
	scanner := bufio.NewScanner(modFile)
	var firstLine string
	if scanner.Scan() {
		firstLine = scanner.Text()
	}
	//字符串替换掉module
	firstLine = strings.Replace(firstLine, "module ", "", 1)
	return firstLine, nil
}
