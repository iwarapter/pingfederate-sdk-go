package pingOneForEnterprise

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type PingOneForEnterpriseAPI interface {
	GetPingOneForEnterpriseSettings() (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)
	GetPingOneForEnterpriseSettingsWithContext(ctx context.Context) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)

	UpdatePingOneSettings(input *UpdatePingOneSettingsInput) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)
	UpdatePingOneSettingsWithContext(ctx context.Context, input *UpdatePingOneSettingsInput) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)

	Disconnect() (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)
	DisconnectWithContext(ctx context.Context) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)

	UpdatePingOneForEnterpriseIdentityRepository() (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)
	UpdatePingOneForEnterpriseIdentityRepositoryWithContext(ctx context.Context) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error)

	GetKeyPairs() (output *models.P14EKeysView, resp *http.Response, err error)
	GetKeyPairsWithContext(ctx context.Context) (output *models.P14EKeysView, resp *http.Response, err error)

	RotateKeys() (output *models.P14EKeysView, resp *http.Response, err error)
	RotateKeysWithContext(ctx context.Context) (output *models.P14EKeysView, resp *http.Response, err error)
}
