package setting

import (
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/burntsushi/toml"
)

var Config *Conf

type Conf struct {
	Le_Domain            string `json:"Le_Domain"`
	Le_CertCreateTime    string `json:"Le_CertCreateTime"`
	Le_CertCreateTimeStr string `json:"Le_CertCreateTimeStr"`
	Le_NextRenewTimeStr  string `json:"Le_NextRenewTimeStr"`
	Le_NextRenewTime     string `json:"Le_NextRenewTime"`
	Le_RealKeyPath       string `json:"Le_RealKeyPath"`
	Le_RealFullChainPath string `json:"Le_RealFullChainPath"`
}

func (c *Conf) Convert2SSLConf() (*SSLConf, error) {
	conf := &SSLConf{}

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

type SSLConf struct {
	Le_Domain            string
	Le_CertCreateTime    time.Time
	Le_CertCreateTimeStr string
	Le_NextRenewTimeStr  string
	Le_NextRenewTime     time.Time
	Le_RealKey           []byte
	Le_RealFullChain     []byte
}

func ParseConf(confPath string) error {
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("read conf file fialed : %v", err)
		return err
	}

	Config = &Conf{}

	err = toml.Unmarshal(data, Config)
	if err != nil {
		log.Printf("parse config file fialed: %v", err)
		return err
	}

	return nil
}

func parseStrUnixToTime(timestampStr string) (time.Time, error) {
	timestampInt64, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timestampInt64, 0), nil
}
