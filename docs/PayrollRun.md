# PayrollRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RunState** | Pointer to [**NullableRunStateEnum**](RunStateEnum.md) | The state of the payroll run  * &#x60;PAID&#x60; - PAID * &#x60;DRAFT&#x60; - DRAFT * &#x60;APPROVED&#x60; - APPROVED * &#x60;FAILED&#x60; - FAILED * &#x60;CLOSED&#x60; - CLOSED | [optional] 
**RunType** | Pointer to [**NullableRunTypeEnum**](RunTypeEnum.md) | The type of the payroll run  * &#x60;REGULAR&#x60; - REGULAR * &#x60;OFF_CYCLE&#x60; - OFF_CYCLE * &#x60;CORRECTION&#x60; - CORRECTION * &#x60;TERMINATION&#x60; - TERMINATION * &#x60;SIGN_ON_BONUS&#x60; - SIGN_ON_BONUS | [optional] 
**StartDate** | Pointer to **NullableTime** | The day and time the payroll run started. | [optional] 
**EndDate** | Pointer to **NullableTime** | The day and time the payroll run ended. | [optional] 
**CheckDate** | Pointer to **NullableTime** | The day and time the payroll run was checked. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewPayrollRun

`func NewPayrollRun() *PayrollRun`

NewPayrollRun instantiates a new PayrollRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayrollRunWithDefaults

`func NewPayrollRunWithDefaults() *PayrollRun`

NewPayrollRunWithDefaults instantiates a new PayrollRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PayrollRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayrollRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayrollRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayrollRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayrollRun) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayrollRun) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayrollRun) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayrollRun) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *PayrollRun) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *PayrollRun) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRunState

`func (o *PayrollRun) GetRunState() RunStateEnum`

GetRunState returns the RunState field if non-nil, zero value otherwise.

### GetRunStateOk

`func (o *PayrollRun) GetRunStateOk() (*RunStateEnum, bool)`

GetRunStateOk returns a tuple with the RunState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunState

`func (o *PayrollRun) SetRunState(v RunStateEnum)`

SetRunState sets RunState field to given value.

### HasRunState

`func (o *PayrollRun) HasRunState() bool`

HasRunState returns a boolean if a field has been set.

### SetRunStateNil

`func (o *PayrollRun) SetRunStateNil(b bool)`

 SetRunStateNil sets the value for RunState to be an explicit nil

### UnsetRunState
`func (o *PayrollRun) UnsetRunState()`

UnsetRunState ensures that no value is present for RunState, not even an explicit nil
### GetRunType

`func (o *PayrollRun) GetRunType() RunTypeEnum`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *PayrollRun) GetRunTypeOk() (*RunTypeEnum, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *PayrollRun) SetRunType(v RunTypeEnum)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *PayrollRun) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *PayrollRun) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *PayrollRun) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetStartDate

`func (o *PayrollRun) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PayrollRun) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PayrollRun) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PayrollRun) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *PayrollRun) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PayrollRun) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *PayrollRun) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PayrollRun) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PayrollRun) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PayrollRun) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *PayrollRun) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PayrollRun) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetCheckDate

`func (o *PayrollRun) GetCheckDate() time.Time`

GetCheckDate returns the CheckDate field if non-nil, zero value otherwise.

### GetCheckDateOk

`func (o *PayrollRun) GetCheckDateOk() (*time.Time, bool)`

GetCheckDateOk returns a tuple with the CheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckDate

`func (o *PayrollRun) SetCheckDate(v time.Time)`

SetCheckDate sets CheckDate field to given value.

### HasCheckDate

`func (o *PayrollRun) HasCheckDate() bool`

HasCheckDate returns a boolean if a field has been set.

### SetCheckDateNil

`func (o *PayrollRun) SetCheckDateNil(b bool)`

 SetCheckDateNil sets the value for CheckDate to be an explicit nil

### UnsetCheckDate
`func (o *PayrollRun) UnsetCheckDate()`

UnsetCheckDate ensures that no value is present for CheckDate, not even an explicit nil
### GetRemoteWasDeleted

`func (o *PayrollRun) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *PayrollRun) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *PayrollRun) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *PayrollRun) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *PayrollRun) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *PayrollRun) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *PayrollRun) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *PayrollRun) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *PayrollRun) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *PayrollRun) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *PayrollRun) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PayrollRun) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PayrollRun) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PayrollRun) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *PayrollRun) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PayrollRun) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PayrollRun) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *PayrollRun) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *PayrollRun) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PayrollRun) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


