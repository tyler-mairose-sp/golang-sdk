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

// checks if the ReassignReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReassignReference{}

// ReassignReference struct for ReassignReference
type ReassignReference struct {
	// The ID of item or identity being reassigned.
	Id string `json:"id"`
	// The type of item or identity being reassigned.
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ReassignReference ReassignReference

// NewReassignReference instantiates a new ReassignReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReassignReference(id string, type_ string) *ReassignReference {
	this := ReassignReference{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewReassignReferenceWithDefaults instantiates a new ReassignReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReassignReferenceWithDefaults() *ReassignReference {
	this := ReassignReference{}
	return &this
}

// GetId returns the Id field value
func (o *ReassignReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReassignReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReassignReference) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ReassignReference) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReassignReference) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReassignReference) SetType(v string) {
	o.Type = v
}

func (o ReassignReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReassignReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReassignReference) UnmarshalJSON(bytes []byte) (err error) {
	varReassignReference := _ReassignReference{}

	if err = json.Unmarshal(bytes, &varReassignReference); err == nil {
		*o = ReassignReference(varReassignReference)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReassignReference struct {
	value *ReassignReference
	isSet bool
}

func (v NullableReassignReference) Get() *ReassignReference {
	return v.value
}

func (v *NullableReassignReference) Set(val *ReassignReference) {
	v.value = val
	v.isSet = true
}

func (v NullableReassignReference) IsSet() bool {
	return v.isSet
}

func (v *NullableReassignReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReassignReference(val *ReassignReference) *NullableReassignReference {
	return &NullableReassignReference{value: val, isSet: true}
}

func (v NullableReassignReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReassignReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


