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

// Tax # The Tax Object ### Description The `Tax` object is used to represent a tax for a given employee's payroll run. One run could include several taxes.  ### Usage Example Fetch from the `LIST Taxes` endpoint and filter by `ID` to show all taxes.
type Tax struct {
	Id *string `json:"id,omitempty"`
	EmployeePayrollRun NullableString `json:"employee_payroll_run,omitempty"`
	// The tax's name.
	Name NullableString `json:"name,omitempty"`
	// The tax amount.
	Amount NullableFloat32 `json:"amount,omitempty"`
	// Whether or not the employer is responsible for paying the tax.
	EmployerTax NullableBool `json:"employer_tax,omitempty"`
	RemoteData []map[string]interface{} `json:"remote_data,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTax instantiates a new Tax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTax() *Tax {
	this := Tax{}
	return &this
}

// NewTaxWithDefaults instantiates a new Tax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxWithDefaults() *Tax {
	this := Tax{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tax) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tax) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tax) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Tax) SetId(v string) {
	o.Id = &v
}

// GetEmployeePayrollRun returns the EmployeePayrollRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tax) GetEmployeePayrollRun() string {
	if o == nil || o.EmployeePayrollRun.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeePayrollRun.Get()
}

// GetEmployeePayrollRunOk returns a tuple with the EmployeePayrollRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tax) GetEmployeePayrollRunOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeePayrollRun.Get(), o.EmployeePayrollRun.IsSet()
}

// HasEmployeePayrollRun returns a boolean if a field has been set.
func (o *Tax) HasEmployeePayrollRun() bool {
	if o != nil && o.EmployeePayrollRun.IsSet() {
		return true
	}

	return false
}

// SetEmployeePayrollRun gets a reference to the given NullableString and assigns it to the EmployeePayrollRun field.
func (o *Tax) SetEmployeePayrollRun(v string) {
	o.EmployeePayrollRun.Set(&v)
}
// SetEmployeePayrollRunNil sets the value for EmployeePayrollRun to be an explicit nil
func (o *Tax) SetEmployeePayrollRunNil() {
	o.EmployeePayrollRun.Set(nil)
}

// UnsetEmployeePayrollRun ensures that no value is present for EmployeePayrollRun, not even an explicit nil
func (o *Tax) UnsetEmployeePayrollRun() {
	o.EmployeePayrollRun.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tax) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tax) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Tax) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Tax) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Tax) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Tax) UnsetName() {
	o.Name.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tax) GetAmount() float32 {
	if o == nil || o.Amount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tax) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Tax) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *Tax) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Tax) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Tax) UnsetAmount() {
	o.Amount.Unset()
}

// GetEmployerTax returns the EmployerTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tax) GetEmployerTax() bool {
	if o == nil || o.EmployerTax.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EmployerTax.Get()
}

// GetEmployerTaxOk returns a tuple with the EmployerTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tax) GetEmployerTaxOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployerTax.Get(), o.EmployerTax.IsSet()
}

// HasEmployerTax returns a boolean if a field has been set.
func (o *Tax) HasEmployerTax() bool {
	if o != nil && o.EmployerTax.IsSet() {
		return true
	}

	return false
}

// SetEmployerTax gets a reference to the given NullableBool and assigns it to the EmployerTax field.
func (o *Tax) SetEmployerTax(v bool) {
	o.EmployerTax.Set(&v)
}
// SetEmployerTaxNil sets the value for EmployerTax to be an explicit nil
func (o *Tax) SetEmployerTaxNil() {
	o.EmployerTax.Set(nil)
}

// UnsetEmployerTax ensures that no value is present for EmployerTax, not even an explicit nil
func (o *Tax) UnsetEmployerTax() {
	o.EmployerTax.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tax) GetRemoteData() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tax) GetRemoteDataOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Tax) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []map[string]interface{} and assigns it to the RemoteData field.
func (o *Tax) SetRemoteData(v []map[string]interface{}) {
	o.RemoteData = v
}

func (o Tax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EmployeePayrollRun.IsSet() {
		toSerialize["employee_payroll_run"] = o.EmployeePayrollRun.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.EmployerTax.IsSet() {
		toSerialize["employer_tax"] = o.EmployerTax.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

func (v *Tax) UnmarshalJSON(src []byte) error {
    type TaxUnmarshalTarget Tax

	var intermediateResult TaxUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Tax(intermediateResult)
	return nil
}
type NullableTax struct {
	value *Tax
	isSet bool
}

func (v NullableTax) Get() *Tax {
	return v.value
}

func (v *NullableTax) Set(val *Tax) {
	v.value = val
	v.isSet = true
}

func (v NullableTax) IsSet() bool {
	return v.isSet
}

func (v *NullableTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTax(val *Tax) *NullableTax {
	return &NullableTax{value: val, isSet: true}
}

func (v NullableTax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


