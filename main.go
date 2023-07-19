package main

import (
	"fmt"
	"github.com/ives22/cloud_station/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}