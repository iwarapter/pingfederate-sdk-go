package pingfederate_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/services/passwordCredentialValidators"
)

func TestPasswordCredentialValidatorsDescriptors(t *testing.T) {
	svc := passwordCredentialValidators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))

	result1, resp1, err1 := svc.GetPasswordCredentialValidatorDescriptors()
	require.Nil(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.True(t, 5 <= len(*result1.Items))

	input2 := passwordCredentialValidators.GetPasswordCredentialValidatorDescriptorInput{
		Id: "org.sourceid.saml20.domain.LDAPUsernamePasswordCredentialValidator",
	}
	result2, resp2, err2 := svc.GetPasswordCredentialValidatorDescriptor(&input2)
	require.Nil(t, err2)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)
	assert.Equal(t, "LDAP Username Password Credential Validator", *result2.Name)
}

func TestPasswordCredentialValidators(t *testing.T) {
	svc := passwordCredentialValidators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))

	_, resp1, err1 := svc.GetPasswordCredentialValidators()
	require.Nil(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)

	input2 := passwordCredentialValidators.CreatePasswordCredentialValidatorInput{
		Body: models.PasswordCredentialValidator{
			Id:   pingfederate.String("pwdval2"),
			Name: pingfederate.String("pwdval2"),
			PluginDescriptorRef: &models.ResourceLink{
				Id:       pingfederate.String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
				Location: pingfederate.String("https://localhost:9999/pf-admin-api/v1/passwordCredentialValidators/descriptors/org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
			},
			Configuration: &models.PluginConfiguration{
				Tables: &[]*models.ConfigTable{
					{
						Name: pingfederate.String("Users"),
						Rows: &[]*models.ConfigRow{
							{
								Fields: &[]*models.ConfigField{
									{
										Name:  pingfederate.String("Username"),
										Value: pingfederate.String("demo"),
									},
									{
										Name:  pingfederate.String("Password"),
										Value: pingfederate.String("Demo1234"),
									},
									{
										Name:  pingfederate.String("Confirm Password"),
										Value: pingfederate.String("Demo1234"),
									},
								},
							},
						},
					},
				},
			},
			AttributeContract: &models.PasswordCredentialValidatorAttributeContract{
				CoreAttributes: &[]*models.PasswordCredentialValidatorAttribute{
					{
						Name: pingfederate.String("username"),
					},
				},
			},
		},
	}

	result2, resp2, err2 := svc.CreatePasswordCredentialValidator(&input2)
	require.Nil(t, err2)
	assert.Equal(t, http.StatusCreated, resp2.StatusCode)
	assert.Equal(t, "pwdval2", *result2.Name)

	input3 := passwordCredentialValidators.UpdatePasswordCredentialValidatorInput{
		Body: input2.Body,
		Id:   *input2.Body.Id,
	}
	(*(*(*input3.Body.Configuration.Tables)[0].Rows)[0].Fields)[0].Value = pingfederate.String("demo update")

	result3, resp3, err3 := svc.UpdatePasswordCredentialValidator(&input3)
	require.Nil(t, err3)
	assert.Equal(t, http.StatusOK, resp3.StatusCode)
	assert.Equal(t, "pwdval2", *result3.Name)
	assert.Equal(t, "demo update", *(*(*(*result3.Configuration.Tables)[0].Rows)[0].Fields)[0].Value)

	input4 := passwordCredentialValidators.GetPasswordCredentialValidatorInput{
		Id: *input2.Body.Id,
	}
	result4, resp4, err4 := svc.GetPasswordCredentialValidator(&input4)
	require.Nil(t, err4)
	assert.Equal(t, http.StatusOK, resp4.StatusCode)
	assert.Equal(t, "pwdval2", *result4.Name)

	input5 := passwordCredentialValidators.DeletePasswordCredentialValidatorInput{
		Id: *input2.Body.Id,
	}
	result5, resp5, err5 := svc.DeletePasswordCredentialValidator(&input5)
	require.Nil(t, err5)
	assert.Equal(t, http.StatusNoContent, resp5.StatusCode)
	assert.Nil(t, result5)
}
