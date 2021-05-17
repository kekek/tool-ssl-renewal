package acme

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/burntsushi/toml"
	"wps.ktkt.com/monitor/tool-ssl-renewal/pkg/util/path"
)

type SSLConf struct {
	Le_Domain            string `json:"Le_Domain"`
	Le_CertCreateTime    string `json:"Le_CertCreateTime"`
	Le_CertCreateTimeStr string `json:"Le_CertCreateTimeStr"`
	Le_NextRenewTimeStr  string `json:"Le_NextRenewTimeStr"`
	Le_NextRenewTime     string `json:"Le_NextRenewTime"`
	Le_RealKeyPath       string `json:"Le_RealKeyPath"`
	Le_RealFullChainPath string `json:"Le_RealFullChainPath"`
}

func (c *SSLConf) ConvertToAcme() (*AcmeConf, error) {
	conf := &AcmeConf{}

	certCreateTime, err := parseStrUnixToTime(c.Le_CertCreateTime)
	if err != nil {
		return conf, err
	}

	nextRenewTime, err := parseStrUnixToTime(c.Le_NextRenewTime)
	if err != nil {
		return conf, err
	}

	conf.Le_CertCreateTime = certCreateTime
	conf.Le_NextRenewTime = nextRenewTime
	conf.Le_CertCreateTimeStr = c.Le_CertCreateTimeStr
	conf.Le_NextRenewTimeStr = c.Le_NextRenewTimeStr
	conf.Le_Domain = c.Le_Domain

	keyContent, _ := ioutil.ReadFile(c.Le_RealKeyPath)
	fullContent, _ := ioutil.ReadFile(c.Le_RealFullChainPath)

	conf.Le_RealFullChain = fullContent
	conf.Le_RealKey = keyContent

	return conf, nil
}

func CheckAndParseSSL(sslPath string) (*SSLConf, error) {

	// parse config file
	if !path.Exists(sslPath) || !path.IsFile(sslPath) {
		fmt.Errorf("ssl config file not exist : %s", sslPath)
	}

	ssl, err := ParseSSLConf(sslPath)
	if err != nil {
		log.Printf("ParseSSLConf failed : %v \n", err)
		return nil, err
	}

	return ssl, nil
}

func ParseSSLConf(path string) (*SSLConf, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("read conf file fialed : %v", err)
		return nil, err
	}

	ssl := &SSLConf{}

	err = toml.Unmarshal(data, ssl)
	if err != nil {
		log.Printf("parse config file fialed: %v", err)
		return nil, err
	}
	return ssl, nil
}

func parseStrUnixToTime(timestampStr string) (time.Time, error) {
	timestampInt64, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timestampInt64, 0), nil
}
