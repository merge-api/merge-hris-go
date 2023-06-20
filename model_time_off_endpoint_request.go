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

// TimeOffEndpointRequest struct for TimeOffEndpointRequest
type TimeOffEndpointRequest struct {
	Model TimeOffRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTimeOffEndpointRequest instantiates a new TimeOffEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOffEndpointRequest(model TimeOffRequest) *TimeOffEndpointRequest {
	this := TimeOffEndpointRequest{}
	this.Model = model
	return &this
}

// NewTimeOffEndpointRequestWithDefaults instantiates a new TimeOffEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOffEndpointRequestWithDefaults() *TimeOffEndpointRequest {
	this := TimeOffEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *TimeOffEndpointRequest) GetModel() TimeOffRequest {
	if o == nil {
		var ret TimeOffRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *TimeOffEndpointRequest) GetModelOk() (*TimeOffRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *TimeOffEndpointRequest) SetModel(v TimeOffRequest) {
	o.Model = v
}

func (o TimeOffEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *TimeOffEndpointRequest) UnmarshalJSON(src []byte) error {
    type TimeOffEndpointRequestUnmarshalTarget TimeOffEndpointRequest

	var intermediateResult TimeOffEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = TimeOffEndpointRequest(intermediateResult)
	return nil
}
type NullableTimeOffEndpointRequest struct {
	value *TimeOffEndpointRequest
	isSet bool
}

func (v NullableTimeOffEndpointRequest) Get() *TimeOffEndpointRequest {
	return v.value
}

func (v *NullableTimeOffEndpointRequest) Set(val *TimeOffEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOffEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOffEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOffEndpointRequest(val *TimeOffEndpointRequest) *NullableTimeOffEndpointRequest {
	return &NullableTimeOffEndpointRequest{value: val, isSet: true}
}

func (v NullableTimeOffEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOffEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


