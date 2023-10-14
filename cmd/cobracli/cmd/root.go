package cobracli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "datatop",
	Short: "An open source tool for finding top data.",
	Long:  `An open source tool for finding top data. http://github.com/felixfwu/datatop`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
