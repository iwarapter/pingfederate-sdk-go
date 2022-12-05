package pingfederate

import (
	"encoding/json"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestWeCanMarshalAttributeSources(t *testing.T) {
	data := `[
        {
          "type": "LDAP",
          "dataStoreRef": {
            "id": "LDAP-0B01156DCB8373F211B5B6B10F04EC006178B620",
            "location": "https://localhost:9999/pf-admin-api/v1/dataStores/LDAP-0B01156DCB8373F211B5B6B10F04EC006178B620"
          },
          "id": "ldap",
          "description": "ldap",
          "baseDn": "",
          "searchScope": "SUBTREE",
          "searchFilter": "uid=${subject}",
          "searchAttributes": [
            "Subject DN"
          ],
          "binaryAttributeSettings": {},
          "memberOfNestedGroup": false
        },
        {
          "type": "JDBC",
          "dataStoreRef": {
            "id": "ProvisionerDS",
            "location": "https://localhost:9999/pf-admin-api/v1/dataStores/ProvisionerDS"
          },
          "id": "jdbc",
          "description": "jdbc",
          "schema": "INFORMATION_SCHEMA",
          "table": "ADMINISTRABLE_ROLE_AUTHORIZATIONS",
          "filter": "uid=${email}",
          "columnNames": [
            "GRANTEE",
            "IS_GRANTABLE"
          ]
        }
      ]`

	var desc []*models.AttributeSource
	err := json.Unmarshal([]byte(data), &desc)
	assert.NoError(t, err)

	assert.NotNil(t, desc[0].DataStoreRef)

	for _, source := range desc {
		switch *source.Type {
		case "LDAP":
			assert.Equal(t, *source.Description, *source.LdapAttributeSource.Description)
			assert.Equal(t, *source.Id, *source.LdapAttributeSource.Id)
			assert.Equal(t, *source.Type, *source.LdapAttributeSource.Type)
			assert.Equal(t, *source.DataStoreRef.Id, *source.LdapAttributeSource.DataStoreRef.Id)
		case "JDBC":
			assert.Equal(t, *source.Description, *source.JdbcAttributeSource.Description)
			assert.Equal(t, *source.Id, *source.JdbcAttributeSource.Id)
			assert.Equal(t, *source.Type, *source.JdbcAttributeSource.Type)
			assert.Equal(t, *source.DataStoreRef.Id, *source.JdbcAttributeSource.DataStoreRef.Id)
		case "CUSTOM":
			assert.Equal(t, *source.Description, *source.CustomAttributeSource.Description)
			assert.Equal(t, *source.Id, *source.CustomAttributeSource.Id)
			assert.Equal(t, *source.Type, *source.CustomAttributeSource.Type)
			assert.Equal(t, *source.DataStoreRef.Id, *source.CustomAttributeSource.DataStoreRef.Id)
		}
	}
}

func TestWeCanMarshalPluginDescriptors(t *testing.T) {
	radioGrpDesc := `[{
        "type": "RADIO_GROUP",
        "name": "Password Reset Type",
        "label": "Password Reset Type",
        "description": "Select the method to use for self-service password reset. Depending on the selected method, additional settings are required to complete the configuration. It is recommended that the method used is not already part of a multi-factor authentication policy that includes a password challenge, as that would indirectly reduce that authentication policy to a single factor. For example, if users normally authenticate with a password challenge and then PingID, the self-service password reset method should not be PingID.",
        "defaultValue": "NONE",
        "advanced": false,
        "required": false,
        "optionValues": [
          {
            "name": "Authentication Policy",
            "value": "POLICY"
          },
          {
            "name": "Email One-Time Link",
            "value": "OTL"
          },
          {
            "name": "Email One-Time Password",
            "value": "OTP"
          },
          {
            "name": "PingID",
            "value": "PingID"
          },
          {
            "name": "Text Message",
            "value": "SMS"
          },
          {
            "name": "None",
            "value": "NONE"
          }
        ]
      }]`

	desc := &[]*models.FieldDescriptor{}
	err := json.Unmarshal([]byte(radioGrpDesc), &desc)
	assert.NoError(t, err)

	require.NotNil(t, (*desc)[0].RadioGroupFieldDescriptor)
	require.NotNil(t, (*desc)[0].RadioGroupFieldDescriptor.OptionValues)
	assert.Len(t, *(*desc)[0].RadioGroupFieldDescriptor.OptionValues, 6)
}

func TestWeCanMarshalPolicyActions(t *testing.T) {
	str := `{
          "attributeRules": {
            "fallbackToSuccess": true,
            "items": [
              {
                "attributeName": "sub",
                "condition": "EQUALS",
                "expectedValue": "boo",
                "result": "Condition"
              }
            ]
          },
          "authenticationSource": {
            "sourceRef": {
              "id": "bart",
              "location": "https://localhost:9999/pf-admin-api/v1/idp/adapters/bart"
            },
            "type": "IDP_ADAPTER"
          },
          "type": "AUTHN_SOURCE"
        }`
	act := &models.PolicyAction{}
	err := json.Unmarshal([]byte(str), &act)
	assert.NoError(t, err)

	require.NotNil(t, act.AuthnSourcePolicyAction)
	require.NotNil(t, act.AuthnSourcePolicyAction.AttributeRules)

	b, err := json.Marshal(&models.PolicyAction{
		ApcMappingPolicyAction:           models.ApcMappingPolicyAction{},
		LocalIdentityMappingPolicyAction: models.LocalIdentityMappingPolicyAction{},
		AuthnSelectorPolicyAction:        models.AuthnSelectorPolicyAction{},
		AuthnSourcePolicyAction: models.AuthnSourcePolicyAction{
			AttributeRules: &models.AttributeRules{
				FallbackToSuccess: Bool(true),
				Items: &[]*models.AttributeRule{
					{
						AttributeName: String("foo1"),
						Condition:     String("foo2"),
						ExpectedValue: String("foo3"),
						Result:        String("foo4"),
					},
				},
			},
			AuthenticationSource: &models.AuthenticationSource{
				SourceRef: &models.ResourceLink{Id: String("foo")},
				Type:      String("IDP_ADAPTER"),
			},
		},
		ContinuePolicyAction: models.ContinuePolicyAction{},
		RestartPolicyAction:  models.RestartPolicyAction{},
		DonePolicyAction:     models.DonePolicyAction{},
		FragmentPolicyAction: models.FragmentPolicyAction{},
		Context:              String("foo"),
		Type:                 String("AUTHN_SOURCE"),
	})
	require.NoError(t, err)
	assert.True(t, strings.Contains(string(b), "attributeRules"))
	assert.True(t, strings.Contains(string(b), "type"))
	assert.True(t, strings.Contains(string(b), "context"))

	b, err = json.Marshal(&models.PolicyAction{
		Context: String("foo"),
		Type:    String("DONE"),
	})
	require.NoError(t, err)
	assert.True(t, strings.Contains(string(b), "type"))
	assert.True(t, strings.Contains(string(b), "context"))
}
