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

// TypeEnum the model 'TypeEnum'
type TypeEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of TypeEnum
const (
    TYPEENUM_MERGE_NONSTANDARD_VALUE TypeEnum = "MERGE_NONSTANDARD_VALUE"
    
	TYPEENUM_SALARY TypeEnum = "SALARY"
	TYPEENUM_REIMBURSEMENT TypeEnum = "REIMBURSEMENT"
	TYPEENUM_OVERTIME TypeEnum = "OVERTIME"
	TYPEENUM_BONUS TypeEnum = "BONUS"
)

var allowedTypeEnumEnumValues = []TypeEnum{
	"SALARY",
	"REIMBURSEMENT",
	"OVERTIME",
	"BONUS",
}

func (v *TypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeEnum(value)
	for _, existing := range allowedTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TYPEENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewTypeEnumFromValue returns a pointer to a valid TypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeEnumFromValue(v string) (*TypeEnum, error) {
	ev := TypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := TYPEENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeEnum) IsValid() bool {
	for _, existing := range allowedTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeEnum value
func (v TypeEnum) Ptr() *TypeEnum {
	return &v
}

type NullableTypeEnum struct {
	value *TypeEnum
	isSet bool
}

func (v NullableTypeEnum) Get() *TypeEnum {
	return v.value
}

func (v *NullableTypeEnum) Set(val *TypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeEnum(val *TypeEnum) *NullableTypeEnum {
	return &NullableTypeEnum{value: val, isSet: true}
}

func (v NullableTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

