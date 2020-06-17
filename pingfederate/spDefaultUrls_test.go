package pingfederate_test

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/spDefaultUrls"
)

func TestSpDefaultUrls(t *testing.T) {
	svc := spDefaultUrls.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

	input1 := spDefaultUrls.UpdateDefaultUrlsInput{
		Body: models.SpDefaultUrls{
			ConfirmSlo: pingfederate.Bool(true),
		},
	}
	result1, resp1, err1 := svc.UpdateDefaultUrls(&input1)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, true, *result1.ConfirmSlo)

	result2, resp2, err2 := svc.GetDefaultUrls()
	equals(t, nil, err2)
	equals(t, 200, resp2.StatusCode)
	equals(t, true, *result2.ConfirmSlo)
}
