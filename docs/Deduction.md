# Deduction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**EmployeePayrollRun** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** | The deduction&#39;s name. | [optional] 
**EmployeeDeduction** | Pointer to **NullableFloat64** | The amount of money that is withheld from an employee&#39;s gross pay by the employee. | [optional] 
**CompanyDeduction** | Pointer to **NullableFloat64** | The amount of money that is withheld on behalf of an employee by the company. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewDeduction

`func NewDeduction() *Deduction`

NewDeduction instantiates a new Deduction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductionWithDefaults

`func NewDeductionWithDefaults() *Deduction`

NewDeductionWithDefaults instantiates a new Deduction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deduction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deduction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deduction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deduction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Deduction) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Deduction) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Deduction) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Deduction) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Deduction) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Deduction) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployeePayrollRun

`func (o *Deduction) GetEmployeePayrollRun() string`

GetEmployeePayrollRun returns the EmployeePayrollRun field if non-nil, zero value otherwise.

### GetEmployeePayrollRunOk

`func (o *Deduction) GetEmployeePayrollRunOk() (*string, bool)`

GetEmployeePayrollRunOk returns a tuple with the EmployeePayrollRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeePayrollRun

`func (o *Deduction) SetEmployeePayrollRun(v string)`

SetEmployeePayrollRun sets EmployeePayrollRun field to given value.

### HasEmployeePayrollRun

`func (o *Deduction) HasEmployeePayrollRun() bool`

HasEmployeePayrollRun returns a boolean if a field has been set.

### SetEmployeePayrollRunNil

`func (o *Deduction) SetEmployeePayrollRunNil(b bool)`

 SetEmployeePayrollRunNil sets the value for EmployeePayrollRun to be an explicit nil

### UnsetEmployeePayrollRun
`func (o *Deduction) UnsetEmployeePayrollRun()`

UnsetEmployeePayrollRun ensures that no value is present for EmployeePayrollRun, not even an explicit nil
### GetName

`func (o *Deduction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deduction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deduction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Deduction) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Deduction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Deduction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmployeeDeduction

`func (o *Deduction) GetEmployeeDeduction() float64`

GetEmployeeDeduction returns the EmployeeDeduction field if non-nil, zero value otherwise.

### GetEmployeeDeductionOk

`func (o *Deduction) GetEmployeeDeductionOk() (*float64, bool)`

GetEmployeeDeductionOk returns a tuple with the EmployeeDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeDeduction

`func (o *Deduction) SetEmployeeDeduction(v float64)`

SetEmployeeDeduction sets EmployeeDeduction field to given value.

### HasEmployeeDeduction

`func (o *Deduction) HasEmployeeDeduction() bool`

HasEmployeeDeduction returns a boolean if a field has been set.

### SetEmployeeDeductionNil

`func (o *Deduction) SetEmployeeDeductionNil(b bool)`

 SetEmployeeDeductionNil sets the value for EmployeeDeduction to be an explicit nil

### UnsetEmployeeDeduction
`func (o *Deduction) UnsetEmployeeDeduction()`

UnsetEmployeeDeduction ensures that no value is present for EmployeeDeduction, not even an explicit nil
### GetCompanyDeduction

`func (o *Deduction) GetCompanyDeduction() float64`

GetCompanyDeduction returns the CompanyDeduction field if non-nil, zero value otherwise.

### GetCompanyDeductionOk

`func (o *Deduction) GetCompanyDeductionOk() (*float64, bool)`

GetCompanyDeductionOk returns a tuple with the CompanyDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDeduction

`func (o *Deduction) SetCompanyDeduction(v float64)`

SetCompanyDeduction sets CompanyDeduction field to given value.

### HasCompanyDeduction

`func (o *Deduction) HasCompanyDeduction() bool`

HasCompanyDeduction returns a boolean if a field has been set.

### SetCompanyDeductionNil

`func (o *Deduction) SetCompanyDeductionNil(b bool)`

 SetCompanyDeductionNil sets the value for CompanyDeduction to be an explicit nil

### UnsetCompanyDeduction
`func (o *Deduction) UnsetCompanyDeduction()`

UnsetCompanyDeduction ensures that no value is present for CompanyDeduction, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Deduction) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Deduction) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Deduction) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Deduction) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Deduction) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Deduction) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Deduction) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Deduction) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Deduction) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Deduction) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Deduction) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Deduction) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Deduction) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Deduction) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Deduction) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Deduction) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Deduction) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Deduction) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Deduction) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Deduction) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


