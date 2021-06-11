# \RegenerateKeyApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegenerateKeyCreate**](RegenerateKeyApi.md#RegenerateKeyCreate) | **Post** /regenerate-key | 



## RegenerateKeyCreate

> RemoteKey RegenerateKeyCreate(ctx).RemoteKeyForRegenerationRequest(remoteKeyForRegenerationRequest).Execute()





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
    remoteKeyForRegenerationRequest := *openapiclient.NewRemoteKeyForRegenerationRequest("Remote Deployment Key 1") // RemoteKeyForRegenerationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegenerateKeyApi.RegenerateKeyCreate(context.Background()).RemoteKeyForRegenerationRequest(remoteKeyForRegenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegenerateKeyApi.RegenerateKeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegenerateKeyCreate`: RemoteKey
    fmt.Fprintf(os.Stdout, "Response from `RegenerateKeyApi.RegenerateKeyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteKeyForRegenerationRequest** | [**RemoteKeyForRegenerationRequest**](RemoteKeyForRegenerationRequest.md) |  | 

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

