package pingfederate

import (
	"testing"
)

func TestIdpAdapters(t *testing.T) {
	svc := basicClient()

	input := GetIdpAdaptersInput{}
	_, resp1, err1 := svc.IdpAdapters.GetIdpAdapters(&input)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	//equals(t, *(*result1.Items)[0].Name, "woot")

	input2 := CreateIdpAdapterInput{
		Body: IdpAdapter{
			AttributeContract: &IdpAdapterAttributeContract{
				CoreAttributes: &[]*IdpAdapterAttribute{
					&IdpAdapterAttribute{
						Masked:    Bool(false),
						Name:      String("subject"),
						Pseudonym: Bool(true),
					},
				},
				ExtendedAttributes: &[]*IdpAdapterAttribute{},
				MaskOgnlValues:     Bool(false),
			},
			AttributeMapping: &IdpAdapterContractMapping{
				AttributeContractFulfillment: map[string]*AttributeFulfillmentValue{
					"subject": &AttributeFulfillmentValue{
						Source: &SourceTypeIdKey{
							Type: String("ADAPTER"),
						},
						Value: String("subject"),
					},
				},
				IssuanceCriteria: &IssuanceCriteria{
					ConditionalCriteria: &[]*ConditionalIssuanceCriteriaEntry{},
				},
			},
			Configuration: &PluginConfiguration{
				Fields: &[]*ConfigField{
					&ConfigField{
						Name:  String("Password"),
						Value: String("Password01"),
					},
					&ConfigField{
						Name:  String("Confirm Password"),
						Value: String("Password01"),
					},
					&ConfigField{
						Name:  String("Authentication Service"),
						Value: String("http://foo.com"),
					},
				},
			},
			Id:   String("woot"),
			Name: String("woot"),
			PluginDescriptorRef: &ResourceLink{
				Id: String("com.pingidentity.adapters.opentoken.IdpAuthnAdapter"),
			},
		},
	}

	result2, resp2, err2 := svc.IdpAdapters.CreateIdpAdapter(&input2)
	equals(t, nil, err2)
	equals(t, 201, resp2.StatusCode)
	equals(t, *result2.Name, "woot")

	input3 := UpdateIdpAdapterInput{
		Body: input2.Body,
		Id:   *input2.Body.Id,
	}
	(*input3.Body.Configuration.Fields)[0].Value = String("Password02")
	(*input3.Body.Configuration.Fields)[1].Value = String("Password02")

	result3, resp3, err3 := svc.IdpAdapters.UpdateIdpAdapter(&input3)
	equals(t, nil, err3)
	equals(t, 200, resp3.StatusCode)
	equals(t, *result3.Name, "woot")

	input4 := DeleteIdpAdapterInput{
		Id: *input2.Body.Id,
	}
	result4, resp4, err4 := svc.IdpAdapters.DeleteIdpAdapter(&input4)
	equals(t, nil, err4)
	equals(t, 204, resp4.StatusCode)
	equals(t, (*ApiResult)(nil), result4)
}
