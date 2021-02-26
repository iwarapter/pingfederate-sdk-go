package kerberosRealms

import (
	"context"
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
	ServiceName = "KerberosRealms"
)

type KerberosRealmsService struct {
	*client.PfClient
}

// New creates a new instance of the KerberosRealmsService client.
func New(cfg *config.Config) *KerberosRealmsService {

	return &KerberosRealmsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a KerberosRealms operation
func (c *KerberosRealmsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetKerberosRealmSettings - Gets the Kerberos Realms Settings.
//RequestType: GET
//Input:
func (s *KerberosRealmsService) GetKerberosRealmSettings() (output *models.KerberosRealmsSettings, resp *http.Response, err error) {
	return s.GetKerberosRealmSettingsWithContext(context.Background())
}

//GetKerberosRealmSettingsWithContext - Gets the Kerberos Realms Settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *KerberosRealmsService) GetKerberosRealmSettingsWithContext(ctx context.Context) (output *models.KerberosRealmsSettings, resp *http.Response, err error) {
	path := "/kerberos/realms/settings"
	op := &request.Operation{
		Name:       "GetKerberosRealmSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KerberosRealmsSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Set/Update the Kerberos Realms Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *KerberosRealmsService) UpdateSettings(input *UpdateSettingsInput) (output *models.KerberosRealmsSettings, resp *http.Response, err error) {
	return s.UpdateSettingsWithContext(context.Background(), input)
}

//UpdateSettingsWithContext - Set/Update the Kerberos Realms Settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSettingsInput
func (s *KerberosRealmsService) UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.KerberosRealmsSettings, resp *http.Response, err error) {
	path := "/kerberos/realms/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.KerberosRealmsSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKerberosRealms - Gets the Kerberos Realms.
//RequestType: GET
//Input:
func (s *KerberosRealmsService) GetKerberosRealms() (output *models.KerberosRealms, resp *http.Response, err error) {
	return s.GetKerberosRealmsWithContext(context.Background())
}

//GetKerberosRealmsWithContext - Gets the Kerberos Realms.
//RequestType: GET
//Input: ctx context.Context,
func (s *KerberosRealmsService) GetKerberosRealmsWithContext(ctx context.Context) (output *models.KerberosRealms, resp *http.Response, err error) {
	path := "/kerberos/realms"
	op := &request.Operation{
		Name:       "GetKerberosRealms",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KerberosRealms{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateKerberosRealm - Create a new Kerberos Realm.
//RequestType: POST
//Input: input *CreateKerberosRealmInput
func (s *KerberosRealmsService) CreateKerberosRealm(input *CreateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error) {
	return s.CreateKerberosRealmWithContext(context.Background(), input)
}

//CreateKerberosRealmWithContext - Create a new Kerberos Realm.
//RequestType: POST
//Input: ctx context.Context, input *CreateKerberosRealmInput
func (s *KerberosRealmsService) CreateKerberosRealmWithContext(ctx context.Context, input *CreateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms"
	op := &request.Operation{
		Name:       "CreateKerberosRealm",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.KerberosRealm{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKerberosRealm - Find a Kerberos Realm by ID.
//RequestType: GET
//Input: input *GetKerberosRealmInput
func (s *KerberosRealmsService) GetKerberosRealm(input *GetKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error) {
	return s.GetKerberosRealmWithContext(context.Background(), input)
}

//GetKerberosRealmWithContext - Find a Kerberos Realm by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetKerberosRealmInput
func (s *KerberosRealmsService) GetKerberosRealmWithContext(ctx context.Context, input *GetKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetKerberosRealm",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KerberosRealm{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateKerberosRealm - Update a Kerberos Realm by ID.
//RequestType: PUT
//Input: input *UpdateKerberosRealmInput
func (s *KerberosRealmsService) UpdateKerberosRealm(input *UpdateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error) {
	return s.UpdateKerberosRealmWithContext(context.Background(), input)
}

//UpdateKerberosRealmWithContext - Update a Kerberos Realm by ID.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateKerberosRealmInput
func (s *KerberosRealmsService) UpdateKerberosRealmWithContext(ctx context.Context, input *UpdateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateKerberosRealm",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.KerberosRealm{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteKerberosRealm - Delete a Kerberos Realm.
//RequestType: DELETE
//Input: input *DeleteKerberosRealmInput
func (s *KerberosRealmsService) DeleteKerberosRealm(input *DeleteKerberosRealmInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteKerberosRealmWithContext(context.Background(), input)
}

//DeleteKerberosRealmWithContext - Delete a Kerberos Realm.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteKerberosRealmInput
func (s *KerberosRealmsService) DeleteKerberosRealmWithContext(ctx context.Context, input *DeleteKerberosRealmInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteKerberosRealm",
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

type CreateKerberosRealmInput struct {
	Body models.KerberosRealm
}

type DeleteKerberosRealmInput struct {
	Id string
}

type GetKerberosRealmInput struct {
	Id string
}

type UpdateKerberosRealmInput struct {
	Body models.KerberosRealm
	Id   string
}

type UpdateSettingsInput struct {
	Body models.KerberosRealmsSettings
}
