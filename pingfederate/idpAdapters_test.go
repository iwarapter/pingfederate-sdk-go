package pingfederate_test

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"
)

func TestIdpAdapters(t *testing.T) {
	svc := idpAdapters.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

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
						Masked:    pingfederate.Bool(false),
						Name:      pingfederate.String("subject"),
						Pseudonym: pingfederate.Bool(true),
					},
				},
				ExtendedAttributes: &[]*models.IdpAdapterAttribute{},
				MaskOgnlValues:     pingfederate.Bool(false),
			},
			AttributeMapping: &models.IdpAdapterContractMapping{
				AttributeContractFulfillment: map[string]*models.AttributeFulfillmentValue{
					"subject": &models.AttributeFulfillmentValue{
						Source: &models.SourceTypeIdKey{
							Type: pingfederate.String("ADAPTER"),
						},
						Value: pingfederate.String("subject"),
					},
				},
				IssuanceCriteria: &models.IssuanceCriteria{
					ConditionalCriteria: &[]*models.ConditionalIssuanceCriteriaEntry{},
				},
			},
			Configuration: &models.PluginConfiguration{
				Fields: &[]*models.ConfigField{
					&models.ConfigField{
						Name:  pingfederate.String("Password"),
						Value: pingfederate.String("Password01"),
					},
					&models.ConfigField{
						Name:  pingfederate.String("Confirm Password"),
						Value: pingfederate.String("Password01"),
					},
					&models.ConfigField{
						Name:  pingfederate.String("Authentication Service"),
						Value: pingfederate.String("http://foo.com"),
					},
				},
			},
			Id:   pingfederate.String("woot"),
			Name: pingfederate.String("woot"),
			PluginDescriptorRef: &models.ResourceLink{
				Id: pingfederate.String("com.pingidentity.adapters.opentoken.IdpAuthnAdapter"),
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
	(*input3.Body.Configuration.Fields)[0].Value = pingfederate.String("Password02")
	(*input3.Body.Configuration.Fields)[1].Value = pingfederate.String("Password02")

	result3, resp3, err3 := svc.UpdateIdpAdapter(&input3)
	equals(t, nil, err3)
	equals(t, 200, resp3.StatusCode)
	equals(t, *result3.Name, "woot")

	input4 := idpAdapters.GetIdpAdaptersInput{
		Filter: *input2.Body.Name,
	}
	result4, resp4, err4 := svc.GetIdpAdapters(&input4)
	ok(t, err4)
	equals(t, 200, resp4.StatusCode)
	equals(t, *(*result4.Items)[0].Name, "woot")

	input5 := idpAdapters.DeleteIdpAdapterInput{
		Id: *input2.Body.Id,
	}
	result5, resp5, err5 := svc.DeleteIdpAdapter(&input5)
	equals(t, nil, err5)
	equals(t, 204, resp5.StatusCode)
	equals(t, (*models.ApiResult)(nil), result5)
}
