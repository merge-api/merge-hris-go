# AccountIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Company name. | 
**Categories** | Pointer to [**[]CategoriesEnum**](CategoriesEnum.md) | Category or categories this integration belongs to. Multiple categories should be comma separated, i.e. [ats, hris]. | [optional] 
**Image** | Pointer to **NullableString** | Company logo in rectangular shape. &lt;b&gt;Upload an image with a clear background.&lt;/b&gt; | [optional] 
**SquareImage** | Pointer to **NullableString** | Company logo in square shape. &lt;b&gt;Upload an image with a white background.&lt;/b&gt; | [optional] 
**Color** | Pointer to **string** | The color of this integration used for buttons and text throughout the app and landing pages. &lt;b&gt;Choose a darker, saturated color.&lt;/b&gt; | [optional] 
**Slug** | Pointer to **string** |  | [optional] [readonly] 
**IsInBeta** | Pointer to **bool** | If checked, this integration will not appear in the linking flow, and will appear elsewhere with a Beta tag. | [optional] 
**ApiEndpointsToDocumentationUrls** | Pointer to **map[string]interface{}** | Mapping of API endpoints to documentation urls for support. Example: {&#39;GET&#39;: [[&#39;/common-model-scopes&#39;, &#39;https://docs.merge.dev/accounting/common-model-scopes/#common_model_scopes_retrieve&#39;],[&#39;/common-model-actions&#39;, &#39;https://docs.merge.dev/accounting/common-model-actions/#common_model_actions_retrieve&#39;]], &#39;POST&#39;: []} | [optional] 

## Methods

### NewAccountIntegration

`func NewAccountIntegration(name string, ) *AccountIntegration`

NewAccountIntegration instantiates a new AccountIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIntegrationWithDefaults

`func NewAccountIntegrationWithDefaults() *AccountIntegration`

NewAccountIntegrationWithDefaults instantiates a new AccountIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetCategories

`func (o *AccountIntegration) GetCategories() []CategoriesEnum`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AccountIntegration) GetCategoriesOk() (*[]CategoriesEnum, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AccountIntegration) SetCategories(v []CategoriesEnum)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AccountIntegration) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetImage

`func (o *AccountIntegration) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AccountIntegration) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AccountIntegration) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AccountIntegration) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *AccountIntegration) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *AccountIntegration) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetSquareImage

`func (o *AccountIntegration) GetSquareImage() string`

GetSquareImage returns the SquareImage field if non-nil, zero value otherwise.

### GetSquareImageOk

`func (o *AccountIntegration) GetSquareImageOk() (*string, bool)`

GetSquareImageOk returns a tuple with the SquareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareImage

`func (o *AccountIntegration) SetSquareImage(v string)`

SetSquareImage sets SquareImage field to given value.

### HasSquareImage

`func (o *AccountIntegration) HasSquareImage() bool`

HasSquareImage returns a boolean if a field has been set.

### SetSquareImageNil

`func (o *AccountIntegration) SetSquareImageNil(b bool)`

 SetSquareImageNil sets the value for SquareImage to be an explicit nil

### UnsetSquareImage
`func (o *AccountIntegration) UnsetSquareImage()`

UnsetSquareImage ensures that no value is present for SquareImage, not even an explicit nil
### GetColor

`func (o *AccountIntegration) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AccountIntegration) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AccountIntegration) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AccountIntegration) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetSlug

`func (o *AccountIntegration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AccountIntegration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AccountIntegration) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AccountIntegration) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetIsInBeta

`func (o *AccountIntegration) GetIsInBeta() bool`

GetIsInBeta returns the IsInBeta field if non-nil, zero value otherwise.

### GetIsInBetaOk

`func (o *AccountIntegration) GetIsInBetaOk() (*bool, bool)`

GetIsInBetaOk returns a tuple with the IsInBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInBeta

`func (o *AccountIntegration) SetIsInBeta(v bool)`

SetIsInBeta sets IsInBeta field to given value.

### HasIsInBeta

`func (o *AccountIntegration) HasIsInBeta() bool`

HasIsInBeta returns a boolean if a field has been set.

### GetApiEndpointsToDocumentationUrls

`func (o *AccountIntegration) GetApiEndpointsToDocumentationUrls() map[string]interface{}`

GetApiEndpointsToDocumentationUrls returns the ApiEndpointsToDocumentationUrls field if non-nil, zero value otherwise.

### GetApiEndpointsToDocumentationUrlsOk

`func (o *AccountIntegration) GetApiEndpointsToDocumentationUrlsOk() (*map[string]interface{}, bool)`

GetApiEndpointsToDocumentationUrlsOk returns a tuple with the ApiEndpointsToDocumentationUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpointsToDocumentationUrls

`func (o *AccountIntegration) SetApiEndpointsToDocumentationUrls(v map[string]interface{})`

SetApiEndpointsToDocumentationUrls sets ApiEndpointsToDocumentationUrls field to given value.

### HasApiEndpointsToDocumentationUrls

`func (o *AccountIntegration) HasApiEndpointsToDocumentationUrls() bool`

HasApiEndpointsToDocumentationUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


