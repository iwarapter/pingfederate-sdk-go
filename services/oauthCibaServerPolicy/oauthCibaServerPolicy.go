package oauthCibaServerPolicy

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
	ServiceName = "OauthCibaServerPolicy"
)

type OauthCibaServerPolicyService struct {
	*client.PfClient
}

// New creates a new instance of the OauthCibaServerPolicyService client.
func New(cfg *config.Config) *OauthCibaServerPolicyService {

	return &OauthCibaServerPolicyService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthCibaServerPolicy operation
func (c *OauthCibaServerPolicyService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get general ciba server request policy settings.
//RequestType: GET
//Input:
func (s *OauthCibaServerPolicyService) GetSettings() (output *models.CibaServerPolicySettings, resp *http.Response, err error) {
	return s.GetSettingsWithContext(context.Background())
}

//GetSettingsWithContext - Get general ciba server request policy settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthCibaServerPolicyService) GetSettingsWithContext(ctx context.Context) (output *models.CibaServerPolicySettings, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CibaServerPolicySettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Update general ciba server request policy settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthCibaServerPolicyService) UpdateSettings(input *UpdateSettingsInput) (output *models.CibaServerPolicySettings, resp *http.Response, err error) {
	return s.UpdateSettingsWithContext(context.Background(), input)
}

//UpdateSettingsWithContext - Update general ciba server request policy settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSettingsInput
func (s *OauthCibaServerPolicyService) UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.CibaServerPolicySettings, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.CibaServerPolicySettings{}
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

//GetPolicies - Get list of request policies.
//RequestType: GET
//Input:
func (s *OauthCibaServerPolicyService) GetPolicies() (output *models.RequestPolicies, resp *http.Response, err error) {
	return s.GetPoliciesWithContext(context.Background())
}

//GetPoliciesWithContext - Get list of request policies.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthCibaServerPolicyService) GetPoliciesWithContext(ctx context.Context) (output *models.RequestPolicies, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies"
	op := &request.Operation{
		Name:       "GetPolicies",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RequestPolicies{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreatePolicy - Create a new request policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthCibaServerPolicyService) CreatePolicy(input *CreatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error) {
	return s.CreatePolicyWithContext(context.Background(), input)
}

//CreatePolicyWithContext - Create a new request policy.
//RequestType: POST
//Input: ctx context.Context, input *CreatePolicyInput
func (s *OauthCibaServerPolicyService) CreatePolicyWithContext(ctx context.Context, input *CreatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies"
	op := &request.Operation{
		Name:       "CreatePolicy",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.RequestPolicy{}
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

//GetPolicy - Find request policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthCibaServerPolicyService) GetPolicy(input *GetPolicyInput) (output *models.RequestPolicy, resp *http.Response, err error) {
	return s.GetPolicyWithContext(context.Background(), input)
}

//GetPolicyWithContext - Find request policy by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetPolicyInput
func (s *OauthCibaServerPolicyService) GetPolicyWithContext(ctx context.Context, input *GetPolicyInput) (output *models.RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RequestPolicy{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePolicy - Update a request policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthCibaServerPolicyService) UpdatePolicy(input *UpdatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error) {
	return s.UpdatePolicyWithContext(context.Background(), input)
}

//UpdatePolicyWithContext - Update a request policy.
//RequestType: PUT
//Input: ctx context.Context, input *UpdatePolicyInput
func (s *OauthCibaServerPolicyService) UpdatePolicyWithContext(ctx context.Context, input *UpdatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdatePolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.RequestPolicy{}
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

//DeletePolicy - Delete a request policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthCibaServerPolicyService) DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeletePolicyWithContext(context.Background(), input)
}

//DeletePolicyWithContext - Delete a request policy.
//RequestType: DELETE
//Input: ctx context.Context, input *DeletePolicyInput
func (s *OauthCibaServerPolicyService) DeletePolicyWithContext(ctx context.Context, input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeletePolicy",
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

type CreatePolicyInput struct {
	Body models.RequestPolicy

	BypassExternalValidation *bool
}

type DeletePolicyInput struct {
	Id string
}

type GetPolicyInput struct {
	Id string
}

type UpdatePolicyInput struct {
	Body models.RequestPolicy
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.CibaServerPolicySettings

	BypassExternalValidation *bool
}
