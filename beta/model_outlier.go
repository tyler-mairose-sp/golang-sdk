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

// checks if the Outlier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Outlier{}

// Outlier struct for Outlier
type Outlier struct {
	// The identity's unique identifier for the outlier record
	Id *string `json:"id,omitempty"`
	// The ID of the identity that is detected as an outlier
	IdentityId *string `json:"identityId,omitempty"`
	// The type of outlier summary
	Type *string `json:"type,omitempty"`
	// The first date the outlier was detected
	FirstDetectionDate *time.Time `json:"firstDetectionDate,omitempty"`
	// The most recent date the outlier was detected
	LatestDetectionDate *time.Time `json:"latestDetectionDate,omitempty"`
	// Flag whether or not the outlier has been ignored
	Ignored *bool `json:"ignored,omitempty"`
	// Object containing mapped identity attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The outlier score determined by the detection engine ranging from 0..1
	Score *float32 `json:"score,omitempty"`
	// Enum value of if the outlier manually or automatically un-ignored. Will be NULL if outlier is not ignored
	UnignoreType *string `json:"unignoreType,omitempty"`
	// shows date when last time has been unignored outlier
	UnignoreDate *time.Time `json:"unignoreDate,omitempty"`
	// shows date when last time has been ignored outlier
	IgnoreDate *time.Time `json:"ignoreDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Outlier Outlier

// NewOutlier instantiates a new Outlier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlier() *Outlier {
	this := Outlier{}
	return &this
}

// NewOutlierWithDefaults instantiates a new Outlier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlierWithDefaults() *Outlier {
	this := Outlier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Outlier) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Outlier) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Outlier) SetId(v string) {
	o.Id = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *Outlier) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *Outlier) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *Outlier) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Outlier) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Outlier) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Outlier) SetType(v string) {
	o.Type = &v
}

// GetFirstDetectionDate returns the FirstDetectionDate field value if set, zero value otherwise.
func (o *Outlier) GetFirstDetectionDate() time.Time {
	if o == nil || isNil(o.FirstDetectionDate) {
		var ret time.Time
		return ret
	}
	return *o.FirstDetectionDate
}

// GetFirstDetectionDateOk returns a tuple with the FirstDetectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetFirstDetectionDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.FirstDetectionDate) {
		return nil, false
	}
	return o.FirstDetectionDate, true
}

// HasFirstDetectionDate returns a boolean if a field has been set.
func (o *Outlier) HasFirstDetectionDate() bool {
	if o != nil && !isNil(o.FirstDetectionDate) {
		return true
	}

	return false
}

// SetFirstDetectionDate gets a reference to the given time.Time and assigns it to the FirstDetectionDate field.
func (o *Outlier) SetFirstDetectionDate(v time.Time) {
	o.FirstDetectionDate = &v
}

// GetLatestDetectionDate returns the LatestDetectionDate field value if set, zero value otherwise.
func (o *Outlier) GetLatestDetectionDate() time.Time {
	if o == nil || isNil(o.LatestDetectionDate) {
		var ret time.Time
		return ret
	}
	return *o.LatestDetectionDate
}

// GetLatestDetectionDateOk returns a tuple with the LatestDetectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetLatestDetectionDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.LatestDetectionDate) {
		return nil, false
	}
	return o.LatestDetectionDate, true
}

// HasLatestDetectionDate returns a boolean if a field has been set.
func (o *Outlier) HasLatestDetectionDate() bool {
	if o != nil && !isNil(o.LatestDetectionDate) {
		return true
	}

	return false
}

// SetLatestDetectionDate gets a reference to the given time.Time and assigns it to the LatestDetectionDate field.
func (o *Outlier) SetLatestDetectionDate(v time.Time) {
	o.LatestDetectionDate = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *Outlier) GetIgnored() bool {
	if o == nil || isNil(o.Ignored) {
		var ret bool
		return ret
	}
	return *o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetIgnoredOk() (*bool, bool) {
	if o == nil || isNil(o.Ignored) {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *Outlier) HasIgnored() bool {
	if o != nil && !isNil(o.Ignored) {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given bool and assigns it to the Ignored field.
func (o *Outlier) SetIgnored(v bool) {
	o.Ignored = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Outlier) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Outlier) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Outlier) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Outlier) GetScore() float32 {
	if o == nil || isNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetScoreOk() (*float32, bool) {
	if o == nil || isNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Outlier) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *Outlier) SetScore(v float32) {
	o.Score = &v
}

// GetUnignoreType returns the UnignoreType field value if set, zero value otherwise.
func (o *Outlier) GetUnignoreType() string {
	if o == nil || isNil(o.UnignoreType) {
		var ret string
		return ret
	}
	return *o.UnignoreType
}

// GetUnignoreTypeOk returns a tuple with the UnignoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetUnignoreTypeOk() (*string, bool) {
	if o == nil || isNil(o.UnignoreType) {
		return nil, false
	}
	return o.UnignoreType, true
}

// HasUnignoreType returns a boolean if a field has been set.
func (o *Outlier) HasUnignoreType() bool {
	if o != nil && !isNil(o.UnignoreType) {
		return true
	}

	return false
}

// SetUnignoreType gets a reference to the given string and assigns it to the UnignoreType field.
func (o *Outlier) SetUnignoreType(v string) {
	o.UnignoreType = &v
}

// GetUnignoreDate returns the UnignoreDate field value if set, zero value otherwise.
func (o *Outlier) GetUnignoreDate() time.Time {
	if o == nil || isNil(o.UnignoreDate) {
		var ret time.Time
		return ret
	}
	return *o.UnignoreDate
}

// GetUnignoreDateOk returns a tuple with the UnignoreDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetUnignoreDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.UnignoreDate) {
		return nil, false
	}
	return o.UnignoreDate, true
}

// HasUnignoreDate returns a boolean if a field has been set.
func (o *Outlier) HasUnignoreDate() bool {
	if o != nil && !isNil(o.UnignoreDate) {
		return true
	}

	return false
}

// SetUnignoreDate gets a reference to the given time.Time and assigns it to the UnignoreDate field.
func (o *Outlier) SetUnignoreDate(v time.Time) {
	o.UnignoreDate = &v
}

// GetIgnoreDate returns the IgnoreDate field value if set, zero value otherwise.
func (o *Outlier) GetIgnoreDate() time.Time {
	if o == nil || isNil(o.IgnoreDate) {
		var ret time.Time
		return ret
	}
	return *o.IgnoreDate
}

// GetIgnoreDateOk returns a tuple with the IgnoreDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outlier) GetIgnoreDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.IgnoreDate) {
		return nil, false
	}
	return o.IgnoreDate, true
}

// HasIgnoreDate returns a boolean if a field has been set.
func (o *Outlier) HasIgnoreDate() bool {
	if o != nil && !isNil(o.IgnoreDate) {
		return true
	}

	return false
}

// SetIgnoreDate gets a reference to the given time.Time and assigns it to the IgnoreDate field.
func (o *Outlier) SetIgnoreDate(v time.Time) {
	o.IgnoreDate = &v
}

func (o Outlier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Outlier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.FirstDetectionDate) {
		toSerialize["firstDetectionDate"] = o.FirstDetectionDate
	}
	if !isNil(o.LatestDetectionDate) {
		toSerialize["latestDetectionDate"] = o.LatestDetectionDate
	}
	if !isNil(o.Ignored) {
		toSerialize["ignored"] = o.Ignored
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !isNil(o.UnignoreType) {
		toSerialize["unignoreType"] = o.UnignoreType
	}
	if !isNil(o.UnignoreDate) {
		toSerialize["unignoreDate"] = o.UnignoreDate
	}
	if !isNil(o.IgnoreDate) {
		toSerialize["ignoreDate"] = o.IgnoreDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Outlier) UnmarshalJSON(bytes []byte) (err error) {
	varOutlier := _Outlier{}

	if err = json.Unmarshal(bytes, &varOutlier); err == nil {
		*o = Outlier(varOutlier)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "firstDetectionDate")
		delete(additionalProperties, "latestDetectionDate")
		delete(additionalProperties, "ignored")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "score")
		delete(additionalProperties, "unignoreType")
		delete(additionalProperties, "unignoreDate")
		delete(additionalProperties, "ignoreDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutlier struct {
	value *Outlier
	isSet bool
}

func (v NullableOutlier) Get() *Outlier {
	return v.value
}

func (v *NullableOutlier) Set(val *Outlier) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlier) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlier(val *Outlier) *NullableOutlier {
	return &NullableOutlier{value: val, isSet: true}
}

func (v NullableOutlier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


