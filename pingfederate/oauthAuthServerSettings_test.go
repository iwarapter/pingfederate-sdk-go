package pingfederate_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthServerSettings"
)

func TestOauthAuthServerSettings(t *testing.T) {
	svc := oauthAuthServerSettings.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))

	result1, resp1, err1 := svc.GetAuthorizationServerSettings()
	require.Nil(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.Len(t, *result1.Scopes, 0)

	input2 := oauthAuthServerSettings.UpdateAuthorizationServerSettingsInput{
		Body: *result1,
	}
	input2.Body.AllowUnidentifiedClientROCreds = pingfederate.Bool(true)
	result2, resp2, err2 := svc.UpdateAuthorizationServerSettings(&input2)
	require.Nil(t, err2)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)
	assert.True(t, *result2.AllowUnidentifiedClientROCreds)
}
