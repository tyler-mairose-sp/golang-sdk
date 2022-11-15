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

// WorkflowTrigger The trigger that starts the workflow
type WorkflowTrigger struct {
	// The trigger type
	Type map[string]interface{} `json:"type"`
	Attributes WorkflowTriggerAttributes `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTrigger WorkflowTrigger

// NewWorkflowTrigger instantiates a new WorkflowTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTrigger(type_ map[string]interface{}, attributes WorkflowTriggerAttributes) *WorkflowTrigger {
	this := WorkflowTrigger{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTriggerWithDefaults() *WorkflowTrigger {
	this := WorkflowTrigger{}
	return &this
}

// GetType returns the Type field value
func (o *WorkflowTrigger) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *WorkflowTrigger) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *WorkflowTrigger) GetAttributes() WorkflowTriggerAttributes {
	if o == nil {
		var ret WorkflowTriggerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetAttributesOk() (*WorkflowTriggerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WorkflowTrigger) SetAttributes(v WorkflowTriggerAttributes) {
	o.Attributes = v
}

func (o WorkflowTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTrigger) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTrigger := _WorkflowTrigger{}

	if err = json.Unmarshal(bytes, &varWorkflowTrigger); err == nil {
		*o = WorkflowTrigger(varWorkflowTrigger)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTrigger struct {
	value *WorkflowTrigger
	isSet bool
}

func (v NullableWorkflowTrigger) Get() *WorkflowTrigger {
	return v.value
}

func (v *NullableWorkflowTrigger) Set(val *WorkflowTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTrigger(val *WorkflowTrigger) *NullableWorkflowTrigger {
	return &NullableWorkflowTrigger{value: val, isSet: true}
}

func (v NullableWorkflowTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


