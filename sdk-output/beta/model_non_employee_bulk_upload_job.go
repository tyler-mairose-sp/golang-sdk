/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
	"time"
)

// NonEmployeeBulkUploadJob struct for NonEmployeeBulkUploadJob
type NonEmployeeBulkUploadJob struct {
	// The bulk upload job's ID. (UUID)
	Id *string `json:"id,omitempty"`
	// The ID of the source to bulk-upload non-employees to. (UUID)
	SourceId *string `json:"sourceId,omitempty"`
	// The date-time the job was submitted.
	Created *time.Time `json:"created,omitempty"`
	// The date-time that the job was last updated.
	Modified *time.Time `json:"modified,omitempty"`
	// Returns the following values indicating the progress or result of the bulk upload job. \"PENDING\" means the job is queued and waiting to be processed. \"IN_PROGRESS\" means the job is currently being processed. \"COMPLETED\" means the job has been completed without any errors. \"ERROR\" means the job failed to process with errors. 
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeBulkUploadJob NonEmployeeBulkUploadJob

// NewNonEmployeeBulkUploadJob instantiates a new NonEmployeeBulkUploadJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeBulkUploadJob() *NonEmployeeBulkUploadJob {
	this := NonEmployeeBulkUploadJob{}
	return &this
}

// NewNonEmployeeBulkUploadJobWithDefaults instantiates a new NonEmployeeBulkUploadJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeBulkUploadJobWithDefaults() *NonEmployeeBulkUploadJob {
	this := NonEmployeeBulkUploadJob{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeBulkUploadJob) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeBulkUploadJob) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeBulkUploadJob) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeBulkUploadJob) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeBulkUploadJob) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeBulkUploadJob) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeBulkUploadJob) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeBulkUploadJob) SetSourceId(v string) {
	o.SourceId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeBulkUploadJob) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeBulkUploadJob) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeBulkUploadJob) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeBulkUploadJob) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeBulkUploadJob) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeBulkUploadJob) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeBulkUploadJob) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeBulkUploadJob) SetModified(v time.Time) {
	o.Modified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NonEmployeeBulkUploadJob) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeBulkUploadJob) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NonEmployeeBulkUploadJob) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NonEmployeeBulkUploadJob) SetStatus(v string) {
	o.Status = &v
}

func (o NonEmployeeBulkUploadJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeBulkUploadJob) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeBulkUploadJob := _NonEmployeeBulkUploadJob{}

	if err = json.Unmarshal(bytes, &varNonEmployeeBulkUploadJob); err == nil {
		*o = NonEmployeeBulkUploadJob(varNonEmployeeBulkUploadJob)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeBulkUploadJob struct {
	value *NonEmployeeBulkUploadJob
	isSet bool
}

func (v NullableNonEmployeeBulkUploadJob) Get() *NonEmployeeBulkUploadJob {
	return v.value
}

func (v *NullableNonEmployeeBulkUploadJob) Set(val *NonEmployeeBulkUploadJob) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeBulkUploadJob) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeBulkUploadJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeBulkUploadJob(val *NonEmployeeBulkUploadJob) *NullableNonEmployeeBulkUploadJob {
	return &NullableNonEmployeeBulkUploadJob{value: val, isSet: true}
}

func (v NullableNonEmployeeBulkUploadJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeBulkUploadJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


