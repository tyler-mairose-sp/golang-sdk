/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ReviewableEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewableEntitlement{}

// ReviewableEntitlement struct for ReviewableEntitlement
type ReviewableEntitlement struct {
	// The id for the entitlement
	Id *string `json:"id,omitempty"`
	// The name of the entitlement
	Name *string `json:"name,omitempty"`
	// Information about the entitlement
	Description NullableString `json:"description,omitempty"`
	// Indicates if the entitlement is a privileged entitlement
	Privileged *bool `json:"privileged,omitempty"`
	Owner NullableIdentityReferenceWithNameAndEmail `json:"owner,omitempty"`
	// The name of the attribute on the source
	AttributeName *string `json:"attributeName,omitempty"`
	// The value of the attribute on the source
	AttributeValue *string `json:"attributeValue,omitempty"`
	// The schema object type on the source used to represent the entitlement and its attributes
	SourceSchemaObjectType *string `json:"sourceSchemaObjectType,omitempty"`
	// The name of the source for which this entitlement belongs
	SourceName *string `json:"sourceName,omitempty"`
	// The type of the source for which the entitlement belongs
	SourceType *string `json:"sourceType,omitempty"`
	// Indicates if the entitlement has permissions
	HasPermissions *bool `json:"hasPermissions,omitempty"`
	// Indicates if the entitlement is a representation of an account permission
	IsPermission *bool `json:"isPermission,omitempty"`
	// Indicates whether the entitlement can be revoked
	Revocable *bool `json:"revocable,omitempty"`
	// True if the entitlement is cloud governed
	CloudGoverned *bool `json:"cloudGoverned,omitempty"`
	Account NullableReviewableEntitlementAccount `json:"account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewableEntitlement ReviewableEntitlement

// NewReviewableEntitlement instantiates a new ReviewableEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewableEntitlement() *ReviewableEntitlement {
	this := ReviewableEntitlement{}
	return &this
}

// NewReviewableEntitlementWithDefaults instantiates a new ReviewableEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewableEntitlementWithDefaults() *ReviewableEntitlement {
	this := ReviewableEntitlement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReviewableEntitlement) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewableEntitlement) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlement) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlement) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ReviewableEntitlement) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ReviewableEntitlement) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ReviewableEntitlement) UnsetDescription() {
	o.Description.Unset()
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *ReviewableEntitlement) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlement) GetOwner() IdentityReferenceWithNameAndEmail {
	if o == nil || isNil(o.Owner.Get()) {
		var ret IdentityReferenceWithNameAndEmail
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlement) GetOwnerOk() (*IdentityReferenceWithNameAndEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableIdentityReferenceWithNameAndEmail and assigns it to the Owner field.
func (o *ReviewableEntitlement) SetOwner(v IdentityReferenceWithNameAndEmail) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ReviewableEntitlement) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ReviewableEntitlement) UnsetOwner() {
	o.Owner.Unset()
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetAttributeName() string {
	if o == nil || isNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasAttributeName() bool {
	if o != nil && !isNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *ReviewableEntitlement) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetAttributeValue() string {
	if o == nil || isNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetAttributeValueOk() (*string, bool) {
	if o == nil || isNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasAttributeValue() bool {
	if o != nil && !isNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *ReviewableEntitlement) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetSourceSchemaObjectType returns the SourceSchemaObjectType field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetSourceSchemaObjectType() string {
	if o == nil || isNil(o.SourceSchemaObjectType) {
		var ret string
		return ret
	}
	return *o.SourceSchemaObjectType
}

// GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetSourceSchemaObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.SourceSchemaObjectType) {
		return nil, false
	}
	return o.SourceSchemaObjectType, true
}

// HasSourceSchemaObjectType returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasSourceSchemaObjectType() bool {
	if o != nil && !isNil(o.SourceSchemaObjectType) {
		return true
	}

	return false
}

// SetSourceSchemaObjectType gets a reference to the given string and assigns it to the SourceSchemaObjectType field.
func (o *ReviewableEntitlement) SetSourceSchemaObjectType(v string) {
	o.SourceSchemaObjectType = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *ReviewableEntitlement) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetSourceType() string {
	if o == nil || isNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetSourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasSourceType() bool {
	if o != nil && !isNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *ReviewableEntitlement) SetSourceType(v string) {
	o.SourceType = &v
}

// GetHasPermissions returns the HasPermissions field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetHasPermissions() bool {
	if o == nil || isNil(o.HasPermissions) {
		var ret bool
		return ret
	}
	return *o.HasPermissions
}

// GetHasPermissionsOk returns a tuple with the HasPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetHasPermissionsOk() (*bool, bool) {
	if o == nil || isNil(o.HasPermissions) {
		return nil, false
	}
	return o.HasPermissions, true
}

// HasHasPermissions returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasHasPermissions() bool {
	if o != nil && !isNil(o.HasPermissions) {
		return true
	}

	return false
}

// SetHasPermissions gets a reference to the given bool and assigns it to the HasPermissions field.
func (o *ReviewableEntitlement) SetHasPermissions(v bool) {
	o.HasPermissions = &v
}

// GetIsPermission returns the IsPermission field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetIsPermission() bool {
	if o == nil || isNil(o.IsPermission) {
		var ret bool
		return ret
	}
	return *o.IsPermission
}

// GetIsPermissionOk returns a tuple with the IsPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetIsPermissionOk() (*bool, bool) {
	if o == nil || isNil(o.IsPermission) {
		return nil, false
	}
	return o.IsPermission, true
}

// HasIsPermission returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasIsPermission() bool {
	if o != nil && !isNil(o.IsPermission) {
		return true
	}

	return false
}

// SetIsPermission gets a reference to the given bool and assigns it to the IsPermission field.
func (o *ReviewableEntitlement) SetIsPermission(v bool) {
	o.IsPermission = &v
}

// GetRevocable returns the Revocable field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetRevocable() bool {
	if o == nil || isNil(o.Revocable) {
		var ret bool
		return ret
	}
	return *o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetRevocableOk() (*bool, bool) {
	if o == nil || isNil(o.Revocable) {
		return nil, false
	}
	return o.Revocable, true
}

// HasRevocable returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasRevocable() bool {
	if o != nil && !isNil(o.Revocable) {
		return true
	}

	return false
}

// SetRevocable gets a reference to the given bool and assigns it to the Revocable field.
func (o *ReviewableEntitlement) SetRevocable(v bool) {
	o.Revocable = &v
}

// GetCloudGoverned returns the CloudGoverned field value if set, zero value otherwise.
func (o *ReviewableEntitlement) GetCloudGoverned() bool {
	if o == nil || isNil(o.CloudGoverned) {
		var ret bool
		return ret
	}
	return *o.CloudGoverned
}

// GetCloudGovernedOk returns a tuple with the CloudGoverned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewableEntitlement) GetCloudGovernedOk() (*bool, bool) {
	if o == nil || isNil(o.CloudGoverned) {
		return nil, false
	}
	return o.CloudGoverned, true
}

// HasCloudGoverned returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasCloudGoverned() bool {
	if o != nil && !isNil(o.CloudGoverned) {
		return true
	}

	return false
}

// SetCloudGoverned gets a reference to the given bool and assigns it to the CloudGoverned field.
func (o *ReviewableEntitlement) SetCloudGoverned(v bool) {
	o.CloudGoverned = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewableEntitlement) GetAccount() ReviewableEntitlementAccount {
	if o == nil || isNil(o.Account.Get()) {
		var ret ReviewableEntitlementAccount
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewableEntitlement) GetAccountOk() (*ReviewableEntitlementAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ReviewableEntitlement) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableReviewableEntitlementAccount and assigns it to the Account field.
func (o *ReviewableEntitlement) SetAccount(v ReviewableEntitlementAccount) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *ReviewableEntitlement) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ReviewableEntitlement) UnsetAccount() {
	o.Account.Unset()
}

func (o ReviewableEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewableEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if !isNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !isNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !isNil(o.SourceSchemaObjectType) {
		toSerialize["sourceSchemaObjectType"] = o.SourceSchemaObjectType
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !isNil(o.HasPermissions) {
		toSerialize["hasPermissions"] = o.HasPermissions
	}
	if !isNil(o.IsPermission) {
		toSerialize["isPermission"] = o.IsPermission
	}
	if !isNil(o.Revocable) {
		toSerialize["revocable"] = o.Revocable
	}
	if !isNil(o.CloudGoverned) {
		toSerialize["cloudGoverned"] = o.CloudGoverned
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewableEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varReviewableEntitlement := _ReviewableEntitlement{}

	if err = json.Unmarshal(bytes, &varReviewableEntitlement); err == nil {
		*o = ReviewableEntitlement(varReviewableEntitlement)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "attributeValue")
		delete(additionalProperties, "sourceSchemaObjectType")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "sourceType")
		delete(additionalProperties, "hasPermissions")
		delete(additionalProperties, "isPermission")
		delete(additionalProperties, "revocable")
		delete(additionalProperties, "cloudGoverned")
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewableEntitlement struct {
	value *ReviewableEntitlement
	isSet bool
}

func (v NullableReviewableEntitlement) Get() *ReviewableEntitlement {
	return v.value
}

func (v *NullableReviewableEntitlement) Set(val *ReviewableEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewableEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewableEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewableEntitlement(val *ReviewableEntitlement) *NullableReviewableEntitlement {
	return &NullableReviewableEntitlement{value: val, isSet: true}
}

func (v NullableReviewableEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewableEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


