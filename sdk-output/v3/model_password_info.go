/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// PasswordInfo struct for PasswordInfo
type PasswordInfo struct {
	// Identity ID
	IdentityId *string `json:"identityId,omitempty"`
	// source ID
	SourceId *string `json:"sourceId,omitempty"`
	// public key ID
	PublicKeyId *string `json:"publicKeyId,omitempty"`
	// User's public key with Base64 encoding
	PublicKey *string `json:"publicKey,omitempty"`
	// Account info related to queried identity and source
	Accounts []PasswordInfoAccount `json:"accounts,omitempty"`
	// Password constraints
	Policies []string `json:"policies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordInfo PasswordInfo

// NewPasswordInfo instantiates a new PasswordInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordInfo() *PasswordInfo {
	this := PasswordInfo{}
	return &this
}

// NewPasswordInfoWithDefaults instantiates a new PasswordInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordInfoWithDefaults() *PasswordInfo {
	this := PasswordInfo{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *PasswordInfo) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfo) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *PasswordInfo) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *PasswordInfo) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *PasswordInfo) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfo) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *PasswordInfo) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *PasswordInfo) SetSourceId(v string) {
	o.SourceId = &v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise.
func (o *PasswordInfo) GetPublicKeyId() string {
	if o == nil || isNil(o.PublicKeyId) {
		var ret string
		return ret
	}
	return *o.PublicKeyId
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfo) GetPublicKeyIdOk() (*string, bool) {
	if o == nil || isNil(o.PublicKeyId) {
		return nil, false
	}
	return o.PublicKeyId, true
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *PasswordInfo) HasPublicKeyId() bool {
	if o != nil && !isNil(o.PublicKeyId) {
		return true
	}

	return false
}

// SetPublicKeyId gets a reference to the given string and assigns it to the PublicKeyId field.
func (o *PasswordInfo) SetPublicKeyId(v string) {
	o.PublicKeyId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PasswordInfo) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfo) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PasswordInfo) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PasswordInfo) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *PasswordInfo) GetAccounts() []PasswordInfoAccount {
	if o == nil || isNil(o.Accounts) {
		var ret []PasswordInfoAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfo) GetAccountsOk() ([]PasswordInfoAccount, bool) {
	if o == nil || isNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *PasswordInfo) HasAccounts() bool {
	if o != nil && !isNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []PasswordInfoAccount and assigns it to the Accounts field.
func (o *PasswordInfo) SetAccounts(v []PasswordInfoAccount) {
	o.Accounts = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *PasswordInfo) GetPolicies() []string {
	if o == nil || isNil(o.Policies) {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordInfo) GetPoliciesOk() ([]string, bool) {
	if o == nil || isNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *PasswordInfo) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *PasswordInfo) SetPolicies(v []string) {
	o.Policies = v
}

func (o PasswordInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.PublicKeyId) {
		toSerialize["publicKeyId"] = o.PublicKeyId
	}
	if !isNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !isNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordInfo) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordInfo := _PasswordInfo{}

	if err = json.Unmarshal(bytes, &varPasswordInfo); err == nil {
		*o = PasswordInfo(varPasswordInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "publicKeyId")
		delete(additionalProperties, "publicKey")
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "policies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordInfo struct {
	value *PasswordInfo
	isSet bool
}

func (v NullablePasswordInfo) Get() *PasswordInfo {
	return v.value
}

func (v *NullablePasswordInfo) Set(val *PasswordInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordInfo(val *PasswordInfo) *NullablePasswordInfo {
	return &NullablePasswordInfo{value: val, isSet: true}
}

func (v NullablePasswordInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


