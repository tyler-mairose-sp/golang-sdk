/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the AccountRequestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountRequestResult{}

// AccountRequestResult struct for AccountRequestResult
type AccountRequestResult struct {
	Errors []string `json:"errors,omitempty"`
	// The status of the account request
	Status *string `json:"status,omitempty"`
	TicketId NullableString `json:"ticketId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountRequestResult AccountRequestResult

// NewAccountRequestResult instantiates a new AccountRequestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRequestResult() *AccountRequestResult {
	this := AccountRequestResult{}
	return &this
}

// NewAccountRequestResultWithDefaults instantiates a new AccountRequestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRequestResultWithDefaults() *AccountRequestResult {
	this := AccountRequestResult{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AccountRequestResult) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRequestResult) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AccountRequestResult) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AccountRequestResult) SetErrors(v []string) {
	o.Errors = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountRequestResult) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRequestResult) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountRequestResult) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountRequestResult) SetStatus(v string) {
	o.Status = &v
}

// GetTicketId returns the TicketId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequestResult) GetTicketId() string {
	if o == nil || isNil(o.TicketId.Get()) {
		var ret string
		return ret
	}
	return *o.TicketId.Get()
}

// GetTicketIdOk returns a tuple with the TicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequestResult) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketId.Get(), o.TicketId.IsSet()
}

// HasTicketId returns a boolean if a field has been set.
func (o *AccountRequestResult) HasTicketId() bool {
	if o != nil && o.TicketId.IsSet() {
		return true
	}

	return false
}

// SetTicketId gets a reference to the given NullableString and assigns it to the TicketId field.
func (o *AccountRequestResult) SetTicketId(v string) {
	o.TicketId.Set(&v)
}
// SetTicketIdNil sets the value for TicketId to be an explicit nil
func (o *AccountRequestResult) SetTicketIdNil() {
	o.TicketId.Set(nil)
}

// UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
func (o *AccountRequestResult) UnsetTicketId() {
	o.TicketId.Unset()
}

func (o AccountRequestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountRequestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.TicketId.IsSet() {
		toSerialize["ticketId"] = o.TicketId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountRequestResult) UnmarshalJSON(bytes []byte) (err error) {
	varAccountRequestResult := _AccountRequestResult{}

	if err = json.Unmarshal(bytes, &varAccountRequestResult); err == nil {
		*o = AccountRequestResult(varAccountRequestResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		delete(additionalProperties, "status")
		delete(additionalProperties, "ticketId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountRequestResult struct {
	value *AccountRequestResult
	isSet bool
}

func (v NullableAccountRequestResult) Get() *AccountRequestResult {
	return v.value
}

func (v *NullableAccountRequestResult) Set(val *AccountRequestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRequestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRequestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRequestResult(val *AccountRequestResult) *NullableAccountRequestResult {
	return &NullableAccountRequestResult{value: val, isSet: true}
}

func (v NullableAccountRequestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRequestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


