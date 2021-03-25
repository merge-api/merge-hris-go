# Tax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EmployeePayrollRun** | Pointer to **NullableString** | The tax&#39;s employee payroll run. | [optional] 
**Name** | Pointer to **NullableString** | The tax&#39;s name. | [optional] 
**Amount** | Pointer to **NullableFloat32** | The tax amount. | [optional] 
**EmployerTax** | Pointer to **NullableBool** | Whether or not the employer is responsible for paying the tax. | [optional] 
**RemoteData** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewTax

`func NewTax() *Tax`

NewTax instantiates a new Tax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithDefaults

`func NewTaxWithDefaults() *Tax`

NewTaxWithDefaults instantiates a new Tax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tax) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tax) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tax) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tax) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmployeePayrollRun

`func (o *Tax) GetEmployeePayrollRun() string`

GetEmployeePayrollRun returns the EmployeePayrollRun field if non-nil, zero value otherwise.

### GetEmployeePayrollRunOk

`func (o *Tax) GetEmployeePayrollRunOk() (*string, bool)`

GetEmployeePayrollRunOk returns a tuple with the EmployeePayrollRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeePayrollRun

`func (o *Tax) SetEmployeePayrollRun(v string)`

SetEmployeePayrollRun sets EmployeePayrollRun field to given value.

### HasEmployeePayrollRun

`func (o *Tax) HasEmployeePayrollRun() bool`

HasEmployeePayrollRun returns a boolean if a field has been set.

### SetEmployeePayrollRunNil

`func (o *Tax) SetEmployeePayrollRunNil(b bool)`

 SetEmployeePayrollRunNil sets the value for EmployeePayrollRun to be an explicit nil

### UnsetEmployeePayrollRun
`func (o *Tax) UnsetEmployeePayrollRun()`

UnsetEmployeePayrollRun ensures that no value is present for EmployeePayrollRun, not even an explicit nil
### GetName

`func (o *Tax) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tax) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tax) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tax) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Tax) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Tax) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAmount

`func (o *Tax) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Tax) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Tax) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Tax) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *Tax) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *Tax) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetEmployerTax

`func (o *Tax) GetEmployerTax() bool`

GetEmployerTax returns the EmployerTax field if non-nil, zero value otherwise.

### GetEmployerTaxOk

`func (o *Tax) GetEmployerTaxOk() (*bool, bool)`

GetEmployerTaxOk returns a tuple with the EmployerTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerTax

`func (o *Tax) SetEmployerTax(v bool)`

SetEmployerTax sets EmployerTax field to given value.

### HasEmployerTax

`func (o *Tax) HasEmployerTax() bool`

HasEmployerTax returns a boolean if a field has been set.

### SetEmployerTaxNil

`func (o *Tax) SetEmployerTaxNil(b bool)`

 SetEmployerTaxNil sets the value for EmployerTax to be an explicit nil

### UnsetEmployerTax
`func (o *Tax) UnsetEmployerTax()`

UnsetEmployerTax ensures that no value is present for EmployerTax, not even an explicit nil
### GetRemoteData

`func (o *Tax) GetRemoteData() []map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Tax) GetRemoteDataOk() (*[]map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Tax) SetRemoteData(v []map[string]interface{})`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Tax) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Tax) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Tax) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


