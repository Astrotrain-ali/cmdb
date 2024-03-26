package cmd

import (
	"errors"
	"fmt"

	version "github.com/Astrotrain-ali/cmdb/version"
	"github.com/spf13/cobra"
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo-api 后端api",
	Long:  "demo-api 后端api",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return errors.New("no flags find")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print demo-api version")
}
