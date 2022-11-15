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

// ConnectorRuleUpdateRequestAllOf struct for ConnectorRuleUpdateRequestAllOf
type ConnectorRuleUpdateRequestAllOf struct {
	// the ID of the rule to update
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorRuleUpdateRequestAllOf ConnectorRuleUpdateRequestAllOf

// NewConnectorRuleUpdateRequestAllOf instantiates a new ConnectorRuleUpdateRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRuleUpdateRequestAllOf(id string) *ConnectorRuleUpdateRequestAllOf {
	this := ConnectorRuleUpdateRequestAllOf{}
	this.Id = id
	return &this
}

// NewConnectorRuleUpdateRequestAllOfWithDefaults instantiates a new ConnectorRuleUpdateRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRuleUpdateRequestAllOfWithDefaults() *ConnectorRuleUpdateRequestAllOf {
	this := ConnectorRuleUpdateRequestAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectorRuleUpdateRequestAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleUpdateRequestAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectorRuleUpdateRequestAllOf) SetId(v string) {
	o.Id = v
}

func (o ConnectorRuleUpdateRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorRuleUpdateRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorRuleUpdateRequestAllOf := _ConnectorRuleUpdateRequestAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorRuleUpdateRequestAllOf); err == nil {
		*o = ConnectorRuleUpdateRequestAllOf(varConnectorRuleUpdateRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorRuleUpdateRequestAllOf struct {
	value *ConnectorRuleUpdateRequestAllOf
	isSet bool
}

func (v NullableConnectorRuleUpdateRequestAllOf) Get() *ConnectorRuleUpdateRequestAllOf {
	return v.value
}

func (v *NullableConnectorRuleUpdateRequestAllOf) Set(val *ConnectorRuleUpdateRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRuleUpdateRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRuleUpdateRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRuleUpdateRequestAllOf(val *ConnectorRuleUpdateRequestAllOf) *NullableConnectorRuleUpdateRequestAllOf {
	return &NullableConnectorRuleUpdateRequestAllOf{value: val, isSet: true}
}

func (v NullableConnectorRuleUpdateRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRuleUpdateRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


