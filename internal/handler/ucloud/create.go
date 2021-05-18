package ucloud

import (
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/acme"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client/ucloud"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
)

// CmdCreateHandler  常见ssl证书
func CmdCreateHandler(sslName string) (sslId string, err error) {
	sslId = ""
	sslconf, err := acme.CheckAndParseSSL(setting.Conf.SSLConfigPath)
	if err != nil {
		return sslId, err
	}

	acme, err := sslconf.ConvertToAcme()
	if err != nil {
		return sslId, err
	}

	var cli client.SSLClienter

	cli = ucloud.GetClient(setting.Conf.UCloudCredential, ucloud.WithProjectId(setting.Conf.ProjectId))

	if len(sslName) == 0 {
		sslName = acme.GetSSLName()
	}
	sslId, err = cli.SendCreateSSL(sslName, acme.GetKeyAndFullChain())
	if err != nil {
		return sslId, err
	}

	//fmt.Fprint(os.Stdout, sslId)
	return sslId, nil
}
