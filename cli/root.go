package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	//ossProvier string
	version bool
)
var RootCmd = &cobra.Command{
	Use:     "cloud-station-cli",
	Long:    "cloud-station-cli 云中转站",
	Short:   "cloud-station-cli 云中转站",
	Example: "cloud-station-cli cmds",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("cloud-station-cli v0.0.1")
		}
		return nil
	},
}

func init() {
	//RootCmd.PersistentFlags().StringVarP(&ossProvier, "provider", "p", "aliyun", "oss storage provier [aliyun/tencent/aws]")
	RootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "cloud station 版本信息")
}