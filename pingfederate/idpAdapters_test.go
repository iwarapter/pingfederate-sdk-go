package pingfederate

import (
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"
)

func TestIdpAdapters(t *testing.T) {
	svc := idpAdapters.New("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)

	input := idpAdapters.GetIdpAdaptersInput{}
	_, resp1, err1 := svc.GetIdpAdapters(&input)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	//equals(t, *(*result1.Items)[0].Name, "woot")

	input2 := idpAdapters.CreateIdpAdapterInput{
		Body: models.IdpAdapter{
			AttributeContract: &models.IdpAdapterAttributeContract{
				CoreAttributes: &[]*models.IdpAdapterAttribute{
					&models.IdpAdapterAttribute{
						Masked:    Bool(false),
						Name:      String("subject"),
						Pseudonym: Bool(true),
					},
				},
				ExtendedAttributes: &[]*models.IdpAdapterAttribute{},
				MaskOgnlValues:     Bool(false),
			},
			AttributeMapping: &models.IdpAdapterContractMapping{
				AttributeContractFulfillment: map[string]*models.AttributeFulfillmentValue{
					"subject": &models.AttributeFulfillmentValue{
						Source: &models.SourceTypeIdKey{
							Type: String("ADAPTER"),
						},
						Value: String("subject"),
					},
				},
				IssuanceCriteria: &models.IssuanceCriteria{
					ConditionalCriteria: &[]*models.ConditionalIssuanceCriteriaEntry{},
				},
			},
			Configuration: &models.PluginConfiguration{
				Fields: &[]*models.ConfigField{
					&models.ConfigField{
						Name:  String("Password"),
						Value: String("Password01"),
					},
					&models.ConfigField{
						Name:  String("Confirm Password"),
						Value: String("Password01"),
					},
					&models.ConfigField{
						Name:  String("Authentication Service"),
						Value: String("http://foo.com"),
					},
				},
			},
			Id:   String("woot"),
			Name: String("woot"),
			PluginDescriptorRef: &models.ResourceLink{
				Id: String("com.pingidentity.adapters.opentoken.IdpAuthnAdapter"),
			},
		},
	}

	result2, resp2, err2 := svc.CreateIdpAdapter(&input2)
	equals(t, nil, err2)
	equals(t, 201, resp2.StatusCode)
	equals(t, *result2.Name, "woot")

	input3 := idpAdapters.UpdateIdpAdapterInput{
		Body: input2.Body,
		Id:   *input2.Body.Id,
	}
	(*input3.Body.Configuration.Fields)[0].Value = String("Password02")
	(*input3.Body.Configuration.Fields)[1].Value = String("Password02")

	result3, resp3, err3 := svc.UpdateIdpAdapter(&input3)
	equals(t, nil, err3)
	equals(t, 200, resp3.StatusCode)
	equals(t, *result3.Name, "woot")

	input4 := idpAdapters.DeleteIdpAdapterInput{
		Id: *input2.Body.Id,
	}
	result4, resp4, err4 := svc.DeleteIdpAdapter(&input4)
	equals(t, nil, err4)
	equals(t, 204, resp4.StatusCode)
	equals(t, (*models.ApiResult)(nil), result4)
}
