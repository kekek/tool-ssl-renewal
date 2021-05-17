package acme

import (
	"fmt"
	"strings"
	"time"
)

type AcmeConf struct {
	Le_Domain            string
	Le_CertCreateTime    time.Time
	Le_CertCreateTimeStr string
	Le_NextRenewTimeStr  string
	Le_NextRenewTime     time.Time
	Le_RealKey           []byte
	Le_RealFullChain     []byte
}

// GetFullChain 返回证书内容
func (a *AcmeConf) GetFullChain() string {
	return string(a.Le_RealFullChain)
}

// GetKey 返回私钥内容
func (a *AcmeConf) GetKey() string {
	return string(a.Le_RealFullChain)
}

// GetKeyAndFullChain 返回合并私钥和证书内容
func (a *AcmeConf) GetKeyAndFullChain() string {
	st := strings.Builder{}
	st.Write(a.Le_RealKey)
	st.WriteByte('\n')
	st.Write(a.Le_RealFullChain)
	return st.String()
}

func (a *AcmeConf) GetSSLName() string {
	return fmt.Sprintf("%s-acme-%s", time.Now().Format("2006-01-02-15-04"), a.Le_Domain)
}
