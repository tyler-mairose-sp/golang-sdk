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

// AccessProfileSummary This is a summary representation of an access profile.
type AccessProfileSummary struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Source *Reference1 `json:"source,omitempty"`
	Owner *DisplayReference `json:"owner,omitempty"`
	Revocable *bool `json:"revocable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileSummary AccessProfileSummary

// NewAccessProfileSummary instantiates a new AccessProfileSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileSummary() *AccessProfileSummary {
	this := AccessProfileSummary{}
	return &this
}

// NewAccessProfileSummaryWithDefaults instantiates a new AccessProfileSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileSummaryWithDefaults() *AccessProfileSummary {
	this := AccessProfileSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessProfileSummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessProfileSummary) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AccessProfileSummary) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *AccessProfileSummary) SetType(v DtoType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileSummary) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileSummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessProfileSummary) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessProfileSummary) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessProfileSummary) UnsetDescription() {
	o.Description.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetSource() Reference1 {
	if o == nil || isNil(o.Source) {
		var ret Reference1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetSourceOk() (*Reference1, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference1 and assigns it to the Source field.
func (o *AccessProfileSummary) SetSource(v Reference1) {
	o.Source = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetOwner() DisplayReference {
	if o == nil || isNil(o.Owner) {
		var ret DisplayReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetOwnerOk() (*DisplayReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DisplayReference and assigns it to the Owner field.
func (o *AccessProfileSummary) SetOwner(v DisplayReference) {
	o.Owner = &v
}

// GetRevocable returns the Revocable field value if set, zero value otherwise.
func (o *AccessProfileSummary) GetRevocable() bool {
	if o == nil || isNil(o.Revocable) {
		var ret bool
		return ret
	}
	return *o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileSummary) GetRevocableOk() (*bool, bool) {
	if o == nil || isNil(o.Revocable) {
		return nil, false
	}
	return o.Revocable, true
}

// HasRevocable returns a boolean if a field has been set.
func (o *AccessProfileSummary) HasRevocable() bool {
	if o != nil && !isNil(o.Revocable) {
		return true
	}

	return false
}

// SetRevocable gets a reference to the given bool and assigns it to the Revocable field.
func (o *AccessProfileSummary) SetRevocable(v bool) {
	o.Revocable = &v
}

func (o AccessProfileSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Revocable) {
		toSerialize["revocable"] = o.Revocable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessProfileSummary) UnmarshalJSON(bytes []byte) (err error) {
	varAccessProfileSummary := _AccessProfileSummary{}

	if err = json.Unmarshal(bytes, &varAccessProfileSummary); err == nil {
		*o = AccessProfileSummary(varAccessProfileSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "source")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "revocable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileSummary struct {
	value *AccessProfileSummary
	isSet bool
}

func (v NullableAccessProfileSummary) Get() *AccessProfileSummary {
	return v.value
}

func (v *NullableAccessProfileSummary) Set(val *AccessProfileSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileSummary(val *AccessProfileSummary) *NullableAccessProfileSummary {
	return &NullableAccessProfileSummary{value: val, isSet: true}
}

func (v NullableAccessProfileSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


