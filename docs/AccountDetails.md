# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Integration** | Pointer to **string** |  | [optional] [readonly] 
**IntegrationSlug** | Pointer to **string** |  | [optional] [readonly] 
**Category** | Pointer to [**NullableCategoryEnum**](CategoryEnum.md) |  | [optional] 
**EndUserOriginId** | Pointer to **string** |  | [optional] [readonly] 
**EndUserOrganizationName** | Pointer to **string** |  | [optional] [readonly] 
**EndUserEmailAddress** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**WebhookListenerUrl** | Pointer to **string** |  | [optional] [readonly] 
**IsDuplicate** | Pointer to **NullableBool** | Whether a Production Linked Account&#39;s credentials match another existing Production Linked Account. This field is &#x60;null&#x60; for Test Linked Accounts, incomplete Production Linked Accounts, and ignored duplicate Production Linked Account sets. | [optional] [readonly] 

## Methods

### NewAccountDetails

`func NewAccountDetails() *AccountDetails`

NewAccountDetails instantiates a new AccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsWithDefaults

`func NewAccountDetailsWithDefaults() *AccountDetails`

NewAccountDetailsWithDefaults instantiates a new AccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *AccountDetails) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AccountDetails) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AccountDetails) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AccountDetails) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationSlug

`func (o *AccountDetails) GetIntegrationSlug() string`

GetIntegrationSlug returns the IntegrationSlug field if non-nil, zero value otherwise.

### GetIntegrationSlugOk

`func (o *AccountDetails) GetIntegrationSlugOk() (*string, bool)`

GetIntegrationSlugOk returns a tuple with the IntegrationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationSlug

`func (o *AccountDetails) SetIntegrationSlug(v string)`

SetIntegrationSlug sets IntegrationSlug field to given value.

### HasIntegrationSlug

`func (o *AccountDetails) HasIntegrationSlug() bool`

HasIntegrationSlug returns a boolean if a field has been set.

### GetCategory

`func (o *AccountDetails) GetCategory() CategoryEnum`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AccountDetails) GetCategoryOk() (*CategoryEnum, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AccountDetails) SetCategory(v CategoryEnum)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AccountDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *AccountDetails) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AccountDetails) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetEndUserOriginId

`func (o *AccountDetails) GetEndUserOriginId() string`

GetEndUserOriginId returns the EndUserOriginId field if non-nil, zero value otherwise.

### GetEndUserOriginIdOk

`func (o *AccountDetails) GetEndUserOriginIdOk() (*string, bool)`

GetEndUserOriginIdOk returns a tuple with the EndUserOriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserOriginId

`func (o *AccountDetails) SetEndUserOriginId(v string)`

SetEndUserOriginId sets EndUserOriginId field to given value.

### HasEndUserOriginId

`func (o *AccountDetails) HasEndUserOriginId() bool`

HasEndUserOriginId returns a boolean if a field has been set.

### GetEndUserOrganizationName

`func (o *AccountDetails) GetEndUserOrganizationName() string`

GetEndUserOrganizationName returns the EndUserOrganizationName field if non-nil, zero value otherwise.

### GetEndUserOrganizationNameOk

`func (o *AccountDetails) GetEndUserOrganizationNameOk() (*string, bool)`

GetEndUserOrganizationNameOk returns a tuple with the EndUserOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserOrganizationName

`func (o *AccountDetails) SetEndUserOrganizationName(v string)`

SetEndUserOrganizationName sets EndUserOrganizationName field to given value.

### HasEndUserOrganizationName

`func (o *AccountDetails) HasEndUserOrganizationName() bool`

HasEndUserOrganizationName returns a boolean if a field has been set.

### GetEndUserEmailAddress

`func (o *AccountDetails) GetEndUserEmailAddress() string`

GetEndUserEmailAddress returns the EndUserEmailAddress field if non-nil, zero value otherwise.

### GetEndUserEmailAddressOk

`func (o *AccountDetails) GetEndUserEmailAddressOk() (*string, bool)`

GetEndUserEmailAddressOk returns a tuple with the EndUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserEmailAddress

`func (o *AccountDetails) SetEndUserEmailAddress(v string)`

SetEndUserEmailAddress sets EndUserEmailAddress field to given value.

### HasEndUserEmailAddress

`func (o *AccountDetails) HasEndUserEmailAddress() bool`

HasEndUserEmailAddress returns a boolean if a field has been set.

### GetStatus

`func (o *AccountDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebhookListenerUrl

`func (o *AccountDetails) GetWebhookListenerUrl() string`

GetWebhookListenerUrl returns the WebhookListenerUrl field if non-nil, zero value otherwise.

### GetWebhookListenerUrlOk

`func (o *AccountDetails) GetWebhookListenerUrlOk() (*string, bool)`

GetWebhookListenerUrlOk returns a tuple with the WebhookListenerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookListenerUrl

`func (o *AccountDetails) SetWebhookListenerUrl(v string)`

SetWebhookListenerUrl sets WebhookListenerUrl field to given value.

### HasWebhookListenerUrl

`func (o *AccountDetails) HasWebhookListenerUrl() bool`

HasWebhookListenerUrl returns a boolean if a field has been set.

### GetIsDuplicate

`func (o *AccountDetails) GetIsDuplicate() bool`

GetIsDuplicate returns the IsDuplicate field if non-nil, zero value otherwise.

### GetIsDuplicateOk

`func (o *AccountDetails) GetIsDuplicateOk() (*bool, bool)`

GetIsDuplicateOk returns a tuple with the IsDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDuplicate

`func (o *AccountDetails) SetIsDuplicate(v bool)`

SetIsDuplicate sets IsDuplicate field to given value.

### HasIsDuplicate

`func (o *AccountDetails) HasIsDuplicate() bool`

HasIsDuplicate returns a boolean if a field has been set.

### SetIsDuplicateNil

`func (o *AccountDetails) SetIsDuplicateNil(b bool)`

 SetIsDuplicateNil sets the value for IsDuplicate to be an explicit nil

### UnsetIsDuplicate
`func (o *AccountDetails) UnsetIsDuplicate()`

UnsetIsDuplicate ensures that no value is present for IsDuplicate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


