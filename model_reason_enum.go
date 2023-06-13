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

// ReasonEnum * `GENERAL_CUSTOMER_REQUEST` - GENERAL_CUSTOMER_REQUEST * `GDPR` - GDPR * `OTHER` - OTHER
type ReasonEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of ReasonEnum
const (
    REASONENUM_MERGE_NONSTANDARD_VALUE ReasonEnum = "MERGE_NONSTANDARD_VALUE"
    
	REASONENUM_GENERAL_CUSTOMER_REQUEST ReasonEnum = "GENERAL_CUSTOMER_REQUEST"
	REASONENUM_GDPR ReasonEnum = "GDPR"
	REASONENUM_OTHER ReasonEnum = "OTHER"
)

var allowedReasonEnumEnumValues = []ReasonEnum{
	"GENERAL_CUSTOMER_REQUEST",
	"GDPR",
	"OTHER",
}

func (v *ReasonEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReasonEnum(value)
	for _, existing := range allowedReasonEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = REASONENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewReasonEnumFromValue returns a pointer to a valid ReasonEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReasonEnumFromValue(v string) (*ReasonEnum, error) {
	ev := ReasonEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := REASONENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReasonEnum) IsValid() bool {
	for _, existing := range allowedReasonEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReasonEnum value
func (v ReasonEnum) Ptr() *ReasonEnum {
	return &v
}

type NullableReasonEnum struct {
	value *ReasonEnum
	isSet bool
}

func (v NullableReasonEnum) Get() *ReasonEnum {
	return v.value
}

func (v *NullableReasonEnum) Set(val *ReasonEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonEnum(val *ReasonEnum) *NullableReasonEnum {
	return &NullableReasonEnum{value: val, isSet: true}
}

func (v NullableReasonEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

