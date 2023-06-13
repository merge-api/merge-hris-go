# EmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeNumber** | Pointer to **NullableString** | The employee&#39;s number that appears in the third-party integration&#39;s UI. | [optional] 
**Company** | Pointer to **NullableString** | The ID of the employee&#39;s company. | [optional] 
**FirstName** | Pointer to **NullableString** | The employee&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The employee&#39;s last name. | [optional] 
**DisplayFullName** | Pointer to **NullableString** | The employee&#39;s full name, to use for display purposes. If a preferred first name is available, the full name will include the preferred first name. | [optional] 
**Username** | Pointer to **NullableString** | The employee&#39;s username that appears in the remote UI. | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**WorkEmail** | Pointer to **NullableString** | The employee&#39;s work email. | [optional] 
**PersonalEmail** | Pointer to **NullableString** | The employee&#39;s personal email. | [optional] 
**MobilePhoneNumber** | Pointer to **NullableString** | The employee&#39;s mobile phone number. | [optional] 
**Employments** | Pointer to **[]string** | Array of &#x60;Employment&#x60; IDs for this Employee. | [optional] 
**HomeLocation** | Pointer to **NullableString** | The employee&#39;s home address. | [optional] 
**WorkLocation** | Pointer to **NullableString** | The employee&#39;s work address. | [optional] 
**Manager** | Pointer to **NullableString** | The employee ID of the employee&#39;s manager. | [optional] 
**Team** | Pointer to **NullableString** | The employee&#39;s team. | [optional] 
**PayGroup** | Pointer to **NullableString** | The employee&#39;s pay group | [optional] 
**Ssn** | Pointer to **NullableString** | The employee&#39;s social security number. | [optional] 
**Gender** | Pointer to [**NullableGenderEnum**](GenderEnum.md) | The employee&#39;s gender.  * &#x60;MALE&#x60; - MALE * &#x60;FEMALE&#x60; - FEMALE * &#x60;NON-BINARY&#x60; - NON-BINARY * &#x60;OTHER&#x60; - OTHER * &#x60;PREFER_NOT_TO_DISCLOSE&#x60; - PREFER_NOT_TO_DISCLOSE | [optional] 
**Ethnicity** | Pointer to [**NullableEthnicityEnum**](EthnicityEnum.md) | The employee&#39;s ethnicity.  * &#x60;AMERICAN_INDIAN_OR_ALASKA_NATIVE&#x60; - AMERICAN_INDIAN_OR_ALASKA_NATIVE * &#x60;ASIAN_OR_INDIAN_SUBCONTINENT&#x60; - ASIAN_OR_INDIAN_SUBCONTINENT * &#x60;BLACK_OR_AFRICAN_AMERICAN&#x60; - BLACK_OR_AFRICAN_AMERICAN * &#x60;HISPANIC_OR_LATINO&#x60; - HISPANIC_OR_LATINO * &#x60;NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER&#x60; - NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER * &#x60;TWO_OR_MORE_RACES&#x60; - TWO_OR_MORE_RACES * &#x60;WHITE&#x60; - WHITE * &#x60;PREFER_NOT_TO_DISCLOSE&#x60; - PREFER_NOT_TO_DISCLOSE | [optional] 
**MaritalStatus** | Pointer to [**NullableMaritalStatusEnum**](MaritalStatusEnum.md) | The employee&#39;s filing status as related to marital status.  * &#x60;SINGLE&#x60; - SINGLE * &#x60;MARRIED_FILING_JOINTLY&#x60; - MARRIED_FILING_JOINTLY * &#x60;MARRIED_FILING_SEPARATELY&#x60; - MARRIED_FILING_SEPARATELY * &#x60;HEAD_OF_HOUSEHOLD&#x60; - HEAD_OF_HOUSEHOLD * &#x60;QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD&#x60; - QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD | [optional] 
**DateOfBirth** | Pointer to **NullableTime** | The employee&#39;s date of birth. | [optional] 
**HireDate** | Pointer to **NullableTime** | The date that the employee was hired, usually the day that an offer letter is signed. If an employee has multiple hire dates from previous employments, this represents the most recent hire date. Note: If you&#39;re looking for the employee&#39;s start date, refer to the start_date field. | [optional] 
**StartDate** | Pointer to **NullableTime** | The date that the employee started working. If an employee was rehired, the most recent start date will be returned. | [optional] 
**EmploymentStatus** | Pointer to [**NullableEmploymentStatusEnum**](EmploymentStatusEnum.md) | The employment status of the employee.  * &#x60;ACTIVE&#x60; - ACTIVE * &#x60;PENDING&#x60; - PENDING * &#x60;INACTIVE&#x60; - INACTIVE | [optional] 
**TerminationDate** | Pointer to **NullableTime** | The employee&#39;s termination date. | [optional] 
**Avatar** | Pointer to **NullableString** | The URL of the employee&#39;s avatar image. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEmployeeRequest

