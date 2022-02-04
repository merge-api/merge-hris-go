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

// EmploymentTypeEnum the model 'EmploymentTypeEnum'
type EmploymentTypeEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of EmploymentTypeEnum
const (
    EMPLOYMENTTYPEENUM_MERGE_NONSTANDARD_VALUE EmploymentTypeEnum = "MERGE_NONSTANDARD_VALUE"
    
	EMPLOYMENTTYPEENUM_FULL_TIME EmploymentTypeEnum = "FULL_TIME"
	EMPLOYMENTTYPEENUM_PART_TIME EmploymentTypeEnum = "PART_TIME"
	EMPLOYMENTTYPEENUM_INTERN EmploymentTypeEnum = "INTERN"
	EMPLOYMENTTYPEENUM_CONTRACTOR EmploymentTypeEnum = "CONTRACTOR"
	EMPLOYMENTTYPEENUM_FREELANCE EmploymentTypeEnum = "FREELANCE"
)

var allowedEmploymentTypeEnumEnumValues = []EmploymentTypeEnum{
	"FULL_TIME",
	"PART_TIME",
	"INTERN",
	"CONTRACTOR",
	"FREELANCE",
}

func (v *EmploymentTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmploymentTypeEnum(value)
	for _, existing := range allowedEmploymentTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = EMPLOYMENTTYPEENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewEmploymentTypeEnumFromValue returns a pointer to a valid EmploymentTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmploymentTypeEnumFromValue(v string) (*EmploymentTypeEnum, error) {
	ev := EmploymentTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := EMPLOYMENTTYPEENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmploymentTypeEnum) IsValid() bool {
	for _, existing := range allowedEmploymentTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmploymentTypeEnum value
func (v EmploymentTypeEnum) Ptr() *EmploymentTypeEnum {
	return &v
}

type NullableEmploymentTypeEnum struct {
	value *EmploymentTypeEnum
	isSet bool
}

func (v NullableEmploymentTypeEnum) Get() *EmploymentTypeEnum {
	return v.value
}

func (v *NullableEmploymentTypeEnum) Set(val *EmploymentTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentTypeEnum(val *EmploymentTypeEnum) *NullableEmploymentTypeEnum {
	return &NullableEmploymentTypeEnum{value: val, isSet: true}
}

func (v NullableEmploymentTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

