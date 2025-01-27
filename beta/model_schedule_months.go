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

// checks if the ScheduleMonths type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleMonths{}

// ScheduleMonths Specifies which months of a schedule are active. Only valid for ANNUALLY schedule types. Examples:  On February and March: * type LIST * values \"2\", \"3\"  Every 3 months, starting in January (quarterly): * type LIST * values \"1\" * interval 3  Every two months between July and December: * type RANGE * values \"7\", \"12\" * interval 2 
type ScheduleMonths struct {
	Type string `json:"type"`
	Values []string `json:"values"`
	Interval *int32 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleMonths ScheduleMonths

// NewScheduleMonths instantiates a new ScheduleMonths object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleMonths(type_ string, values []string) *ScheduleMonths {
	this := ScheduleMonths{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewScheduleMonthsWithDefaults instantiates a new ScheduleMonths object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleMonthsWithDefaults() *ScheduleMonths {
	this := ScheduleMonths{}
	return &this
}

// GetType returns the Type field value
func (o *ScheduleMonths) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleMonths) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleMonths) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *ScheduleMonths) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ScheduleMonths) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ScheduleMonths) SetValues(v []string) {
	o.Values = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ScheduleMonths) GetInterval() int32 {
	if o == nil || isNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleMonths) GetIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ScheduleMonths) HasInterval() bool {
	if o != nil && !isNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *ScheduleMonths) SetInterval(v int32) {
	o.Interval = &v
}

func (o ScheduleMonths) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleMonths) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values
	if !isNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduleMonths) UnmarshalJSON(bytes []byte) (err error) {
	varScheduleMonths := _ScheduleMonths{}

	if err = json.Unmarshal(bytes, &varScheduleMonths); err == nil {
		*o = ScheduleMonths(varScheduleMonths)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		delete(additionalProperties, "interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduleMonths struct {
	value *ScheduleMonths
	isSet bool
}

func (v NullableScheduleMonths) Get() *ScheduleMonths {
	return v.value
}

func (v *NullableScheduleMonths) Set(val *ScheduleMonths) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleMonths) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleMonths) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleMonths(val *ScheduleMonths) *NullableScheduleMonths {
	return &NullableScheduleMonths{value: val, isSet: true}
}

func (v NullableScheduleMonths) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleMonths) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


