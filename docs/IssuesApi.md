# \IssuesApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssuesList**](IssuesApi.md#IssuesList) | **Get** /issues | 
[**IssuesRetrieve**](IssuesApi.md#IssuesRetrieve) | **Get** /issues/{id} | 



## IssuesList

> PaginatedIssueList IssuesList(ctx).AccountToken(accountToken).Cursor(cursor).EndDate(endDate).EndUserOrganizationName(endUserOrganizationName).FirstIncidentTimeAfter(firstIncidentTimeAfter).FirstIncidentTimeBefore(firstIncidentTimeBefore).IncludeMuted(includeMuted).IntegrationName(integrationName).LastIncidentTimeAfter(lastIncidentTimeAfter).LastIncidentTimeBefore(lastIncidentTimeBefore).PageSize(pageSize).StartDate(startDate).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountToken := "accountToken_example" // string |  (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    endDate := "endDate_example" // string | If included, will only include issues whose most recent action occurred before this time (optional)
    endUserOrganizationName := "endUserOrganizationName_example" // string |  (optional)
    firstIncidentTimeAfter := time.Now() // time.Time | If provided, will only return issues whose first incident time was after this datetime. (optional)
    firstIncidentTimeBefore := time.Now() // time.Time | If provided, will only return issues whose first incident time was before this datetime. (optional)
    includeMuted := "includeMuted_example" // string | If True, will include muted issues (optional)
    integrationName := "integrationName_example" // string |  (optional)
    lastIncidentTimeAfter := time.Now() // time.Time | If provided, will only return issues whose last incident time was after this datetime. (optional)
    lastIncidentTimeBefore := time.Now() // time.Time | If provided, will only return issues whose last incident time was before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    startDate := "startDate_example" // string | If included, will only include issues whose most recent action occurred after this time (optional)
    status := "status_example" // string | Status of the issue. Options: ('ONGOING', 'RESOLVED')  * `ONGOING` - ONGOING * `RESOLVED` - RESOLVED (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IssuesApi.IssuesList(context.Background()).AccountToken(accountToken).Cursor(cursor).EndDate(endDate).EndUserOrganizationName(endUserOrganizationName).FirstIncidentTimeAfter(firstIncidentTimeAfter).FirstIncidentTimeBefore(firstIncidentTimeBefore).IncludeMuted(includeMuted).IntegrationName(integrationName).LastIncidentTimeAfter(lastIncidentTimeAfter).LastIncidentTimeBefore(lastIncidentTimeBefore).PageSize(pageSize).StartDate(startDate).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesList`: PaginatedIssueList
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountToken** | **string** |  | 
 **cursor** | **string** | The pagination cursor value. | 
 **endDate** | **string** | If included, will only include issues whose most recent action occurred before this time | 
 **endUserOrganizationName** | **string** |  | 
 **firstIncidentTimeAfter** | **time.Time** | If provided, will only return issues whose first incident time was after this datetime. | 
 **firstIncidentTimeBefore** | **time.Time** | If provided, will only return issues whose first incident time was before this datetime. | 
 **includeMuted** | **string** | If True, will include muted issues | 
 **integrationName** | **string** |  | 
 **lastIncidentTimeAfter** | **time.Time** | If provided, will only return issues whose last incident time was after this datetime. | 
 **lastIncidentTimeBefore** | **time.Time** | If provided, will only return issues whose last incident time was before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **startDate** | **string** | If included, will only include issues whose most recent action occurred after this time | 
 **status** | **string** | Status of the issue. Options: (&#39;ONGOING&#39;, &#39;RESOLVED&#39;)  * &#x60;ONGOING&#x60; - ONGOING * &#x60;RESOLVED&#x60; - RESOLVED | 

### Return type

[**PaginatedIssueList**](PaginatedIssueList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesRetrieve

> Issue IssuesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IssuesApi.IssuesRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesRetrieve`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Issue**](Issue.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

