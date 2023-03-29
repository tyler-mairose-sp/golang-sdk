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

// checks if the IdentityReferenceWithId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityReferenceWithId{}

// IdentityReferenceWithId struct for IdentityReferenceWithId
type IdentityReferenceWithId struct {
	Type *DtoType `json:"type,omitempty"`
	// Identity id
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityReferenceWithId IdentityReferenceWithId

// NewIdentityReferenceWithId instantiates a new IdentityReferenceWithId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityReferenceWithId() *IdentityReferenceWithId {
	this := IdentityReferenceWithId{}
	return &this
}

// NewIdentityReferenceWithIdWithDefaults instantiates a new IdentityReferenceWithId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityReferenceWithIdWithDefaults() *IdentityReferenceWithId {
	this := IdentityReferenceWithId{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityReferenceWithId) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReferenceWithId) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityReferenceWithId) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *IdentityReferenceWithId) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityReferenceWithId) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReferenceWithId) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityReferenceWithId) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityReferenceWithId) SetId(v string) {
	o.Id = &v
}

func (o IdentityReferenceWithId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityReferenceWithId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityReferenceWithId) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityReferenceWithId := _IdentityReferenceWithId{}

	if err = json.Unmarshal(bytes, &varIdentityReferenceWithId); err == nil {
		*o = IdentityReferenceWithId(varIdentityReferenceWithId)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityReferenceWithId struct {
	value *IdentityReferenceWithId
	isSet bool
}

func (v NullableIdentityReferenceWithId) Get() *IdentityReferenceWithId {
	return v.value
}

func (v *NullableIdentityReferenceWithId) Set(val *IdentityReferenceWithId) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityReferenceWithId) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityReferenceWithId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityReferenceWithId(val *IdentityReferenceWithId) *NullableIdentityReferenceWithId {
	return &NullableIdentityReferenceWithId{value: val, isSet: true}
}

func (v NullableIdentityReferenceWithId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityReferenceWithId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


