/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// AccountAction Object for specifying Actions to be performed on a specified list of sources' account.
type AccountAction struct {
	// Describes if action will be enabled or disabled
	Action *string `json:"action,omitempty"`
	// List of unique source IDs. The sources must have the ENABLE feature or flat file source. See \"/sources\" endpoint for source features.
	SourceIds []string `json:"sourceIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountAction AccountAction

// NewAccountAction instantiates a new AccountAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAction() *AccountAction {
	this := AccountAction{}
	return &this
}

// NewAccountActionWithDefaults instantiates a new AccountAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountActionWithDefaults() *AccountAction {
	this := AccountAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AccountAction) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAction) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AccountAction) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AccountAction) SetAction(v string) {
	o.Action = &v
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise.
func (o *AccountAction) GetSourceIds() []string {
	if o == nil || isNil(o.SourceIds) {
		var ret []string
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAction) GetSourceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SourceIds) {
		return nil, false
	}
	return o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *AccountAction) HasSourceIds() bool {
	if o != nil && !isNil(o.SourceIds) {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []string and assigns it to the SourceIds field.
func (o *AccountAction) SetSourceIds(v []string) {
	o.SourceIds = v
}

func (o AccountAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !isNil(o.SourceIds) {
		toSerialize["sourceIds"] = o.SourceIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountAction) UnmarshalJSON(bytes []byte) (err error) {
	varAccountAction := _AccountAction{}

	if err = json.Unmarshal(bytes, &varAccountAction); err == nil {
		*o = AccountAction(varAccountAction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "sourceIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAction struct {
	value *AccountAction
	isSet bool
}

func (v NullableAccountAction) Get() *AccountAction {
	return v.value
}

func (v *NullableAccountAction) Set(val *AccountAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAction(val *AccountAction) *NullableAccountAction {
	return &NullableAccountAction{value: val, isSet: true}
}

func (v NullableAccountAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


