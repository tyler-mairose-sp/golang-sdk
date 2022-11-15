/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// SourceCluster Reference to the associated Cluster
type SourceCluster struct {
	// The type of object being referenced
	Type *string `json:"type,omitempty"`
	// ID of the cluster
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the cluster
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceCluster SourceCluster

// NewSourceCluster instantiates a new SourceCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceCluster() *SourceCluster {
	this := SourceCluster{}
	return &this
}

// NewSourceClusterWithDefaults instantiates a new SourceCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceClusterWithDefaults() *SourceCluster {
	this := SourceCluster{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceCluster) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCluster) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceCluster) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceCluster) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceCluster) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCluster) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceCluster) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceCluster) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceCluster) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCluster) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceCluster) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceCluster) SetName(v string) {
	o.Name = &v
}

func (o SourceCluster) MarshalJSON() ([]byte, error) {
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

	return json.Marshal(toSerialize)
}

func (o *SourceCluster) UnmarshalJSON(bytes []byte) (err error) {
	varSourceCluster := _SourceCluster{}

	if err = json.Unmarshal(bytes, &varSourceCluster); err == nil {
		*o = SourceCluster(varSourceCluster)
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

type NullableSourceCluster struct {
	value *SourceCluster
	isSet bool
}

func (v NullableSourceCluster) Get() *SourceCluster {
	return v.value
}

func (v *NullableSourceCluster) Set(val *SourceCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceCluster(val *SourceCluster) *NullableSourceCluster {
	return &NullableSourceCluster{value: val, isSet: true}
}

func (v NullableSourceCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


