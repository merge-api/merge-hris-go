/*
 * Merge HRIS API
 *
 * The unified API for building rich integrations with multiple HR Information System platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_hris_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// EmployeesApiService EmployeesApi service
type EmployeesApiService service

type ApiEmployeesListRequest struct {
	ctx _context.Context
	ApiService *EmployeesApiService
	xAccountToken *string
	companyId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	expand *string
	managerId *string
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteId *string
	teamId *string
	workLocationId *string
}

func (r ApiEmployeesListRequest) XAccountToken(xAccountToken string) ApiEmployeesListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeesListRequest) CompanyId(companyId string) ApiEmployeesListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiEmployeesListRequest) CreatedAfter(createdAfter time.Time) ApiEmployeesListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiEmployeesListRequest) CreatedBefore(createdBefore time.Time) ApiEmployeesListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiEmployeesListRequest) Cursor(cursor string) ApiEmployeesListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiEmployeesListRequest) Expand(expand string) ApiEmployeesListRequest {
	r.expand = &expand
	return r
}
func (r ApiEmployeesListRequest) ManagerId(managerId string) ApiEmployeesListRequest {
	r.managerId = &managerId
	return r
}
func (r ApiEmployeesListRequest) ModifiedAfter(modifiedAfter time.Time) ApiEmployeesListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiEmployeesListRequest) ModifiedBefore(modifiedBefore time.Time) ApiEmployeesListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiEmployeesListRequest) PageSize(pageSize int32) ApiEmployeesListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiEmployeesListRequest) RemoteId(remoteId string) ApiEmployeesListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiEmployeesListRequest) TeamId(teamId string) ApiEmployeesListRequest {
	r.teamId = &teamId
	return r
}
func (r ApiEmployeesListRequest) WorkLocationId(workLocationId string) ApiEmployeesListRequest {
	r.workLocationId = &workLocationId
	return r
}

func (r ApiEmployeesListRequest) Execute() (PaginatedEmployeeList, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.EmployeesListExecute(r)
}

/*
 * EmployeesList Method for EmployeesList
 * Returns a list of `Employee` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEmployeesListRequest
 */
func (a *EmployeesApiService) EmployeesList(ctx _context.Context) ApiEmployeesListRequest {
	return ApiEmployeesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedEmployeeList
 */
func (a *EmployeesApiService) EmployeesListExecute(r ApiEmployeesListRequest) (PaginatedEmployeeList, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  PaginatedEmployeeList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeesApiService.EmployeesList")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/employees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		executionError.error = "xAccountToken is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	if r.companyId != nil {
		localVarQueryParams.Add("company_id", parameterToString(*r.companyId, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("created_after", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created_before", parameterToString(*r.createdBefore, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
	}
	if r.managerId != nil {
		localVarQueryParams.Add("manager_id", parameterToString(*r.managerId, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
	}
	if r.teamId != nil {
		localVarQueryParams.Add("team_id", parameterToString(*r.teamId, ""))
	}
	if r.workLocationId != nil {
		localVarQueryParams.Add("work_location_id", parameterToString(*r.workLocationId, ""))
	}
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiEmployeesRetrieveRequest struct {
	ctx _context.Context
	ApiService *EmployeesApiService
	xAccountToken *string
	id string
	expand *string
}

func (r ApiEmployeesRetrieveRequest) XAccountToken(xAccountToken string) ApiEmployeesRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeesRetrieveRequest) Expand(expand string) ApiEmployeesRetrieveRequest {
	r.expand = &expand
	return r
}

func (r ApiEmployeesRetrieveRequest) Execute() (Employee, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.EmployeesRetrieveExecute(r)
}

/*
 * EmployeesRetrieve Method for EmployeesRetrieve
 * Returns an `Employee` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiEmployeesRetrieveRequest
 */
func (a *EmployeesApiService) EmployeesRetrieve(ctx _context.Context, id string) ApiEmployeesRetrieveRequest {
	return ApiEmployeesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Employee
 */
func (a *EmployeesApiService) EmployeesRetrieveExecute(r ApiEmployeesRetrieveRequest) (Employee, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Employee
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeesApiService.EmployeesRetrieve")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/employees/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		executionError.error = "xAccountToken is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
	}
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}