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

// checks if the TriggerInputCampaignGeneratedCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerInputCampaignGeneratedCampaign{}

// TriggerInputCampaignGeneratedCampaign Details about the campaign that was generated.
type TriggerInputCampaignGeneratedCampaign struct {
	// The unique ID of the campaign.
	Id string `json:"id"`
	// Human friendly name of the campaign.
	Name string `json:"name"`
	// Extended description of the campaign.
	Description string `json:"description"`
	// The date and time the campaign was created.
	Created time.Time `json:"created"`
	// The date and time the campaign was last modified.
	Modified NullableString `json:"modified,omitempty"`
	// The date and time when the campaign must be finished by.
	Deadline NullableString `json:"deadline,omitempty"`
	// The type of campaign that was generated.
	Type map[string]interface{} `json:"type"`
	CampaignOwner TriggerInputCampaignGeneratedCampaignCampaignOwner `json:"campaignOwner"`
	// The current status of the campaign.
	Status map[string]interface{} `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputCampaignGeneratedCampaign TriggerInputCampaignGeneratedCampaign

// NewTriggerInputCampaignGeneratedCampaign instantiates a new TriggerInputCampaignGeneratedCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputCampaignGeneratedCampaign(id string, name string, description string, created time.Time, type_ map[string]interface{}, campaignOwner TriggerInputCampaignGeneratedCampaignCampaignOwner, status map[string]interface{}) *TriggerInputCampaignGeneratedCampaign {
	this := TriggerInputCampaignGeneratedCampaign{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Created = created
	this.Type = type_
	this.CampaignOwner = campaignOwner
	this.Status = status
	return &this
}

// NewTriggerInputCampaignGeneratedCampaignWithDefaults instantiates a new TriggerInputCampaignGeneratedCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputCampaignGeneratedCampaignWithDefaults() *TriggerInputCampaignGeneratedCampaign {
	this := TriggerInputCampaignGeneratedCampaign{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputCampaignGeneratedCampaign) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputCampaignGeneratedCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *TriggerInputCampaignGeneratedCampaign) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetDescription(v string) {
	o.Description = v
}

// GetCreated returns the Created field value
func (o *TriggerInputCampaignGeneratedCampaign) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputCampaignGeneratedCampaign) GetModified() string {
	if o == nil || isNil(o.Modified.Get()) {
		var ret string
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputCampaignGeneratedCampaign) GetModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *TriggerInputCampaignGeneratedCampaign) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableString and assigns it to the Modified field.
func (o *TriggerInputCampaignGeneratedCampaign) SetModified(v string) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *TriggerInputCampaignGeneratedCampaign) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *TriggerInputCampaignGeneratedCampaign) UnsetModified() {
	o.Modified.Unset()
}

// GetDeadline returns the Deadline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputCampaignGeneratedCampaign) GetDeadline() string {
	if o == nil || isNil(o.Deadline.Get()) {
		var ret string
		return ret
	}
	return *o.Deadline.Get()
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputCampaignGeneratedCampaign) GetDeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deadline.Get(), o.Deadline.IsSet()
}

// HasDeadline returns a boolean if a field has been set.
func (o *TriggerInputCampaignGeneratedCampaign) HasDeadline() bool {
	if o != nil && o.Deadline.IsSet() {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given NullableString and assigns it to the Deadline field.
func (o *TriggerInputCampaignGeneratedCampaign) SetDeadline(v string) {
	o.Deadline.Set(&v)
}
// SetDeadlineNil sets the value for Deadline to be an explicit nil
func (o *TriggerInputCampaignGeneratedCampaign) SetDeadlineNil() {
	o.Deadline.Set(nil)
}

// UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
func (o *TriggerInputCampaignGeneratedCampaign) UnsetDeadline() {
	o.Deadline.Unset()
}

// GetType returns the Type field value
func (o *TriggerInputCampaignGeneratedCampaign) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetCampaignOwner returns the CampaignOwner field value
func (o *TriggerInputCampaignGeneratedCampaign) GetCampaignOwner() TriggerInputCampaignGeneratedCampaignCampaignOwner {
	if o == nil {
		var ret TriggerInputCampaignGeneratedCampaignCampaignOwner
		return ret
	}

	return o.CampaignOwner
}

// GetCampaignOwnerOk returns a tuple with the CampaignOwner field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetCampaignOwnerOk() (*TriggerInputCampaignGeneratedCampaignCampaignOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignOwner, true
}

// SetCampaignOwner sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetCampaignOwner(v TriggerInputCampaignGeneratedCampaignCampaignOwner) {
	o.CampaignOwner = v
}

// GetStatus returns the Status field value
func (o *TriggerInputCampaignGeneratedCampaign) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TriggerInputCampaignGeneratedCampaign) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *TriggerInputCampaignGeneratedCampaign) SetStatus(v map[string]interface{}) {
	o.Status = v
}

func (o TriggerInputCampaignGeneratedCampaign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerInputCampaignGeneratedCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["created"] = o.Created
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Deadline.IsSet() {
		toSerialize["deadline"] = o.Deadline.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["campaignOwner"] = o.CampaignOwner
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerInputCampaignGeneratedCampaign) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputCampaignGeneratedCampaign := _TriggerInputCampaignGeneratedCampaign{}

	if err = json.Unmarshal(bytes, &varTriggerInputCampaignGeneratedCampaign); err == nil {
		*o = TriggerInputCampaignGeneratedCampaign(varTriggerInputCampaignGeneratedCampaign)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "deadline")
		delete(additionalProperties, "type")
		delete(additionalProperties, "campaignOwner")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputCampaignGeneratedCampaign struct {
	value *TriggerInputCampaignGeneratedCampaign
	isSet bool
}

func (v NullableTriggerInputCampaignGeneratedCampaign) Get() *TriggerInputCampaignGeneratedCampaign {
	return v.value
}

func (v *NullableTriggerInputCampaignGeneratedCampaign) Set(val *TriggerInputCampaignGeneratedCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputCampaignGeneratedCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputCampaignGeneratedCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputCampaignGeneratedCampaign(val *TriggerInputCampaignGeneratedCampaign) *NullableTriggerInputCampaignGeneratedCampaign {
	return &NullableTriggerInputCampaignGeneratedCampaign{value: val, isSet: true}
}

func (v NullableTriggerInputCampaignGeneratedCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputCampaignGeneratedCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


