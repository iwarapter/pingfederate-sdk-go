package dataStores

import (
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
	ServiceName = "DataStores"
)

type DataStoresService struct {
	*client.PfClient
}

// New creates a new instance of the DataStoresService client.
func New(cfg *config.Config) *DataStoresService {

	return &DataStoresService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a DataStores operation
func (c *DataStoresService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetCustomDataStoreDescriptors - Get the list of available custom data store descriptors.
//RequestType: GET
//Input:
func (s *DataStoresService) GetCustomDataStoreDescriptors() (output *models.CustomDataStoreDescriptors, resp *http.Response, err error) {
	path := "/dataStores/descriptors"
	op := &request.Operation{
		Name:       "GetCustomDataStoreDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CustomDataStoreDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetCustomDataStoreDescriptor - Get the description of a custom data store plugin by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreDescriptorInput
func (s *DataStoresService) GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (output *models.CustomDataStoreDescriptor, resp *http.Response, err error) {
	path := "/dataStores/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetCustomDataStoreDescriptor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CustomDataStoreDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDataStores - Get list of data stores.
//RequestType: GET
//Input:
func (s *DataStoresService) GetDataStores() (output *models.DataStores, resp *http.Response, err error) {
	path := "/dataStores"
	op := &request.Operation{
		Name:       "GetDataStores",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DataStores{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateJdbcDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateJdbcDataStoreInput
func (s *DataStoresService) CreateJdbcDataStore(input *CreateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores"
	op := &request.Operation{
		Name:       "CreateJdbcDataStore",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.JdbcDataStore{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateLdapDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateLdapDataStoreInput
func (s *DataStoresService) CreateLdapDataStore(input *CreateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores"
	op := &request.Operation{
		Name:       "CreateLdapDataStore",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.LdapDataStore{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateCustomDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateCustomDataStoreInput
func (s *DataStoresService) CreateCustomDataStore(input *CreateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores"
	op := &request.Operation{
		Name:       "CreateCustomDataStore",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.CustomDataStore{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetDataStoreInput
func (s *DataStoresService) GetDataStore(input *GetDataStoreInput) (output *models.DataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetDataStore",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DataStore{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteDataStore - Delete a data store.
//RequestType: DELETE
//Input: input *DeleteDataStoreInput
func (s *DataStoresService) DeleteDataStore(input *DeleteDataStoreInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteDataStore",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetJdbcDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetJdbcDataStoreInput
func (s *DataStoresService) GetJdbcDataStore(input *GetJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetJdbcDataStore",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.JdbcDataStore{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetLdapDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetLdapDataStoreInput
func (s *DataStoresService) GetLdapDataStore(input *GetLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetLdapDataStore",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LdapDataStore{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetCustomDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreInput
func (s *DataStoresService) GetCustomDataStore(input *GetCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetCustomDataStore",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CustomDataStore{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateJdbcDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateJdbcDataStoreInput
func (s *DataStoresService) UpdateJdbcDataStore(input *UpdateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateJdbcDataStore",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.JdbcDataStore{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateLdapDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateLdapDataStoreInput
func (s *DataStoresService) UpdateLdapDataStore(input *UpdateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateLdapDataStore",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.LdapDataStore{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateCustomDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateCustomDataStoreInput
func (s *DataStoresService) UpdateCustomDataStore(input *UpdateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateCustomDataStore",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.CustomDataStore{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetActions - List the actions for a data store instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *DataStoresService) GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetActions",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Actions{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAction - Find a data store instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *DataStoresService) GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	op := &request.Operation{
		Name:       "GetAction",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Action{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//InvokeAction - Invokes an action for a data source instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *DataStoresService) InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	op := &request.Operation{
		Name:       "InvokeAction",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ActionResult{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateCustomDataStoreInput struct {
	Body models.CustomDataStore

	BypassExternalValidation *bool
}

type CreateDataStoreInput struct {
	Body models.CustomDataStore

	BypassExternalValidation *bool
}

type CreateJdbcDataStoreInput struct {
	Body models.JdbcDataStore

	BypassExternalValidation *bool
}

type CreateLdapDataStoreInput struct {
	Body models.LdapDataStore

	BypassExternalValidation *bool
}

type DeleteDataStoreInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetCustomDataStoreInput struct {
	Id string
}

type GetCustomDataStoreDescriptorInput struct {
	Id string
}

type GetDataStoreInput struct {
	Id string
}

type GetJdbcDataStoreInput struct {
	Id string
}

type GetLdapDataStoreInput struct {
	Id string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateCustomDataStoreInput struct {
	Body models.CustomDataStore
	Id   string

	BypassExternalValidation *bool
}

type UpdateDataStoreInput struct {
	Body models.LdapDataStore
	Id   string

	BypassExternalValidation *bool
}

type UpdateJdbcDataStoreInput struct {
	Body models.JdbcDataStore
	Id   string

	BypassExternalValidation *bool
}

type UpdateLdapDataStoreInput struct {
	Body models.LdapDataStore
	Id   string

	BypassExternalValidation *bool
}
