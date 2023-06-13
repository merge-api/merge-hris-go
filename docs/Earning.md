# Earning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**EmployeePayrollRun** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** | The amount earned. | [optional] 
**Type** | Pointer to [**NullableEarningTypeEnum**](EarningTypeEnum.md) | The type of earning.  * &#x60;SALARY&#x60; - SALARY * &#x60;REIMBURSEMENT&#x60; - REIMBURSEMENT * &#x60;OVERTIME&#x60; - OVERTIME * &#x60;BONUS&#x60; - BONUS | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewEarning

`func NewEarning() *Earning`

NewEarning instantiates a new Earning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningWithDefaults

`func NewEarningWithDefaults() *Earning`

NewEarningWithDefaults instantiates a new Earning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Earning) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Earning) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Earning) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Earning) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Earning) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Earning) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Earning) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Earning) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Earning) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Earning) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployeePayrollRun

`func (o *Earning) GetEmployeePayrollRun() string`

GetEmployeePayrollRun returns the EmployeePayrollRun field if non-nil, zero value otherwise.

### GetEmployeePayrollRunOk

`func (o *Earning) GetEmployeePayrollRunOk() (*string, bool)`

GetEmployeePayrollRunOk returns a tuple with the EmployeePayrollRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeePayrollRun

`func (o *Earning) SetEmployeePayrollRun(v string)`

SetEmployeePayrollRun sets EmployeePayrollRun field to given value.

### HasEmployeePayrollRun

`func (o *Earning) HasEmployeePayrollRun() bool`

HasEmployeePayrollRun returns a boolean if a field has been set.

### SetEmployeePayrollRunNil

`func (o *Earning) SetEmployeePayrollRunNil(b bool)`

 SetEmployeePayrollRunNil sets the value for EmployeePayrollRun to be an explicit nil

### UnsetEmployeePayrollRun
`func (o *Earning) UnsetEmployeePayrollRun()`

UnsetEmployeePayrollRun ensures that no value is present for EmployeePayrollRun, not even an explicit nil
### GetAmount

`func (o *Earning) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Earning) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Earning) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Earning) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *Earning) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *Earning) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetType

`func (o *Earning) GetType() EarningTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Earning) GetTypeOk() (*EarningTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Earning) SetType(v EarningTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Earning) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Earning) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Earning) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Earning) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Earning) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Earning) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Earning) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Earning) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Earning) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Earning) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Earning) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Earning) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Earning) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Earning) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Earning) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Earning) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Earning) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Earning) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Earning) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Earning) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Earning) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Earning) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Earning) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


