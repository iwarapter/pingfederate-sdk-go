package pingfederate

import (
	"encoding/json"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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
