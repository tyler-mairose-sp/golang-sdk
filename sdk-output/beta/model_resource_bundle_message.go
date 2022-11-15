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

// ResourceBundleMessage struct for ResourceBundleMessage
type ResourceBundleMessage struct {
	// The key of the message
	Key *string `json:"key,omitempty"`
	// The format of the message
	Format *string `json:"format,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceBundleMessage ResourceBundleMessage

// NewResourceBundleMessage instantiates a new ResourceBundleMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceBundleMessage() *ResourceBundleMessage {
	this := ResourceBundleMessage{}
	return &this
}

// NewResourceBundleMessageWithDefaults instantiates a new ResourceBundleMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceBundleMessageWithDefaults() *ResourceBundleMessage {
	this := ResourceBundleMessage{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ResourceBundleMessage) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleMessage) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ResourceBundleMessage) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ResourceBundleMessage) SetKey(v string) {
	o.Key = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ResourceBundleMessage) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleMessage) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ResourceBundleMessage) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ResourceBundleMessage) SetFormat(v string) {
	o.Format = &v
}

func (o ResourceBundleMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceBundleMessage) UnmarshalJSON(bytes []byte) (err error) {
	varResourceBundleMessage := _ResourceBundleMessage{}

	if err = json.Unmarshal(bytes, &varResourceBundleMessage); err == nil {
		*o = ResourceBundleMessage(varResourceBundleMessage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "format")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceBundleMessage struct {
	value *ResourceBundleMessage
	isSet bool
}

func (v NullableResourceBundleMessage) Get() *ResourceBundleMessage {
	return v.value
}

func (v *NullableResourceBundleMessage) Set(val *ResourceBundleMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceBundleMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceBundleMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceBundleMessage(val *ResourceBundleMessage) *NullableResourceBundleMessage {
	return &NullableResourceBundleMessage{value: val, isSet: true}
}

func (v NullableResourceBundleMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceBundleMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