`func NewEmployeeRequest() *EmployeeRequest`

NewEmployeeRequest instantiates a new EmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeRequestWithDefaults

`func NewEmployeeRequestWithDefaults() *EmployeeRequest`

NewEmployeeRequestWithDefaults instantiates a new EmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeNumber

`func (o *EmployeeRequest) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *EmployeeRequest) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *EmployeeRequest) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *EmployeeRequest) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### SetEmployeeNumberNil

`func (o *EmployeeRequest) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *EmployeeRequest) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetCompany

`func (o *EmployeeRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EmployeeRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EmployeeRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EmployeeRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *EmployeeRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *EmployeeRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFirstName

`func (o *EmployeeRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EmployeeRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *EmployeeRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *EmployeeRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *EmployeeRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployeeRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployeeRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *EmployeeRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *EmployeeRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *EmployeeRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayFullName

`func (o *EmployeeRequest) GetDisplayFullName() string`

GetDisplayFullName returns the DisplayFullName field if non-nil, zero value otherwise.

### GetDisplayFullNameOk

`func (o *EmployeeRequest) GetDisplayFullNameOk() (*string, bool)`

GetDisplayFullNameOk returns a tuple with the DisplayFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFullName

`func (o *EmployeeRequest) SetDisplayFullName(v string)`

SetDisplayFullName sets DisplayFullName field to given value.

### HasDisplayFullName

`func (o *EmployeeRequest) HasDisplayFullName() bool`

HasDisplayFullName returns a boolean if a field has been set.

### SetDisplayFullNameNil

`func (o *EmployeeRequest) SetDisplayFullNameNil(b bool)`

 SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil

### UnsetDisplayFullName
`func (o *EmployeeRequest) UnsetDisplayFullName()`

UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
### GetUsername

`func (o *EmployeeRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmployeeRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmployeeRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmployeeRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *EmployeeRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *EmployeeRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetGroups

`func (o *EmployeeRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmployeeRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmployeeRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmployeeRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetWorkEmail

`func (o *EmployeeRequest) GetWorkEmail() string`

GetWorkEmail returns the WorkEmail field if non-nil, zero value otherwise.

### GetWorkEmailOk

`func (o *EmployeeRequest) GetWorkEmailOk() (*string, bool)`

GetWorkEmailOk returns a tuple with the WorkEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkEmail

`func (o *EmployeeRequest) SetWorkEmail(v string)`

SetWorkEmail sets WorkEmail field to given value.

### HasWorkEmail

`func (o *EmployeeRequest) HasWorkEmail() bool`

HasWorkEmail returns a boolean if a field has been set.

### SetWorkEmailNil

`func (o *EmployeeRequest) SetWorkEmailNil(b bool)`

 SetWorkEmailNil sets the value for WorkEmail to be an explicit nil

### UnsetWorkEmail
`func (o *EmployeeRequest) UnsetWorkEmail()`

UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
### GetPersonalEmail

`func (o *EmployeeRequest) GetPersonalEmail() string`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *EmployeeRequest) GetPersonalEmailOk() (*string, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *EmployeeRequest) SetPersonalEmail(v string)`

SetPersonalEmail sets PersonalEmail field to given value.

### HasPersonalEmail

`func (o *EmployeeRequest) HasPersonalEmail() bool`

HasPersonalEmail returns a boolean if a field has been set.

### SetPersonalEmailNil

`func (o *EmployeeRequest) SetPersonalEmailNil(b bool)`

 SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil

### UnsetPersonalEmail
`func (o *EmployeeRequest) UnsetPersonalEmail()`

UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
### GetMobilePhoneNumber

`func (o *EmployeeRequest) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *EmployeeRequest) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *EmployeeRequest) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *EmployeeRequest) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### SetMobilePhoneNumberNil

