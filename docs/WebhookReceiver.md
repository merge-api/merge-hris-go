# WebhookReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**IsActive** | **bool** |  | 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookReceiver

`func NewWebhookReceiver(event string, isActive bool, ) *WebhookReceiver`

NewWebhookReceiver instantiates a new WebhookReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookReceiverWithDefaults

`func NewWebhookReceiverWithDefaults() *WebhookReceiver`

NewWebhookReceiverWithDefaults instantiates a new WebhookReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookReceiver) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookReceiver) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookReceiver) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetIsActive

`func (o *WebhookReceiver) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhookReceiver) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhookReceiver) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetKey

`func (o *WebhookReceiver) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookReceiver) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookReceiver) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WebhookReceiver) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


