package acme

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}

func (s *ClientTestSuite) TestObtainSSL() {
	ctx := context.Background()
	client, err := NewRegisterAccount(ctx, "ci@haozi.net", CALetsEncryptStaging, nil, KeyEC256, slog.Default())
	s.Nil(err)

	client.UseDns(AliYun, DNSParam{
		AK: "123456",
		SK: "654321",
	})

	/*client.UseManualDns(2)

	resolves, err := client.GetDNSRecords(ctx, []string{"*.haozi.net", "haozi.net"}, KeyEC256)
	debug.Dump(resolves)
	s.Nil(err)
	s.NotNil(resolves)

	time.Sleep(2 * time.Minute)

	ssl, err := client.ObtainCertificateManual()*/
	ssl, err := client.ObtainCertificate(ctx, []string{"*.haozi.net", "haozi.net"}, KeyEC256)
	s.Error(err)
	s.NotNil(ssl)
}
