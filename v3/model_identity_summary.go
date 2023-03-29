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

// checks if the IdentitySummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySummary{}

// IdentitySummary struct for IdentitySummary
type IdentitySummary struct {
	// ID of this identity summary
	Id *string `json:"id,omitempty"`
	// Human-readable display name of identity
	Name *string `json:"name,omitempty"`
	// ID of the identity that this summary represents
	IdentityId *string `json:"identityId,omitempty"`
	// Indicates if all access items for this summary have been decided on
	Completed *bool `json:"completed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySummary IdentitySummary

// NewIdentitySummary instantiates a new IdentitySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySummary() *IdentitySummary {
	this := IdentitySummary{}
	return &this
}

// NewIdentitySummaryWithDefaults instantiates a new IdentitySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySummaryWithDefaults() *IdentitySummary {
	this := IdentitySummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentitySummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentitySummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentitySummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentitySummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentitySummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentitySummary) SetName(v string) {
	o.Name = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *IdentitySummary) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySummary) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *IdentitySummary) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *IdentitySummary) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *IdentitySummary) GetCompleted() bool {
	if o == nil || isNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySummary) GetCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *IdentitySummary) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *IdentitySummary) SetCompleted(v bool) {
	o.Completed = &v
}

func (o IdentitySummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySummary) UnmarshalJSON(bytes []byte) (err error) {
	varIdentitySummary := _IdentitySummary{}

	if err = json.Unmarshal(bytes, &varIdentitySummary); err == nil {
		*o = IdentitySummary(varIdentitySummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "completed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySummary struct {
	value *IdentitySummary
	isSet bool
}

func (v NullableIdentitySummary) Get() *IdentitySummary {
	return v.value
}

func (v *NullableIdentitySummary) Set(val *IdentitySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySummary(val *IdentitySummary) *NullableIdentitySummary {
	return &NullableIdentitySummary{value: val, isSet: true}
}

func (v NullableIdentitySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


