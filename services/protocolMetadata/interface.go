package protocolMetadata

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ProtocolMetadataAPI interface {
	GetLifetimeSettings() (output *models.MetadataLifetimeSettings, resp *http.Response, err error)
	GetLifetimeSettingsWithContext(ctx context.Context) (output *models.MetadataLifetimeSettings, resp *http.Response, err error)

	UpdateLifetimeSettings(input *UpdateLifetimeSettingsInput) (output *models.MetadataLifetimeSettings, resp *http.Response, err error)
	UpdateLifetimeSettingsWithContext(ctx context.Context, input *UpdateLifetimeSettingsInput) (output *models.MetadataLifetimeSettings, resp *http.Response, err error)

	GetSigningSettings() (output *models.MetadataSigningSettings, resp *http.Response, err error)
	GetSigningSettingsWithContext(ctx context.Context) (output *models.MetadataSigningSettings, resp *http.Response, err error)

	UpdateSigningSettings(input *UpdateSigningSettingsInput) (output *models.MetadataSigningSettings, resp *http.Response, err error)
	UpdateSigningSettingsWithContext(ctx context.Context, input *UpdateSigningSettingsInput) (output *models.MetadataSigningSettings, resp *http.Response, err error)
}
