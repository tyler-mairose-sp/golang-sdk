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

// checks if the EntitlementSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementSummary{}

// EntitlementSummary EntitlementReference
type EntitlementSummary struct {
	// The unique ID of the referenced object.
	Id *string `json:"id,omitempty"`
	// The human readable name of the referenced object.
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Source *Reference `json:"source,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	Attribute *string `json:"attribute,omitempty"`
	Value *string `json:"value,omitempty"`
	Standalone *bool `json:"standalone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementSummary EntitlementSummary

// NewEntitlementSummary instantiates a new EntitlementSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementSummary() *EntitlementSummary {
	this := EntitlementSummary{}
	return &this
}

// NewEntitlementSummaryWithDefaults instantiates a new EntitlementSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementSummaryWithDefaults() *EntitlementSummary {
	this := EntitlementSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementSummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementSummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementSummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementSummary) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EntitlementSummary) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EntitlementSummary) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EntitlementSummary) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitlementSummary) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitlementSummary) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *EntitlementSummary) SetType(v DtoType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementSummary) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementSummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementSummary) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EntitlementSummary) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EntitlementSummary) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EntitlementSummary) UnsetDescription() {
	o.Description.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntitlementSummary) GetSource() Reference {
	if o == nil || isNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetSourceOk() (*Reference, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntitlementSummary) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *EntitlementSummary) SetSource(v Reference) {
	o.Source = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *EntitlementSummary) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *EntitlementSummary) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *EntitlementSummary) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *EntitlementSummary) GetAttribute() string {
	if o == nil || isNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *EntitlementSummary) HasAttribute() bool {
	if o != nil && !isNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *EntitlementSummary) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementSummary) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementSummary) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementSummary) SetValue(v string) {
	o.Value = &v
}

// GetStandalone returns the Standalone field value if set, zero value otherwise.
func (o *EntitlementSummary) GetStandalone() bool {
	if o == nil || isNil(o.Standalone) {
		var ret bool
		return ret
	}
	return *o.Standalone
}

// GetStandaloneOk returns a tuple with the Standalone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementSummary) GetStandaloneOk() (*bool, bool) {
	if o == nil || isNil(o.Standalone) {
		return nil, false
	}
	return o.Standalone, true
}

// HasStandalone returns a boolean if a field has been set.
func (o *EntitlementSummary) HasStandalone() bool {
	if o != nil && !isNil(o.Standalone) {
		return true
	}

	return false
}

// SetStandalone gets a reference to the given bool and assigns it to the Standalone field.
func (o *EntitlementSummary) SetStandalone(v bool) {
	o.Standalone = &v
}

func (o EntitlementSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementSummary) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Standalone) {
		toSerialize["standalone"] = o.Standalone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementSummary) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementSummary := _EntitlementSummary{}

	if err = json.Unmarshal(bytes, &varEntitlementSummary); err == nil {
		*o = EntitlementSummary(varEntitlementSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "source")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "standalone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementSummary struct {
	value *EntitlementSummary
	isSet bool
}

func (v NullableEntitlementSummary) Get() *EntitlementSummary {
	return v.value
}

func (v *NullableEntitlementSummary) Set(val *EntitlementSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementSummary(val *EntitlementSummary) *NullableEntitlementSummary {
	return &NullableEntitlementSummary{value: val, isSet: true}
}

func (v NullableEntitlementSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


