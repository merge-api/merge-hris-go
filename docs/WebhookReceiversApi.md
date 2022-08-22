# \WebhookReceiversApi

All URIs are relative to *https://api.merge.dev/api/hris/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookReceiversCreate**](WebhookReceiversApi.md#WebhookReceiversCreate) | **Post** /webhook-receivers | 
[**WebhookReceiversList**](WebhookReceiversApi.md#WebhookReceiversList) | **Get** /webhook-receivers | 



## WebhookReceiversCreate

> WebhookReceiver WebhookReceiversCreate(ctx).XAccountToken(xAccountToken).WebhookReceiverRequest(webhookReceiverRequest).Execute()





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
    webhookReceiverRequest := *openapiclient.NewWebhookReceiverRequest("Event_example", false) // WebhookReceiverRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookReceiversApi.WebhookReceiversCreate(context.Background()).XAccountToken(xAccountToken).WebhookReceiverRequest(webhookReceiverRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookReceiversApi.WebhookReceiversCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookReceiversCreate`: WebhookReceiver
    fmt.Fprintf(os.Stdout, "Response from `WebhookReceiversApi.WebhookReceiversCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookReceiversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **webhookReceiverRequest** | [**WebhookReceiverRequest**](WebhookReceiverRequest.md) |  | 

### Return type

[**WebhookReceiver**](WebhookReceiver.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookReceiversList

> []WebhookReceiver WebhookReceiversList(ctx).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.WebhookReceiversApi.WebhookReceiversList(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookReceiversApi.WebhookReceiversList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookReceiversList`: []WebhookReceiver
    fmt.Fprintf(os.Stdout, "Response from `WebhookReceiversApi.WebhookReceiversList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookReceiversListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

### Return type

[**[]WebhookReceiver**](WebhookReceiver.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

