package certificatesGroups

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
	ServiceName = "CertificatesGroups"
)

type CertificatesGroupsService struct {
	*client.PfClient
}

// New creates a new instance of the CertificatesGroupsService client.
func New(cfg *config.Config) *CertificatesGroupsService {

	return &CertificatesGroupsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a CertificatesGroups operation
func (c *CertificatesGroupsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetCertificatesForGroup - Get list of all certificates for a group.
//RequestType: GET
//Input: input *GetCertificatesForGroupInput
func (s *CertificatesGroupsService) GetCertificatesForGroup(input *GetCertificatesForGroupInput) (output *models.CertViews, resp *http.Response, err error) {
	return s.GetCertificatesForGroupWithContext(context.Background(), input)
}

//GetCertificatesForGroupWithContext - Get list of all certificates for a group.
//RequestType: GET
//Input: ctx context.Context, input *GetCertificatesForGroupInput
func (s *CertificatesGroupsService) GetCertificatesForGroupWithContext(ctx context.Context, input *GetCertificatesForGroupInput) (output *models.CertViews, resp *http.Response, err error) {
	path := "/certificates/groups/{groupName}"
	path = strings.Replace(path, "{groupName}", input.GroupName, -1)

	op := &request.Operation{
		Name:       "GetCertificatesForGroup",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CertViews{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportFeatureCert - Import a new certificate to a group.
//RequestType: POST
//Input: input *ImportFeatureCertInput
func (s *CertificatesGroupsService) ImportFeatureCert(input *ImportFeatureCertInput) (output *models.CertView, resp *http.Response, err error) {
	return s.ImportFeatureCertWithContext(context.Background(), input)
}

//ImportFeatureCertWithContext - Import a new certificate to a group.
//RequestType: POST
//Input: ctx context.Context, input *ImportFeatureCertInput
func (s *CertificatesGroupsService) ImportFeatureCertWithContext(ctx context.Context, input *ImportFeatureCertInput) (output *models.CertView, resp *http.Response, err error) {
	path := "/certificates/groups/{groupName}/import"
	path = strings.Replace(path, "{groupName}", input.GroupName, -1)

	op := &request.Operation{
		Name:       "ImportFeatureCert",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.CertView{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetCertificateFromGroup - Retrieve details of a certificate.
//RequestType: GET
//Input: input *GetCertificateFromGroupInput
func (s *CertificatesGroupsService) GetCertificateFromGroup(input *GetCertificateFromGroupInput) (output *models.CertView, resp *http.Response, err error) {
	return s.GetCertificateFromGroupWithContext(context.Background(), input)
}

//GetCertificateFromGroupWithContext - Retrieve details of a certificate.
//RequestType: GET
//Input: ctx context.Context, input *GetCertificateFromGroupInput
func (s *CertificatesGroupsService) GetCertificateFromGroupWithContext(ctx context.Context, input *GetCertificateFromGroupInput) (output *models.CertView, resp *http.Response, err error) {
	path := "/certificates/groups/{groupName}/{id}"
	path = strings.Replace(path, "{groupName}", input.GroupName, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetCertificateFromGroup",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CertView{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteCertificateFromGroup - Delete a certificate from a group.
//RequestType: DELETE
//Input: input *DeleteCertificateFromGroupInput
func (s *CertificatesGroupsService) DeleteCertificateFromGroup(input *DeleteCertificateFromGroupInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteCertificateFromGroupWithContext(context.Background(), input)
}

//DeleteCertificateFromGroupWithContext - Delete a certificate from a group.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteCertificateFromGroupInput
func (s *CertificatesGroupsService) DeleteCertificateFromGroupWithContext(ctx context.Context, input *DeleteCertificateFromGroupInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/certificates/groups/{groupName}/{id}"
	path = strings.Replace(path, "{groupName}", input.GroupName, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteCertificateFromGroup",
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

type DeleteCertificateFromGroupInput struct {
	GroupName string
	Id        string
}

type GetCertificateFromGroupInput struct {
	GroupName string
	Id        string
}

type GetCertificatesForGroupInput struct {
	GroupName string
}

type ImportFeatureCertInput struct {
	Body      models.X509File
	GroupName string
}
