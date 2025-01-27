/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the BaseAccountAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseAccountAllOf{}

// BaseAccountAllOf struct for BaseAccountAllOf
type BaseAccountAllOf struct {
	// The ID of the account
	AccountId *string `json:"accountId,omitempty"`
	Source *AccountSource `json:"source,omitempty"`
	// Indicates if the account is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Indicates if the account is locked
	Locked *bool `json:"locked,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	// Indicates if the account has been manually correlated to an identity
	ManuallyCorrelated *bool `json:"manuallyCorrelated,omitempty"`
	// A date-time in ISO-8601 format
	PasswordLastSet NullableTime `json:"passwordLastSet,omitempty"`
	// a map or dictionary of key/value pairs
	EntitlementAttributes map[string]interface{} `json:"entitlementAttributes,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseAccountAllOf BaseAccountAllOf

// NewBaseAccountAllOf instantiates a new BaseAccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAccountAllOf() *BaseAccountAllOf {
	this := BaseAccountAllOf{}
	return &this
}

// NewBaseAccountAllOfWithDefaults instantiates a new BaseAccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAccountAllOfWithDefaults() *BaseAccountAllOf {
	this := BaseAccountAllOf{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *BaseAccountAllOf) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountAllOf) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *BaseAccountAllOf) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BaseAccountAllOf) GetSource() AccountSource {
	if o == nil || isNil(o.Source) {
		var ret AccountSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountAllOf) GetSourceOk() (*AccountSource, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given AccountSource and assigns it to the Source field.
func (o *BaseAccountAllOf) SetSource(v AccountSource) {
	o.Source = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *BaseAccountAllOf) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountAllOf) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *BaseAccountAllOf) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BaseAccountAllOf) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountAllOf) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BaseAccountAllOf) SetLocked(v bool) {
	o.Locked = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *BaseAccountAllOf) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountAllOf) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *BaseAccountAllOf) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetManuallyCorrelated returns the ManuallyCorrelated field value if set, zero value otherwise.
func (o *BaseAccountAllOf) GetManuallyCorrelated() bool {
	if o == nil || isNil(o.ManuallyCorrelated) {
		var ret bool
		return ret
	}
	return *o.ManuallyCorrelated
}

// GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountAllOf) GetManuallyCorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.ManuallyCorrelated) {
		return nil, false
	}
	return o.ManuallyCorrelated, true
}

// HasManuallyCorrelated returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasManuallyCorrelated() bool {
	if o != nil && !isNil(o.ManuallyCorrelated) {
		return true
	}

	return false
}

// SetManuallyCorrelated gets a reference to the given bool and assigns it to the ManuallyCorrelated field.
func (o *BaseAccountAllOf) SetManuallyCorrelated(v bool) {
	o.ManuallyCorrelated = &v
}

// GetPasswordLastSet returns the PasswordLastSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseAccountAllOf) GetPasswordLastSet() time.Time {
	if o == nil || isNil(o.PasswordLastSet.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PasswordLastSet.Get()
}

// GetPasswordLastSetOk returns a tuple with the PasswordLastSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseAccountAllOf) GetPasswordLastSetOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordLastSet.Get(), o.PasswordLastSet.IsSet()
}

// HasPasswordLastSet returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasPasswordLastSet() bool {
	if o != nil && o.PasswordLastSet.IsSet() {
		return true
	}

	return false
}

// SetPasswordLastSet gets a reference to the given NullableTime and assigns it to the PasswordLastSet field.
func (o *BaseAccountAllOf) SetPasswordLastSet(v time.Time) {
	o.PasswordLastSet.Set(&v)
}
// SetPasswordLastSetNil sets the value for PasswordLastSet to be an explicit nil
func (o *BaseAccountAllOf) SetPasswordLastSetNil() {
	o.PasswordLastSet.Set(nil)
}

// UnsetPasswordLastSet ensures that no value is present for PasswordLastSet, not even an explicit nil
func (o *BaseAccountAllOf) UnsetPasswordLastSet() {
	o.PasswordLastSet.Unset()
}

// GetEntitlementAttributes returns the EntitlementAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseAccountAllOf) GetEntitlementAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EntitlementAttributes
}

// GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseAccountAllOf) GetEntitlementAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.EntitlementAttributes) {
		return map[string]interface{}{}, false
	}
	return o.EntitlementAttributes, true
}

// HasEntitlementAttributes returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasEntitlementAttributes() bool {
	if o != nil && isNil(o.EntitlementAttributes) {
		return true
	}

	return false
}

// SetEntitlementAttributes gets a reference to the given map[string]interface{} and assigns it to the EntitlementAttributes field.
func (o *BaseAccountAllOf) SetEntitlementAttributes(v map[string]interface{}) {
	o.EntitlementAttributes = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseAccountAllOf) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseAccountAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *BaseAccountAllOf) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *BaseAccountAllOf) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *BaseAccountAllOf) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *BaseAccountAllOf) UnsetCreated() {
	o.Created.Unset()
}

func (o BaseAccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseAccountAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.ManuallyCorrelated) {
		toSerialize["manuallyCorrelated"] = o.ManuallyCorrelated
	}
	if o.PasswordLastSet.IsSet() {
		toSerialize["passwordLastSet"] = o.PasswordLastSet.Get()
	}
	if o.EntitlementAttributes != nil {
		toSerialize["entitlementAttributes"] = o.EntitlementAttributes
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseAccountAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBaseAccountAllOf := _BaseAccountAllOf{}

	if err = json.Unmarshal(bytes, &varBaseAccountAllOf); err == nil {
		*o = BaseAccountAllOf(varBaseAccountAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "source")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "manuallyCorrelated")
		delete(additionalProperties, "passwordLastSet")
		delete(additionalProperties, "entitlementAttributes")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseAccountAllOf struct {
	value *BaseAccountAllOf
	isSet bool
}

func (v NullableBaseAccountAllOf) Get() *BaseAccountAllOf {
	return v.value
}

func (v *NullableBaseAccountAllOf) Set(val *BaseAccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAccountAllOf(val *BaseAccountAllOf) *NullableBaseAccountAllOf {
	return &NullableBaseAccountAllOf{value: val, isSet: true}
}

func (v NullableBaseAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


