package cmds

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func apiHandler(ctx context.Context, args []string, opts *apiOptions) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	modName, err := readCurrentModName(dir)
	if err != nil {
		return err
	}
	fileNotExistsAndCreate("./api")
	dir = dir + "/api"
	p = &project{
		ModName:     modName,
		ModuleName:  opts.Name,
		UModuleName: strings.ToUpper(string(opts.Name[0])) + opts.Name[1:],
	}
	err = p.createTemplateFile(fmt.Sprintf("%s/%s.go", dir, opts.Name), "tml/api/api.tml")
	if err != nil {
		return err
	}
	return nil
}