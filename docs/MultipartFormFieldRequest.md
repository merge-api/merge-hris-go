# MultipartFormFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the form field | 
**Data** | **string** | The data for the form field. | 
**Encoding** | Pointer to [**NullableEncodingEnum**](EncodingEnum.md) | The encoding of the value of &#x60;data&#x60;. Defaults to &#x60;RAW&#x60; if not defined.  * &#x60;RAW&#x60; - RAW * &#x60;BASE64&#x60; - BASE64 * &#x60;GZIP_BASE64&#x60; - GZIP_BASE64 | [optional] 
**FileName** | Pointer to **NullableString** | The file name of the form field, if the field is for a file. | [optional] 
**ContentType** | Pointer to **NullableString** | The MIME type of the file, if the field is for a file. | [optional] 

## Methods

### NewMultipartFormFieldRequest

`func NewMultipartFormFieldRequest(name string, data string, ) *MultipartFormFieldRequest`

NewMultipartFormFieldRequest instantiates a new MultipartFormFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipartFormFieldRequestWithDefaults

`func NewMultipartFormFieldRequestWithDefaults() *MultipartFormFieldRequest`

NewMultipartFormFieldRequestWithDefaults instantiates a new MultipartFormFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MultipartFormFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultipartFormFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultipartFormFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *MultipartFormFieldRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultipartFormFieldRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultipartFormFieldRequest) SetData(v string)`

SetData sets Data field to given value.


### GetEncoding

`func (o *MultipartFormFieldRequest) GetEncoding() EncodingEnum`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *MultipartFormFieldRequest) GetEncodingOk() (*EncodingEnum, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *MultipartFormFieldRequest) SetEncoding(v EncodingEnum)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *MultipartFormFieldRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### SetEncodingNil

`func (o *MultipartFormFieldRequest) SetEncodingNil(b bool)`

 SetEncodingNil sets the value for Encoding to be an explicit nil

### UnsetEncoding
`func (o *MultipartFormFieldRequest) UnsetEncoding()`

UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
### GetFileName

`func (o *MultipartFormFieldRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MultipartFormFieldRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MultipartFormFieldRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MultipartFormFieldRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MultipartFormFieldRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MultipartFormFieldRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetContentType

`func (o *MultipartFormFieldRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MultipartFormFieldRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MultipartFormFieldRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MultipartFormFieldRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MultipartFormFieldRequest) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MultipartFormFieldRequest) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


