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

// RunTypeEnum the model 'RunTypeEnum'
type RunTypeEnum string

// List of RunTypeEnum
const (
	RUNTYPEENUM_REGULAR RunTypeEnum = "REGULAR"
	RUNTYPEENUM_OFF_CYCLE RunTypeEnum = "OFF_CYCLE"
	RUNTYPEENUM_CORRECTION RunTypeEnum = "CORRECTION"
	RUNTYPEENUM_TERMINATION RunTypeEnum = "TERMINATION"
	RUNTYPEENUM_SIGN_ON_BONUS RunTypeEnum = "SIGN_ON_BONUS"
)

var allowedRunTypeEnumEnumValues = []RunTypeEnum{
	"REGULAR",
	"OFF_CYCLE",
	"CORRECTION",
	"TERMINATION",
	"SIGN_ON_BONUS",
}

func (v *RunTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RunTypeEnum(value)
	for _, existing := range allowedRunTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RunTypeEnum", value)
}

// NewRunTypeEnumFromValue returns a pointer to a valid RunTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRunTypeEnumFromValue(v string) (*RunTypeEnum, error) {
	ev := RunTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RunTypeEnum: valid values are %v", v, allowedRunTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RunTypeEnum) IsValid() bool {
	for _, existing := range allowedRunTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunTypeEnum value
func (v RunTypeEnum) Ptr() *RunTypeEnum {
	return &v
}

type NullableRunTypeEnum struct {
	value *RunTypeEnum
	isSet bool
}

func (v NullableRunTypeEnum) Get() *RunTypeEnum {
	return v.value
}

func (v *NullableRunTypeEnum) Set(val *RunTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTypeEnum(val *RunTypeEnum) *NullableRunTypeEnum {
	return &NullableRunTypeEnum{value: val, isSet: true}
}

func (v NullableRunTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

