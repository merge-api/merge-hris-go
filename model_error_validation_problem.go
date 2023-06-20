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
)

// ErrorValidationProblem struct for ErrorValidationProblem
type ErrorValidationProblem struct {
	Source *ValidationProblemSource `json:"source,omitempty"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	ProblemType string `json:"problem_type"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewErrorValidationProblem instantiates a new ErrorValidationProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorValidationProblem(title string, detail string, problemType string) *ErrorValidationProblem {
	this := ErrorValidationProblem{}
	this.Title = title
	this.Detail = detail
	this.ProblemType = problemType
	return &this
}

// NewErrorValidationProblemWithDefaults instantiates a new ErrorValidationProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorValidationProblemWithDefaults() *ErrorValidationProblem {
	this := ErrorValidationProblem{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ErrorValidationProblem) GetSource() ValidationProblemSource {
	if o == nil || o.Source == nil {
		var ret ValidationProblemSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorValidationProblem) GetSourceOk() (*ValidationProblemSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ErrorValidationProblem) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given ValidationProblemSource and assigns it to the Source field.
func (o *ErrorValidationProblem) SetSource(v ValidationProblemSource) {
	o.Source = &v
}

// GetTitle returns the Title field value
func (o *ErrorValidationProblem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ErrorValidationProblem) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ErrorValidationProblem) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value
func (o *ErrorValidationProblem) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorValidationProblem) GetDetailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorValidationProblem) SetDetail(v string) {
	o.Detail = v
}

// GetProblemType returns the ProblemType field value
func (o *ErrorValidationProblem) GetProblemType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProblemType
}

// GetProblemTypeOk returns a tuple with the ProblemType field value
// and a boolean to check if the value has been set.
func (o *ErrorValidationProblem) GetProblemTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProblemType, true
}

// SetProblemType sets field value
func (o *ErrorValidationProblem) SetProblemType(v string) {
	o.ProblemType = v
}

func (o ErrorValidationProblem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["detail"] = o.Detail
	}
	if true {
		toSerialize["problem_type"] = o.ProblemType
	}
	return json.Marshal(toSerialize)
}

func (v *ErrorValidationProblem) UnmarshalJSON(src []byte) error {
    type ErrorValidationProblemUnmarshalTarget ErrorValidationProblem

	var intermediateResult ErrorValidationProblemUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = ErrorValidationProblem(intermediateResult)
	return nil
}
type NullableErrorValidationProblem struct {
	value *ErrorValidationProblem
	isSet bool
}

func (v NullableErrorValidationProblem) Get() *ErrorValidationProblem {
	return v.value
}

func (v *NullableErrorValidationProblem) Set(val *ErrorValidationProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorValidationProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorValidationProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorValidationProblem(val *ErrorValidationProblem) *NullableErrorValidationProblem {
	return &NullableErrorValidationProblem{value: val, isSet: true}
}

func (v NullableErrorValidationProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorValidationProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


