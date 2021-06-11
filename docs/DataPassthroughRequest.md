# DataPassthroughRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**MethodEnum**](MethodEnum.md) |  | 
**Path** | **string** |  | 
**BaseUrlOverride** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDataPassthroughRequest

`func NewDataPassthroughRequest(method MethodEnum, path string, ) *DataPassthroughRequest`

NewDataPassthroughRequest instantiates a new DataPassthroughRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPassthroughRequestWithDefaults

`func NewDataPassthroughRequestWithDefaults() *DataPassthroughRequest`

NewDataPassthroughRequestWithDefaults instantiates a new DataPassthroughRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *DataPassthroughRequest) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DataPassthroughRequest) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DataPassthroughRequest) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.


### GetPath

`func (o *DataPassthroughRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DataPassthroughRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DataPassthroughRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetBaseUrlOverride

`func (o *DataPassthroughRequest) GetBaseUrlOverride() string`

GetBaseUrlOverride returns the BaseUrlOverride field if non-nil, zero value otherwise.

### GetBaseUrlOverrideOk

`func (o *DataPassthroughRequest) GetBaseUrlOverrideOk() (*string, bool)`

GetBaseUrlOverrideOk returns a tuple with the BaseUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrlOverride

`func (o *DataPassthroughRequest) SetBaseUrlOverride(v string)`

SetBaseUrlOverride sets BaseUrlOverride field to given value.

### HasBaseUrlOverride

`func (o *DataPassthroughRequest) HasBaseUrlOverride() bool`

HasBaseUrlOverride returns a boolean if a field has been set.

### SetBaseUrlOverrideNil

`func (o *DataPassthroughRequest) SetBaseUrlOverrideNil(b bool)`

 SetBaseUrlOverrideNil sets the value for BaseUrlOverride to be an explicit nil

### UnsetBaseUrlOverride
`func (o *DataPassthroughRequest) UnsetBaseUrlOverride()`

UnsetBaseUrlOverride ensures that no value is present for BaseUrlOverride, not even an explicit nil
### GetData

`func (o *DataPassthroughRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataPassthroughRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataPassthroughRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DataPassthroughRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *DataPassthroughRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DataPassthroughRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetHeaders

`func (o *DataPassthroughRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DataPassthroughRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DataPassthroughRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DataPassthroughRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *DataPassthroughRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *DataPassthroughRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


