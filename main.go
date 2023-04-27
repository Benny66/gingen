package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type modelTemplate struct {
	Model  string
	XModel string
}

var modelData modelTemplate

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
	modelData.XModel = fmt.Sprintf("%s%s", strings.ToLower(string(model[0])), string(model[1:]))

}

func main() {
	var err error

	//获取当前文件夹绝对路径
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("获取当前文件夹失败", err)
	}

	filePath := fmt.Sprintf("%s/model/%s.go", dir, modelData.XModel)
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		fmt.Println("File is exist")
		return
	}
	modelByte, err := os.ReadFile("./tml/model.tml")
	if err != nil {
		log.Fatal(err.Error())
	}
	template := template.Must(template.New("").Parse(string(modelByte)))
	var b strings.Builder
	err = template.Execute(&b, modelData)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	n, err := file.Write([]byte(b.String()))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(n)
	return
}
