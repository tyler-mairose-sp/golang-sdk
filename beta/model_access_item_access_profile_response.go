/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the AccessItemAccessProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemAccessProfileResponse{}

// AccessItemAccessProfileResponse struct for AccessItemAccessProfileResponse
type AccessItemAccessProfileResponse struct {
	// the access item type. accessProfile in this case
	AccessType *string `json:"accessType,omitempty"`
	// the access item id
	Id *string `json:"id,omitempty"`
	// the access profile name
	Name *string `json:"name,omitempty"`
	// the name of the source
	SourceName *string `json:"sourceName,omitempty"`
	// the id of the source
	SourceId *string `json:"sourceId,omitempty"`
	// the description for the access profile
	Description *string `json:"description,omitempty"`
	// the display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// the number of entitlements the access profile will create
	EntitlementCount *string `json:"entitlementCount,omitempty"`
	// the name of app
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemAccessProfileResponse AccessItemAccessProfileResponse

// NewAccessItemAccessProfileResponse instantiates a new AccessItemAccessProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemAccessProfileResponse() *AccessItemAccessProfileResponse {
	this := AccessItemAccessProfileResponse{}
	return &this
}

// NewAccessItemAccessProfileResponseWithDefaults instantiates a new AccessItemAccessProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemAccessProfileResponseWithDefaults() *AccessItemAccessProfileResponse {
	this := AccessItemAccessProfileResponse{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetAccessType() string {
	if o == nil || isNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetAccessTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *AccessItemAccessProfileResponse) SetAccessType(v string) {
	o.AccessType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemAccessProfileResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessItemAccessProfileResponse) SetName(v string) {
	o.Name = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessItemAccessProfileResponse) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AccessItemAccessProfileResponse) SetSourceId(v string) {
	o.SourceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessItemAccessProfileResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AccessItemAccessProfileResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetEntitlementCount() string {
	if o == nil || isNil(o.EntitlementCount) {
		var ret string
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetEntitlementCountOk() (*string, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given string and assigns it to the EntitlementCount field.
func (o *AccessItemAccessProfileResponse) SetEntitlementCount(v string) {
	o.EntitlementCount = &v
}

// GetAppDisplayName returns the AppDisplayName field value if set, zero value otherwise.
func (o *AccessItemAccessProfileResponse) GetAppDisplayName() string {
	if o == nil || isNil(o.AppDisplayName) {
		var ret string
		return ret
	}
	return *o.AppDisplayName
}

// GetAppDisplayNameOk returns a tuple with the AppDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccessProfileResponse) GetAppDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.AppDisplayName) {
		return nil, false
	}
	return o.AppDisplayName, true
}

// HasAppDisplayName returns a boolean if a field has been set.
func (o *AccessItemAccessProfileResponse) HasAppDisplayName() bool {
	if o != nil && !isNil(o.AppDisplayName) {
		return true
	}

	return false
}

// SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.
func (o *AccessItemAccessProfileResponse) SetAppDisplayName(v string) {
	o.AppDisplayName = &v
}

func (o AccessItemAccessProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemAccessProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.AppDisplayName) {
		toSerialize["appDisplayName"] = o.AppDisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemAccessProfileResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccessItemAccessProfileResponse := _AccessItemAccessProfileResponse{}

	if err = json.Unmarshal(bytes, &varAccessItemAccessProfileResponse); err == nil {
		*o = AccessItemAccessProfileResponse(varAccessItemAccessProfileResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "appDisplayName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemAccessProfileResponse struct {
	value *AccessItemAccessProfileResponse
	isSet bool
}

func (v NullableAccessItemAccessProfileResponse) Get() *AccessItemAccessProfileResponse {
	return v.value
}

func (v *NullableAccessItemAccessProfileResponse) Set(val *AccessItemAccessProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemAccessProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemAccessProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemAccessProfileResponse(val *AccessItemAccessProfileResponse) *NullableAccessItemAccessProfileResponse {
	return &NullableAccessItemAccessProfileResponse{value: val, isSet: true}
}

func (v NullableAccessItemAccessProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemAccessProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


