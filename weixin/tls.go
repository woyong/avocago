package weixin

import (
	"crypto/tls"
)

func NewTLSConfig(certPath string, keyPath string) (tlsConfig *tls.Config, err error) {
	tlsConfig = new(tls.Config)
	var cert tls.Certificate
	cert, err = tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return
	}
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	return
}
