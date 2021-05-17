package client

type SSLClienter interface {
	SendCreateSSL(name, content string) (sslId string, err error)
	SendBindingSSL(sslId, ulbId, vServerId string) error
}
