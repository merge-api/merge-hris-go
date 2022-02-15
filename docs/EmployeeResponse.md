# EmployeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**Employee**](Employee.md) |  | 
**Warnings** | [**[]WarningValidationProblem**](WarningValidationProblem.md) |  | 
**Errors** | [**[]ErrorValidationProblem**](ErrorValidationProblem.md) |  | 
**Logs** | Pointer to [**[]DebugModeLog**](DebugModeLog.md) |  | [optional] 

## Methods

### NewEmployeeResponse

`func NewEmployeeResponse(model Employee, warnings []WarningValidationProblem, errors []ErrorValidationProblem, ) *EmployeeResponse`

NewEmployeeResponse instantiates a new EmployeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeResponseWithDefaults

`func NewEmployeeResponseWithDefaults() *EmployeeResponse`

NewEmployeeResponseWithDefaults instantiates a new EmployeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EmployeeResponse) GetModel() Employee`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmployeeResponse) GetModelOk() (*Employee, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmployeeResponse) SetModel(v Employee)`

SetModel sets Model field to given value.


### GetWarnings

`func (o *EmployeeResponse) GetWarnings() []WarningValidationProblem`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *EmployeeResponse) GetWarningsOk() (*[]WarningValidationProblem, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *EmployeeResponse) SetWarnings(v []WarningValidationProblem)`

SetWarnings sets Warnings field to given value.


### GetErrors

`func (o *EmployeeResponse) GetErrors() []ErrorValidationProblem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EmployeeResponse) GetErrorsOk() (*[]ErrorValidationProblem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EmployeeResponse) SetErrors(v []ErrorValidationProblem)`

SetErrors sets Errors field to given value.


### GetLogs

`func (o *EmployeeResponse) GetLogs() []DebugModeLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *EmployeeResponse) GetLogsOk() (*[]DebugModeLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *EmployeeResponse) SetLogs(v []DebugModeLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *EmployeeResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


