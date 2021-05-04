package serviceAuthentication

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
	ServiceName = "ServiceAuthentication"
)

type ServiceAuthenticationService struct {
	*client.PfClient
}

// New creates a new instance of the ServiceAuthenticationService client.
func New(cfg *config.Config) *ServiceAuthenticationService {

	return &ServiceAuthenticationService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ServiceAuthentication operation
func (c *ServiceAuthenticationService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetServiceAuthentication - Get the service authentication settings.
//RequestType: GET
//Input:
func (s *ServiceAuthenticationService) GetServiceAuthentication() (output *models.ServiceAuthentication, resp *http.Response, err error) {
	return s.GetServiceAuthenticationWithContext(context.Background())
}

//GetServiceAuthenticationWithContext - Get the service authentication settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *ServiceAuthenticationService) GetServiceAuthenticationWithContext(ctx context.Context) (output *models.ServiceAuthentication, resp *http.Response, err error) {
	path := "/serviceAuthentication"
	op := &request.Operation{
		Name:       "GetServiceAuthentication",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ServiceAuthentication{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateServiceAuthentication - Update the service authentication settings.
//RequestType: PUT
//Input: input *UpdateServiceAuthenticationInput
func (s *ServiceAuthenticationService) UpdateServiceAuthentication(input *UpdateServiceAuthenticationInput) (output *models.ServiceAuthentication, resp *http.Response, err error) {
	return s.UpdateServiceAuthenticationWithContext(context.Background(), input)
}

//UpdateServiceAuthenticationWithContext - Update the service authentication settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateServiceAuthenticationInput
func (s *ServiceAuthenticationService) UpdateServiceAuthenticationWithContext(ctx context.Context, input *UpdateServiceAuthenticationInput) (output *models.ServiceAuthentication, resp *http.Response, err error) {
	path := "/serviceAuthentication"
	op := &request.Operation{
		Name:       "UpdateServiceAuthentication",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ServiceAuthentication{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateServiceAuthenticationInput struct {
	Body models.ServiceAuthentication
}
