package cmds

import (
	"context"
	"fmt"
	"os"
	"strings"
)

var newApi []string = []string{
	"api",
	"schemas",
	"service",
}

func apiHandler(ctx context.Context, args []string, opts *apiOptions) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	modName, err := readCurrentModName(dir)
	if err != nil {
		return err
	}
	lowerCamelCaseName := toLowerCamelCase(opts.Name)
	p := &project{
		ModName:     modName,
		ModuleName:  lowerCamelCaseName,
		UModuleName: strings.ToUpper(string(lowerCamelCaseName[0])) + lowerCamelCaseName[1:],
	}
	for _, item := range newApi {
		fileNotExistsAndCreate("./" + item)
		filePath := fmt.Sprintf("%s/%s/%s.go", dir, item, opts.Name)
		p.FuncName = "Test"
		if item != "schemas" {
			err = p.createTemplateFile(filePath, fmt.Sprintf("tml/%s/%s.tml", item, item))
			if err != nil && opts.Func == "" {
				fmt.Println(err.Error())
				return err
			}
		} else {
			filePath = fmt.Sprintf("%s/%s/%s.go", dir, item, "schemas")
		}
		if opts.Func != "" {
			p.FuncName = strings.ToUpper(string(opts.Func[0])) + opts.Func[1:]
			var lineNumber, startLine int
			p.FuncText, lineNumber, startLine = getFuncParma(p.FuncName, item)
			err = p.createTemplateFunc(filePath, fmt.Sprintf("tml/%s/%s.tml", item, item), lineNumber, startLine)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getFuncParma(funcName, t string) ([]string, int, int) {
	var array []string
	var lineNumber, startLine int
	switch t {
	case "api":
		array = []string{"\t" + funcName + "(context *gin.Context)"}
		lineNumber = 11
		startLine = 19
	case "service":
		array = []string{fmt.Sprintf("\t%s(req *schemas.%sApiReq) (data interface{}, err IServiceError)", funcName, funcName)}
		lineNumber = 7
		startLine = 14
	case "schemas":
		lineNumber = 2
		startLine = 2
	default:
		// array = []string{"\n"}
	}
	return array, lineNumber, startLine
}
