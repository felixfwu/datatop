package cobracli

import (
	"os"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed VERSION
var version string

var rootCmd = &cobra.Command{
	Use:          "datatop",
	Version:      version,
	SilenceUsage: true,
	Short:        "An open source tool for finding top data.",
	Long:         `An open source tool for finding top data. http://github.com/felixfwu/datatop`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
