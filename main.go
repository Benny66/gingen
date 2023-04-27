package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/Benny66/gingen/asset"
)

type modelTemplate struct {
	Type   string
	Model  string
	XModel string
}

var modelData modelTemplate

// 初始化参数
// model文件
func init() {
	var model string
	flag.StringVar(&model, "model", "", "model name")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if model == "" {
		log.Fatal("Model is empty")
	}

	modelData.Model = model
	modelData.XModel = fmt.Sprintf("%s%s", strings.ToUpper(string(model[0])), string(model[1:]))

}

func main() {
	var err error

	//获取当前文件夹绝对路径
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("获取当前文件夹失败", err)
	}
	modelData.Type = "dao"
	err = modelData.createModelTemplateFile(dir, "tml/dao.tml")
	if err != nil {
		log.Fatal(err)
	}
	modelData.Type = "model"
	err = modelData.createModelTemplateFile(dir, "tml/model.tml")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *modelTemplate) createModelTemplateFile(dirPath string, tmlPath string) error {
	var filePath string = fmt.Sprintf("%s/models/%s.go", dirPath, m.Model)
	if m.Type == "model" {
		filePath = fmt.Sprintf("%s/models/%s_model.go", dirPath, m.Model)
	}
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return err
	}
	modelByte, err := asset.ReadFile(tmlPath)
	if err != nil {
		fmt.Println(1111, tmlPath)
		return err
	}
	template := template.Must(template.New("").Parse(string(modelByte)))
	var b strings.Builder
	err = template.Execute(&b, modelData)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(b.String()))
	if err != nil {
		return err
	}
	return nil
}
