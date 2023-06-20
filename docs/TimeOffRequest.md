# TimeOffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employee** | Pointer to **NullableString** | The employee requesting time off. | [optional] 
**Approver** | Pointer to **NullableString** | The Merge ID of the employee with the ability to approve the time off request. | [optional] 
**Status** | Pointer to [**NullableTimeOffStatusEnum**](TimeOffStatusEnum.md) | The status of this time off request.  * &#x60;REQUESTED&#x60; - REQUESTED * &#x60;APPROVED&#x60; - APPROVED * &#x60;DECLINED&#x60; - DECLINED * &#x60;CANCELLED&#x60; - CANCELLED * &#x60;DELETED&#x60; - DELETED | [optional] 
**EmployeeNote** | Pointer to **NullableString** | The employee note for this time off request. | [optional] 
**Units** | Pointer to [**NullableUnitsEnum**](UnitsEnum.md) | The measurement that the third-party integration uses to count time requested.  * &#x60;HOURS&#x60; - HOURS * &#x60;DAYS&#x60; - DAYS | [optional] 
**Amount** | Pointer to **NullableFloat64** | The time off quantity measured by the prescribed “units”. | [optional] 
**RequestType** | Pointer to [**NullableRequestTypeEnum**](RequestTypeEnum.md) | The type of time off request.  * &#x60;VACATION&#x60; - VACATION * &#x60;SICK&#x60; - SICK * &#x60;PERSONAL&#x60; - PERSONAL * &#x60;JURY_DUTY&#x60; - JURY_DUTY * &#x60;VOLUNTEER&#x60; - VOLUNTEER * &#x60;BEREAVEMENT&#x60; - BEREAVEMENT | [optional] 
**StartTime** | Pointer to **NullableTime** | The day and time of the start of the time requested off. | [optional] 
**EndTime** | Pointer to **NullableTime** | The day and time of the end of the time requested off. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTimeOffRequest

`func NewTimeOffRequest() *TimeOffRequest`

NewTimeOffRequest instantiates a new TimeOffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffRequestWithDefaults

`func NewTimeOffRequestWithDefaults() *TimeOffRequest`

NewTimeOffRequestWithDefaults instantiates a new TimeOffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *TimeOffRequest) GetStatus() TimeOffStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeOffRequest) GetStatusOk() (*TimeOffStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeOffRequest) SetStatus(v TimeOffStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TimeOffRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TimeOffRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TimeOffRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
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

`func (o *TimeOffRequest) GetUnits() UnitsEnum`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TimeOffRequest) GetUnitsOk() (*UnitsEnum, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TimeOffRequest) SetUnits(v UnitsEnum)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *TimeOffRequest) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *TimeOffRequest) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *TimeOffRequest) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetAmount

`func (o *TimeOffRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TimeOffRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TimeOffRequest) SetAmount(v float64)`

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

`func (o *TimeOffRequest) GetRequestType() RequestTypeEnum`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TimeOffRequest) GetRequestTypeOk() (*RequestTypeEnum, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TimeOffRequest) SetRequestType(v RequestTypeEnum)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TimeOffRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *TimeOffRequest) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *TimeOffRequest) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
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
### GetIntegrationParams

`func (o *TimeOffRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *TimeOffRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *TimeOffRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *TimeOffRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *TimeOffRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *TimeOffRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *TimeOffRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *TimeOffRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *TimeOffRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *TimeOffRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *TimeOffRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *TimeOffRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


