# \EmployeesApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmployeesCreate**](EmployeesApi.md#EmployeesCreate) | **Post** /employees | 
[**EmployeesIgnoreCreate**](EmployeesApi.md#EmployeesIgnoreCreate) | **Post** /employees/ignore/{model_id} | 
[**EmployeesList**](EmployeesApi.md#EmployeesList) | **Get** /employees | 
[**EmployeesMetaPostRetrieve**](EmployeesApi.md#EmployeesMetaPostRetrieve) | **Get** /employees/meta/post | 
[**EmployeesRetrieve**](EmployeesApi.md#EmployeesRetrieve) | **Get** /employees/{id} | 



## EmployeesCreate

> EmployeeResponse EmployeesCreate(ctx).XAccountToken(xAccountToken).EmployeeEndpointRequest(employeeEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    employeeEndpointRequest := *openapiclient.NewEmployeeEndpointRequest(*openapiclient.NewEmployeeRequest()) // EmployeeEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesCreate(context.Background()).XAccountToken(xAccountToken).EmployeeEndpointRequest(employeeEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeesApi.EmployeesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmployeesCreate`: EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `EmployeesApi.EmployeesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmployeesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **employeeEndpointRequest** | [**EmployeeEndpointRequest**](EmployeeEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployeesIgnoreCreate

> EmployeesIgnoreCreate(ctx, modelId).XAccountToken(xAccountToken).IgnoreCommonModelRequest(ignoreCommonModelRequest).Execute()





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
    modelId := TODO // string | 
    ignoreCommonModelRequest := *openapiclient.NewIgnoreCommonModelRequest(openapiclient.ReasonEnum("GENERAL_CUSTOMER_REQUEST")) // IgnoreCommonModelRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesIgnoreCreate(context.Background(), modelId).XAccountToken(xAccountToken).IgnoreCommonModelRequest(ignoreCommonModelRequest).Execute()
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
 **xAccountToken** | **string** | Token identifying the end user. | 

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

> PaginatedEmployeeList EmployeesList(ctx).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).DisplayFullName(displayFullName).EmploymentStatus(employmentStatus).FirstName(firstName).Groups(groups).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).LastName(lastName).ManagerId(managerId).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).PayGroupId(payGroupId).PersonalEmail(personalEmail).RemoteFields(remoteFields).RemoteId(remoteId).ShowEnumOrigins(showEnumOrigins).StartedAfter(startedAfter).StartedBefore(startedBefore).TeamId(teamId).TerminatedAfter(terminatedAfter).TerminatedBefore(terminatedBefore).WorkEmail(workEmail).WorkLocationId(workLocationId).Execute()





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
    displayFullName := "displayFullName_example" // string | If provided, will only return employees with this display name. (optional)
    employmentStatus := "employmentStatus_example" // string | If provided, will only return employees with this employment status.  * `ACTIVE` - ACTIVE * `PENDING` - PENDING * `INACTIVE` - INACTIVE (optional)
    firstName := "firstName_example" // string | If provided, will only return employees with this first name. (optional)
    groups := "groups_example" // string | If provided, will only return employees matching the group ids; multiple groups can be separated by commas. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    includeSensitiveFields := true // bool | Whether to include sensitive fields (such as social security numbers) in the response. (optional)
    lastName := "lastName_example" // string | If provided, will only return employees with this last name. (optional)
    managerId := "managerId_example" // string | If provided, will only return employees for this manager. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    payGroupId := "payGroupId_example" // string | If provided, will only return employees for this pay group (optional)
    personalEmail := "personalEmail@example.com" // string | If provided, will only return Employees with this personal email (optional)
    remoteFields := "employment_status,ethnicity,gender,marital_status" // string | Deprecated. Use show_enum_origins. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    showEnumOrigins := "employment_status,ethnicity,gender,marital_status" // string | Which fields should be returned in non-normalized form. (optional)
    startedAfter := time.Now() // time.Time | If provided, will only return employees that started after this datetime. (optional)
    startedBefore := time.Now() // time.Time | If provided, will only return employees that started before this datetime. (optional)
    teamId := "teamId_example" // string | If provided, will only return employees for this team. (optional)
    terminatedAfter := time.Now() // time.Time | If provided, will only return employees that were terminated after this datetime. (optional)
    terminatedBefore := time.Now() // time.Time | If provided, will only return employees that were terminated before this datetime. (optional)
    workEmail := "workEmail@example.com" // string | If provided, will only return Employees with this work email (optional)
    workLocationId := "workLocationId_example" // string | If provided, will only return employees for this location. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesList(context.Background()).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).DisplayFullName(displayFullName).EmploymentStatus(employmentStatus).FirstName(firstName).Groups(groups).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).LastName(lastName).ManagerId(managerId).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).PayGroupId(payGroupId).PersonalEmail(personalEmail).RemoteFields(remoteFields).RemoteId(remoteId).ShowEnumOrigins(showEnumOrigins).StartedAfter(startedAfter).StartedBefore(startedBefore).TeamId(teamId).TerminatedAfter(terminatedAfter).TerminatedBefore(terminatedBefore).WorkEmail(workEmail).WorkLocationId(workLocationId).Execute()
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
 **displayFullName** | **string** | If provided, will only return employees with this display name. | 
 **employmentStatus** | **string** | If provided, will only return employees with this employment status.  * &#x60;ACTIVE&#x60; - ACTIVE * &#x60;PENDING&#x60; - PENDING * &#x60;INACTIVE&#x60; - INACTIVE | 
 **firstName** | **string** | If provided, will only return employees with this first name. | 
 **groups** | **string** | If provided, will only return employees matching the group ids; multiple groups can be separated by commas. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeSensitiveFields** | **bool** | Whether to include sensitive fields (such as social security numbers) in the response. | 
 **lastName** | **string** | If provided, will only return employees with this last name. | 
 **managerId** | **string** | If provided, will only return employees for this manager. | 
 **modifiedAfter** | **time.Time** | If provided, only objects synced by Merge after this date time will be returned. | 
 **modifiedBefore** | **time.Time** | If provided, only objects synced by Merge before this date time will be returned. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **payGroupId** | **string** | If provided, will only return employees for this pay group | 
 **personalEmail** | [**string**](string.md) | If provided, will only return Employees with this personal email | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 
 **startedAfter** | **time.Time** | If provided, will only return employees that started after this datetime. | 
 **startedBefore** | **time.Time** | If provided, will only return employees that started before this datetime. | 
 **teamId** | **string** | If provided, will only return employees for this team. | 
 **terminatedAfter** | **time.Time** | If provided, will only return employees that were terminated after this datetime. | 
 **terminatedBefore** | **time.Time** | If provided, will only return employees that were terminated before this datetime. | 
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


## EmployeesMetaPostRetrieve

> MetaResponse EmployeesMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.EmployeesApi.EmployeesMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeesApi.EmployeesMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmployeesMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `EmployeesApi.EmployeesMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmployeesMetaPostRetrieveRequest struct via the builder pattern


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


## EmployeesRetrieve

> Employee EmployeesRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()





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
    remoteFields := "employment_status,ethnicity,gender,marital_status" // string | Deprecated. Use show_enum_origins. (optional)
    showEnumOrigins := "employment_status,ethnicity,gender,marital_status" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmployeesApi.EmployeesRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeSensitiveFields(includeSensitiveFields).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()
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
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

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

