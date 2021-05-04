package serviceAuthentication

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ServiceAuthenticationAPI interface {
	GetServiceAuthentication() (output *models.ServiceAuthentication, resp *http.Response, err error)
	GetServiceAuthenticationWithContext(ctx context.Context) (output *models.ServiceAuthentication, resp *http.Response, err error)

	UpdateServiceAuthentication(input *UpdateServiceAuthenticationInput) (output *models.ServiceAuthentication, resp *http.Response, err error)
	UpdateServiceAuthenticationWithContext(ctx context.Context, input *UpdateServiceAuthenticationInput) (output *models.ServiceAuthentication, resp *http.Response, err error)
}
