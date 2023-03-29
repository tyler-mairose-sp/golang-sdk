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

// checks if the TriggerInputVAClusterStatusChangeEventApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputVAClusterStatusChangeEventApplication{}

// TriggerInputVAClusterStatusChangeEventApplication Details about the `CLUSTER` or `SOURCE` that initiated this event.
type TriggerInputVAClusterStatusChangeEventApplication struct {
	// The GUID of the application
	Id string `json:"id"`
	// The name of the application
	Name string `json:"name"`
	// Custom map of attributes for a source.  This will only be populated if type is `SOURCE` and the source has a proxy.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputVAClusterStatusChangeEventApplication TriggerInputVAClusterStatusChangeEventApplication

// NewTriggerInputVAClusterStatusChangeEventApplication instantiates a new TriggerInputVAClusterStatusChangeEventApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputVAClusterStatusChangeEventApplication(id string, name string, attributes map[string]interface{}) *TriggerInputVAClusterStatusChangeEventApplication {
	this := TriggerInputVAClusterStatusChangeEventApplication{}
	this.Id = id
	this.Name = name
	this.Attributes = attributes
	return &this
}

// NewTriggerInputVAClusterStatusChangeEventApplicationWithDefaults instantiates a new TriggerInputVAClusterStatusChangeEventApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputVAClusterStatusChangeEventApplicationWithDefaults() *TriggerInputVAClusterStatusChangeEventApplication {
	this := TriggerInputVAClusterStatusChangeEventApplication{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputVAClusterStatusChangeEventApplication) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputVAClusterStatusChangeEventApplication) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputVAClusterStatusChangeEventApplication) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputVAClusterStatusChangeEventApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputVAClusterStatusChangeEventApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputVAClusterStatusChangeEventApplication) SetName(v string) {
	o.Name = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *TriggerInputVAClusterStatusChangeEventApplication) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputVAClusterStatusChangeEventApplication) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TriggerInputVAClusterStatusChangeEventApplication) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TriggerInputVAClusterStatusChangeEventApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputVAClusterStatusChangeEventApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputVAClusterStatusChangeEventApplication) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputVAClusterStatusChangeEventApplication := _TriggerInputVAClusterStatusChangeEventApplication{}

	if err = json.Unmarshal(bytes, &varTriggerInputVAClusterStatusChangeEventApplication); err == nil {
		*o = TriggerInputVAClusterStatusChangeEventApplication(varTriggerInputVAClusterStatusChangeEventApplication)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputVAClusterStatusChangeEventApplication struct {
	value *TriggerInputVAClusterStatusChangeEventApplication
	isSet bool
}

func (v NullableTriggerInputVAClusterStatusChangeEventApplication) Get() *TriggerInputVAClusterStatusChangeEventApplication {
	return v.value
}

func (v *NullableTriggerInputVAClusterStatusChangeEventApplication) Set(val *TriggerInputVAClusterStatusChangeEventApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputVAClusterStatusChangeEventApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputVAClusterStatusChangeEventApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputVAClusterStatusChangeEventApplication(val *TriggerInputVAClusterStatusChangeEventApplication) *NullableTriggerInputVAClusterStatusChangeEventApplication {
	return &NullableTriggerInputVAClusterStatusChangeEventApplication{value: val, isSet: true}
}

func (v NullableTriggerInputVAClusterStatusChangeEventApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputVAClusterStatusChangeEventApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


