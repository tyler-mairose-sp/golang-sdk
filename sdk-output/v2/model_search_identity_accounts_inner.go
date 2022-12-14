/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
	"time"
)

// SearchIdentityAccountsInner struct for SearchIdentityAccountsInner
type SearchIdentityAccountsInner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	Source *SearchIdentityAccountsInnerSource `json:"source,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Locked *bool `json:"locked,omitempty"`
	ManuallyCorrelated *bool `json:"manuallyCorrelated,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	PasswordLastSet *time.Time `json:"passwordLastSet,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	EntitlementAttributes *SearchIdentityAttributes `json:"entitlementAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchIdentityAccountsInner SearchIdentityAccountsInner

// NewSearchIdentityAccountsInner instantiates a new SearchIdentityAccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIdentityAccountsInner() *SearchIdentityAccountsInner {
	this := SearchIdentityAccountsInner{}
	return &this
}

// NewSearchIdentityAccountsInnerWithDefaults instantiates a new SearchIdentityAccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIdentityAccountsInnerWithDefaults() *SearchIdentityAccountsInner {
	this := SearchIdentityAccountsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchIdentityAccountsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchIdentityAccountsInner) SetName(v string) {
	o.Name = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SearchIdentityAccountsInner) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetSource() SearchIdentityAccountsInnerSource {
	if o == nil || isNil(o.Source) {
		var ret SearchIdentityAccountsInnerSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetSourceOk() (*SearchIdentityAccountsInnerSource, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given SearchIdentityAccountsInnerSource and assigns it to the Source field.
func (o *SearchIdentityAccountsInner) SetSource(v SearchIdentityAccountsInnerSource) {
	o.Source = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *SearchIdentityAccountsInner) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *SearchIdentityAccountsInner) SetLocked(v bool) {
	o.Locked = &v
}

// GetManuallyCorrelated returns the ManuallyCorrelated field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetManuallyCorrelated() bool {
	if o == nil || isNil(o.ManuallyCorrelated) {
		var ret bool
		return ret
	}
	return *o.ManuallyCorrelated
}

// GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetManuallyCorrelatedOk() (*bool, bool) {
	if o == nil || isNil(o.ManuallyCorrelated) {
		return nil, false
	}
	return o.ManuallyCorrelated, true
}

// HasManuallyCorrelated returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasManuallyCorrelated() bool {
	if o != nil && !isNil(o.ManuallyCorrelated) {
		return true
	}

	return false
}

// SetManuallyCorrelated gets a reference to the given bool and assigns it to the ManuallyCorrelated field.
func (o *SearchIdentityAccountsInner) SetManuallyCorrelated(v bool) {
	o.ManuallyCorrelated = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *SearchIdentityAccountsInner) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetPasswordLastSet returns the PasswordLastSet field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetPasswordLastSet() time.Time {
	if o == nil || isNil(o.PasswordLastSet) {
		var ret time.Time
		return ret
	}
	return *o.PasswordLastSet
}

// GetPasswordLastSetOk returns a tuple with the PasswordLastSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetPasswordLastSetOk() (*time.Time, bool) {
	if o == nil || isNil(o.PasswordLastSet) {
		return nil, false
	}
	return o.PasswordLastSet, true
}

// HasPasswordLastSet returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasPasswordLastSet() bool {
	if o != nil && !isNil(o.PasswordLastSet) {
		return true
	}

	return false
}

// SetPasswordLastSet gets a reference to the given time.Time and assigns it to the PasswordLastSet field.
func (o *SearchIdentityAccountsInner) SetPasswordLastSet(v time.Time) {
	o.PasswordLastSet = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SearchIdentityAccountsInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEntitlementAttributes returns the EntitlementAttributes field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInner) GetEntitlementAttributes() SearchIdentityAttributes {
	if o == nil || isNil(o.EntitlementAttributes) {
		var ret SearchIdentityAttributes
		return ret
	}
	return *o.EntitlementAttributes
}

// GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInner) GetEntitlementAttributesOk() (*SearchIdentityAttributes, bool) {
	if o == nil || isNil(o.EntitlementAttributes) {
		return nil, false
	}
	return o.EntitlementAttributes, true
}

// HasEntitlementAttributes returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInner) HasEntitlementAttributes() bool {
	if o != nil && !isNil(o.EntitlementAttributes) {
		return true
	}

	return false
}

// SetEntitlementAttributes gets a reference to the given SearchIdentityAttributes and assigns it to the EntitlementAttributes field.
func (o *SearchIdentityAccountsInner) SetEntitlementAttributes(v SearchIdentityAttributes) {
	o.EntitlementAttributes = &v
}

func (o SearchIdentityAccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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
	if !isNil(o.ManuallyCorrelated) {
		toSerialize["manuallyCorrelated"] = o.ManuallyCorrelated
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.PasswordLastSet) {
		toSerialize["passwordLastSet"] = o.PasswordLastSet
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.EntitlementAttributes) {
		toSerialize["entitlementAttributes"] = o.EntitlementAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchIdentityAccountsInner) UnmarshalJSON(bytes []byte) (err error) {
	varSearchIdentityAccountsInner := _SearchIdentityAccountsInner{}

	if err = json.Unmarshal(bytes, &varSearchIdentityAccountsInner); err == nil {
		*o = SearchIdentityAccountsInner(varSearchIdentityAccountsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "source")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "manuallyCorrelated")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "passwordLastSet")
		delete(additionalProperties, "created")
		delete(additionalProperties, "entitlementAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchIdentityAccountsInner struct {
	value *SearchIdentityAccountsInner
	isSet bool
}

func (v NullableSearchIdentityAccountsInner) Get() *SearchIdentityAccountsInner {
	return v.value
}

func (v *NullableSearchIdentityAccountsInner) Set(val *SearchIdentityAccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIdentityAccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIdentityAccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIdentityAccountsInner(val *SearchIdentityAccountsInner) *NullableSearchIdentityAccountsInner {
	return &NullableSearchIdentityAccountsInner{value: val, isSet: true}
}

func (v NullableSearchIdentityAccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIdentityAccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


