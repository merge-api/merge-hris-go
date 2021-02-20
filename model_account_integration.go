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

// AccountIntegration struct for AccountIntegration
type AccountIntegration struct {
	// Company name.
	Name string `json:"name"`
	// Category or categories this integration belongs to. Multiple categories should be comma separated.<br />For [ats, hris], enter <i>ats,hris</i>
	Categories *[]string `json:"categories,omitempty"`
	// Company logo.
	Image NullableString `json:"image,omitempty"`
	// Company logo in square shape.
	SquareImage NullableString `json:"square_image,omitempty"`
	// The color of this integration used for buttons and text throughout the app and landing pages. Choose a darker, saturated color.
	Color *string `json:"color,omitempty"`
}

// NewAccountIntegration instantiates a new AccountIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIntegration(name string, ) *AccountIntegration {
	this := AccountIntegration{}
	this.Name = name
	return &this
}

// NewAccountIntegrationWithDefaults instantiates a new AccountIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIntegrationWithDefaults() *AccountIntegration {
	this := AccountIntegration{}
	return &this
}

// GetName returns the Name field value
func (o *AccountIntegration) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountIntegration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountIntegration) SetName(v string) {
	o.Name = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *AccountIntegration) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIntegration) GetCategoriesOk() (*[]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *AccountIntegration) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *AccountIntegration) SetCategories(v []string) {
	o.Categories = &v
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIntegration) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIntegration) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *AccountIntegration) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *AccountIntegration) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *AccountIntegration) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *AccountIntegration) UnsetImage() {
	o.Image.Unset()
}

// GetSquareImage returns the SquareImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIntegration) GetSquareImage() string {
	if o == nil || o.SquareImage.Get() == nil {
		var ret string
		return ret
	}
	return *o.SquareImage.Get()
}

// GetSquareImageOk returns a tuple with the SquareImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIntegration) GetSquareImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SquareImage.Get(), o.SquareImage.IsSet()
}

// HasSquareImage returns a boolean if a field has been set.
func (o *AccountIntegration) HasSquareImage() bool {
	if o != nil && o.SquareImage.IsSet() {
		return true
	}

	return false
}

// SetSquareImage gets a reference to the given NullableString and assigns it to the SquareImage field.
func (o *AccountIntegration) SetSquareImage(v string) {
	o.SquareImage.Set(&v)
}
// SetSquareImageNil sets the value for SquareImage to be an explicit nil
func (o *AccountIntegration) SetSquareImageNil() {
	o.SquareImage.Set(nil)
}

// UnsetSquareImage ensures that no value is present for SquareImage, not even an explicit nil
func (o *AccountIntegration) UnsetSquareImage() {
	o.SquareImage.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *AccountIntegration) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIntegration) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *AccountIntegration) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *AccountIntegration) SetColor(v string) {
	o.Color = &v
}

func (o AccountIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.SquareImage.IsSet() {
		toSerialize["square_image"] = o.SquareImage.Get()
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	return json.Marshal(toSerialize)
}

type NullableAccountIntegration struct {
	value *AccountIntegration
	isSet bool
}

func (v NullableAccountIntegration) Get() *AccountIntegration {
	return v.value
}

func (v *NullableAccountIntegration) Set(val *AccountIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIntegration(val *AccountIntegration) *NullableAccountIntegration {
	return &NullableAccountIntegration{value: val, isSet: true}
}

func (v NullableAccountIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


