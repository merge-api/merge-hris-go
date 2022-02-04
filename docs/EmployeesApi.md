# \EmployeesApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmployeesIgnoreCreate**](EmployeesApi.md#EmployeesIgnoreCreate) | **Post** /employees/ignore/{model_id} | 
[**EmployeesList**](EmployeesApi.md#EmployeesList) | **Get** /employees | 
[**EmployeesRetrieve**](EmployeesApi.md#EmployeesRetrieve) | **Get** /employees/{id} | 



## EmployeesIgnoreCreate

> EmployeesIgnoreCreate(ctx, modelId).IgnoreCommonModelRequest(ignoreCommonModelRequest).Execute()





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
    modelId := TODO // string | 
    ignoreCommonModelRequest := *openapiclient.NewIgnoreCommonModelRequest(openapiclient.ReasonEnum("GENERAL_CUSTOMER_REQUEST")) // IgnoreCommonModelRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesIgnoreCreate(context.Background(), modelId).IgnoreCommonModelRequest(ignoreCommonModelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeesApi.EmployeesIgnoreCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmployeesIgnoreCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ignoreCommonModelRequest** | [**IgnoreCommonModelRequest**](IgnoreCommonModelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployeesList

> PaginatedEmployeeList EmployeesList(ctx).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).ManagerId(managerId).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).PayGroupId(payGroupId).PersonalEmail(personalEmail).RemoteId(remoteId).TeamId(teamId).WorkEmail(workEmail).WorkLocationId(workLocationId).Execute()





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
    companyId := "companyId_example" // string | If provided, will only return employees for this company. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was deleted in the third-party service. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    includeSensitiveFields := true // bool | Whether to include sensitive fields (such as social security numbers) in the response. (optional)
    managerId := "managerId_example" // string | If provided, will only return employees for this manager. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    payGroupId := "payGroupId_example" // string | If provided, will only return employees for this pay group (optional)
    personalEmail := "personalEmail@example.com" // string | If provided, will only return Employees with this personal email (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    teamId := "teamId_example" // string | If provided, will only return employees for this team. (optional)
    workEmail := "workEmail@example.com" // string | If provided, will only return Employees with this work email (optional)
    workLocationId := "workLocationId_example" // string | If provided, will only return employees for this location. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesList(context.Background()).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).ManagerId(managerId).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).PayGroupId(payGroupId).PersonalEmail(personalEmail).RemoteId(remoteId).TeamId(teamId).WorkEmail(workEmail).WorkLocationId(workLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeesApi.EmployeesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmployeesList`: PaginatedEmployeeList
    fmt.Fprintf(os.Stdout, "Response from `EmployeesApi.EmployeesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmployeesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **companyId** | **string** | If provided, will only return employees for this company. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was deleted in the third-party service. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeSensitiveFields** | **bool** | Whether to include sensitive fields (such as social security numbers) in the response. | 
 **managerId** | **string** | If provided, will only return employees for this manager. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **payGroupId** | **string** | If provided, will only return employees for this pay group | 
 **personalEmail** | [**string**](string.md) | If provided, will only return Employees with this personal email | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **teamId** | **string** | If provided, will only return employees for this team. | 
 **workEmail** | [**string**](string.md) | If provided, will only return Employees with this work email | 
 **workLocationId** | **string** | If provided, will only return employees for this location. | 

### Return type

[**PaginatedEmployeeList**](PaginatedEmployeeList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployeesRetrieve

> Employee EmployeesRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).Execute()





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
    includeSensitiveFields := true // bool | Whether to include sensitive fields (such as social security numbers) in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeesApi.EmployeesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmployeesRetrieve`: Employee
    fmt.Fprintf(os.Stdout, "Response from `EmployeesApi.EmployeesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmployeesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeSensitiveFields** | **bool** | Whether to include sensitive fields (such as social security numbers) in the response. | 

### Return type

[**Employee**](Employee.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

