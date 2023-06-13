# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Name** | Pointer to **NullableString** | The location&#39;s name. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The location&#39;s phone number. | [optional] 
**Street1** | Pointer to **NullableString** | Line 1 of the location&#39;s street address. | [optional] 
**Street2** | Pointer to **NullableString** | Line 2 of the location&#39;s street address. | [optional] 
**City** | Pointer to **NullableString** | The location&#39;s city. | [optional] 
**State** | Pointer to **NullableString** | The location&#39;s state. Represents a region if outside of the US. | [optional] 
**ZipCode** | Pointer to **NullableString** | The location&#39;s zip code or postal code. | [optional] 
**Country** | Pointer to [**NullableCountryEnum**](CountryEnum.md) | The location&#39;s country.  * &#x60;AF&#x60; - Afghanistan * &#x60;AX&#x60; - Åland Islands * &#x60;AL&#x60; - Albania * &#x60;DZ&#x60; - Algeria * &#x60;AS&#x60; - American Samoa * &#x60;AD&#x60; - Andorra * &#x60;AO&#x60; - Angola * &#x60;AI&#x60; - Anguilla * &#x60;AQ&#x60; - Antarctica * &#x60;AG&#x60; - Antigua and Barbuda * &#x60;AR&#x60; - Argentina * &#x60;AM&#x60; - Armenia * &#x60;AW&#x60; - Aruba * &#x60;AU&#x60; - Australia * &#x60;AT&#x60; - Austria * &#x60;AZ&#x60; - Azerbaijan * &#x60;BS&#x60; - Bahamas * &#x60;BH&#x60; - Bahrain * &#x60;BD&#x60; - Bangladesh * &#x60;BB&#x60; - Barbados * &#x60;BY&#x60; - Belarus * &#x60;BE&#x60; - Belgium * &#x60;BZ&#x60; - Belize * &#x60;BJ&#x60; - Benin * &#x60;BM&#x60; - Bermuda * &#x60;BT&#x60; - Bhutan * &#x60;BO&#x60; - Bolivia * &#x60;BQ&#x60; - Bonaire, Sint Eustatius and Saba * &#x60;BA&#x60; - Bosnia and Herzegovina * &#x60;BW&#x60; - Botswana * &#x60;BV&#x60; - Bouvet Island * &#x60;BR&#x60; - Brazil * &#x60;IO&#x60; - British Indian Ocean Territory * &#x60;BN&#x60; - Brunei * &#x60;BG&#x60; - Bulgaria * &#x60;BF&#x60; - Burkina Faso * &#x60;BI&#x60; - Burundi * &#x60;CV&#x60; - Cabo Verde * &#x60;KH&#x60; - Cambodia * &#x60;CM&#x60; - Cameroon * &#x60;CA&#x60; - Canada * &#x60;KY&#x60; - Cayman Islands * &#x60;CF&#x60; - Central African Republic * &#x60;TD&#x60; - Chad * &#x60;CL&#x60; - Chile * &#x60;CN&#x60; - China * &#x60;CX&#x60; - Christmas Island * &#x60;CC&#x60; - Cocos (Keeling) Islands * &#x60;CO&#x60; - Colombia * &#x60;KM&#x60; - Comoros * &#x60;CG&#x60; - Congo * &#x60;CD&#x60; - Congo (the Democratic Republic of the) * &#x60;CK&#x60; - Cook Islands * &#x60;CR&#x60; - Costa Rica * &#x60;CI&#x60; - Côte d&#39;Ivoire * &#x60;HR&#x60; - Croatia * &#x60;CU&#x60; - Cuba * &#x60;CW&#x60; - Curaçao * &#x60;CY&#x60; - Cyprus * &#x60;CZ&#x60; - Czechia * &#x60;DK&#x60; - Denmark * &#x60;DJ&#x60; - Djibouti * &#x60;DM&#x60; - Dominica * &#x60;DO&#x60; - Dominican Republic * &#x60;EC&#x60; - Ecuador * &#x60;EG&#x60; - Egypt * &#x60;SV&#x60; - El Salvador * &#x60;GQ&#x60; - Equatorial Guinea * &#x60;ER&#x60; - Eritrea * &#x60;EE&#x60; - Estonia * &#x60;SZ&#x60; - Eswatini * &#x60;ET&#x60; - Ethiopia * &#x60;FK&#x60; - Falkland Islands (Malvinas) * &#x60;FO&#x60; - Faroe Islands * &#x60;FJ&#x60; - Fiji * &#x60;FI&#x60; - Finland * &#x60;FR&#x60; - France * &#x60;GF&#x60; - French Guiana * &#x60;PF&#x60; - French Polynesia * &#x60;TF&#x60; - French Southern Territories * &#x60;GA&#x60; - Gabon * &#x60;GM&#x60; - Gambia * &#x60;GE&#x60; - Georgia * &#x60;DE&#x60; - Germany * &#x60;GH&#x60; - Ghana * &#x60;GI&#x60; - Gibraltar * &#x60;GR&#x60; - Greece * &#x60;GL&#x60; - Greenland * &#x60;GD&#x60; - Grenada * &#x60;GP&#x60; - Guadeloupe * &#x60;GU&#x60; - Guam * &#x60;GT&#x60; - Guatemala * &#x60;GG&#x60; - Guernsey * &#x60;GN&#x60; - Guinea * &#x60;GW&#x60; - Guinea-Bissau * &#x60;GY&#x60; - Guyana * &#x60;HT&#x60; - Haiti * &#x60;HM&#x60; - Heard Island and McDonald Islands * &#x60;VA&#x60; - Holy See * &#x60;HN&#x60; - Honduras * &#x60;HK&#x60; - Hong Kong * &#x60;HU&#x60; - Hungary * &#x60;IS&#x60; - Iceland * &#x60;IN&#x60; - India * &#x60;ID&#x60; - Indonesia * &#x60;IR&#x60; - Iran * &#x60;IQ&#x60; - Iraq * &#x60;IE&#x60; - Ireland * &#x60;IM&#x60; - Isle of Man * &#x60;IL&#x60; - Israel * &#x60;IT&#x60; - Italy * &#x60;JM&#x60; - Jamaica * &#x60;JP&#x60; - Japan * &#x60;JE&#x60; - Jersey * &#x60;JO&#x60; - Jordan * &#x60;KZ&#x60; - Kazakhstan * &#x60;KE&#x60; - Kenya * &#x60;KI&#x60; - Kiribati * &#x60;KW&#x60; - Kuwait * &#x60;KG&#x60; - Kyrgyzstan * &#x60;LA&#x60; - Laos * &#x60;LV&#x60; - Latvia * &#x60;LB&#x60; - Lebanon * &#x60;LS&#x60; - Lesotho * &#x60;LR&#x60; - Liberia * &#x60;LY&#x60; - Libya * &#x60;LI&#x60; - Liechtenstein * &#x60;LT&#x60; - Lithuania * &#x60;LU&#x60; - Luxembourg * &#x60;MO&#x60; - Macao * &#x60;MG&#x60; - Madagascar * &#x60;MW&#x60; - Malawi * &#x60;MY&#x60; - Malaysia * &#x60;MV&#x60; - Maldives * &#x60;ML&#x60; - Mali * &#x60;MT&#x60; - Malta * &#x60;MH&#x60; - Marshall Islands * &#x60;MQ&#x60; - Martinique * &#x60;MR&#x60; - Mauritania * &#x60;MU&#x60; - Mauritius * &#x60;YT&#x60; - Mayotte * &#x60;MX&#x60; - Mexico * &#x60;FM&#x60; - Micronesia (Federated States of) * &#x60;MD&#x60; - Moldova * &#x60;MC&#x60; - Monaco * &#x60;MN&#x60; - Mongolia * &#x60;ME&#x60; - Montenegro * &#x60;MS&#x60; - Montserrat * &#x60;MA&#x60; - Morocco * &#x60;MZ&#x60; - Mozambique * &#x60;MM&#x60; - Myanmar * &#x60;NA&#x60; - Namibia * &#x60;NR&#x60; - Nauru * &#x60;NP&#x60; - Nepal * &#x60;NL&#x60; - Netherlands * &#x60;NC&#x60; - New Caledonia * &#x60;NZ&#x60; - New Zealand * &#x60;NI&#x60; - Nicaragua * &#x60;NE&#x60; - Niger * &#x60;NG&#x60; - Nigeria * &#x60;NU&#x60; - Niue * &#x60;NF&#x60; - Norfolk Island * &#x60;KP&#x60; - North Korea * &#x60;MK&#x60; - North Macedonia * &#x60;MP&#x60; - Northern Mariana Islands * &#x60;NO&#x60; - Norway * &#x60;OM&#x60; - Oman * &#x60;PK&#x60; - Pakistan * &#x60;PW&#x60; - Palau * &#x60;PS&#x60; - Palestine, State of * &#x60;PA&#x60; - Panama * &#x60;PG&#x60; - Papua New Guinea * &#x60;PY&#x60; - Paraguay * &#x60;PE&#x60; - Peru * &#x60;PH&#x60; - Philippines * &#x60;PN&#x60; - Pitcairn * &#x60;PL&#x60; - Poland * &#x60;PT&#x60; - Portugal * &#x60;PR&#x60; - Puerto Rico * &#x60;QA&#x60; - Qatar * &#x60;RE&#x60; - Réunion * &#x60;RO&#x60; - Romania * &#x60;RU&#x60; - Russia * &#x60;RW&#x60; - Rwanda * &#x60;BL&#x60; - Saint Barthélemy * &#x60;SH&#x60; - Saint Helena, Ascension and Tristan da Cunha * &#x60;KN&#x60; - Saint Kitts and Nevis * &#x60;LC&#x60; - Saint Lucia * &#x60;MF&#x60; - Saint Martin (French part) * &#x60;PM&#x60; - Saint Pierre and Miquelon * &#x60;VC&#x60; - Saint Vincent and the Grenadines * &#x60;WS&#x60; - Samoa * &#x60;SM&#x60; - San Marino * &#x60;ST&#x60; - Sao Tome and Principe * &#x60;SA&#x60; - Saudi Arabia * &#x60;SN&#x60; - Senegal * &#x60;RS&#x60; - Serbia * &#x60;SC&#x60; - Seychelles * &#x60;SL&#x60; - Sierra Leone * &#x60;SG&#x60; - Singapore * &#x60;SX&#x60; - Sint Maarten (Dutch part) * &#x60;SK&#x60; - Slovakia * &#x60;SI&#x60; - Slovenia * &#x60;SB&#x60; - Solomon Islands * &#x60;SO&#x60; - Somalia * &#x60;ZA&#x60; - South Africa * &#x60;GS&#x60; - South Georgia and the South Sandwich Islands * &#x60;KR&#x60; - South Korea * &#x60;SS&#x60; - South Sudan * &#x60;ES&#x60; - Spain * &#x60;LK&#x60; - Sri Lanka * &#x60;SD&#x60; - Sudan * &#x60;SR&#x60; - Suriname * &#x60;SJ&#x60; - Svalbard and Jan Mayen * &#x60;SE&#x60; - Sweden * &#x60;CH&#x60; - Switzerland * &#x60;SY&#x60; - Syria * &#x60;TW&#x60; - Taiwan * &#x60;TJ&#x60; - Tajikistan * &#x60;TZ&#x60; - Tanzania * &#x60;TH&#x60; - Thailand * &#x60;TL&#x60; - Timor-Leste * &#x60;TG&#x60; - Togo * &#x60;TK&#x60; - Tokelau * &#x60;TO&#x60; - Tonga * &#x60;TT&#x60; - Trinidad and Tobago * &#x60;TN&#x60; - Tunisia * &#x60;TR&#x60; - Turkey * &#x60;TM&#x60; - Turkmenistan * &#x60;TC&#x60; - Turks and Caicos Islands * &#x60;TV&#x60; - Tuvalu * &#x60;UG&#x60; - Uganda * &#x60;UA&#x60; - Ukraine * &#x60;AE&#x60; - United Arab Emirates * &#x60;GB&#x60; - United Kingdom * &#x60;UM&#x60; - United States Minor Outlying Islands * &#x60;US&#x60; - United States of America * &#x60;UY&#x60; - Uruguay * &#x60;UZ&#x60; - Uzbekistan * &#x60;VU&#x60; - Vanuatu * &#x60;VE&#x60; - Venezuela * &#x60;VN&#x60; - Vietnam * &#x60;VG&#x60; - Virgin Islands (British) * &#x60;VI&#x60; - Virgin Islands (U.S.) * &#x60;WF&#x60; - Wallis and Futuna * &#x60;EH&#x60; - Western Sahara * &#x60;YE&#x60; - Yemen * &#x60;ZM&#x60; - Zambia * &#x60;ZW&#x60; - Zimbabwe | [optional] 
**LocationType** | Pointer to [**NullableLocationTypeEnum**](LocationTypeEnum.md) | The location&#39;s type. Can be either WORK or HOME  * &#x60;HOME&#x60; - HOME * &#x60;WORK&#x60; - WORK | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Location) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Location) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Location) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Location) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Location) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Location) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Location) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Location) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhoneNumber

