package pingfederate

import (
	"testing"
)

func TestOauthAuthServerSettings(t *testing.T) {
	svc := basicClient()

	result1, resp1, err1 := svc.OauthAuthServerSettings.GetAuthorizationServerSettings()
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, len(*result1.Scopes), 5)

	input2 := UpdateAuthorizationServerSettingsInput{
		Body: *result1,
	}
	input2.Body.AllowUnidentifiedClientROCreds = Bool(true)
	result2, resp2, err2 := svc.OauthAuthServerSettings.UpdateAuthorizationServerSettings(&input2)
	equals(t, nil, err2)
	equals(t, 200, resp2.StatusCode)
	equals(t, true, *result2.AllowUnidentifiedClientROCreds)
}
