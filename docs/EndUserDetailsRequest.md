# EndUserDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserEmailAddress** | **string** |  | 
**EndUserOrganizationName** | **string** |  | 
**EndUserOriginId** | **string** |  | 
**Categories** | Pointer to [**[]CategoriesEnum**](CategoriesEnum.md) |  | [optional] 
**Integration** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEndUserDetailsRequest

`func NewEndUserDetailsRequest(endUserEmailAddress string, endUserOrganizationName string, endUserOriginId string, ) *EndUserDetailsRequest`

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

### HasCategories

`func (o *EndUserDetailsRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


