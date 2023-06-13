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

// EndUserDetailsRequest struct for EndUserDetailsRequest
type EndUserDetailsRequest struct {
	// Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
	EndUserEmailAddress string `json:"end_user_email_address"`
	// Your end user's organization.
	EndUserOrganizationName string `json:"end_user_organization_name"`
	// This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
	EndUserOriginId string `json:"end_user_origin_id"`
	// The integration categories to show in Merge Link.
	Categories []CategoriesEnum `json:"categories"`
	// The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://www.merge.dev/docs/basics/integration-metadata/.
	Integration NullableString `json:"integration,omitempty"`
	// An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
	LinkExpiryMins *int32 `json:"link_expiry_mins,omitempty"`
	// Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/product/integrations,-fast.-say-hello-to-magic-link/.
	ShouldCreateMagicLinkUrl NullableBool `json:"should_create_magic_link_url,omitempty"`
	// An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
	CommonModels []CommonModelScopesBodyRequest `json:"common_models,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewEndUserDetailsRequest instantiates a new EndUserDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserDetailsRequest(endUserEmailAddress string, endUserOrganizationName string, endUserOriginId string, categories []CategoriesEnum) *EndUserDetailsRequest {
	this := EndUserDetailsRequest{}
	this.EndUserEmailAddress = endUserEmailAddress
	this.EndUserOrganizationName = endUserOrganizationName
	this.EndUserOriginId = endUserOriginId
	this.Categories = categories
	var linkExpiryMins int32 = 30
	this.LinkExpiryMins = &linkExpiryMins
	var shouldCreateMagicLinkUrl bool = false
	this.ShouldCreateMagicLinkUrl = *NewNullableBool(&shouldCreateMagicLinkUrl)
	return &this
}

// NewEndUserDetailsRequestWithDefaults instantiates a new EndUserDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserDetailsRequestWithDefaults() *EndUserDetailsRequest {
	this := EndUserDetailsRequest{}
	var linkExpiryMins int32 = 30
	this.LinkExpiryMins = &linkExpiryMins
	var shouldCreateMagicLinkUrl bool = false
	this.ShouldCreateMagicLinkUrl = *NewNullableBool(&shouldCreateMagicLinkUrl)
	return &this
}

// GetEndUserEmailAddress returns the EndUserEmailAddress field value
func (o *EndUserDetailsRequest) GetEndUserEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndUserEmailAddress
}

// GetEndUserEmailAddressOk returns a tuple with the EndUserEmailAddress field value
// and a boolean to check if the value has been set.
func (o *EndUserDetailsRequest) GetEndUserEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndUserEmailAddress, true
}

// SetEndUserEmailAddress sets field value
func (o *EndUserDetailsRequest) SetEndUserEmailAddress(v string) {
	o.EndUserEmailAddress = v
}

// GetEndUserOrganizationName returns the EndUserOrganizationName field value
func (o *EndUserDetailsRequest) GetEndUserOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndUserOrganizationName
}

// GetEndUserOrganizationNameOk returns a tuple with the EndUserOrganizationName field value
// and a boolean to check if the value has been set.
func (o *EndUserDetailsRequest) GetEndUserOrganizationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndUserOrganizationName, true
}

// SetEndUserOrganizationName sets field value
func (o *EndUserDetailsRequest) SetEndUserOrganizationName(v string) {
	o.EndUserOrganizationName = v
}

// GetEndUserOriginId returns the EndUserOriginId field value
func (o *EndUserDetailsRequest) GetEndUserOriginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndUserOriginId
}

// GetEndUserOriginIdOk returns a tuple with the EndUserOriginId field value
// and a boolean to check if the value has been set.
func (o *EndUserDetailsRequest) GetEndUserOriginIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndUserOriginId, true
}

// SetEndUserOriginId sets field value
func (o *EndUserDetailsRequest) SetEndUserOriginId(v string) {
	o.EndUserOriginId = v
}

// GetCategories returns the Categories field value
func (o *EndUserDetailsRequest) GetCategories() []CategoriesEnum {
	if o == nil {
		var ret []CategoriesEnum
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *EndUserDetailsRequest) GetCategoriesOk() (*[]CategoriesEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *EndUserDetailsRequest) SetCategories(v []CategoriesEnum) {
	o.Categories = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndUserDetailsRequest) GetIntegration() string {
	if o == nil || o.Integration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Integration.Get()
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndUserDetailsRequest) GetIntegrationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Integration.Get(), o.Integration.IsSet()
}

// HasIntegration returns a boolean if a field has been set.
func (o *EndUserDetailsRequest) HasIntegration() bool {
	if o != nil && o.Integration.IsSet() {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given NullableString and assigns it to the Integration field.
func (o *EndUserDetailsRequest) SetIntegration(v string) {
	o.Integration.Set(&v)
}
// SetIntegrationNil sets the value for Integration to be an explicit nil
func (o *EndUserDetailsRequest) SetIntegrationNil() {
	o.Integration.Set(nil)
}

// UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
func (o *EndUserDetailsRequest) UnsetIntegration() {
	o.Integration.Unset()
}

// GetLinkExpiryMins returns the LinkExpiryMins field value if set, zero value otherwise.
func (o *EndUserDetailsRequest) GetLinkExpiryMins() int32 {
	if o == nil || o.LinkExpiryMins == nil {
		var ret int32
		return ret
	}
	return *o.LinkExpiryMins
}

// GetLinkExpiryMinsOk returns a tuple with the LinkExpiryMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserDetailsRequest) GetLinkExpiryMinsOk() (*int32, bool) {
	if o == nil || o.LinkExpiryMins == nil {
		return nil, false
	}
	return o.LinkExpiryMins, true
}

// HasLinkExpiryMins returns a boolean if a field has been set.
func (o *EndUserDetailsRequest) HasLinkExpiryMins() bool {
	if o != nil && o.LinkExpiryMins != nil {
		return true
	}

	return false
}

// SetLinkExpiryMins gets a reference to the given int32 and assigns it to the LinkExpiryMins field.
func (o *EndUserDetailsRequest) SetLinkExpiryMins(v int32) {
	o.LinkExpiryMins = &v
}

// GetShouldCreateMagicLinkUrl returns the ShouldCreateMagicLinkUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndUserDetailsRequest) GetShouldCreateMagicLinkUrl() bool {
	if o == nil || o.ShouldCreateMagicLinkUrl.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ShouldCreateMagicLinkUrl.Get()
}

// GetShouldCreateMagicLinkUrlOk returns a tuple with the ShouldCreateMagicLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndUserDetailsRequest) GetShouldCreateMagicLinkUrlOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShouldCreateMagicLinkUrl.Get(), o.ShouldCreateMagicLinkUrl.IsSet()
}

// HasShouldCreateMagicLinkUrl returns a boolean if a field has been set.
func (o *EndUserDetailsRequest) HasShouldCreateMagicLinkUrl() bool {
	if o != nil && o.ShouldCreateMagicLinkUrl.IsSet() {
		return true
	}

	return false
}

// SetShouldCreateMagicLinkUrl gets a reference to the given NullableBool and assigns it to the ShouldCreateMagicLinkUrl field.
func (o *EndUserDetailsRequest) SetShouldCreateMagicLinkUrl(v bool) {
	o.ShouldCreateMagicLinkUrl.Set(&v)
}
// SetShouldCreateMagicLinkUrlNil sets the value for ShouldCreateMagicLinkUrl to be an explicit nil
func (o *EndUserDetailsRequest) SetShouldCreateMagicLinkUrlNil() {
	o.ShouldCreateMagicLinkUrl.Set(nil)
}

// UnsetShouldCreateMagicLinkUrl ensures that no value is present for ShouldCreateMagicLinkUrl, not even an explicit nil
func (o *EndUserDetailsRequest) UnsetShouldCreateMagicLinkUrl() {
	o.ShouldCreateMagicLinkUrl.Unset()
}

// GetCommonModels returns the CommonModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndUserDetailsRequest) GetCommonModels() []CommonModelScopesBodyRequest {
	if o == nil  {
		var ret []CommonModelScopesBodyRequest
		return ret
	}
	return o.CommonModels
}

// GetCommonModelsOk returns a tuple with the CommonModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndUserDetailsRequest) GetCommonModelsOk() (*[]CommonModelScopesBodyRequest, bool) {
	if o == nil || o.CommonModels == nil {
		return nil, false
	}
	return &o.CommonModels, true
}

// HasCommonModels returns a boolean if a field has been set.
func (o *EndUserDetailsRequest) HasCommonModels() bool {
	if o != nil && o.CommonModels != nil {
		return true
	}

	return false
}

// SetCommonModels gets a reference to the given []CommonModelScopesBodyRequest and assigns it to the CommonModels field.
func (o *EndUserDetailsRequest) SetCommonModels(v []CommonModelScopesBodyRequest) {
	o.CommonModels = v
}

func (o EndUserDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["end_user_email_address"] = o.EndUserEmailAddress
	}
	if true {
		toSerialize["end_user_organization_name"] = o.EndUserOrganizationName
	}
	if true {
		toSerialize["end_user_origin_id"] = o.EndUserOriginId
	}
	if true {
		toSerialize["categories"] = o.Categories
	}
	if o.Integration.IsSet() {
		toSerialize["integration"] = o.Integration.Get()
	}
	if o.LinkExpiryMins != nil {
		toSerialize["link_expiry_mins"] = o.LinkExpiryMins
	}
	if o.ShouldCreateMagicLinkUrl.IsSet() {
		toSerialize["should_create_magic_link_url"] = o.ShouldCreateMagicLinkUrl.Get()
	}
	if o.CommonModels != nil {
		toSerialize["common_models"] = o.CommonModels
	}
	return json.Marshal(toSerialize)
}

func (v *EndUserDetailsRequest) UnmarshalJSON(src []byte) error {
    type EndUserDetailsRequestUnmarshalTarget EndUserDetailsRequest

	var intermediateResult EndUserDetailsRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = EndUserDetailsRequest(intermediateResult)
	return nil
}
type NullableEndUserDetailsRequest struct {
	value *EndUserDetailsRequest
	isSet bool
}

func (v NullableEndUserDetailsRequest) Get() *EndUserDetailsRequest {
	return v.value
}

func (v *NullableEndUserDetailsRequest) Set(val *EndUserDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserDetailsRequest(val *EndUserDetailsRequest) *NullableEndUserDetailsRequest {
	return &NullableEndUserDetailsRequest{value: val, isSet: true}
}

func (v NullableEndUserDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


