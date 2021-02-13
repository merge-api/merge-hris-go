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

// PayPeriodEnum the model 'PayPeriodEnum'
type PayPeriodEnum string

// List of PayPeriodEnum
const (
	PAYPERIODENUM_HOUR PayPeriodEnum = "HOUR"
	PAYPERIODENUM_DAY PayPeriodEnum = "DAY"
	PAYPERIODENUM_WEEK PayPeriodEnum = "WEEK"
	PAYPERIODENUM_EVERY_TWO_WEEKS PayPeriodEnum = "EVERY_TWO_WEEKS"
	PAYPERIODENUM_MONTH PayPeriodEnum = "MONTH"
	PAYPERIODENUM_QUARTER PayPeriodEnum = "QUARTER"
	PAYPERIODENUM_EVERY_SIX_MONTHS PayPeriodEnum = "EVERY_SIX_MONTHS"
	PAYPERIODENUM_YEAR PayPeriodEnum = "YEAR"
)

func (v *PayPeriodEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PayPeriodEnum(value)
	for _, existing := range []PayPeriodEnum{ "HOUR", "DAY", "WEEK", "EVERY_TWO_WEEKS", "MONTH", "QUARTER", "EVERY_SIX_MONTHS", "YEAR",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PayPeriodEnum", value)
}

// Ptr returns reference to PayPeriodEnum value
func (v PayPeriodEnum) Ptr() *PayPeriodEnum {
	return &v
}

type NullablePayPeriodEnum struct {
	value *PayPeriodEnum
	isSet bool
}

func (v NullablePayPeriodEnum) Get() *PayPeriodEnum {
	return v.value
}

func (v *NullablePayPeriodEnum) Set(val *PayPeriodEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePayPeriodEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePayPeriodEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayPeriodEnum(val *PayPeriodEnum) *NullablePayPeriodEnum {
	return &NullablePayPeriodEnum{value: val, isSet: true}
}

func (v NullablePayPeriodEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayPeriodEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

