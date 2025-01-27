/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the EmailNotificationOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailNotificationOption{}

// EmailNotificationOption struct for EmailNotificationOption
type EmailNotificationOption struct {
	// If true, then the manager is notified of the lifecycle state change.
	NotifyManagers *bool `json:"notifyManagers,omitempty"`
	// If true, then all the admins are notified of the lifecycle state change.
	NotifyAllAdmins *bool `json:"notifyAllAdmins,omitempty"`
	// If true, then the users specified in \"emailAddressList\" below are notified of lifecycle state change.
	NotifySpecificUsers *bool `json:"notifySpecificUsers,omitempty"`
	// List of user email addresses. If \"notifySpecificUsers\" option is true, then these users are notified of lifecycle state change.
	EmailAddressList []string `json:"emailAddressList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailNotificationOption EmailNotificationOption

// NewEmailNotificationOption instantiates a new EmailNotificationOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationOption() *EmailNotificationOption {
	this := EmailNotificationOption{}
	return &this
}

// NewEmailNotificationOptionWithDefaults instantiates a new EmailNotificationOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationOptionWithDefaults() *EmailNotificationOption {
	this := EmailNotificationOption{}
	return &this
}

// GetNotifyManagers returns the NotifyManagers field value if set, zero value otherwise.
func (o *EmailNotificationOption) GetNotifyManagers() bool {
	if o == nil || isNil(o.NotifyManagers) {
		var ret bool
		return ret
	}
	return *o.NotifyManagers
}

// GetNotifyManagersOk returns a tuple with the NotifyManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOption) GetNotifyManagersOk() (*bool, bool) {
	if o == nil || isNil(o.NotifyManagers) {
		return nil, false
	}
	return o.NotifyManagers, true
}

// HasNotifyManagers returns a boolean if a field has been set.
func (o *EmailNotificationOption) HasNotifyManagers() bool {
	if o != nil && !isNil(o.NotifyManagers) {
		return true
	}

	return false
}

// SetNotifyManagers gets a reference to the given bool and assigns it to the NotifyManagers field.
func (o *EmailNotificationOption) SetNotifyManagers(v bool) {
	o.NotifyManagers = &v
}

// GetNotifyAllAdmins returns the NotifyAllAdmins field value if set, zero value otherwise.
func (o *EmailNotificationOption) GetNotifyAllAdmins() bool {
	if o == nil || isNil(o.NotifyAllAdmins) {
		var ret bool
		return ret
	}
	return *o.NotifyAllAdmins
}

// GetNotifyAllAdminsOk returns a tuple with the NotifyAllAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOption) GetNotifyAllAdminsOk() (*bool, bool) {
	if o == nil || isNil(o.NotifyAllAdmins) {
		return nil, false
	}
	return o.NotifyAllAdmins, true
}

// HasNotifyAllAdmins returns a boolean if a field has been set.
func (o *EmailNotificationOption) HasNotifyAllAdmins() bool {
	if o != nil && !isNil(o.NotifyAllAdmins) {
		return true
	}

	return false
}

// SetNotifyAllAdmins gets a reference to the given bool and assigns it to the NotifyAllAdmins field.
func (o *EmailNotificationOption) SetNotifyAllAdmins(v bool) {
	o.NotifyAllAdmins = &v
}

// GetNotifySpecificUsers returns the NotifySpecificUsers field value if set, zero value otherwise.
func (o *EmailNotificationOption) GetNotifySpecificUsers() bool {
	if o == nil || isNil(o.NotifySpecificUsers) {
		var ret bool
		return ret
	}
	return *o.NotifySpecificUsers
}

// GetNotifySpecificUsersOk returns a tuple with the NotifySpecificUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOption) GetNotifySpecificUsersOk() (*bool, bool) {
	if o == nil || isNil(o.NotifySpecificUsers) {
		return nil, false
	}
	return o.NotifySpecificUsers, true
}

// HasNotifySpecificUsers returns a boolean if a field has been set.
func (o *EmailNotificationOption) HasNotifySpecificUsers() bool {
	if o != nil && !isNil(o.NotifySpecificUsers) {
		return true
	}

	return false
}

// SetNotifySpecificUsers gets a reference to the given bool and assigns it to the NotifySpecificUsers field.
func (o *EmailNotificationOption) SetNotifySpecificUsers(v bool) {
	o.NotifySpecificUsers = &v
}

// GetEmailAddressList returns the EmailAddressList field value if set, zero value otherwise.
func (o *EmailNotificationOption) GetEmailAddressList() []string {
	if o == nil || isNil(o.EmailAddressList) {
		var ret []string
		return ret
	}
	return o.EmailAddressList
}

// GetEmailAddressListOk returns a tuple with the EmailAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOption) GetEmailAddressListOk() ([]string, bool) {
	if o == nil || isNil(o.EmailAddressList) {
		return nil, false
	}
	return o.EmailAddressList, true
}

// HasEmailAddressList returns a boolean if a field has been set.
func (o *EmailNotificationOption) HasEmailAddressList() bool {
	if o != nil && !isNil(o.EmailAddressList) {
		return true
	}

	return false
}

// SetEmailAddressList gets a reference to the given []string and assigns it to the EmailAddressList field.
func (o *EmailNotificationOption) SetEmailAddressList(v []string) {
	o.EmailAddressList = v
}

func (o EmailNotificationOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNotificationOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NotifyManagers) {
		toSerialize["notifyManagers"] = o.NotifyManagers
	}
	if !isNil(o.NotifyAllAdmins) {
		toSerialize["notifyAllAdmins"] = o.NotifyAllAdmins
	}
	if !isNil(o.NotifySpecificUsers) {
		toSerialize["notifySpecificUsers"] = o.NotifySpecificUsers
	}
	if !isNil(o.EmailAddressList) {
		toSerialize["emailAddressList"] = o.EmailAddressList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailNotificationOption) UnmarshalJSON(bytes []byte) (err error) {
	varEmailNotificationOption := _EmailNotificationOption{}

	if err = json.Unmarshal(bytes, &varEmailNotificationOption); err == nil {
		*o = EmailNotificationOption(varEmailNotificationOption)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "notifyManagers")
		delete(additionalProperties, "notifyAllAdmins")
		delete(additionalProperties, "notifySpecificUsers")
		delete(additionalProperties, "emailAddressList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailNotificationOption struct {
	value *EmailNotificationOption
	isSet bool
}

func (v NullableEmailNotificationOption) Get() *EmailNotificationOption {
	return v.value
}

func (v *NullableEmailNotificationOption) Set(val *EmailNotificationOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationOption(val *EmailNotificationOption) *NullableEmailNotificationOption {
	return &NullableEmailNotificationOption{value: val, isSet: true}
}

func (v NullableEmailNotificationOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


