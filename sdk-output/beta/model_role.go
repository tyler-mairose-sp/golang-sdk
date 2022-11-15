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

// Role A Role
type Role struct {
	// The id of the Role. This field must be left null when creating an Role, otherwise a 400 Bad Request error will result.
	Id *string `json:"id,omitempty"`
	// The human-readable display name of the Role
	Name string `json:"name"`
	// Date the Role was created
	Created *time.Time `json:"created,omitempty"`
	// Date the Role was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// A human-readable description of the Role
	Description NullableString `json:"description,omitempty"`
	Owner *OwnerReference `json:"owner,omitempty"`
	AccessProfiles []AccessProfileRef `json:"accessProfiles,omitempty"`
	Membership *RoleMembershipSelector `json:"membership,omitempty"`
	// This field is not directly modifiable and is generally expected to be *null*. In very rare instances, some Roles may have been created using membership selection criteria that are no longer fully supported. While these Roles will still work, they should be migrated to STANDARD or IDENTITY_LIST selection criteria. This field exists for informational purposes as an aid to such migration.
	LegacyMembershipInfo map[string]interface{} `json:"legacyMembershipInfo,omitempty"`
	// Whether the Role is enabled or not. This field is false by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether the Role can be the target of Access Requests. This field is false by default.
	Requestable *bool `json:"requestable,omitempty"`
	AccessRequestConfig *RequestabilityForRole `json:"accessRequestConfig,omitempty"`
	RevocationRequestConfig *Revocability `json:"revocationRequestConfig,omitempty"`
	// List of IDs of segments, if any, to which this Role is assigned.
	Segments []string `json:"segments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Role Role

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole(name string) *Role {
	this := Role{}
	this.Name = name
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Role) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Role) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Role) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Role) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Role) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Role) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Role) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Role) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Role) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Role) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Role) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Role) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Role) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Role) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Role) UnsetDescription() {
	o.Description.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Role) GetOwner() OwnerReference {
	if o == nil || isNil(o.Owner) {
		var ret OwnerReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetOwnerOk() (*OwnerReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Role) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given OwnerReference and assigns it to the Owner field.
func (o *Role) SetOwner(v OwnerReference) {
	o.Owner = &v
}

// GetAccessProfiles returns the AccessProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetAccessProfiles() []AccessProfileRef {
	if o == nil {
		var ret []AccessProfileRef
		return ret
	}
	return o.AccessProfiles
}

// GetAccessProfilesOk returns a tuple with the AccessProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetAccessProfilesOk() ([]AccessProfileRef, bool) {
	if o == nil || isNil(o.AccessProfiles) {
		return nil, false
	}
	return o.AccessProfiles, true
}

// HasAccessProfiles returns a boolean if a field has been set.
func (o *Role) HasAccessProfiles() bool {
	if o != nil && isNil(o.AccessProfiles) {
		return true
	}

	return false
}

// SetAccessProfiles gets a reference to the given []AccessProfileRef and assigns it to the AccessProfiles field.
func (o *Role) SetAccessProfiles(v []AccessProfileRef) {
	o.AccessProfiles = v
}

// GetMembership returns the Membership field value if set, zero value otherwise.
func (o *Role) GetMembership() RoleMembershipSelector {
	if o == nil || isNil(o.Membership) {
		var ret RoleMembershipSelector
		return ret
	}
	return *o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetMembershipOk() (*RoleMembershipSelector, bool) {
	if o == nil || isNil(o.Membership) {
		return nil, false
	}
	return o.Membership, true
}

// HasMembership returns a boolean if a field has been set.
func (o *Role) HasMembership() bool {
	if o != nil && !isNil(o.Membership) {
		return true
	}

	return false
}

// SetMembership gets a reference to the given RoleMembershipSelector and assigns it to the Membership field.
func (o *Role) SetMembership(v RoleMembershipSelector) {
	o.Membership = &v
}

// GetLegacyMembershipInfo returns the LegacyMembershipInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetLegacyMembershipInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LegacyMembershipInfo
}

// GetLegacyMembershipInfoOk returns a tuple with the LegacyMembershipInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetLegacyMembershipInfoOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.LegacyMembershipInfo) {
		return map[string]interface{}{}, false
	}
	return o.LegacyMembershipInfo, true
}

// HasLegacyMembershipInfo returns a boolean if a field has been set.
func (o *Role) HasLegacyMembershipInfo() bool {
	if o != nil && isNil(o.LegacyMembershipInfo) {
		return true
	}

	return false
}

// SetLegacyMembershipInfo gets a reference to the given map[string]interface{} and assigns it to the LegacyMembershipInfo field.
func (o *Role) SetLegacyMembershipInfo(v map[string]interface{}) {
	o.LegacyMembershipInfo = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Role) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Role) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Role) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *Role) GetRequestable() bool {
	if o == nil || isNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRequestableOk() (*bool, bool) {
	if o == nil || isNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *Role) HasRequestable() bool {
	if o != nil && !isNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *Role) SetRequestable(v bool) {
	o.Requestable = &v
}

// GetAccessRequestConfig returns the AccessRequestConfig field value if set, zero value otherwise.
func (o *Role) GetAccessRequestConfig() RequestabilityForRole {
	if o == nil || isNil(o.AccessRequestConfig) {
		var ret RequestabilityForRole
		return ret
	}
	return *o.AccessRequestConfig
}

// GetAccessRequestConfigOk returns a tuple with the AccessRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetAccessRequestConfigOk() (*RequestabilityForRole, bool) {
	if o == nil || isNil(o.AccessRequestConfig) {
		return nil, false
	}
	return o.AccessRequestConfig, true
}

// HasAccessRequestConfig returns a boolean if a field has been set.
func (o *Role) HasAccessRequestConfig() bool {
	if o != nil && !isNil(o.AccessRequestConfig) {
		return true
	}

	return false
}

// SetAccessRequestConfig gets a reference to the given RequestabilityForRole and assigns it to the AccessRequestConfig field.
func (o *Role) SetAccessRequestConfig(v RequestabilityForRole) {
	o.AccessRequestConfig = &v
}

// GetRevocationRequestConfig returns the RevocationRequestConfig field value if set, zero value otherwise.
func (o *Role) GetRevocationRequestConfig() Revocability {
	if o == nil || isNil(o.RevocationRequestConfig) {
		var ret Revocability
		return ret
	}
	return *o.RevocationRequestConfig
}

// GetRevocationRequestConfigOk returns a tuple with the RevocationRequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRevocationRequestConfigOk() (*Revocability, bool) {
	if o == nil || isNil(o.RevocationRequestConfig) {
		return nil, false
	}
	return o.RevocationRequestConfig, true
}

// HasRevocationRequestConfig returns a boolean if a field has been set.
func (o *Role) HasRevocationRequestConfig() bool {
	if o != nil && !isNil(o.RevocationRequestConfig) {
		return true
	}

	return false
}

// SetRevocationRequestConfig gets a reference to the given Revocability and assigns it to the RevocationRequestConfig field.
func (o *Role) SetRevocationRequestConfig(v Revocability) {
	o.RevocationRequestConfig = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetSegments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetSegmentsOk() ([]string, bool) {
	if o == nil || isNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *Role) HasSegments() bool {
	if o != nil && isNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *Role) SetSegments(v []string) {
	o.Segments = v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if o.AccessProfiles != nil {
		toSerialize["accessProfiles"] = o.AccessProfiles
	}
	if !isNil(o.Membership) {
		toSerialize["membership"] = o.Membership
	}
	if o.LegacyMembershipInfo != nil {
		toSerialize["legacyMembershipInfo"] = o.LegacyMembershipInfo
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if !isNil(o.AccessRequestConfig) {
		toSerialize["accessRequestConfig"] = o.AccessRequestConfig
	}
	if !isNil(o.RevocationRequestConfig) {
		toSerialize["revocationRequestConfig"] = o.RevocationRequestConfig
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Role) UnmarshalJSON(bytes []byte) (err error) {
	varRole := _Role{}

	if err = json.Unmarshal(bytes, &varRole); err == nil {
		*o = Role(varRole)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "accessProfiles")
		delete(additionalProperties, "membership")
		delete(additionalProperties, "legacyMembershipInfo")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "accessRequestConfig")
		delete(additionalProperties, "revocationRequestConfig")
		delete(additionalProperties, "segments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


