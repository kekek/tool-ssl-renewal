package cmd

import (
	"github.com/spf13/cobra"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/handler/ucloud"
)

var (
	vServerId string
	ulbId     string
	sslId     string
)

var bindingCmd = &cobra.Command{
	Use:   "binding",
	Short: "给ucloud负载均衡下的虚拟服务，绑定证书",
	RunE:  ucloud.CmdBindingHandler(vServerId, ulbId, sslId),
}

func init() {
	rootCmd.AddCommand(bindingCmd)

	bindingCmd.PersistentFlags().StringVarP(&vServerId, "vServerId", "s", "", "ucloud 负载均衡下的vServerId, 不能为空")
	bindingCmd.MarkPersistentFlagRequired("vServerId")

	bindingCmd.PersistentFlags().StringVarP(&ulbId, "ulbId", "b", "", "ucloud 负载均衡ID, 不能为空")
	bindingCmd.MarkPersistentFlagRequired("ulbId")

	bindingCmd.PersistentFlags().StringVarP(&sslId, "sslId", "l", "", "待绑定的证书ID， 不能为空")
	bindingCmd.MarkPersistentFlagRequired("sslId")
}
