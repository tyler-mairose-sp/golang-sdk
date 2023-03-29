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

// checks if the NonEmployeeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeSource{}

// NonEmployeeSource struct for NonEmployeeSource
type NonEmployeeSource struct {
	// Non-Employee source id.
	Id *string `json:"id,omitempty"`
	// Source Id associated with this non-employee source.
	SourceId *string `json:"sourceId,omitempty"`
	// Source name associated with this non-employee source.
	Name *string `json:"name,omitempty"`
	// Source description associated with this non-employee source.
	Description *string `json:"description,omitempty"`
	// List of approvers
	Approvers []IdentityReferenceWithId `json:"approvers,omitempty"`
	// List of account managers
	AccountManagers []IdentityReferenceWithId `json:"accountManagers,omitempty"`
	// When the request was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	// The number of non-employee records on all sources that *requested-for* user manages.
	NonEmployeeCount NullableInt32 `json:"nonEmployeeCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSource NonEmployeeSource

// NewNonEmployeeSource instantiates a new NonEmployeeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSource() *NonEmployeeSource {
	this := NonEmployeeSource{}
	return &this
}

// NewNonEmployeeSourceWithDefaults instantiates a new NonEmployeeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceWithDefaults() *NonEmployeeSource {
	this := NonEmployeeSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeSource) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeSource) SetSourceId(v string) {
	o.SourceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NonEmployeeSource) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonEmployeeSource) SetDescription(v string) {
	o.Description = &v
}

// GetApprovers returns the Approvers field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetApprovers() []IdentityReferenceWithId {
	if o == nil || isNil(o.Approvers) {
		var ret []IdentityReferenceWithId
		return ret
	}
	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetApproversOk() ([]IdentityReferenceWithId, bool) {
	if o == nil || isNil(o.Approvers) {
		return nil, false
	}
	return o.Approvers, true
}

// HasApprovers returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasApprovers() bool {
	if o != nil && !isNil(o.Approvers) {
		return true
	}

	return false
}

// SetApprovers gets a reference to the given []IdentityReferenceWithId and assigns it to the Approvers field.
func (o *NonEmployeeSource) SetApprovers(v []IdentityReferenceWithId) {
	o.Approvers = v
}

// GetAccountManagers returns the AccountManagers field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetAccountManagers() []IdentityReferenceWithId {
	if o == nil || isNil(o.AccountManagers) {
		var ret []IdentityReferenceWithId
		return ret
	}
	return o.AccountManagers
}

// GetAccountManagersOk returns a tuple with the AccountManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetAccountManagersOk() ([]IdentityReferenceWithId, bool) {
	if o == nil || isNil(o.AccountManagers) {
		return nil, false
	}
	return o.AccountManagers, true
}

// HasAccountManagers returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasAccountManagers() bool {
	if o != nil && !isNil(o.AccountManagers) {
		return true
	}

	return false
}

// SetAccountManagers gets a reference to the given []IdentityReferenceWithId and assigns it to the AccountManagers field.
func (o *NonEmployeeSource) SetAccountManagers(v []IdentityReferenceWithId) {
	o.AccountManagers = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeSource) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeSource) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSource) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeSource) SetCreated(v time.Time) {
	o.Created = &v
}

// GetNonEmployeeCount returns the NonEmployeeCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NonEmployeeSource) GetNonEmployeeCount() int32 {
	if o == nil || isNil(o.NonEmployeeCount.Get()) {
		var ret int32
		return ret
	}
	return *o.NonEmployeeCount.Get()
}

// GetNonEmployeeCountOk returns a tuple with the NonEmployeeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NonEmployeeSource) GetNonEmployeeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonEmployeeCount.Get(), o.NonEmployeeCount.IsSet()
}

// HasNonEmployeeCount returns a boolean if a field has been set.
func (o *NonEmployeeSource) HasNonEmployeeCount() bool {
	if o != nil && o.NonEmployeeCount.IsSet() {
		return true
	}

	return false
}

// SetNonEmployeeCount gets a reference to the given NullableInt32 and assigns it to the NonEmployeeCount field.
func (o *NonEmployeeSource) SetNonEmployeeCount(v int32) {
	o.NonEmployeeCount.Set(&v)
}
// SetNonEmployeeCountNil sets the value for NonEmployeeCount to be an explicit nil
func (o *NonEmployeeSource) SetNonEmployeeCountNil() {
	o.NonEmployeeCount.Set(nil)
}

// UnsetNonEmployeeCount ensures that no value is present for NonEmployeeCount, not even an explicit nil
func (o *NonEmployeeSource) UnsetNonEmployeeCount() {
	o.NonEmployeeCount.Unset()
}

func (o NonEmployeeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Approvers) {
		toSerialize["approvers"] = o.Approvers
	}
	if !isNil(o.AccountManagers) {
		toSerialize["accountManagers"] = o.AccountManagers
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.NonEmployeeCount.IsSet() {
		toSerialize["nonEmployeeCount"] = o.NonEmployeeCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeSource) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeSource := _NonEmployeeSource{}

	if err = json.Unmarshal(bytes, &varNonEmployeeSource); err == nil {
		*o = NonEmployeeSource(varNonEmployeeSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approvers")
		delete(additionalProperties, "accountManagers")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		delete(additionalProperties, "nonEmployeeCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSource struct {
	value *NonEmployeeSource
	isSet bool
}

func (v NullableNonEmployeeSource) Get() *NonEmployeeSource {
	return v.value
}

func (v *NullableNonEmployeeSource) Set(val *NonEmployeeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSource(val *NonEmployeeSource) *NullableNonEmployeeSource {
	return &NullableNonEmployeeSource{value: val, isSet: true}
}

func (v NullableNonEmployeeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


