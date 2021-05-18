package cmd

import (
	"github.com/spf13/cobra"

	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/handler/ucloud"
)

var sslName string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a ssl certificate to ULB",
	RunE:  ucloud.CmdCreateHandler(sslName),
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&sslName, "name", "", "name of the ssl certificate to create, if empty, it will be auto generated")
}
