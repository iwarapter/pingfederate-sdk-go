package pingfederate

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/spIdpConnections"
)

func TestSpIdpConnections(t *testing.T) {
	svc := spIdpConnections.New("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)

	input1 := spIdpConnections.GetConnectionsInput{}
	result1, resp1, err1 := svc.GetConnections(&input1)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, 0, len(*result1.Items))

	input2 := spIdpConnections.CreateConnectionInput{
		Body: models.IdpConnection{
			Type:                   String("IDP"),
			Name:                   String("foo"),
			EntityId:               String("foo"),
			Active:                 Bool(true),
			ContactInfo:            &models.ContactInfo{},
			LoggingMode:            String("STANDARD"),
			LicenseConnectionGroup: String(""),
			Credentials: &models.ConnectionCredentials{
				Certs: &[]*models.ConnectionCert{},
				OutboundBackChannelAuth: &models.OutboundBackChannelAuth{
					Type:                String("OUTBOUND"),
					DigitalSignature:    Bool(false),
					ValidatePartnerCert: Bool(true),
				},
			},
			AttributeQuery: &models.IdpAttributeQuery{
				Url:          String("http://foo.com"),
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
