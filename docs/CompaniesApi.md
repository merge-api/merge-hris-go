# \CompaniesApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesList**](CompaniesApi.md#CompaniesList) | **Get** /companies | 
[**CompaniesRetrieve**](CompaniesApi.md#CompaniesRetrieve) | **Get** /companies/{id} | 



## CompaniesList

> PaginatedCompanyList CompaniesList(ctx).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    resp, r, err := api_client.CompaniesApi.CompaniesList(context.Background()).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CompaniesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompaniesList`: PaginatedCompanyList
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CompaniesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesListRequest struct via the builder pattern


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

[**PaginatedCompanyList**](PaginatedCompanyList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesRetrieve

> Company CompaniesRetrieve(ctx, id).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.CompaniesApi.CompaniesRetrieve(context.Background(), id).XAccountToken(xAccountToken).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CompaniesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompaniesRetrieve`: Company
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CompaniesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


### Return type

[**Company**](Company.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

