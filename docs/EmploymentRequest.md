# EmploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**JobTitle** | Pointer to **NullableString** | The position&#39;s title. | [optional] 
**PayRate** | Pointer to **NullableFloat32** | The position&#39;s pay rate in dollars. | [optional] 
**PayPeriod** | Pointer to [**NullablePayPeriodEnum**](PayPeriodEnum.md) | The time period this pay rate encompasses. | [optional] 
**PayFrequency** | Pointer to [**NullablePayFrequencyEnum**](PayFrequencyEnum.md) | The position&#39;s pay frequency. | [optional] 
**PayCurrency** | Pointer to [**NullablePayCurrencyEnum**](PayCurrencyEnum.md) | The position&#39;s currency code. | [optional] 
**FlsaStatus** | Pointer to [**NullableFlsaStatusEnum**](FlsaStatusEnum.md) | The position&#39;s FLSA status. | [optional] 
**EffectiveDate** | Pointer to **NullableTime** | The position&#39;s effective date. | [optional] 
**EmploymentType** | Pointer to [**NullableEmploymentTypeEnum**](EmploymentTypeEnum.md) | The position&#39;s type of employment. | [optional] 

## Methods

### NewEmploymentRequest

`func NewEmploymentRequest() *EmploymentRequest`

NewEmploymentRequest instantiates a new EmploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentRequestWithDefaults

`func NewEmploymentRequestWithDefaults() *EmploymentRequest`

NewEmploymentRequestWithDefaults instantiates a new EmploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *EmploymentRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *EmploymentRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *EmploymentRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *EmploymentRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *EmploymentRequest) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *EmploymentRequest) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetJobTitle

`func (o *EmploymentRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *EmploymentRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *EmploymentRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *EmploymentRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *EmploymentRequest) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *EmploymentRequest) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetPayRate

`func (o *EmploymentRequest) GetPayRate() float32`

GetPayRate returns the PayRate field if non-nil, zero value otherwise.

### GetPayRateOk

`func (o *EmploymentRequest) GetPayRateOk() (*float32, bool)`

GetPayRateOk returns a tuple with the PayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRate

`func (o *EmploymentRequest) SetPayRate(v float32)`

SetPayRate sets PayRate field to given value.

### HasPayRate

`func (o *EmploymentRequest) HasPayRate() bool`

HasPayRate returns a boolean if a field has been set.

### SetPayRateNil

`func (o *EmploymentRequest) SetPayRateNil(b bool)`

 SetPayRateNil sets the value for PayRate to be an explicit nil

### UnsetPayRate
`func (o *EmploymentRequest) UnsetPayRate()`

UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
### GetPayPeriod

`func (o *EmploymentRequest) GetPayPeriod() PayPeriodEnum`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *EmploymentRequest) GetPayPeriodOk() (*PayPeriodEnum, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *EmploymentRequest) SetPayPeriod(v PayPeriodEnum)`

SetPayPeriod sets PayPeriod field to given value.

### HasPayPeriod

`func (o *EmploymentRequest) HasPayPeriod() bool`

HasPayPeriod returns a boolean if a field has been set.

### SetPayPeriodNil

`func (o *EmploymentRequest) SetPayPeriodNil(b bool)`

 SetPayPeriodNil sets the value for PayPeriod to be an explicit nil

### UnsetPayPeriod
`func (o *EmploymentRequest) UnsetPayPeriod()`

UnsetPayPeriod ensures that no value is present for PayPeriod, not even an explicit nil
### GetPayFrequency

`func (o *EmploymentRequest) GetPayFrequency() PayFrequencyEnum`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *EmploymentRequest) GetPayFrequencyOk() (*PayFrequencyEnum, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *EmploymentRequest) SetPayFrequency(v PayFrequencyEnum)`

SetPayFrequency sets PayFrequency field to given value.

### HasPayFrequency

`func (o *EmploymentRequest) HasPayFrequency() bool`

HasPayFrequency returns a boolean if a field has been set.

### SetPayFrequencyNil

`func (o *EmploymentRequest) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *EmploymentRequest) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
### GetPayCurrency

`func (o *EmploymentRequest) GetPayCurrency() PayCurrencyEnum`

GetPayCurrency returns the PayCurrency field if non-nil, zero value otherwise.

### GetPayCurrencyOk

`func (o *EmploymentRequest) GetPayCurrencyOk() (*PayCurrencyEnum, bool)`

GetPayCurrencyOk returns a tuple with the PayCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayCurrency

`func (o *EmploymentRequest) SetPayCurrency(v PayCurrencyEnum)`

SetPayCurrency sets PayCurrency field to given value.

### HasPayCurrency

`func (o *EmploymentRequest) HasPayCurrency() bool`

HasPayCurrency returns a boolean if a field has been set.

### SetPayCurrencyNil

`func (o *EmploymentRequest) SetPayCurrencyNil(b bool)`

 SetPayCurrencyNil sets the value for PayCurrency to be an explicit nil

### UnsetPayCurrency
`func (o *EmploymentRequest) UnsetPayCurrency()`

UnsetPayCurrency ensures that no value is present for PayCurrency, not even an explicit nil
### GetFlsaStatus

`func (o *EmploymentRequest) GetFlsaStatus() FlsaStatusEnum`

GetFlsaStatus returns the FlsaStatus field if non-nil, zero value otherwise.

### GetFlsaStatusOk

`func (o *EmploymentRequest) GetFlsaStatusOk() (*FlsaStatusEnum, bool)`

GetFlsaStatusOk returns a tuple with the FlsaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlsaStatus

`func (o *EmploymentRequest) SetFlsaStatus(v FlsaStatusEnum)`

SetFlsaStatus sets FlsaStatus field to given value.

### HasFlsaStatus

`func (o *EmploymentRequest) HasFlsaStatus() bool`

HasFlsaStatus returns a boolean if a field has been set.

### SetFlsaStatusNil

`func (o *EmploymentRequest) SetFlsaStatusNil(b bool)`

 SetFlsaStatusNil sets the value for FlsaStatus to be an explicit nil

### UnsetFlsaStatus
`func (o *EmploymentRequest) UnsetFlsaStatus()`

UnsetFlsaStatus ensures that no value is present for FlsaStatus, not even an explicit nil
### GetEffectiveDate

`func (o *EmploymentRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EmploymentRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EmploymentRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *EmploymentRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *EmploymentRequest) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *EmploymentRequest) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetEmploymentType

`func (o *EmploymentRequest) GetEmploymentType() EmploymentTypeEnum`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *EmploymentRequest) GetEmploymentTypeOk() (*EmploymentTypeEnum, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *EmploymentRequest) SetEmploymentType(v EmploymentTypeEnum)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *EmploymentRequest) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### SetEmploymentTypeNil

`func (o *EmploymentRequest) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *EmploymentRequest) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


