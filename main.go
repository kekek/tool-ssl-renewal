package main

import (
	"flag"
	"log"
	"os"

	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/acme"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/client"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
	"wps.ktkt.com/monitor/tool-ssl-renewal/version"
)

var printVersion = flag.Bool("v", false, "show build version for the program")
var confPath = flag.String("c", "data/run.toml", "path to conf that create by acme.sh")


func main() {
	flag.Parse()

	if *printVersion {
		version.PrintVersion()
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

	client := client.GetClient(setting.Conf.UCloudCredential, client.WithProjectId(setting.Conf.ProjectId))

	sslId, err := client.SendCreateSSL(acme.GetSSLName(), acme.GetKeyAndFullChain())
	if err != nil {
		log.Printf("SendCreateSSL failed : %v", err)
		panic(err)
	}

	err = client.SendBindingSSL(sslId, setting.Conf.UlbId, setting.Conf.VServerId)
	if err != nil {
		log.Printf("SendBindingSSL ssl failed : %v", err)
	}

	log.Printf("binding success! sslId :  %s \n", sslId)
}


