# ModelOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** |  | 
**AvailableOperations** | **[]string** |  | 
**RequiredPostParameters** | **[]string** |  | 
**SupportedFields** | **[]string** |  | 

## Methods

### NewModelOperation

`func NewModelOperation(modelName string, availableOperations []string, requiredPostParameters []string, supportedFields []string, ) *ModelOperation`

NewModelOperation instantiates a new ModelOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOperationWithDefaults

`func NewModelOperationWithDefaults() *ModelOperation`

NewModelOperationWithDefaults instantiates a new ModelOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *ModelOperation) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModelOperation) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModelOperation) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetAvailableOperations

`func (o *ModelOperation) GetAvailableOperations() []string`

GetAvailableOperations returns the AvailableOperations field if non-nil, zero value otherwise.

### GetAvailableOperationsOk

`func (o *ModelOperation) GetAvailableOperationsOk() (*[]string, bool)`

GetAvailableOperationsOk returns a tuple with the AvailableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOperations

`func (o *ModelOperation) SetAvailableOperations(v []string)`

SetAvailableOperations sets AvailableOperations field to given value.


### GetRequiredPostParameters

`func (o *ModelOperation) GetRequiredPostParameters() []string`

GetRequiredPostParameters returns the RequiredPostParameters field if non-nil, zero value otherwise.

### GetRequiredPostParametersOk

`func (o *ModelOperation) GetRequiredPostParametersOk() (*[]string, bool)`

GetRequiredPostParametersOk returns a tuple with the RequiredPostParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPostParameters

`func (o *ModelOperation) SetRequiredPostParameters(v []string)`

SetRequiredPostParameters sets RequiredPostParameters field to given value.


### GetSupportedFields

`func (o *ModelOperation) GetSupportedFields() []string`

GetSupportedFields returns the SupportedFields field if non-nil, zero value otherwise.

### GetSupportedFieldsOk

`func (o *ModelOperation) GetSupportedFieldsOk() (*[]string, bool)`

GetSupportedFieldsOk returns a tuple with the SupportedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFields

`func (o *ModelOperation) SetSupportedFields(v []string)`

SetSupportedFields sets SupportedFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


