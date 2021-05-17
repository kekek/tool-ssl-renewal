package main

import (
	"flag"
	"log"
	"os"

	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/acme"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client/ucloud"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
	"wps.ktkt.com/monitor/tool-ssl-renewal/version"
)

var (
	printVersion = flag.Bool("v", false, "show build version for the program")
	confPath     = flag.String("c", "data/run.toml", "path to conf that create by acme.sh")
	ulbId        = flag.String("ulb_id", "", "ucloud ulb id, must not empty")
	vServerId    = flag.String("v_server_id", "", "ucloud vServer id, must not empty ")
)

func main() {
	flag.Parse()

	if *printVersion {
		version.PrintVersion()
		os.Exit(0)
	}

	if len(*vServerId) == 0 || len(*ulbId) == 0 {
		log.Printf("ulb_id 或 v_server_id 不能为空\n")
		flag.Usage()
		os.Exit(0)
	}

	err := setting.CheckAndParseConf(*confPath)
	if err != nil {
		log.Printf("setting.CheckAndParseConf failed : %v", err)
		panic(err)
	}

	sslconf, err := acme.CheckAndParseSSL(setting.Conf.SSLConfigPath)
	if err != nil {
		log.Printf("acme.CheckAndParseSSL failed : %v", err)
		panic(err)
	}

	acme, err := sslconf.ConvertToAcme()
	if err != nil {
		log.Printf("sslconf.ConvertToAcme failed : %v", err)
		panic(err)
	}

	log.Printf("\n acme conf : %#v \n", acme)

	log.Printf("get acme name :  %s \n", acme.GetSSLName())

	log.Printf("get acme content : \n %s \n", acme.GetKeyAndFullChain())

	var cli client.SSLClienter

	cli = ucloud.GetClient(setting.Conf.UCloudCredential, ucloud.WithProjectId(setting.Conf.ProjectId))

	sslId, err := cli.SendCreateSSL(acme.GetSSLName(), acme.GetKeyAndFullChain())
	if err != nil {
		log.Printf("SendCreateSSL failed : %v", err)
		panic(err)
	}
	//sslId := "ssl-5jtnrhjy"
	err = cli.SendBindingSSL(sslId, *ulbId, *vServerId)
	if err != nil {
		log.Printf("SendBindingSSL ssl failed : %v", err)
	}

	log.Printf("binding success! sslId :  %s \n", sslId)
}
