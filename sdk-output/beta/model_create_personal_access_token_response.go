/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
	"time"
)

// CreatePersonalAccessTokenResponse struct for CreatePersonalAccessTokenResponse
type CreatePersonalAccessTokenResponse struct {
	// The ID of the personal access token (to be used as the username for Basic Auth).
	Id string `json:"id"`
	// The secret of the personal access token (to be used as the password for Basic Auth).
	Secret string `json:"secret"`
	// Scopes of the personal  access token.
	Scope []string `json:"scope"`
	// The name of the personal access token. Cannot be the same as other personal access tokens owned by a user.
	Name string `json:"name"`
	Owner BaseReferenceDto `json:"owner"`
	// The date and time, down to the millisecond, when this personal access token was created.
	Created time.Time `json:"created"`
	AdditionalProperties map[string]interface{}
}

type _CreatePersonalAccessTokenResponse CreatePersonalAccessTokenResponse

// NewCreatePersonalAccessTokenResponse instantiates a new CreatePersonalAccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePersonalAccessTokenResponse(id string, secret string, scope []string, name string, owner BaseReferenceDto, created time.Time) *CreatePersonalAccessTokenResponse {
	this := CreatePersonalAccessTokenResponse{}
	this.Id = id
	this.Secret = secret
	this.Scope = scope
	this.Name = name
	this.Owner = owner
	this.Created = created
	return &this
}

// NewCreatePersonalAccessTokenResponseWithDefaults instantiates a new CreatePersonalAccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePersonalAccessTokenResponseWithDefaults() *CreatePersonalAccessTokenResponse {
	this := CreatePersonalAccessTokenResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreatePersonalAccessTokenResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreatePersonalAccessTokenResponse) SetId(v string) {
	o.Id = v
}

// GetSecret returns the Secret field value
func (o *CreatePersonalAccessTokenResponse) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenResponse) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreatePersonalAccessTokenResponse) SetSecret(v string) {
	o.Secret = v
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CreatePersonalAccessTokenResponse) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePersonalAccessTokenResponse) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *CreatePersonalAccessTokenResponse) SetScope(v []string) {
	o.Scope = v
}

// GetName returns the Name field value
func (o *CreatePersonalAccessTokenResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePersonalAccessTokenResponse) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value
func (o *CreatePersonalAccessTokenResponse) GetOwner() BaseReferenceDto {
	if o == nil {
		var ret BaseReferenceDto
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenResponse) GetOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *CreatePersonalAccessTokenResponse) SetOwner(v BaseReferenceDto) {
	o.Owner = v
}

// GetCreated returns the Created field value
func (o *CreatePersonalAccessTokenResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CreatePersonalAccessTokenResponse) SetCreated(v time.Time) {
	o.Created = v
}

func (o CreatePersonalAccessTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["created"] = o.Created
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreatePersonalAccessTokenResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreatePersonalAccessTokenResponse := _CreatePersonalAccessTokenResponse{}

	if err = json.Unmarshal(bytes, &varCreatePersonalAccessTokenResponse); err == nil {
		*o = CreatePersonalAccessTokenResponse(varCreatePersonalAccessTokenResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePersonalAccessTokenResponse struct {
	value *CreatePersonalAccessTokenResponse
	isSet bool
}

func (v NullableCreatePersonalAccessTokenResponse) Get() *CreatePersonalAccessTokenResponse {
	return v.value
}

func (v *NullableCreatePersonalAccessTokenResponse) Set(val *CreatePersonalAccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePersonalAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePersonalAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePersonalAccessTokenResponse(val *CreatePersonalAccessTokenResponse) *NullableCreatePersonalAccessTokenResponse {
	return &NullableCreatePersonalAccessTokenResponse{value: val, isSet: true}
}

func (v NullableCreatePersonalAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePersonalAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


