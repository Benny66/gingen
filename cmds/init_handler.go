package cmds

import (
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

var p = &project{}

type project struct {
	ModName     string
	ModuleName  string
	UModuleName string
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
	p = &project{
		ModName:     opts.ModName,
		ModuleName:  "user",
		UModuleName: "User",
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
