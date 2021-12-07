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

// IssueStatusEnum the model 'IssueStatusEnum'
type IssueStatusEnum string

// List of IssueStatusEnum
const (
	ISSUESTATUSENUM_ONGOING IssueStatusEnum = "ONGOING"
	ISSUESTATUSENUM_RESOLVED IssueStatusEnum = "RESOLVED"
)

var allowedIssueStatusEnumEnumValues = []IssueStatusEnum{
	"ONGOING",
	"RESOLVED",
}

func (v *IssueStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssueStatusEnum(value)
	for _, existing := range allowedIssueStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssueStatusEnum", value)
}

// NewIssueStatusEnumFromValue returns a pointer to a valid IssueStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssueStatusEnumFromValue(v string) (*IssueStatusEnum, error) {
	ev := IssueStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssueStatusEnum: valid values are %v", v, allowedIssueStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssueStatusEnum) IsValid() bool {
	for _, existing := range allowedIssueStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStatusEnum value
func (v IssueStatusEnum) Ptr() *IssueStatusEnum {
	return &v
}

type NullableIssueStatusEnum struct {
	value *IssueStatusEnum
	isSet bool
}

func (v NullableIssueStatusEnum) Get() *IssueStatusEnum {
	return v.value
}

func (v *NullableIssueStatusEnum) Set(val *IssueStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueStatusEnum(val *IssueStatusEnum) *NullableIssueStatusEnum {
	return &NullableIssueStatusEnum{value: val, isSet: true}
}

func (v NullableIssueStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

