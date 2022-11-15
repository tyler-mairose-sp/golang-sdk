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

// ForwardApprovalDto struct for ForwardApprovalDto
type ForwardApprovalDto struct {
	// The Id of the new owner
	NewOwnerId string `json:"newOwnerId"`
	// The comment provided by the forwarder
	Comment string `json:"comment"`
	AdditionalProperties map[string]interface{}
}

type _ForwardApprovalDto ForwardApprovalDto

// NewForwardApprovalDto instantiates a new ForwardApprovalDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardApprovalDto(newOwnerId string, comment string) *ForwardApprovalDto {
	this := ForwardApprovalDto{}
	this.NewOwnerId = newOwnerId
	this.Comment = comment
	return &this
}

// NewForwardApprovalDtoWithDefaults instantiates a new ForwardApprovalDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardApprovalDtoWithDefaults() *ForwardApprovalDto {
	this := ForwardApprovalDto{}
	return &this
}

// GetNewOwnerId returns the NewOwnerId field value
func (o *ForwardApprovalDto) GetNewOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewOwnerId
}

// GetNewOwnerIdOk returns a tuple with the NewOwnerId field value
// and a boolean to check if the value has been set.
func (o *ForwardApprovalDto) GetNewOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewOwnerId, true
}

// SetNewOwnerId sets field value
func (o *ForwardApprovalDto) SetNewOwnerId(v string) {
	o.NewOwnerId = v
}

// GetComment returns the Comment field value
func (o *ForwardApprovalDto) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *ForwardApprovalDto) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *ForwardApprovalDto) SetComment(v string) {
	o.Comment = v
}

func (o ForwardApprovalDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["newOwnerId"] = o.NewOwnerId
	}
	if true {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForwardApprovalDto) UnmarshalJSON(bytes []byte) (err error) {
	varForwardApprovalDto := _ForwardApprovalDto{}

	if err = json.Unmarshal(bytes, &varForwardApprovalDto); err == nil {
		*o = ForwardApprovalDto(varForwardApprovalDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "newOwnerId")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForwardApprovalDto struct {
	value *ForwardApprovalDto
	isSet bool
}

func (v NullableForwardApprovalDto) Get() *ForwardApprovalDto {
	return v.value
}

func (v *NullableForwardApprovalDto) Set(val *ForwardApprovalDto) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardApprovalDto) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardApprovalDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardApprovalDto(val *ForwardApprovalDto) *NullableForwardApprovalDto {
	return &NullableForwardApprovalDto{value: val, isSet: true}
}

func (v NullableForwardApprovalDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardApprovalDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


