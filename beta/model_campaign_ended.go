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

// checks if the CampaignEnded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignEnded{}

// CampaignEnded struct for CampaignEnded
type CampaignEnded struct {
	Campaign TriggerInputCampaignEndedCampaign `json:"campaign"`
	AdditionalProperties map[string]interface{}
}

type _CampaignEnded CampaignEnded

// NewCampaignEnded instantiates a new CampaignEnded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignEnded(campaign TriggerInputCampaignEndedCampaign) *CampaignEnded {
	this := CampaignEnded{}
	this.Campaign = campaign
	return &this
}

// NewCampaignEndedWithDefaults instantiates a new CampaignEnded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignEndedWithDefaults() *CampaignEnded {
	this := CampaignEnded{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *CampaignEnded) GetCampaign() TriggerInputCampaignEndedCampaign {
	if o == nil {
		var ret TriggerInputCampaignEndedCampaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *CampaignEnded) GetCampaignOk() (*TriggerInputCampaignEndedCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *CampaignEnded) SetCampaign(v TriggerInputCampaignEndedCampaign) {
	o.Campaign = v
}

func (o CampaignEnded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignEnded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignEnded) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignEnded := _CampaignEnded{}

	if err = json.Unmarshal(bytes, &varCampaignEnded); err == nil {
		*o = CampaignEnded(varCampaignEnded)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignEnded struct {
	value *CampaignEnded
	isSet bool
}

func (v NullableCampaignEnded) Get() *CampaignEnded {
	return v.value
}

func (v *NullableCampaignEnded) Set(val *CampaignEnded) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignEnded) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignEnded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignEnded(val *CampaignEnded) *NullableCampaignEnded {
	return &NullableCampaignEnded{value: val, isSet: true}
}

func (v NullableCampaignEnded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignEnded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


