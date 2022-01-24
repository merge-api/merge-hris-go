# TimeOffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Employee** | Pointer to **NullableString** | The employee requesting time off. | [optional] 
**Approver** | Pointer to **NullableString** | The employee approving the time off request. | [optional] 
**Status** | **string** |  | 
**EmployeeNote** | Pointer to **NullableString** | The employee note for this time off request. | [optional] 
**Units** | **string** |  | 
**Amount** | Pointer to **NullableFloat32** | The number of time off units requested. | [optional] 
**RequestType** | **string** |  | 
**StartTime** | Pointer to **NullableTime** | The day and time of the start of the time requested off. | [optional] 
**EndTime** | Pointer to **NullableTime** | The day and time of the end of the time requested off. | [optional] 

## Methods

### NewTimeOffRequest

`func NewTimeOffRequest(status string, units string, requestType string, ) *TimeOffRequest`

NewTimeOffRequest instantiates a new TimeOffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffRequestWithDefaults

`func NewTimeOffRequestWithDefaults() *TimeOffRequest`

NewTimeOffRequestWithDefaults instantiates a new TimeOffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *TimeOffRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TimeOffRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TimeOffRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *TimeOffRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *TimeOffRequest) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *TimeOffRequest) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployee

`func (o *TimeOffRequest) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *TimeOffRequest) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *TimeOffRequest) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *TimeOffRequest) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *TimeOffRequest) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *TimeOffRequest) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetApprover

`func (o *TimeOffRequest) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *TimeOffRequest) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *TimeOffRequest) SetApprover(v string)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *TimeOffRequest) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### SetApproverNil

`func (o *TimeOffRequest) SetApproverNil(b bool)`

 SetApproverNil sets the value for Approver to be an explicit nil

### UnsetApprover
`func (o *TimeOffRequest) UnsetApprover()`

UnsetApprover ensures that no value is present for Approver, not even an explicit nil
### GetStatus

`func (o *TimeOffRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeOffRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeOffRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEmployeeNote

`func (o *TimeOffRequest) GetEmployeeNote() string`

GetEmployeeNote returns the EmployeeNote field if non-nil, zero value otherwise.

### GetEmployeeNoteOk

`func (o *TimeOffRequest) GetEmployeeNoteOk() (*string, bool)`

GetEmployeeNoteOk returns a tuple with the EmployeeNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNote

`func (o *TimeOffRequest) SetEmployeeNote(v string)`

SetEmployeeNote sets EmployeeNote field to given value.

### HasEmployeeNote

`func (o *TimeOffRequest) HasEmployeeNote() bool`

HasEmployeeNote returns a boolean if a field has been set.

### SetEmployeeNoteNil

`func (o *TimeOffRequest) SetEmployeeNoteNil(b bool)`

 SetEmployeeNoteNil sets the value for EmployeeNote to be an explicit nil

### UnsetEmployeeNote
`func (o *TimeOffRequest) UnsetEmployeeNote()`

UnsetEmployeeNote ensures that no value is present for EmployeeNote, not even an explicit nil
### GetUnits

`func (o *TimeOffRequest) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TimeOffRequest) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TimeOffRequest) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetAmount

`func (o *TimeOffRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TimeOffRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TimeOffRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TimeOffRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *TimeOffRequest) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *TimeOffRequest) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetRequestType

`func (o *TimeOffRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TimeOffRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TimeOffRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetStartTime

`func (o *TimeOffRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeOffRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeOffRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TimeOffRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TimeOffRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TimeOffRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *TimeOffRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimeOffRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimeOffRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimeOffRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TimeOffRequest) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TimeOffRequest) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


