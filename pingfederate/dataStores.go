package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type DataStoresService service

//GetCustomDataStoreDescriptors - Get the list of available custom data store descriptors.
//RequestType: GET
//Input:
func (s *DataStoresService) GetCustomDataStoreDescriptors() (result *CustomDataStoreDescriptors, resp *http.Response, err error) {
	path := "/dataStores/descriptors"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCustomDataStoreDescriptor - Get the description of a custom data store plugin by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreDescriptorInput
func (s *DataStoresService) GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (result *CustomDataStoreDescriptor, resp *http.Response, err error) {
	path := "/dataStores/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetDataStores - Get list of data stores.
//RequestType: GET
//Input:
func (s *DataStoresService) GetDataStores() (result *DataStores, resp *http.Response, err error) {
	path := "/dataStores"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreateDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateDataStoreInput
func (s *DataStoresService) CreateDataStore(input *CreateDataStoreInput) (result *DataStore, resp *http.Response, err error) {
	path := "/dataStores"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	var req *http.Request
	if input.Body.LdapDataStore != nil {
		req, err = s.client.newRequest("POST", rel, input.Body.LdapDataStore)
	}
	if input.Body.JdbcDataStore != nil {
		req, err = s.client.newRequest("POST", rel, input.Body.JdbcDataStore)
	}
	if input.Body.CustomDataStore != nil {
		req, err = s.client.newRequest("POST", rel, input.Body.CustomDataStore)
	}
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetDataStoreInput
func (s *DataStoresService) GetDataStore(input *GetDataStoreInput) (result *DataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateDataStoreInput
func (s *DataStoresService) UpdateDataStore(input *UpdateDataStoreInput) (result *DataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	var req *http.Request
	if input.Body.LdapDataStore != nil {
		req, err = s.client.newRequest("PUT", rel, input.Body.LdapDataStore)
	}
	if input.Body.JdbcDataStore != nil {
		req, err = s.client.newRequest("PUT", rel, input.Body.JdbcDataStore)
	}
	if input.Body.CustomDataStore != nil {
		req, err = s.client.newRequest("PUT", rel, input.Body.CustomDataStore)
	}
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteDataStore - Delete a data store.
//RequestType: DELETE
//Input: input *DeleteDataStoreInput
func (s *DataStoresService) DeleteDataStore(input *DeleteDataStoreInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/dataStores/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetActions - List the actions for a data store instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *DataStoresService) GetActions(input *GetActionsInput) (result *Actions, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetAction - Find a data store instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *DataStoresService) GetAction(input *GetActionInput) (result *Action, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//InvokeAction - Invokes an action for a data source instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *DataStoresService) InvokeAction(input *InvokeActionInput) (result *ActionResult, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
