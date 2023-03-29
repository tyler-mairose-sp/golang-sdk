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

// checks if the CompleteInvocationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteInvocationInput{}

// CompleteInvocationInput struct for CompleteInvocationInput
type CompleteInvocationInput struct {
	LocalizedError *LocalizedMessage `json:"localizedError,omitempty"`
	// Trigger output that completed the invocation. Its schema is defined in the trigger definition.
	Output map[string]interface{} `json:"output,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompleteInvocationInput CompleteInvocationInput

// NewCompleteInvocationInput instantiates a new CompleteInvocationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteInvocationInput() *CompleteInvocationInput {
	this := CompleteInvocationInput{}
	return &this
}

// NewCompleteInvocationInputWithDefaults instantiates a new CompleteInvocationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteInvocationInputWithDefaults() *CompleteInvocationInput {
	this := CompleteInvocationInput{}
	return &this
}

// GetLocalizedError returns the LocalizedError field value if set, zero value otherwise.
func (o *CompleteInvocationInput) GetLocalizedError() LocalizedMessage {
	if o == nil || isNil(o.LocalizedError) {
		var ret LocalizedMessage
		return ret
	}
	return *o.LocalizedError
}

// GetLocalizedErrorOk returns a tuple with the LocalizedError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteInvocationInput) GetLocalizedErrorOk() (*LocalizedMessage, bool) {
	if o == nil || isNil(o.LocalizedError) {
		return nil, false
	}
	return o.LocalizedError, true
}

// HasLocalizedError returns a boolean if a field has been set.
func (o *CompleteInvocationInput) HasLocalizedError() bool {
	if o != nil && !isNil(o.LocalizedError) {
		return true
	}

	return false
}

// SetLocalizedError gets a reference to the given LocalizedMessage and assigns it to the LocalizedError field.
func (o *CompleteInvocationInput) SetLocalizedError(v LocalizedMessage) {
	o.LocalizedError = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *CompleteInvocationInput) GetOutput() map[string]interface{} {
	if o == nil || isNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteInvocationInput) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *CompleteInvocationInput) HasOutput() bool {
	if o != nil && !isNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *CompleteInvocationInput) SetOutput(v map[string]interface{}) {
	o.Output = v
}

func (o CompleteInvocationInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteInvocationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocalizedError) {
		toSerialize["localizedError"] = o.LocalizedError
	}
	if !isNil(o.Output) {
		toSerialize["output"] = o.Output
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompleteInvocationInput) UnmarshalJSON(bytes []byte) (err error) {
	varCompleteInvocationInput := _CompleteInvocationInput{}

	if err = json.Unmarshal(bytes, &varCompleteInvocationInput); err == nil {
		*o = CompleteInvocationInput(varCompleteInvocationInput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "localizedError")
		delete(additionalProperties, "output")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompleteInvocationInput struct {
	value *CompleteInvocationInput
	isSet bool
}

func (v NullableCompleteInvocationInput) Get() *CompleteInvocationInput {
	return v.value
}

func (v *NullableCompleteInvocationInput) Set(val *CompleteInvocationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteInvocationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteInvocationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteInvocationInput(val *CompleteInvocationInput) *NullableCompleteInvocationInput {
	return &NullableCompleteInvocationInput{value: val, isSet: true}
}

func (v NullableCompleteInvocationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteInvocationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


