# \PayrollRunsApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayrollRunsList**](PayrollRunsApi.md#PayrollRunsList) | **Get** /payroll-runs | 
[**PayrollRunsRetrieve**](PayrollRunsApi.md#PayrollRunsRetrieve) | **Get** /payroll-runs/{id} | 



## PayrollRunsList

> PaginatedPayrollRunList PayrollRunsList(ctx).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EndedAfter(endedAfter).EndedBefore(endedBefore).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).RunType(runType).ShowEnumOrigins(showEnumOrigins).StartedAfter(startedAfter).StartedBefore(startedBefore).Execute()





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
    endedAfter := time.Now() // time.Time | If provided, will only return payroll runs ended after this datetime. (optional)
    endedBefore := time.Now() // time.Time | If provided, will only return payroll runs ended before this datetime. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteFields := "run_state,run_type" // string | Deprecated. Use show_enum_origins. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    runType := "runType_example" // string | If provided, will only return PayrollRun's with this status. Options: ('REGULAR', 'OFF_CYCLE', 'CORRECTION', 'TERMINATION', 'SIGN_ON_BONUS')  * `REGULAR` - REGULAR * `OFF_CYCLE` - OFF_CYCLE * `CORRECTION` - CORRECTION * `TERMINATION` - TERMINATION * `SIGN_ON_BONUS` - SIGN_ON_BONUS (optional)
    showEnumOrigins := "run_state,run_type" // string | Which fields should be returned in non-normalized form. (optional)
    startedAfter := time.Now() // time.Time | If provided, will only return payroll runs started after this datetime. (optional)
    startedBefore := time.Now() // time.Time | If provided, will only return payroll runs started before this datetime. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayrollRunsApi.PayrollRunsList(context.Background()).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EndedAfter(endedAfter).EndedBefore(endedBefore).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).RunType(runType).ShowEnumOrigins(showEnumOrigins).StartedAfter(startedAfter).StartedBefore(startedBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayrollRunsApi.PayrollRunsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayrollRunsList`: PaginatedPayrollRunList
    fmt.Fprintf(os.Stdout, "Response from `PayrollRunsApi.PayrollRunsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayrollRunsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **endedAfter** | **time.Time** | If provided, will only return payroll runs ended after this datetime. | 
 **endedBefore** | **time.Time** | If provided, will only return payroll runs ended before this datetime. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, only objects synced by Merge after this date time will be returned. | 
 **modifiedBefore** | **time.Time** | If provided, only objects synced by Merge before this date time will be returned. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **runType** | **string** | If provided, will only return PayrollRun&#39;s with this status. Options: (&#39;REGULAR&#39;, &#39;OFF_CYCLE&#39;, &#39;CORRECTION&#39;, &#39;TERMINATION&#39;, &#39;SIGN_ON_BONUS&#39;)  * &#x60;REGULAR&#x60; - REGULAR * &#x60;OFF_CYCLE&#x60; - OFF_CYCLE * &#x60;CORRECTION&#x60; - CORRECTION * &#x60;TERMINATION&#x60; - TERMINATION * &#x60;SIGN_ON_BONUS&#x60; - SIGN_ON_BONUS | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 
 **startedAfter** | **time.Time** | If provided, will only return payroll runs started after this datetime. | 
 **startedBefore** | **time.Time** | If provided, will only return payroll runs started before this datetime. | 

### Return type

[**PaginatedPayrollRunList**](PaginatedPayrollRunList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayrollRunsRetrieve

> PayrollRun PayrollRunsRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()





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
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    remoteFields := "run_state,run_type" // string | Deprecated. Use show_enum_origins. (optional)
    showEnumOrigins := "run_state,run_type" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayrollRunsApi.PayrollRunsRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayrollRunsApi.PayrollRunsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayrollRunsRetrieve`: PayrollRun
    fmt.Fprintf(os.Stdout, "Response from `PayrollRunsApi.PayrollRunsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayrollRunsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

### Return type

[**PayrollRun**](PayrollRun.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

