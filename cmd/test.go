package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
)

var testCmd = &cobra.Command{
	Use:                "test",
	Short:              "测试命令的运行情况 ",
	DisableSuggestions: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("test error")
		key := viper.GetString(keyName)
		fmt.Fprint(os.Stdout, keyName, ": ", key, "\n")

		fmt.Fprint(os.Stdout, *setting.Conf)

		return nil
	},
}

var keyName string

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.PersistentFlags().StringVarP(&keyName, "key", "k", "SSLConfigPath", "获取配置文件参数")
}
