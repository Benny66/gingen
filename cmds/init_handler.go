package cmds

import (
	"bufio"
	"context"
	"fmt"
	"html"
	"html/template"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Benny66/gingen/asset"
)

// var p ProjectInterface = &project{}

type project struct {
	ModName     string
	ModuleName  string
	UModuleName string
	FuncName    string
	FuncText    []string
}

var modules []string = []string{
	"api",
	"format",
	"config",
	"db",
	"log",
	"middleware",
	"middleware",
	"models",
	"routers",
	"schemas",
	"service",
	"main",
}

func initHandler(ctx context.Context, args []string, opts *initOptions) error {
	p := &project{
		ModName:     opts.ModName,
		ModuleName:  "user",
		UModuleName: "User",
	}
	if p.FuncName == "" {
		p.FuncName = "Test"
	}
	var sync sync.WaitGroup
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	for _, module := range modules {
		sync.Add(1)
		go func(module string) {
			defer sync.Done()
			modulePath := dir + "/" + module
			if module != "main" {
				fileNotExistsAndCreate(modulePath)
				files, err := asset.ReadDir(fmt.Sprintf("tml/%s", module))
				if err != nil {
					log.Fatal(err)
				}
				for _, file := range files {
					array := strings.Split(file, ".")
					err = p.createTemplateFile(fmt.Sprintf("%s/%s.go", modulePath, array[0]), fmt.Sprintf("tml/%s/%s", module, file))
					if err != nil {
						log.Fatal("创建文件异常", err)
					}
					fmt.Println("create file success", file)
				}
			} else {
				err := p.createTemplateFile(fmt.Sprintf("%s/%s.go", dir, "main"), fmt.Sprintf("tml/%s", "main.tml"))
				if err != nil {
					log.Fatal("创建文件异常", err)
				}
				err = p.createTemplateFile(fmt.Sprintf("%s/go.mod", dir), fmt.Sprintf("tml/%s", "mod.tml"))
				if err != nil {
					log.Fatal("创建文件异常", err)
				}
				fmt.Println("create main.go go.mod success")
			}

		}(module)
	}
	sync.Wait()
	return nil
}

func fileNotExistsAndCreate(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}

type ProjectInterface interface {
	createTemplateFile(filePath, tmlPath string) error
	createTemplateFunc(filePath, tmlPath string, startLine int) error
}

func (p *project) createTemplateFile(filePath, tmlPath string) error {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return err
	}
	modelByte, err := asset.ReadFile(tmlPath)
	if err != nil {
		return err
	}
	template := template.Must(template.New("").Parse(string(modelByte)))
	var b strings.Builder
	err = template.Execute(&b, p)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(html.UnescapeString(b.String())))
	if err != nil {
		return err
	}
	return nil
}

func (p *project) createTemplateFunc(filePath, tmlPath string, lineNumber, startLine int) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return err
	}
	fileBytes, err := asset.ReadLinesFromFile(tmlPath, startLine)
	if err != nil {
		return err
	}

	template := template.Must(template.New("").Parse(string(fileBytes)))
	var b strings.Builder
	err = template.Execute(&b, p)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// 创建一个 scanner，用于逐行读取文件内容
	scanner := bufio.NewScanner(file)

	// 逐行读取文件并存储到一个字符串切片中
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// 利用存储的字符串切片，在指定行后插入新行
	//{{.FuncName}}(context *gin.Context)
	lines = append(lines[:lineNumber+1], append(p.FuncText, lines[lineNumber+1:]...)...)

	// 将修改后的字符串切片写回到文件中
	file.Truncate(0)
	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	_, err = writer.WriteString(b.String() + "\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}

func toLowerCamelCase(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 1 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	for i := range parts {
		if i != 0 {
			parts[i] = strings.Title(parts[i])
		}
	}
	return strings.Join(parts, "")
}

func toUpCamelCase(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 1 {
		return strings.Title(parts[0])
	}
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
