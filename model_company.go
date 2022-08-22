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

// Company # The Company Object ### Description The `Company` object is used to represent a Company within the HRIS / Payroll system.  ### Usage Example Fetch from the `LIST Companies` endpoint and filter by `ID` to show all companies.
type Company struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The company's legal name.
	LegalName NullableString `json:"legal_name,omitempty"`
	// The company's display name.
	DisplayName NullableString `json:"display_name,omitempty"`
	// The company's Employer Identification Numbers.
	Eins []string `json:"eins,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany() *Company {
	this := Company{}
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Company) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Company) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Company) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Company) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Company) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Company) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Company) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetLegalName() string {
	if o == nil || o.LegalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *Company) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableString and assigns it to the LegalName field.
func (o *Company) SetLegalName(v string) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *Company) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *Company) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Company) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Company) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Company) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Company) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEins returns the Eins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetEins() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Eins
}

// GetEinsOk returns a tuple with the Eins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetEinsOk() (*[]string, bool) {
	if o == nil || o.Eins == nil {
		return nil, false
	}
	return &o.Eins, true
}

// HasEins returns a boolean if a field has been set.
func (o *Company) HasEins() bool {
	if o != nil && o.Eins != nil {
		return true
	}

	return false
}

// SetEins gets a reference to the given []string and assigns it to the Eins field.
func (o *Company) SetEins(v []string) {
	o.Eins = v
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Company) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Company) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Company) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Company) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Company) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.LegalName.IsSet() {
		toSerialize["legal_name"] = o.LegalName.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.Eins != nil {
		toSerialize["eins"] = o.Eins
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	return json.Marshal(toSerialize)
}

func (v *Company) UnmarshalJSON(src []byte) error {
    type CompanyUnmarshalTarget Company

	var intermediateResult CompanyUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Company(intermediateResult)
	return nil
}
type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


