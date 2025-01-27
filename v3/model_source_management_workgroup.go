/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the SourceManagementWorkgroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceManagementWorkgroup{}

// SourceManagementWorkgroup Reference to Management Workgroup for this Source
type SourceManagementWorkgroup struct {
	// The type of object being referenced
	Type *string `json:"type,omitempty"`
	// ID of the management workgroup
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the management workgroup
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceManagementWorkgroup SourceManagementWorkgroup

// NewSourceManagementWorkgroup instantiates a new SourceManagementWorkgroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceManagementWorkgroup() *SourceManagementWorkgroup {
	this := SourceManagementWorkgroup{}
	return &this
}

// NewSourceManagementWorkgroupWithDefaults instantiates a new SourceManagementWorkgroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceManagementWorkgroupWithDefaults() *SourceManagementWorkgroup {
	this := SourceManagementWorkgroup{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceManagementWorkgroup) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManagementWorkgroup) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceManagementWorkgroup) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceManagementWorkgroup) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceManagementWorkgroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManagementWorkgroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceManagementWorkgroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceManagementWorkgroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceManagementWorkgroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManagementWorkgroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceManagementWorkgroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceManagementWorkgroup) SetName(v string) {
	o.Name = &v
}

func (o SourceManagementWorkgroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceManagementWorkgroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceManagementWorkgroup) UnmarshalJSON(bytes []byte) (err error) {
	varSourceManagementWorkgroup := _SourceManagementWorkgroup{}

	if err = json.Unmarshal(bytes, &varSourceManagementWorkgroup); err == nil {
		*o = SourceManagementWorkgroup(varSourceManagementWorkgroup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceManagementWorkgroup struct {
	value *SourceManagementWorkgroup
	isSet bool
}

func (v NullableSourceManagementWorkgroup) Get() *SourceManagementWorkgroup {
	return v.value
}

func (v *NullableSourceManagementWorkgroup) Set(val *SourceManagementWorkgroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceManagementWorkgroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceManagementWorkgroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceManagementWorkgroup(val *SourceManagementWorkgroup) *NullableSourceManagementWorkgroup {
	return &NullableSourceManagementWorkgroup{value: val, isSet: true}
}

func (v NullableSourceManagementWorkgroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceManagementWorkgroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


