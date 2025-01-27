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

// checks if the TriggerInputSavedSearchCompleteSearchResultsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputSavedSearchCompleteSearchResultsEntitlement{}

// TriggerInputSavedSearchCompleteSearchResultsEntitlement A table of entitlements that match the search criteria.
type TriggerInputSavedSearchCompleteSearchResultsEntitlement struct {
	// The number of rows in the table.
	Count string `json:"count"`
	// The type of object represented in the table.
	Noun string `json:"noun"`
	// A sample of the data in the table.
	Preview [][]string `json:"preview"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputSavedSearchCompleteSearchResultsEntitlement TriggerInputSavedSearchCompleteSearchResultsEntitlement

// NewTriggerInputSavedSearchCompleteSearchResultsEntitlement instantiates a new TriggerInputSavedSearchCompleteSearchResultsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputSavedSearchCompleteSearchResultsEntitlement(count string, noun string, preview [][]string) *TriggerInputSavedSearchCompleteSearchResultsEntitlement {
	this := TriggerInputSavedSearchCompleteSearchResultsEntitlement{}
	this.Count = count
	this.Noun = noun
	this.Preview = preview
	return &this
}

// NewTriggerInputSavedSearchCompleteSearchResultsEntitlementWithDefaults instantiates a new TriggerInputSavedSearchCompleteSearchResultsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputSavedSearchCompleteSearchResultsEntitlementWithDefaults() *TriggerInputSavedSearchCompleteSearchResultsEntitlement {
	this := TriggerInputSavedSearchCompleteSearchResultsEntitlement{}
	return &this
}

// GetCount returns the Count field value
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) GetCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) GetCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) SetCount(v string) {
	o.Count = v
}

// GetNoun returns the Noun field value
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) GetNoun() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Noun
}

// GetNounOk returns a tuple with the Noun field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) GetNounOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Noun, true
}

// SetNoun sets field value
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) SetNoun(v string) {
	o.Noun = v
}

// GetPreview returns the Preview field value
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) GetPreview() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) GetPreviewOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preview, true
}

// SetPreview sets field value
func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) SetPreview(v [][]string) {
	o.Preview = v
}

func (o TriggerInputSavedSearchCompleteSearchResultsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputSavedSearchCompleteSearchResultsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["noun"] = o.Noun
	toSerialize["preview"] = o.Preview

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputSavedSearchCompleteSearchResultsEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputSavedSearchCompleteSearchResultsEntitlement := _TriggerInputSavedSearchCompleteSearchResultsEntitlement{}

	if err = json.Unmarshal(bytes, &varTriggerInputSavedSearchCompleteSearchResultsEntitlement); err == nil {
		*o = TriggerInputSavedSearchCompleteSearchResultsEntitlement(varTriggerInputSavedSearchCompleteSearchResultsEntitlement)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "noun")
		delete(additionalProperties, "preview")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement struct {
	value *TriggerInputSavedSearchCompleteSearchResultsEntitlement
	isSet bool
}

func (v NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement) Get() *TriggerInputSavedSearchCompleteSearchResultsEntitlement {
	return v.value
}

func (v *NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement) Set(val *TriggerInputSavedSearchCompleteSearchResultsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputSavedSearchCompleteSearchResultsEntitlement(val *TriggerInputSavedSearchCompleteSearchResultsEntitlement) *NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement {
	return &NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement{value: val, isSet: true}
}

func (v NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


