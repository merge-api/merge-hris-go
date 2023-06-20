# RemoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** |  | 
**Path** | **string** |  | 
**Status** | **int32** |  | 
**Response** | **interface{}** |  | 
**ResponseHeaders** | Pointer to **map[string]interface{}** |  | [optional] 
**ResponseType** | Pointer to [**ResponseTypeEnum**](ResponseTypeEnum.md) |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRemoteResponse

`func NewRemoteResponse(method string, path string, status int32, response interface{}, ) *RemoteResponse`

NewRemoteResponse instantiates a new RemoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteResponseWithDefaults

`func NewRemoteResponseWithDefaults() *RemoteResponse`

NewRemoteResponseWithDefaults instantiates a new RemoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *RemoteResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RemoteResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RemoteResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *RemoteResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RemoteResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RemoteResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetStatus

`func (o *RemoteResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetResponse

`func (o *RemoteResponse) GetResponse() interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RemoteResponse) GetResponseOk() (*interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RemoteResponse) SetResponse(v interface{})`

SetResponse sets Response field to given value.


### SetResponseNil

`func (o *RemoteResponse) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *RemoteResponse) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseHeaders

`func (o *RemoteResponse) GetResponseHeaders() map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *RemoteResponse) GetResponseHeadersOk() (*map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *RemoteResponse) SetResponseHeaders(v map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *RemoteResponse) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResponseType

`func (o *RemoteResponse) GetResponseType() ResponseTypeEnum`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *RemoteResponse) GetResponseTypeOk() (*ResponseTypeEnum, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *RemoteResponse) SetResponseType(v ResponseTypeEnum)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *RemoteResponse) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetHeaders

`func (o *RemoteResponse) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RemoteResponse) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RemoteResponse) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RemoteResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


