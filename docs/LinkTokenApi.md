# \LinkTokenApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LinkTokenCreate**](LinkTokenApi.md#LinkTokenCreate) | **Post** /link-token | 



## LinkTokenCreate

> LinkToken LinkTokenCreate(ctx).EndUserDetails(endUserDetails).Execute()





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
    endUserDetails := *openapiclient.NewEndUserDetails("EndUserEmailAddress_example", "EndUserOrganizationName_example", "EndUserOriginId_example", []string{"Categories_example"}) // EndUserDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkTokenApi.LinkTokenCreate(context.Background()).EndUserDetails(endUserDetails).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkTokenApi.LinkTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkTokenCreate`: LinkToken
    fmt.Fprintf(os.Stdout, "Response from `LinkTokenApi.LinkTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endUserDetails** | [**EndUserDetails**](EndUserDetails.md) |  | 

### Return type

[**LinkToken**](LinkToken.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

