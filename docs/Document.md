# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Employee** | Pointer to **NullableString** | The employee this document belongs to. | [optional] 
**FileName** | Pointer to **NullableString** | The document&#39;s filename. | [optional] 
**FileUrl** | Pointer to **NullableString** | The URL from which the document can be retrieved. | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Document) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmployee

`func (o *Document) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *Document) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *Document) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *Document) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *Document) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *Document) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetFileName

`func (o *Document) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Document) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Document) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Document) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *Document) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *Document) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileUrl

`func (o *Document) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *Document) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *Document) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *Document) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *Document) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *Document) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


