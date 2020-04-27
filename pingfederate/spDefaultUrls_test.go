package pingfederate

import "testing"

func TestSpDefaultUrls(t *testing.T) {
	svc := basicClient()

	input1 := UpdateDefaultUrlsInput{
		Body: SpDefaultUrls{
			ConfirmSlo: Bool(true),
		},
	}
	result1, resp1, err1 := svc.SpDefaultUrls.UpdateDefaultUrls(&input1)
	equals(t, nil, err1)
	equals(t, 200, resp1.StatusCode)
	equals(t, true, *result1.ConfirmSlo)

	result2, resp2, err2 := svc.SpDefaultUrls.GetDefaultUrls()
	equals(t, nil, err2)
	equals(t, 200, resp2.StatusCode)
	equals(t, true, *result2.ConfirmSlo)
}
