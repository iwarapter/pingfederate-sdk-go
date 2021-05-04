package pingOneForEnterprise

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
	ServiceName = "PingOneForEnterprise"
)

type PingOneForEnterpriseService struct {
	*client.PfClient
}

// New creates a new instance of the PingOneForEnterpriseService client.
func New(cfg *config.Config) *PingOneForEnterpriseService {

	return &PingOneForEnterpriseService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a PingOneForEnterprise operation
func (c *PingOneForEnterpriseService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetPingOneForEnterpriseSettings - Get the PingOne for Enterprise settings
//RequestType: GET
//Input:
func (s *PingOneForEnterpriseService) GetPingOneForEnterpriseSettings() (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	return s.GetPingOneForEnterpriseSettingsWithContext(context.Background())
}

//GetPingOneForEnterpriseSettingsWithContext - Get the PingOne for Enterprise settings
//RequestType: GET
//Input: ctx context.Context,
func (s *PingOneForEnterpriseService) GetPingOneForEnterpriseSettingsWithContext(ctx context.Context) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	path := "/pingOneForEnterprise"
	op := &request.Operation{
		Name:       "GetPingOneForEnterpriseSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingOneForEnterpriseSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingOneSettings - Update the PingOne for Enterprise settings.
//RequestType: PUT
//Input: input *UpdatePingOneSettingsInput
func (s *PingOneForEnterpriseService) UpdatePingOneSettings(input *UpdatePingOneSettingsInput) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	return s.UpdatePingOneSettingsWithContext(context.Background(), input)
}

//UpdatePingOneSettingsWithContext - Update the PingOne for Enterprise settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdatePingOneSettingsInput
func (s *PingOneForEnterpriseService) UpdatePingOneSettingsWithContext(ctx context.Context, input *UpdatePingOneSettingsInput) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	path := "/pingOneForEnterprise"
	op := &request.Operation{
		Name:       "UpdatePingOneSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.PingOneForEnterpriseSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//Disconnect - Disconnect from PingOne for Enterprise
//RequestType: POST
//Input:
func (s *PingOneForEnterpriseService) Disconnect() (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	return s.DisconnectWithContext(context.Background())
}

//DisconnectWithContext - Disconnect from PingOne for Enterprise
//RequestType: POST
//Input: ctx context.Context,
func (s *PingOneForEnterpriseService) DisconnectWithContext(ctx context.Context) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	path := "/pingOneForEnterprise/disconnect"
	op := &request.Operation{
		Name:       "Disconnect",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.PingOneForEnterpriseSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingOneForEnterpriseIdentityRepository - Update the PingOne Identity Repository
//RequestType: POST
//Input:
func (s *PingOneForEnterpriseService) UpdatePingOneForEnterpriseIdentityRepository() (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	return s.UpdatePingOneForEnterpriseIdentityRepositoryWithContext(context.Background())
}

//UpdatePingOneForEnterpriseIdentityRepositoryWithContext - Update the PingOne Identity Repository
//RequestType: POST
//Input: ctx context.Context,
func (s *PingOneForEnterpriseService) UpdatePingOneForEnterpriseIdentityRepositoryWithContext(ctx context.Context) (output *models.PingOneForEnterpriseSettings, resp *http.Response, err error) {
	path := "/pingOneForEnterprise/updateIdentityRepository"
	op := &request.Operation{
		Name:       "UpdatePingOneForEnterpriseIdentityRepository",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.PingOneForEnterpriseSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKeyPairs - Get the PingOne for Enterprise key pair settings
//RequestType: GET
//Input:
func (s *PingOneForEnterpriseService) GetKeyPairs() (output *models.P14EKeysView, resp *http.Response, err error) {
	return s.GetKeyPairsWithContext(context.Background())
}

//GetKeyPairsWithContext - Get the PingOne for Enterprise key pair settings
//RequestType: GET
//Input: ctx context.Context,
func (s *PingOneForEnterpriseService) GetKeyPairsWithContext(ctx context.Context) (output *models.P14EKeysView, resp *http.Response, err error) {
	path := "/pingOneForEnterprise/keyPairs"
	op := &request.Operation{
		Name:       "GetKeyPairs",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.P14EKeysView{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//RotateKeys - Rotate the authentication key
//RequestType: POST
//Input:
func (s *PingOneForEnterpriseService) RotateKeys() (output *models.P14EKeysView, resp *http.Response, err error) {
	return s.RotateKeysWithContext(context.Background())
}

//RotateKeysWithContext - Rotate the authentication key
//RequestType: POST
//Input: ctx context.Context,
func (s *PingOneForEnterpriseService) RotateKeysWithContext(ctx context.Context) (output *models.P14EKeysView, resp *http.Response, err error) {
	path := "/pingOneForEnterprise/keyPairs/rotate"
	op := &request.Operation{
		Name:       "RotateKeys",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.P14EKeysView{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdatePingOneSettingsInput struct {
	Body models.PingOneForEnterpriseSettings
}
