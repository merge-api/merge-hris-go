# LinkedAccountCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionSchemaId** | **string** | The ID indicating which condition schema to use for a specific condition. | 
**CommonModel** | Pointer to **string** | The common model for a specific condition. | [optional] [readonly] 
**NativeName** | **NullableString** | User-facing *native condition* name. e.g. \&quot;Skip Manager\&quot;. | 
**Operator** | **string** | The operator for a specific condition. | 
**Value** | Pointer to **interface{}** | The value for a condition. | [optional] [readonly] 
**FieldName** | **NullableString** | The name of the field on the common model that this condition corresponds to, if they conceptually match. e.g. \&quot;location_type\&quot;. | 

## Methods

### NewLinkedAccountCondition

`func NewLinkedAccountCondition(conditionSchemaId string, nativeName NullableString, operator string, fieldName NullableString, ) *LinkedAccountCondition`

NewLinkedAccountCondition instantiates a new LinkedAccountCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAccountConditionWithDefaults

`func NewLinkedAccountConditionWithDefaults() *LinkedAccountCondition`

NewLinkedAccountConditionWithDefaults instantiates a new LinkedAccountCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionSchemaId

`func (o *LinkedAccountCondition) GetConditionSchemaId() string`

GetConditionSchemaId returns the ConditionSchemaId field if non-nil, zero value otherwise.

### GetConditionSchemaIdOk

`func (o *LinkedAccountCondition) GetConditionSchemaIdOk() (*string, bool)`

GetConditionSchemaIdOk returns a tuple with the ConditionSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSchemaId

`func (o *LinkedAccountCondition) SetConditionSchemaId(v string)`

SetConditionSchemaId sets ConditionSchemaId field to given value.


### GetCommonModel

`func (o *LinkedAccountCondition) GetCommonModel() string`

GetCommonModel returns the CommonModel field if non-nil, zero value otherwise.

### GetCommonModelOk

`func (o *LinkedAccountCondition) GetCommonModelOk() (*string, bool)`

GetCommonModelOk returns a tuple with the CommonModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonModel

`func (o *LinkedAccountCondition) SetCommonModel(v string)`

SetCommonModel sets CommonModel field to given value.

### HasCommonModel

`func (o *LinkedAccountCondition) HasCommonModel() bool`

HasCommonModel returns a boolean if a field has been set.

### GetNativeName

`func (o *LinkedAccountCondition) GetNativeName() string`

GetNativeName returns the NativeName field if non-nil, zero value otherwise.

### GetNativeNameOk

`func (o *LinkedAccountCondition) GetNativeNameOk() (*string, bool)`

GetNativeNameOk returns a tuple with the NativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeName

`func (o *LinkedAccountCondition) SetNativeName(v string)`

SetNativeName sets NativeName field to given value.


### SetNativeNameNil

`func (o *LinkedAccountCondition) SetNativeNameNil(b bool)`

 SetNativeNameNil sets the value for NativeName to be an explicit nil

### UnsetNativeName
`func (o *LinkedAccountCondition) UnsetNativeName()`

UnsetNativeName ensures that no value is present for NativeName, not even an explicit nil
### GetOperator

`func (o *LinkedAccountCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LinkedAccountCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LinkedAccountCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *LinkedAccountCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LinkedAccountCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LinkedAccountCondition) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *LinkedAccountCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *LinkedAccountCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LinkedAccountCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetFieldName

`func (o *LinkedAccountCondition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *LinkedAccountCondition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *LinkedAccountCondition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### SetFieldNameNil

`func (o *LinkedAccountCondition) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *LinkedAccountCondition) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


