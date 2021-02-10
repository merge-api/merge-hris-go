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

// AccountToken struct for AccountToken
type AccountToken struct {
	AccountToken string `json:"account_token"`
}

// NewAccountToken instantiates a new AccountToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountToken(accountToken string, ) *AccountToken {
	this := AccountToken{}
	this.AccountToken = accountToken
	return &this
}

// NewAccountTokenWithDefaults instantiates a new AccountToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTokenWithDefaults() *AccountToken {
	this := AccountToken{}
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *AccountToken) GetAccountToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *AccountToken) GetAccountTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *AccountToken) SetAccountToken(v string) {
	o.AccountToken = v
}

func (o AccountToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_token"] = o.AccountToken
	}
	return json.Marshal(toSerialize)
}

type NullableAccountToken struct {
	value *AccountToken
	isSet bool
}

func (v NullableAccountToken) Get() *AccountToken {
	return v.value
}

func (v *NullableAccountToken) Set(val *AccountToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountToken(val *AccountToken) *NullableAccountToken {
	return &NullableAccountToken{value: val, isSet: true}
}

func (v NullableAccountToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

