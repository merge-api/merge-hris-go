# EndUserDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserEmailAddress** | **string** | Your end user&#39;s email address. This is purely for identification purposes - setting this value will not cause any emails to be sent. | 
**EndUserOrganizationName** | **string** | Your end user&#39;s organization. | 
**EndUserOriginId** | **string** | This unique identifier typically represents the ID for your end user in your product&#39;s database. This value must be distinct from other Linked Accounts&#39; unique identifiers. | 
**Categories** | [**[]CategoriesEnum**](CategoriesEnum.md) | The integration categories to show in Merge Link. | 
**Integration** | Pointer to **NullableString** | The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://www.merge.dev/docs/basics/integration-metadata/. | [optional] 
**LinkExpiryMins** | Pointer to **int32** | An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30. | [optional] [default to 30]
**ShouldCreateMagicLinkUrl** | Pointer to **NullableBool** | Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/product/integrations,-fast.-say-hello-to-magic-link/. | [optional] [default to false]
**CommonModels** | Pointer to [**[]CommonModelScopesBodyRequest**](CommonModelScopesBodyRequest.md) | An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account. | [optional] 

## Methods

### NewEndUserDetailsRequest

`func NewEndUserDetailsRequest(endUserEmailAddress string, endUserOrganizationName string, endUserOriginId string, categories []CategoriesEnum, ) *EndUserDetailsRequest`

NewEndUserDetailsRequest instantiates a new EndUserDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserDetailsRequestWithDefaults

`func NewEndUserDetailsRequestWithDefaults() *EndUserDetailsRequest`

NewEndUserDetailsRequestWithDefaults instantiates a new EndUserDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserEmailAddress

`func (o *EndUserDetailsRequest) GetEndUserEmailAddress() string`

GetEndUserEmailAddress returns the EndUserEmailAddress field if non-nil, zero value otherwise.

### GetEndUserEmailAddressOk

`func (o *EndUserDetailsRequest) GetEndUserEmailAddressOk() (*string, bool)`

GetEndUserEmailAddressOk returns a tuple with the EndUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserEmailAddress

`func (o *EndUserDetailsRequest) SetEndUserEmailAddress(v string)`

SetEndUserEmailAddress sets EndUserEmailAddress field to given value.


### GetEndUserOrganizationName

`func (o *EndUserDetailsRequest) GetEndUserOrganizationName() string`

GetEndUserOrganizationName returns the EndUserOrganizationName field if non-nil, zero value otherwise.

### GetEndUserOrganizationNameOk

`func (o *EndUserDetailsRequest) GetEndUserOrganizationNameOk() (*string, bool)`

GetEndUserOrganizationNameOk returns a tuple with the EndUserOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserOrganizationName

`func (o *EndUserDetailsRequest) SetEndUserOrganizationName(v string)`

SetEndUserOrganizationName sets EndUserOrganizationName field to given value.


### GetEndUserOriginId

`func (o *EndUserDetailsRequest) GetEndUserOriginId() string`

GetEndUserOriginId returns the EndUserOriginId field if non-nil, zero value otherwise.

### GetEndUserOriginIdOk

`func (o *EndUserDetailsRequest) GetEndUserOriginIdOk() (*string, bool)`

GetEndUserOriginIdOk returns a tuple with the EndUserOriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserOriginId

`func (o *EndUserDetailsRequest) SetEndUserOriginId(v string)`

SetEndUserOriginId sets EndUserOriginId field to given value.


### GetCategories

`func (o *EndUserDetailsRequest) GetCategories() []CategoriesEnum`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *EndUserDetailsRequest) GetCategoriesOk() (*[]CategoriesEnum, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *EndUserDetailsRequest) SetCategories(v []CategoriesEnum)`

SetCategories sets Categories field to given value.


### GetIntegration

`func (o *EndUserDetailsRequest) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *EndUserDetailsRequest) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *EndUserDetailsRequest) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *EndUserDetailsRequest) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *EndUserDetailsRequest) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *EndUserDetailsRequest) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetLinkExpiryMins

`func (o *EndUserDetailsRequest) GetLinkExpiryMins() int32`

GetLinkExpiryMins returns the LinkExpiryMins field if non-nil, zero value otherwise.

### GetLinkExpiryMinsOk

`func (o *EndUserDetailsRequest) GetLinkExpiryMinsOk() (*int32, bool)`

GetLinkExpiryMinsOk returns a tuple with the LinkExpiryMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkExpiryMins

`func (o *EndUserDetailsRequest) SetLinkExpiryMins(v int32)`

SetLinkExpiryMins sets LinkExpiryMins field to given value.

### HasLinkExpiryMins

`func (o *EndUserDetailsRequest) HasLinkExpiryMins() bool`

HasLinkExpiryMins returns a boolean if a field has been set.

### GetShouldCreateMagicLinkUrl

`func (o *EndUserDetailsRequest) GetShouldCreateMagicLinkUrl() bool`

GetShouldCreateMagicLinkUrl returns the ShouldCreateMagicLinkUrl field if non-nil, zero value otherwise.

### GetShouldCreateMagicLinkUrlOk

`func (o *EndUserDetailsRequest) GetShouldCreateMagicLinkUrlOk() (*bool, bool)`

GetShouldCreateMagicLinkUrlOk returns a tuple with the ShouldCreateMagicLinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateMagicLinkUrl

`func (o *EndUserDetailsRequest) SetShouldCreateMagicLinkUrl(v bool)`

SetShouldCreateMagicLinkUrl sets ShouldCreateMagicLinkUrl field to given value.

### HasShouldCreateMagicLinkUrl

`func (o *EndUserDetailsRequest) HasShouldCreateMagicLinkUrl() bool`

HasShouldCreateMagicLinkUrl returns a boolean if a field has been set.

### SetShouldCreateMagicLinkUrlNil

`func (o *EndUserDetailsRequest) SetShouldCreateMagicLinkUrlNil(b bool)`

 SetShouldCreateMagicLinkUrlNil sets the value for ShouldCreateMagicLinkUrl to be an explicit nil

### UnsetShouldCreateMagicLinkUrl
`func (o *EndUserDetailsRequest) UnsetShouldCreateMagicLinkUrl()`

UnsetShouldCreateMagicLinkUrl ensures that no value is present for ShouldCreateMagicLinkUrl, not even an explicit nil
### GetCommonModels

`func (o *EndUserDetailsRequest) GetCommonModels() []CommonModelScopesBodyRequest`

GetCommonModels returns the CommonModels field if non-nil, zero value otherwise.

### GetCommonModelsOk

`func (o *EndUserDetailsRequest) GetCommonModelsOk() (*[]CommonModelScopesBodyRequest, bool)`

GetCommonModelsOk returns a tuple with the CommonModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonModels

`func (o *EndUserDetailsRequest) SetCommonModels(v []CommonModelScopesBodyRequest)`

SetCommonModels sets CommonModels field to given value.

### HasCommonModels

`func (o *EndUserDetailsRequest) HasCommonModels() bool`

HasCommonModels returns a boolean if a field has been set.

### SetCommonModelsNil

`func (o *EndUserDetailsRequest) SetCommonModelsNil(b bool)`

 SetCommonModelsNil sets the value for CommonModels to be an explicit nil

### UnsetCommonModels
`func (o *EndUserDetailsRequest) UnsetCommonModels()`

UnsetCommonModels ensures that no value is present for CommonModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


