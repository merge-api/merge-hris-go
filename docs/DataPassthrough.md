# DataPassthrough

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**MethodEnum**](MethodEnum.md) |  | 
**Path** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDataPassthrough

`func NewDataPassthrough(method MethodEnum, path string, ) *DataPassthrough`

NewDataPassthrough instantiates a new DataPassthrough object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPassthroughWithDefaults

`func NewDataPassthroughWithDefaults() *DataPassthrough`

NewDataPassthroughWithDefaults instantiates a new DataPassthrough object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *DataPassthrough) GetMethod() MethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DataPassthrough) GetMethodOk() (*MethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DataPassthrough) SetMethod(v MethodEnum)`

SetMethod sets Method field to given value.


### GetPath

`func (o *DataPassthrough) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DataPassthrough) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DataPassthrough) SetPath(v string)`

SetPath sets Path field to given value.


### GetData

`func (o *DataPassthrough) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataPassthrough) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataPassthrough) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DataPassthrough) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *DataPassthrough) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DataPassthrough) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DataPassthrough) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DataPassthrough) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


