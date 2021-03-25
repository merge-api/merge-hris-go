# EndUserDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserEmailAddress** | **string** |  | 
**EndUserOrganizationName** | **string** |  | 
**EndUserOriginId** | **string** |  | 
**Categories** | **[]string** |  | 

## Methods

### NewEndUserDetailsRequest

`func NewEndUserDetailsRequest(endUserEmailAddress string, endUserOrganizationName string, endUserOriginId string, categories []string, ) *EndUserDetailsRequest`

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

`func (o *EndUserDetailsRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *EndUserDetailsRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *EndUserDetailsRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


