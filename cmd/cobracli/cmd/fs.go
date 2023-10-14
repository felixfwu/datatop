package cobracli

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
		fs := filesystem.FileSystem{}
		if len(args) == 0 {
			fs.Root = "."
		} else {
			fs.Root = args[0]
		}
		ds, err := datatop.Top(n, &fs)
		if err != nil {
			return errors.Join(errors.New("fsCmd error"), err)
		}

		fds := (ds).([]filesystem.Dir)
		for _, f := range fds {
			fmt.Printf("%d\t\t%s\n", f.FileCount, f.Path)
		}
		return nil
	},
}
