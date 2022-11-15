/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointccsdk

import (
	"encoding/json"
)

// CreateAccountSchemaAttributeRequest struct for CreateAccountSchemaAttributeRequest
type CreateAccountSchemaAttributeRequest struct {
	ObjectType *string `json:"objectType,omitempty"`
	Entitlement *bool `json:"entitlement,omitempty"`
	Multi *bool `json:"multi,omitempty"`
	Names *string `json:"names,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccountSchemaAttributeRequest CreateAccountSchemaAttributeRequest

// NewCreateAccountSchemaAttributeRequest instantiates a new CreateAccountSchemaAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSchemaAttributeRequest() *CreateAccountSchemaAttributeRequest {
	this := CreateAccountSchemaAttributeRequest{}
	return &this
}

// NewCreateAccountSchemaAttributeRequestWithDefaults instantiates a new CreateAccountSchemaAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSchemaAttributeRequestWithDefaults() *CreateAccountSchemaAttributeRequest {
	this := CreateAccountSchemaAttributeRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *CreateAccountSchemaAttributeRequest) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSchemaAttributeRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *CreateAccountSchemaAttributeRequest) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *CreateAccountSchemaAttributeRequest) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetEntitlement returns the Entitlement field value if set, zero value otherwise.
func (o *CreateAccountSchemaAttributeRequest) GetEntitlement() bool {
	if o == nil || isNil(o.Entitlement) {
		var ret bool
		return ret
	}
	return *o.Entitlement
}

// GetEntitlementOk returns a tuple with the Entitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSchemaAttributeRequest) GetEntitlementOk() (*bool, bool) {
	if o == nil || isNil(o.Entitlement) {
		return nil, false
	}
	return o.Entitlement, true
}

// HasEntitlement returns a boolean if a field has been set.
func (o *CreateAccountSchemaAttributeRequest) HasEntitlement() bool {
	if o != nil && !isNil(o.Entitlement) {
		return true
	}

	return false
}

// SetEntitlement gets a reference to the given bool and assigns it to the Entitlement field.
func (o *CreateAccountSchemaAttributeRequest) SetEntitlement(v bool) {
	o.Entitlement = &v
}

// GetMulti returns the Multi field value if set, zero value otherwise.
func (o *CreateAccountSchemaAttributeRequest) GetMulti() bool {
	if o == nil || isNil(o.Multi) {
		var ret bool
		return ret
	}
	return *o.Multi
}

// GetMultiOk returns a tuple with the Multi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSchemaAttributeRequest) GetMultiOk() (*bool, bool) {
	if o == nil || isNil(o.Multi) {
		return nil, false
	}
	return o.Multi, true
}

// HasMulti returns a boolean if a field has been set.
func (o *CreateAccountSchemaAttributeRequest) HasMulti() bool {
	if o != nil && !isNil(o.Multi) {
		return true
	}

	return false
}

// SetMulti gets a reference to the given bool and assigns it to the Multi field.
func (o *CreateAccountSchemaAttributeRequest) SetMulti(v bool) {
	o.Multi = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *CreateAccountSchemaAttributeRequest) GetNames() string {
	if o == nil || isNil(o.Names) {
		var ret string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSchemaAttributeRequest) GetNamesOk() (*string, bool) {
	if o == nil || isNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *CreateAccountSchemaAttributeRequest) HasNames() bool {
	if o != nil && !isNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given string and assigns it to the Names field.
func (o *CreateAccountSchemaAttributeRequest) SetNames(v string) {
	o.Names = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateAccountSchemaAttributeRequest) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSchemaAttributeRequest) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateAccountSchemaAttributeRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateAccountSchemaAttributeRequest) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAccountSchemaAttributeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSchemaAttributeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAccountSchemaAttributeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAccountSchemaAttributeRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreateAccountSchemaAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.Entitlement) {
		toSerialize["entitlement"] = o.Entitlement
	}
	if !isNil(o.Multi) {
		toSerialize["multi"] = o.Multi
	}
	if !isNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateAccountSchemaAttributeRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateAccountSchemaAttributeRequest := _CreateAccountSchemaAttributeRequest{}

	if err = json.Unmarshal(bytes, &varCreateAccountSchemaAttributeRequest); err == nil {
		*o = CreateAccountSchemaAttributeRequest(varCreateAccountSchemaAttributeRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "entitlement")
		delete(additionalProperties, "multi")
		delete(additionalProperties, "names")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccountSchemaAttributeRequest struct {
	value *CreateAccountSchemaAttributeRequest
	isSet bool
}

func (v NullableCreateAccountSchemaAttributeRequest) Get() *CreateAccountSchemaAttributeRequest {
	return v.value
}

func (v *NullableCreateAccountSchemaAttributeRequest) Set(val *CreateAccountSchemaAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSchemaAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSchemaAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSchemaAttributeRequest(val *CreateAccountSchemaAttributeRequest) *NullableCreateAccountSchemaAttributeRequest {
	return &NullableCreateAccountSchemaAttributeRequest{value: val, isSet: true}
}

func (v NullableCreateAccountSchemaAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSchemaAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


