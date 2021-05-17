package setting

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/burntsushi/toml"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"wps.ktkt.com/monitor/tool-ssl-renewal/pkg/util/path"
)

type Config struct {
	ProjectId        string          // 项目projectID
	//VServerId        string          // 绑定操作的v_server id值
	//UlbId            string          // 绑定操作的 ulb_id
	SSLConfigPath    string          // acme 生成的域名配置路径
	UCloudCredential auth.Credential // u_cloud 的安全验证配置
}

var Conf *Config

func CheckAndParseConf(confPath string) error {
	// parse config file
	if !path.Exists(confPath) || !path.IsFile(confPath) {
		return fmt.Errorf("conf file not exist : %s", confPath)
	}

	err := ParseConf(confPath)
	if err != nil {
		log.Printf("ParseConf failed : %v \n", err)
		return err
	}
	log.Printf("配置文件解析完成  %#v \n", Conf)

	return nil
}

func ParseConf(confPath string) error {
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("read conf file fialed : %v", err)
		return err
	}

	Conf = &Config{}

	err = toml.Unmarshal(data, Conf)
	if err != nil {
		log.Printf("parse config file fialed: %v", err)
		return err
	}

	return nil

}
