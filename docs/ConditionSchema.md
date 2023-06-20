# ConditionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the condition schema. This ID is used when updating selective syncs for a linked account. | 
**CommonModel** | Pointer to **string** | The common model for which a condition schema is defined. | [optional] [readonly] 
**NativeName** | **NullableString** | User-facing *native condition* name. e.g. \&quot;Skip Manager\&quot;. | 
**FieldName** | **NullableString** | The name of the field on the common model that this condition corresponds to, if they conceptually match. e.g. \&quot;location_type\&quot;. | 
**IsUnique** | Pointer to **bool** | Whether this condition can only be applied once. If false, the condition can be AND&#39;d together multiple times. | [optional] 
**ConditionType** | [**ConditionTypeEnum**](ConditionTypeEnum.md) | The type of value(s) that can be set for this condition.  * &#x60;BOOLEAN&#x60; - BOOLEAN * &#x60;DATE&#x60; - DATE * &#x60;DATE_TIME&#x60; - DATE_TIME * &#x60;INTEGER&#x60; - INTEGER * &#x60;FLOAT&#x60; - FLOAT * &#x60;STRING&#x60; - STRING * &#x60;LIST_OF_STRINGS&#x60; - LIST_OF_STRINGS | 
**Operators** | [**[]OperatorSchema**](OperatorSchema.md) | The schemas for the operators that can be used on a condition. | 

## Methods

### NewConditionSchema

`func NewConditionSchema(id string, nativeName NullableString, fieldName NullableString, conditionType ConditionTypeEnum, operators []OperatorSchema, ) *ConditionSchema`

NewConditionSchema instantiates a new ConditionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionSchemaWithDefaults

`func NewConditionSchemaWithDefaults() *ConditionSchema`

NewConditionSchemaWithDefaults instantiates a new ConditionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConditionSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConditionSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConditionSchema) SetId(v string)`

SetId sets Id field to given value.


### GetCommonModel

`func (o *ConditionSchema) GetCommonModel() string`

GetCommonModel returns the CommonModel field if non-nil, zero value otherwise.

### GetCommonModelOk

`func (o *ConditionSchema) GetCommonModelOk() (*string, bool)`

GetCommonModelOk returns a tuple with the CommonModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonModel

`func (o *ConditionSchema) SetCommonModel(v string)`

SetCommonModel sets CommonModel field to given value.

### HasCommonModel

`func (o *ConditionSchema) HasCommonModel() bool`

HasCommonModel returns a boolean if a field has been set.

### GetNativeName

`func (o *ConditionSchema) GetNativeName() string`

GetNativeName returns the NativeName field if non-nil, zero value otherwise.

### GetNativeNameOk

`func (o *ConditionSchema) GetNativeNameOk() (*string, bool)`

GetNativeNameOk returns a tuple with the NativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeName

`func (o *ConditionSchema) SetNativeName(v string)`

SetNativeName sets NativeName field to given value.


### SetNativeNameNil

`func (o *ConditionSchema) SetNativeNameNil(b bool)`

 SetNativeNameNil sets the value for NativeName to be an explicit nil

### UnsetNativeName
`func (o *ConditionSchema) UnsetNativeName()`

UnsetNativeName ensures that no value is present for NativeName, not even an explicit nil
### GetFieldName

`func (o *ConditionSchema) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ConditionSchema) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ConditionSchema) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### SetFieldNameNil

`func (o *ConditionSchema) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *ConditionSchema) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetIsUnique

`func (o *ConditionSchema) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *ConditionSchema) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *ConditionSchema) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *ConditionSchema) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetConditionType

`func (o *ConditionSchema) GetConditionType() ConditionTypeEnum`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ConditionSchema) GetConditionTypeOk() (*ConditionTypeEnum, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ConditionSchema) SetConditionType(v ConditionTypeEnum)`

SetConditionType sets ConditionType field to given value.


### GetOperators

`func (o *ConditionSchema) GetOperators() []OperatorSchema`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *ConditionSchema) GetOperatorsOk() (*[]OperatorSchema, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *ConditionSchema) SetOperators(v []OperatorSchema)`

SetOperators sets Operators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


