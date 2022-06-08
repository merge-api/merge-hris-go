# TimeOff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Employee** | Pointer to **NullableString** |  | [optional] 
**Approver** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableTimeOffStatusEnum**](TimeOffStatusEnum.md) | The status of this time off request. | [optional] 
**EmployeeNote** | Pointer to **NullableString** | The employee note for this time off request. | [optional] 
**Units** | Pointer to [**NullableUnitsEnum**](UnitsEnum.md) | The unit of time requested. | [optional] 
**Amount** | Pointer to **NullableFloat32** | The number of time off units requested. | [optional] 
**RequestType** | Pointer to [**NullableRequestTypeEnum**](RequestTypeEnum.md) | The type of time off request. | [optional] 
**StartTime** | Pointer to **NullableTime** | The day and time of the start of the time requested off. | [optional] 
**EndTime** | Pointer to **NullableTime** | The day and time of the end of the time requested off. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewTimeOff

`func NewTimeOff() *TimeOff`

NewTimeOff instantiates a new TimeOff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffWithDefaults

`func NewTimeOffWithDefaults() *TimeOff`

NewTimeOffWithDefaults instantiates a new TimeOff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeOff) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeOff) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeOff) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimeOff) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *TimeOff) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TimeOff) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TimeOff) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *TimeOff) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *TimeOff) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *TimeOff) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployee

`func (o *TimeOff) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *TimeOff) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *TimeOff) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *TimeOff) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *TimeOff) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *TimeOff) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetApprover

`func (o *TimeOff) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *TimeOff) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *TimeOff) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *TimeOff) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *TimeOff) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *TimeOff) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetStatus

`func (o *TimeOff) GetStatus() TimeOffStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeOff) GetStatusOk() (*TimeOffStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeOff) SetStatus(v TimeOffStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TimeOff) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TimeOff) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TimeOff) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEmployeeNote

`func (o *TimeOff) GetEmployeeNote() string`

GetEmployeeNote returns the EmployeeNote field if non-nil, zero value otherwise.

### GetEmployeeNoteOk

`func (o *TimeOff) GetEmployeeNoteOk() (*string, bool)`

GetEmployeeNoteOk returns a tuple with the EmployeeNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNote

`func (o *TimeOff) SetEmployeeNote(v string)`

SetEmployeeNote sets EmployeeNote field to given value.

### HasEmployeeNote

`func (o *TimeOff) HasEmployeeNote() bool`

HasEmployeeNote returns a boolean if a field has been set.

### SetEmployeeNoteNil

`func (o *TimeOff) SetEmployeeNoteNil(b bool)`

 SetEmployeeNoteNil sets the value for EmployeeNote to be an explicit nil

### UnsetEmployeeNote
`func (o *TimeOff) UnsetEmployeeNote()`

UnsetEmployeeNote ensures that no value is present for EmployeeNote, not even an explicit nil
### GetUnits

`func (o *TimeOff) GetUnits() UnitsEnum`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TimeOff) GetUnitsOk() (*UnitsEnum, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TimeOff) SetUnits(v UnitsEnum)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *TimeOff) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *TimeOff) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *TimeOff) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetAmount

`func (o *TimeOff) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TimeOff) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TimeOff) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TimeOff) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *TimeOff) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *TimeOff) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetRequestType

`func (o *TimeOff) GetRequestType() RequestTypeEnum`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TimeOff) GetRequestTypeOk() (*RequestTypeEnum, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TimeOff) SetRequestType(v RequestTypeEnum)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TimeOff) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *TimeOff) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *TimeOff) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetStartTime

`func (o *TimeOff) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeOff) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeOff) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TimeOff) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TimeOff) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TimeOff) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *TimeOff) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimeOff) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimeOff) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimeOff) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TimeOff) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TimeOff) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetRemoteWasDeleted

`func (o *TimeOff) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *TimeOff) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *TimeOff) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *TimeOff) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


