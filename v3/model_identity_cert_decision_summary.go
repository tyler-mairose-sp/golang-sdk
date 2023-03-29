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

// checks if the IdentityCertDecisionSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityCertDecisionSummary{}

// IdentityCertDecisionSummary struct for IdentityCertDecisionSummary
type IdentityCertDecisionSummary struct {
	// Number of entitlement decisions that have been made
	EntitlementDecisionsMade *int32 `json:"entitlementDecisionsMade,omitempty"`
	// Number of access profile decisions that have been made
	AccessProfileDecisionsMade *int32 `json:"accessProfileDecisionsMade,omitempty"`
	// Number of role decisions that have been made
	RoleDecisionsMade *int32 `json:"roleDecisionsMade,omitempty"`
	// Number of account decisions that have been made
	AccountDecisionsMade *int32 `json:"accountDecisionsMade,omitempty"`
	// The total number of entitlement decisions on the certification, both complete and incomplete
	EntitlementDecisionsTotal *int32 `json:"entitlementDecisionsTotal,omitempty"`
	// The total number of access profile decisions on the certification, both complete and incomplete
	AccessProfileDecisionsTotal *int32 `json:"accessProfileDecisionsTotal,omitempty"`
	// The total number of role decisions on the certification, both complete and incomplete
	RoleDecisionsTotal *int32 `json:"roleDecisionsTotal,omitempty"`
	// The total number of account decisions on the certification, both complete and incomplete
	AccountDecisionsTotal *int32 `json:"accountDecisionsTotal,omitempty"`
	// The number of entitlement decisions that have been made which were approved
	EntitlementsApproved *int32 `json:"entitlementsApproved,omitempty"`
	// The number of entitlement decisions that have been made which were revoked
	EntitlementsRevoked *int32 `json:"entitlementsRevoked,omitempty"`
	// The number of access profile decisions that have been made which were approved
	AccessProfilesApproved *int32 `json:"accessProfilesApproved,omitempty"`
	// The number of access profile decisions that have been made which were revoked
	AccessProfilesRevoked *int32 `json:"accessProfilesRevoked,omitempty"`
	// The number of role decisions that have been made which were approved
	RolesApproved *int32 `json:"rolesApproved,omitempty"`
	// The number of role decisions that have been made which were revoked
	RolesRevoked *int32 `json:"rolesRevoked,omitempty"`
	// The number of account decisions that have been made which were approved
	AccountsApproved *int32 `json:"accountsApproved,omitempty"`
	// The number of account decisions that have been made which were revoked
	AccountsRevoked *int32 `json:"accountsRevoked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityCertDecisionSummary IdentityCertDecisionSummary

// NewIdentityCertDecisionSummary instantiates a new IdentityCertDecisionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCertDecisionSummary() *IdentityCertDecisionSummary {
	this := IdentityCertDecisionSummary{}
	return &this
}

// NewIdentityCertDecisionSummaryWithDefaults instantiates a new IdentityCertDecisionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCertDecisionSummaryWithDefaults() *IdentityCertDecisionSummary {
	this := IdentityCertDecisionSummary{}
	return &this
}

// GetEntitlementDecisionsMade returns the EntitlementDecisionsMade field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetEntitlementDecisionsMade() int32 {
	if o == nil || isNil(o.EntitlementDecisionsMade) {
		var ret int32
		return ret
	}
	return *o.EntitlementDecisionsMade
}

// GetEntitlementDecisionsMadeOk returns a tuple with the EntitlementDecisionsMade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetEntitlementDecisionsMadeOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementDecisionsMade) {
		return nil, false
	}
	return o.EntitlementDecisionsMade, true
}

// HasEntitlementDecisionsMade returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasEntitlementDecisionsMade() bool {
	if o != nil && !isNil(o.EntitlementDecisionsMade) {
		return true
	}

	return false
}

// SetEntitlementDecisionsMade gets a reference to the given int32 and assigns it to the EntitlementDecisionsMade field.
func (o *IdentityCertDecisionSummary) SetEntitlementDecisionsMade(v int32) {
	o.EntitlementDecisionsMade = &v
}

// GetAccessProfileDecisionsMade returns the AccessProfileDecisionsMade field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccessProfileDecisionsMade() int32 {
	if o == nil || isNil(o.AccessProfileDecisionsMade) {
		var ret int32
		return ret
	}
	return *o.AccessProfileDecisionsMade
}

// GetAccessProfileDecisionsMadeOk returns a tuple with the AccessProfileDecisionsMade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccessProfileDecisionsMadeOk() (*int32, bool) {
	if o == nil || isNil(o.AccessProfileDecisionsMade) {
		return nil, false
	}
	return o.AccessProfileDecisionsMade, true
}

// HasAccessProfileDecisionsMade returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccessProfileDecisionsMade() bool {
	if o != nil && !isNil(o.AccessProfileDecisionsMade) {
		return true
	}

	return false
}

// SetAccessProfileDecisionsMade gets a reference to the given int32 and assigns it to the AccessProfileDecisionsMade field.
func (o *IdentityCertDecisionSummary) SetAccessProfileDecisionsMade(v int32) {
	o.AccessProfileDecisionsMade = &v
}

// GetRoleDecisionsMade returns the RoleDecisionsMade field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetRoleDecisionsMade() int32 {
	if o == nil || isNil(o.RoleDecisionsMade) {
		var ret int32
		return ret
	}
	return *o.RoleDecisionsMade
}

// GetRoleDecisionsMadeOk returns a tuple with the RoleDecisionsMade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetRoleDecisionsMadeOk() (*int32, bool) {
	if o == nil || isNil(o.RoleDecisionsMade) {
		return nil, false
	}
	return o.RoleDecisionsMade, true
}

// HasRoleDecisionsMade returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasRoleDecisionsMade() bool {
	if o != nil && !isNil(o.RoleDecisionsMade) {
		return true
	}

	return false
}

// SetRoleDecisionsMade gets a reference to the given int32 and assigns it to the RoleDecisionsMade field.
func (o *IdentityCertDecisionSummary) SetRoleDecisionsMade(v int32) {
	o.RoleDecisionsMade = &v
}

// GetAccountDecisionsMade returns the AccountDecisionsMade field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccountDecisionsMade() int32 {
	if o == nil || isNil(o.AccountDecisionsMade) {
		var ret int32
		return ret
	}
	return *o.AccountDecisionsMade
}

// GetAccountDecisionsMadeOk returns a tuple with the AccountDecisionsMade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccountDecisionsMadeOk() (*int32, bool) {
	if o == nil || isNil(o.AccountDecisionsMade) {
		return nil, false
	}
	return o.AccountDecisionsMade, true
}

// HasAccountDecisionsMade returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccountDecisionsMade() bool {
	if o != nil && !isNil(o.AccountDecisionsMade) {
		return true
	}

	return false
}

// SetAccountDecisionsMade gets a reference to the given int32 and assigns it to the AccountDecisionsMade field.
func (o *IdentityCertDecisionSummary) SetAccountDecisionsMade(v int32) {
	o.AccountDecisionsMade = &v
}

// GetEntitlementDecisionsTotal returns the EntitlementDecisionsTotal field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetEntitlementDecisionsTotal() int32 {
	if o == nil || isNil(o.EntitlementDecisionsTotal) {
		var ret int32
		return ret
	}
	return *o.EntitlementDecisionsTotal
}

// GetEntitlementDecisionsTotalOk returns a tuple with the EntitlementDecisionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetEntitlementDecisionsTotalOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementDecisionsTotal) {
		return nil, false
	}
	return o.EntitlementDecisionsTotal, true
}

// HasEntitlementDecisionsTotal returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasEntitlementDecisionsTotal() bool {
	if o != nil && !isNil(o.EntitlementDecisionsTotal) {
		return true
	}

	return false
}

// SetEntitlementDecisionsTotal gets a reference to the given int32 and assigns it to the EntitlementDecisionsTotal field.
func (o *IdentityCertDecisionSummary) SetEntitlementDecisionsTotal(v int32) {
	o.EntitlementDecisionsTotal = &v
}

// GetAccessProfileDecisionsTotal returns the AccessProfileDecisionsTotal field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccessProfileDecisionsTotal() int32 {
	if o == nil || isNil(o.AccessProfileDecisionsTotal) {
		var ret int32
		return ret
	}
	return *o.AccessProfileDecisionsTotal
}

// GetAccessProfileDecisionsTotalOk returns a tuple with the AccessProfileDecisionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccessProfileDecisionsTotalOk() (*int32, bool) {
	if o == nil || isNil(o.AccessProfileDecisionsTotal) {
		return nil, false
	}
	return o.AccessProfileDecisionsTotal, true
}

// HasAccessProfileDecisionsTotal returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccessProfileDecisionsTotal() bool {
	if o != nil && !isNil(o.AccessProfileDecisionsTotal) {
		return true
	}

	return false
}

// SetAccessProfileDecisionsTotal gets a reference to the given int32 and assigns it to the AccessProfileDecisionsTotal field.
func (o *IdentityCertDecisionSummary) SetAccessProfileDecisionsTotal(v int32) {
	o.AccessProfileDecisionsTotal = &v
}

// GetRoleDecisionsTotal returns the RoleDecisionsTotal field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetRoleDecisionsTotal() int32 {
	if o == nil || isNil(o.RoleDecisionsTotal) {
		var ret int32
		return ret
	}
	return *o.RoleDecisionsTotal
}

// GetRoleDecisionsTotalOk returns a tuple with the RoleDecisionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetRoleDecisionsTotalOk() (*int32, bool) {
	if o == nil || isNil(o.RoleDecisionsTotal) {
		return nil, false
	}
	return o.RoleDecisionsTotal, true
}

// HasRoleDecisionsTotal returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasRoleDecisionsTotal() bool {
	if o != nil && !isNil(o.RoleDecisionsTotal) {
		return true
	}

	return false
}

// SetRoleDecisionsTotal gets a reference to the given int32 and assigns it to the RoleDecisionsTotal field.
func (o *IdentityCertDecisionSummary) SetRoleDecisionsTotal(v int32) {
	o.RoleDecisionsTotal = &v
}

// GetAccountDecisionsTotal returns the AccountDecisionsTotal field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccountDecisionsTotal() int32 {
	if o == nil || isNil(o.AccountDecisionsTotal) {
		var ret int32
		return ret
	}
	return *o.AccountDecisionsTotal
}

// GetAccountDecisionsTotalOk returns a tuple with the AccountDecisionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccountDecisionsTotalOk() (*int32, bool) {
	if o == nil || isNil(o.AccountDecisionsTotal) {
		return nil, false
	}
	return o.AccountDecisionsTotal, true
}

// HasAccountDecisionsTotal returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccountDecisionsTotal() bool {
	if o != nil && !isNil(o.AccountDecisionsTotal) {
		return true
	}

	return false
}

// SetAccountDecisionsTotal gets a reference to the given int32 and assigns it to the AccountDecisionsTotal field.
func (o *IdentityCertDecisionSummary) SetAccountDecisionsTotal(v int32) {
	o.AccountDecisionsTotal = &v
}

// GetEntitlementsApproved returns the EntitlementsApproved field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetEntitlementsApproved() int32 {
	if o == nil || isNil(o.EntitlementsApproved) {
		var ret int32
		return ret
	}
	return *o.EntitlementsApproved
}

// GetEntitlementsApprovedOk returns a tuple with the EntitlementsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetEntitlementsApprovedOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementsApproved) {
		return nil, false
	}
	return o.EntitlementsApproved, true
}

// HasEntitlementsApproved returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasEntitlementsApproved() bool {
	if o != nil && !isNil(o.EntitlementsApproved) {
		return true
	}

	return false
}

// SetEntitlementsApproved gets a reference to the given int32 and assigns it to the EntitlementsApproved field.
func (o *IdentityCertDecisionSummary) SetEntitlementsApproved(v int32) {
	o.EntitlementsApproved = &v
}

// GetEntitlementsRevoked returns the EntitlementsRevoked field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetEntitlementsRevoked() int32 {
	if o == nil || isNil(o.EntitlementsRevoked) {
		var ret int32
		return ret
	}
	return *o.EntitlementsRevoked
}

// GetEntitlementsRevokedOk returns a tuple with the EntitlementsRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetEntitlementsRevokedOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementsRevoked) {
		return nil, false
	}
	return o.EntitlementsRevoked, true
}

// HasEntitlementsRevoked returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasEntitlementsRevoked() bool {
	if o != nil && !isNil(o.EntitlementsRevoked) {
		return true
	}

	return false
}

// SetEntitlementsRevoked gets a reference to the given int32 and assigns it to the EntitlementsRevoked field.
func (o *IdentityCertDecisionSummary) SetEntitlementsRevoked(v int32) {
	o.EntitlementsRevoked = &v
}

// GetAccessProfilesApproved returns the AccessProfilesApproved field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccessProfilesApproved() int32 {
	if o == nil || isNil(o.AccessProfilesApproved) {
		var ret int32
		return ret
	}
	return *o.AccessProfilesApproved
}

// GetAccessProfilesApprovedOk returns a tuple with the AccessProfilesApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccessProfilesApprovedOk() (*int32, bool) {
	if o == nil || isNil(o.AccessProfilesApproved) {
		return nil, false
	}
	return o.AccessProfilesApproved, true
}

// HasAccessProfilesApproved returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccessProfilesApproved() bool {
	if o != nil && !isNil(o.AccessProfilesApproved) {
		return true
	}

	return false
}

// SetAccessProfilesApproved gets a reference to the given int32 and assigns it to the AccessProfilesApproved field.
func (o *IdentityCertDecisionSummary) SetAccessProfilesApproved(v int32) {
	o.AccessProfilesApproved = &v
}

// GetAccessProfilesRevoked returns the AccessProfilesRevoked field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccessProfilesRevoked() int32 {
	if o == nil || isNil(o.AccessProfilesRevoked) {
		var ret int32
		return ret
	}
	return *o.AccessProfilesRevoked
}

// GetAccessProfilesRevokedOk returns a tuple with the AccessProfilesRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccessProfilesRevokedOk() (*int32, bool) {
	if o == nil || isNil(o.AccessProfilesRevoked) {
		return nil, false
	}
	return o.AccessProfilesRevoked, true
}

// HasAccessProfilesRevoked returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccessProfilesRevoked() bool {
	if o != nil && !isNil(o.AccessProfilesRevoked) {
		return true
	}

	return false
}

// SetAccessProfilesRevoked gets a reference to the given int32 and assigns it to the AccessProfilesRevoked field.
func (o *IdentityCertDecisionSummary) SetAccessProfilesRevoked(v int32) {
	o.AccessProfilesRevoked = &v
}

// GetRolesApproved returns the RolesApproved field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetRolesApproved() int32 {
	if o == nil || isNil(o.RolesApproved) {
		var ret int32
		return ret
	}
	return *o.RolesApproved
}

// GetRolesApprovedOk returns a tuple with the RolesApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetRolesApprovedOk() (*int32, bool) {
	if o == nil || isNil(o.RolesApproved) {
		return nil, false
	}
	return o.RolesApproved, true
}

// HasRolesApproved returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasRolesApproved() bool {
	if o != nil && !isNil(o.RolesApproved) {
		return true
	}

	return false
}

// SetRolesApproved gets a reference to the given int32 and assigns it to the RolesApproved field.
func (o *IdentityCertDecisionSummary) SetRolesApproved(v int32) {
	o.RolesApproved = &v
}

// GetRolesRevoked returns the RolesRevoked field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetRolesRevoked() int32 {
	if o == nil || isNil(o.RolesRevoked) {
		var ret int32
		return ret
	}
	return *o.RolesRevoked
}

// GetRolesRevokedOk returns a tuple with the RolesRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetRolesRevokedOk() (*int32, bool) {
	if o == nil || isNil(o.RolesRevoked) {
		return nil, false
	}
	return o.RolesRevoked, true
}

// HasRolesRevoked returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasRolesRevoked() bool {
	if o != nil && !isNil(o.RolesRevoked) {
		return true
	}

	return false
}

// SetRolesRevoked gets a reference to the given int32 and assigns it to the RolesRevoked field.
func (o *IdentityCertDecisionSummary) SetRolesRevoked(v int32) {
	o.RolesRevoked = &v
}

// GetAccountsApproved returns the AccountsApproved field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccountsApproved() int32 {
	if o == nil || isNil(o.AccountsApproved) {
		var ret int32
		return ret
	}
	return *o.AccountsApproved
}

// GetAccountsApprovedOk returns a tuple with the AccountsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccountsApprovedOk() (*int32, bool) {
	if o == nil || isNil(o.AccountsApproved) {
		return nil, false
	}
	return o.AccountsApproved, true
}

// HasAccountsApproved returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccountsApproved() bool {
	if o != nil && !isNil(o.AccountsApproved) {
		return true
	}

	return false
}

// SetAccountsApproved gets a reference to the given int32 and assigns it to the AccountsApproved field.
func (o *IdentityCertDecisionSummary) SetAccountsApproved(v int32) {
	o.AccountsApproved = &v
}

// GetAccountsRevoked returns the AccountsRevoked field value if set, zero value otherwise.
func (o *IdentityCertDecisionSummary) GetAccountsRevoked() int32 {
	if o == nil || isNil(o.AccountsRevoked) {
		var ret int32
		return ret
	}
	return *o.AccountsRevoked
}

// GetAccountsRevokedOk returns a tuple with the AccountsRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertDecisionSummary) GetAccountsRevokedOk() (*int32, bool) {
	if o == nil || isNil(o.AccountsRevoked) {
		return nil, false
	}
	return o.AccountsRevoked, true
}

// HasAccountsRevoked returns a boolean if a field has been set.
func (o *IdentityCertDecisionSummary) HasAccountsRevoked() bool {
	if o != nil && !isNil(o.AccountsRevoked) {
		return true
	}

	return false
}

// SetAccountsRevoked gets a reference to the given int32 and assigns it to the AccountsRevoked field.
func (o *IdentityCertDecisionSummary) SetAccountsRevoked(v int32) {
	o.AccountsRevoked = &v
}

func (o IdentityCertDecisionSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityCertDecisionSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EntitlementDecisionsMade) {
		toSerialize["entitlementDecisionsMade"] = o.EntitlementDecisionsMade
	}
	if !isNil(o.AccessProfileDecisionsMade) {
		toSerialize["accessProfileDecisionsMade"] = o.AccessProfileDecisionsMade
	}
	if !isNil(o.RoleDecisionsMade) {
		toSerialize["roleDecisionsMade"] = o.RoleDecisionsMade
	}
	if !isNil(o.AccountDecisionsMade) {
		toSerialize["accountDecisionsMade"] = o.AccountDecisionsMade
	}
	if !isNil(o.EntitlementDecisionsTotal) {
		toSerialize["entitlementDecisionsTotal"] = o.EntitlementDecisionsTotal
	}
	if !isNil(o.AccessProfileDecisionsTotal) {
		toSerialize["accessProfileDecisionsTotal"] = o.AccessProfileDecisionsTotal
	}
	if !isNil(o.RoleDecisionsTotal) {
		toSerialize["roleDecisionsTotal"] = o.RoleDecisionsTotal
	}
	if !isNil(o.AccountDecisionsTotal) {
		toSerialize["accountDecisionsTotal"] = o.AccountDecisionsTotal
	}
	if !isNil(o.EntitlementsApproved) {
		toSerialize["entitlementsApproved"] = o.EntitlementsApproved
	}
	if !isNil(o.EntitlementsRevoked) {
		toSerialize["entitlementsRevoked"] = o.EntitlementsRevoked
	}
	if !isNil(o.AccessProfilesApproved) {
		toSerialize["accessProfilesApproved"] = o.AccessProfilesApproved
	}
	if !isNil(o.AccessProfilesRevoked) {
		toSerialize["accessProfilesRevoked"] = o.AccessProfilesRevoked
	}
	if !isNil(o.RolesApproved) {
		toSerialize["rolesApproved"] = o.RolesApproved
	}
	if !isNil(o.RolesRevoked) {
		toSerialize["rolesRevoked"] = o.RolesRevoked
	}
	if !isNil(o.AccountsApproved) {
		toSerialize["accountsApproved"] = o.AccountsApproved
	}
	if !isNil(o.AccountsRevoked) {
		toSerialize["accountsRevoked"] = o.AccountsRevoked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityCertDecisionSummary) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityCertDecisionSummary := _IdentityCertDecisionSummary{}

	if err = json.Unmarshal(bytes, &varIdentityCertDecisionSummary); err == nil {
		*o = IdentityCertDecisionSummary(varIdentityCertDecisionSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementDecisionsMade")
		delete(additionalProperties, "accessProfileDecisionsMade")
		delete(additionalProperties, "roleDecisionsMade")
		delete(additionalProperties, "accountDecisionsMade")
		delete(additionalProperties, "entitlementDecisionsTotal")
		delete(additionalProperties, "accessProfileDecisionsTotal")
		delete(additionalProperties, "roleDecisionsTotal")
		delete(additionalProperties, "accountDecisionsTotal")
		delete(additionalProperties, "entitlementsApproved")
		delete(additionalProperties, "entitlementsRevoked")
		delete(additionalProperties, "accessProfilesApproved")
		delete(additionalProperties, "accessProfilesRevoked")
		delete(additionalProperties, "rolesApproved")
		delete(additionalProperties, "rolesRevoked")
		delete(additionalProperties, "accountsApproved")
		delete(additionalProperties, "accountsRevoked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityCertDecisionSummary struct {
	value *IdentityCertDecisionSummary
	isSet bool
}

func (v NullableIdentityCertDecisionSummary) Get() *IdentityCertDecisionSummary {
	return v.value
}

func (v *NullableIdentityCertDecisionSummary) Set(val *IdentityCertDecisionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCertDecisionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCertDecisionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCertDecisionSummary(val *IdentityCertDecisionSummary) *NullableIdentityCertDecisionSummary {
	return &NullableIdentityCertDecisionSummary{value: val, isSet: true}
}

func (v NullableIdentityCertDecisionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCertDecisionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


