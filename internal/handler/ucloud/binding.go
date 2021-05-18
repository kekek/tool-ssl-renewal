package ucloud

import (
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client/ucloud"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
)

// CmdBindingHandler 创建绑定
func CmdBindingHandler(vServerId, ulbId, sslId string) error {
	var cli client.SSLClienter
	cli = ucloud.GetClient(setting.Conf.UCloudCredential, ucloud.WithProjectId(setting.Conf.ProjectId))
	//sslId := "ssl-5jtnrhjy"
	err := cli.SendBindingSSL(sslId, ulbId, vServerId)
	if err != nil {
		return err
	}
	return nil
}
