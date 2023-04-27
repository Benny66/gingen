package cmds

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Benny66/gingen/options"
)

var _ = os.Exit
var _ = fmt.Fprintf
var _ pflag.FlagSet
var _ = context.Background
var _ options.Options

type rootPersistentOptions struct {
	LogLevel string
}

var rootCmd = &cobra.Command{
	Use: "gingen",
}

var (
	AppName   = "gingen"
	Version   = "1.0.0"
	GitHash   = "1.0.0"
	BuildTime = "2023-05-01 00:00:00"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	persistentFlagSet := rootCmd.PersistentFlags()
	persistentFlagSet.SortFlags = false
	persistentFlagSet.String(
		"log-level",
		"info",
		"log level",
	)
}
