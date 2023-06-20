# OperatorSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The operator for which an operator schema is defined. | [optional] [readonly] 
**IsUnique** | Pointer to **bool** | Whether the operator can be repeated multiple times. | [optional] [readonly] 

## Methods

### NewOperatorSchema

`func NewOperatorSchema() *OperatorSchema`

NewOperatorSchema instantiates a new OperatorSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorSchemaWithDefaults

`func NewOperatorSchemaWithDefaults() *OperatorSchema`

NewOperatorSchemaWithDefaults instantiates a new OperatorSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *OperatorSchema) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *OperatorSchema) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *OperatorSchema) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *OperatorSchema) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetIsUnique

`func (o *OperatorSchema) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *OperatorSchema) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *OperatorSchema) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *OperatorSchema) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


