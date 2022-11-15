/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// Schedule1 The schedule information. 
type Schedule1 struct {
	Type ScheduleType `json:"type"`
	Months NullableSelector `json:"months,omitempty"`
	Days NullableSelector `json:"days,omitempty"`
	Hours NullableSelector `json:"hours"`
	// A date-time in ISO-8601 format
	Expiration NullableTime `json:"expiration,omitempty"`
	// The ID of the time zone for the schedule. 
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule1 Schedule1

// NewSchedule1 instantiates a new Schedule1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule1(type_ ScheduleType, hours NullableSelector) *Schedule1 {
	this := Schedule1{}
	this.Type = type_
	this.Hours = hours
	return &this
}

// NewSchedule1WithDefaults instantiates a new Schedule1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule1WithDefaults() *Schedule1 {
	this := Schedule1{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule1) GetType() ScheduleType {
	if o == nil {
		var ret ScheduleType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetTypeOk() (*ScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule1) SetType(v ScheduleType) {
	o.Type = v
}

// GetMonths returns the Months field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1) GetMonths() Selector {
	if o == nil || isNil(o.Months.Get()) {
		var ret Selector
		return ret
	}
	return *o.Months.Get()
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1) GetMonthsOk() (*Selector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Months.Get(), o.Months.IsSet()
}

// HasMonths returns a boolean if a field has been set.
func (o *Schedule1) HasMonths() bool {
	if o != nil && o.Months.IsSet() {
		return true
	}

	return false
}

// SetMonths gets a reference to the given NullableSelector and assigns it to the Months field.
func (o *Schedule1) SetMonths(v Selector) {
	o.Months.Set(&v)
}
// SetMonthsNil sets the value for Months to be an explicit nil
func (o *Schedule1) SetMonthsNil() {
	o.Months.Set(nil)
}

// UnsetMonths ensures that no value is present for Months, not even an explicit nil
func (o *Schedule1) UnsetMonths() {
	o.Months.Unset()
}

// GetDays returns the Days field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1) GetDays() Selector {
	if o == nil || isNil(o.Days.Get()) {
		var ret Selector
		return ret
	}
	return *o.Days.Get()
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1) GetDaysOk() (*Selector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days.Get(), o.Days.IsSet()
}

// HasDays returns a boolean if a field has been set.
func (o *Schedule1) HasDays() bool {
	if o != nil && o.Days.IsSet() {
		return true
	}

	return false
}

// SetDays gets a reference to the given NullableSelector and assigns it to the Days field.
func (o *Schedule1) SetDays(v Selector) {
	o.Days.Set(&v)
}
// SetDaysNil sets the value for Days to be an explicit nil
func (o *Schedule1) SetDaysNil() {
	o.Days.Set(nil)
}

// UnsetDays ensures that no value is present for Days, not even an explicit nil
func (o *Schedule1) UnsetDays() {
	o.Days.Unset()
}

// GetHours returns the Hours field value
// If the value is explicit nil, the zero value for Selector will be returned
func (o *Schedule1) GetHours() Selector {
	if o == nil || o.Hours.Get() == nil {
		var ret Selector
		return ret
	}

	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1) GetHoursOk() (*Selector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// SetHours sets field value
func (o *Schedule1) SetHours(v Selector) {
	o.Hours.Set(&v)
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1) GetExpiration() time.Time {
	if o == nil || isNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *Schedule1) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *Schedule1) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}
// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *Schedule1) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *Schedule1) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetTimeZoneId returns the TimeZoneId field value if set, zero value otherwise.
func (o *Schedule1) GetTimeZoneId() string {
	if o == nil || isNil(o.TimeZoneId) {
		var ret string
		return ret
	}
	return *o.TimeZoneId
}

// GetTimeZoneIdOk returns a tuple with the TimeZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule1) GetTimeZoneIdOk() (*string, bool) {
	if o == nil || isNil(o.TimeZoneId) {
		return nil, false
	}
	return o.TimeZoneId, true
}

// HasTimeZoneId returns a boolean if a field has been set.
func (o *Schedule1) HasTimeZoneId() bool {
	if o != nil && !isNil(o.TimeZoneId) {
		return true
	}

	return false
}

// SetTimeZoneId gets a reference to the given string and assigns it to the TimeZoneId field.
func (o *Schedule1) SetTimeZoneId(v string) {
	o.TimeZoneId = &v
}

func (o Schedule1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Months.IsSet() {
		toSerialize["months"] = o.Months.Get()
	}
	if o.Days.IsSet() {
		toSerialize["days"] = o.Days.Get()
	}
	if true {
		toSerialize["hours"] = o.Hours.Get()
	}
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if !isNil(o.TimeZoneId) {
		toSerialize["timeZoneId"] = o.TimeZoneId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Schedule1) UnmarshalJSON(bytes []byte) (err error) {
	varSchedule1 := _Schedule1{}

	if err = json.Unmarshal(bytes, &varSchedule1); err == nil {
		*o = Schedule1(varSchedule1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "months")
		delete(additionalProperties, "days")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "timeZoneId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule1 struct {
	value *Schedule1
	isSet bool
}

func (v NullableSchedule1) Get() *Schedule1 {
	return v.value
}

func (v *NullableSchedule1) Set(val *Schedule1) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule1) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule1(val *Schedule1) *NullableSchedule1 {
	return &NullableSchedule1{value: val, isSet: true}
}

func (v NullableSchedule1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


