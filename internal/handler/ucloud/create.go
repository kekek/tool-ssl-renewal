package ucloud

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/acme"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client/ucloud"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
)

// CmdCreateHandler  常见ssl证书
func CmdCreateHandler(sslName string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		sslconf, err := acme.CheckAndParseSSL(setting.Conf.SSLConfigPath)
		if err != nil {
			return err
		}

		acme, err := sslconf.ConvertToAcme()
		if err != nil {
			return err
		}

		var cli client.SSLClienter

		cli = ucloud.GetClient(setting.Conf.UCloudCredential, ucloud.WithProjectId(setting.Conf.ProjectId))

		if len(sslName) == 0 {
			sslName = acme.GetSSLName()
		}
		sslId, err := cli.SendCreateSSL(sslName, acme.GetKeyAndFullChain())
		if err != nil {
			return err
		}

		fmt.Fprint(os.Stdout, sslId)
		return nil
	}
}
