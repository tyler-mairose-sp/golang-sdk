/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
	"time"
)

// SearchEvent struct for SearchEvent
type SearchEvent struct {
	Id *string `json:"id,omitempty"`
	Action *string `json:"action,omitempty"`
	Type *string `json:"type,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Synced *time.Time `json:"synced,omitempty"`
	Actor *SearchEventActor `json:"actor,omitempty"`
	Target *SearchEventActor `json:"target,omitempty"`
	Stack *string `json:"stack,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Attributes *SearchEventAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchEvent SearchEvent

// NewSearchEvent instantiates a new SearchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEvent() *SearchEvent {
	this := SearchEvent{}
	return &this
}

// NewSearchEventWithDefaults instantiates a new SearchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEventWithDefaults() *SearchEvent {
	this := SearchEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchEvent) SetId(v string) {
	o.Id = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SearchEvent) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SearchEvent) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SearchEvent) SetAction(v string) {
	o.Action = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SearchEvent) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchEvent) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SearchEvent) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SearchEvent) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SearchEvent) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SearchEvent) SetCreated(v time.Time) {
	o.Created = &v
}

// GetSynced returns the Synced field value if set, zero value otherwise.
func (o *SearchEvent) GetSynced() time.Time {
	if o == nil || isNil(o.Synced) {
		var ret time.Time
		return ret
	}
	return *o.Synced
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetSyncedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Synced) {
		return nil, false
	}
	return o.Synced, true
}

// HasSynced returns a boolean if a field has been set.
func (o *SearchEvent) HasSynced() bool {
	if o != nil && !isNil(o.Synced) {
		return true
	}

	return false
}

// SetSynced gets a reference to the given time.Time and assigns it to the Synced field.
func (o *SearchEvent) SetSynced(v time.Time) {
	o.Synced = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *SearchEvent) GetActor() SearchEventActor {
	if o == nil || isNil(o.Actor) {
		var ret SearchEventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetActorOk() (*SearchEventActor, bool) {
	if o == nil || isNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *SearchEvent) HasActor() bool {
	if o != nil && !isNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given SearchEventActor and assigns it to the Actor field.
func (o *SearchEvent) SetActor(v SearchEventActor) {
	o.Actor = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SearchEvent) GetTarget() SearchEventActor {
	if o == nil || isNil(o.Target) {
		var ret SearchEventActor
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetTargetOk() (*SearchEventActor, bool) {
	if o == nil || isNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SearchEvent) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given SearchEventActor and assigns it to the Target field.
func (o *SearchEvent) SetTarget(v SearchEventActor) {
	o.Target = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *SearchEvent) GetStack() string {
	if o == nil || isNil(o.Stack) {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetStackOk() (*string, bool) {
	if o == nil || isNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *SearchEvent) HasStack() bool {
	if o != nil && !isNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *SearchEvent) SetStack(v string) {
	o.Stack = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SearchEvent) GetRequestId() string {
	if o == nil || isNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SearchEvent) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SearchEvent) SetRequestId(v string) {
	o.RequestId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SearchEvent) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SearchEvent) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SearchEvent) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SearchEvent) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SearchEvent) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SearchEvent) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SearchEvent) GetAttributes() SearchEventAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret SearchEventAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEvent) GetAttributesOk() (*SearchEventAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SearchEvent) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SearchEventAttributes and assigns it to the Attributes field.
func (o *SearchEvent) SetAttributes(v SearchEventAttributes) {
	o.Attributes = &v
}

func (o SearchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Synced) {
		toSerialize["synced"] = o.Synced
	}
	if !isNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	if !isNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchEvent) UnmarshalJSON(bytes []byte) (err error) {
	varSearchEvent := _SearchEvent{}

	if err = json.Unmarshal(bytes, &varSearchEvent); err == nil {
		*o = SearchEvent(varSearchEvent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "action")
		delete(additionalProperties, "type")
		delete(additionalProperties, "created")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "target")
		delete(additionalProperties, "stack")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchEvent struct {
	value *SearchEvent
	isSet bool
}

func (v NullableSearchEvent) Get() *SearchEvent {
	return v.value
}

func (v *NullableSearchEvent) Set(val *SearchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEvent(val *SearchEvent) *NullableSearchEvent {
	return &NullableSearchEvent{value: val, isSet: true}
}

func (v NullableSearchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


