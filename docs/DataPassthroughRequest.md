# DataPassthroughRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**MethodEnum**](MethodEnum.md) |  | 
**Path** | **string** |  | 
**BaseUrlOverride** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**MultipartFormData** | Pointer to [**[]MultipartFormFieldRequest**](MultipartFormFieldRequest.md) | Pass an array of &#x60;MultipartFormField&#x60; objects in here instead of using the &#x60;data&#x60; param if &#x60;request_format&#x60; is set to &#x60;MULTIPART&#x60;. | [optional] 
**Headers** | Pointer to **map[string]interface{}** | The headers to use for the request (Merge will handle the account&#39;s authorization headers). &#x60;Content-Type&#x60; header is required for passthrough. Choose content type corresponding to expected format of receiving server. | [optional] 
**RequestFormat** | Pointer to [**NullableRequestFormatEnum**](RequestFormatEnum.md) |  | [optional] 
**NormalizeResponse** | Pointer to **bool** |  | [optional] 

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

`func (o *DataPassthroughRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataPassthroughRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataPassthroughRequest) SetData(v string)`

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
### GetMultipartFormData

`func (o *DataPassthroughRequest) GetMultipartFormData() []MultipartFormFieldRequest`

GetMultipartFormData returns the MultipartFormData field if non-nil, zero value otherwise.

### GetMultipartFormDataOk

`func (o *DataPassthroughRequest) GetMultipartFormDataOk() (*[]MultipartFormFieldRequest, bool)`

GetMultipartFormDataOk returns a tuple with the MultipartFormData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipartFormData

`func (o *DataPassthroughRequest) SetMultipartFormData(v []MultipartFormFieldRequest)`

SetMultipartFormData sets MultipartFormData field to given value.

### HasMultipartFormData

`func (o *DataPassthroughRequest) HasMultipartFormData() bool`

HasMultipartFormData returns a boolean if a field has been set.

### SetMultipartFormDataNil

`func (o *DataPassthroughRequest) SetMultipartFormDataNil(b bool)`

 SetMultipartFormDataNil sets the value for MultipartFormData to be an explicit nil

### UnsetMultipartFormData
`func (o *DataPassthroughRequest) UnsetMultipartFormData()`

UnsetMultipartFormData ensures that no value is present for MultipartFormData, not even an explicit nil
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
### GetRequestFormat

`func (o *DataPassthroughRequest) GetRequestFormat() RequestFormatEnum`

GetRequestFormat returns the RequestFormat field if non-nil, zero value otherwise.

### GetRequestFormatOk

`func (o *DataPassthroughRequest) GetRequestFormatOk() (*RequestFormatEnum, bool)`

GetRequestFormatOk returns a tuple with the RequestFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFormat

`func (o *DataPassthroughRequest) SetRequestFormat(v RequestFormatEnum)`

SetRequestFormat sets RequestFormat field to given value.

### HasRequestFormat

`func (o *DataPassthroughRequest) HasRequestFormat() bool`

HasRequestFormat returns a boolean if a field has been set.

### SetRequestFormatNil

`func (o *DataPassthroughRequest) SetRequestFormatNil(b bool)`

 SetRequestFormatNil sets the value for RequestFormat to be an explicit nil

### UnsetRequestFormat
`func (o *DataPassthroughRequest) UnsetRequestFormat()`

UnsetRequestFormat ensures that no value is present for RequestFormat, not even an explicit nil
### GetNormalizeResponse

`func (o *DataPassthroughRequest) GetNormalizeResponse() bool`

GetNormalizeResponse returns the NormalizeResponse field if non-nil, zero value otherwise.

### GetNormalizeResponseOk

`func (o *DataPassthroughRequest) GetNormalizeResponseOk() (*bool, bool)`

GetNormalizeResponseOk returns a tuple with the NormalizeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeResponse

`func (o *DataPassthroughRequest) SetNormalizeResponse(v bool)`

SetNormalizeResponse sets NormalizeResponse field to given value.

### HasNormalizeResponse

`func (o *DataPassthroughRequest) HasNormalizeResponse() bool`

HasNormalizeResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


