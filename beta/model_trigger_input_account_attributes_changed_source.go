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

// checks if the TriggerInputAccountAttributesChangedSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputAccountAttributesChangedSource{}

// TriggerInputAccountAttributesChangedSource The source that contains the account.
type TriggerInputAccountAttributesChangedSource struct {
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// The type of object that is referenced
	Type string `json:"type"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountAttributesChangedSource TriggerInputAccountAttributesChangedSource

// NewTriggerInputAccountAttributesChangedSource instantiates a new TriggerInputAccountAttributesChangedSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountAttributesChangedSource(id string, type_ string, name string) *TriggerInputAccountAttributesChangedSource {
	this := TriggerInputAccountAttributesChangedSource{}
	this.Id = id
	this.Type = type_
	this.Name = name
	return &this
}

// NewTriggerInputAccountAttributesChangedSourceWithDefaults instantiates a new TriggerInputAccountAttributesChangedSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountAttributesChangedSourceWithDefaults() *TriggerInputAccountAttributesChangedSource {
	this := TriggerInputAccountAttributesChangedSource{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputAccountAttributesChangedSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccountAttributesChangedSource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TriggerInputAccountAttributesChangedSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccountAttributesChangedSource) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TriggerInputAccountAttributesChangedSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccountAttributesChangedSource) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccountAttributesChangedSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputAccountAttributesChangedSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputAccountAttributesChangedSource) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountAttributesChangedSource := _TriggerInputAccountAttributesChangedSource{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountAttributesChangedSource); err == nil {
		*o = TriggerInputAccountAttributesChangedSource(varTriggerInputAccountAttributesChangedSource)
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

type NullableTriggerInputAccountAttributesChangedSource struct {
	value *TriggerInputAccountAttributesChangedSource
	isSet bool
}

func (v NullableTriggerInputAccountAttributesChangedSource) Get() *TriggerInputAccountAttributesChangedSource {
	return v.value
}

func (v *NullableTriggerInputAccountAttributesChangedSource) Set(val *TriggerInputAccountAttributesChangedSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountAttributesChangedSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountAttributesChangedSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountAttributesChangedSource(val *TriggerInputAccountAttributesChangedSource) *NullableTriggerInputAccountAttributesChangedSource {
	return &NullableTriggerInputAccountAttributesChangedSource{value: val, isSet: true}
}

func (v NullableTriggerInputAccountAttributesChangedSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountAttributesChangedSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


