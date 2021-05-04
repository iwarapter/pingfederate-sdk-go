package certificatesGroups

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesGroupsAPI interface {
	GetCertificatesForGroup(input *GetCertificatesForGroupInput) (output *models.CertViews, resp *http.Response, err error)
	GetCertificatesForGroupWithContext(ctx context.Context, input *GetCertificatesForGroupInput) (output *models.CertViews, resp *http.Response, err error)

	ImportFeatureCert(input *ImportFeatureCertInput) (output *models.CertView, resp *http.Response, err error)
	ImportFeatureCertWithContext(ctx context.Context, input *ImportFeatureCertInput) (output *models.CertView, resp *http.Response, err error)

	GetCertificateFromGroup(input *GetCertificateFromGroupInput) (output *models.CertView, resp *http.Response, err error)
	GetCertificateFromGroupWithContext(ctx context.Context, input *GetCertificateFromGroupInput) (output *models.CertView, resp *http.Response, err error)

	DeleteCertificateFromGroup(input *DeleteCertificateFromGroupInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteCertificateFromGroupWithContext(ctx context.Context, input *DeleteCertificateFromGroupInput) (output *models.ApiResult, resp *http.Response, err error)
}
