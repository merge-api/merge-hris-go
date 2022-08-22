# Benefit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Employee** | Pointer to **NullableString** |  | [optional] 
**ProviderName** | Pointer to **NullableString** | The name of the benefit provider. | [optional] 
**BenefitPlanType** | Pointer to **NullableString** | The type of benefit plan | [optional] 
**EmployeeContribution** | Pointer to **NullableFloat32** | The employee&#39;s contribution. | [optional] 
**CompanyContribution** | Pointer to **NullableFloat32** | The company&#39;s contribution. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 

## Methods

### NewBenefit

`func NewBenefit() *Benefit`

NewBenefit instantiates a new Benefit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenefitWithDefaults

`func NewBenefitWithDefaults() *Benefit`

NewBenefitWithDefaults instantiates a new Benefit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Benefit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Benefit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Benefit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Benefit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Benefit) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Benefit) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Benefit) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Benefit) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Benefit) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Benefit) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployee

`func (o *Benefit) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *Benefit) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *Benefit) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *Benefit) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *Benefit) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *Benefit) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetProviderName

`func (o *Benefit) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *Benefit) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *Benefit) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *Benefit) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *Benefit) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *Benefit) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetBenefitPlanType

`func (o *Benefit) GetBenefitPlanType() string`

GetBenefitPlanType returns the BenefitPlanType field if non-nil, zero value otherwise.

### GetBenefitPlanTypeOk

`func (o *Benefit) GetBenefitPlanTypeOk() (*string, bool)`

GetBenefitPlanTypeOk returns a tuple with the BenefitPlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitPlanType

`func (o *Benefit) SetBenefitPlanType(v string)`

SetBenefitPlanType sets BenefitPlanType field to given value.

### HasBenefitPlanType

`func (o *Benefit) HasBenefitPlanType() bool`

HasBenefitPlanType returns a boolean if a field has been set.

### SetBenefitPlanTypeNil

`func (o *Benefit) SetBenefitPlanTypeNil(b bool)`

 SetBenefitPlanTypeNil sets the value for BenefitPlanType to be an explicit nil

### UnsetBenefitPlanType
`func (o *Benefit) UnsetBenefitPlanType()`

UnsetBenefitPlanType ensures that no value is present for BenefitPlanType, not even an explicit nil
### GetEmployeeContribution

`func (o *Benefit) GetEmployeeContribution() float32`

GetEmployeeContribution returns the EmployeeContribution field if non-nil, zero value otherwise.

### GetEmployeeContributionOk

`func (o *Benefit) GetEmployeeContributionOk() (*float32, bool)`

GetEmployeeContributionOk returns a tuple with the EmployeeContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeContribution

`func (o *Benefit) SetEmployeeContribution(v float32)`

SetEmployeeContribution sets EmployeeContribution field to given value.

### HasEmployeeContribution

`func (o *Benefit) HasEmployeeContribution() bool`

HasEmployeeContribution returns a boolean if a field has been set.

### SetEmployeeContributionNil

`func (o *Benefit) SetEmployeeContributionNil(b bool)`

 SetEmployeeContributionNil sets the value for EmployeeContribution to be an explicit nil

### UnsetEmployeeContribution
`func (o *Benefit) UnsetEmployeeContribution()`

UnsetEmployeeContribution ensures that no value is present for EmployeeContribution, not even an explicit nil
### GetCompanyContribution

`func (o *Benefit) GetCompanyContribution() float32`

GetCompanyContribution returns the CompanyContribution field if non-nil, zero value otherwise.

### GetCompanyContributionOk

`func (o *Benefit) GetCompanyContributionOk() (*float32, bool)`

GetCompanyContributionOk returns a tuple with the CompanyContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContribution

`func (o *Benefit) SetCompanyContribution(v float32)`

SetCompanyContribution sets CompanyContribution field to given value.

### HasCompanyContribution

`func (o *Benefit) HasCompanyContribution() bool`

HasCompanyContribution returns a boolean if a field has been set.

### SetCompanyContributionNil

`func (o *Benefit) SetCompanyContributionNil(b bool)`

 SetCompanyContributionNil sets the value for CompanyContribution to be an explicit nil

### UnsetCompanyContribution
`func (o *Benefit) UnsetCompanyContribution()`

UnsetCompanyContribution ensures that no value is present for CompanyContribution, not even an explicit nil
### GetRemoteData

`func (o *Benefit) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Benefit) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Benefit) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Benefit) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Benefit) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Benefit) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Benefit) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Benefit) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Benefit) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Benefit) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