`func (o *EmployeeRequest) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *EmployeeRequest) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetEmployments

`func (o *EmployeeRequest) GetEmployments() []string`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *EmployeeRequest) GetEmploymentsOk() (*[]string, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *EmployeeRequest) SetEmployments(v []string)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *EmployeeRequest) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### GetHomeLocation

`func (o *EmployeeRequest) GetHomeLocation() string`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *EmployeeRequest) GetHomeLocationOk() (*string, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *EmployeeRequest) SetHomeLocation(v string)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *EmployeeRequest) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.

### SetHomeLocationNil

`func (o *EmployeeRequest) SetHomeLocationNil(b bool)`

 SetHomeLocationNil sets the value for HomeLocation to be an explicit nil

### UnsetHomeLocation
`func (o *EmployeeRequest) UnsetHomeLocation()`

UnsetHomeLocation ensures that no value is present for HomeLocation, not even an explicit nil
### GetWorkLocation

`func (o *EmployeeRequest) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *EmployeeRequest) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *EmployeeRequest) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *EmployeeRequest) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *EmployeeRequest) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *EmployeeRequest) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetManager

`func (o *EmployeeRequest) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *EmployeeRequest) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *EmployeeRequest) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *EmployeeRequest) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *EmployeeRequest) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *EmployeeRequest) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetTeam

`func (o *EmployeeRequest) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *EmployeeRequest) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *EmployeeRequest) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *EmployeeRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *EmployeeRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *EmployeeRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetPayGroup

`func (o *EmployeeRequest) GetPayGroup() string`

GetPayGroup returns the PayGroup field if non-nil, zero value otherwise.

### GetPayGroupOk

`func (o *EmployeeRequest) GetPayGroupOk() (*string, bool)`

GetPayGroupOk returns a tuple with the PayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayGroup

`func (o *EmployeeRequest) SetPayGroup(v string)`

SetPayGroup sets PayGroup field to given value.

### HasPayGroup

`func (o *EmployeeRequest) HasPayGroup() bool`

HasPayGroup returns a boolean if a field has been set.

### SetPayGroupNil

`func (o *EmployeeRequest) SetPayGroupNil(b bool)`

 SetPayGroupNil sets the value for PayGroup to be an explicit nil

### UnsetPayGroup
`func (o *EmployeeRequest) UnsetPayGroup()`

UnsetPayGroup ensures that no value is present for PayGroup, not even an explicit nil
### GetSsn

