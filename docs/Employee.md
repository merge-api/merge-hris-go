# Employee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Company** | Pointer to **NullableString** | The ID of the Employee&#39;s company. | [optional] 
**FirstName** | Pointer to **NullableString** | The employee&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The employee&#39;s last name. | [optional] 
**DisplayFullName** | Pointer to **NullableString** | The employee&#39;s full name, to use for display purposes. | [optional] 
**WorkEmail** | Pointer to **NullableString** | The employee&#39;s work email. | [optional] 
**PersonalEmail** | Pointer to **NullableString** | The employee&#39;s personal email. | [optional] 
**MobilePhoneNumber** | Pointer to **NullableString** | The employee&#39;s mobile phone number. | [optional] 
**Employments** | Pointer to **[]string** |  | [optional] [readonly] 
**HomeLocation** | Pointer to **NullableString** | The employee&#39;s home address. | [optional] 
**WorkLocation** | Pointer to **NullableString** | The employee&#39;s work address. | [optional] 
**Manager** | Pointer to **NullableString** | The employeee ID of the employee&#39;s manager. | [optional] 
**Team** | Pointer to **NullableString** | The employee&#39;s team. | [optional] 
**Ssn** | Pointer to **NullableString** | The employee&#39;s social security number. | [optional] 
**Gender** | Pointer to [**NullableGenderEnum**](GenderEnum.md) | The employee&#39;s gender. | [optional] 
**Ethnicity** | Pointer to [**NullableEthnicityEnum**](EthnicityEnum.md) | The employee&#39;s ethnicity. | [optional] 
**MaritalStatus** | Pointer to [**NullableMaritalStatusEnum**](MaritalStatusEnum.md) | The employee&#39;s marital status. | [optional] 
**DateOfBirth** | Pointer to **NullableTime** | The employee&#39;s date of birth. | [optional] 
**HireDate** | Pointer to **NullableTime** | The employee&#39;s hire date. | [optional] 
**EmploymentStatus** | Pointer to [**NullableEmploymentStatusEnum**](EmploymentStatusEnum.md) | The employment status of the employee. | [optional] 
**TerminationDate** | Pointer to **NullableTime** | The employee&#39;s termination date. | [optional] 
**Avatar** | Pointer to **NullableString** | The URL of the employee&#39;s avatar image. | [optional] 
**Documents** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewEmployee

`func NewEmployee() *Employee`

NewEmployee instantiates a new Employee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeWithDefaults

`func NewEmployeeWithDefaults() *Employee`

NewEmployeeWithDefaults instantiates a new Employee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Employee) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Employee) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Employee) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Employee) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Employee) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Employee) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Employee) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Employee) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Employee) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Employee) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetCompany

`func (o *Employee) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Employee) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Employee) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Employee) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Employee) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Employee) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFirstName

`func (o *Employee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Employee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Employee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Employee) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Employee) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Employee) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Employee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Employee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Employee) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Employee) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Employee) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Employee) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayFullName

`func (o *Employee) GetDisplayFullName() string`

GetDisplayFullName returns the DisplayFullName field if non-nil, zero value otherwise.

### GetDisplayFullNameOk

`func (o *Employee) GetDisplayFullNameOk() (*string, bool)`

GetDisplayFullNameOk returns a tuple with the DisplayFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFullName

`func (o *Employee) SetDisplayFullName(v string)`

SetDisplayFullName sets DisplayFullName field to given value.

### HasDisplayFullName

`func (o *Employee) HasDisplayFullName() bool`

HasDisplayFullName returns a boolean if a field has been set.

### SetDisplayFullNameNil

`func (o *Employee) SetDisplayFullNameNil(b bool)`

 SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil

### UnsetDisplayFullName
`func (o *Employee) UnsetDisplayFullName()`

UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
### GetWorkEmail

`func (o *Employee) GetWorkEmail() string`

GetWorkEmail returns the WorkEmail field if non-nil, zero value otherwise.

### GetWorkEmailOk

`func (o *Employee) GetWorkEmailOk() (*string, bool)`

GetWorkEmailOk returns a tuple with the WorkEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkEmail

`func (o *Employee) SetWorkEmail(v string)`

SetWorkEmail sets WorkEmail field to given value.

### HasWorkEmail

`func (o *Employee) HasWorkEmail() bool`

HasWorkEmail returns a boolean if a field has been set.

### SetWorkEmailNil

`func (o *Employee) SetWorkEmailNil(b bool)`

 SetWorkEmailNil sets the value for WorkEmail to be an explicit nil

### UnsetWorkEmail
`func (o *Employee) UnsetWorkEmail()`

UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
### GetPersonalEmail

