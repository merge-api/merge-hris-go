# AccountDetailsAndActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Category** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**StatusDetail** | Pointer to **string** |  | [optional] 
**EndUserOriginId** | Pointer to **string** |  | [optional] 
**EndUserOrganizationName** | **string** |  | 
**EndUserEmailAddress** | **string** |  | 
**Integration** | Pointer to [**AccountDetailsAndActionsIntegration**](AccountDetailsAndActionsIntegration.md) |  | [optional] 

## Methods

### NewAccountDetailsAndActions

`func NewAccountDetailsAndActions(id string, status string, endUserOrganizationName string, endUserEmailAddress string, ) *AccountDetailsAndActions`

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

`func (o *AccountDetailsAndActions) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AccountDetailsAndActions) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AccountDetailsAndActions) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AccountDetailsAndActions) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *AccountDetailsAndActions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountDetailsAndActions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountDetailsAndActions) SetStatus(v string)`

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


