package pingfederate_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/spDefaultUrls"
)

func TestSpDefaultUrls(t *testing.T) {
	svc := spDefaultUrls.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))

	input1 := spDefaultUrls.UpdateDefaultUrlsInput{
		Body: models.SpDefaultUrls{
			ConfirmSlo: pingfederate.Bool(true),
		},
	}
	result1, resp1, err1 := svc.UpdateDefaultUrls(&input1)
	require.Nil(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.True(t, *result1.ConfirmSlo)

	result2, resp2, err2 := svc.GetDefaultUrls()
	require.Nil(t, err2)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)
	assert.True(t, *result2.ConfirmSlo)
}