`func (o *Location) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Location) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Location) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Location) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *Location) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *Location) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetStreet1

`func (o *Location) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *Location) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *Location) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.

### HasStreet1

`func (o *Location) HasStreet1() bool`

HasStreet1 returns a boolean if a field has been set.

### SetStreet1Nil

`func (o *Location) SetStreet1Nil(b bool)`

 SetStreet1Nil sets the value for Street1 to be an explicit nil

### UnsetStreet1
`func (o *Location) UnsetStreet1()`

UnsetStreet1 ensures that no value is present for Street1, not even an explicit nil
### GetStreet2

`func (o *Location) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *Location) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *Location) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *Location) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### SetStreet2Nil

`func (o *Location) SetStreet2Nil(b bool)`

 SetStreet2Nil sets the value for Street2 to be an explicit nil

### UnsetStreet2
`func (o *Location) UnsetStreet2()`

UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
### GetCity

`func (o *Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Location) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Location) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Location) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Location) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *Location) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Location) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Location) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Location) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Location) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Location) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetZipCode

`func (o *Location) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Location) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Location) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Location) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *Location) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *Location) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCountry

`func (o *Location) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Location) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Location) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Location) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLocationType

`func (o *Location) GetLocationType() LocationTypeEnum`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *Location) GetLocationTypeOk() (*LocationTypeEnum, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *Location) SetLocationType(v LocationTypeEnum)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *Location) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *Location) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *Location) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Location) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Location) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Location) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Location) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Location) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Location) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Location) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Location) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Location) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Location) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Location) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Location) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Location) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Location) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Location) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Location) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Location) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Location) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Location) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Location) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


