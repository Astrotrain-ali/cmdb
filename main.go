package main

import (
	"fmt"

	"github.com/Astrotrain-ali/cmdb/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
