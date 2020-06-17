package spIdpConnections

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpIdpConnectionsAPI interface {
	GetConnections(input *GetConnectionsInput) (output *models.IdpConnections, resp *http.Response, err error)
	CreateConnection(input *CreateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)
	GetConnection(input *GetConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)
	UpdateConnection(input *UpdateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)
	DeleteConnection(input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error)
	GetSigningSettings(input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error)
	UpdateSigningSettings(input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error)
	AddConnectionCert(input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error)
	GetConnectionCerts(input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error)
	UpdateConnectionCerts(input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error)
	GetDecryptionKeys(input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error)
	UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error)
}
