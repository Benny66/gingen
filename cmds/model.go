package cmds

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type modelOptions struct {
	flagSet *pflag.FlagSet
	rootPersistentOptions
	ModelName string
}

func unmarshalModelOptions(flagSet *pflag.FlagSet) *modelOptions {
	o := &modelOptions{flagSet: flagSet}
	modelName, err := flagSet.GetString("name")
	if err != nil || modelName == "" {
		log.Fatal("model is not set")
	}
	o.ModelName = modelName
	return o
}

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "create a model/dao module",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := unmarshalModelOptions(cmd.Flags())
		return modelHandler(context.Background(), args, opts)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
	flagSet := modelCmd.Flags()
	flagSet.SortFlags = false
	flagSet.String("name", "", "新增一个数据模型,例如：user,goods,coupon")
}
