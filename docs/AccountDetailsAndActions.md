# AccountDetailsAndActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Category** | Pointer to [**CategoryEnum**](CategoryEnum.md) |  | [optional] 
**Status** | [**AccountDetailsAndActionsStatusEnum**](AccountDetailsAndActionsStatusEnum.md) |  | 
**StatusDetail** | Pointer to **string** |  | [optional] 
**EndUserOriginId** | Pointer to **string** |  | [optional] 
**EndUserOrganizationName** | **string** |  | 
**EndUserEmailAddress** | **string** |  | 
**WebhookListenerUrl** | **string** |  | 
**IsDuplicate** | Pointer to **NullableBool** | Whether a Production Linked Account&#39;s credentials match another existing Production Linked Account. This field is &#x60;null&#x60; for Test Linked Accounts, incomplete Production Linked Accounts, and ignored duplicate Production Linked Account sets. | [optional] 
**Integration** | Pointer to [**AccountDetailsAndActionsIntegration**](AccountDetailsAndActionsIntegration.md) |  | [optional] 

## Methods

### NewAccountDetailsAndActions

`func NewAccountDetailsAndActions(id string, status AccountDetailsAndActionsStatusEnum, endUserOrganizationName string, endUserEmailAddress string, webhookListenerUrl string, ) *AccountDetailsAndActions`

NewAccountDetailsAndActions instantiates a new AccountDetailsAndActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsAndActionsWithDefaults

`func NewAccountDetailsAndActionsWithDefaults() *AccountDetailsAndActions`

NewAccountDetailsAndActionsWithDefaults instantiates a new AccountDetailsAndActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountDetailsAndActions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDetailsAndActions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDetailsAndActions) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *AccountDetailsAndActions) GetCategory() CategoryEnum`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AccountDetailsAndActions) GetCategoryOk() (*CategoryEnum, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AccountDetailsAndActions) SetCategory(v CategoryEnum)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AccountDetailsAndActions) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *AccountDetailsAndActions) GetStatus() AccountDetailsAndActionsStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountDetailsAndActions) GetStatusOk() (*AccountDetailsAndActionsStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountDetailsAndActions) SetStatus(v AccountDetailsAndActionsStatusEnum)`

SetStatus sets Status field to given value.


### GetStatusDetail

`func (o *AccountDetailsAndActions) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *AccountDetailsAndActions) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *AccountDetailsAndActions) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *AccountDetailsAndActions) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetEndUserOriginId

`func (o *AccountDetailsAndActions) GetEndUserOriginId() string`

GetEndUserOriginId returns the EndUserOriginId field if non-nil, zero value otherwise.

### GetEndUserOriginIdOk

`func (o *AccountDetailsAndActions) GetEndUserOriginIdOk() (*string, bool)`

GetEndUserOriginIdOk returns a tuple with the EndUserOriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserOriginId

`func (o *AccountDetailsAndActions) SetEndUserOriginId(v string)`

SetEndUserOriginId sets EndUserOriginId field to given value.

### HasEndUserOriginId

`func (o *AccountDetailsAndActions) HasEndUserOriginId() bool`

HasEndUserOriginId returns a boolean if a field has been set.

### GetEndUserOrganizationName

`func (o *AccountDetailsAndActions) GetEndUserOrganizationName() string`

GetEndUserOrganizationName returns the EndUserOrganizationName field if non-nil, zero value otherwise.

### GetEndUserOrganizationNameOk

`func (o *AccountDetailsAndActions) GetEndUserOrganizationNameOk() (*string, bool)`

GetEndUserOrganizationNameOk returns a tuple with the EndUserOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserOrganizationName

`func (o *AccountDetailsAndActions) SetEndUserOrganizationName(v string)`

SetEndUserOrganizationName sets EndUserOrganizationName field to given value.


### GetEndUserEmailAddress

`func (o *AccountDetailsAndActions) GetEndUserEmailAddress() string`

GetEndUserEmailAddress returns the EndUserEmailAddress field if non-nil, zero value otherwise.

### GetEndUserEmailAddressOk

`func (o *AccountDetailsAndActions) GetEndUserEmailAddressOk() (*string, bool)`

GetEndUserEmailAddressOk returns a tuple with the EndUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserEmailAddress

`func (o *AccountDetailsAndActions) SetEndUserEmailAddress(v string)`

SetEndUserEmailAddress sets EndUserEmailAddress field to given value.


### GetWebhookListenerUrl

`func (o *AccountDetailsAndActions) GetWebhookListenerUrl() string`

GetWebhookListenerUrl returns the WebhookListenerUrl field if non-nil, zero value otherwise.

### GetWebhookListenerUrlOk

`func (o *AccountDetailsAndActions) GetWebhookListenerUrlOk() (*string, bool)`

GetWebhookListenerUrlOk returns a tuple with the WebhookListenerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookListenerUrl

`func (o *AccountDetailsAndActions) SetWebhookListenerUrl(v string)`

SetWebhookListenerUrl sets WebhookListenerUrl field to given value.


### GetIsDuplicate

`func (o *AccountDetailsAndActions) GetIsDuplicate() bool`

GetIsDuplicate returns the IsDuplicate field if non-nil, zero value otherwise.

### GetIsDuplicateOk

`func (o *AccountDetailsAndActions) GetIsDuplicateOk() (*bool, bool)`

GetIsDuplicateOk returns a tuple with the IsDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDuplicate

`func (o *AccountDetailsAndActions) SetIsDuplicate(v bool)`

SetIsDuplicate sets IsDuplicate field to given value.

### HasIsDuplicate

`func (o *AccountDetailsAndActions) HasIsDuplicate() bool`

HasIsDuplicate returns a boolean if a field has been set.

### SetIsDuplicateNil

`func (o *AccountDetailsAndActions) SetIsDuplicateNil(b bool)`

 SetIsDuplicateNil sets the value for IsDuplicate to be an explicit nil

### UnsetIsDuplicate
`func (o *AccountDetailsAndActions) UnsetIsDuplicate()`

UnsetIsDuplicate ensures that no value is present for IsDuplicate, not even an explicit nil
### GetIntegration

`func (o *AccountDetailsAndActions) GetIntegration() AccountDetailsAndActionsIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AccountDetailsAndActions) GetIntegrationOk() (*AccountDetailsAndActionsIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AccountDetailsAndActions) SetIntegration(v AccountDetailsAndActionsIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AccountDetailsAndActions) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


