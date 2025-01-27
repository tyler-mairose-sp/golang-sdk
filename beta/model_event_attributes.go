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

// checks if the EventAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventAttributes{}

// EventAttributes struct for EventAttributes
type EventAttributes struct {
	// The unique ID of the trigger
	Id string `json:"id"`
	// JSON path expression that will limit which events the trigger will fire on
	Filter *string `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventAttributes EventAttributes

// NewEventAttributes instantiates a new EventAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventAttributes(id string) *EventAttributes {
	this := EventAttributes{}
	this.Id = id
	return &this
}

// NewEventAttributesWithDefaults instantiates a new EventAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventAttributesWithDefaults() *EventAttributes {
	this := EventAttributes{}
	return &this
}

// GetId returns the Id field value
func (o *EventAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventAttributes) SetId(v string) {
	o.Id = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *EventAttributes) GetFilter() string {
	if o == nil || isNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetFilterOk() (*string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *EventAttributes) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *EventAttributes) SetFilter(v string) {
	o.Filter = &v
}

func (o EventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varEventAttributes := _EventAttributes{}

	if err = json.Unmarshal(bytes, &varEventAttributes); err == nil {
		*o = EventAttributes(varEventAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventAttributes struct {
	value *EventAttributes
	isSet bool
}

func (v NullableEventAttributes) Get() *EventAttributes {
	return v.value
}

func (v *NullableEventAttributes) Set(val *EventAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAttributes(val *EventAttributes) *NullableEventAttributes {
	return &NullableEventAttributes{value: val, isSet: true}
}

func (v NullableEventAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


