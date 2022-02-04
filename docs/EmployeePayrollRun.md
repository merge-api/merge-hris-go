# EmployeePayrollRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Employee** | Pointer to **NullableString** |  | [optional] 
**PayrollRun** | Pointer to **NullableString** |  | [optional] 
**GrossPay** | Pointer to **NullableFloat32** | The gross pay from the run. | [optional] 
**NetPay** | Pointer to **NullableFloat32** | The net pay from the run. | [optional] 
**StartDate** | Pointer to **NullableTime** | The day and time the payroll run started. | [optional] 
**EndDate** | Pointer to **NullableTime** | The day and time the payroll run ended. | [optional] 
**CheckDate** | Pointer to **NullableTime** | The day and time the payroll run was checked. | [optional] 
**Earnings** | Pointer to [**[]Earning**](Earning.md) |  | [optional] [readonly] 
**Deductions** | Pointer to [**[]Deduction**](Deduction.md) |  | [optional] [readonly] 
**Taxes** | Pointer to [**[]Tax**](Tax.md) |  | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewEmployeePayrollRun

`func NewEmployeePayrollRun() *EmployeePayrollRun`

NewEmployeePayrollRun instantiates a new EmployeePayrollRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeePayrollRunWithDefaults

`func NewEmployeePayrollRunWithDefaults() *EmployeePayrollRun`

NewEmployeePayrollRunWithDefaults instantiates a new EmployeePayrollRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeePayrollRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeePayrollRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeePayrollRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeePayrollRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *EmployeePayrollRun) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *EmployeePayrollRun) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *EmployeePayrollRun) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *EmployeePayrollRun) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *EmployeePayrollRun) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *EmployeePayrollRun) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployee

`func (o *EmployeePayrollRun) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *EmployeePayrollRun) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *EmployeePayrollRun) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *EmployeePayrollRun) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *EmployeePayrollRun) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *EmployeePayrollRun) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetPayrollRun

`func (o *EmployeePayrollRun) GetPayrollRun() string`

GetPayrollRun returns the PayrollRun field if non-nil, zero value otherwise.

### GetPayrollRunOk

`func (o *EmployeePayrollRun) GetPayrollRunOk() (*string, bool)`

GetPayrollRunOk returns a tuple with the PayrollRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollRun

`func (o *EmployeePayrollRun) SetPayrollRun(v string)`

SetPayrollRun sets PayrollRun field to given value.

### HasPayrollRun

`func (o *EmployeePayrollRun) HasPayrollRun() bool`

HasPayrollRun returns a boolean if a field has been set.

### SetPayrollRunNil

`func (o *EmployeePayrollRun) SetPayrollRunNil(b bool)`

 SetPayrollRunNil sets the value for PayrollRun to be an explicit nil

### UnsetPayrollRun
`func (o *EmployeePayrollRun) UnsetPayrollRun()`

UnsetPayrollRun ensures that no value is present for PayrollRun, not even an explicit nil
### GetGrossPay

`func (o *EmployeePayrollRun) GetGrossPay() float32`

GetGrossPay returns the GrossPay field if non-nil, zero value otherwise.

### GetGrossPayOk

`func (o *EmployeePayrollRun) GetGrossPayOk() (*float32, bool)`

GetGrossPayOk returns a tuple with the GrossPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPay

`func (o *EmployeePayrollRun) SetGrossPay(v float32)`

SetGrossPay sets GrossPay field to given value.

### HasGrossPay

`func (o *EmployeePayrollRun) HasGrossPay() bool`

HasGrossPay returns a boolean if a field has been set.

### SetGrossPayNil

`func (o *EmployeePayrollRun) SetGrossPayNil(b bool)`

 SetGrossPayNil sets the value for GrossPay to be an explicit nil

### UnsetGrossPay
`func (o *EmployeePayrollRun) UnsetGrossPay()`

UnsetGrossPay ensures that no value is present for GrossPay, not even an explicit nil
### GetNetPay

`func (o *EmployeePayrollRun) GetNetPay() float32`

GetNetPay returns the NetPay field if non-nil, zero value otherwise.

### GetNetPayOk

`func (o *EmployeePayrollRun) GetNetPayOk() (*float32, bool)`

GetNetPayOk returns a tuple with the NetPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPay

`func (o *EmployeePayrollRun) SetNetPay(v float32)`

SetNetPay sets NetPay field to given value.

### HasNetPay

`func (o *EmployeePayrollRun) HasNetPay() bool`

HasNetPay returns a boolean if a field has been set.

### SetNetPayNil

`func (o *EmployeePayrollRun) SetNetPayNil(b bool)`

 SetNetPayNil sets the value for NetPay to be an explicit nil

### UnsetNetPay
`func (o *EmployeePayrollRun) UnsetNetPay()`

UnsetNetPay ensures that no value is present for NetPay, not even an explicit nil
### GetStartDate

`func (o *EmployeePayrollRun) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmployeePayrollRun) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmployeePayrollRun) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EmployeePayrollRun) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *EmployeePayrollRun) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EmployeePayrollRun) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *EmployeePayrollRun) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EmployeePayrollRun) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EmployeePayrollRun) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EmployeePayrollRun) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *EmployeePayrollRun) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *EmployeePayrollRun) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetCheckDate

`func (o *EmployeePayrollRun) GetCheckDate() time.Time`

GetCheckDate returns the CheckDate field if non-nil, zero value otherwise.

### GetCheckDateOk

`func (o *EmployeePayrollRun) GetCheckDateOk() (*time.Time, bool)`

GetCheckDateOk returns a tuple with the CheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckDate

`func (o *EmployeePayrollRun) SetCheckDate(v time.Time)`

SetCheckDate sets CheckDate field to given value.

### HasCheckDate

`func (o *EmployeePayrollRun) HasCheckDate() bool`

HasCheckDate returns a boolean if a field has been set.

### SetCheckDateNil

`func (o *EmployeePayrollRun) SetCheckDateNil(b bool)`

 SetCheckDateNil sets the value for CheckDate to be an explicit nil

### UnsetCheckDate
`func (o *EmployeePayrollRun) UnsetCheckDate()`

UnsetCheckDate ensures that no value is present for CheckDate, not even an explicit nil
### GetEarnings

`func (o *EmployeePayrollRun) GetEarnings() []Earning`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *EmployeePayrollRun) GetEarningsOk() (*[]Earning, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *EmployeePayrollRun) SetEarnings(v []Earning)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *EmployeePayrollRun) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetDeductions

`func (o *EmployeePayrollRun) GetDeductions() []Deduction`

GetDeductions returns the Deductions field if non-nil, zero value otherwise.

### GetDeductionsOk

`func (o *EmployeePayrollRun) GetDeductionsOk() (*[]Deduction, bool)`

GetDeductionsOk returns a tuple with the Deductions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductions

`func (o *EmployeePayrollRun) SetDeductions(v []Deduction)`

SetDeductions sets Deductions field to given value.

### HasDeductions

`func (o *EmployeePayrollRun) HasDeductions() bool`

HasDeductions returns a boolean if a field has been set.

### GetTaxes

`func (o *EmployeePayrollRun) GetTaxes() []Tax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *EmployeePayrollRun) GetTaxesOk() (*[]Tax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *EmployeePayrollRun) SetTaxes(v []Tax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *EmployeePayrollRun) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetRemoteData

`func (o *EmployeePayrollRun) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *EmployeePayrollRun) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *EmployeePayrollRun) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *EmployeePayrollRun) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *EmployeePayrollRun) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *EmployeePayrollRun) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


