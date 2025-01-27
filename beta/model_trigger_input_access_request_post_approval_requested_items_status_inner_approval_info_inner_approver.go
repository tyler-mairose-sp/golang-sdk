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

// checks if the TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}

// TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover The identity of the approver.
type TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type"`
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover

// NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover instantiates a new TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover(type_ map[string]interface{}, id string, name string) *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	this := TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApproverWithDefaults instantiates a new TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApproverWithDefaults() *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	this := TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}
	return &this
}

// GetType returns the Type field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover := _TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover); err == nil {
		*o = TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover(varTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover)
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

type NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover struct {
	value *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover
	isSet bool
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) Get() *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	return v.value
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) Set(val *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover(val *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover {
	return &NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover{value: val, isSet: true}
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


