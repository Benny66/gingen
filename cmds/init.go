package cmds

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Benny66/gingen/options"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var _ = os.Exit
var _ = fmt.Fprintf
var _ pflag.FlagSet
var _ = context.Background
var _ options.Options

type initOptions struct {
	flagSet *pflag.FlagSet
	rootPersistentOptions
	ModName string
}

func unmarshalInitOptions(flagSet *pflag.FlagSet) *initOptions {
	o := &initOptions{flagSet: flagSet}
	modName, err := flagSet.GetString("mod")
	if err != nil {
		log.Fatal(err)
	}
	if modName == "" {
		log.Fatal("mod name not specified")
	}
	o.ModName = modName
	return o
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a ginServer module",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := unmarshalInitOptions(cmd.Flags())
		return initHandler(context.Background(), args, opts)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	flagSet := initCmd.Flags()
	flagSet.SortFlags = false
	flagSet.String(
		"mod",
		"",
		"新增一个项目并命名初始化mod",
	)
}
