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

// checks if the AccountsCollectedForAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsCollectedForAggregation{}

// AccountsCollectedForAggregation struct for AccountsCollectedForAggregation
type AccountsCollectedForAggregation struct {
	Source TriggerInputAccountsCollectedForAggregationSource `json:"source"`
	// The overall status of the collection.
	Status map[string]interface{} `json:"status"`
	// The date and time when the account collection started.
	Started time.Time `json:"started"`
	// The date and time when the account collection finished.
	Completed time.Time `json:"completed"`
	// A list of errors that occurred during the collection.
	Errors []string `json:"errors"`
	// A list of warnings that occurred during the collection.
	Warnings []string `json:"warnings"`
	Stats TriggerInputAccountsCollectedForAggregationStats `json:"stats"`
	AdditionalProperties map[string]interface{}
}

type _AccountsCollectedForAggregation AccountsCollectedForAggregation

// NewAccountsCollectedForAggregation instantiates a new AccountsCollectedForAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsCollectedForAggregation(source TriggerInputAccountsCollectedForAggregationSource, status map[string]interface{}, started time.Time, completed time.Time, errors []string, warnings []string, stats TriggerInputAccountsCollectedForAggregationStats) *AccountsCollectedForAggregation {
	this := AccountsCollectedForAggregation{}
	this.Source = source
	this.Status = status
	this.Started = started
	this.Completed = completed
	this.Errors = errors
	this.Warnings = warnings
	this.Stats = stats
	return &this
}

// NewAccountsCollectedForAggregationWithDefaults instantiates a new AccountsCollectedForAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsCollectedForAggregationWithDefaults() *AccountsCollectedForAggregation {
	this := AccountsCollectedForAggregation{}
	return &this
}

// GetSource returns the Source field value
func (o *AccountsCollectedForAggregation) GetSource() TriggerInputAccountsCollectedForAggregationSource {
	if o == nil {
		var ret TriggerInputAccountsCollectedForAggregationSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregation) GetSourceOk() (*TriggerInputAccountsCollectedForAggregationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AccountsCollectedForAggregation) SetSource(v TriggerInputAccountsCollectedForAggregationSource) {
	o.Source = v
}

// GetStatus returns the Status field value
func (o *AccountsCollectedForAggregation) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregation) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *AccountsCollectedForAggregation) SetStatus(v map[string]interface{}) {
	o.Status = v
}

// GetStarted returns the Started field value
func (o *AccountsCollectedForAggregation) GetStarted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Started
}

// GetStartedOk returns a tuple with the Started field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregation) GetStartedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Started, true
}

// SetStarted sets field value
func (o *AccountsCollectedForAggregation) SetStarted(v time.Time) {
	o.Started = v
}

// GetCompleted returns the Completed field value
func (o *AccountsCollectedForAggregation) GetCompleted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregation) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *AccountsCollectedForAggregation) SetCompleted(v time.Time) {
	o.Completed = v
}

// GetErrors returns the Errors field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AccountsCollectedForAggregation) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountsCollectedForAggregation) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *AccountsCollectedForAggregation) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AccountsCollectedForAggregation) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountsCollectedForAggregation) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *AccountsCollectedForAggregation) SetWarnings(v []string) {
	o.Warnings = v
}

// GetStats returns the Stats field value
func (o *AccountsCollectedForAggregation) GetStats() TriggerInputAccountsCollectedForAggregationStats {
	if o == nil {
		var ret TriggerInputAccountsCollectedForAggregationStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *AccountsCollectedForAggregation) GetStatsOk() (*TriggerInputAccountsCollectedForAggregationStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *AccountsCollectedForAggregation) SetStats(v TriggerInputAccountsCollectedForAggregationStats) {
	o.Stats = v
}

func (o AccountsCollectedForAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsCollectedForAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status
	toSerialize["started"] = o.Started
	toSerialize["completed"] = o.Completed
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	toSerialize["stats"] = o.Stats

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountsCollectedForAggregation) UnmarshalJSON(bytes []byte) (err error) {
	varAccountsCollectedForAggregation := _AccountsCollectedForAggregation{}

	if err = json.Unmarshal(bytes, &varAccountsCollectedForAggregation); err == nil {
		*o = AccountsCollectedForAggregation(varAccountsCollectedForAggregation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "status")
		delete(additionalProperties, "started")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "stats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountsCollectedForAggregation struct {
	value *AccountsCollectedForAggregation
	isSet bool
}

func (v NullableAccountsCollectedForAggregation) Get() *AccountsCollectedForAggregation {
	return v.value
}

func (v *NullableAccountsCollectedForAggregation) Set(val *AccountsCollectedForAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsCollectedForAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsCollectedForAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsCollectedForAggregation(val *AccountsCollectedForAggregation) *NullableAccountsCollectedForAggregation {
	return &NullableAccountsCollectedForAggregation{value: val, isSet: true}
}

func (v NullableAccountsCollectedForAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsCollectedForAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


