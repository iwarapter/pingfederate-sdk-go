package session

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
	ServiceName = "Session"
)

type SessionService struct {
	*client.PfClient
}

// New creates a new instance of the SessionService client.
func New(cfg *config.Config) *SessionService {

	return &SessionService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Session operation
func (c *SessionService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSessionSettings - Get general session management settings.
//RequestType: GET
//Input:
func (s *SessionService) GetSessionSettings() (output *models.SessionSettings, resp *http.Response, err error) {
	return s.GetSessionSettingsWithContext(context.Background())
}

//GetSessionSettingsWithContext - Get general session management settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *SessionService) GetSessionSettingsWithContext(ctx context.Context) (output *models.SessionSettings, resp *http.Response, err error) {
	path := "/session/settings"
	op := &request.Operation{
		Name:       "GetSessionSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SessionSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSessionSettings - Update general session management settings.
//RequestType: PUT
//Input: input *UpdateSessionSettingsInput
func (s *SessionService) UpdateSessionSettings(input *UpdateSessionSettingsInput) (output *models.SessionSettings, resp *http.Response, err error) {
	return s.UpdateSessionSettingsWithContext(context.Background(), input)
}

//UpdateSessionSettingsWithContext - Update general session management settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSessionSettingsInput
func (s *SessionService) UpdateSessionSettingsWithContext(ctx context.Context, input *UpdateSessionSettingsInput) (output *models.SessionSettings, resp *http.Response, err error) {
	path := "/session/settings"
	op := &request.Operation{
		Name:       "UpdateSessionSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SessionSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetGlobalPolicy - Get the global authentication session policy.
//RequestType: GET
//Input:
func (s *SessionService) GetGlobalPolicy() (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	return s.GetGlobalPolicyWithContext(context.Background())
}

//GetGlobalPolicyWithContext - Get the global authentication session policy.
//RequestType: GET
//Input: ctx context.Context,
func (s *SessionService) GetGlobalPolicyWithContext(ctx context.Context) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/global"
	op := &request.Operation{
		Name:       "GetGlobalPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.GlobalAuthenticationSessionPolicy{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateGlobalPolicy - Update the global authentication session policy.
//RequestType: PUT
//Input: input *UpdateGlobalPolicyInput
func (s *SessionService) UpdateGlobalPolicy(input *UpdateGlobalPolicyInput) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	return s.UpdateGlobalPolicyWithContext(context.Background(), input)
}

//UpdateGlobalPolicyWithContext - Update the global authentication session policy.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateGlobalPolicyInput
func (s *SessionService) UpdateGlobalPolicyWithContext(ctx context.Context, input *UpdateGlobalPolicyInput) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/global"
	op := &request.Operation{
		Name:       "UpdateGlobalPolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.GlobalAuthenticationSessionPolicy{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetApplicationPolicy - Get the application session policy.
//RequestType: GET
//Input:
func (s *SessionService) GetApplicationPolicy() (output *models.ApplicationSessionPolicy, resp *http.Response, err error) {
	return s.GetApplicationPolicyWithContext(context.Background())
}

//GetApplicationPolicyWithContext - Get the application session policy.
//RequestType: GET
//Input: ctx context.Context,
func (s *SessionService) GetApplicationPolicyWithContext(ctx context.Context) (output *models.ApplicationSessionPolicy, resp *http.Response, err error) {
	path := "/session/applicationSessionPolicy"
	op := &request.Operation{
		Name:       "GetApplicationPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ApplicationSessionPolicy{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateApplicationPolicy - Update the application session policy.
//RequestType: PUT
//Input: input *UpdateApplicationPolicyInput
func (s *SessionService) UpdateApplicationPolicy(input *UpdateApplicationPolicyInput) (output *models.ApplicationSessionPolicy, resp *http.Response, err error) {
	return s.UpdateApplicationPolicyWithContext(context.Background(), input)
}

//UpdateApplicationPolicyWithContext - Update the application session policy.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateApplicationPolicyInput
func (s *SessionService) UpdateApplicationPolicyWithContext(ctx context.Context, input *UpdateApplicationPolicyInput) (output *models.ApplicationSessionPolicy, resp *http.Response, err error) {
	path := "/session/applicationSessionPolicy"
	op := &request.Operation{
		Name:       "UpdateApplicationPolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ApplicationSessionPolicy{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSourcePolicies - Get list of session policies.
//RequestType: GET
//Input:
func (s *SessionService) GetSourcePolicies() (output *models.AuthenticationSessionPolicies, resp *http.Response, err error) {
	return s.GetSourcePoliciesWithContext(context.Background())
}

//GetSourcePoliciesWithContext - Get list of session policies.
//RequestType: GET
//Input: ctx context.Context,
func (s *SessionService) GetSourcePoliciesWithContext(ctx context.Context) (output *models.AuthenticationSessionPolicies, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies"
	op := &request.Operation{
		Name:       "GetSourcePolicies",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSessionPolicies{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateSourcePolicy - Create a new session policy.
//RequestType: POST
//Input: input *CreateSourcePolicyInput
func (s *SessionService) CreateSourcePolicy(input *CreateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	return s.CreateSourcePolicyWithContext(context.Background(), input)
}

//CreateSourcePolicyWithContext - Create a new session policy.
//RequestType: POST
//Input: ctx context.Context, input *CreateSourcePolicyInput
func (s *SessionService) CreateSourcePolicyWithContext(ctx context.Context, input *CreateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies"
	op := &request.Operation{
		Name:       "CreateSourcePolicy",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSessionPolicy{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSourcePolicy - Find session policy by ID.
//RequestType: GET
//Input: input *GetSourcePolicyInput
func (s *SessionService) GetSourcePolicy(input *GetSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	return s.GetSourcePolicyWithContext(context.Background(), input)
}

//GetSourcePolicyWithContext - Find session policy by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetSourcePolicyInput
func (s *SessionService) GetSourcePolicyWithContext(ctx context.Context, input *GetSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetSourcePolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSessionPolicy{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSourcePolicy - Update a session policy.
//RequestType: PUT
//Input: input *UpdateSourcePolicyInput
func (s *SessionService) UpdateSourcePolicy(input *UpdateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	return s.UpdateSourcePolicyWithContext(context.Background(), input)
}

//UpdateSourcePolicyWithContext - Update a session policy.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSourcePolicyInput
func (s *SessionService) UpdateSourcePolicyWithContext(ctx context.Context, input *UpdateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateSourcePolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSessionPolicy{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteSourcePolicy - Delete a session policy.
//RequestType: DELETE
//Input: input *DeleteSourcePolicyInput
func (s *SessionService) DeleteSourcePolicy(input *DeleteSourcePolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteSourcePolicyWithContext(context.Background(), input)
}

//DeleteSourcePolicyWithContext - Delete a session policy.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteSourcePolicyInput
func (s *SessionService) DeleteSourcePolicyWithContext(ctx context.Context, input *DeleteSourcePolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteSourcePolicy",
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

type CreateSourcePolicyInput struct {
	Body models.AuthenticationSessionPolicy
}

type DeleteSourcePolicyInput struct {
	Id string
}

type GetSourcePolicyInput struct {
	Id string
}

type UpdateApplicationPolicyInput struct {
	Body models.ApplicationSessionPolicy
}

type UpdateGlobalPolicyInput struct {
	Body models.GlobalAuthenticationSessionPolicy
}

type UpdateSessionSettingsInput struct {
	Body models.SessionSettings
}

type UpdateSourcePolicyInput struct {
	Body models.AuthenticationSessionPolicy
	Id   string
}
