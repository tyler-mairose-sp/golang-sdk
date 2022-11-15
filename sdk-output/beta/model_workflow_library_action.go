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

// WorkflowLibraryAction struct for WorkflowLibraryAction
type WorkflowLibraryAction struct {
	// Action ID. This is a static namespaced ID for the action
	Id *string `json:"id,omitempty"`
	// Action Name
	Name *string `json:"name,omitempty"`
	// Action Description
	Description *string `json:"description,omitempty"`
	// One or more inputs that the action accepts
	FormFields []WorkflowLibraryFormFields `json:"formFields,omitempty"`
	// Defines the output schema, if any, that this action produces.
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLibraryAction WorkflowLibraryAction

// NewWorkflowLibraryAction instantiates a new WorkflowLibraryAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLibraryAction() *WorkflowLibraryAction {
	this := WorkflowLibraryAction{}
	return &this
}

// NewWorkflowLibraryActionWithDefaults instantiates a new WorkflowLibraryAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLibraryActionWithDefaults() *WorkflowLibraryAction {
	this := WorkflowLibraryAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowLibraryAction) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryAction) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowLibraryAction) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowLibraryAction) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowLibraryAction) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryAction) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowLibraryAction) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowLibraryAction) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowLibraryAction) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryAction) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowLibraryAction) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowLibraryAction) SetDescription(v string) {
	o.Description = &v
}

// GetFormFields returns the FormFields field value if set, zero value otherwise.
func (o *WorkflowLibraryAction) GetFormFields() []WorkflowLibraryFormFields {
	if o == nil || isNil(o.FormFields) {
		var ret []WorkflowLibraryFormFields
		return ret
	}
	return o.FormFields
}

// GetFormFieldsOk returns a tuple with the FormFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryAction) GetFormFieldsOk() ([]WorkflowLibraryFormFields, bool) {
	if o == nil || isNil(o.FormFields) {
		return nil, false
	}
	return o.FormFields, true
}

// HasFormFields returns a boolean if a field has been set.
func (o *WorkflowLibraryAction) HasFormFields() bool {
	if o != nil && !isNil(o.FormFields) {
		return true
	}

	return false
}

// SetFormFields gets a reference to the given []WorkflowLibraryFormFields and assigns it to the FormFields field.
func (o *WorkflowLibraryAction) SetFormFields(v []WorkflowLibraryFormFields) {
	o.FormFields = v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *WorkflowLibraryAction) GetOutputSchema() map[string]interface{} {
	if o == nil || isNil(o.OutputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryAction) GetOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.OutputSchema) {
		return map[string]interface{}{}, false
	}
	return o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *WorkflowLibraryAction) HasOutputSchema() bool {
	if o != nil && !isNil(o.OutputSchema) {
		return true
	}

	return false
}

// SetOutputSchema gets a reference to the given map[string]interface{} and assigns it to the OutputSchema field.
func (o *WorkflowLibraryAction) SetOutputSchema(v map[string]interface{}) {
	o.OutputSchema = v
}

func (o WorkflowLibraryAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.FormFields) {
		toSerialize["formFields"] = o.FormFields
	}
	if !isNil(o.OutputSchema) {
		toSerialize["outputSchema"] = o.OutputSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowLibraryAction) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowLibraryAction := _WorkflowLibraryAction{}

	if err = json.Unmarshal(bytes, &varWorkflowLibraryAction); err == nil {
		*o = WorkflowLibraryAction(varWorkflowLibraryAction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "formFields")
		delete(additionalProperties, "outputSchema")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLibraryAction struct {
	value *WorkflowLibraryAction
	isSet bool
}

func (v NullableWorkflowLibraryAction) Get() *WorkflowLibraryAction {
	return v.value
}

func (v *NullableWorkflowLibraryAction) Set(val *WorkflowLibraryAction) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryAction) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryAction(val *WorkflowLibraryAction) *NullableWorkflowLibraryAction {
	return &NullableWorkflowLibraryAction{value: val, isSet: true}
}

func (v NullableWorkflowLibraryAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


