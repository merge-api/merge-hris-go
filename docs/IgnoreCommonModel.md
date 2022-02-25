# IgnoreCommonModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**ReasonEnum**](ReasonEnum.md) |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewIgnoreCommonModel

`func NewIgnoreCommonModel(reason ReasonEnum, ) *IgnoreCommonModel`

NewIgnoreCommonModel instantiates a new IgnoreCommonModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgnoreCommonModelWithDefaults

`func NewIgnoreCommonModelWithDefaults() *IgnoreCommonModel`

NewIgnoreCommonModelWithDefaults instantiates a new IgnoreCommonModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *IgnoreCommonModel) GetReason() ReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IgnoreCommonModel) GetReasonOk() (*ReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IgnoreCommonModel) SetReason(v ReasonEnum)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *IgnoreCommonModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IgnoreCommonModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IgnoreCommonModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IgnoreCommonModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


