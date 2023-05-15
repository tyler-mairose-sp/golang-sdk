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

// checks if the ManagedClusterAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedClusterAttributes{}

// ManagedClusterAttributes Managed Cluster Attributes for Cluster Configuration. Supported Cluster Types [sqsCluster, spConnectCluster]
type ManagedClusterAttributes struct {
	Queue *ManagedClusterQueue `json:"queue,omitempty"`
	// ManagedCluster keystore for spConnectCluster type
	Keystore NullableString `json:"keystore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClusterAttributes ManagedClusterAttributes

// NewManagedClusterAttributes instantiates a new ManagedClusterAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClusterAttributes() *ManagedClusterAttributes {
	this := ManagedClusterAttributes{}
	return &this
}

// NewManagedClusterAttributesWithDefaults instantiates a new ManagedClusterAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClusterAttributesWithDefaults() *ManagedClusterAttributes {
	this := ManagedClusterAttributes{}
	return &this
}

// GetQueue returns the Queue field value if set, zero value otherwise.
func (o *ManagedClusterAttributes) GetQueue() ManagedClusterQueue {
	if o == nil || isNil(o.Queue) {
		var ret ManagedClusterQueue
		return ret
	}
	return *o.Queue
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedClusterAttributes) GetQueueOk() (*ManagedClusterQueue, bool) {
	if o == nil || isNil(o.Queue) {
		return nil, false
	}
	return o.Queue, true
}

// HasQueue returns a boolean if a field has been set.
func (o *ManagedClusterAttributes) HasQueue() bool {
	if o != nil && !isNil(o.Queue) {
		return true
	}

	return false
}

// SetQueue gets a reference to the given ManagedClusterQueue and assigns it to the Queue field.
func (o *ManagedClusterAttributes) SetQueue(v ManagedClusterQueue) {
	o.Queue = &v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClusterAttributes) GetKeystore() string {
	if o == nil || isNil(o.Keystore.Get()) {
		var ret string
		return ret
	}
	return *o.Keystore.Get()
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClusterAttributes) GetKeystoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keystore.Get(), o.Keystore.IsSet()
}

// HasKeystore returns a boolean if a field has been set.
func (o *ManagedClusterAttributes) HasKeystore() bool {
	if o != nil && o.Keystore.IsSet() {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given NullableString and assigns it to the Keystore field.
func (o *ManagedClusterAttributes) SetKeystore(v string) {
	o.Keystore.Set(&v)
}
// SetKeystoreNil sets the value for Keystore to be an explicit nil
func (o *ManagedClusterAttributes) SetKeystoreNil() {
	o.Keystore.Set(nil)
}

// UnsetKeystore ensures that no value is present for Keystore, not even an explicit nil
func (o *ManagedClusterAttributes) UnsetKeystore() {
	o.Keystore.Unset()
}

func (o ManagedClusterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedClusterAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Queue) {
		toSerialize["queue"] = o.Queue
	}
	if o.Keystore.IsSet() {
		toSerialize["keystore"] = o.Keystore.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedClusterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varManagedClusterAttributes := _ManagedClusterAttributes{}

	if err = json.Unmarshal(bytes, &varManagedClusterAttributes); err == nil {
		*o = ManagedClusterAttributes(varManagedClusterAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queue")
		delete(additionalProperties, "keystore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedClusterAttributes struct {
	value *ManagedClusterAttributes
	isSet bool
}

func (v NullableManagedClusterAttributes) Get() *ManagedClusterAttributes {
	return v.value
}

func (v *NullableManagedClusterAttributes) Set(val *ManagedClusterAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClusterAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClusterAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClusterAttributes(val *ManagedClusterAttributes) *NullableManagedClusterAttributes {
	return &NullableManagedClusterAttributes{value: val, isSet: true}
}

func (v NullableManagedClusterAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClusterAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


