# \SelectiveSyncApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SelectiveSyncConfigurationsList**](SelectiveSyncApi.md#SelectiveSyncConfigurationsList) | **Get** /selective-sync/configurations | 
[**SelectiveSyncConfigurationsUpdate**](SelectiveSyncApi.md#SelectiveSyncConfigurationsUpdate) | **Put** /selective-sync/configurations | 
[**SelectiveSyncMetaList**](SelectiveSyncApi.md#SelectiveSyncMetaList) | **Get** /selective-sync/meta | 



## SelectiveSyncConfigurationsList

> []LinkedAccountSelectiveSyncConfiguration SelectiveSyncConfigurationsList(ctx).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.SelectiveSyncApi.SelectiveSyncConfigurationsList(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncApi.SelectiveSyncConfigurationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelectiveSyncConfigurationsList`: []LinkedAccountSelectiveSyncConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncApi.SelectiveSyncConfigurationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectiveSyncConfigurationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

### Return type

[**[]LinkedAccountSelectiveSyncConfiguration**](LinkedAccountSelectiveSyncConfiguration.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectiveSyncConfigurationsUpdate

> []LinkedAccountSelectiveSyncConfiguration SelectiveSyncConfigurationsUpdate(ctx).XAccountToken(xAccountToken).LinkedAccountSelectiveSyncConfigurationListRequest(linkedAccountSelectiveSyncConfigurationListRequest).Execute()





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
    linkedAccountSelectiveSyncConfigurationListRequest := *openapiclient.NewLinkedAccountSelectiveSyncConfigurationListRequest([]openapiclient.LinkedAccountSelectiveSyncConfigurationRequest{*openapiclient.NewLinkedAccountSelectiveSyncConfigurationRequest([]openapiclient.LinkedAccountConditionRequest{*openapiclient.NewLinkedAccountConditionRequest("ConditionSchemaId_example", "Operator_example", interface{}(123))})}) // LinkedAccountSelectiveSyncConfigurationListRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelectiveSyncApi.SelectiveSyncConfigurationsUpdate(context.Background()).XAccountToken(xAccountToken).LinkedAccountSelectiveSyncConfigurationListRequest(linkedAccountSelectiveSyncConfigurationListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncApi.SelectiveSyncConfigurationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelectiveSyncConfigurationsUpdate`: []LinkedAccountSelectiveSyncConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncApi.SelectiveSyncConfigurationsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectiveSyncConfigurationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **linkedAccountSelectiveSyncConfigurationListRequest** | [**LinkedAccountSelectiveSyncConfigurationListRequest**](LinkedAccountSelectiveSyncConfigurationListRequest.md) |  | 

### Return type

[**[]LinkedAccountSelectiveSyncConfiguration**](LinkedAccountSelectiveSyncConfiguration.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectiveSyncMetaList

> PaginatedConditionSchemaList SelectiveSyncMetaList(ctx).XAccountToken(xAccountToken).CommonModel(commonModel).Cursor(cursor).PageSize(pageSize).Execute()





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
    commonModel := "commonModel_example" // string |  (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelectiveSyncApi.SelectiveSyncMetaList(context.Background()).XAccountToken(xAccountToken).CommonModel(commonModel).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncApi.SelectiveSyncMetaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelectiveSyncMetaList`: PaginatedConditionSchemaList
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncApi.SelectiveSyncMetaList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectiveSyncMetaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **commonModel** | **string** |  | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedConditionSchemaList**](PaginatedConditionSchemaList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

