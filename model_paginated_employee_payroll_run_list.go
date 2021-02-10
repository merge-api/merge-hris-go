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

// PaginatedEmployeePayrollRunList struct for PaginatedEmployeePayrollRunList
type PaginatedEmployeePayrollRunList struct {
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results *[]EmployeePayrollRun `json:"results,omitempty"`
}

// NewPaginatedEmployeePayrollRunList instantiates a new PaginatedEmployeePayrollRunList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEmployeePayrollRunList() *PaginatedEmployeePayrollRunList {
	this := PaginatedEmployeePayrollRunList{}
	return &this
}

// NewPaginatedEmployeePayrollRunListWithDefaults instantiates a new PaginatedEmployeePayrollRunList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEmployeePayrollRunListWithDefaults() *PaginatedEmployeePayrollRunList {
	this := PaginatedEmployeePayrollRunList{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedEmployeePayrollRunList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedEmployeePayrollRunList) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedEmployeePayrollRunList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedEmployeePayrollRunList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedEmployeePayrollRunList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedEmployeePayrollRunList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedEmployeePayrollRunList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedEmployeePayrollRunList) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedEmployeePayrollRunList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedEmployeePayrollRunList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedEmployeePayrollRunList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedEmployeePayrollRunList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedEmployeePayrollRunList) GetResults() []EmployeePayrollRun {
	if o == nil || o.Results == nil {
		var ret []EmployeePayrollRun
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEmployeePayrollRunList) GetResultsOk() (*[]EmployeePayrollRun, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedEmployeePayrollRunList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EmployeePayrollRun and assigns it to the Results field.
func (o *PaginatedEmployeePayrollRunList) SetResults(v []EmployeePayrollRun) {
	o.Results = &v
}

func (o PaginatedEmployeePayrollRunList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEmployeePayrollRunList struct {
	value *PaginatedEmployeePayrollRunList
	isSet bool
}

func (v NullablePaginatedEmployeePayrollRunList) Get() *PaginatedEmployeePayrollRunList {
	return v.value
}

func (v *NullablePaginatedEmployeePayrollRunList) Set(val *PaginatedEmployeePayrollRunList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEmployeePayrollRunList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEmployeePayrollRunList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEmployeePayrollRunList(val *PaginatedEmployeePayrollRunList) *NullablePaginatedEmployeePayrollRunList {
	return &NullablePaginatedEmployeePayrollRunList{value: val, isSet: true}
}

func (v NullablePaginatedEmployeePayrollRunList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEmployeePayrollRunList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


