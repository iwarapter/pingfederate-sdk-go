package protocolMetadata

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "ProtocolMetadata"
)

type ProtocolMetadataService struct {
	*client.PfClient
}

// New creates a new instance of the ProtocolMetadataService client.
func New(cfg *config.Config) *ProtocolMetadataService {

	return &ProtocolMetadataService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ProtocolMetadata operation
func (c *ProtocolMetadataService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetLifetimeSettings - Get metadata cache duration and reload delay for automated reloading.
//RequestType: GET
//Input:
func (s *ProtocolMetadataService) GetLifetimeSettings() (output *models.MetadataLifetimeSettings, resp *http.Response, err error) {
	return s.GetLifetimeSettingsWithContext(context.Background())
}

//GetLifetimeSettingsWithContext - Get metadata cache duration and reload delay for automated reloading.
//RequestType: GET
//Input: ctx context.Context,
func (s *ProtocolMetadataService) GetLifetimeSettingsWithContext(ctx context.Context) (output *models.MetadataLifetimeSettings, resp *http.Response, err error) {
	path := "/protocolMetadata/lifetimeSettings"
	op := &request.Operation{
		Name:       "GetLifetimeSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.MetadataLifetimeSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateLifetimeSettings - Update metadata cache duration and reload delay for automated reloading.
//RequestType: PUT
//Input: input *UpdateLifetimeSettingsInput
func (s *ProtocolMetadataService) UpdateLifetimeSettings(input *UpdateLifetimeSettingsInput) (output *models.MetadataLifetimeSettings, resp *http.Response, err error) {
	return s.UpdateLifetimeSettingsWithContext(context.Background(), input)
}

//UpdateLifetimeSettingsWithContext - Update metadata cache duration and reload delay for automated reloading.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateLifetimeSettingsInput
func (s *ProtocolMetadataService) UpdateLifetimeSettingsWithContext(ctx context.Context, input *UpdateLifetimeSettingsInput) (output *models.MetadataLifetimeSettings, resp *http.Response, err error) {
	path := "/protocolMetadata/lifetimeSettings"
	op := &request.Operation{
		Name:       "UpdateLifetimeSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.MetadataLifetimeSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSigningSettings - Get the certificate ID and algorithm used for metadata signing.
//RequestType: GET
//Input:
func (s *ProtocolMetadataService) GetSigningSettings() (output *models.MetadataSigningSettings, resp *http.Response, err error) {
	return s.GetSigningSettingsWithContext(context.Background())
}

//GetSigningSettingsWithContext - Get the certificate ID and algorithm used for metadata signing.
//RequestType: GET
//Input: ctx context.Context,
func (s *ProtocolMetadataService) GetSigningSettingsWithContext(ctx context.Context) (output *models.MetadataSigningSettings, resp *http.Response, err error) {
	path := "/protocolMetadata/signingSettings"
	op := &request.Operation{
		Name:       "GetSigningSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.MetadataSigningSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSigningSettings - Update the certificate and algorithm for metadata signing.
//RequestType: PUT
//Input: input *UpdateSigningSettingsInput
func (s *ProtocolMetadataService) UpdateSigningSettings(input *UpdateSigningSettingsInput) (output *models.MetadataSigningSettings, resp *http.Response, err error) {
	return s.UpdateSigningSettingsWithContext(context.Background(), input)
}

//UpdateSigningSettingsWithContext - Update the certificate and algorithm for metadata signing.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSigningSettingsInput
func (s *ProtocolMetadataService) UpdateSigningSettingsWithContext(ctx context.Context, input *UpdateSigningSettingsInput) (output *models.MetadataSigningSettings, resp *http.Response, err error) {
	path := "/protocolMetadata/signingSettings"
	op := &request.Operation{
		Name:       "UpdateSigningSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.MetadataSigningSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateLifetimeSettingsInput struct {
	Body models.MetadataLifetimeSettings
}

type UpdateSigningSettingsInput struct {
	Body models.MetadataSigningSettings
}
