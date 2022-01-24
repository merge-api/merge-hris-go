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
	"time"
)

// Employee # The Employee Object ### Description The `Employee` object is used to represent an Employee for a company.  ### Usage Example Fetch from the `LIST Employee` endpoint and filter by `ID` to show all employees.
type Employee struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The employee's number that appears in the remote UI. Note: This is distinct from the remote_id field, which is a unique identifier for the employee set by the remote API, and is not exposed to the user.
	EmployeeNumber NullableString `json:"employee_number,omitempty"`
	// The ID of the employee's company.
	Company NullableString `json:"company,omitempty"`
	// The employee's first name.
	FirstName NullableString `json:"first_name,omitempty"`
	// The employee's last name.
	LastName NullableString `json:"last_name,omitempty"`
	// The employee's full name, to use for display purposes. If a preferred first name is available, the full name will include the preferred first name.
	DisplayFullName NullableString `json:"display_full_name,omitempty"`
	// The employee's work email.
	WorkEmail NullableString `json:"work_email,omitempty"`
	// The employee's personal email.
	PersonalEmail NullableString `json:"personal_email,omitempty"`
	// The employee's mobile phone number.
	MobilePhoneNumber NullableString `json:"mobile_phone_number,omitempty"`
	// Array of `Employment` IDs for this Employee.
	Employments *[]string `json:"employments,omitempty"`
	// The employee's home address.
	HomeLocation NullableString `json:"home_location,omitempty"`
	// The employee's work address.
	WorkLocation NullableString `json:"work_location,omitempty"`
	// The employee ID of the employee's manager.
	Manager NullableString `json:"manager,omitempty"`
	// The employee's team.
	Team NullableString `json:"team,omitempty"`
	// The employee's pay group
	PayGroup NullableString `json:"pay_group,omitempty"`
	// The employee's social security number.
	Ssn NullableString `json:"ssn,omitempty"`
	Gender string `json:"gender"`
	Ethnicity string `json:"ethnicity"`
	MaritalStatus string `json:"marital_status"`
	// The employee's date of birth.
	DateOfBirth NullableTime `json:"date_of_birth,omitempty"`
	// The date that the employee was hired, usually the day that an offer letter is signed. If an employee has multiple hire dates from previous employments, this represents the most recent hire date. Note: If you're looking for the employee's start date, refer to the start_date field.
	HireDate NullableTime `json:"hire_date,omitempty"`
	// The date that the employee started working. If an employee has multiple start dates from previous employments, this represents the most recent start date.
	StartDate NullableTime `json:"start_date,omitempty"`
	EmploymentStatus string `json:"employment_status"`
	// The employee's termination date.
	TerminationDate NullableTime `json:"termination_date,omitempty"`
	// The URL of the employee's avatar image.
	Avatar NullableString `json:"avatar,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// Custom fields configured for a given model.
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewEmployee instantiates a new Employee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployee(gender string, ethnicity string, maritalStatus string, employmentStatus string) *Employee {
	this := Employee{}
	this.Gender = gender
	this.Ethnicity = ethnicity
	this.MaritalStatus = maritalStatus
	this.EmploymentStatus = employmentStatus
	return &this
}

// NewEmployeeWithDefaults instantiates a new Employee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeWithDefaults() *Employee {
	this := Employee{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Employee) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employee) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Employee) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Employee) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Employee) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Employee) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Employee) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Employee) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetEmployeeNumber() string {
	if o == nil || o.EmployeeNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeeNumber.Get()
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetEmployeeNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeNumber.Get(), o.EmployeeNumber.IsSet()
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *Employee) HasEmployeeNumber() bool {
	if o != nil && o.EmployeeNumber.IsSet() {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given NullableString and assigns it to the EmployeeNumber field.
func (o *Employee) SetEmployeeNumber(v string) {
	o.EmployeeNumber.Set(&v)
}
// SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil
func (o *Employee) SetEmployeeNumberNil() {
	o.EmployeeNumber.Set(nil)
}

// UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
func (o *Employee) UnsetEmployeeNumber() {
	o.EmployeeNumber.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *Employee) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *Employee) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *Employee) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *Employee) UnsetCompany() {
	o.Company.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *Employee) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *Employee) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *Employee) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *Employee) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *Employee) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *Employee) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *Employee) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *Employee) UnsetLastName() {
	o.LastName.Unset()
}

// GetDisplayFullName returns the DisplayFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetDisplayFullName() string {
	if o == nil || o.DisplayFullName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayFullName.Get()
}

// GetDisplayFullNameOk returns a tuple with the DisplayFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetDisplayFullNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayFullName.Get(), o.DisplayFullName.IsSet()
}

// HasDisplayFullName returns a boolean if a field has been set.
func (o *Employee) HasDisplayFullName() bool {
	if o != nil && o.DisplayFullName.IsSet() {
		return true
	}

	return false
}

// SetDisplayFullName gets a reference to the given NullableString and assigns it to the DisplayFullName field.
func (o *Employee) SetDisplayFullName(v string) {
	o.DisplayFullName.Set(&v)
}
// SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil
func (o *Employee) SetDisplayFullNameNil() {
	o.DisplayFullName.Set(nil)
}

// UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
func (o *Employee) UnsetDisplayFullName() {
	o.DisplayFullName.Unset()
}

// GetWorkEmail returns the WorkEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetWorkEmail() string {
	if o == nil || o.WorkEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.WorkEmail.Get()
}

// GetWorkEmailOk returns a tuple with the WorkEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetWorkEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkEmail.Get(), o.WorkEmail.IsSet()
}

// HasWorkEmail returns a boolean if a field has been set.
func (o *Employee) HasWorkEmail() bool {
	if o != nil && o.WorkEmail.IsSet() {
		return true
	}

	return false
}

// SetWorkEmail gets a reference to the given NullableString and assigns it to the WorkEmail field.
func (o *Employee) SetWorkEmail(v string) {
	o.WorkEmail.Set(&v)
}
// SetWorkEmailNil sets the value for WorkEmail to be an explicit nil
func (o *Employee) SetWorkEmailNil() {
	o.WorkEmail.Set(nil)
}

// UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
func (o *Employee) UnsetWorkEmail() {
	o.WorkEmail.Unset()
}

// GetPersonalEmail returns the PersonalEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetPersonalEmail() string {
	if o == nil || o.PersonalEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.PersonalEmail.Get()
}

// GetPersonalEmailOk returns a tuple with the PersonalEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetPersonalEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalEmail.Get(), o.PersonalEmail.IsSet()
}

// HasPersonalEmail returns a boolean if a field has been set.
func (o *Employee) HasPersonalEmail() bool {
	if o != nil && o.PersonalEmail.IsSet() {
		return true
	}

	return false
}

// SetPersonalEmail gets a reference to the given NullableString and assigns it to the PersonalEmail field.
func (o *Employee) SetPersonalEmail(v string) {
	o.PersonalEmail.Set(&v)
}
// SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil
func (o *Employee) SetPersonalEmailNil() {
	o.PersonalEmail.Set(nil)
}

// UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
func (o *Employee) UnsetPersonalEmail() {
	o.PersonalEmail.Unset()
}

// GetMobilePhoneNumber returns the MobilePhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetMobilePhoneNumber() string {
	if o == nil || o.MobilePhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.MobilePhoneNumber.Get()
}

// GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetMobilePhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MobilePhoneNumber.Get(), o.MobilePhoneNumber.IsSet()
}

// HasMobilePhoneNumber returns a boolean if a field has been set.
func (o *Employee) HasMobilePhoneNumber() bool {
	if o != nil && o.MobilePhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetMobilePhoneNumber gets a reference to the given NullableString and assigns it to the MobilePhoneNumber field.
func (o *Employee) SetMobilePhoneNumber(v string) {
	o.MobilePhoneNumber.Set(&v)
}
// SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil
func (o *Employee) SetMobilePhoneNumberNil() {
	o.MobilePhoneNumber.Set(nil)
}

// UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
func (o *Employee) UnsetMobilePhoneNumber() {
	o.MobilePhoneNumber.Unset()
}

// GetEmployments returns the Employments field value if set, zero value otherwise.
func (o *Employee) GetEmployments() []string {
	if o == nil || o.Employments == nil {
		var ret []string
		return ret
	}
	return *o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employee) GetEmploymentsOk() (*[]string, bool) {
	if o == nil || o.Employments == nil {
		return nil, false
	}
	return o.Employments, true
}

// HasEmployments returns a boolean if a field has been set.
func (o *Employee) HasEmployments() bool {
	if o != nil && o.Employments != nil {
		return true
	}

	return false
}

// SetEmployments gets a reference to the given []string and assigns it to the Employments field.
func (o *Employee) SetEmployments(v []string) {
	o.Employments = &v
}

// GetHomeLocation returns the HomeLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetHomeLocation() string {
	if o == nil || o.HomeLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.HomeLocation.Get()
}

// GetHomeLocationOk returns a tuple with the HomeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetHomeLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HomeLocation.Get(), o.HomeLocation.IsSet()
}

// HasHomeLocation returns a boolean if a field has been set.
func (o *Employee) HasHomeLocation() bool {
	if o != nil && o.HomeLocation.IsSet() {
		return true
	}

	return false
}

// SetHomeLocation gets a reference to the given NullableString and assigns it to the HomeLocation field.
func (o *Employee) SetHomeLocation(v string) {
	o.HomeLocation.Set(&v)
}
// SetHomeLocationNil sets the value for HomeLocation to be an explicit nil
func (o *Employee) SetHomeLocationNil() {
	o.HomeLocation.Set(nil)
}

// UnsetHomeLocation ensures that no value is present for HomeLocation, not even an explicit nil
func (o *Employee) UnsetHomeLocation() {
	o.HomeLocation.Unset()
}

// GetWorkLocation returns the WorkLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetWorkLocation() string {
	if o == nil || o.WorkLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.WorkLocation.Get()
}

// GetWorkLocationOk returns a tuple with the WorkLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetWorkLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkLocation.Get(), o.WorkLocation.IsSet()
}

// HasWorkLocation returns a boolean if a field has been set.
func (o *Employee) HasWorkLocation() bool {
	if o != nil && o.WorkLocation.IsSet() {
		return true
	}

	return false
}

// SetWorkLocation gets a reference to the given NullableString and assigns it to the WorkLocation field.
func (o *Employee) SetWorkLocation(v string) {
	o.WorkLocation.Set(&v)
}
// SetWorkLocationNil sets the value for WorkLocation to be an explicit nil
func (o *Employee) SetWorkLocationNil() {
	o.WorkLocation.Set(nil)
}

// UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
func (o *Employee) UnsetWorkLocation() {
	o.WorkLocation.Unset()
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetManager() string {
	if o == nil || o.Manager.Get() == nil {
		var ret string
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetManagerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *Employee) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableString and assigns it to the Manager field.
func (o *Employee) SetManager(v string) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *Employee) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *Employee) UnsetManager() {
	o.Manager.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetTeam() string {
	if o == nil || o.Team.Get() == nil {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetTeamOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *Employee) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *Employee) SetTeam(v string) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *Employee) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *Employee) UnsetTeam() {
	o.Team.Unset()
}

// GetPayGroup returns the PayGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetPayGroup() string {
	if o == nil || o.PayGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayGroup.Get()
}

// GetPayGroupOk returns a tuple with the PayGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetPayGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayGroup.Get(), o.PayGroup.IsSet()
}

// HasPayGroup returns a boolean if a field has been set.
func (o *Employee) HasPayGroup() bool {
	if o != nil && o.PayGroup.IsSet() {
		return true
	}

	return false
}

// SetPayGroup gets a reference to the given NullableString and assigns it to the PayGroup field.
func (o *Employee) SetPayGroup(v string) {
	o.PayGroup.Set(&v)
}
// SetPayGroupNil sets the value for PayGroup to be an explicit nil
func (o *Employee) SetPayGroupNil() {
	o.PayGroup.Set(nil)
}

// UnsetPayGroup ensures that no value is present for PayGroup, not even an explicit nil
func (o *Employee) UnsetPayGroup() {
	o.PayGroup.Unset()
}

// GetSsn returns the Ssn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetSsn() string {
	if o == nil || o.Ssn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ssn.Get()
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetSsnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ssn.Get(), o.Ssn.IsSet()
}

// HasSsn returns a boolean if a field has been set.
func (o *Employee) HasSsn() bool {
	if o != nil && o.Ssn.IsSet() {
		return true
	}

	return false
}

// SetSsn gets a reference to the given NullableString and assigns it to the Ssn field.
func (o *Employee) SetSsn(v string) {
	o.Ssn.Set(&v)
}
// SetSsnNil sets the value for Ssn to be an explicit nil
func (o *Employee) SetSsnNil() {
	o.Ssn.Set(nil)
}

// UnsetSsn ensures that no value is present for Ssn, not even an explicit nil
func (o *Employee) UnsetSsn() {
	o.Ssn.Unset()
}

// GetGender returns the Gender field value
func (o *Employee) GetGender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *Employee) GetGenderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *Employee) SetGender(v string) {
	o.Gender = v
}

// GetEthnicity returns the Ethnicity field value
func (o *Employee) GetEthnicity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ethnicity
}

// GetEthnicityOk returns a tuple with the Ethnicity field value
// and a boolean to check if the value has been set.
func (o *Employee) GetEthnicityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ethnicity, true
}

// SetEthnicity sets field value
func (o *Employee) SetEthnicity(v string) {
	o.Ethnicity = v
}

// GetMaritalStatus returns the MaritalStatus field value
func (o *Employee) GetMaritalStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaritalStatus
}

// GetMaritalStatusOk returns a tuple with the MaritalStatus field value
// and a boolean to check if the value has been set.
func (o *Employee) GetMaritalStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaritalStatus, true
}

// SetMaritalStatus sets field value
func (o *Employee) SetMaritalStatus(v string) {
	o.MaritalStatus = v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetDateOfBirth() time.Time {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetDateOfBirthOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *Employee) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth.IsSet() {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given NullableTime and assigns it to the DateOfBirth field.
func (o *Employee) SetDateOfBirth(v time.Time) {
	o.DateOfBirth.Set(&v)
}
// SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil
func (o *Employee) SetDateOfBirthNil() {
	o.DateOfBirth.Set(nil)
}

// UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
func (o *Employee) UnsetDateOfBirth() {
	o.DateOfBirth.Unset()
}

// GetHireDate returns the HireDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetHireDate() time.Time {
	if o == nil || o.HireDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.HireDate.Get()
}

// GetHireDateOk returns a tuple with the HireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetHireDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HireDate.Get(), o.HireDate.IsSet()
}

// HasHireDate returns a boolean if a field has been set.
func (o *Employee) HasHireDate() bool {
	if o != nil && o.HireDate.IsSet() {
		return true
	}

	return false
}

// SetHireDate gets a reference to the given NullableTime and assigns it to the HireDate field.
func (o *Employee) SetHireDate(v time.Time) {
	o.HireDate.Set(&v)
}
// SetHireDateNil sets the value for HireDate to be an explicit nil
func (o *Employee) SetHireDateNil() {
	o.HireDate.Set(nil)
}

// UnsetHireDate ensures that no value is present for HireDate, not even an explicit nil
func (o *Employee) UnsetHireDate() {
	o.HireDate.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *Employee) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *Employee) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *Employee) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *Employee) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEmploymentStatus returns the EmploymentStatus field value
func (o *Employee) GetEmploymentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmploymentStatus
}

// GetEmploymentStatusOk returns a tuple with the EmploymentStatus field value
// and a boolean to check if the value has been set.
func (o *Employee) GetEmploymentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmploymentStatus, true
}

// SetEmploymentStatus sets field value
func (o *Employee) SetEmploymentStatus(v string) {
	o.EmploymentStatus = v
}

// GetTerminationDate returns the TerminationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetTerminationDate() time.Time {
	if o == nil || o.TerminationDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TerminationDate.Get()
}

// GetTerminationDateOk returns a tuple with the TerminationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetTerminationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TerminationDate.Get(), o.TerminationDate.IsSet()
}

// HasTerminationDate returns a boolean if a field has been set.
func (o *Employee) HasTerminationDate() bool {
	if o != nil && o.TerminationDate.IsSet() {
		return true
	}

	return false
}

// SetTerminationDate gets a reference to the given NullableTime and assigns it to the TerminationDate field.
func (o *Employee) SetTerminationDate(v time.Time) {
	o.TerminationDate.Set(&v)
}
// SetTerminationDateNil sets the value for TerminationDate to be an explicit nil
func (o *Employee) SetTerminationDateNil() {
	o.TerminationDate.Set(nil)
}

// UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
func (o *Employee) UnsetTerminationDate() {
	o.TerminationDate.Unset()
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetAvatarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *Employee) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *Employee) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *Employee) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *Employee) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Employee) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Employee) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetCustomFields() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetCustomFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return &o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Employee) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Employee) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Employee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.EmployeeNumber.IsSet() {
		toSerialize["employee_number"] = o.EmployeeNumber.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.DisplayFullName.IsSet() {
		toSerialize["display_full_name"] = o.DisplayFullName.Get()
	}
	if o.WorkEmail.IsSet() {
		toSerialize["work_email"] = o.WorkEmail.Get()
	}
	if o.PersonalEmail.IsSet() {
		toSerialize["personal_email"] = o.PersonalEmail.Get()
	}
	if o.MobilePhoneNumber.IsSet() {
		toSerialize["mobile_phone_number"] = o.MobilePhoneNumber.Get()
	}
	if o.Employments != nil {
		toSerialize["employments"] = o.Employments
	}
	if o.HomeLocation.IsSet() {
		toSerialize["home_location"] = o.HomeLocation.Get()
	}
	if o.WorkLocation.IsSet() {
		toSerialize["work_location"] = o.WorkLocation.Get()
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.PayGroup.IsSet() {
		toSerialize["pay_group"] = o.PayGroup.Get()
	}
	if o.Ssn.IsSet() {
		toSerialize["ssn"] = o.Ssn.Get()
	}
	if true {
		toSerialize["gender"] = o.Gender
	}
	if true {
		toSerialize["ethnicity"] = o.Ethnicity
	}
	if true {
		toSerialize["marital_status"] = o.MaritalStatus
	}
	if o.DateOfBirth.IsSet() {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if o.HireDate.IsSet() {
		toSerialize["hire_date"] = o.HireDate.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if true {
		toSerialize["employment_status"] = o.EmploymentStatus
	}
	if o.TerminationDate.IsSet() {
		toSerialize["termination_date"] = o.TerminationDate.Get()
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return json.Marshal(toSerialize)
}

type NullableEmployee struct {
	value *Employee
	isSet bool
}

func (v NullableEmployee) Get() *Employee {
	return v.value
}

func (v *NullableEmployee) Set(val *Employee) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployee) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployee(val *Employee) *NullableEmployee {
	return &NullableEmployee{value: val, isSet: true}
}

func (v NullableEmployee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


