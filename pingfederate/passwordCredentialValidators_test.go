package pingfederate

import (
	"testing"
)

func TestPasswordCredentialValidatorsDescriptors(t *testing.T) {
	svc := basicClient()

	result1, resp1, err1 := svc.PasswordCredentialValidators.GetPasswordCredentialValidatorDescriptors()
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, 5, len(*result1.Items))

	input2 := GetPasswordCredentialValidatorDescriptorInput{
		Id: "org.sourceid.saml20.domain.LDAPUsernamePasswordCredentialValidator",
	}
	result2, resp2, err2 := svc.PasswordCredentialValidators.GetPasswordCredentialValidatorDescriptor(&input2)
	equals(t, nil, err2)
	equals(t, 200, resp2.StatusCode)
	equals(t, "LDAP Username Password Credential Validator", *result2.Name)
}

func TestPasswordCredentialValidators(t *testing.T) {
	svc := basicClient()

	_, resp1, err1 := svc.PasswordCredentialValidators.GetPasswordCredentialValidators()
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)

	input2 := CreatePasswordCredentialValidatorInput{
		Body: PasswordCredentialValidator{
			Id:   String("pwdval2"),
			Name: String("pwdval2"),
			PluginDescriptorRef: &ResourceLink{
				Id:       String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
				Location: String("https://localhost:9999/pf-admin-api/v1/passwordCredentialValidators/descriptors/org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
			},
			Configuration: &PluginConfiguration{
				Tables: &[]*ConfigTable{
					&ConfigTable{
						Name: String("Users"),
						Rows: &[]*ConfigRow{
							&ConfigRow{
								Fields: &[]*ConfigField{
									&ConfigField{
										Name:  String("Username"),
										Value: String("demo"),
									},
									&ConfigField{
										Name:  String("Password"),
										Value: String("Demo1234"),
									},
									&ConfigField{
										Name:  String("Confirm Password"),
										Value: String("Demo1234"),
									},
								},
							},
						},
					},
				},
			},
			AttributeContract: &PasswordCredentialValidatorAttributeContract{
				CoreAttributes: &[]*PasswordCredentialValidatorAttribute{
					&PasswordCredentialValidatorAttribute{
						Name: String("username"),
					},
				},
			},
		},
	}

	result2, resp2, err2 := svc.PasswordCredentialValidators.CreatePasswordCredentialValidator(&input2)
	equals(t, nil, err2)
	equals(t, 201, resp2.StatusCode)
	equals(t, *result2.Name, "pwdval2")

	input3 := UpdatePasswordCredentialValidatorInput{
		Body: input2.Body,
		Id:   *input2.Body.Id,
	}
	(*(*(*input3.Body.Configuration.Tables)[0].Rows)[0].Fields)[0].Value = String("demo update")

	result3, resp3, err3 := svc.PasswordCredentialValidators.UpdatePasswordCredentialValidator(&input3)
	equals(t, nil, err3)
	equals(t, 200, resp3.StatusCode)
	equals(t, *result3.Name, "pwdval2")

	input4 := GetPasswordCredentialValidatorInput{
		Id: *input2.Body.Id,
	}
	result4, resp4, err4 := svc.PasswordCredentialValidators.GetPasswordCredentialValidator(&input4)
	equals(t, nil, err4)
	equals(t, 200, resp4.StatusCode)
	equals(t, *result4.Name, "pwdval2")

	input5 := DeletePasswordCredentialValidatorInput{
		Id: *input2.Body.Id,
	}
	result5, resp5, err5 := svc.PasswordCredentialValidators.DeletePasswordCredentialValidator(&input5)
	equals(t, nil, err5)
	equals(t, 204, resp5.StatusCode)
	equals(t, (*ApiResult)(nil), result5)
}
