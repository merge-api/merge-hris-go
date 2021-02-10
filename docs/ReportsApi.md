# \ReportsApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsList**](ReportsApi.md#ReportsList) | **Get** /reports | 
[**ReportsRetrieve**](ReportsApi.md#ReportsRetrieve) | **Get** /reports/{id} | 



## ReportsList

> PaginatedReportList ReportsList(ctx).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ReportsList(context.Background()).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsList`: PaginatedReportList
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedReportList**](PaginatedReportList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsRetrieve

> Report ReportsRetrieve(ctx, id).XAccountToken(xAccountToken).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ReportsRetrieve(context.Background(), id).XAccountToken(xAccountToken).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsRetrieve`: Report
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


### Return type

[**Report**](Report.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
