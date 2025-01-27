/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"time"
)

// checks if the TriggerInputSourceCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputSourceCreated{}

// TriggerInputSourceCreated struct for TriggerInputSourceCreated
type TriggerInputSourceCreated struct {
	// The unique ID of the source.
	Id string `json:"id"`
	// Human friendly name of the source.
	Name string `json:"name"`
	// The connection type.
	Type string `json:"type"`
	// The date and time the source was created.
	Created time.Time `json:"created"`
	// The connector type used to connect to the source.
	Connector string `json:"connector"`
	Actor TriggerInputSourceCreatedActor `json:"actor"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputSourceCreated TriggerInputSourceCreated

// NewTriggerInputSourceCreated instantiates a new TriggerInputSourceCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputSourceCreated(id string, name string, type_ string, created time.Time, connector string, actor TriggerInputSourceCreatedActor) *TriggerInputSourceCreated {
	this := TriggerInputSourceCreated{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Created = created
	this.Connector = connector
	this.Actor = actor
	return &this
}

// NewTriggerInputSourceCreatedWithDefaults instantiates a new TriggerInputSourceCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputSourceCreatedWithDefaults() *TriggerInputSourceCreated {
	this := TriggerInputSourceCreated{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputSourceCreated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceCreated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputSourceCreated) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputSourceCreated) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceCreated) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputSourceCreated) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TriggerInputSourceCreated) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceCreated) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TriggerInputSourceCreated) SetType(v string) {
	o.Type = v
}

// GetCreated returns the Created field value
func (o *TriggerInputSourceCreated) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceCreated) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TriggerInputSourceCreated) SetCreated(v time.Time) {
	o.Created = v
}

// GetConnector returns the Connector field value
func (o *TriggerInputSourceCreated) GetConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceCreated) GetConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *TriggerInputSourceCreated) SetConnector(v string) {
	o.Connector = v
}

// GetActor returns the Actor field value
func (o *TriggerInputSourceCreated) GetActor() TriggerInputSourceCreatedActor {
	if o == nil {
		var ret TriggerInputSourceCreatedActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceCreated) GetActorOk() (*TriggerInputSourceCreatedActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *TriggerInputSourceCreated) SetActor(v TriggerInputSourceCreatedActor) {
	o.Actor = v
}

func (o TriggerInputSourceCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputSourceCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["created"] = o.Created
	toSerialize["connector"] = o.Connector
	toSerialize["actor"] = o.Actor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputSourceCreated) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputSourceCreated := _TriggerInputSourceCreated{}

	if err = json.Unmarshal(bytes, &varTriggerInputSourceCreated); err == nil {
		*o = TriggerInputSourceCreated(varTriggerInputSourceCreated)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "created")
		delete(additionalProperties, "connector")
		delete(additionalProperties, "actor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputSourceCreated struct {
	value *TriggerInputSourceCreated
	isSet bool
}

func (v NullableTriggerInputSourceCreated) Get() *TriggerInputSourceCreated {
	return v.value
}

func (v *NullableTriggerInputSourceCreated) Set(val *TriggerInputSourceCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputSourceCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputSourceCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputSourceCreated(val *TriggerInputSourceCreated) *NullableTriggerInputSourceCreated {
	return &NullableTriggerInputSourceCreated{value: val, isSet: true}
}

func (v NullableTriggerInputSourceCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputSourceCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


