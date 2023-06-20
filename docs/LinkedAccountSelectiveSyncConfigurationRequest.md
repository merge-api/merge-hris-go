# LinkedAccountSelectiveSyncConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedAccountConditions** | [**[]LinkedAccountConditionRequest**](LinkedAccountConditionRequest.md) | The conditions belonging to a selective sync. | 

## Methods

### NewLinkedAccountSelectiveSyncConfigurationRequest

`func NewLinkedAccountSelectiveSyncConfigurationRequest(linkedAccountConditions []LinkedAccountConditionRequest, ) *LinkedAccountSelectiveSyncConfigurationRequest`

NewLinkedAccountSelectiveSyncConfigurationRequest instantiates a new LinkedAccountSelectiveSyncConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAccountSelectiveSyncConfigurationRequestWithDefaults

`func NewLinkedAccountSelectiveSyncConfigurationRequestWithDefaults() *LinkedAccountSelectiveSyncConfigurationRequest`

NewLinkedAccountSelectiveSyncConfigurationRequestWithDefaults instantiates a new LinkedAccountSelectiveSyncConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedAccountConditions

`func (o *LinkedAccountSelectiveSyncConfigurationRequest) GetLinkedAccountConditions() []LinkedAccountConditionRequest`

GetLinkedAccountConditions returns the LinkedAccountConditions field if non-nil, zero value otherwise.

### GetLinkedAccountConditionsOk

`func (o *LinkedAccountSelectiveSyncConfigurationRequest) GetLinkedAccountConditionsOk() (*[]LinkedAccountConditionRequest, bool)`

GetLinkedAccountConditionsOk returns a tuple with the LinkedAccountConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountConditions

`func (o *LinkedAccountSelectiveSyncConfigurationRequest) SetLinkedAccountConditions(v []LinkedAccountConditionRequest)`

SetLinkedAccountConditions sets LinkedAccountConditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


