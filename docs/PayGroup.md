# PayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**PayGroupName** | Pointer to **NullableString** | The pay group name. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 

## Methods

### NewPayGroup

`func NewPayGroup() *PayGroup`

NewPayGroup instantiates a new PayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayGroupWithDefaults

`func NewPayGroupWithDefaults() *PayGroup`

NewPayGroupWithDefaults instantiates a new PayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PayGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayGroup) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayGroup) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayGroup) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayGroup) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *PayGroup) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *PayGroup) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetPayGroupName

`func (o *PayGroup) GetPayGroupName() string`

GetPayGroupName returns the PayGroupName field if non-nil, zero value otherwise.

### GetPayGroupNameOk

`func (o *PayGroup) GetPayGroupNameOk() (*string, bool)`

GetPayGroupNameOk returns a tuple with the PayGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayGroupName

`func (o *PayGroup) SetPayGroupName(v string)`

SetPayGroupName sets PayGroupName field to given value.

### HasPayGroupName

`func (o *PayGroup) HasPayGroupName() bool`

HasPayGroupName returns a boolean if a field has been set.

### SetPayGroupNameNil

`func (o *PayGroup) SetPayGroupNameNil(b bool)`

 SetPayGroupNameNil sets the value for PayGroupName to be an explicit nil

### UnsetPayGroupName
`func (o *PayGroup) UnsetPayGroupName()`

UnsetPayGroupName ensures that no value is present for PayGroupName, not even an explicit nil
### GetRemoteData

`func (o *PayGroup) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PayGroup) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PayGroup) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *PayGroup) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *PayGroup) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PayGroup) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetRemoteWasDeleted

`func (o *PayGroup) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *PayGroup) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *PayGroup) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *PayGroup) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


