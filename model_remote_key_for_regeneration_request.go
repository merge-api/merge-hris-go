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

// RemoteKeyForRegenerationRequest # The RemoteKeyForRegeneration Object ### Description The `RemoteKeyForRegeneration` object is used to exchange an old remote key for a new one  ### Usage Example Post a `RemoteKeyForRegeneration` to swap out an old remote key for a new one
type RemoteKeyForRegenerationRequest struct {
	Name string `json:"name"`
}

// NewRemoteKeyForRegenerationRequest instantiates a new RemoteKeyForRegenerationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteKeyForRegenerationRequest(name string, ) *RemoteKeyForRegenerationRequest {
	this := RemoteKeyForRegenerationRequest{}
	this.Name = name
	return &this
}

// NewRemoteKeyForRegenerationRequestWithDefaults instantiates a new RemoteKeyForRegenerationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteKeyForRegenerationRequestWithDefaults() *RemoteKeyForRegenerationRequest {
	this := RemoteKeyForRegenerationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RemoteKeyForRegenerationRequest) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RemoteKeyForRegenerationRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RemoteKeyForRegenerationRequest) SetName(v string) {
	o.Name = v
}

func (o RemoteKeyForRegenerationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteKeyForRegenerationRequest struct {
	value *RemoteKeyForRegenerationRequest
	isSet bool
}

func (v NullableRemoteKeyForRegenerationRequest) Get() *RemoteKeyForRegenerationRequest {
	return v.value
}

func (v *NullableRemoteKeyForRegenerationRequest) Set(val *RemoteKeyForRegenerationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteKeyForRegenerationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteKeyForRegenerationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteKeyForRegenerationRequest(val *RemoteKeyForRegenerationRequest) *NullableRemoteKeyForRegenerationRequest {
	return &NullableRemoteKeyForRegenerationRequest{value: val, isSet: true}
}

func (v NullableRemoteKeyForRegenerationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteKeyForRegenerationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


