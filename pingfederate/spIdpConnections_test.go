package pingfederate

import (
	"testing"
)

func TestSpIdpConnections(t *testing.T) {
	svc := basicClient()

	input1 := GetConnectionsInput{}
	result1, resp1, err1 := svc.SpIdpConnections.GetConnections(&input1)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, 0, len(*result1.Items))

	input2 := CreateConnectionInput{
		Body: IdpConnection{
			Type:                   String("IDP"),
			Name:                   String("foo"),
			EntityId:               String("foo"),
			Active:                 Bool(true),
			ContactInfo:            &ContactInfo{},
			LoggingMode:            String("STANDARD"),
			LicenseConnectionGroup: String(""),
			Credentials: &ConnectionCredentials{
				Certs: &[]*ConnectionCert{},
				OutboundBackChannelAuth: &OutboundBackChannelAuth{
					Type:                String("OUTBOUND"),
					DigitalSignature:    Bool(false),
					ValidatePartnerCert: Bool(true),
				},
			},
			AttributeQuery: &IdpAttributeQuery{
				Url:          String("http://foo.com"),
				NameMappings: &[]*AttributeQueryNameMapping{},
				// Policy:       &IdpAttributeQueryPolicy{
				// 	SignAttributeQuery: ,
				// },
			},
		},
	}
	result2, resp2, err2 := svc.SpIdpConnections.CreateConnection(&input2)
	equals(t, nil, err2)
	equals(t, 201, resp2.StatusCode)
	equals(t, true, *result2.Active)

	result3, resp3, err3 := svc.SpIdpConnections.CreateConnection(&input2)
	equals(t, "Validation error(s) occurred. Please review the error(s) and address accordingly.\nThe Connection ID you specified is already in use.\nThe Connection Name you specified is already in use.", err3.Error())
	equals(t, 422, resp3.StatusCode)
	equals(t, (*IdpConnection)(nil), result3)
}
