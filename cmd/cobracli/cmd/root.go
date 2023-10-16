package cobracli

import (
	"os"

	"github.com/spf13/cobra"
)

var BuildVersion string

var rootCmd = &cobra.Command{
	Use:          "datatop",
	Version:      BuildVersion,
	SilenceUsage: true,
	Short:        "An open source tool for finding top data.",
	Long:         `An open source tool for finding top data. http://github.com/felixfwu/datatop`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
