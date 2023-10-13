package cmd

import (
	"errors"
	"fmt"

	"github.com/felixfwu/datatop"
	"github.com/felixfwu/datatop/filesystem"
	"github.com/spf13/cobra"
)

var n int

func init() {
	rootCmd.AddCommand(fsCmd)
	fsCmd.Flags().IntVarP(&n, "num", "n", 5, "Top N directory")
}

var fsCmd = &cobra.Command{
	Use:   "fs",
	Short: "Find the directory with the most files.",
	Long:  `Find the directory with the most files in the file system.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fs := filesystem.File{}
		if len(args) == 0 {
			fs.Root = "./"
		} else {
			fs.Root = args[0]
		}
		d, err := datatop.Top(n, &fs)
		if err != nil {
			return errors.Join(fmt.Errorf("fsCmd error: %s", err), err)
		}

		fds := (d).([]filesystem.Dir)
		for _, f := range fds {
			fmt.Printf("%s\t%d\n", f.Name, f.Cnt)
		}
		return nil
	},
}
