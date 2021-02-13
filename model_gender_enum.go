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

// GenderEnum the model 'GenderEnum'
type GenderEnum string

// List of GenderEnum
const (
	GENDERENUM_MALE GenderEnum = "MALE"
	GENDERENUM_FEMALE GenderEnum = "FEMALE"
	GENDERENUM_NON_BINARY GenderEnum = "NON-BINARY"
	GENDERENUM_OTHER GenderEnum = "OTHER"
	GENDERENUM_PREFER_NOT_TO_DISCLOSE GenderEnum = "PREFER_NOT_TO_DISCLOSE"
)

func (v *GenderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GenderEnum(value)
	for _, existing := range []GenderEnum{ "MALE", "FEMALE", "NON-BINARY", "OTHER", "PREFER_NOT_TO_DISCLOSE",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GenderEnum", value)
}

// Ptr returns reference to GenderEnum value
func (v GenderEnum) Ptr() *GenderEnum {
	return &v
}

type NullableGenderEnum struct {
	value *GenderEnum
	isSet bool
}

func (v NullableGenderEnum) Get() *GenderEnum {
	return v.value
}

func (v *NullableGenderEnum) Set(val *GenderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGenderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGenderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenderEnum(val *GenderEnum) *NullableGenderEnum {
	return &NullableGenderEnum{value: val, isSet: true}
}

func (v NullableGenderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

