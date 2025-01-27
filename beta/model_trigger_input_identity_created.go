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

// checks if the TriggerInputIdentityCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputIdentityCreated{}

// TriggerInputIdentityCreated struct for TriggerInputIdentityCreated
type TriggerInputIdentityCreated struct {
	Identity TriggerInputIdentityCreatedIdentity `json:"identity"`
	// The attributes assigned to the identity.  Attributes are determined by the identity profile.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputIdentityCreated TriggerInputIdentityCreated

// NewTriggerInputIdentityCreated instantiates a new TriggerInputIdentityCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputIdentityCreated(identity TriggerInputIdentityCreatedIdentity, attributes map[string]interface{}) *TriggerInputIdentityCreated {
	this := TriggerInputIdentityCreated{}
	this.Identity = identity
	this.Attributes = attributes
	return &this
}

// NewTriggerInputIdentityCreatedWithDefaults instantiates a new TriggerInputIdentityCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputIdentityCreatedWithDefaults() *TriggerInputIdentityCreated {
	this := TriggerInputIdentityCreated{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *TriggerInputIdentityCreated) GetIdentity() TriggerInputIdentityCreatedIdentity {
	if o == nil {
		var ret TriggerInputIdentityCreatedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputIdentityCreated) GetIdentityOk() (*TriggerInputIdentityCreatedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *TriggerInputIdentityCreated) SetIdentity(v TriggerInputIdentityCreatedIdentity) {
	o.Identity = v
}

// GetAttributes returns the Attributes field value
func (o *TriggerInputIdentityCreated) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TriggerInputIdentityCreated) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TriggerInputIdentityCreated) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TriggerInputIdentityCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputIdentityCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identity"] = o.Identity
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputIdentityCreated) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputIdentityCreated := _TriggerInputIdentityCreated{}

	if err = json.Unmarshal(bytes, &varTriggerInputIdentityCreated); err == nil {
		*o = TriggerInputIdentityCreated(varTriggerInputIdentityCreated)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputIdentityCreated struct {
	value *TriggerInputIdentityCreated
	isSet bool
}

func (v NullableTriggerInputIdentityCreated) Get() *TriggerInputIdentityCreated {
	return v.value
}

func (v *NullableTriggerInputIdentityCreated) Set(val *TriggerInputIdentityCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputIdentityCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputIdentityCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputIdentityCreated(val *TriggerInputIdentityCreated) *NullableTriggerInputIdentityCreated {
	return &NullableTriggerInputIdentityCreated{value: val, isSet: true}
}

func (v NullableTriggerInputIdentityCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputIdentityCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


