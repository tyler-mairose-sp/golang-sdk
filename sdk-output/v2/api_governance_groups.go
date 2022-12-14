/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// GovernanceGroupsApiService GovernanceGroupsApi service
type GovernanceGroupsApiService service

type ApiBulkDeleteWorkGroupsRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
	bulkDeleteWorkGroupsRequest *BulkDeleteWorkGroupsRequest
}

// Work group ids to delete
func (r ApiBulkDeleteWorkGroupsRequest) BulkDeleteWorkGroupsRequest(bulkDeleteWorkGroupsRequest BulkDeleteWorkGroupsRequest) ApiBulkDeleteWorkGroupsRequest {
	r.bulkDeleteWorkGroupsRequest = &bulkDeleteWorkGroupsRequest
	return r
}

func (r ApiBulkDeleteWorkGroupsRequest) Execute() (*BulkDeleteWorkGroups200Response, *http.Response, error) {
	return r.ApiService.BulkDeleteWorkGroupsExecute(r)
}

/*
BulkDeleteWorkGroups Bulk delete work groups

This API allows you to bulk-delete work groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkDeleteWorkGroupsRequest
*/
func (a *GovernanceGroupsApiService) BulkDeleteWorkGroups(ctx context.Context) ApiBulkDeleteWorkGroupsRequest {
	return ApiBulkDeleteWorkGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BulkDeleteWorkGroups200Response
func (a *GovernanceGroupsApiService) BulkDeleteWorkGroupsExecute(r ApiBulkDeleteWorkGroupsRequest) (*BulkDeleteWorkGroups200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BulkDeleteWorkGroups200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.BulkDeleteWorkGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/bulk-delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bulkDeleteWorkGroupsRequest == nil {
		return localVarReturnValue, nil, reportError("bulkDeleteWorkGroupsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkDeleteWorkGroupsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
	createWorkgroupRequest *CreateWorkgroupRequest
}

// Work group to create.
func (r ApiCreateWorkgroupRequest) CreateWorkgroupRequest(createWorkgroupRequest CreateWorkgroupRequest) ApiCreateWorkgroupRequest {
	r.createWorkgroupRequest = &createWorkgroupRequest
	return r
}

func (r ApiCreateWorkgroupRequest) Execute() ([]Workgroup, *http.Response, error) {
	return r.ApiService.CreateWorkgroupExecute(r)
}

/*
CreateWorkgroup Create Work Group

This API allows you to create a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWorkgroupRequest
*/
func (a *GovernanceGroupsApiService) CreateWorkgroup(ctx context.Context) ApiCreateWorkgroupRequest {
	return ApiCreateWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Workgroup
func (a *GovernanceGroupsApiService) CreateWorkgroupExecute(r ApiCreateWorkgroupRequest) ([]Workgroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Workgroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.CreateWorkgroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createWorkgroupRequest == nil {
		return localVarReturnValue, nil, reportError("createWorkgroupRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createWorkgroupRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
}

func (r ApiDeleteWorkgroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWorkgroupExecute(r)
}

/*
DeleteWorkgroup Delete Work Group By Id

This API deletes a single workgroup based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteWorkgroupRequest
*/
func (a *GovernanceGroupsApiService) DeleteWorkgroup(ctx context.Context) ApiDeleteWorkgroupRequest {
	return ApiDeleteWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *GovernanceGroupsApiService) DeleteWorkgroupExecute(r ApiDeleteWorkgroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.DeleteWorkgroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workgroups/{workgroupId}"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
}

func (r ApiGetWorkgroupRequest) Execute() (*Workgroup, *http.Response, error) {
	return r.ApiService.GetWorkgroupExecute(r)
}

/*
GetWorkgroup Get Work Group By Id

This API returns the details for a single workgroup based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWorkgroupRequest
*/
func (a *GovernanceGroupsApiService) GetWorkgroup(ctx context.Context) ApiGetWorkgroupRequest {
	return ApiGetWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Workgroup
func (a *GovernanceGroupsApiService) GetWorkgroupExecute(r ApiGetWorkgroupRequest) (*Workgroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Workgroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.GetWorkgroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workgroups/{workgroupId}"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListWorkgroupConnectionsRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
}

func (r ApiListWorkgroupConnectionsRequest) Execute() ([]WorkgroupConnection, *http.Response, error) {
	return r.ApiService.ListWorkgroupConnectionsExecute(r)
}

/*
ListWorkgroupConnections List Work Group Connections

This API returns the connections of a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWorkgroupConnectionsRequest
*/
func (a *GovernanceGroupsApiService) ListWorkgroupConnections(ctx context.Context) ApiListWorkgroupConnectionsRequest {
	return ApiListWorkgroupConnectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []WorkgroupConnection
func (a *GovernanceGroupsApiService) ListWorkgroupConnectionsExecute(r ApiListWorkgroupConnectionsRequest) ([]WorkgroupConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkgroupConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.ListWorkgroupConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}/connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListWorkgroupMembersRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
}

func (r ApiListWorkgroupMembersRequest) Execute() ([]WorkgroupMember, *http.Response, error) {
	return r.ApiService.ListWorkgroupMembersExecute(r)
}

/*
ListWorkgroupMembers List Work Group Members

This API returns the members of a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWorkgroupMembersRequest
*/
func (a *GovernanceGroupsApiService) ListWorkgroupMembers(ctx context.Context) ApiListWorkgroupMembersRequest {
	return ApiListWorkgroupMembersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []WorkgroupMember
func (a *GovernanceGroupsApiService) ListWorkgroupMembersExecute(r ApiListWorkgroupMembersRequest) ([]WorkgroupMember, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkgroupMember
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.ListWorkgroupMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}/members"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListWorkgroupsRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
}

func (r ApiListWorkgroupsRequest) Execute() ([]Workgroup, *http.Response, error) {
	return r.ApiService.ListWorkgroupsExecute(r)
}

/*
ListWorkgroups List Work Groups

This API returns the details for a single account based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWorkgroupsRequest
*/
func (a *GovernanceGroupsApiService) ListWorkgroups(ctx context.Context) ApiListWorkgroupsRequest {
	return ApiListWorkgroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Workgroup
func (a *GovernanceGroupsApiService) ListWorkgroupsExecute(r ApiListWorkgroupsRequest) ([]Workgroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Workgroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.ListWorkgroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiModifyWorkgroupMembersRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
	modifyWorkgroupMembersRequest *ModifyWorkgroupMembersRequest
}

// Add/Remove workgroup member ids.
func (r ApiModifyWorkgroupMembersRequest) ModifyWorkgroupMembersRequest(modifyWorkgroupMembersRequest ModifyWorkgroupMembersRequest) ApiModifyWorkgroupMembersRequest {
	r.modifyWorkgroupMembersRequest = &modifyWorkgroupMembersRequest
	return r
}

func (r ApiModifyWorkgroupMembersRequest) Execute() (*http.Response, error) {
	return r.ApiService.ModifyWorkgroupMembersExecute(r)
}

/*
ModifyWorkgroupMembers Modify Work Group Members

This API allows you to modify the members of a work group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiModifyWorkgroupMembersRequest
*/
func (a *GovernanceGroupsApiService) ModifyWorkgroupMembers(ctx context.Context) ApiModifyWorkgroupMembersRequest {
	return ApiModifyWorkgroupMembersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *GovernanceGroupsApiService) ModifyWorkgroupMembersExecute(r ApiModifyWorkgroupMembersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.ModifyWorkgroupMembers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workgroups/{workgroupId}/members"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modifyWorkgroupMembersRequest == nil {
		return nil, reportError("modifyWorkgroupMembersRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.modifyWorkgroupMembersRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateWorkgroupRequest struct {
	ctx context.Context
	ApiService *GovernanceGroupsApiService
	createWorkgroupRequest *CreateWorkgroupRequest
}

// Work group to modify.
func (r ApiUpdateWorkgroupRequest) CreateWorkgroupRequest(createWorkgroupRequest CreateWorkgroupRequest) ApiUpdateWorkgroupRequest {
	r.createWorkgroupRequest = &createWorkgroupRequest
	return r
}

func (r ApiUpdateWorkgroupRequest) Execute() (*Workgroup, *http.Response, error) {
	return r.ApiService.UpdateWorkgroupExecute(r)
}

/*
UpdateWorkgroup Update Work Group By Id

This API updates and returns the details for a single workgroup based on the ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateWorkgroupRequest
*/
func (a *GovernanceGroupsApiService) UpdateWorkgroup(ctx context.Context) ApiUpdateWorkgroupRequest {
	return ApiUpdateWorkgroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Workgroup
func (a *GovernanceGroupsApiService) UpdateWorkgroupExecute(r ApiUpdateWorkgroupRequest) (*Workgroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Workgroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GovernanceGroupsApiService.UpdateWorkgroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workgroups/{workgroupId}"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createWorkgroupRequest == nil {
		return localVarReturnValue, nil, reportError("createWorkgroupRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createWorkgroupRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
