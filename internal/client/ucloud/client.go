package ucloud

import (
	"log"

	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type Client struct {
	*ulb.ULBClient
	copts options
}

func GetClient(conf auth.Credential, opts ...ClientOption) *Client {
	cfg := ucloud.NewConfig()
	cfg.Region = "cn-bj2"

	c := &Client{ULBClient: ulb.NewClient(&cfg, &conf)}
	for _, opt := range opts {
		opt.apply(&c.copts)
	}

	// replace the public/private key by your own

	return c
}

// SendCreateSSL create new ssl
func (c *Client) SendCreateSSL(name, content string) (sslId string, err error) {
	req := c.NewCreateSSLRequest()
	req.ProjectId = ucloud.String(c.copts.projectId)
	req.SSLContent = ucloud.String(content)
	req.SSLName = ucloud.String(name)
	req.Zone = ucloud.String("cn-bj2-05")

	// send request
	newSSL, err := c.CreateSSL(req)
	if err != nil {
		return "", err
	}

	return newSSL.SSLId, nil
}

// SendBindingSSL 先删除，在绑定
func (c *Client) SendBindingSSL(sslId, ulbId, vServerId string) error {

	// 查询旧sslID
	oldSslId, err := c.DoGetOldSSL(ulbId, vServerId)
	if err != nil {
		return err
	}

	if len(sslId) > 0 {
		// 删除
		err = c.DoDelOldSSL(oldSslId, ulbId, vServerId)
		if err != nil {
			return err
		}
	}

	// 绑定新的
	err = c.DoBindingSSL(sslId, ulbId, vServerId)
	if err != nil {
		return err
	}

	return nil
}

// DoGetOldSSL get old ssl id
func (c *Client) DoGetOldSSL(ulbId, vServerId string) (sslId string, err error) {

	sslId = ""

	req := c.NewDescribeVServerRequest()
	req.ProjectId = ucloud.String(c.copts.projectId)
	req.ULBId = ucloud.String(ulbId)
	req.VServerId = ucloud.String(vServerId)

	info, err := c.DescribeVServer(req)
	if err != nil {
		return sslId, err
	}

	if info != nil && len(info.DataSet) > 0 && len(info.DataSet[0].SSLSet) > 0 {
		sslId = info.DataSet[0].SSLSet[0].SSLId
	}

	return sslId, nil
}

// DoDelOldSSL delete old ssl id
func (c *Client) DoDelOldSSL(sslId, ulbId, vServerId string) error {
	req := c.NewUnbindSSLRequest()
	req.ProjectId = ucloud.String(c.copts.projectId)
	req.ULBId = ucloud.String(ulbId)
	req.VServerId = ucloud.String(vServerId)
	req.SSLId = ucloud.String(sslId)

	_, err := c.UnbindSSL(req)
	if err != nil {
		return err
	}

	return nil
}

// DoBindingSSL bind ssl
func (c *Client) DoBindingSSL(sslId, ulbId, vServerId string) error {
	req := c.NewBindSSLRequest()
	req.ProjectId = ucloud.String(c.copts.projectId)

	req.SSLId = ucloud.String(sslId)
	req.ULBId = ucloud.String(ulbId)
	req.VServerId = ucloud.String(vServerId)

	res, err := c.BindSSL(req)
	if err != nil {
		return err
	}

	log.Printf("bind success : %#v", res)

	return nil
}
