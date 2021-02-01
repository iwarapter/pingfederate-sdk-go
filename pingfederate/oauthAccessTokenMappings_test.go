package pingfederate_test

import (
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenMappings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
)

func TestOauthAccessTokenMappings(t *testing.T) {
	svc := oauthAccessTokenMappings.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

	result1, resp1, err1 := svc.GetMappings()
	require.Nil(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.Len(t, *result1, 0)

	//setup test data
	mgrs := oauthAccessTokenManagers.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))
	_, _, err := mgrs.CreateTokenManager(&oauthAccessTokenManagers.CreateTokenManagerInput{Body: models.AccessTokenManager{
		Id:   pf.String("foo"),
		Name: pf.String("foo"),
		PluginDescriptorRef: &models.ResourceLink{
			Id: pf.String("org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin"),
		},
		Configuration: &models.PluginConfiguration{},
		AttributeContract: &models.AccessTokenAttributeContract{
			ExtendedAttributes: &[]*models.AccessTokenAttribute{
				{
					Name: pf.String("foo"),
				},
			},
		},
	}})
	require.Nil(t, err)
	defer mgrs.DeleteTokenManager(&oauthAccessTokenManagers.DeleteTokenManagerInput{Id: "foo"})

	result2, resp2, err2 := svc.CreateMapping(&oauthAccessTokenMappings.CreateMappingInput{
		Body: models.AccessTokenMapping{
			AccessTokenManagerRef: &models.ResourceLink{Id: pf.String("foo")},
			AttributeContractFulfillment: map[string]*models.AttributeFulfillmentValue{
				"foo": {
					Source: &models.SourceTypeIdKey{
						Type: pf.String("CONTEXT"),
					},
					Value: pf.String("ClientId"),
				},
			},
		},
		BypassExternalValidation: pf.Bool(true),
	})
	require.Nil(t, err2)
	assert.Equal(t, http.StatusCreated, resp2.StatusCode)
	assert.NotNil(t, *result2.Id)

	result3, resp3, err3 := svc.UpdateMapping(&oauthAccessTokenMappings.UpdateMappingInput{
		Id: *result2.Id,
		Body: models.AccessTokenMapping{
			AccessTokenManagerRef: &models.ResourceLink{Id: pf.String("foo")},
			AttributeContractFulfillment: map[string]*models.AttributeFulfillmentValue{
				"foo": {
					Source: &models.SourceTypeIdKey{
						Type: pf.String("CONTEXT"),
					},
					Value: pf.String("ClientIp"),
				},
			},
		},
		BypassExternalValidation: pf.Bool(true),
	})
	require.Nil(t, err3)
	assert.Equal(t, http.StatusOK, resp3.StatusCode)
	assert.NotNil(t, *result3.Id)

	result4, resp4, err4 := svc.DeleteMapping(&oauthAccessTokenMappings.DeleteMappingInput{Id: *result3.Id})
	require.Nil(t, err4)
	assert.Equal(t, http.StatusNoContent, resp4.StatusCode)
	assert.Nil(t, result4)
}
