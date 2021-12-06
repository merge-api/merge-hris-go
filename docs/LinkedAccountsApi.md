# \LinkedAccountsApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LinkedAccountsList**](LinkedAccountsApi.md#LinkedAccountsList) | **Get** /linked-accounts | 



## LinkedAccountsList

> PaginatedAccountDetailsAndActionsList LinkedAccountsList(ctx).Category(category).Cursor(cursor).EndUserEmailAddress(endUserEmailAddress).EndUserOrganizationName(endUserOrganizationName).EndUserOriginId(endUserOriginId).EndUserOriginIds(endUserOriginIds).Id(id).Ids(ids).IntegrationName(integrationName).IsTestAccount(isTestAccount).PageSize(pageSize).Status(status).Execute()





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
    category := "category_example" // string |  (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    endUserEmailAddress := "endUserEmailAddress_example" // string |  (optional)
    endUserOrganizationName := "endUserOrganizationName_example" // string |  (optional)
    endUserOriginId := "endUserOriginId_example" // string |  (optional)
    endUserOriginIds := "endUserOriginIds_example" // string | Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once (optional)
    id := TODO // string |  (optional)
    ids := "ids_example" // string | Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once (optional)
    integrationName := "integrationName_example" // string |  (optional)
    isTestAccount := "isTestAccount_example" // string | If included, will only include test linked accounts. If not included, will only include non-test linked accounts (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    status := "status_example" // string | Filter by status. Options: `COMPLETE`, `INCOMPLETE`, `RELINK_NEEDED` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkedAccountsApi.LinkedAccountsList(context.Background()).Category(category).Cursor(cursor).EndUserEmailAddress(endUserEmailAddress).EndUserOrganizationName(endUserOrganizationName).EndUserOriginId(endUserOriginId).EndUserOriginIds(endUserOriginIds).Id(id).Ids(ids).IntegrationName(integrationName).IsTestAccount(isTestAccount).PageSize(pageSize).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedAccountsApi.LinkedAccountsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkedAccountsList`: PaginatedAccountDetailsAndActionsList
    fmt.Fprintf(os.Stdout, "Response from `LinkedAccountsApi.LinkedAccountsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkedAccountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** |  | 
 **cursor** | **string** | The pagination cursor value. | 
 **endUserEmailAddress** | **string** |  | 
 **endUserOrganizationName** | **string** |  | 
 **endUserOriginId** | **string** |  | 
 **endUserOriginIds** | **string** | Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once | 
 **id** | [**string**](string.md) |  | 
 **ids** | **string** | Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once | 
 **integrationName** | **string** |  | 
 **isTestAccount** | **string** | If included, will only include test linked accounts. If not included, will only include non-test linked accounts | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **status** | **string** | Filter by status. Options: &#x60;COMPLETE&#x60;, &#x60;INCOMPLETE&#x60;, &#x60;RELINK_NEEDED&#x60; | 

### Return type

[**PaginatedAccountDetailsAndActionsList**](PaginatedAccountDetailsAndActionsList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

