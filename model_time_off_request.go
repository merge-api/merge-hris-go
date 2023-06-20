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
	"time"
)

// TimeOffRequest # The TimeOff Object ### Description The `TimeOff` object is used to represent all employees' Time Off entries.  ### Usage Example Fetch from the `LIST TimeOffs` endpoint and filter by `ID` to show all time off requests.
type TimeOffRequest struct {
	// The employee requesting time off.
	Employee NullableString `json:"employee,omitempty"`
	// The Merge ID of the employee with the ability to approve the time off request.
	Approver NullableString `json:"approver,omitempty"`
	// The status of this time off request.  * `REQUESTED` - REQUESTED * `APPROVED` - APPROVED * `DECLINED` - DECLINED * `CANCELLED` - CANCELLED * `DELETED` - DELETED
	Status NullableTimeOffStatusEnum `json:"status,omitempty"`
	// The employee note for this time off request.
	EmployeeNote NullableString `json:"employee_note,omitempty"`
	// The measurement that the third-party integration uses to count time requested.  * `HOURS` - HOURS * `DAYS` - DAYS
	Units NullableUnitsEnum `json:"units,omitempty"`
	// The time off quantity measured by the prescribed “units”.
	Amount NullableFloat64 `json:"amount,omitempty"`
	// The type of time off request.  * `VACATION` - VACATION * `SICK` - SICK * `PERSONAL` - PERSONAL * `JURY_DUTY` - JURY_DUTY * `VOLUNTEER` - VOLUNTEER * `BEREAVEMENT` - BEREAVEMENT
	RequestType NullableRequestTypeEnum `json:"request_type,omitempty"`
	// The day and time of the start of the time requested off.
	StartTime NullableTime `json:"start_time,omitempty"`
	// The day and time of the end of the time requested off.
	EndTime NullableTime `json:"end_time,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTimeOffRequest instantiates a new TimeOffRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOffRequest() *TimeOffRequest {
	this := TimeOffRequest{}
	return &this
}

// NewTimeOffRequestWithDefaults instantiates a new TimeOffRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOffRequestWithDefaults() *TimeOffRequest {
	this := TimeOffRequest{}
	return &this
}

// GetEmployee returns the Employee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetEmployee() string {
	if o == nil || o.Employee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Employee.Get()
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employee.Get(), o.Employee.IsSet()
}

// HasEmployee returns a boolean if a field has been set.
func (o *TimeOffRequest) HasEmployee() bool {
	if o != nil && o.Employee.IsSet() {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given NullableString and assigns it to the Employee field.
func (o *TimeOffRequest) SetEmployee(v string) {
	o.Employee.Set(&v)
}
// SetEmployeeNil sets the value for Employee to be an explicit nil
func (o *TimeOffRequest) SetEmployeeNil() {
	o.Employee.Set(nil)
}

// UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
func (o *TimeOffRequest) UnsetEmployee() {
	o.Employee.Unset()
}

// GetApprover returns the Approver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetApprover() string {
	if o == nil || o.Approver.Get() == nil {
		var ret string
		return ret
	}
	return *o.Approver.Get()
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetApproverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Approver.Get(), o.Approver.IsSet()
}

// HasApprover returns a boolean if a field has been set.
func (o *TimeOffRequest) HasApprover() bool {
	if o != nil && o.Approver.IsSet() {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NullableString and assigns it to the Approver field.
func (o *TimeOffRequest) SetApprover(v string) {
	o.Approver.Set(&v)
}
// SetApproverNil sets the value for Approver to be an explicit nil
func (o *TimeOffRequest) SetApproverNil() {
	o.Approver.Set(nil)
}

// UnsetApprover ensures that no value is present for Approver, not even an explicit nil
func (o *TimeOffRequest) UnsetApprover() {
	o.Approver.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetStatus() TimeOffStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret TimeOffStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetStatusOk() (*TimeOffStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TimeOffRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableTimeOffStatusEnum and assigns it to the Status field.
func (o *TimeOffRequest) SetStatus(v TimeOffStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TimeOffRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TimeOffRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetEmployeeNote returns the EmployeeNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetEmployeeNote() string {
	if o == nil || o.EmployeeNote.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeeNote.Get()
}

// GetEmployeeNoteOk returns a tuple with the EmployeeNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetEmployeeNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeNote.Get(), o.EmployeeNote.IsSet()
}

// HasEmployeeNote returns a boolean if a field has been set.
func (o *TimeOffRequest) HasEmployeeNote() bool {
	if o != nil && o.EmployeeNote.IsSet() {
		return true
	}

	return false
}

// SetEmployeeNote gets a reference to the given NullableString and assigns it to the EmployeeNote field.
func (o *TimeOffRequest) SetEmployeeNote(v string) {
	o.EmployeeNote.Set(&v)
}
// SetEmployeeNoteNil sets the value for EmployeeNote to be an explicit nil
func (o *TimeOffRequest) SetEmployeeNoteNil() {
	o.EmployeeNote.Set(nil)
}

// UnsetEmployeeNote ensures that no value is present for EmployeeNote, not even an explicit nil
func (o *TimeOffRequest) UnsetEmployeeNote() {
	o.EmployeeNote.Unset()
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetUnits() UnitsEnum {
	if o == nil || o.Units.Get() == nil {
		var ret UnitsEnum
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetUnitsOk() (*UnitsEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *TimeOffRequest) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableUnitsEnum and assigns it to the Units field.
func (o *TimeOffRequest) SetUnits(v UnitsEnum) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *TimeOffRequest) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *TimeOffRequest) UnsetUnits() {
	o.Units.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetAmount() float64 {
	if o == nil || o.Amount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *TimeOffRequest) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *TimeOffRequest) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *TimeOffRequest) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *TimeOffRequest) UnsetAmount() {
	o.Amount.Unset()
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetRequestType() RequestTypeEnum {
	if o == nil || o.RequestType.Get() == nil {
		var ret RequestTypeEnum
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetRequestTypeOk() (*RequestTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *TimeOffRequest) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableRequestTypeEnum and assigns it to the RequestType field.
func (o *TimeOffRequest) SetRequestType(v RequestTypeEnum) {
	o.RequestType.Set(&v)
}
// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *TimeOffRequest) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *TimeOffRequest) UnsetRequestType() {
	o.RequestType.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeOffRequest) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *TimeOffRequest) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *TimeOffRequest) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *TimeOffRequest) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeOffRequest) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.
func (o *TimeOffRequest) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *TimeOffRequest) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *TimeOffRequest) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *TimeOffRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *TimeOffRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *TimeOffRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *TimeOffRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o TimeOffRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.EndTime.IsSet() {
		toSerialize["end_time"] = o.EndTime.Get()
	}
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *TimeOffRequest) UnmarshalJSON(src []byte) error {
    type TimeOffRequestUnmarshalTarget TimeOffRequest

	var intermediateResult TimeOffRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = TimeOffRequest(intermediateResult)
	return nil
}
type NullableTimeOffRequest struct {
	value *TimeOffRequest
	isSet bool
}

func (v NullableTimeOffRequest) Get() *TimeOffRequest {
	return v.value
}

func (v *NullableTimeOffRequest) Set(val *TimeOffRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOffRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOffRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOffRequest(val *TimeOffRequest) *NullableTimeOffRequest {
	return &NullableTimeOffRequest{value: val, isSet: true}
}

func (v NullableTimeOffRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOffRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


