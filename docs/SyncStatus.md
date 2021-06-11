# SyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** |  | 
**ModelId** | **string** |  | 
**LastSyncStart** | **time.Time** |  | 
**NextSyncStart** | **time.Time** |  | 
**Status** | [**SyncStatusStatusEnum**](SyncStatusStatusEnum.md) |  | 
**IsInitialSync** | **bool** |  | 

## Methods

### NewSyncStatus

`func NewSyncStatus(modelName string, modelId string, lastSyncStart time.Time, nextSyncStart time.Time, status SyncStatusStatusEnum, isInitialSync bool, ) *SyncStatus`

NewSyncStatus instantiates a new SyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStatusWithDefaults

`func NewSyncStatusWithDefaults() *SyncStatus`

NewSyncStatusWithDefaults instantiates a new SyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *SyncStatus) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *SyncStatus) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *SyncStatus) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelId

`func (o *SyncStatus) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *SyncStatus) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *SyncStatus) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetLastSyncStart

`func (o *SyncStatus) GetLastSyncStart() time.Time`

GetLastSyncStart returns the LastSyncStart field if non-nil, zero value otherwise.

### GetLastSyncStartOk

`func (o *SyncStatus) GetLastSyncStartOk() (*time.Time, bool)`

GetLastSyncStartOk returns a tuple with the LastSyncStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStart

`func (o *SyncStatus) SetLastSyncStart(v time.Time)`

SetLastSyncStart sets LastSyncStart field to given value.


### GetNextSyncStart

`func (o *SyncStatus) GetNextSyncStart() time.Time`

GetNextSyncStart returns the NextSyncStart field if non-nil, zero value otherwise.

### GetNextSyncStartOk

`func (o *SyncStatus) GetNextSyncStartOk() (*time.Time, bool)`

GetNextSyncStartOk returns a tuple with the NextSyncStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSyncStart

`func (o *SyncStatus) SetNextSyncStart(v time.Time)`

SetNextSyncStart sets NextSyncStart field to given value.


### GetStatus

`func (o *SyncStatus) GetStatus() SyncStatusStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncStatus) GetStatusOk() (*SyncStatusStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncStatus) SetStatus(v SyncStatusStatusEnum)`

SetStatus sets Status field to given value.


### GetIsInitialSync

`func (o *SyncStatus) GetIsInitialSync() bool`

GetIsInitialSync returns the IsInitialSync field if non-nil, zero value otherwise.

### GetIsInitialSyncOk

`func (o *SyncStatus) GetIsInitialSyncOk() (*bool, bool)`

GetIsInitialSyncOk returns a tuple with the IsInitialSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitialSync

`func (o *SyncStatus) SetIsInitialSync(v bool)`

SetIsInitialSync sets IsInitialSync field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


