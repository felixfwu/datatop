package cobracli

import (
	"os"

	_ "embed"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "datatop",
	SilenceUsage: true,
	Short:        "An open source tool for finding top data.",
	Long:         `An open source tool for finding top data. http://github.com/felixfwu/datatop`,
}

func Execute(v string) {
	if v == "" {
		v = "unknow"
	}
	rootCmd.Version = v
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
