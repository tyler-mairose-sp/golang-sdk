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

// checks if the TriggerInputAccountCorrelated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputAccountCorrelated{}

// TriggerInputAccountCorrelated struct for TriggerInputAccountCorrelated
type TriggerInputAccountCorrelated struct {
	Identity TriggerInputAccountCorrelatedIdentity `json:"identity"`
	Source TriggerInputAccountCorrelatedSource `json:"source"`
	Account TriggerInputAccountCorrelatedAccount `json:"account"`
	// The attributes associated with the account.  Attributes are unique per source.
	Attributes map[string]interface{} `json:"attributes"`
	// The number of entitlements associated with this account.
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountCorrelated TriggerInputAccountCorrelated

// NewTriggerInputAccountCorrelated instantiates a new TriggerInputAccountCorrelated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountCorrelated(identity TriggerInputAccountCorrelatedIdentity, source TriggerInputAccountCorrelatedSource, account TriggerInputAccountCorrelatedAccount, attributes map[string]interface{}) *TriggerInputAccountCorrelated {
	this := TriggerInputAccountCorrelated{}
	this.Identity = identity
	this.Source = source
	this.Account = account
	this.Attributes = attributes
	return &this
}

// NewTriggerInputAccountCorrelatedWithDefaults instantiates a new TriggerInputAccountCorrelated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountCorrelatedWithDefaults() *TriggerInputAccountCorrelated {
	this := TriggerInputAccountCorrelated{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *TriggerInputAccountCorrelated) GetIdentity() TriggerInputAccountCorrelatedIdentity {
	if o == nil {
		var ret TriggerInputAccountCorrelatedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelated) GetIdentityOk() (*TriggerInputAccountCorrelatedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *TriggerInputAccountCorrelated) SetIdentity(v TriggerInputAccountCorrelatedIdentity) {
	o.Identity = v
}

// GetSource returns the Source field value
func (o *TriggerInputAccountCorrelated) GetSource() TriggerInputAccountCorrelatedSource {
	if o == nil {
		var ret TriggerInputAccountCorrelatedSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelated) GetSourceOk() (*TriggerInputAccountCorrelatedSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TriggerInputAccountCorrelated) SetSource(v TriggerInputAccountCorrelatedSource) {
	o.Source = v
}

// GetAccount returns the Account field value
func (o *TriggerInputAccountCorrelated) GetAccount() TriggerInputAccountCorrelatedAccount {
	if o == nil {
		var ret TriggerInputAccountCorrelatedAccount
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelated) GetAccountOk() (*TriggerInputAccountCorrelatedAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *TriggerInputAccountCorrelated) SetAccount(v TriggerInputAccountCorrelatedAccount) {
	o.Account = v
}

// GetAttributes returns the Attributes field value
func (o *TriggerInputAccountCorrelated) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelated) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TriggerInputAccountCorrelated) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *TriggerInputAccountCorrelated) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelated) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *TriggerInputAccountCorrelated) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *TriggerInputAccountCorrelated) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

func (o TriggerInputAccountCorrelated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputAccountCorrelated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identity"] = o.Identity
	toSerialize["source"] = o.Source
	toSerialize["account"] = o.Account
	toSerialize["attributes"] = o.Attributes
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputAccountCorrelated) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountCorrelated := _TriggerInputAccountCorrelated{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountCorrelated); err == nil {
		*o = TriggerInputAccountCorrelated(varTriggerInputAccountCorrelated)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "account")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "entitlementCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountCorrelated struct {
	value *TriggerInputAccountCorrelated
	isSet bool
}

func (v NullableTriggerInputAccountCorrelated) Get() *TriggerInputAccountCorrelated {
	return v.value
}

func (v *NullableTriggerInputAccountCorrelated) Set(val *TriggerInputAccountCorrelated) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountCorrelated) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountCorrelated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountCorrelated(val *TriggerInputAccountCorrelated) *NullableTriggerInputAccountCorrelated {
	return &NullableTriggerInputAccountCorrelated{value: val, isSet: true}
}

func (v NullableTriggerInputAccountCorrelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountCorrelated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


