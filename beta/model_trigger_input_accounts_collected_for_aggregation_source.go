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

// checks if the TriggerInputAccountsCollectedForAggregationSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputAccountsCollectedForAggregationSource{}

// TriggerInputAccountsCollectedForAggregationSource Reference to the source that has been aggregated.
type TriggerInputAccountsCollectedForAggregationSource struct {
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// The type of object that is referenced
	Type string `json:"type"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountsCollectedForAggregationSource TriggerInputAccountsCollectedForAggregationSource

// NewTriggerInputAccountsCollectedForAggregationSource instantiates a new TriggerInputAccountsCollectedForAggregationSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountsCollectedForAggregationSource(id string, type_ string, name string) *TriggerInputAccountsCollectedForAggregationSource {
	this := TriggerInputAccountsCollectedForAggregationSource{}
	this.Id = id
	this.Type = type_
	this.Name = name
	return &this
}

// NewTriggerInputAccountsCollectedForAggregationSourceWithDefaults instantiates a new TriggerInputAccountsCollectedForAggregationSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountsCollectedForAggregationSourceWithDefaults() *TriggerInputAccountsCollectedForAggregationSource {
	this := TriggerInputAccountsCollectedForAggregationSource{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputAccountsCollectedForAggregationSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountsCollectedForAggregationSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccountsCollectedForAggregationSource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TriggerInputAccountsCollectedForAggregationSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountsCollectedForAggregationSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccountsCollectedForAggregationSource) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TriggerInputAccountsCollectedForAggregationSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountsCollectedForAggregationSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccountsCollectedForAggregationSource) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccountsCollectedForAggregationSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputAccountsCollectedForAggregationSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputAccountsCollectedForAggregationSource) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountsCollectedForAggregationSource := _TriggerInputAccountsCollectedForAggregationSource{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountsCollectedForAggregationSource); err == nil {
		*o = TriggerInputAccountsCollectedForAggregationSource(varTriggerInputAccountsCollectedForAggregationSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountsCollectedForAggregationSource struct {
	value *TriggerInputAccountsCollectedForAggregationSource
	isSet bool
}

func (v NullableTriggerInputAccountsCollectedForAggregationSource) Get() *TriggerInputAccountsCollectedForAggregationSource {
	return v.value
}

func (v *NullableTriggerInputAccountsCollectedForAggregationSource) Set(val *TriggerInputAccountsCollectedForAggregationSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountsCollectedForAggregationSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountsCollectedForAggregationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountsCollectedForAggregationSource(val *TriggerInputAccountsCollectedForAggregationSource) *NullableTriggerInputAccountsCollectedForAggregationSource {
	return &NullableTriggerInputAccountsCollectedForAggregationSource{value: val, isSet: true}
}

func (v NullableTriggerInputAccountsCollectedForAggregationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountsCollectedForAggregationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


