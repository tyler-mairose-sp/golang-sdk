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

// checks if the AccessRequestRecommendationActionItemResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestRecommendationActionItemResponseDto{}

// AccessRequestRecommendationActionItemResponseDto struct for AccessRequestRecommendationActionItemResponseDto
type AccessRequestRecommendationActionItemResponseDto struct {
	// The identity ID taking the action.
	IdentityId *string `json:"identityId,omitempty"`
	Access *AccessRequestRecommendationItem `json:"access,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestRecommendationActionItemResponseDto AccessRequestRecommendationActionItemResponseDto

// NewAccessRequestRecommendationActionItemResponseDto instantiates a new AccessRequestRecommendationActionItemResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestRecommendationActionItemResponseDto() *AccessRequestRecommendationActionItemResponseDto {
	this := AccessRequestRecommendationActionItemResponseDto{}
	return &this
}

// NewAccessRequestRecommendationActionItemResponseDtoWithDefaults instantiates a new AccessRequestRecommendationActionItemResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestRecommendationActionItemResponseDtoWithDefaults() *AccessRequestRecommendationActionItemResponseDto {
	this := AccessRequestRecommendationActionItemResponseDto{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *AccessRequestRecommendationActionItemResponseDto) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationActionItemResponseDto) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *AccessRequestRecommendationActionItemResponseDto) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *AccessRequestRecommendationActionItemResponseDto) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AccessRequestRecommendationActionItemResponseDto) GetAccess() AccessRequestRecommendationItem {
	if o == nil || isNil(o.Access) {
		var ret AccessRequestRecommendationItem
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationActionItemResponseDto) GetAccessOk() (*AccessRequestRecommendationItem, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AccessRequestRecommendationActionItemResponseDto) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AccessRequestRecommendationItem and assigns it to the Access field.
func (o *AccessRequestRecommendationActionItemResponseDto) SetAccess(v AccessRequestRecommendationItem) {
	o.Access = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AccessRequestRecommendationActionItemResponseDto) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationActionItemResponseDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AccessRequestRecommendationActionItemResponseDto) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AccessRequestRecommendationActionItemResponseDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o AccessRequestRecommendationActionItemResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestRecommendationActionItemResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestRecommendationActionItemResponseDto) UnmarshalJSON(bytes []byte) (err error) {
	varAccessRequestRecommendationActionItemResponseDto := _AccessRequestRecommendationActionItemResponseDto{}

	if err = json.Unmarshal(bytes, &varAccessRequestRecommendationActionItemResponseDto); err == nil {
		*o = AccessRequestRecommendationActionItemResponseDto(varAccessRequestRecommendationActionItemResponseDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "access")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestRecommendationActionItemResponseDto struct {
	value *AccessRequestRecommendationActionItemResponseDto
	isSet bool
}

func (v NullableAccessRequestRecommendationActionItemResponseDto) Get() *AccessRequestRecommendationActionItemResponseDto {
	return v.value
}

func (v *NullableAccessRequestRecommendationActionItemResponseDto) Set(val *AccessRequestRecommendationActionItemResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestRecommendationActionItemResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestRecommendationActionItemResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestRecommendationActionItemResponseDto(val *AccessRequestRecommendationActionItemResponseDto) *NullableAccessRequestRecommendationActionItemResponseDto {
	return &NullableAccessRequestRecommendationActionItemResponseDto{value: val, isSet: true}
}

func (v NullableAccessRequestRecommendationActionItemResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestRecommendationActionItemResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


