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

// MaritalStatusEnum the model 'MaritalStatusEnum'
type MaritalStatusEnum string

// List of MaritalStatusEnum
const (
	SINGLE MaritalStatusEnum = "SINGLE"
	MARRIED_FILING_JOINTLY MaritalStatusEnum = "MARRIED_FILING_JOINTLY"
	MARRIED_FILING_SEPARATELY MaritalStatusEnum = "MARRIED_FILING_SEPARATELY"
	HEAD_OF_HOUSEHOLD MaritalStatusEnum = "HEAD_OF_HOUSEHOLD"
	QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD MaritalStatusEnum = "QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD"
)

func (v *MaritalStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MaritalStatusEnum(value)
	for _, existing := range []MaritalStatusEnum{ "SINGLE", "MARRIED_FILING_JOINTLY", "MARRIED_FILING_SEPARATELY", "HEAD_OF_HOUSEHOLD", "QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MaritalStatusEnum", value)
}

// Ptr returns reference to MaritalStatusEnum value
func (v MaritalStatusEnum) Ptr() *MaritalStatusEnum {
	return &v
}

type NullableMaritalStatusEnum struct {
	value *MaritalStatusEnum
	isSet bool
}

func (v NullableMaritalStatusEnum) Get() *MaritalStatusEnum {
	return v.value
}

func (v *NullableMaritalStatusEnum) Set(val *MaritalStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMaritalStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMaritalStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaritalStatusEnum(val *MaritalStatusEnum) *NullableMaritalStatusEnum {
	return &NullableMaritalStatusEnum{value: val, isSet: true}
}

func (v NullableMaritalStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaritalStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
