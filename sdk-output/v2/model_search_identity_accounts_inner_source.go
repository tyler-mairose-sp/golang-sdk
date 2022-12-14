/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
)

// SearchIdentityAccountsInnerSource 
type SearchIdentityAccountsInnerSource struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchIdentityAccountsInnerSource SearchIdentityAccountsInnerSource

// NewSearchIdentityAccountsInnerSource instantiates a new SearchIdentityAccountsInnerSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIdentityAccountsInnerSource() *SearchIdentityAccountsInnerSource {
	this := SearchIdentityAccountsInnerSource{}
	return &this
}

// NewSearchIdentityAccountsInnerSourceWithDefaults instantiates a new SearchIdentityAccountsInnerSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIdentityAccountsInnerSourceWithDefaults() *SearchIdentityAccountsInnerSource {
	this := SearchIdentityAccountsInnerSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInnerSource) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInnerSource) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInnerSource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchIdentityAccountsInnerSource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInnerSource) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInnerSource) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInnerSource) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchIdentityAccountsInnerSource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SearchIdentityAccountsInnerSource) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAccountsInnerSource) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchIdentityAccountsInnerSource) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SearchIdentityAccountsInnerSource) SetType(v string) {
	o.Type = &v
}

func (o SearchIdentityAccountsInnerSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchIdentityAccountsInnerSource) UnmarshalJSON(bytes []byte) (err error) {
	varSearchIdentityAccountsInnerSource := _SearchIdentityAccountsInnerSource{}

	if err = json.Unmarshal(bytes, &varSearchIdentityAccountsInnerSource); err == nil {
		*o = SearchIdentityAccountsInnerSource(varSearchIdentityAccountsInnerSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchIdentityAccountsInnerSource struct {
	value *SearchIdentityAccountsInnerSource
	isSet bool
}

func (v NullableSearchIdentityAccountsInnerSource) Get() *SearchIdentityAccountsInnerSource {
	return v.value
}

func (v *NullableSearchIdentityAccountsInnerSource) Set(val *SearchIdentityAccountsInnerSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIdentityAccountsInnerSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIdentityAccountsInnerSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIdentityAccountsInnerSource(val *SearchIdentityAccountsInnerSource) *NullableSearchIdentityAccountsInnerSource {
	return &NullableSearchIdentityAccountsInnerSource{value: val, isSet: true}
}

func (v NullableSearchIdentityAccountsInnerSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIdentityAccountsInnerSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


