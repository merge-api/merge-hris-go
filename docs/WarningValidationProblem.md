# WarningValidationProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**ValidationProblemSource**](ValidationProblemSource.md) |  | [optional] 
**Title** | **string** |  | 
**Detail** | **string** |  | 
**ProblemType** | **string** |  | 

## Methods

### NewWarningValidationProblem

`func NewWarningValidationProblem(title string, detail string, problemType string, ) *WarningValidationProblem`

NewWarningValidationProblem instantiates a new WarningValidationProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningValidationProblemWithDefaults

`func NewWarningValidationProblemWithDefaults() *WarningValidationProblem`

NewWarningValidationProblemWithDefaults instantiates a new WarningValidationProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *WarningValidationProblem) GetSource() ValidationProblemSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WarningValidationProblem) GetSourceOk() (*ValidationProblemSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WarningValidationProblem) SetSource(v ValidationProblemSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *WarningValidationProblem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *WarningValidationProblem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WarningValidationProblem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WarningValidationProblem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *WarningValidationProblem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *WarningValidationProblem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *WarningValidationProblem) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetProblemType

`func (o *WarningValidationProblem) GetProblemType() string`

GetProblemType returns the ProblemType field if non-nil, zero value otherwise.

### GetProblemTypeOk

`func (o *WarningValidationProblem) GetProblemTypeOk() (*string, bool)`

GetProblemTypeOk returns a tuple with the ProblemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemType

`func (o *WarningValidationProblem) SetProblemType(v string)`

SetProblemType sets ProblemType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


