# \GenerateKeyApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateKeyCreate**](GenerateKeyApi.md#GenerateKeyCreate) | **Post** /generate-key | 



## GenerateKeyCreate

> RemoteKey GenerateKeyCreate(ctx).GenerateRemoteKeyRequest(generateRemoteKeyRequest).Execute()





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
    generateRemoteKeyRequest := *openapiclient.NewGenerateRemoteKeyRequest("Remote Deployment Key 1") // GenerateRemoteKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenerateKeyApi.GenerateKeyCreate(context.Background()).GenerateRemoteKeyRequest(generateRemoteKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateKeyApi.GenerateKeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateKeyCreate`: RemoteKey
    fmt.Fprintf(os.Stdout, "Response from `GenerateKeyApi.GenerateKeyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateRemoteKeyRequest** | [**GenerateRemoteKeyRequest**](GenerateRemoteKeyRequest.md) |  | 

### Return type

[**RemoteKey**](RemoteKey.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

