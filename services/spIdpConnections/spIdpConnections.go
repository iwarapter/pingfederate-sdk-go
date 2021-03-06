package spIdpConnections

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "SpIdpConnections"
)

type SpIdpConnectionsService struct {
	*client.PfClient
}

// New creates a new instance of the SpIdpConnectionsService client.
func New(cfg *config.Config) *SpIdpConnectionsService {

	return &SpIdpConnectionsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SpIdpConnections operation
func (c *SpIdpConnectionsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetConnections - Get list of IdP connections.
//RequestType: GET
//Input: input *GetConnectionsInput
func (s *SpIdpConnectionsService) GetConnections(input *GetConnectionsInput) (output *models.IdpConnections, resp *http.Response, err error) {
	return s.GetConnectionsWithContext(context.Background(), input)
}

//GetConnectionsWithContext - Get list of IdP connections.
//RequestType: GET
//Input: ctx context.Context, input *GetConnectionsInput
func (s *SpIdpConnectionsService) GetConnectionsWithContext(ctx context.Context, input *GetConnectionsInput) (output *models.IdpConnections, resp *http.Response, err error) {
	path := "/sp/idpConnections"
	op := &request.Operation{
		Name:       "GetConnections",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"entityId":      input.EntityId,
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.IdpConnections{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateConnection - Create a new IdP connection.
//RequestType: POST
//Input: input *CreateConnectionInput
func (s *SpIdpConnectionsService) CreateConnection(input *CreateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error) {
	return s.CreateConnectionWithContext(context.Background(), input)
}

//CreateConnectionWithContext - Create a new IdP connection.
//RequestType: POST
//Input: ctx context.Context, input *CreateConnectionInput
func (s *SpIdpConnectionsService) CreateConnectionWithContext(ctx context.Context, input *CreateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections"
	op := &request.Operation{
		Name:       "CreateConnection",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.IdpConnection{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetConnection - Find IdP connection by ID.
//RequestType: GET
//Input: input *GetConnectionInput
func (s *SpIdpConnectionsService) GetConnection(input *GetConnectionInput) (output *models.IdpConnection, resp *http.Response, err error) {
	return s.GetConnectionWithContext(context.Background(), input)
}

//GetConnectionWithContext - Find IdP connection by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetConnectionInput
func (s *SpIdpConnectionsService) GetConnectionWithContext(ctx context.Context, input *GetConnectionInput) (output *models.IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetConnection",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpConnection{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateConnection - Update an IdP connection.
//RequestType: PUT
//Input: input *UpdateConnectionInput
func (s *SpIdpConnectionsService) UpdateConnection(input *UpdateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error) {
	return s.UpdateConnectionWithContext(context.Background(), input)
}

//UpdateConnectionWithContext - Update an IdP connection.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateConnectionInput
func (s *SpIdpConnectionsService) UpdateConnectionWithContext(ctx context.Context, input *UpdateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateConnection",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.IdpConnection{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteConnection - Delete an IdP connection.
//RequestType: DELETE
//Input: input *DeleteConnectionInput
func (s *SpIdpConnectionsService) DeleteConnection(input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteConnectionWithContext(context.Background(), input)
}

//DeleteConnectionWithContext - Delete an IdP connection.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteConnectionInput
func (s *SpIdpConnectionsService) DeleteConnectionWithContext(ctx context.Context, input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteConnection",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSigningSettings - Get the IdP connection's signature settings.
//RequestType: GET
//Input: input *GetSigningSettingsInput
func (s *SpIdpConnectionsService) GetSigningSettings(input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	return s.GetSigningSettingsWithContext(context.Background(), input)
}

//GetSigningSettingsWithContext - Get the IdP connection's signature settings.
//RequestType: GET
//Input: ctx context.Context, input *GetSigningSettingsInput
func (s *SpIdpConnectionsService) GetSigningSettingsWithContext(ctx context.Context, input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetSigningSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SigningSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSigningSettings - Update the IdP connection's signature settings.
//RequestType: PUT
//Input: input *UpdateSigningSettingsInput
func (s *SpIdpConnectionsService) UpdateSigningSettings(input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	return s.UpdateSigningSettingsWithContext(context.Background(), input)
}

//UpdateSigningSettingsWithContext - Update the IdP connection's signature settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSigningSettingsInput
func (s *SpIdpConnectionsService) UpdateSigningSettingsWithContext(ctx context.Context, input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateSigningSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SigningSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddConnectionCert - Add a new IdP connection certificate.
//RequestType: POST
//Input: input *AddConnectionCertInput
func (s *SpIdpConnectionsService) AddConnectionCert(input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error) {
	return s.AddConnectionCertWithContext(context.Background(), input)
}

//AddConnectionCertWithContext - Add a new IdP connection certificate.
//RequestType: POST
//Input: ctx context.Context, input *AddConnectionCertInput
func (s *SpIdpConnectionsService) AddConnectionCertWithContext(ctx context.Context, input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "AddConnectionCert",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ConnectionCert{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetConnectionCerts - Get the IdP connection's certificates.
//RequestType: GET
//Input: input *GetConnectionCertsInput
func (s *SpIdpConnectionsService) GetConnectionCerts(input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	return s.GetConnectionCertsWithContext(context.Background(), input)
}

//GetConnectionCertsWithContext - Get the IdP connection's certificates.
//RequestType: GET
//Input: ctx context.Context, input *GetConnectionCertsInput
func (s *SpIdpConnectionsService) GetConnectionCertsWithContext(ctx context.Context, input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetConnectionCerts",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ConnectionCerts{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateConnectionCerts - Update the IdP connection's certificates.
//RequestType: PUT
//Input: input *UpdateConnectionCertsInput
func (s *SpIdpConnectionsService) UpdateConnectionCerts(input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	return s.UpdateConnectionCertsWithContext(context.Background(), input)
}

//UpdateConnectionCertsWithContext - Update the IdP connection's certificates.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateConnectionCertsInput
func (s *SpIdpConnectionsService) UpdateConnectionCertsWithContext(ctx context.Context, input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateConnectionCerts",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ConnectionCerts{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDecryptionKeys - Get the decryption keys of an IdP connection.
//RequestType: GET
//Input: input *GetDecryptionKeysInput
func (s *SpIdpConnectionsService) GetDecryptionKeys(input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	return s.GetDecryptionKeysWithContext(context.Background(), input)
}

//GetDecryptionKeysWithContext - Get the decryption keys of an IdP connection.
//RequestType: GET
//Input: ctx context.Context, input *GetDecryptionKeysInput
func (s *SpIdpConnectionsService) GetDecryptionKeysWithContext(ctx context.Context, input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetDecryptionKeys",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DecryptionKeys{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDecryptionKeys - Updating the IdP connection's decryption keys.
//RequestType: PUT
//Input: input *UpdateDecryptionKeysInput
func (s *SpIdpConnectionsService) UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	return s.UpdateDecryptionKeysWithContext(context.Background(), input)
}

//UpdateDecryptionKeysWithContext - Updating the IdP connection's decryption keys.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateDecryptionKeysInput
func (s *SpIdpConnectionsService) UpdateDecryptionKeysWithContext(ctx context.Context, input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateDecryptionKeys",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.DecryptionKeys{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddConnectionCertInput struct {
	Body models.ConnectionCert
	Id   string
}

type CreateConnectionInput struct {
	Body models.IdpConnection

	BypassExternalValidation *bool
}

type DeleteConnectionInput struct {
	Id string
}

type GetConnectionInput struct {
	Id string
}

type GetConnectionCertsInput struct {
	Id string
}

type GetConnectionsInput struct {
	EntityId      string
	Page          string
	NumberPerPage string
	Filter        string
}

type GetDecryptionKeysInput struct {
	Id string
}

type GetSigningSettingsInput struct {
	Id string
}

type UpdateConnectionInput struct {
	Body models.IdpConnection
	Id   string

	BypassExternalValidation *bool
}

type UpdateConnectionCertsInput struct {
	Body models.ConnectionCerts
	Id   string
}

type UpdateDecryptionKeysInput struct {
	Body models.DecryptionKeys
	Id   string
}

type UpdateSigningSettingsInput struct {
	Body models.SigningSettings
	Id   string
}
