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

// ExpansionItem struct for ExpansionItem
type ExpansionItem struct {
	// The ID of the account
	AccountId *string `json:"accountId,omitempty"`
	Cause *string `json:"cause,omitempty"`
	// The name of the item
	Name *string `json:"name,omitempty"`
	AttributeRequests []AttributeRequest `json:"attributeRequests,omitempty"`
	Source *Source1 `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExpansionItem ExpansionItem

// NewExpansionItem instantiates a new ExpansionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpansionItem() *ExpansionItem {
	this := ExpansionItem{}
	return &this
}

// NewExpansionItemWithDefaults instantiates a new ExpansionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpansionItemWithDefaults() *ExpansionItem {
	this := ExpansionItem{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ExpansionItem) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ExpansionItem) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ExpansionItem) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ExpansionItem) GetCause() string {
	if o == nil || isNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetCauseOk() (*string, bool) {
	if o == nil || isNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ExpansionItem) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ExpansionItem) SetCause(v string) {
	o.Cause = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExpansionItem) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExpansionItem) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExpansionItem) SetName(v string) {
	o.Name = &v
}

// GetAttributeRequests returns the AttributeRequests field value if set, zero value otherwise.
func (o *ExpansionItem) GetAttributeRequests() []AttributeRequest {
	if o == nil || isNil(o.AttributeRequests) {
		var ret []AttributeRequest
		return ret
	}
	return o.AttributeRequests
}

// GetAttributeRequestsOk returns a tuple with the AttributeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetAttributeRequestsOk() ([]AttributeRequest, bool) {
	if o == nil || isNil(o.AttributeRequests) {
		return nil, false
	}
	return o.AttributeRequests, true
}

// HasAttributeRequests returns a boolean if a field has been set.
func (o *ExpansionItem) HasAttributeRequests() bool {
	if o != nil && !isNil(o.AttributeRequests) {
		return true
	}

	return false
}

// SetAttributeRequests gets a reference to the given []AttributeRequest and assigns it to the AttributeRequests field.
func (o *ExpansionItem) SetAttributeRequests(v []AttributeRequest) {
	o.AttributeRequests = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ExpansionItem) GetSource() Source1 {
	if o == nil || isNil(o.Source) {
		var ret Source1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetSourceOk() (*Source1, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ExpansionItem) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source1 and assigns it to the Source field.
func (o *ExpansionItem) SetSource(v Source1) {
	o.Source = &v
}

func (o ExpansionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AttributeRequests) {
		toSerialize["attributeRequests"] = o.AttributeRequests
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExpansionItem) UnmarshalJSON(bytes []byte) (err error) {
	varExpansionItem := _ExpansionItem{}

	if err = json.Unmarshal(bytes, &varExpansionItem); err == nil {
		*o = ExpansionItem(varExpansionItem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "cause")
		delete(additionalProperties, "name")
		delete(additionalProperties, "attributeRequests")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpansionItem struct {
	value *ExpansionItem
	isSet bool
}

func (v NullableExpansionItem) Get() *ExpansionItem {
	return v.value
}

func (v *NullableExpansionItem) Set(val *ExpansionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExpansionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExpansionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpansionItem(val *ExpansionItem) *NullableExpansionItem {
	return &NullableExpansionItem{value: val, isSet: true}
}

func (v NullableExpansionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpansionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


