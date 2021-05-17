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

// SendBindingSSL bind ssl
func (c *Client) SendBindingSSL(sslId, ulbId, vServerId string) error {
	req := c.NewBindSSLRequest()

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
