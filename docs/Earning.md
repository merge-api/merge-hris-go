# Earning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EmployeePayrollRun** | Pointer to **NullableString** | The earning&#39;s employee payroll run. | [optional] 
**Amount** | Pointer to **NullableFloat32** | The amount earned. | [optional] 
**Type** | Pointer to [**NullableTypeEnum**](TypeEnum.md) | The type of earning. | [optional] 

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

`func (o *Earning) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Earning) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Earning) SetAmount(v float32)`

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

`func (o *Earning) GetType() TypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Earning) GetTypeOk() (*TypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Earning) SetType(v TypeEnum)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


