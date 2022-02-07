# AccountDetailsAndActionsIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Categories** | [**[]CategoriesEnum**](CategoriesEnum.md) |  | 
**Image** | Pointer to **string** |  | [optional] 
**SquareImage** | Pointer to **string** |  | [optional] 
**Color** | **string** |  | 
**Slug** | **string** |  | 
**PassthroughAvailable** | **bool** |  | 
**AvailableModelOperations** | Pointer to [**[]ModelOperation**](ModelOperation.md) |  | [optional] 

## Methods

### NewAccountDetailsAndActionsIntegration

`func NewAccountDetailsAndActionsIntegration(name string, categories []CategoriesEnum, color string, slug string, passthroughAvailable bool, ) *AccountDetailsAndActionsIntegration`

NewAccountDetailsAndActionsIntegration instantiates a new AccountDetailsAndActionsIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsAndActionsIntegrationWithDefaults

`func NewAccountDetailsAndActionsIntegrationWithDefaults() *AccountDetailsAndActionsIntegration`

NewAccountDetailsAndActionsIntegrationWithDefaults instantiates a new AccountDetailsAndActionsIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountDetailsAndActionsIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDetailsAndActionsIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDetailsAndActionsIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetCategories

`func (o *AccountDetailsAndActionsIntegration) GetCategories() []CategoriesEnum`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AccountDetailsAndActionsIntegration) GetCategoriesOk() (*[]CategoriesEnum, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AccountDetailsAndActionsIntegration) SetCategories(v []CategoriesEnum)`

SetCategories sets Categories field to given value.


### GetImage

`func (o *AccountDetailsAndActionsIntegration) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AccountDetailsAndActionsIntegration) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AccountDetailsAndActionsIntegration) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AccountDetailsAndActionsIntegration) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSquareImage

`func (o *AccountDetailsAndActionsIntegration) GetSquareImage() string`

GetSquareImage returns the SquareImage field if non-nil, zero value otherwise.

### GetSquareImageOk

`func (o *AccountDetailsAndActionsIntegration) GetSquareImageOk() (*string, bool)`

GetSquareImageOk returns a tuple with the SquareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareImage

`func (o *AccountDetailsAndActionsIntegration) SetSquareImage(v string)`

SetSquareImage sets SquareImage field to given value.

### HasSquareImage

`func (o *AccountDetailsAndActionsIntegration) HasSquareImage() bool`

HasSquareImage returns a boolean if a field has been set.

### GetColor

`func (o *AccountDetailsAndActionsIntegration) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AccountDetailsAndActionsIntegration) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AccountDetailsAndActionsIntegration) SetColor(v string)`

SetColor sets Color field to given value.


### GetSlug

`func (o *AccountDetailsAndActionsIntegration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountDetailsAndActionsIntegration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountDetailsAndActionsIntegration) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPassthroughAvailable

`func (o *AccountDetailsAndActionsIntegration) GetPassthroughAvailable() bool`

GetPassthroughAvailable returns the PassthroughAvailable field if non-nil, zero value otherwise.

### GetPassthroughAvailableOk

`func (o *AccountDetailsAndActionsIntegration) GetPassthroughAvailableOk() (*bool, bool)`

GetPassthroughAvailableOk returns a tuple with the PassthroughAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughAvailable

`func (o *AccountDetailsAndActionsIntegration) SetPassthroughAvailable(v bool)`

SetPassthroughAvailable sets PassthroughAvailable field to given value.


### GetAvailableModelOperations

`func (o *AccountDetailsAndActionsIntegration) GetAvailableModelOperations() []ModelOperation`

GetAvailableModelOperations returns the AvailableModelOperations field if non-nil, zero value otherwise.

### GetAvailableModelOperationsOk

`func (o *AccountDetailsAndActionsIntegration) GetAvailableModelOperationsOk() (*[]ModelOperation, bool)`

GetAvailableModelOperationsOk returns a tuple with the AvailableModelOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableModelOperations

`func (o *AccountDetailsAndActionsIntegration) SetAvailableModelOperations(v []ModelOperation)`

SetAvailableModelOperations sets AvailableModelOperations field to given value.

### HasAvailableModelOperations

`func (o *AccountDetailsAndActionsIntegration) HasAvailableModelOperations() bool`

HasAvailableModelOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


