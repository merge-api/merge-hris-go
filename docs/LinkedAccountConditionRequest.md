# LinkedAccountConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionSchemaId** | **string** | The ID indicating which condition schema to use for a specific condition. | 
**Operator** | **string** | The operator for a specific condition. | 
**Value** | **interface{}** | The value for a specific condition. | 

## Methods

### NewLinkedAccountConditionRequest

`func NewLinkedAccountConditionRequest(conditionSchemaId string, operator string, value interface{}, ) *LinkedAccountConditionRequest`

NewLinkedAccountConditionRequest instantiates a new LinkedAccountConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAccountConditionRequestWithDefaults

`func NewLinkedAccountConditionRequestWithDefaults() *LinkedAccountConditionRequest`

NewLinkedAccountConditionRequestWithDefaults instantiates a new LinkedAccountConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionSchemaId

`func (o *LinkedAccountConditionRequest) GetConditionSchemaId() string`

GetConditionSchemaId returns the ConditionSchemaId field if non-nil, zero value otherwise.

### GetConditionSchemaIdOk

`func (o *LinkedAccountConditionRequest) GetConditionSchemaIdOk() (*string, bool)`

GetConditionSchemaIdOk returns a tuple with the ConditionSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSchemaId

`func (o *LinkedAccountConditionRequest) SetConditionSchemaId(v string)`

SetConditionSchemaId sets ConditionSchemaId field to given value.


### GetOperator

`func (o *LinkedAccountConditionRequest) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LinkedAccountConditionRequest) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LinkedAccountConditionRequest) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *LinkedAccountConditionRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LinkedAccountConditionRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LinkedAccountConditionRequest) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *LinkedAccountConditionRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LinkedAccountConditionRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


