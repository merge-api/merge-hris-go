# \LinkedAccountsApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LinkedAccountsList**](LinkedAccountsApi.md#LinkedAccountsList) | **Get** /linked-accounts | 



## LinkedAccountsList

> PaginatedAccountDetailsAndActionsList LinkedAccountsList(ctx).Category(category).Cursor(cursor).EndUserEmailAddress(endUserEmailAddress).EndUserOrganizationName(endUserOrganizationName).EndUserOriginId(endUserOriginId).EndUserOriginIds(endUserOriginIds).Id(id).Ids(ids).IncludeDuplicates(includeDuplicates).IntegrationName(integrationName).IsTestAccount(isTestAccount).PageSize(pageSize).Status(status).Execute()





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
    category := "category_example" // string | Options: ('hris', 'ats', 'accounting', 'ticketing', 'crm', 'mktg', 'filestorage')  * `hris` - hris * `ats` - ats * `accounting` - accounting * `ticketing` - ticketing * `crm` - crm * `mktg` - mktg * `filestorage` - filestorage (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    endUserEmailAddress := "endUserEmailAddress_example" // string | If provided, will only return linked accounts associated with the given email address. (optional)
    endUserOrganizationName := "endUserOrganizationName_example" // string | If provided, will only return linked accounts associated with the given organization name. (optional)
    endUserOriginId := "endUserOriginId_example" // string | If provided, will only return linked accounts associated with the given origin ID. (optional)
    endUserOriginIds := "endUserOriginIds_example" // string | Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once. (optional)
    id := TODO // string |  (optional)
    ids := "ids_example" // string | Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once. (optional)
    includeDuplicates := true // bool | If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account. (optional)
    integrationName := "integrationName_example" // string | If provided, will only return linked accounts associated with the given integration name. (optional)
    isTestAccount := "isTestAccount_example" // string | If included, will only include test linked accounts. If not included, will only include non-test linked accounts. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    status := "status_example" // string | Filter by status. Options: `COMPLETE`, `INCOMPLETE`, `RELINK_NEEDED` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkedAccountsApi.LinkedAccountsList(context.Background()).Category(category).Cursor(cursor).EndUserEmailAddress(endUserEmailAddress).EndUserOrganizationName(endUserOrganizationName).EndUserOriginId(endUserOriginId).EndUserOriginIds(endUserOriginIds).Id(id).Ids(ids).IncludeDuplicates(includeDuplicates).IntegrationName(integrationName).IsTestAccount(isTestAccount).PageSize(pageSize).Status(status).Execute()
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
 **category** | **string** | Options: (&#39;hris&#39;, &#39;ats&#39;, &#39;accounting&#39;, &#39;ticketing&#39;, &#39;crm&#39;, &#39;mktg&#39;, &#39;filestorage&#39;)  * &#x60;hris&#x60; - hris * &#x60;ats&#x60; - ats * &#x60;accounting&#x60; - accounting * &#x60;ticketing&#x60; - ticketing * &#x60;crm&#x60; - crm * &#x60;mktg&#x60; - mktg * &#x60;filestorage&#x60; - filestorage | 
 **cursor** | **string** | The pagination cursor value. | 
 **endUserEmailAddress** | **string** | If provided, will only return linked accounts associated with the given email address. | 
 **endUserOrganizationName** | **string** | If provided, will only return linked accounts associated with the given organization name. | 
 **endUserOriginId** | **string** | If provided, will only return linked accounts associated with the given origin ID. | 
 **endUserOriginIds** | **string** | Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once. | 
 **id** | [**string**](string.md) |  | 
 **ids** | **string** | Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once. | 
 **includeDuplicates** | **bool** | If &#x60;true&#x60;, will include complete production duplicates of the account specified by the &#x60;id&#x60; query parameter in the response. &#x60;id&#x60; must be for a complete production linked account. | 
 **integrationName** | **string** | If provided, will only return linked accounts associated with the given integration name. | 
 **isTestAccount** | **string** | If included, will only include test linked accounts. If not included, will only include non-test linked accounts. | 
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

