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

// IssuesApiService IssuesApi service
type IssuesApiService service

type ApiIssuesListRequest struct {
	ctx _context.Context
	ApiService *IssuesApiService
	accountToken *string
	cursor *string
	endDate *string
	endUserOrganizationName *string
	firstIncidentTimeAfter *time.Time
	firstIncidentTimeBefore *time.Time
	includeMuted *string
	integrationName *string
	lastIncidentTimeAfter *time.Time
	lastIncidentTimeBefore *time.Time
	pageSize *int32
	startDate *string
	status *string
}

func (r ApiIssuesListRequest) AccountToken(accountToken string) ApiIssuesListRequest {
	r.accountToken = &accountToken
	return r
}
func (r ApiIssuesListRequest) Cursor(cursor string) ApiIssuesListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiIssuesListRequest) EndDate(endDate string) ApiIssuesListRequest {
	r.endDate = &endDate
	return r
}
func (r ApiIssuesListRequest) EndUserOrganizationName(endUserOrganizationName string) ApiIssuesListRequest {
	r.endUserOrganizationName = &endUserOrganizationName
	return r
}
func (r ApiIssuesListRequest) FirstIncidentTimeAfter(firstIncidentTimeAfter time.Time) ApiIssuesListRequest {
	r.firstIncidentTimeAfter = &firstIncidentTimeAfter
	return r
}
func (r ApiIssuesListRequest) FirstIncidentTimeBefore(firstIncidentTimeBefore time.Time) ApiIssuesListRequest {
	r.firstIncidentTimeBefore = &firstIncidentTimeBefore
	return r
}
func (r ApiIssuesListRequest) IncludeMuted(includeMuted string) ApiIssuesListRequest {
	r.includeMuted = &includeMuted
	return r
}
func (r ApiIssuesListRequest) IntegrationName(integrationName string) ApiIssuesListRequest {
	r.integrationName = &integrationName
	return r
}
func (r ApiIssuesListRequest) LastIncidentTimeAfter(lastIncidentTimeAfter time.Time) ApiIssuesListRequest {
	r.lastIncidentTimeAfter = &lastIncidentTimeAfter
	return r
}
func (r ApiIssuesListRequest) LastIncidentTimeBefore(lastIncidentTimeBefore time.Time) ApiIssuesListRequest {
	r.lastIncidentTimeBefore = &lastIncidentTimeBefore
	return r
}
func (r ApiIssuesListRequest) PageSize(pageSize int32) ApiIssuesListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiIssuesListRequest) StartDate(startDate string) ApiIssuesListRequest {
	r.startDate = &startDate
	return r
}
func (r ApiIssuesListRequest) Status(status string) ApiIssuesListRequest {
	r.status = &status
	return r
}

func (r ApiIssuesListRequest) Execute() (PaginatedIssueList, *_nethttp.Response, error) {
	return r.ApiService.IssuesListExecute(r)
}

/*
 * IssuesList Method for IssuesList
 * Gets issues.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIssuesListRequest
 */
func (a *IssuesApiService) IssuesList(ctx _context.Context) ApiIssuesListRequest {
	return ApiIssuesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedIssueList
 */
func (a *IssuesApiService) IssuesListExecute(r ApiIssuesListRequest) (PaginatedIssueList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedIssueList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesApiService.IssuesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.accountToken != nil {
		localVarQueryParams.Add("account_token", parameterToString(*r.accountToken, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("end_date", parameterToString(*r.endDate, ""))
	}
	if r.endUserOrganizationName != nil {
		localVarQueryParams.Add("end_user_organization_name", parameterToString(*r.endUserOrganizationName, ""))
	}
	if r.firstIncidentTimeAfter != nil {
		localVarQueryParams.Add("first_incident_time_after", parameterToString(*r.firstIncidentTimeAfter, ""))
	}
	if r.firstIncidentTimeBefore != nil {
		localVarQueryParams.Add("first_incident_time_before", parameterToString(*r.firstIncidentTimeBefore, ""))
	}
	if r.includeMuted != nil {
		localVarQueryParams.Add("include_muted", parameterToString(*r.includeMuted, ""))
	}
	if r.integrationName != nil {
		localVarQueryParams.Add("integration_name", parameterToString(*r.integrationName, ""))
	}
	if r.lastIncidentTimeAfter != nil {
		localVarQueryParams.Add("last_incident_time_after", parameterToString(*r.lastIncidentTimeAfter, ""))
	}
	if r.lastIncidentTimeBefore != nil {
		localVarQueryParams.Add("last_incident_time_before", parameterToString(*r.lastIncidentTimeBefore, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("start_date", parameterToString(*r.startDate, ""))
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

type ApiIssuesRetrieveRequest struct {
	ctx _context.Context
	ApiService *IssuesApiService
	id string
}


func (r ApiIssuesRetrieveRequest) Execute() (Issue, *_nethttp.Response, error) {
	return r.ApiService.IssuesRetrieveExecute(r)
}

/*
 * IssuesRetrieve Method for IssuesRetrieve
 * Get a specific issue.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIssuesRetrieveRequest
 */
func (a *IssuesApiService) IssuesRetrieve(ctx _context.Context, id string) ApiIssuesRetrieveRequest {
	return ApiIssuesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Issue
 */
func (a *IssuesApiService) IssuesRetrieveExecute(r ApiIssuesRetrieveRequest) (Issue, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Issue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesApiService.IssuesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
