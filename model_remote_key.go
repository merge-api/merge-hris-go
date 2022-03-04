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

// RemoteKey # The RemoteKey Object ### Description The `RemoteKey` object is used to represent a request for a new remote key.  ### Usage Example Post a `GenerateRemoteKey` to receive a new `RemoteKey`.
type RemoteKey struct {
	Name string `json:"name"`
	Key string `json:"key"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewRemoteKey instantiates a new RemoteKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteKey(name string, key string) *RemoteKey {
	this := RemoteKey{}
	this.Name = name
	this.Key = key
	return &this
}

// NewRemoteKeyWithDefaults instantiates a new RemoteKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteKeyWithDefaults() *RemoteKey {
	this := RemoteKey{}
	return &this
}

// GetName returns the Name field value
func (o *RemoteKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RemoteKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RemoteKey) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *RemoteKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RemoteKey) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RemoteKey) SetKey(v string) {
	o.Key = v
}

func (o RemoteKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

func (v *RemoteKey) UnmarshalJSON(src []byte) error {
    type RemoteKeyUnmarshalTarget RemoteKey

	var intermediateResult RemoteKeyUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = RemoteKey(intermediateResult)
	return nil
}
type NullableRemoteKey struct {
	value *RemoteKey
	isSet bool
}

func (v NullableRemoteKey) Get() *RemoteKey {
	return v.value
}

func (v *NullableRemoteKey) Set(val *RemoteKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteKey(val *RemoteKey) *NullableRemoteKey {
	return &NullableRemoteKey{value: val, isSet: true}
}

func (v NullableRemoteKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


