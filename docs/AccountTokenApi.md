# \AccountTokenApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountTokenRetrieve**](AccountTokenApi.md#AccountTokenRetrieve) | **Get** /account-token/{public_token} | 



## AccountTokenRetrieve

> AccountToken AccountTokenRetrieve(ctx, publicToken).Execute()





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
    publicToken := "publicToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTokenApi.AccountTokenRetrieve(context.Background(), publicToken).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTokenApi.AccountTokenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountTokenRetrieve`: AccountToken
    fmt.Fprintf(os.Stdout, "Response from `AccountTokenApi.AccountTokenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountTokenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountToken**](AccountToken.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

