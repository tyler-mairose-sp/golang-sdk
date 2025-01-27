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

// checks if the TaggedObjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaggedObjectDto{}

// TaggedObjectDto struct for TaggedObjectDto
type TaggedObjectDto struct {
	Type *DtoType `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaggedObjectDto TaggedObjectDto

// NewTaggedObjectDto instantiates a new TaggedObjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaggedObjectDto() *TaggedObjectDto {
	this := TaggedObjectDto{}
	return &this
}

// NewTaggedObjectDtoWithDefaults instantiates a new TaggedObjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggedObjectDtoWithDefaults() *TaggedObjectDto {
	this := TaggedObjectDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaggedObjectDto) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggedObjectDto) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaggedObjectDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *TaggedObjectDto) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaggedObjectDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggedObjectDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaggedObjectDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaggedObjectDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaggedObjectDto) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaggedObjectDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TaggedObjectDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TaggedObjectDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TaggedObjectDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TaggedObjectDto) UnsetName() {
	o.Name.Unset()
}

func (o TaggedObjectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaggedObjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaggedObjectDto) UnmarshalJSON(bytes []byte) (err error) {
	varTaggedObjectDto := _TaggedObjectDto{}

	if err = json.Unmarshal(bytes, &varTaggedObjectDto); err == nil {
		*o = TaggedObjectDto(varTaggedObjectDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaggedObjectDto struct {
	value *TaggedObjectDto
	isSet bool
}

func (v NullableTaggedObjectDto) Get() *TaggedObjectDto {
	return v.value
}

func (v *NullableTaggedObjectDto) Set(val *TaggedObjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTaggedObjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTaggedObjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaggedObjectDto(val *TaggedObjectDto) *NullableTaggedObjectDto {
	return &NullableTaggedObjectDto{value: val, isSet: true}
}

func (v NullableTaggedObjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaggedObjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


