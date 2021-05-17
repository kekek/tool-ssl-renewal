package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Config *Conf

type Conf struct {
	Le_CertCreateTime    string `json:"Le_CertCreateTime"`
	Le_CertCreateTimeStr string `json:"Le_CertCreateTimeStr"`
	Le_NextRenewTimeStr  string `json:"Le_NextRenewTimeStr"`
	Le_NextRenewTime     string `json:"Le_NextRenewTime"`
	Le_RealKeyPath       string `json:"Le_RealKeyPath"`
	Le_RealFullChainPath string `json:"Le_RealFullChainPath"`
}

func ParseConf(confPath string) error {
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("read conf file fialed : %v", err)
		return err
	}

	Config = &Conf{}

	err = yaml.Unmarshal(data, Config)
	if err != nil {
		log.Printf("parse config file fialed: %v", err)
		return err
	}

	return nil
}
