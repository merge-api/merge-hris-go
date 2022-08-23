/*
 * Merge HRIS API
 *
 * The unified API for building rich integrations with multiple HR Information System platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_hris_client

import (
	"encoding/json"
)

// Benefit # The Benefit Object ### Description The `Benefit` object is used to represent a Benefit for an employee.  ### Usage Example Fetch from the `LIST Benefits` endpoint and filter by `ID` to show all benefits.
type Benefit struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	Employee NullableString `json:"employee,omitempty"`
	// The name of the benefit provider.
	ProviderName NullableString `json:"provider_name,omitempty"`
	// The type of benefit plan
	BenefitPlanType NullableString `json:"benefit_plan_type,omitempty"`
	// The employee's contribution.
	EmployeeContribution NullableFloat32 `json:"employee_contribution,omitempty"`
	// The company's contribution.
	CompanyContribution NullableFloat32 `json:"company_contribution,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewBenefit instantiates a new Benefit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenefit() *Benefit {
	this := Benefit{}
	return &this
}

// NewBenefitWithDefaults instantiates a new Benefit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenefitWithDefaults() *Benefit {
	this := Benefit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Benefit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benefit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Benefit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Benefit) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Benefit) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Benefit) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Benefit) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Benefit) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetEmployee returns the Employee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetEmployee() string {
	if o == nil || o.Employee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Employee.Get()
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employee.Get(), o.Employee.IsSet()
}

// HasEmployee returns a boolean if a field has been set.
func (o *Benefit) HasEmployee() bool {
	if o != nil && o.Employee.IsSet() {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given NullableString and assigns it to the Employee field.
func (o *Benefit) SetEmployee(v string) {
	o.Employee.Set(&v)
}
// SetEmployeeNil sets the value for Employee to be an explicit nil
func (o *Benefit) SetEmployeeNil() {
	o.Employee.Set(nil)
}

// UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
func (o *Benefit) UnsetEmployee() {
	o.Employee.Unset()
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetProviderName() string {
	if o == nil || o.ProviderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProviderName.Get()
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetProviderNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProviderName.Get(), o.ProviderName.IsSet()
}

// HasProviderName returns a boolean if a field has been set.
func (o *Benefit) HasProviderName() bool {
	if o != nil && o.ProviderName.IsSet() {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given NullableString and assigns it to the ProviderName field.
func (o *Benefit) SetProviderName(v string) {
	o.ProviderName.Set(&v)
}
// SetProviderNameNil sets the value for ProviderName to be an explicit nil
func (o *Benefit) SetProviderNameNil() {
	o.ProviderName.Set(nil)
}

// UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
func (o *Benefit) UnsetProviderName() {
	o.ProviderName.Unset()
}

// GetBenefitPlanType returns the BenefitPlanType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetBenefitPlanType() string {
	if o == nil || o.BenefitPlanType.Get() == nil {
		var ret string
		return ret
	}
	return *o.BenefitPlanType.Get()
}

// GetBenefitPlanTypeOk returns a tuple with the BenefitPlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetBenefitPlanTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BenefitPlanType.Get(), o.BenefitPlanType.IsSet()
}

// HasBenefitPlanType returns a boolean if a field has been set.
func (o *Benefit) HasBenefitPlanType() bool {
	if o != nil && o.BenefitPlanType.IsSet() {
		return true
	}

	return false
}

// SetBenefitPlanType gets a reference to the given NullableString and assigns it to the BenefitPlanType field.
func (o *Benefit) SetBenefitPlanType(v string) {
	o.BenefitPlanType.Set(&v)
}
// SetBenefitPlanTypeNil sets the value for BenefitPlanType to be an explicit nil
func (o *Benefit) SetBenefitPlanTypeNil() {
	o.BenefitPlanType.Set(nil)
}

// UnsetBenefitPlanType ensures that no value is present for BenefitPlanType, not even an explicit nil
func (o *Benefit) UnsetBenefitPlanType() {
	o.BenefitPlanType.Unset()
}

// GetEmployeeContribution returns the EmployeeContribution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetEmployeeContribution() float32 {
	if o == nil || o.EmployeeContribution.Get() == nil {
		var ret float32
		return ret
	}
	return *o.EmployeeContribution.Get()
}

// GetEmployeeContributionOk returns a tuple with the EmployeeContribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetEmployeeContributionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeContribution.Get(), o.EmployeeContribution.IsSet()
}

// HasEmployeeContribution returns a boolean if a field has been set.
func (o *Benefit) HasEmployeeContribution() bool {
	if o != nil && o.EmployeeContribution.IsSet() {
		return true
	}

	return false
}

// SetEmployeeContribution gets a reference to the given NullableFloat32 and assigns it to the EmployeeContribution field.
func (o *Benefit) SetEmployeeContribution(v float32) {
	o.EmployeeContribution.Set(&v)
}
// SetEmployeeContributionNil sets the value for EmployeeContribution to be an explicit nil
func (o *Benefit) SetEmployeeContributionNil() {
	o.EmployeeContribution.Set(nil)
}

// UnsetEmployeeContribution ensures that no value is present for EmployeeContribution, not even an explicit nil
func (o *Benefit) UnsetEmployeeContribution() {
	o.EmployeeContribution.Unset()
}

// GetCompanyContribution returns the CompanyContribution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetCompanyContribution() float32 {
	if o == nil || o.CompanyContribution.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CompanyContribution.Get()
}

// GetCompanyContributionOk returns a tuple with the CompanyContribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetCompanyContributionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompanyContribution.Get(), o.CompanyContribution.IsSet()
}

// HasCompanyContribution returns a boolean if a field has been set.
func (o *Benefit) HasCompanyContribution() bool {
	if o != nil && o.CompanyContribution.IsSet() {
		return true
	}

	return false
}

// SetCompanyContribution gets a reference to the given NullableFloat32 and assigns it to the CompanyContribution field.
func (o *Benefit) SetCompanyContribution(v float32) {
	o.CompanyContribution.Set(&v)
}
// SetCompanyContributionNil sets the value for CompanyContribution to be an explicit nil
func (o *Benefit) SetCompanyContributionNil() {
	o.CompanyContribution.Set(nil)
}

// UnsetCompanyContribution ensures that no value is present for CompanyContribution, not even an explicit nil
func (o *Benefit) UnsetCompanyContribution() {
	o.CompanyContribution.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Benefit) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Benefit) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Benefit) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benefit) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Benefit) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Benefit) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

func (o Benefit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Employee.IsSet() {
		toSerialize["employee"] = o.Employee.Get()
	}
	if o.ProviderName.IsSet() {
		toSerialize["provider_name"] = o.ProviderName.Get()
	}
	if o.BenefitPlanType.IsSet() {
		toSerialize["benefit_plan_type"] = o.BenefitPlanType.Get()
	}
	if o.EmployeeContribution.IsSet() {
		toSerialize["employee_contribution"] = o.EmployeeContribution.Get()
	}
	if o.CompanyContribution.IsSet() {
		toSerialize["company_contribution"] = o.CompanyContribution.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	return json.Marshal(toSerialize)
}

func (v *Benefit) UnmarshalJSON(src []byte) error {
    type BenefitUnmarshalTarget Benefit

	var intermediateResult BenefitUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Benefit(intermediateResult)
	return nil
}
type NullableBenefit struct {
	value *Benefit
	isSet bool
}

func (v NullableBenefit) Get() *Benefit {
	return v.value
}

func (v *NullableBenefit) Set(val *Benefit) {
	v.value = val
	v.isSet = true
}

func (v NullableBenefit) IsSet() bool {
	return v.isSet
}

func (v *NullableBenefit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenefit(val *Benefit) *NullableBenefit {
	return &NullableBenefit{value: val, isSet: true}
}

func (v NullableBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBenefit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


