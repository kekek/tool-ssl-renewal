package main

import (
	"flag"
	"log"
	"os"

	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
	"wps.ktkt.com/monitor/tool-ssl-renewal/pkg/util/path"
	"wps.ktkt.com/monitor/tool-ssl-renewal/version"
)

var printVersion = flag.Bool("v", false, "show build version for the program")
var confPath = flag.String("c", "ilaiyang.conf", "path to conf that create by acme.sh")



func main() {
	flag.Parse()

	if *printVersion {
		version.PrintVersion()
		os.Exit(0)
	}

	// parse config file
	if !path.Exists(*confPath) || !path.IsFile(*confPath) {
		log.Fatalf("conf file not exist : %s", *confPath)
	}

	err := setting.ParseConf(*confPath)
	if err != nil {
		log.Fatalf("ParseConf failed : %v", err)
	}

	log.Printf("配置文件解析完成  %#v", setting.Config)

	log.Print("start create ucloud ssl conf")

}
