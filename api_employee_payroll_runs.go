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

// EmployeePayrollRunsApiService EmployeePayrollRunsApi service
type EmployeePayrollRunsApiService service

type ApiEmployeePayrollRunsListRequest struct {
	ctx _context.Context
	ApiService *EmployeePayrollRunsApiService
	xAccountToken *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	employeeId *string
	endedAfter *time.Time
	endedBefore *time.Time
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	payrollRunId *string
	remoteId *string
	startedAfter *time.Time
	startedBefore *time.Time
}

func (r ApiEmployeePayrollRunsListRequest) XAccountToken(xAccountToken string) ApiEmployeePayrollRunsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeePayrollRunsListRequest) CreatedAfter(createdAfter time.Time) ApiEmployeePayrollRunsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiEmployeePayrollRunsListRequest) CreatedBefore(createdBefore time.Time) ApiEmployeePayrollRunsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiEmployeePayrollRunsListRequest) Cursor(cursor string) ApiEmployeePayrollRunsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiEmployeePayrollRunsListRequest) EmployeeId(employeeId string) ApiEmployeePayrollRunsListRequest {
	r.employeeId = &employeeId
	return r
}
func (r ApiEmployeePayrollRunsListRequest) EndedAfter(endedAfter time.Time) ApiEmployeePayrollRunsListRequest {
	r.endedAfter = &endedAfter
	return r
}
func (r ApiEmployeePayrollRunsListRequest) EndedBefore(endedBefore time.Time) ApiEmployeePayrollRunsListRequest {
	r.endedBefore = &endedBefore
	return r
}
func (r ApiEmployeePayrollRunsListRequest) IncludeRemoteData(includeRemoteData bool) ApiEmployeePayrollRunsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiEmployeePayrollRunsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiEmployeePayrollRunsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiEmployeePayrollRunsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiEmployeePayrollRunsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiEmployeePayrollRunsListRequest) PageSize(pageSize int32) ApiEmployeePayrollRunsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiEmployeePayrollRunsListRequest) PayrollRunId(payrollRunId string) ApiEmployeePayrollRunsListRequest {
	r.payrollRunId = &payrollRunId
	return r
}
func (r ApiEmployeePayrollRunsListRequest) RemoteId(remoteId string) ApiEmployeePayrollRunsListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiEmployeePayrollRunsListRequest) StartedAfter(startedAfter time.Time) ApiEmployeePayrollRunsListRequest {
	r.startedAfter = &startedAfter
	return r
}
func (r ApiEmployeePayrollRunsListRequest) StartedBefore(startedBefore time.Time) ApiEmployeePayrollRunsListRequest {
	r.startedBefore = &startedBefore
	return r
}

func (r ApiEmployeePayrollRunsListRequest) Execute() (PaginatedEmployeePayrollRunList, *_nethttp.Response, error) {
	return r.ApiService.EmployeePayrollRunsListExecute(r)
}

/*
 * EmployeePayrollRunsList Method for EmployeePayrollRunsList
 * Returns a list of `EmployeePayrollRun` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEmployeePayrollRunsListRequest
 */
func (a *EmployeePayrollRunsApiService) EmployeePayrollRunsList(ctx _context.Context) ApiEmployeePayrollRunsListRequest {
	return ApiEmployeePayrollRunsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedEmployeePayrollRunList
 */
func (a *EmployeePayrollRunsApiService) EmployeePayrollRunsListExecute(r ApiEmployeePayrollRunsListRequest) (PaginatedEmployeePayrollRunList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedEmployeePayrollRunList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeePayrollRunsApiService.EmployeePayrollRunsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employee-payroll-runs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
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
	if r.endedAfter != nil {
		localVarQueryParams.Add("ended_after", parameterToString(*r.endedAfter, ""))
	}
	if r.endedBefore != nil {
		localVarQueryParams.Add("ended_before", parameterToString(*r.endedBefore, ""))
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
	if r.payrollRunId != nil {
		localVarQueryParams.Add("payroll_run_id", parameterToString(*r.payrollRunId, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
	}
	if r.startedAfter != nil {
		localVarQueryParams.Add("started_after", parameterToString(*r.startedAfter, ""))
	}
	if r.startedBefore != nil {
		localVarQueryParams.Add("started_before", parameterToString(*r.startedBefore, ""))
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

type ApiEmployeePayrollRunsRetrieveRequest struct {
	ctx _context.Context
	ApiService *EmployeePayrollRunsApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiEmployeePayrollRunsRetrieveRequest) XAccountToken(xAccountToken string) ApiEmployeePayrollRunsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeePayrollRunsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiEmployeePayrollRunsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiEmployeePayrollRunsRetrieveRequest) Execute() (EmployeePayrollRun, *_nethttp.Response, error) {
	return r.ApiService.EmployeePayrollRunsRetrieveExecute(r)
}

/*
 * EmployeePayrollRunsRetrieve Method for EmployeePayrollRunsRetrieve
 * Returns an `EmployeePayrollRun` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiEmployeePayrollRunsRetrieveRequest
 */
func (a *EmployeePayrollRunsApiService) EmployeePayrollRunsRetrieve(ctx _context.Context, id string) ApiEmployeePayrollRunsRetrieveRequest {
	return ApiEmployeePayrollRunsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return EmployeePayrollRun
 */
func (a *EmployeePayrollRunsApiService) EmployeePayrollRunsRetrieveExecute(r ApiEmployeePayrollRunsRetrieveRequest) (EmployeePayrollRun, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EmployeePayrollRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeePayrollRunsApiService.EmployeePayrollRunsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employee-payroll-runs/{id}"
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
