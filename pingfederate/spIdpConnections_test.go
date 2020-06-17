package pingfederate_test

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/spIdpConnections"
)

func TestSpIdpConnections(t *testing.T) {
	svc := spIdpConnections.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

	input1 := spIdpConnections.GetConnectionsInput{}
	result1, resp1, err1 := svc.GetConnections(&input1)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, 0, len(*result1.Items))

	input2 := spIdpConnections.CreateConnectionInput{
		Body: models.IdpConnection{
			Type:                   pingfederate.String("IDP"),
			Name:                   pingfederate.String("foo"),
			EntityId:               pingfederate.String("foo"),
			Active:                 pingfederate.Bool(true),
			ContactInfo:            &models.ContactInfo{},
			LoggingMode:            pingfederate.String("STANDARD"),
			LicenseConnectionGroup: pingfederate.String(""),
			Credentials: &models.ConnectionCredentials{
				Certs: &[]*models.ConnectionCert{},
				OutboundBackChannelAuth: &models.OutboundBackChannelAuth{
					Type:                pingfederate.String("OUTBOUND"),
					DigitalSignature:    pingfederate.Bool(false),
					ValidatePartnerCert: pingfederate.Bool(true),
				},
			},
			AttributeQuery: &models.IdpAttributeQuery{
				Url:          pingfederate.String("http://foo.com"),
				NameMappings: &[]*models.AttributeQueryNameMapping{},
				// Policy:       &IdpAttributeQueryPolicy{
				// 	SignAttributeQuery: ,
				// },
			},
		},
	}
	result2, resp2, err2 := svc.CreateConnection(&input2)
	equals(t, nil, err2)
	equals(t, 201, resp2.StatusCode)
	equals(t, true, *result2.Active)

	result3, resp3, err3 := svc.CreateConnection(&input2)
	equals(t, "Validation error(s) occurred. Please review the error(s) and address accordingly.\nThe Connection ID you specified is already in use.\nThe Connection Name you specified is already in use.", err3.Error())
	equals(t, 422, resp3.StatusCode)
	equals(t, (*models.IdpConnection)(nil), result3)
}
