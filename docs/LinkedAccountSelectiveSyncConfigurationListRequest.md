# LinkedAccountSelectiveSyncConfigurationListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncConfigurations** | [**[]LinkedAccountSelectiveSyncConfigurationRequest**](LinkedAccountSelectiveSyncConfigurationRequest.md) | The selective syncs associated with a linked account. | 

## Methods

### NewLinkedAccountSelectiveSyncConfigurationListRequest

`func NewLinkedAccountSelectiveSyncConfigurationListRequest(syncConfigurations []LinkedAccountSelectiveSyncConfigurationRequest, ) *LinkedAccountSelectiveSyncConfigurationListRequest`

NewLinkedAccountSelectiveSyncConfigurationListRequest instantiates a new LinkedAccountSelectiveSyncConfigurationListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAccountSelectiveSyncConfigurationListRequestWithDefaults

`func NewLinkedAccountSelectiveSyncConfigurationListRequestWithDefaults() *LinkedAccountSelectiveSyncConfigurationListRequest`

NewLinkedAccountSelectiveSyncConfigurationListRequestWithDefaults instantiates a new LinkedAccountSelectiveSyncConfigurationListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncConfigurations

`func (o *LinkedAccountSelectiveSyncConfigurationListRequest) GetSyncConfigurations() []LinkedAccountSelectiveSyncConfigurationRequest`

GetSyncConfigurations returns the SyncConfigurations field if non-nil, zero value otherwise.

### GetSyncConfigurationsOk

`func (o *LinkedAccountSelectiveSyncConfigurationListRequest) GetSyncConfigurationsOk() (*[]LinkedAccountSelectiveSyncConfigurationRequest, bool)`

GetSyncConfigurationsOk returns a tuple with the SyncConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncConfigurations

`func (o *LinkedAccountSelectiveSyncConfigurationListRequest) SetSyncConfigurations(v []LinkedAccountSelectiveSyncConfigurationRequest)`

SetSyncConfigurations sets SyncConfigurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


