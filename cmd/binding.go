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
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Fprint(os.Stdout, "sslId:", sslId, "\n")
		//fmt.Fprint(os.Stdout, "ulbId:", ulbId, "\n")
		//fmt.Fprint(os.Stdout, "vServerId:", vServerId, "\n")

		err := ucloud.CmdBindingHandler(vServerId, ulbId, sslId)
		if err != nil {
			return err
		}
		return nil
	},
	//RunE:  ucloud.CmdBindingHandler(vServerId, ulbId, sslId),
}

func init() {
	rootCmd.AddCommand(bindingCmd)

	bindingCmd.Flags().StringVarP(&vServerId, "vserver", "v", "", "id of vserver, must not empty")
	//bindingCmd.MarkPersistentFlagRequired("vServerId")

	bindingCmd.Flags().StringVarP(&ulbId, "ulb", "u", "", "id of ulb, must not empty")
	//bindingCmd.MarkPersistentFlagRequired("ulbId")

	bindingCmd.Flags().StringVarP(&sslId, "ssl", "s", "", "id of new ssl， must not empty")
	//bindingCmd.MarkPersistentFlagRequired("sslId")
}
