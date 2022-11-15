/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// PasswordInfoQueryDTO struct for PasswordInfoQueryDTO
type PasswordInfoQueryDTO struct {
	// The login name of the user
	UserName *string `json:"userName,omitempty"`
	// The display name of the source
	SourceName *string `json:"sourceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordInfoQueryDTO PasswordInfoQueryDTO

// NewPasswordInfoQueryDTO instantiates a new PasswordInfoQueryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordInfoQueryDTO() *PasswordInfoQueryDTO {
	this := PasswordInfoQueryDTO{}
	return &this
}

// NewPasswordInfoQueryDTOWithDefaults instantiates a new PasswordInfoQueryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordInfoQueryDTOWithDefaults() *PasswordInfoQueryDTO {
	this := PasswordInfoQueryDTO{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *PasswordInfoQueryDTO) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfoQueryDTO) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *PasswordInfoQueryDTO) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *PasswordInfoQueryDTO) SetUserName(v string) {
	o.UserName = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *PasswordInfoQueryDTO) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfoQueryDTO) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *PasswordInfoQueryDTO) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *PasswordInfoQueryDTO) SetSourceName(v string) {
	o.SourceName = &v
}

func (o PasswordInfoQueryDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordInfoQueryDTO) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordInfoQueryDTO := _PasswordInfoQueryDTO{}

	if err = json.Unmarshal(bytes, &varPasswordInfoQueryDTO); err == nil {
		*o = PasswordInfoQueryDTO(varPasswordInfoQueryDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "userName")
		delete(additionalProperties, "sourceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordInfoQueryDTO struct {
	value *PasswordInfoQueryDTO
	isSet bool
}

func (v NullablePasswordInfoQueryDTO) Get() *PasswordInfoQueryDTO {
	return v.value
}

func (v *NullablePasswordInfoQueryDTO) Set(val *PasswordInfoQueryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordInfoQueryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordInfoQueryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordInfoQueryDTO(val *PasswordInfoQueryDTO) *NullablePasswordInfoQueryDTO {
	return &NullablePasswordInfoQueryDTO{value: val, isSet: true}
}

func (v NullablePasswordInfoQueryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordInfoQueryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


