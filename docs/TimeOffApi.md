# \TimeOffApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimeOffList**](TimeOffApi.md#TimeOffList) | **Get** /time-off | 
[**TimeOffRetrieve**](TimeOffApi.md#TimeOffRetrieve) | **Get** /time-off/{id} | 



## TimeOffList

> PaginatedTimeOffList TimeOffList(ctx).XAccountToken(xAccountToken).ApproverId(approverId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EmployeeId(employeeId).Expand(expand).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    approverId := "approverId_example" // string | If provided, will only return time off for this approver. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    employeeId := "employeeId_example" // string | If provided, will only return time off for this employee. (optional)
    expand := "employee,approver" // string | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeOffApi.TimeOffList(context.Background()).XAccountToken(xAccountToken).ApproverId(approverId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EmployeeId(employeeId).Expand(expand).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeOffApi.TimeOffList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeOffList`: PaginatedTimeOffList
    fmt.Fprintf(os.Stdout, "Response from `TimeOffApi.TimeOffList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimeOffListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **approverId** | **string** | If provided, will only return time off for this approver. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **employeeId** | **string** | If provided, will only return time off for this employee. | 
 **expand** | **string** | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedTimeOffList**](PaginatedTimeOffList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimeOffRetrieve

> TimeOff TimeOffRetrieve(ctx, id).XAccountToken(xAccountToken).Expand(expand).Execute()





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
    expand := "employee,approver" // string | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeOffApi.TimeOffRetrieve(context.Background(), id).XAccountToken(xAccountToken).Expand(expand).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeOffApi.TimeOffRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeOffRetrieve`: TimeOff
    fmt.Fprintf(os.Stdout, "Response from `TimeOffApi.TimeOffRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimeOffRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **expand** | **string** | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. | 

### Return type

[**TimeOff**](TimeOff.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

