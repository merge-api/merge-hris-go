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

// Group # The Group Object ### Description The `Group` object is used to represent Group information that employees belong to. This is often referenced with an Employee object.  Please note: The teams object will fulfill most use cases. The Groups object is for power-users that want all types of groups at a company and the optionality of pulling multiple groups for an employee.  ### Usage Example Fetch from the `LIST Employee` endpoint and expand groups to view an employee's groups.
type Group struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The parent group for this group.
	ParentGroup NullableString `json:"parent_group,omitempty"`
	// The group name.
	Name NullableString `json:"name,omitempty"`
	// The group type
	Type NullableGroupTypeEnum `json:"type,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Group) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Group) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Group) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Group) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetParentGroup returns the ParentGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetParentGroup() string {
	if o == nil || o.ParentGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentGroup.Get()
}

// GetParentGroupOk returns a tuple with the ParentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetParentGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentGroup.Get(), o.ParentGroup.IsSet()
}

// HasParentGroup returns a boolean if a field has been set.
func (o *Group) HasParentGroup() bool {
	if o != nil && o.ParentGroup.IsSet() {
		return true
	}

	return false
}

// SetParentGroup gets a reference to the given NullableString and assigns it to the ParentGroup field.
func (o *Group) SetParentGroup(v string) {
	o.ParentGroup.Set(&v)
}
// SetParentGroupNil sets the value for ParentGroup to be an explicit nil
func (o *Group) SetParentGroupNil() {
	o.ParentGroup.Set(nil)
}

// UnsetParentGroup ensures that no value is present for ParentGroup, not even an explicit nil
func (o *Group) UnsetParentGroup() {
	o.ParentGroup.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Group) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Group) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Group) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Group) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetType() GroupTypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret GroupTypeEnum
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetTypeOk() (*GroupTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Group) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableGroupTypeEnum and assigns it to the Type field.
func (o *Group) SetType(v GroupTypeEnum) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Group) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Group) UnsetType() {
	o.Type.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Group) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Group) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Group) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Group) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Group) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.ParentGroup.IsSet() {
		toSerialize["parent_group"] = o.ParentGroup.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	return json.Marshal(toSerialize)
}

func (v *Group) UnmarshalJSON(src []byte) error {
    type GroupUnmarshalTarget Group

	var intermediateResult GroupUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Group(intermediateResult)
	return nil
}
type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