`func (o *EmployeeRequest) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *EmployeeRequest) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *EmployeeRequest) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *EmployeeRequest) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### SetSsnNil

`func (o *EmployeeRequest) SetSsnNil(b bool)`

 SetSsnNil sets the value for Ssn to be an explicit nil

### UnsetSsn
`func (o *EmployeeRequest) UnsetSsn()`

UnsetSsn ensures that no value is present for Ssn, not even an explicit nil
### GetGender

`func (o *EmployeeRequest) GetGender() GenderEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *EmployeeRequest) GetGenderOk() (*GenderEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *EmployeeRequest) SetGender(v GenderEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *EmployeeRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *EmployeeRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *EmployeeRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetEthnicity

`func (o *EmployeeRequest) GetEthnicity() EthnicityEnum`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *EmployeeRequest) GetEthnicityOk() (*EthnicityEnum, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *EmployeeRequest) SetEthnicity(v EthnicityEnum)`

SetEthnicity sets Ethnicity field to given value.

### HasEthnicity

`func (o *EmployeeRequest) HasEthnicity() bool`

HasEthnicity returns a boolean if a field has been set.

### SetEthnicityNil

`func (o *EmployeeRequest) SetEthnicityNil(b bool)`

 SetEthnicityNil sets the value for Ethnicity to be an explicit nil

### UnsetEthnicity
`func (o *EmployeeRequest) UnsetEthnicity()`

UnsetEthnicity ensures that no value is present for Ethnicity, not even an explicit nil
### GetMaritalStatus

`func (o *EmployeeRequest) GetMaritalStatus() MaritalStatusEnum`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *EmployeeRequest) GetMaritalStatusOk() (*MaritalStatusEnum, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *EmployeeRequest) SetMaritalStatus(v MaritalStatusEnum)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *EmployeeRequest) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### SetMaritalStatusNil

`func (o *EmployeeRequest) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *EmployeeRequest) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetDateOfBirth

`func (o *EmployeeRequest) GetDateOfBirth() time.Time`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *EmployeeRequest) GetDateOfBirthOk() (*time.Time, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *EmployeeRequest) SetDateOfBirth(v time.Time)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *EmployeeRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *EmployeeRequest) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *EmployeeRequest) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetHireDate

`func (o *EmployeeRequest) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *EmployeeRequest) GetHireDateOk() (*time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *EmployeeRequest) SetHireDate(v time.Time)`

SetHireDate sets HireDate field to given value.

### HasHireDate

`func (o *EmployeeRequest) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### SetHireDateNil

`func (o *EmployeeRequest) SetHireDateNil(b bool)`

 SetHireDateNil sets the value for HireDate to be an explicit nil

### UnsetHireDate
`func (o *EmployeeRequest) UnsetHireDate()`

UnsetHireDate ensures that no value is present for HireDate, not even an explicit nil
### GetStartDate

`func (o *EmployeeRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmployeeRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmployeeRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EmployeeRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *EmployeeRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EmployeeRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEmploymentStatus

`func (o *EmployeeRequest) GetEmploymentStatus() EmploymentStatusEnum`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *EmployeeRequest) GetEmploymentStatusOk() (*EmploymentStatusEnum, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *EmployeeRequest) SetEmploymentStatus(v EmploymentStatusEnum)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *EmployeeRequest) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### SetEmploymentStatusNil

`func (o *EmployeeRequest) SetEmploymentStatusNil(b bool)`

 SetEmploymentStatusNil sets the value for EmploymentStatus to be an explicit nil

### UnsetEmploymentStatus
`func (o *EmployeeRequest) UnsetEmploymentStatus()`

UnsetEmploymentStatus ensures that no value is present for EmploymentStatus, not even an explicit nil
### GetTerminationDate

`func (o *EmployeeRequest) GetTerminationDate() time.Time`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *EmployeeRequest) GetTerminationDateOk() (*time.Time, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *EmployeeRequest) SetTerminationDate(v time.Time)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *EmployeeRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *EmployeeRequest) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *EmployeeRequest) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetAvatar

`func (o *EmployeeRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *EmployeeRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *EmployeeRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *EmployeeRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *EmployeeRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *EmployeeRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetIntegrationParams

`func (o *EmployeeRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *EmployeeRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *EmployeeRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *EmployeeRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *EmployeeRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *EmployeeRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *EmployeeRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *EmployeeRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *EmployeeRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *EmployeeRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *EmployeeRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *EmployeeRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


