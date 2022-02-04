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

// TimeOffApiService TimeOffApi service
type TimeOffApiService service

type ApiTimeOffCreateRequest struct {
	ctx _context.Context
	ApiService *TimeOffApiService
	xAccountToken *string
	timeOffEndpointRequest *TimeOffEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiTimeOffCreateRequest) XAccountToken(xAccountToken string) ApiTimeOffCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiTimeOffCreateRequest) TimeOffEndpointRequest(timeOffEndpointRequest TimeOffEndpointRequest) ApiTimeOffCreateRequest {
	r.timeOffEndpointRequest = &timeOffEndpointRequest
	return r
}
func (r ApiTimeOffCreateRequest) IsDebugMode(isDebugMode bool) ApiTimeOffCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiTimeOffCreateRequest) RunAsync(runAsync bool) ApiTimeOffCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiTimeOffCreateRequest) Execute() (TimeOffResponse, *_nethttp.Response, error) {
	return r.ApiService.TimeOffCreateExecute(r)
}

/*
 * TimeOffCreate Method for TimeOffCreate
 * Creates a `TimeOff` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTimeOffCreateRequest
 */
func (a *TimeOffApiService) TimeOffCreate(ctx _context.Context) ApiTimeOffCreateRequest {
	return ApiTimeOffCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TimeOffResponse
 */
func (a *TimeOffApiService) TimeOffCreateExecute(r ApiTimeOffCreateRequest) (TimeOffResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TimeOffResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeOffApiService.TimeOffCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time-off"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.timeOffEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("timeOffEndpointRequest is required and must be specified")
	}

	if r.isDebugMode != nil {
		localVarQueryParams.Add("is_debug_mode", parameterToString(*r.isDebugMode, ""))
	}
	if r.runAsync != nil {
		localVarQueryParams.Add("run_async", parameterToString(*r.runAsync, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.timeOffEndpointRequest
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTimeOffListRequest struct {
	ctx _context.Context
	ApiService *TimeOffApiService
	xAccountToken *string
	approverId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	employeeId *string
	includeDeletedData *bool
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteId *string
	requestType *string
	status *string
}

func (r ApiTimeOffListRequest) XAccountToken(xAccountToken string) ApiTimeOffListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiTimeOffListRequest) ApproverId(approverId string) ApiTimeOffListRequest {
	r.approverId = &approverId
	return r
}
func (r ApiTimeOffListRequest) CreatedAfter(createdAfter time.Time) ApiTimeOffListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiTimeOffListRequest) CreatedBefore(createdBefore time.Time) ApiTimeOffListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiTimeOffListRequest) Cursor(cursor string) ApiTimeOffListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiTimeOffListRequest) EmployeeId(employeeId string) ApiTimeOffListRequest {
	r.employeeId = &employeeId
	return r
}
func (r ApiTimeOffListRequest) IncludeDeletedData(includeDeletedData bool) ApiTimeOffListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiTimeOffListRequest) IncludeRemoteData(includeRemoteData bool) ApiTimeOffListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiTimeOffListRequest) ModifiedAfter(modifiedAfter time.Time) ApiTimeOffListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiTimeOffListRequest) ModifiedBefore(modifiedBefore time.Time) ApiTimeOffListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiTimeOffListRequest) PageSize(pageSize int32) ApiTimeOffListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiTimeOffListRequest) RemoteId(remoteId string) ApiTimeOffListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiTimeOffListRequest) RequestType(requestType string) ApiTimeOffListRequest {
	r.requestType = &requestType
	return r
}
func (r ApiTimeOffListRequest) Status(status string) ApiTimeOffListRequest {
	r.status = &status
	return r
}

func (r ApiTimeOffListRequest) Execute() (PaginatedTimeOffList, *_nethttp.Response, error) {
	return r.ApiService.TimeOffListExecute(r)
}

/*
 * TimeOffList Method for TimeOffList
 * Returns a list of `TimeOff` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTimeOffListRequest
 */
func (a *TimeOffApiService) TimeOffList(ctx _context.Context) ApiTimeOffListRequest {
	return ApiTimeOffListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedTimeOffList
 */
func (a *TimeOffApiService) TimeOffListExecute(r ApiTimeOffListRequest) (PaginatedTimeOffList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedTimeOffList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeOffApiService.TimeOffList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time-off"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.approverId != nil {
		localVarQueryParams.Add("approver_id", parameterToString(*r.approverId, ""))
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
	if r.employeeId != nil {
		localVarQueryParams.Add("employee_id", parameterToString(*r.employeeId, ""))
	}
	if r.includeDeletedData != nil {
		localVarQueryParams.Add("include_deleted_data", parameterToString(*r.includeDeletedData, ""))
	}
	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
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
	if r.requestType != nil {
		localVarQueryParams.Add("request_type", parameterToString(*r.requestType, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTimeOffRetrieveRequest struct {
	ctx _context.Context
	ApiService *TimeOffApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiTimeOffRetrieveRequest) XAccountToken(xAccountToken string) ApiTimeOffRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiTimeOffRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiTimeOffRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiTimeOffRetrieveRequest) Execute() (TimeOff, *_nethttp.Response, error) {
	return r.ApiService.TimeOffRetrieveExecute(r)
}

/*
 * TimeOffRetrieve Method for TimeOffRetrieve
 * Returns a `TimeOff` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiTimeOffRetrieveRequest
 */
func (a *TimeOffApiService) TimeOffRetrieve(ctx _context.Context, id string) ApiTimeOffRetrieveRequest {
	return ApiTimeOffRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return TimeOff
 */
func (a *TimeOffApiService) TimeOffRetrieveExecute(r ApiTimeOffRetrieveRequest) (TimeOff, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TimeOff
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeOffApiService.TimeOffRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time-off/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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

	return localVarReturnValue, localVarHTTPResponse, nil
}
