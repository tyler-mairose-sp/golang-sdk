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

// checks if the ReviewableAccessProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewableAccessProfile{}

// ReviewableAccessProfile struct for ReviewableAccessProfile
type ReviewableAccessProfile struct {
	// The id of the Access Profile
	Id *string `json:"id,omitempty"`
	// Name of the Access Profile
	Name *string `json:"name,omitempty"`
	// Information about the Access Profile
	Description *string `json:"description,omitempty"`
	// Indicates if the entitlement is a privileged entitlement
	Privileged *bool `json:"privileged,omitempty"`
	// True if the entitlement is cloud governed
	CloudGoverned *bool `json:"cloudGoverned,omitempty"`
	// The date at which a user's access expires
	EndDate NullableTime `json:"endDate,omitempty"`
	Owner NullableIdentityReferenceWithNameAndEmail `json:"owner,omitempty"`
	// A list of entitlements associated with this Access Profile
	Entitlements []ReviewableEntitlement `json:"entitlements,omitempty"`
	// Date the Access Profile was created.
	Created *time.Time `json:"created,omitempty"`
	// Date the Access Profile was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewableAccessProfile ReviewableAccessProfile

// NewReviewableAccessProfile instantiates a new ReviewableAccessProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewableAccessProfile() *ReviewableAccessProfile {
	this := ReviewableAccessProfile{}
	return &this
}

// NewReviewableAccessProfileWithDefaults instantiates a new ReviewableAccessProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewableAccessProfileWithDefaults() *ReviewableAccessProfile {
	this := ReviewableAccessProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReviewableAccessProfile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewableAccessProfile) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReviewableAccessProfile) SetDescription(v string) {
	o.Description = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *ReviewableAccessProfile) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetCloudGoverned returns the CloudGoverned field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetCloudGoverned() bool {
	if o == nil || isNil(o.CloudGoverned) {
		var ret bool
		return ret
	}
	return *o.CloudGoverned
}

// GetCloudGovernedOk returns a tuple with the CloudGoverned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetCloudGovernedOk() (*bool, bool) {
	if o == nil || isNil(o.CloudGoverned) {
		return nil, false
	}
	return o.CloudGoverned, true
}

// HasCloudGoverned returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasCloudGoverned() bool {
	if o != nil && !isNil(o.CloudGoverned) {
		return true
	}

	return false
}

// SetCloudGoverned gets a reference to the given bool and assigns it to the CloudGoverned field.
func (o *ReviewableAccessProfile) SetCloudGoverned(v bool) {
	o.CloudGoverned = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableAccessProfile) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableAccessProfile) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *ReviewableAccessProfile) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ReviewableAccessProfile) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ReviewableAccessProfile) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableAccessProfile) GetOwner() IdentityReferenceWithNameAndEmail {
	if o == nil || isNil(o.Owner.Get()) {
		var ret IdentityReferenceWithNameAndEmail
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableAccessProfile) GetOwnerOk() (*IdentityReferenceWithNameAndEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableIdentityReferenceWithNameAndEmail and assigns it to the Owner field.
func (o *ReviewableAccessProfile) SetOwner(v IdentityReferenceWithNameAndEmail) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ReviewableAccessProfile) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ReviewableAccessProfile) UnsetOwner() {
	o.Owner.Unset()
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetEntitlements() []ReviewableEntitlement {
	if o == nil || isNil(o.Entitlements) {
		var ret []ReviewableEntitlement
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetEntitlementsOk() ([]ReviewableEntitlement, bool) {
	if o == nil || isNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []ReviewableEntitlement and assigns it to the Entitlements field.
func (o *ReviewableAccessProfile) SetEntitlements(v []ReviewableEntitlement) {
	o.Entitlements = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ReviewableAccessProfile) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ReviewableAccessProfile) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableAccessProfile) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ReviewableAccessProfile) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ReviewableAccessProfile) SetModified(v time.Time) {
	o.Modified = &v
}

func (o ReviewableAccessProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewableAccessProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.CloudGoverned) {
		toSerialize["cloudGoverned"] = o.CloudGoverned
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if !isNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewableAccessProfile) UnmarshalJSON(bytes []byte) (err error) {
	varReviewableAccessProfile := _ReviewableAccessProfile{}

	if err = json.Unmarshal(bytes, &varReviewableAccessProfile); err == nil {
		*o = ReviewableAccessProfile(varReviewableAccessProfile)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "cloudGoverned")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewableAccessProfile struct {
	value *ReviewableAccessProfile
	isSet bool
}

func (v NullableReviewableAccessProfile) Get() *ReviewableAccessProfile {
	return v.value
}

func (v *NullableReviewableAccessProfile) Set(val *ReviewableAccessProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewableAccessProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewableAccessProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewableAccessProfile(val *ReviewableAccessProfile) *NullableReviewableAccessProfile {
	return &NullableReviewableAccessProfile{value: val, isSet: true}
}

func (v NullableReviewableAccessProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewableAccessProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


