# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**IssueStatusEnum**](IssueStatusEnum.md) | Status of the issue. Options: (&#39;ONGOING&#39;, &#39;RESOLVED&#39;)  * &#x60;ONGOING&#x60; - ONGOING * &#x60;RESOLVED&#x60; - RESOLVED | [optional] 
**ErrorDescription** | **string** |  | 
**EndUser** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**FirstIncidentTime** | Pointer to **NullableTime** |  | [optional] 
**LastIncidentTime** | Pointer to **NullableTime** |  | [optional] 
**IsMuted** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorDetails** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewIssue

`func NewIssue(errorDescription string, ) *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Issue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Issue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Issue) GetStatus() IssueStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Issue) GetStatusOk() (*IssueStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Issue) SetStatus(v IssueStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Issue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorDescription

`func (o *Issue) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *Issue) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *Issue) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.


### GetEndUser

`func (o *Issue) GetEndUser() map[string]interface{}`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *Issue) GetEndUserOk() (*map[string]interface{}, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *Issue) SetEndUser(v map[string]interface{})`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *Issue) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetFirstIncidentTime

`func (o *Issue) GetFirstIncidentTime() time.Time`

GetFirstIncidentTime returns the FirstIncidentTime field if non-nil, zero value otherwise.

### GetFirstIncidentTimeOk

`func (o *Issue) GetFirstIncidentTimeOk() (*time.Time, bool)`

GetFirstIncidentTimeOk returns a tuple with the FirstIncidentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstIncidentTime

`func (o *Issue) SetFirstIncidentTime(v time.Time)`

SetFirstIncidentTime sets FirstIncidentTime field to given value.

### HasFirstIncidentTime

`func (o *Issue) HasFirstIncidentTime() bool`

HasFirstIncidentTime returns a boolean if a field has been set.

### SetFirstIncidentTimeNil

`func (o *Issue) SetFirstIncidentTimeNil(b bool)`

 SetFirstIncidentTimeNil sets the value for FirstIncidentTime to be an explicit nil

### UnsetFirstIncidentTime
`func (o *Issue) UnsetFirstIncidentTime()`

UnsetFirstIncidentTime ensures that no value is present for FirstIncidentTime, not even an explicit nil
### GetLastIncidentTime

`func (o *Issue) GetLastIncidentTime() time.Time`

GetLastIncidentTime returns the LastIncidentTime field if non-nil, zero value otherwise.

### GetLastIncidentTimeOk

`func (o *Issue) GetLastIncidentTimeOk() (*time.Time, bool)`

GetLastIncidentTimeOk returns a tuple with the LastIncidentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIncidentTime

`func (o *Issue) SetLastIncidentTime(v time.Time)`

SetLastIncidentTime sets LastIncidentTime field to given value.

### HasLastIncidentTime

`func (o *Issue) HasLastIncidentTime() bool`

HasLastIncidentTime returns a boolean if a field has been set.

### SetLastIncidentTimeNil

`func (o *Issue) SetLastIncidentTimeNil(b bool)`

 SetLastIncidentTimeNil sets the value for LastIncidentTime to be an explicit nil

### UnsetLastIncidentTime
`func (o *Issue) UnsetLastIncidentTime()`

UnsetLastIncidentTime ensures that no value is present for LastIncidentTime, not even an explicit nil
### GetIsMuted

`func (o *Issue) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *Issue) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *Issue) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *Issue) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetErrorDetails

`func (o *Issue) GetErrorDetails() []string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *Issue) GetErrorDetailsOk() (*[]string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *Issue) SetErrorDetails(v []string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *Issue) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


