package pingfederate_test

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/services/passwordCredentialValidators"
)

func TestPasswordCredentialValidatorsDescriptors(t *testing.T) {
	svc := passwordCredentialValidators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

	result1, resp1, err1 := svc.GetPasswordCredentialValidatorDescriptors()
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, 5, len(*result1.Items))

	input2 := passwordCredentialValidators.GetPasswordCredentialValidatorDescriptorInput{
		Id: "org.sourceid.saml20.domain.LDAPUsernamePasswordCredentialValidator",
	}
	result2, resp2, err2 := svc.GetPasswordCredentialValidatorDescriptor(&input2)
	equals(t, nil, err2)
	equals(t, 200, resp2.StatusCode)
	equals(t, "LDAP Username Password Credential Validator", *result2.Name)
}

func TestPasswordCredentialValidators(t *testing.T) {
	svc := passwordCredentialValidators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

	_, resp1, err1 := svc.GetPasswordCredentialValidators()
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)

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
	equals(t, nil, err2)
	equals(t, 201, resp2.StatusCode)
	equals(t, *result2.Name, "pwdval2")

	input3 := passwordCredentialValidators.UpdatePasswordCredentialValidatorInput{
		Body: input2.Body,
		Id:   *input2.Body.Id,
	}
	(*(*(*input3.Body.Configuration.Tables)[0].Rows)[0].Fields)[0].Value = pingfederate.String("demo update")

	result3, resp3, err3 := svc.UpdatePasswordCredentialValidator(&input3)
	equals(t, nil, err3)
	equals(t, 200, resp3.StatusCode)
	equals(t, *result3.Name, "pwdval2")

	input4 := passwordCredentialValidators.GetPasswordCredentialValidatorInput{
		Id: *input2.Body.Id,
	}
	result4, resp4, err4 := svc.GetPasswordCredentialValidator(&input4)
	equals(t, nil, err4)
	equals(t, 200, resp4.StatusCode)
	equals(t, *result4.Name, "pwdval2")

	input5 := passwordCredentialValidators.DeletePasswordCredentialValidatorInput{
		Id: *input2.Body.Id,
	}
	result5, resp5, err5 := svc.DeletePasswordCredentialValidator(&input5)
	equals(t, nil, err5)
	equals(t, 204, resp5.StatusCode)
	equals(t, (*models.ApiResult)(nil), result5)
}
