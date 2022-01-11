# TimeOffResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**TimeOff**](TimeOff.md) |  | 
**Warnings** | [**[]WarningValidationProblem**](WarningValidationProblem.md) |  | 
**Errors** | [**[]ErrorValidationProblem**](ErrorValidationProblem.md) |  | 

## Methods

### NewTimeOffResponse

`func NewTimeOffResponse(model TimeOff, warnings []WarningValidationProblem, errors []ErrorValidationProblem, ) *TimeOffResponse`

NewTimeOffResponse instantiates a new TimeOffResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffResponseWithDefaults

`func NewTimeOffResponseWithDefaults() *TimeOffResponse`

NewTimeOffResponseWithDefaults instantiates a new TimeOffResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TimeOffResponse) GetModel() TimeOff`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TimeOffResponse) GetModelOk() (*TimeOff, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TimeOffResponse) SetModel(v TimeOff)`

SetModel sets Model field to given value.


### GetWarnings

`func (o *TimeOffResponse) GetWarnings() []WarningValidationProblem`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TimeOffResponse) GetWarningsOk() (*[]WarningValidationProblem, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TimeOffResponse) SetWarnings(v []WarningValidationProblem)`

SetWarnings sets Warnings field to given value.


### GetErrors

`func (o *TimeOffResponse) GetErrors() []ErrorValidationProblem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TimeOffResponse) GetErrorsOk() (*[]ErrorValidationProblem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TimeOffResponse) SetErrors(v []ErrorValidationProblem)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


