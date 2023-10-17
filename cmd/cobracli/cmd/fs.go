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
	Use:   "fs [flags] [path]",
	Short: "Find the directory with the most files.",
	Long:  `Find the directory with the most files in the file system.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ft := filesystem.FileToper{}
		if len(args) == 0 {
			ft.Root = "."
		} else {
			ft.Root = args[0]
		}
		ds, err := datatop.Top(n, &ft)
		if err != nil {
			return errors.Join(errors.New("fsCmd error"), err)
		}

		fs := (ds).([]filesystem.File)
		for _, f := range fs {
			fmt.Printf("%d\t\t%s\n", f.FileCount, f.Path)
		}
		return nil
	},
}
