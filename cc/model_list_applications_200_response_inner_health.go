/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the ListApplications200ResponseInnerHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplications200ResponseInnerHealth{}

// ListApplications200ResponseInnerHealth struct for ListApplications200ResponseInnerHealth
type ListApplications200ResponseInnerHealth struct {
	Status *string `json:"status,omitempty"`
	LastChanged *string `json:"lastChanged,omitempty"`
	Since *float32 `json:"since,omitempty"`
	Healthy *bool `json:"healthy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListApplications200ResponseInnerHealth ListApplications200ResponseInnerHealth

// NewListApplications200ResponseInnerHealth instantiates a new ListApplications200ResponseInnerHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplications200ResponseInnerHealth() *ListApplications200ResponseInnerHealth {
	this := ListApplications200ResponseInnerHealth{}
	return &this
}

// NewListApplications200ResponseInnerHealthWithDefaults instantiates a new ListApplications200ResponseInnerHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplications200ResponseInnerHealthWithDefaults() *ListApplications200ResponseInnerHealth {
	this := ListApplications200ResponseInnerHealth{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerHealth) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerHealth) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerHealth) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListApplications200ResponseInnerHealth) SetStatus(v string) {
	o.Status = &v
}

// GetLastChanged returns the LastChanged field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerHealth) GetLastChanged() string {
	if o == nil || isNil(o.LastChanged) {
		var ret string
		return ret
	}
	return *o.LastChanged
}

// GetLastChangedOk returns a tuple with the LastChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerHealth) GetLastChangedOk() (*string, bool) {
	if o == nil || isNil(o.LastChanged) {
		return nil, false
	}
	return o.LastChanged, true
}

// HasLastChanged returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerHealth) HasLastChanged() bool {
	if o != nil && !isNil(o.LastChanged) {
		return true
	}

	return false
}

// SetLastChanged gets a reference to the given string and assigns it to the LastChanged field.
func (o *ListApplications200ResponseInnerHealth) SetLastChanged(v string) {
	o.LastChanged = &v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerHealth) GetSince() float32 {
	if o == nil || isNil(o.Since) {
		var ret float32
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerHealth) GetSinceOk() (*float32, bool) {
	if o == nil || isNil(o.Since) {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerHealth) HasSince() bool {
	if o != nil && !isNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given float32 and assigns it to the Since field.
func (o *ListApplications200ResponseInnerHealth) SetSince(v float32) {
	o.Since = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerHealth) GetHealthy() bool {
	if o == nil || isNil(o.Healthy) {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerHealth) GetHealthyOk() (*bool, bool) {
	if o == nil || isNil(o.Healthy) {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerHealth) HasHealthy() bool {
	if o != nil && !isNil(o.Healthy) {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *ListApplications200ResponseInnerHealth) SetHealthy(v bool) {
	o.Healthy = &v
}

func (o ListApplications200ResponseInnerHealth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplications200ResponseInnerHealth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.LastChanged) {
		toSerialize["lastChanged"] = o.LastChanged
	}
	if !isNil(o.Since) {
		toSerialize["since"] = o.Since
	}
	if !isNil(o.Healthy) {
		toSerialize["healthy"] = o.Healthy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListApplications200ResponseInnerHealth) UnmarshalJSON(bytes []byte) (err error) {
	varListApplications200ResponseInnerHealth := _ListApplications200ResponseInnerHealth{}

	if err = json.Unmarshal(bytes, &varListApplications200ResponseInnerHealth); err == nil {
		*o = ListApplications200ResponseInnerHealth(varListApplications200ResponseInnerHealth)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastChanged")
		delete(additionalProperties, "since")
		delete(additionalProperties, "healthy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListApplications200ResponseInnerHealth struct {
	value *ListApplications200ResponseInnerHealth
	isSet bool
}

func (v NullableListApplications200ResponseInnerHealth) Get() *ListApplications200ResponseInnerHealth {
	return v.value
}

func (v *NullableListApplications200ResponseInnerHealth) Set(val *ListApplications200ResponseInnerHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInnerHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInnerHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInnerHealth(val *ListApplications200ResponseInnerHealth) *NullableListApplications200ResponseInnerHealth {
	return &NullableListApplications200ResponseInnerHealth{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInnerHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInnerHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


