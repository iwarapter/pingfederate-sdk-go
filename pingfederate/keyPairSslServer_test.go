package pingfederate_test

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"
)

func TestKeyPairSslServer(t *testing.T) {
	svc := keyPairsSslServer.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))

	result1, resp1, err1 := svc.CreateKeyPair(&keyPairsSslServer.CreateKeyPairInput{Body: models.NewKeyPairSettings{
		City:             pingfederate.String("Test"),
		CommonName:       pingfederate.String("Test"),
		Country:          pingfederate.String("GB"),
		KeyAlgorithm:     pingfederate.String("RSA"),
		KeySize:          pingfederate.Int(2048),
		Organization:     pingfederate.String("Test"),
		OrganizationUnit: pingfederate.String("Test"),
		State:            pingfederate.String("Test"),
		ValidDays:        pingfederate.Int(365),
	}})
	equals(t, nil, err1)
	equals(t, http.StatusCreated, resp1.StatusCode)

	result2, resp2, err2 := svc.ExportCsr(&keyPairsSslServer.ExportCsrInput{
		Id: *result1.Id,
	})
	equals(t, nil, err2)
	equals(t, http.StatusOK, resp2.StatusCode)
	b, _ := pem.Decode([]byte(*result2))
	_, err := x509.ParseCertificateRequest(b.Bytes)
	ok(t, err)
}
