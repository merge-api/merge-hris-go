# \TimeOffApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimeOffCreate**](TimeOffApi.md#TimeOffCreate) | **Post** /time-off | 
[**TimeOffList**](TimeOffApi.md#TimeOffList) | **Get** /time-off | 
[**TimeOffMetaPostRetrieve**](TimeOffApi.md#TimeOffMetaPostRetrieve) | **Get** /time-off/meta/post | 
[**TimeOffRetrieve**](TimeOffApi.md#TimeOffRetrieve) | **Get** /time-off/{id} | 



## TimeOffCreate

> TimeOffResponse TimeOffCreate(ctx).XAccountToken(xAccountToken).TimeOffEndpointRequest(timeOffEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    timeOffEndpointRequest := *openapiclient.NewTimeOffEndpointRequest(*openapiclient.NewTimeOffRequest()) // TimeOffEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeOffApi.TimeOffCreate(context.Background()).XAccountToken(xAccountToken).TimeOffEndpointRequest(timeOffEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeOffApi.TimeOffCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeOffCreate`: TimeOffResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeOffApi.TimeOffCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimeOffCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **timeOffEndpointRequest** | [**TimeOffEndpointRequest**](TimeOffEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**TimeOffResponse**](TimeOffResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimeOffList

> PaginatedTimeOffList TimeOffList(ctx).XAccountToken(xAccountToken).ApproverId(approverId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EmployeeId(employeeId).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).RequestType(requestType).ShowEnumOrigins(showEnumOrigins).Status(status).Execute()





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
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteFields := "request_type,status,units" // string | Deprecated. Use show_enum_origins. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    requestType := "requestType_example" // string | If provided, will only return TimeOff with this request type. Options: ('VACATION', 'SICK', 'PERSONAL', 'JURY_DUTY', 'VOLUNTEER', 'BEREAVEMENT')  * `VACATION` - VACATION * `SICK` - SICK * `PERSONAL` - PERSONAL * `JURY_DUTY` - JURY_DUTY * `VOLUNTEER` - VOLUNTEER * `BEREAVEMENT` - BEREAVEMENT (optional)
    showEnumOrigins := "request_type,status,units" // string | Which fields should be returned in non-normalized form. (optional)
    status := "status_example" // string | If provided, will only return TimeOff with this status. Options: ('REQUESTED', 'APPROVED', 'DECLINED', 'CANCELLED', 'DELETED')  * `REQUESTED` - REQUESTED * `APPROVED` - APPROVED * `DECLINED` - DECLINED * `CANCELLED` - CANCELLED * `DELETED` - DELETED (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeOffApi.TimeOffList(context.Background()).XAccountToken(xAccountToken).ApproverId(approverId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EmployeeId(employeeId).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).RequestType(requestType).ShowEnumOrigins(showEnumOrigins).Status(status).Execute()
    if err != nil {
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
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, only objects synced by Merge after this date time will be returned. | 
 **modifiedBefore** | **time.Time** | If provided, only objects synced by Merge before this date time will be returned. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **requestType** | **string** | If provided, will only return TimeOff with this request type. Options: (&#39;VACATION&#39;, &#39;SICK&#39;, &#39;PERSONAL&#39;, &#39;JURY_DUTY&#39;, &#39;VOLUNTEER&#39;, &#39;BEREAVEMENT&#39;)  * &#x60;VACATION&#x60; - VACATION * &#x60;SICK&#x60; - SICK * &#x60;PERSONAL&#x60; - PERSONAL * &#x60;JURY_DUTY&#x60; - JURY_DUTY * &#x60;VOLUNTEER&#x60; - VOLUNTEER * &#x60;BEREAVEMENT&#x60; - BEREAVEMENT | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 
 **status** | **string** | If provided, will only return TimeOff with this status. Options: (&#39;REQUESTED&#39;, &#39;APPROVED&#39;, &#39;DECLINED&#39;, &#39;CANCELLED&#39;, &#39;DELETED&#39;)  * &#x60;REQUESTED&#x60; - REQUESTED * &#x60;APPROVED&#x60; - APPROVED * &#x60;DECLINED&#x60; - DECLINED * &#x60;CANCELLED&#x60; - CANCELLED * &#x60;DELETED&#x60; - DELETED | 

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


## TimeOffMetaPostRetrieve

> MetaResponse TimeOffMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeOffApi.TimeOffMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeOffApi.TimeOffMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeOffMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeOffApi.TimeOffMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimeOffMetaPostRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

### Return type

[**MetaResponse**](MetaResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimeOffRetrieve

> TimeOff TimeOffRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()





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
    remoteFields := "request_type,status,units" // string | Deprecated. Use show_enum_origins. (optional)
    showEnumOrigins := "request_type,status,units" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeOffApi.TimeOffRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()
    if err != nil {
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

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

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

