/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// ApprovalInfoResponse struct for ApprovalInfoResponse
type ApprovalInfoResponse struct {
	// the id of approver
	Id *string `json:"id,omitempty"`
	// the name of approver
	Name *string `json:"name,omitempty"`
	// the status of the approval request
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalInfoResponse ApprovalInfoResponse

// NewApprovalInfoResponse instantiates a new ApprovalInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalInfoResponse() *ApprovalInfoResponse {
	this := ApprovalInfoResponse{}
	return &this
}

// NewApprovalInfoResponseWithDefaults instantiates a new ApprovalInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalInfoResponseWithDefaults() *ApprovalInfoResponse {
	this := ApprovalInfoResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApprovalInfoResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalInfoResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApprovalInfoResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApprovalInfoResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApprovalInfoResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalInfoResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApprovalInfoResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApprovalInfoResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApprovalInfoResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalInfoResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApprovalInfoResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApprovalInfoResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ApprovalInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApprovalInfoResponse) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalInfoResponse := _ApprovalInfoResponse{}

	if err = json.Unmarshal(bytes, &varApprovalInfoResponse); err == nil {
		*o = ApprovalInfoResponse(varApprovalInfoResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalInfoResponse struct {
	value *ApprovalInfoResponse
	isSet bool
}

func (v NullableApprovalInfoResponse) Get() *ApprovalInfoResponse {
	return v.value
}

func (v *NullableApprovalInfoResponse) Set(val *ApprovalInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalInfoResponse(val *ApprovalInfoResponse) *NullableApprovalInfoResponse {
	return &NullableApprovalInfoResponse{value: val, isSet: true}
}

func (v NullableApprovalInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


