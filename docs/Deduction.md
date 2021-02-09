# Deduction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EmployeePayrollRun** | Pointer to **NullableString** | The deduction&#39;s employee payroll run. | [optional] 
**Name** | Pointer to **NullableString** | The deduction&#39;s name. | [optional] 
**EmployeeDeduction** | Pointer to **NullableFloat32** | The amount the employee is deducting. | [optional] 
**CompanyDeduction** | Pointer to **NullableFloat32** | The amount the company is deducting. | [optional] 

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

`func (o *Deduction) GetEmployeeDeduction() float32`

GetEmployeeDeduction returns the EmployeeDeduction field if non-nil, zero value otherwise.

### GetEmployeeDeductionOk

`func (o *Deduction) GetEmployeeDeductionOk() (*float32, bool)`

GetEmployeeDeductionOk returns a tuple with the EmployeeDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeDeduction

`func (o *Deduction) SetEmployeeDeduction(v float32)`

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

`func (o *Deduction) GetCompanyDeduction() float32`

GetCompanyDeduction returns the CompanyDeduction field if non-nil, zero value otherwise.

### GetCompanyDeductionOk

`func (o *Deduction) GetCompanyDeductionOk() (*float32, bool)`

GetCompanyDeductionOk returns a tuple with the CompanyDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDeduction

`func (o *Deduction) SetCompanyDeduction(v float32)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


