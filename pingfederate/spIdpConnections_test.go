package pingfederate_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/spIdpConnections"
)

func TestSpIdpConnections(t *testing.T) {
	svc := spIdpConnections.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))

	input1 := spIdpConnections.GetConnectionsInput{}
	result1, resp1, err1 := svc.GetConnections(&input1)
	require.Nil(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.Len(t, *result1.Items, 0)

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
	require.Nil(t, err2)
	assert.Equal(t, http.StatusCreated, resp2.StatusCode)
	assert.True(t, *result2.Active)

	result3, resp3, err3 := svc.CreateConnection(&input2)
	require.NotNil(t, err3)
	assert.Equal(t, "Validation error(s) occurred. Please review the error(s) and address accordingly.\nThe Connection ID you specified is already in use.\nThe Connection Name you specified is already in use.", err3.Error())
	assert.Equal(t, http.StatusUnprocessableEntity, resp3.StatusCode)
	assert.Nil(t, result3)

	result4, resp4, err4 := svc.DeleteConnection(&spIdpConnections.DeleteConnectionInput{Id: *result2.Id})
	require.Nil(t, err4)
	assert.Equal(t, http.StatusNoContent, resp4.StatusCode)
	assert.Nil(t, result4)
}
