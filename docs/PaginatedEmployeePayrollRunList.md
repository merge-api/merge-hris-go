# PaginatedEmployeePayrollRunList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]EmployeePayrollRun**](EmployeePayrollRun.md) |  | [optional] 

## Methods

### NewPaginatedEmployeePayrollRunList

`func NewPaginatedEmployeePayrollRunList() *PaginatedEmployeePayrollRunList`

NewPaginatedEmployeePayrollRunList instantiates a new PaginatedEmployeePayrollRunList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEmployeePayrollRunListWithDefaults

`func NewPaginatedEmployeePayrollRunListWithDefaults() *PaginatedEmployeePayrollRunList`

NewPaginatedEmployeePayrollRunListWithDefaults instantiates a new PaginatedEmployeePayrollRunList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedEmployeePayrollRunList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedEmployeePayrollRunList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedEmployeePayrollRunList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedEmployeePayrollRunList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedEmployeePayrollRunList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedEmployeePayrollRunList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedEmployeePayrollRunList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedEmployeePayrollRunList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedEmployeePayrollRunList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedEmployeePayrollRunList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedEmployeePayrollRunList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedEmployeePayrollRunList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedEmployeePayrollRunList) GetResults() []EmployeePayrollRun`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedEmployeePayrollRunList) GetResultsOk() (*[]EmployeePayrollRun, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedEmployeePayrollRunList) SetResults(v []EmployeePayrollRun)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedEmployeePayrollRunList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


