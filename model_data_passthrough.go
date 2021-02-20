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

// DataPassthrough # The DataPassthrough Object ### Description The `DataPassthrough` object is used to send information to an otherwise-unsupported third-party endpoint.  ### Usage Example Create a `DataPassthrough` to get team hierarchies from your Rippling integration.
type DataPassthrough struct {
	Method MethodEnum `json:"method"`
	Path string `json:"path"`
	Data *map[string]interface{} `json:"data,omitempty"`
	Headers *map[string]interface{} `json:"headers,omitempty"`
}

// NewDataPassthrough instantiates a new DataPassthrough object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPassthrough(method MethodEnum, path string, ) *DataPassthrough {
	this := DataPassthrough{}
	this.Method = method
	this.Path = path
	return &this
}

// NewDataPassthroughWithDefaults instantiates a new DataPassthrough object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPassthroughWithDefaults() *DataPassthrough {
	this := DataPassthrough{}
	return &this
}

// GetMethod returns the Method field value
func (o *DataPassthrough) GetMethod() MethodEnum {
	if o == nil  {
		var ret MethodEnum
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *DataPassthrough) GetMethodOk() (*MethodEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *DataPassthrough) SetMethod(v MethodEnum) {
	o.Method = v
}

// GetPath returns the Path field value
func (o *DataPassthrough) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DataPassthrough) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DataPassthrough) SetPath(v string) {
	o.Path = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DataPassthrough) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPassthrough) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DataPassthrough) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *DataPassthrough) SetData(v map[string]interface{}) {
	o.Data = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *DataPassthrough) GetHeaders() map[string]interface{} {
	if o == nil || o.Headers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPassthrough) GetHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *DataPassthrough) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *DataPassthrough) SetHeaders(v map[string]interface{}) {
	o.Headers = &v
}

func (o DataPassthrough) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	return json.Marshal(toSerialize)
}

type NullableDataPassthrough struct {
	value *DataPassthrough
	isSet bool
}

func (v NullableDataPassthrough) Get() *DataPassthrough {
	return v.value
}

func (v *NullableDataPassthrough) Set(val *DataPassthrough) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPassthrough) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPassthrough) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPassthrough(val *DataPassthrough) *NullableDataPassthrough {
	return &NullableDataPassthrough{value: val, isSet: true}
}

func (v NullableDataPassthrough) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPassthrough) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


