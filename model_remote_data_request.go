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

// RemoteDataRequest struct for RemoteDataRequest
type RemoteDataRequest struct {
	Path string `json:"path"`
	Data *map[string]interface{} `json:"data,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewRemoteDataRequest instantiates a new RemoteDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteDataRequest(path string) *RemoteDataRequest {
	this := RemoteDataRequest{}
	this.Path = path
	return &this
}

// NewRemoteDataRequestWithDefaults instantiates a new RemoteDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteDataRequestWithDefaults() *RemoteDataRequest {
	this := RemoteDataRequest{}
	return &this
}

// GetPath returns the Path field value
func (o *RemoteDataRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RemoteDataRequest) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RemoteDataRequest) SetPath(v string) {
	o.Path = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoteDataRequest) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteDataRequest) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoteDataRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *RemoteDataRequest) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o RemoteDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (v *RemoteDataRequest) UnmarshalJSON(src []byte) error {
    type RemoteDataRequestUnmarshalTarget RemoteDataRequest

	var intermediateResult RemoteDataRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = RemoteDataRequest(intermediateResult)
	return nil
}
type NullableRemoteDataRequest struct {
	value *RemoteDataRequest
	isSet bool
}

func (v NullableRemoteDataRequest) Get() *RemoteDataRequest {
	return v.value
}

func (v *NullableRemoteDataRequest) Set(val *RemoteDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteDataRequest(val *RemoteDataRequest) *NullableRemoteDataRequest {
	return &NullableRemoteDataRequest{value: val, isSet: true}
}

func (v NullableRemoteDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


