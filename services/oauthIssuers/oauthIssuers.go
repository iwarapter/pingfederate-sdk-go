package oauthIssuers

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
	ServiceName = "OauthIssuers"
)

type OauthIssuersService struct {
	*client.PfClient
}

// New creates a new instance of the OauthIssuersService client.
func New(cfg *config.Config) *OauthIssuersService {

	return &OauthIssuersService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthIssuers operation
func (c *OauthIssuersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIssuers - Get the list of virtual issuers.
//RequestType: GET
//Input:
func (s *OauthIssuersService) GetIssuers() (output *models.Issuers, resp *http.Response, err error) {
	return s.GetIssuersWithContext(context.Background())
}

//GetIssuersWithContext - Get the list of virtual issuers.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthIssuersService) GetIssuersWithContext(ctx context.Context) (output *models.Issuers, resp *http.Response, err error) {
	path := "/oauth/issuers"
	op := &request.Operation{
		Name:       "GetIssuers",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Issuers{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddIssuer - Create a new virtual issuer.
//RequestType: POST
//Input: input *AddIssuerInput
func (s *OauthIssuersService) AddIssuer(input *AddIssuerInput) (output *models.Issuer, resp *http.Response, err error) {
	return s.AddIssuerWithContext(context.Background(), input)
}

//AddIssuerWithContext - Create a new virtual issuer.
//RequestType: POST
//Input: ctx context.Context, input *AddIssuerInput
func (s *OauthIssuersService) AddIssuerWithContext(ctx context.Context, input *AddIssuerInput) (output *models.Issuer, resp *http.Response, err error) {
	path := "/oauth/issuers"
	op := &request.Operation{
		Name:       "AddIssuer",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.Issuer{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIssuerById - Find a virtual issuer by ID.
//RequestType: GET
//Input: input *GetIssuerByIdInput
func (s *OauthIssuersService) GetIssuerById(input *GetIssuerByIdInput) (output *models.Issuer, resp *http.Response, err error) {
	return s.GetIssuerByIdWithContext(context.Background(), input)
}

//GetIssuerByIdWithContext - Find a virtual issuer by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetIssuerByIdInput
func (s *OauthIssuersService) GetIssuerByIdWithContext(ctx context.Context, input *GetIssuerByIdInput) (output *models.Issuer, resp *http.Response, err error) {
	path := "/oauth/issuers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIssuerById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Issuer{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIssuer - Update a virtual issuer.
//RequestType: PUT
//Input: input *UpdateIssuerInput
func (s *OauthIssuersService) UpdateIssuer(input *UpdateIssuerInput) (output *models.Issuer, resp *http.Response, err error) {
	return s.UpdateIssuerWithContext(context.Background(), input)
}

//UpdateIssuerWithContext - Update a virtual issuer.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateIssuerInput
func (s *OauthIssuersService) UpdateIssuerWithContext(ctx context.Context, input *UpdateIssuerInput) (output *models.Issuer, resp *http.Response, err error) {
	path := "/oauth/issuers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateIssuer",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.Issuer{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteIssuer - Delete a virtual issuer.
//RequestType: DELETE
//Input: input *DeleteIssuerInput
func (s *OauthIssuersService) DeleteIssuer(input *DeleteIssuerInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteIssuerWithContext(context.Background(), input)
}

//DeleteIssuerWithContext - Delete a virtual issuer.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteIssuerInput
func (s *OauthIssuersService) DeleteIssuerWithContext(ctx context.Context, input *DeleteIssuerInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/issuers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteIssuer",
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

type AddIssuerInput struct {
	Body models.Issuer
}

type DeleteIssuerInput struct {
	Id string
}

type GetIssuerByIdInput struct {
	Id string
}

type UpdateIssuerInput struct {
	Body models.Issuer
	Id   string
}
