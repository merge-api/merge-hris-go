# \PassthroughApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PassthroughCreate**](PassthroughApi.md#PassthroughCreate) | **Post** /passthrough | 



## PassthroughCreate

> RemoteResponse PassthroughCreate(ctx).XAccountToken(xAccountToken).DataPassthroughRequest(dataPassthroughRequest).Execute()





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
    dataPassthroughRequest := *openapiclient.NewDataPassthroughRequest("POST", "/scooters") // DataPassthroughRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PassthroughApi.PassthroughCreate(context.Background()).XAccountToken(xAccountToken).DataPassthroughRequest(dataPassthroughRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassthroughApi.PassthroughCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PassthroughCreate`: RemoteResponse
    fmt.Fprintf(os.Stdout, "Response from `PassthroughApi.PassthroughCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPassthroughCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **dataPassthroughRequest** | [**DataPassthroughRequest**](DataPassthroughRequest.md) |  | 

### Return type

[**RemoteResponse**](RemoteResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

