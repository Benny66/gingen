package cmds

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type apiOptions struct {
	flagSet *pflag.FlagSet
	rootPersistentOptions
	Name string
	Func string
}

func unmarshalApiOptions(flagSet *pflag.FlagSet) *apiOptions {
	o := &apiOptions{flagSet: flagSet}
	moduleName, err := flagSet.GetString("name")
	if err != nil || moduleName == "" {
		log.Fatal("api name is not set")
	}
	funcName, _ := flagSet.GetString("func")
	o.Name = moduleName
	o.Func = toUpCamelCase(funcName)
	return o
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "create a api module",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := unmarshalApiOptions(cmd.Flags())
		return apiHandler(context.Background(), args, opts)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	flagSet := apiCmd.Flags()
	flagSet.SortFlags = false
	flagSet.String("name", "", "新增一个api接口模块，例如：user,goods,coupon")
	flagSet.String("func", "", "新增一个函数方法，例如：user,goods,coupon")
}
