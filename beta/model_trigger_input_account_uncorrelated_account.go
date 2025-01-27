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

// checks if the TriggerInputAccountUncorrelatedAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputAccountUncorrelatedAccount{}

// TriggerInputAccountUncorrelatedAccount The account that was uncorrelated.
type TriggerInputAccountUncorrelatedAccount struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type"`
	// Unique ID of the account on the source.
	NativeIdentity string `json:"nativeIdentity"`
	// The source's unique identifier for the account. UUID is generated by the source system.
	Uuid NullableString `json:"uuid,omitempty"`
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountUncorrelatedAccount TriggerInputAccountUncorrelatedAccount

// NewTriggerInputAccountUncorrelatedAccount instantiates a new TriggerInputAccountUncorrelatedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountUncorrelatedAccount(type_ map[string]interface{}, nativeIdentity string, id string, name string) *TriggerInputAccountUncorrelatedAccount {
	this := TriggerInputAccountUncorrelatedAccount{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewTriggerInputAccountUncorrelatedAccountWithDefaults instantiates a new TriggerInputAccountUncorrelatedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountUncorrelatedAccountWithDefaults() *TriggerInputAccountUncorrelatedAccount {
	this := TriggerInputAccountUncorrelatedAccount{}
	return &this
}

// GetType returns the Type field value
func (o *TriggerInputAccountUncorrelatedAccount) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountUncorrelatedAccount) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccountUncorrelatedAccount) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetNativeIdentity returns the NativeIdentity field value
func (o *TriggerInputAccountUncorrelatedAccount) GetNativeIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountUncorrelatedAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentity, true
}

// SetNativeIdentity sets field value
func (o *TriggerInputAccountUncorrelatedAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputAccountUncorrelatedAccount) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccountUncorrelatedAccount) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *TriggerInputAccountUncorrelatedAccount) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *TriggerInputAccountUncorrelatedAccount) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *TriggerInputAccountUncorrelatedAccount) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *TriggerInputAccountUncorrelatedAccount) UnsetUuid() {
	o.Uuid.Unset()
}

// GetId returns the Id field value
func (o *TriggerInputAccountUncorrelatedAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountUncorrelatedAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccountUncorrelatedAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputAccountUncorrelatedAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountUncorrelatedAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccountUncorrelatedAccount) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccountUncorrelatedAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputAccountUncorrelatedAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["nativeIdentity"] = o.NativeIdentity
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputAccountUncorrelatedAccount) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountUncorrelatedAccount := _TriggerInputAccountUncorrelatedAccount{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountUncorrelatedAccount); err == nil {
		*o = TriggerInputAccountUncorrelatedAccount(varTriggerInputAccountUncorrelatedAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountUncorrelatedAccount struct {
	value *TriggerInputAccountUncorrelatedAccount
	isSet bool
}

func (v NullableTriggerInputAccountUncorrelatedAccount) Get() *TriggerInputAccountUncorrelatedAccount {
	return v.value
}

func (v *NullableTriggerInputAccountUncorrelatedAccount) Set(val *TriggerInputAccountUncorrelatedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountUncorrelatedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountUncorrelatedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountUncorrelatedAccount(val *TriggerInputAccountUncorrelatedAccount) *NullableTriggerInputAccountUncorrelatedAccount {
	return &NullableTriggerInputAccountUncorrelatedAccount{value: val, isSet: true}
}

func (v NullableTriggerInputAccountUncorrelatedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountUncorrelatedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


