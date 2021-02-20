# AvailableActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | [**AccountIntegration**](AccountIntegration.md) |  | 
**PassthroughAvailable** | **bool** |  | 
**AvailableModelOperations** | Pointer to [**[]ModelOperation**](ModelOperation.md) |  | [optional] 

## Methods

### NewAvailableActions

`func NewAvailableActions(integration AccountIntegration, passthroughAvailable bool, ) *AvailableActions`

NewAvailableActions instantiates a new AvailableActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableActionsWithDefaults

`func NewAvailableActionsWithDefaults() *AvailableActions`

NewAvailableActionsWithDefaults instantiates a new AvailableActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *AvailableActions) GetIntegration() AccountIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AvailableActions) GetIntegrationOk() (*AccountIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AvailableActions) SetIntegration(v AccountIntegration)`

SetIntegration sets Integration field to given value.


### GetPassthroughAvailable

`func (o *AvailableActions) GetPassthroughAvailable() bool`

GetPassthroughAvailable returns the PassthroughAvailable field if non-nil, zero value otherwise.

### GetPassthroughAvailableOk

`func (o *AvailableActions) GetPassthroughAvailableOk() (*bool, bool)`

GetPassthroughAvailableOk returns a tuple with the PassthroughAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughAvailable

`func (o *AvailableActions) SetPassthroughAvailable(v bool)`

SetPassthroughAvailable sets PassthroughAvailable field to given value.


### GetAvailableModelOperations

`func (o *AvailableActions) GetAvailableModelOperations() []ModelOperation`

GetAvailableModelOperations returns the AvailableModelOperations field if non-nil, zero value otherwise.

### GetAvailableModelOperationsOk

`func (o *AvailableActions) GetAvailableModelOperationsOk() (*[]ModelOperation, bool)`

GetAvailableModelOperationsOk returns a tuple with the AvailableModelOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableModelOperations

`func (o *AvailableActions) SetAvailableModelOperations(v []ModelOperation)`

SetAvailableModelOperations sets AvailableModelOperations field to given value.

### HasAvailableModelOperations

`func (o *AvailableActions) HasAvailableModelOperations() bool`

HasAvailableModelOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


