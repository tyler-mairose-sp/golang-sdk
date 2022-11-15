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

// SubscriptionPutRequest struct for SubscriptionPutRequest
type SubscriptionPutRequest struct {
	// Subscription name.
	Name *string `json:"name,omitempty"`
	// Subscription description.
	Description *string `json:"description,omitempty"`
	Type *SubscriptionType `json:"type,omitempty"`
	// Deadline for completing REQUEST_RESPONSE trigger invocation, represented in ISO-8601 duration format.
	ResponseDeadline *string `json:"responseDeadline,omitempty"`
	HttpConfig *HttpConfig `json:"httpConfig,omitempty"`
	EventBridgeConfig *EventBridgeConfig `json:"eventBridgeConfig,omitempty"`
	// Whether subscription should receive real-time trigger invocations or not.  Test trigger invocations are always enabled regardless of this option.
	Enabled *bool `json:"enabled,omitempty"`
	// JSONPath filter to conditionally invoke trigger when expression evaluates to true.
	Filter *string `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionPutRequest SubscriptionPutRequest

// NewSubscriptionPutRequest instantiates a new SubscriptionPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPutRequest() *SubscriptionPutRequest {
	this := SubscriptionPutRequest{}
	var responseDeadline string = "PT1H"
	this.ResponseDeadline = &responseDeadline
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewSubscriptionPutRequestWithDefaults instantiates a new SubscriptionPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPutRequestWithDefaults() *SubscriptionPutRequest {
	this := SubscriptionPutRequest{}
	var responseDeadline string = "PT1H"
	this.ResponseDeadline = &responseDeadline
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriptionPutRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubscriptionPutRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetType() SubscriptionType {
	if o == nil || isNil(o.Type) {
		var ret SubscriptionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetTypeOk() (*SubscriptionType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SubscriptionType and assigns it to the Type field.
func (o *SubscriptionPutRequest) SetType(v SubscriptionType) {
	o.Type = &v
}

// GetResponseDeadline returns the ResponseDeadline field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetResponseDeadline() string {
	if o == nil || isNil(o.ResponseDeadline) {
		var ret string
		return ret
	}
	return *o.ResponseDeadline
}

// GetResponseDeadlineOk returns a tuple with the ResponseDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetResponseDeadlineOk() (*string, bool) {
	if o == nil || isNil(o.ResponseDeadline) {
		return nil, false
	}
	return o.ResponseDeadline, true
}

// HasResponseDeadline returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasResponseDeadline() bool {
	if o != nil && !isNil(o.ResponseDeadline) {
		return true
	}

	return false
}

// SetResponseDeadline gets a reference to the given string and assigns it to the ResponseDeadline field.
func (o *SubscriptionPutRequest) SetResponseDeadline(v string) {
	o.ResponseDeadline = &v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetHttpConfig() HttpConfig {
	if o == nil || isNil(o.HttpConfig) {
		var ret HttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetHttpConfigOk() (*HttpConfig, bool) {
	if o == nil || isNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasHttpConfig() bool {
	if o != nil && !isNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given HttpConfig and assigns it to the HttpConfig field.
func (o *SubscriptionPutRequest) SetHttpConfig(v HttpConfig) {
	o.HttpConfig = &v
}

// GetEventBridgeConfig returns the EventBridgeConfig field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetEventBridgeConfig() EventBridgeConfig {
	if o == nil || isNil(o.EventBridgeConfig) {
		var ret EventBridgeConfig
		return ret
	}
	return *o.EventBridgeConfig
}

// GetEventBridgeConfigOk returns a tuple with the EventBridgeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetEventBridgeConfigOk() (*EventBridgeConfig, bool) {
	if o == nil || isNil(o.EventBridgeConfig) {
		return nil, false
	}
	return o.EventBridgeConfig, true
}

// HasEventBridgeConfig returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasEventBridgeConfig() bool {
	if o != nil && !isNil(o.EventBridgeConfig) {
		return true
	}

	return false
}

// SetEventBridgeConfig gets a reference to the given EventBridgeConfig and assigns it to the EventBridgeConfig field.
func (o *SubscriptionPutRequest) SetEventBridgeConfig(v EventBridgeConfig) {
	o.EventBridgeConfig = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SubscriptionPutRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SubscriptionPutRequest) GetFilter() string {
	if o == nil || isNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPutRequest) GetFilterOk() (*string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SubscriptionPutRequest) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *SubscriptionPutRequest) SetFilter(v string) {
	o.Filter = &v
}

func (o SubscriptionPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ResponseDeadline) {
		toSerialize["responseDeadline"] = o.ResponseDeadline
	}
	if !isNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	if !isNil(o.EventBridgeConfig) {
		toSerialize["eventBridgeConfig"] = o.EventBridgeConfig
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SubscriptionPutRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSubscriptionPutRequest := _SubscriptionPutRequest{}

	if err = json.Unmarshal(bytes, &varSubscriptionPutRequest); err == nil {
		*o = SubscriptionPutRequest(varSubscriptionPutRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "responseDeadline")
		delete(additionalProperties, "httpConfig")
		delete(additionalProperties, "eventBridgeConfig")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionPutRequest struct {
	value *SubscriptionPutRequest
	isSet bool
}

func (v NullableSubscriptionPutRequest) Get() *SubscriptionPutRequest {
	return v.value
}

func (v *NullableSubscriptionPutRequest) Set(val *SubscriptionPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPutRequest(val *SubscriptionPutRequest) *NullableSubscriptionPutRequest {
	return &NullableSubscriptionPutRequest{value: val, isSet: true}
}

func (v NullableSubscriptionPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


