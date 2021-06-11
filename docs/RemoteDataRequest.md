# RemoteDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRemoteDataRequest

`func NewRemoteDataRequest(path string, ) *RemoteDataRequest`

NewRemoteDataRequest instantiates a new RemoteDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteDataRequestWithDefaults

`func NewRemoteDataRequestWithDefaults() *RemoteDataRequest`

NewRemoteDataRequestWithDefaults instantiates a new RemoteDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RemoteDataRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RemoteDataRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RemoteDataRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetData

`func (o *RemoteDataRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RemoteDataRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RemoteDataRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *RemoteDataRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


