package ucloud

import (
	"github.com/spf13/cobra"

	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client/ucloud"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
)

// CmdBindingHandler 创建绑定
func CmdBindingHandler(vServerId, ulbId, sslId string) func(*cobra.Command, []string) error {
	return func(command *cobra.Command, strings []string) error {
		var cli client.SSLClienter
		cli = ucloud.GetClient(setting.Conf.UCloudCredential, ucloud.WithProjectId(setting.Conf.ProjectId))
		//sslId := "ssl-5jtnrhjy"
		err := cli.SendBindingSSL(sslId, ulbId, vServerId)
		if err != nil {
			return err
		}
		return nil
	}
}