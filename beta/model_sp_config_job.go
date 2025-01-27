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

// checks if the SpConfigJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpConfigJob{}

// SpConfigJob struct for SpConfigJob
type SpConfigJob struct {
	// Unique id assigned to this job.
	JobId string `json:"jobId"`
	// Status of the job.
	Status string `json:"status"`
	// Type of the job, either export or import.
	Type string `json:"type"`
	// This message contains additional information about the overall status of the job.
	Message string `json:"message"`
	// Optional user defined description/name for export job.
	Description string `json:"description"`
	// The time until which the artifacts will be available for download.
	Expiration time.Time `json:"expiration"`
	// The time the job was started.
	Created time.Time `json:"created"`
	// The time of the last update to the job.
	Modified time.Time `json:"modified"`
	// The time the job was completed.
	Completed time.Time `json:"completed"`
	AdditionalProperties map[string]interface{}
}

type _SpConfigJob SpConfigJob

// NewSpConfigJob instantiates a new SpConfigJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpConfigJob(jobId string, status string, type_ string, message string, description string, expiration time.Time, created time.Time, modified time.Time, completed time.Time) *SpConfigJob {
	this := SpConfigJob{}
	this.JobId = jobId
	this.Status = status
	this.Type = type_
	this.Message = message
	this.Description = description
	this.Expiration = expiration
	this.Created = created
	this.Modified = modified
	this.Completed = completed
	return &this
}

// NewSpConfigJobWithDefaults instantiates a new SpConfigJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpConfigJobWithDefaults() *SpConfigJob {
	this := SpConfigJob{}
	return &this
}

// GetJobId returns the JobId field value
func (o *SpConfigJob) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *SpConfigJob) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *SpConfigJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SpConfigJob) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *SpConfigJob) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpConfigJob) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *SpConfigJob) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SpConfigJob) SetMessage(v string) {
	o.Message = v
}

// GetDescription returns the Description field value
func (o *SpConfigJob) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SpConfigJob) SetDescription(v string) {
	o.Description = v
}

// GetExpiration returns the Expiration field value
func (o *SpConfigJob) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *SpConfigJob) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetCreated returns the Created field value
func (o *SpConfigJob) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SpConfigJob) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *SpConfigJob) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *SpConfigJob) SetModified(v time.Time) {
	o.Modified = v
}

// GetCompleted returns the Completed field value
func (o *SpConfigJob) GetCompleted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *SpConfigJob) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *SpConfigJob) SetCompleted(v time.Time) {
	o.Completed = v
}

func (o SpConfigJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpConfigJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["message"] = o.Message
	toSerialize["description"] = o.Description
	toSerialize["expiration"] = o.Expiration
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	toSerialize["completed"] = o.Completed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpConfigJob) UnmarshalJSON(bytes []byte) (err error) {
	varSpConfigJob := _SpConfigJob{}

	if err = json.Unmarshal(bytes, &varSpConfigJob); err == nil {
		*o = SpConfigJob(varSpConfigJob)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "message")
		delete(additionalProperties, "description")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "completed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpConfigJob struct {
	value *SpConfigJob
	isSet bool
}

func (v NullableSpConfigJob) Get() *SpConfigJob {
	return v.value
}

func (v *NullableSpConfigJob) Set(val *SpConfigJob) {
	v.value = val
	v.isSet = true
}

func (v NullableSpConfigJob) IsSet() bool {
	return v.isSet
}

func (v *NullableSpConfigJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpConfigJob(val *SpConfigJob) *NullableSpConfigJob {
	return &NullableSpConfigJob{value: val, isSet: true}
}

func (v NullableSpConfigJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpConfigJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


