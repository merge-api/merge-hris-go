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

// PayFrequencyEnum the model 'PayFrequencyEnum'
type PayFrequencyEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of PayFrequencyEnum
const (
    PAYFREQUENCYENUM_MERGE_NONSTANDARD_VALUE PayFrequencyEnum = "MERGE_NONSTANDARD_VALUE"
    
	PAYFREQUENCYENUM_WEEKLY PayFrequencyEnum = "WEEKLY"
	PAYFREQUENCYENUM_BIWEEKLY PayFrequencyEnum = "BIWEEKLY"
	PAYFREQUENCYENUM_MONTHLY PayFrequencyEnum = "MONTHLY"
	PAYFREQUENCYENUM_QUARTERLY PayFrequencyEnum = "QUARTERLY"
	PAYFREQUENCYENUM_SEMIANNUALLY PayFrequencyEnum = "SEMIANNUALLY"
	PAYFREQUENCYENUM_ANNUALLY PayFrequencyEnum = "ANNUALLY"
	PAYFREQUENCYENUM_THIRTEEN_MONTHLY PayFrequencyEnum = "THIRTEEN-MONTHLY"
	PAYFREQUENCYENUM_PRO_RATA PayFrequencyEnum = "PRO_RATA"
)

var allowedPayFrequencyEnumEnumValues = []PayFrequencyEnum{
	"WEEKLY",
	"BIWEEKLY",
	"MONTHLY",
	"QUARTERLY",
	"SEMIANNUALLY",
	"ANNUALLY",
	"THIRTEEN-MONTHLY",
	"PRO_RATA",
}

func (v *PayFrequencyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PayFrequencyEnum(value)
	for _, existing := range allowedPayFrequencyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = PAYFREQUENCYENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewPayFrequencyEnumFromValue returns a pointer to a valid PayFrequencyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPayFrequencyEnumFromValue(v string) (*PayFrequencyEnum, error) {
	ev := PayFrequencyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := PAYFREQUENCYENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PayFrequencyEnum) IsValid() bool {
	for _, existing := range allowedPayFrequencyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PayFrequencyEnum value
func (v PayFrequencyEnum) Ptr() *PayFrequencyEnum {
	return &v
}

type NullablePayFrequencyEnum struct {
	value *PayFrequencyEnum
	isSet bool
}

func (v NullablePayFrequencyEnum) Get() *PayFrequencyEnum {
	return v.value
}

func (v *NullablePayFrequencyEnum) Set(val *PayFrequencyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePayFrequencyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePayFrequencyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayFrequencyEnum(val *PayFrequencyEnum) *NullablePayFrequencyEnum {
	return &NullablePayFrequencyEnum{value: val, isSet: true}
}

func (v NullablePayFrequencyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayFrequencyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

