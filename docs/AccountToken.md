# AccountToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** |  | 
**Integration** | [**AccountIntegration**](AccountIntegration.md) |  | 

## Methods

### NewAccountToken

`func NewAccountToken(accountToken string, integration AccountIntegration, ) *AccountToken`

NewAccountToken instantiates a new AccountToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenWithDefaults

`func NewAccountTokenWithDefaults() *AccountToken`

NewAccountTokenWithDefaults instantiates a new AccountToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *AccountToken) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *AccountToken) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *AccountToken) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetIntegration

`func (o *AccountToken) GetIntegration() AccountIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AccountToken) GetIntegrationOk() (*AccountIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AccountToken) SetIntegration(v AccountIntegration)`

SetIntegration sets Integration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


