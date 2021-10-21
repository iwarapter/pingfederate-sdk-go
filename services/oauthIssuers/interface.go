package oauthIssuers

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthIssuersAPI interface {
	GetIssuers() (output *models.Issuers, resp *http.Response, err error)
	GetIssuersWithContext(ctx context.Context) (output *models.Issuers, resp *http.Response, err error)

	AddIssuer(input *AddIssuerInput) (output *models.Issuer, resp *http.Response, err error)
	AddIssuerWithContext(ctx context.Context, input *AddIssuerInput) (output *models.Issuer, resp *http.Response, err error)

	GetIssuerById(input *GetIssuerByIdInput) (output *models.Issuer, resp *http.Response, err error)
	GetIssuerByIdWithContext(ctx context.Context, input *GetIssuerByIdInput) (output *models.Issuer, resp *http.Response, err error)

	UpdateIssuer(input *UpdateIssuerInput) (output *models.Issuer, resp *http.Response, err error)
	UpdateIssuerWithContext(ctx context.Context, input *UpdateIssuerInput) (output *models.Issuer, resp *http.Response, err error)

	DeleteIssuer(input *DeleteIssuerInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteIssuerWithContext(ctx context.Context, input *DeleteIssuerInput) (output *models.ApiResult, resp *http.Response, err error)
}
