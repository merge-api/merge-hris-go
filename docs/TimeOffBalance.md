# TimeOffBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Employee** | Pointer to **NullableString** |  | [optional] 
**Balance** | Pointer to **NullableFloat32** | The current remaining PTO balance in terms of hours. This does not represent the starting PTO balance. If the API provider only provides PTO balance in terms of days, we estimate 8 hours per day. | [optional] 
**Used** | Pointer to **NullableFloat32** | The amount of PTO used in terms of hours. | [optional] 
**PolicyType** | Pointer to [**NullablePolicyTypeEnum**](PolicyTypeEnum.md) | The policy type of this time off balance. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 

## Methods

### NewTimeOffBalance

`func NewTimeOffBalance() *TimeOffBalance`

NewTimeOffBalance instantiates a new TimeOffBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffBalanceWithDefaults

`func NewTimeOffBalanceWithDefaults() *TimeOffBalance`

NewTimeOffBalanceWithDefaults instantiates a new TimeOffBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeOffBalance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeOffBalance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeOffBalance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimeOffBalance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *TimeOffBalance) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TimeOffBalance) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TimeOffBalance) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *TimeOffBalance) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *TimeOffBalance) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *TimeOffBalance) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployee

`func (o *TimeOffBalance) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *TimeOffBalance) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *TimeOffBalance) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *TimeOffBalance) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *TimeOffBalance) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *TimeOffBalance) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetBalance

`func (o *TimeOffBalance) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TimeOffBalance) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TimeOffBalance) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *TimeOffBalance) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *TimeOffBalance) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *TimeOffBalance) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetUsed

`func (o *TimeOffBalance) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *TimeOffBalance) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *TimeOffBalance) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *TimeOffBalance) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### SetUsedNil

`func (o *TimeOffBalance) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *TimeOffBalance) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil
### GetPolicyType

`func (o *TimeOffBalance) GetPolicyType() PolicyTypeEnum`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *TimeOffBalance) GetPolicyTypeOk() (*PolicyTypeEnum, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *TimeOffBalance) SetPolicyType(v PolicyTypeEnum)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *TimeOffBalance) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### SetPolicyTypeNil

`func (o *TimeOffBalance) SetPolicyTypeNil(b bool)`

 SetPolicyTypeNil sets the value for PolicyType to be an explicit nil

### UnsetPolicyType
`func (o *TimeOffBalance) UnsetPolicyType()`

UnsetPolicyType ensures that no value is present for PolicyType, not even an explicit nil
### GetRemoteData

`func (o *TimeOffBalance) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *TimeOffBalance) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *TimeOffBalance) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *TimeOffBalance) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *TimeOffBalance) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *TimeOffBalance) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetRemoteWasDeleted

`func (o *TimeOffBalance) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *TimeOffBalance) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *TimeOffBalance) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *TimeOffBalance) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


