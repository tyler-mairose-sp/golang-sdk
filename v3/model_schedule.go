/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the Schedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule{}

// Schedule The schedule information.
type Schedule struct {
	Type ScheduleType `json:"type"`
	Days *ScheduleDays `json:"days,omitempty"`
	Hours ScheduleHours `json:"hours"`
	// A date-time in ISO-8601 format
	Expiration NullableTime `json:"expiration,omitempty"`
	// The GMT formatted timezone the schedule will run in (ex. GMT-06:00).  If no timezone is specified, the org's default timezone is used.
	TimeZoneId NullableString `json:"timeZoneId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule Schedule

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule(type_ ScheduleType, hours ScheduleHours) *Schedule {
	this := Schedule{}
	this.Type = type_
	this.Hours = hours
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule) GetType() ScheduleType {
	if o == nil {
		var ret ScheduleType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetTypeOk() (*ScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule) SetType(v ScheduleType) {
	o.Type = v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *Schedule) GetDays() ScheduleDays {
	if o == nil || isNil(o.Days) {
		var ret ScheduleDays
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDaysOk() (*ScheduleDays, bool) {
	if o == nil || isNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *Schedule) HasDays() bool {
	if o != nil && !isNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given ScheduleDays and assigns it to the Days field.
func (o *Schedule) SetDays(v ScheduleDays) {
	o.Days = &v
}

// GetHours returns the Hours field value
func (o *Schedule) GetHours() ScheduleHours {
	if o == nil {
		var ret ScheduleHours
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetHoursOk() (*ScheduleHours, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value
func (o *Schedule) SetHours(v ScheduleHours) {
	o.Hours = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetExpiration() time.Time {
	if o == nil || isNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *Schedule) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *Schedule) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}
// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *Schedule) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *Schedule) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetTimeZoneId returns the TimeZoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetTimeZoneId() string {
	if o == nil || isNil(o.TimeZoneId.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZoneId.Get()
}

// GetTimeZoneIdOk returns a tuple with the TimeZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetTimeZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZoneId.Get(), o.TimeZoneId.IsSet()
}

// HasTimeZoneId returns a boolean if a field has been set.
func (o *Schedule) HasTimeZoneId() bool {
	if o != nil && o.TimeZoneId.IsSet() {
		return true
	}

	return false
}

// SetTimeZoneId gets a reference to the given NullableString and assigns it to the TimeZoneId field.
func (o *Schedule) SetTimeZoneId(v string) {
	o.TimeZoneId.Set(&v)
}
// SetTimeZoneIdNil sets the value for TimeZoneId to be an explicit nil
func (o *Schedule) SetTimeZoneIdNil() {
	o.TimeZoneId.Set(nil)
}

// UnsetTimeZoneId ensures that no value is present for TimeZoneId, not even an explicit nil
func (o *Schedule) UnsetTimeZoneId() {
	o.TimeZoneId.Unset()
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !isNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	toSerialize["hours"] = o.Hours
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.TimeZoneId.IsSet() {
		toSerialize["timeZoneId"] = o.TimeZoneId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schedule) UnmarshalJSON(bytes []byte) (err error) {
	varSchedule := _Schedule{}

	if err = json.Unmarshal(bytes, &varSchedule); err == nil {
		*o = Schedule(varSchedule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "days")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "timeZoneId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


