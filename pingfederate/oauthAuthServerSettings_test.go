package pingfederate

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthServerSettings"
)

func TestOauthAuthServerSettings(t *testing.T) {
	svc := oauthAuthServerSettings.New("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)

	result1, resp1, err1 := svc.GetAuthorizationServerSettings()
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, len(*result1.Scopes), 5)

	input2 := oauthAuthServerSettings.UpdateAuthorizationServerSettingsInput{
		Body: *result1,
	}
	input2.Body.AllowUnidentifiedClientROCreds = Bool(true)
	result2, resp2, err2 := svc.UpdateAuthorizationServerSettings(&input2)
	equals(t, nil, err2)
	equals(t, 200, resp2.StatusCode)
	equals(t, true, *result2.AllowUnidentifiedClientROCreds)
}
