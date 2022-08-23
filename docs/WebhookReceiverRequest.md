# WebhookReceiverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**IsActive** | **bool** |  | 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookReceiverRequest

`func NewWebhookReceiverRequest(event string, isActive bool, ) *WebhookReceiverRequest`

NewWebhookReceiverRequest instantiates a new WebhookReceiverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookReceiverRequestWithDefaults

`func NewWebhookReceiverRequestWithDefaults() *WebhookReceiverRequest`

NewWebhookReceiverRequestWithDefaults instantiates a new WebhookReceiverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookReceiverRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookReceiverRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookReceiverRequest) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetIsActive

`func (o *WebhookReceiverRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhookReceiverRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhookReceiverRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetKey

`func (o *WebhookReceiverRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookReceiverRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookReceiverRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WebhookReceiverRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


