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
	"fmt"
)

// MethodEnum the model 'MethodEnum'
type MethodEnum string

// List of MethodEnum
const (
	METHODENUM_GET MethodEnum = "GET"
	METHODENUM_OPTIONS MethodEnum = "OPTIONS"
	METHODENUM_HEAD MethodEnum = "HEAD"
	METHODENUM_POST MethodEnum = "POST"
	METHODENUM_PUT MethodEnum = "PUT"
	METHODENUM_PATCH MethodEnum = "PATCH"
	METHODENUM_DELETE MethodEnum = "DELETE"
)

var allowedMethodEnumEnumValues = []MethodEnum{
	"GET",
	"OPTIONS",
	"HEAD",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
}

func (v *MethodEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MethodEnum(value)
	for _, existing := range allowedMethodEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MethodEnum", value)
}

// NewMethodEnumFromValue returns a pointer to a valid MethodEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMethodEnumFromValue(v string) (*MethodEnum, error) {
	ev := MethodEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MethodEnum: valid values are %v", v, allowedMethodEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MethodEnum) IsValid() bool {
	for _, existing := range allowedMethodEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MethodEnum value
func (v MethodEnum) Ptr() *MethodEnum {
	return &v
}

type NullableMethodEnum struct {
	value *MethodEnum
	isSet bool
}

func (v NullableMethodEnum) Get() *MethodEnum {
	return v.value
}

func (v *NullableMethodEnum) Set(val *MethodEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodEnum(val *MethodEnum) *NullableMethodEnum {
	return &NullableMethodEnum{value: val, isSet: true}
}

func (v NullableMethodEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

