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

// TimeOff # The TimeOff Object ### Description The `TimeOff` object is used to represent a Time Off Request filed by an employee.  ### Usage Example Fetch from the `LIST TimeOffs` endpoint and filter by `ID` to show all time off requests.
type TimeOff struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The employee requesting time off.
	Employee NullableString `json:"employee,omitempty"`
	// The employee approving the time off request.
	Approver NullableString `json:"approver,omitempty"`
	// The status of this time off request.
	Status NullableTimeOffStatusEnum `json:"status,omitempty"`
	// The status of this time off request.
	EmployeeNote NullableString `json:"employee_note,omitempty"`
	// The unit of time requested.
	Units NullableUnitsEnum `json:"units,omitempty"`
	// The number of time off units requested.
	Amount NullableFloat32 `json:"amount,omitempty"`
	// The type of time off request.
	RequestType NullableRequestTypeEnum `json:"request_type,omitempty"`
}

// NewTimeOff instantiates a new TimeOff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOff() *TimeOff {
	this := TimeOff{}
	return &this
}

// NewTimeOffWithDefaults instantiates a new TimeOff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOffWithDefaults() *TimeOff {
	this := TimeOff{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimeOff) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeOff) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimeOff) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimeOff) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *TimeOff) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *TimeOff) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *TimeOff) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *TimeOff) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetEmployee returns the Employee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetEmployee() string {
	if o == nil || o.Employee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Employee.Get()
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employee.Get(), o.Employee.IsSet()
}

// HasEmployee returns a boolean if a field has been set.
func (o *TimeOff) HasEmployee() bool {
	if o != nil && o.Employee.IsSet() {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given NullableString and assigns it to the Employee field.
func (o *TimeOff) SetEmployee(v string) {
	o.Employee.Set(&v)
}
// SetEmployeeNil sets the value for Employee to be an explicit nil
func (o *TimeOff) SetEmployeeNil() {
	o.Employee.Set(nil)
}

// UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
func (o *TimeOff) UnsetEmployee() {
	o.Employee.Unset()
}

// GetApprover returns the Approver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetApprover() string {
	if o == nil || o.Approver.Get() == nil {
		var ret string
		return ret
	}
	return *o.Approver.Get()
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetApproverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Approver.Get(), o.Approver.IsSet()
}

// HasApprover returns a boolean if a field has been set.
func (o *TimeOff) HasApprover() bool {
	if o != nil && o.Approver.IsSet() {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NullableString and assigns it to the Approver field.
func (o *TimeOff) SetApprover(v string) {
	o.Approver.Set(&v)
}
// SetApproverNil sets the value for Approver to be an explicit nil
func (o *TimeOff) SetApproverNil() {
	o.Approver.Set(nil)
}

// UnsetApprover ensures that no value is present for Approver, not even an explicit nil
func (o *TimeOff) UnsetApprover() {
	o.Approver.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetStatus() TimeOffStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret TimeOffStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetStatusOk() (*TimeOffStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TimeOff) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableTimeOffStatusEnum and assigns it to the Status field.
func (o *TimeOff) SetStatus(v TimeOffStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TimeOff) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TimeOff) UnsetStatus() {
	o.Status.Unset()
}

// GetEmployeeNote returns the EmployeeNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetEmployeeNote() string {
	if o == nil || o.EmployeeNote.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeeNote.Get()
}

// GetEmployeeNoteOk returns a tuple with the EmployeeNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetEmployeeNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeNote.Get(), o.EmployeeNote.IsSet()
}

// HasEmployeeNote returns a boolean if a field has been set.
func (o *TimeOff) HasEmployeeNote() bool {
	if o != nil && o.EmployeeNote.IsSet() {
		return true
	}

	return false
}

// SetEmployeeNote gets a reference to the given NullableString and assigns it to the EmployeeNote field.
func (o *TimeOff) SetEmployeeNote(v string) {
	o.EmployeeNote.Set(&v)
}
// SetEmployeeNoteNil sets the value for EmployeeNote to be an explicit nil
func (o *TimeOff) SetEmployeeNoteNil() {
	o.EmployeeNote.Set(nil)
}

// UnsetEmployeeNote ensures that no value is present for EmployeeNote, not even an explicit nil
func (o *TimeOff) UnsetEmployeeNote() {
	o.EmployeeNote.Unset()
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetUnits() UnitsEnum {
	if o == nil || o.Units.Get() == nil {
		var ret UnitsEnum
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetUnitsOk() (*UnitsEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *TimeOff) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableUnitsEnum and assigns it to the Units field.
func (o *TimeOff) SetUnits(v UnitsEnum) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *TimeOff) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *TimeOff) UnsetUnits() {
	o.Units.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetAmount() float32 {
	if o == nil || o.Amount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *TimeOff) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *TimeOff) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *TimeOff) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *TimeOff) UnsetAmount() {
	o.Amount.Unset()
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOff) GetRequestType() RequestTypeEnum {
	if o == nil || o.RequestType.Get() == nil {
		var ret RequestTypeEnum
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOff) GetRequestTypeOk() (*RequestTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *TimeOff) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableRequestTypeEnum and assigns it to the RequestType field.
func (o *TimeOff) SetRequestType(v RequestTypeEnum) {
	o.RequestType.Set(&v)
}
// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *TimeOff) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *TimeOff) UnsetRequestType() {
	o.RequestType.Unset()
}

func (o TimeOff) MarshalJSON() ([]byte, error) {
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
	if o.Approver.IsSet() {
		toSerialize["approver"] = o.Approver.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.EmployeeNote.IsSet() {
		toSerialize["employee_note"] = o.EmployeeNote.Get()
	}
	if o.Units.IsSet() {
		toSerialize["units"] = o.Units.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.RequestType.IsSet() {
		toSerialize["request_type"] = o.RequestType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTimeOff struct {
	value *TimeOff
	isSet bool
}

func (v NullableTimeOff) Get() *TimeOff {
	return v.value
}

func (v *NullableTimeOff) Set(val *TimeOff) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOff) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOff(val *TimeOff) *NullableTimeOff {
	return &NullableTimeOff{value: val, isSet: true}
}

func (v NullableTimeOff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

