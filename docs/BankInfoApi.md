# \BankInfoApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankInfoList**](BankInfoApi.md#BankInfoList) | **Get** /bank-info | 
[**BankInfoRetrieve**](BankInfoApi.md#BankInfoRetrieve) | **Get** /bank-info/{id} | 



## BankInfoList

> PaginatedBankInfoList BankInfoList(ctx).XAccountToken(xAccountToken).AccountType(accountType).BankName(bankName).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).Employee(employee).EmployeeId(employeeId).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).OrderBy(orderBy).PageSize(pageSize).RemoteCreatedAt(remoteCreatedAt).RemoteId(remoteId).Execute()





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
    accountType := "accountType_example" // string | The bank account type: [CHECKING, SAVINGS] (optional)
    bankName := "bankName_example" // string |  (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    employee := TODO // string | If provided, will only return bank accounts for this employee. (optional)
    employeeId := "employeeId_example" // string | If provided, will only return bank accounts for this employee. (optional)
    includeDeletedData := true // bool | Whether to include data that was deleted in the third-party service. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    orderBy := "orderBy_example" // string | Overrides the default ordering for this endpoint. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteCreatedAt := time.Now() // time.Time |  (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BankInfoApi.BankInfoList(context.Background()).XAccountToken(xAccountToken).AccountType(accountType).BankName(bankName).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).Employee(employee).EmployeeId(employeeId).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).OrderBy(orderBy).PageSize(pageSize).RemoteCreatedAt(remoteCreatedAt).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankInfoApi.BankInfoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankInfoList`: PaginatedBankInfoList
    fmt.Fprintf(os.Stdout, "Response from `BankInfoApi.BankInfoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **accountType** | **string** | The bank account type: [CHECKING, SAVINGS] | 
 **bankName** | **string** |  | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **employee** | [**string**](string.md) | If provided, will only return bank accounts for this employee. | 
 **employeeId** | **string** | If provided, will only return bank accounts for this employee. | 
 **includeDeletedData** | **bool** | Whether to include data that was deleted in the third-party service. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **orderBy** | **string** | Overrides the default ordering for this endpoint. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteCreatedAt** | **time.Time** |  | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedBankInfoList**](PaginatedBankInfoList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankInfoRetrieve

> BankInfo BankInfoRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BankInfoApi.BankInfoRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BankInfoApi.BankInfoRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankInfoRetrieve`: BankInfo
    fmt.Fprintf(os.Stdout, "Response from `BankInfoApi.BankInfoRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankInfoRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 

### Return type

[**BankInfo**](BankInfo.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

