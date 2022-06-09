# Employment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Employee** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** | The position&#39;s title. | [optional] 
**PayRate** | Pointer to **NullableFloat32** | The position&#39;s pay rate in dollars. | [optional] 
**PayPeriod** | Pointer to [**NullablePayPeriodEnum**](PayPeriodEnum.md) | The time period this pay rate encompasses. | [optional] 
**PayFrequency** | Pointer to [**NullablePayFrequencyEnum**](PayFrequencyEnum.md) | The position&#39;s pay frequency. | [optional] 
**PayCurrency** | Pointer to [**NullablePayCurrencyEnum**](PayCurrencyEnum.md) | The position&#39;s currency code. | [optional] 
**PayGroup** | Pointer to **NullableString** |  | [optional] 
**FlsaStatus** | Pointer to [**NullableFlsaStatusEnum**](FlsaStatusEnum.md) | The position&#39;s FLSA status. | [optional] 
**EffectiveDate** | Pointer to **NullableTime** | The position&#39;s effective date. | [optional] 
**EmploymentType** | Pointer to [**NullableEmploymentTypeEnum**](EmploymentTypeEnum.md) | The position&#39;s type of employment. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted on the third-party. | [optional] [readonly] 

## Methods

### NewEmployment

`func NewEmployment() *Employment`

NewEmployment instantiates a new Employment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentWithDefaults

`func NewEmploymentWithDefaults() *Employment`

NewEmploymentWithDefaults instantiates a new Employment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Employment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Employment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Employment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Employment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Employment) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Employment) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Employment) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Employment) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Employment) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Employment) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetEmployee

`func (o *Employment) GetEmployee() string`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *Employment) GetEmployeeOk() (*string, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *Employment) SetEmployee(v string)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *Employment) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *Employment) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *Employment) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetJobTitle

`func (o *Employment) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Employment) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Employment) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Employment) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *Employment) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *Employment) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetPayRate

`func (o *Employment) GetPayRate() float32`

GetPayRate returns the PayRate field if non-nil, zero value otherwise.

### GetPayRateOk

`func (o *Employment) GetPayRateOk() (*float32, bool)`

GetPayRateOk returns a tuple with the PayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRate

`func (o *Employment) SetPayRate(v float32)`

SetPayRate sets PayRate field to given value.

### HasPayRate

`func (o *Employment) HasPayRate() bool`

HasPayRate returns a boolean if a field has been set.

### SetPayRateNil

`func (o *Employment) SetPayRateNil(b bool)`

 SetPayRateNil sets the value for PayRate to be an explicit nil

### UnsetPayRate
`func (o *Employment) UnsetPayRate()`

UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
### GetPayPeriod

`func (o *Employment) GetPayPeriod() PayPeriodEnum`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *Employment) GetPayPeriodOk() (*PayPeriodEnum, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *Employment) SetPayPeriod(v PayPeriodEnum)`

SetPayPeriod sets PayPeriod field to given value.

### HasPayPeriod

`func (o *Employment) HasPayPeriod() bool`

HasPayPeriod returns a boolean if a field has been set.

### SetPayPeriodNil

`func (o *Employment) SetPayPeriodNil(b bool)`

 SetPayPeriodNil sets the value for PayPeriod to be an explicit nil

### UnsetPayPeriod
`func (o *Employment) UnsetPayPeriod()`

UnsetPayPeriod ensures that no value is present for PayPeriod, not even an explicit nil
### GetPayFrequency

`func (o *Employment) GetPayFrequency() PayFrequencyEnum`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *Employment) GetPayFrequencyOk() (*PayFrequencyEnum, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *Employment) SetPayFrequency(v PayFrequencyEnum)`

SetPayFrequency sets PayFrequency field to given value.

### HasPayFrequency

`func (o *Employment) HasPayFrequency() bool`

HasPayFrequency returns a boolean if a field has been set.

### SetPayFrequencyNil

`func (o *Employment) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *Employment) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
### GetPayCurrency

`func (o *Employment) GetPayCurrency() PayCurrencyEnum`

GetPayCurrency returns the PayCurrency field if non-nil, zero value otherwise.

### GetPayCurrencyOk

`func (o *Employment) GetPayCurrencyOk() (*PayCurrencyEnum, bool)`

GetPayCurrencyOk returns a tuple with the PayCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayCurrency

`func (o *Employment) SetPayCurrency(v PayCurrencyEnum)`

SetPayCurrency sets PayCurrency field to given value.

### HasPayCurrency

`func (o *Employment) HasPayCurrency() bool`

HasPayCurrency returns a boolean if a field has been set.

### SetPayCurrencyNil

`func (o *Employment) SetPayCurrencyNil(b bool)`

 SetPayCurrencyNil sets the value for PayCurrency to be an explicit nil

### UnsetPayCurrency
`func (o *Employment) UnsetPayCurrency()`

UnsetPayCurrency ensures that no value is present for PayCurrency, not even an explicit nil
### GetPayGroup

`func (o *Employment) GetPayGroup() string`

GetPayGroup returns the PayGroup field if non-nil, zero value otherwise.

### GetPayGroupOk

`func (o *Employment) GetPayGroupOk() (*string, bool)`

GetPayGroupOk returns a tuple with the PayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayGroup

`func (o *Employment) SetPayGroup(v string)`

SetPayGroup sets PayGroup field to given value.

### HasPayGroup

`func (o *Employment) HasPayGroup() bool`

HasPayGroup returns a boolean if a field has been set.

### SetPayGroupNil

`func (o *Employment) SetPayGroupNil(b bool)`

 SetPayGroupNil sets the value for PayGroup to be an explicit nil

### UnsetPayGroup
`func (o *Employment) UnsetPayGroup()`

UnsetPayGroup ensures that no value is present for PayGroup, not even an explicit nil
### GetFlsaStatus

`func (o *Employment) GetFlsaStatus() FlsaStatusEnum`

GetFlsaStatus returns the FlsaStatus field if non-nil, zero value otherwise.

### GetFlsaStatusOk

`func (o *Employment) GetFlsaStatusOk() (*FlsaStatusEnum, bool)`

GetFlsaStatusOk returns a tuple with the FlsaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlsaStatus

`func (o *Employment) SetFlsaStatus(v FlsaStatusEnum)`

SetFlsaStatus sets FlsaStatus field to given value.

### HasFlsaStatus

`func (o *Employment) HasFlsaStatus() bool`

HasFlsaStatus returns a boolean if a field has been set.

### SetFlsaStatusNil

`func (o *Employment) SetFlsaStatusNil(b bool)`

 SetFlsaStatusNil sets the value for FlsaStatus to be an explicit nil

### UnsetFlsaStatus
`func (o *Employment) UnsetFlsaStatus()`

UnsetFlsaStatus ensures that no value is present for FlsaStatus, not even an explicit nil
### GetEffectiveDate

`func (o *Employment) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Employment) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Employment) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *Employment) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *Employment) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *Employment) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetEmploymentType

`func (o *Employment) GetEmploymentType() EmploymentTypeEnum`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *Employment) GetEmploymentTypeOk() (*EmploymentTypeEnum, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *Employment) SetEmploymentType(v EmploymentTypeEnum)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *Employment) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### SetEmploymentTypeNil

`func (o *Employment) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *Employment) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Employment) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Employment) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Employment) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Employment) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