`func (o *Employee) GetPersonalEmail() string`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *Employee) GetPersonalEmailOk() (*string, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *Employee) SetPersonalEmail(v string)`

SetPersonalEmail sets PersonalEmail field to given value.

### HasPersonalEmail

`func (o *Employee) HasPersonalEmail() bool`

HasPersonalEmail returns a boolean if a field has been set.

### SetPersonalEmailNil

`func (o *Employee) SetPersonalEmailNil(b bool)`

 SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil

### UnsetPersonalEmail
`func (o *Employee) UnsetPersonalEmail()`

UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
### GetMobilePhoneNumber

`func (o *Employee) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *Employee) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *Employee) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *Employee) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### SetMobilePhoneNumberNil

`func (o *Employee) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *Employee) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetEmployments

`func (o *Employee) GetEmployments() []string`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *Employee) GetEmploymentsOk() (*[]string, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *Employee) SetEmployments(v []string)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *Employee) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### GetHomeLocation

`func (o *Employee) GetHomeLocation() string`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *Employee) GetHomeLocationOk() (*string, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *Employee) SetHomeLocation(v string)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *Employee) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.

### SetHomeLocationNil

`func (o *Employee) SetHomeLocationNil(b bool)`

 SetHomeLocationNil sets the value for HomeLocation to be an explicit nil

### UnsetHomeLocation
`func (o *Employee) UnsetHomeLocation()`

UnsetHomeLocation ensures that no value is present for HomeLocation, not even an explicit nil
### GetWorkLocation

`func (o *Employee) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *Employee) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *Employee) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *Employee) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *Employee) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *Employee) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetManager

`func (o *Employee) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Employee) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Employee) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Employee) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *Employee) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *Employee) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetTeam

`func (o *Employee) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Employee) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *Employee) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *Employee) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *Employee) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *Employee) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetSsn

`func (o *Employee) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *Employee) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *Employee) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *Employee) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### SetSsnNil

`func (o *Employee) SetSsnNil(b bool)`

 SetSsnNil sets the value for Ssn to be an explicit nil

### UnsetSsn
`func (o *Employee) UnsetSsn()`

UnsetSsn ensures that no value is present for Ssn, not even an explicit nil
### GetGender

`func (o *Employee) GetGender() GenderEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Employee) GetGenderOk() (*GenderEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Employee) SetGender(v GenderEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Employee) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *Employee) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *Employee) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetEthnicity

`func (o *Employee) GetEthnicity() EthnicityEnum`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *Employee) GetEthnicityOk() (*EthnicityEnum, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *Employee) SetEthnicity(v EthnicityEnum)`

SetEthnicity sets Ethnicity field to given value.

### HasEthnicity

`func (o *Employee) HasEthnicity() bool`

HasEthnicity returns a boolean if a field has been set.

### SetEthnicityNil

`func (o *Employee) SetEthnicityNil(b bool)`

 SetEthnicityNil sets the value for Ethnicity to be an explicit nil

### UnsetEthnicity
`func (o *Employee) UnsetEthnicity()`

UnsetEthnicity ensures that no value is present for Ethnicity, not even an explicit nil
### GetMaritalStatus

`func (o *Employee) GetMaritalStatus() MaritalStatusEnum`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *Employee) GetMaritalStatusOk() (*MaritalStatusEnum, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *Employee) SetMaritalStatus(v MaritalStatusEnum)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *Employee) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### SetMaritalStatusNil

`func (o *Employee) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *Employee) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetDateOfBirth

`func (o *Employee) GetDateOfBirth() time.Time`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Employee) GetDateOfBirthOk() (*time.Time, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Employee) SetDateOfBirth(v time.Time)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Employee) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *Employee) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *Employee) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetHireDate

`func (o *Employee) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *Employee) GetHireDateOk() (*time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *Employee) SetHireDate(v time.Time)`

SetHireDate sets HireDate field to given value.

### HasHireDate

`func (o *Employee) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### SetHireDateNil

`func (o *Employee) SetHireDateNil(b bool)`

 SetHireDateNil sets the value for HireDate to be an explicit nil

### UnsetHireDate
`func (o *Employee) UnsetHireDate()`

UnsetHireDate ensures that no value is present for HireDate, not even an explicit nil
### GetEmploymentStatus

`func (o *Employee) GetEmploymentStatus() EmploymentStatusEnum`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *Employee) GetEmploymentStatusOk() (*EmploymentStatusEnum, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *Employee) SetEmploymentStatus(v EmploymentStatusEnum)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *Employee) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### SetEmploymentStatusNil

`func (o *Employee) SetEmploymentStatusNil(b bool)`

 SetEmploymentStatusNil sets the value for EmploymentStatus to be an explicit nil

### UnsetEmploymentStatus
`func (o *Employee) UnsetEmploymentStatus()`

UnsetEmploymentStatus ensures that no value is present for EmploymentStatus, not even an explicit nil
### GetTerminationDate

`func (o *Employee) GetTerminationDate() time.Time`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *Employee) GetTerminationDateOk() (*time.Time, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *Employee) SetTerminationDate(v time.Time)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *Employee) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *Employee) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *Employee) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetAvatar

`func (o *Employee) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Employee) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Employee) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Employee) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *Employee) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *Employee) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetDocuments

`func (o *Employee) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Employee) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Employee) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Employee) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


