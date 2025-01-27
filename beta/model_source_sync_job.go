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

// checks if the SourceSyncJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceSyncJob{}

// SourceSyncJob struct for SourceSyncJob
type SourceSyncJob struct {
	// Job ID.
	Id string `json:"id"`
	// The job status.
	Status string `json:"status"`
	Payload SourceSyncPayload `json:"payload"`
	AdditionalProperties map[string]interface{}
}

type _SourceSyncJob SourceSyncJob

// NewSourceSyncJob instantiates a new SourceSyncJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSyncJob(id string, status string, payload SourceSyncPayload) *SourceSyncJob {
	this := SourceSyncJob{}
	this.Id = id
	this.Status = status
	this.Payload = payload
	return &this
}

// NewSourceSyncJobWithDefaults instantiates a new SourceSyncJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSyncJobWithDefaults() *SourceSyncJob {
	this := SourceSyncJob{}
	return &this
}

// GetId returns the Id field value
func (o *SourceSyncJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceSyncJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceSyncJob) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *SourceSyncJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SourceSyncJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SourceSyncJob) SetStatus(v string) {
	o.Status = v
}

// GetPayload returns the Payload field value
func (o *SourceSyncJob) GetPayload() SourceSyncPayload {
	if o == nil {
		var ret SourceSyncPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *SourceSyncJob) GetPayloadOk() (*SourceSyncPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *SourceSyncJob) SetPayload(v SourceSyncPayload) {
	o.Payload = v
}

func (o SourceSyncJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceSyncJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["payload"] = o.Payload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceSyncJob) UnmarshalJSON(bytes []byte) (err error) {
	varSourceSyncJob := _SourceSyncJob{}

	if err = json.Unmarshal(bytes, &varSourceSyncJob); err == nil {
		*o = SourceSyncJob(varSourceSyncJob)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceSyncJob struct {
	value *SourceSyncJob
	isSet bool
}

func (v NullableSourceSyncJob) Get() *SourceSyncJob {
	return v.value
}

func (v *NullableSourceSyncJob) Set(val *SourceSyncJob) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSyncJob) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSyncJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSyncJob(val *SourceSyncJob) *NullableSourceSyncJob {
	return &NullableSourceSyncJob{value: val, isSet: true}
}

func (v NullableSourceSyncJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSyncJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


