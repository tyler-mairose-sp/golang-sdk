/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the ListApplications200ResponseInnerAppProfilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplications200ResponseInnerAppProfilesInner{}

// ListApplications200ResponseInnerAppProfilesInner struct for ListApplications200ResponseInnerAppProfilesInner
type ListApplications200ResponseInnerAppProfilesInner struct {
	Id *float32 `json:"id,omitempty"`
	Filename *string `json:"filename,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	DateCreated *string `json:"dateCreated,omitempty"`
	XsdVersion *string `json:"xsdVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListApplications200ResponseInnerAppProfilesInner ListApplications200ResponseInnerAppProfilesInner

// NewListApplications200ResponseInnerAppProfilesInner instantiates a new ListApplications200ResponseInnerAppProfilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplications200ResponseInnerAppProfilesInner() *ListApplications200ResponseInnerAppProfilesInner {
	this := ListApplications200ResponseInnerAppProfilesInner{}
	return &this
}

// NewListApplications200ResponseInnerAppProfilesInnerWithDefaults instantiates a new ListApplications200ResponseInnerAppProfilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplications200ResponseInnerAppProfilesInnerWithDefaults() *ListApplications200ResponseInnerAppProfilesInner {
	this := ListApplications200ResponseInnerAppProfilesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ListApplications200ResponseInnerAppProfilesInner) SetId(v float32) {
	o.Id = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetFilename() string {
	if o == nil || isNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetFilenameOk() (*string, bool) {
	if o == nil || isNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) HasFilename() bool {
	if o != nil && !isNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ListApplications200ResponseInnerAppProfilesInner) SetFilename(v string) {
	o.Filename = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetCreatedByOk() (*string, bool) {
	if o == nil || isNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ListApplications200ResponseInnerAppProfilesInner) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetDateCreated() string {
	if o == nil || isNil(o.DateCreated) {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetDateCreatedOk() (*string, bool) {
	if o == nil || isNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) HasDateCreated() bool {
	if o != nil && !isNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *ListApplications200ResponseInnerAppProfilesInner) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetXsdVersion returns the XsdVersion field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetXsdVersion() string {
	if o == nil || isNil(o.XsdVersion) {
		var ret string
		return ret
	}
	return *o.XsdVersion
}

// GetXsdVersionOk returns a tuple with the XsdVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) GetXsdVersionOk() (*string, bool) {
	if o == nil || isNil(o.XsdVersion) {
		return nil, false
	}
	return o.XsdVersion, true
}

// HasXsdVersion returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerAppProfilesInner) HasXsdVersion() bool {
	if o != nil && !isNil(o.XsdVersion) {
		return true
	}

	return false
}

// SetXsdVersion gets a reference to the given string and assigns it to the XsdVersion field.
func (o *ListApplications200ResponseInnerAppProfilesInner) SetXsdVersion(v string) {
	o.XsdVersion = &v
}

func (o ListApplications200ResponseInnerAppProfilesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplications200ResponseInnerAppProfilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !isNil(o.XsdVersion) {
		toSerialize["xsdVersion"] = o.XsdVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListApplications200ResponseInnerAppProfilesInner) UnmarshalJSON(bytes []byte) (err error) {
	varListApplications200ResponseInnerAppProfilesInner := _ListApplications200ResponseInnerAppProfilesInner{}

	if err = json.Unmarshal(bytes, &varListApplications200ResponseInnerAppProfilesInner); err == nil {
		*o = ListApplications200ResponseInnerAppProfilesInner(varListApplications200ResponseInnerAppProfilesInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "filename")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "dateCreated")
		delete(additionalProperties, "xsdVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListApplications200ResponseInnerAppProfilesInner struct {
	value *ListApplications200ResponseInnerAppProfilesInner
	isSet bool
}

func (v NullableListApplications200ResponseInnerAppProfilesInner) Get() *ListApplications200ResponseInnerAppProfilesInner {
	return v.value
}

func (v *NullableListApplications200ResponseInnerAppProfilesInner) Set(val *ListApplications200ResponseInnerAppProfilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInnerAppProfilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInnerAppProfilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInnerAppProfilesInner(val *ListApplications200ResponseInnerAppProfilesInner) *NullableListApplications200ResponseInnerAppProfilesInner {
	return &NullableListApplications200ResponseInnerAppProfilesInner{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInnerAppProfilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInnerAppProfilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

