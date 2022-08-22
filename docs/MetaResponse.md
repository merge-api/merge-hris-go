# MetaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSchema** | **map[string]interface{}** |  | 
**Status** | Pointer to [**LinkedAccountStatus**](LinkedAccountStatus.md) |  | [optional] 
**HasConditionalParams** | **bool** |  | 
**HasRequiredLinkedAccountParams** | **bool** |  | 

## Methods

### NewMetaResponse

`func NewMetaResponse(requestSchema map[string]interface{}, hasConditionalParams bool, hasRequiredLinkedAccountParams bool, ) *MetaResponse`

NewMetaResponse instantiates a new MetaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaResponseWithDefaults

`func NewMetaResponseWithDefaults() *MetaResponse`

NewMetaResponseWithDefaults instantiates a new MetaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSchema

`func (o *MetaResponse) GetRequestSchema() map[string]interface{}`

GetRequestSchema returns the RequestSchema field if non-nil, zero value otherwise.

### GetRequestSchemaOk

`func (o *MetaResponse) GetRequestSchemaOk() (*map[string]interface{}, bool)`

GetRequestSchemaOk returns a tuple with the RequestSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSchema

`func (o *MetaResponse) SetRequestSchema(v map[string]interface{})`

SetRequestSchema sets RequestSchema field to given value.


### GetStatus

`func (o *MetaResponse) GetStatus() LinkedAccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetaResponse) GetStatusOk() (*LinkedAccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetaResponse) SetStatus(v LinkedAccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetaResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHasConditionalParams

`func (o *MetaResponse) GetHasConditionalParams() bool`

GetHasConditionalParams returns the HasConditionalParams field if non-nil, zero value otherwise.

### GetHasConditionalParamsOk

`func (o *MetaResponse) GetHasConditionalParamsOk() (*bool, bool)`

GetHasConditionalParamsOk returns a tuple with the HasConditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConditionalParams

`func (o *MetaResponse) SetHasConditionalParams(v bool)`

SetHasConditionalParams sets HasConditionalParams field to given value.


### GetHasRequiredLinkedAccountParams

`func (o *MetaResponse) GetHasRequiredLinkedAccountParams() bool`

GetHasRequiredLinkedAccountParams returns the HasRequiredLinkedAccountParams field if non-nil, zero value otherwise.

### GetHasRequiredLinkedAccountParamsOk

`func (o *MetaResponse) GetHasRequiredLinkedAccountParamsOk() (*bool, bool)`

GetHasRequiredLinkedAccountParamsOk returns a tuple with the HasRequiredLinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRequiredLinkedAccountParams

`func (o *MetaResponse) SetHasRequiredLinkedAccountParams(v bool)`

SetHasRequiredLinkedAccountParams sets HasRequiredLinkedAccountParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


