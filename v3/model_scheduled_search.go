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

// checks if the ScheduledSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledSearch{}

// ScheduledSearch struct for ScheduledSearch
type ScheduledSearch struct {
	// The scheduled search ID.
	Id string `json:"id"`
	Owner ScheduledSearchAllOfOwner `json:"owner"`
	// The ID of the scheduled search owner.  Please use the `id` in the `owner` object instead. 
	// Deprecated
	OwnerId string `json:"ownerId"`
	// The name of the scheduled search. 
	Name NullableString `json:"name,omitempty"`
	// The description of the scheduled search. 
	Description NullableString `json:"description,omitempty"`
	// The ID of the saved search that will be executed.
	SavedSearchId string `json:"savedSearchId"`
	// The date the scheduled search was initially created.
	Created *time.Time `json:"created,omitempty"`
	// The last date the scheduled search was modified.
	Modified *time.Time `json:"modified,omitempty"`
	Schedule Schedule `json:"schedule"`
	// A list of identities that should receive the scheduled search report via email.
	Recipients []SearchScheduleRecipientsInner `json:"recipients"`
	// Indicates if the scheduled search is enabled. 
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates if email generation should not be suppressed if search returns no results. 
	EmailEmptyResults *bool `json:"emailEmptyResults,omitempty"`
	// Indicates if the generated email should include the query and search results preview (which could include PII). 
	DisplayQueryDetails *bool `json:"displayQueryDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledSearch ScheduledSearch

// NewScheduledSearch instantiates a new ScheduledSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledSearch(id string, owner ScheduledSearchAllOfOwner, ownerId string, savedSearchId string, schedule Schedule, recipients []SearchScheduleRecipientsInner) *ScheduledSearch {
	this := ScheduledSearch{}
	this.Id = id
	this.Owner = owner
	this.OwnerId = ownerId
	this.SavedSearchId = savedSearchId
	this.Schedule = schedule
	this.Recipients = recipients
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// NewScheduledSearchWithDefaults instantiates a new ScheduledSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledSearchWithDefaults() *ScheduledSearch {
	this := ScheduledSearch{}
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// GetId returns the Id field value
func (o *ScheduledSearch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScheduledSearch) SetId(v string) {
	o.Id = v
}

// GetOwner returns the Owner field value
func (o *ScheduledSearch) GetOwner() ScheduledSearchAllOfOwner {
	if o == nil {
		var ret ScheduledSearchAllOfOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetOwnerOk() (*ScheduledSearchAllOfOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *ScheduledSearch) SetOwner(v ScheduledSearchAllOfOwner) {
	o.Owner = v
}

// GetOwnerId returns the OwnerId field value
// Deprecated
func (o *ScheduledSearch) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ScheduledSearch) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
// Deprecated
func (o *ScheduledSearch) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledSearch) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledSearch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ScheduledSearch) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ScheduledSearch) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ScheduledSearch) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ScheduledSearch) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledSearch) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledSearch) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ScheduledSearch) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ScheduledSearch) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ScheduledSearch) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ScheduledSearch) UnsetDescription() {
	o.Description.Unset()
}

// GetSavedSearchId returns the SavedSearchId field value
func (o *ScheduledSearch) GetSavedSearchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SavedSearchId
}

// GetSavedSearchIdOk returns a tuple with the SavedSearchId field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetSavedSearchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavedSearchId, true
}

// SetSavedSearchId sets field value
func (o *ScheduledSearch) SetSavedSearchId(v string) {
	o.SavedSearchId = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ScheduledSearch) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ScheduledSearch) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ScheduledSearch) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ScheduledSearch) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ScheduledSearch) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ScheduledSearch) SetModified(v time.Time) {
	o.Modified = &v
}

// GetSchedule returns the Schedule field value
func (o *ScheduledSearch) GetSchedule() Schedule {
	if o == nil {
		var ret Schedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetScheduleOk() (*Schedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ScheduledSearch) SetSchedule(v Schedule) {
	o.Schedule = v
}

// GetRecipients returns the Recipients field value
func (o *ScheduledSearch) GetRecipients() []SearchScheduleRecipientsInner {
	if o == nil {
		var ret []SearchScheduleRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetRecipientsOk() ([]SearchScheduleRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ScheduledSearch) SetRecipients(v []SearchScheduleRecipientsInner) {
	o.Recipients = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScheduledSearch) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScheduledSearch) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScheduledSearch) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEmailEmptyResults returns the EmailEmptyResults field value if set, zero value otherwise.
func (o *ScheduledSearch) GetEmailEmptyResults() bool {
	if o == nil || isNil(o.EmailEmptyResults) {
		var ret bool
		return ret
	}
	return *o.EmailEmptyResults
}

// GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetEmailEmptyResultsOk() (*bool, bool) {
	if o == nil || isNil(o.EmailEmptyResults) {
		return nil, false
	}
	return o.EmailEmptyResults, true
}

// HasEmailEmptyResults returns a boolean if a field has been set.
func (o *ScheduledSearch) HasEmailEmptyResults() bool {
	if o != nil && !isNil(o.EmailEmptyResults) {
		return true
	}

	return false
}

// SetEmailEmptyResults gets a reference to the given bool and assigns it to the EmailEmptyResults field.
func (o *ScheduledSearch) SetEmailEmptyResults(v bool) {
	o.EmailEmptyResults = &v
}

// GetDisplayQueryDetails returns the DisplayQueryDetails field value if set, zero value otherwise.
func (o *ScheduledSearch) GetDisplayQueryDetails() bool {
	if o == nil || isNil(o.DisplayQueryDetails) {
		var ret bool
		return ret
	}
	return *o.DisplayQueryDetails
}

// GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetDisplayQueryDetailsOk() (*bool, bool) {
	if o == nil || isNil(o.DisplayQueryDetails) {
		return nil, false
	}
	return o.DisplayQueryDetails, true
}

// HasDisplayQueryDetails returns a boolean if a field has been set.
func (o *ScheduledSearch) HasDisplayQueryDetails() bool {
	if o != nil && !isNil(o.DisplayQueryDetails) {
		return true
	}

	return false
}

// SetDisplayQueryDetails gets a reference to the given bool and assigns it to the DisplayQueryDetails field.
func (o *ScheduledSearch) SetDisplayQueryDetails(v bool) {
	o.DisplayQueryDetails = &v
}

func (o ScheduledSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["owner"] = o.Owner
	// skip: ownerId is readOnly
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["savedSearchId"] = o.SavedSearchId
	// skip: created is readOnly
	// skip: modified is readOnly
	toSerialize["schedule"] = o.Schedule
	toSerialize["recipients"] = o.Recipients
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.EmailEmptyResults) {
		toSerialize["emailEmptyResults"] = o.EmailEmptyResults
	}
	if !isNil(o.DisplayQueryDetails) {
		toSerialize["displayQueryDetails"] = o.DisplayQueryDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduledSearch) UnmarshalJSON(bytes []byte) (err error) {
	varScheduledSearch := _ScheduledSearch{}

	if err = json.Unmarshal(bytes, &varScheduledSearch); err == nil {
		*o = ScheduledSearch(varScheduledSearch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "savedSearchId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "emailEmptyResults")
		delete(additionalProperties, "displayQueryDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledSearch struct {
	value *ScheduledSearch
	isSet bool
}

func (v NullableScheduledSearch) Get() *ScheduledSearch {
	return v.value
}

func (v *NullableScheduledSearch) Set(val *ScheduledSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledSearch(val *ScheduledSearch) *NullableScheduledSearch {
	return &NullableScheduledSearch{value: val, isSet: true}
}

func (v NullableScheduledSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


