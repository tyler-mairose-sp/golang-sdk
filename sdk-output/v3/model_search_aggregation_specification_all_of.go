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

// SearchAggregationSpecificationAllOf struct for SearchAggregationSpecificationAllOf
type SearchAggregationSpecificationAllOf struct {
	SubAggregation *SubSearchAggregationSpecification `json:"subAggregation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchAggregationSpecificationAllOf SearchAggregationSpecificationAllOf

// NewSearchAggregationSpecificationAllOf instantiates a new SearchAggregationSpecificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAggregationSpecificationAllOf() *SearchAggregationSpecificationAllOf {
	this := SearchAggregationSpecificationAllOf{}
	return &this
}

// NewSearchAggregationSpecificationAllOfWithDefaults instantiates a new SearchAggregationSpecificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAggregationSpecificationAllOfWithDefaults() *SearchAggregationSpecificationAllOf {
	this := SearchAggregationSpecificationAllOf{}
	return &this
}

// GetSubAggregation returns the SubAggregation field value if set, zero value otherwise.
func (o *SearchAggregationSpecificationAllOf) GetSubAggregation() SubSearchAggregationSpecification {
	if o == nil || isNil(o.SubAggregation) {
		var ret SubSearchAggregationSpecification
		return ret
	}
	return *o.SubAggregation
}

// GetSubAggregationOk returns a tuple with the SubAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAggregationSpecificationAllOf) GetSubAggregationOk() (*SubSearchAggregationSpecification, bool) {
	if o == nil || isNil(o.SubAggregation) {
		return nil, false
	}
	return o.SubAggregation, true
}

// HasSubAggregation returns a boolean if a field has been set.
func (o *SearchAggregationSpecificationAllOf) HasSubAggregation() bool {
	if o != nil && !isNil(o.SubAggregation) {
		return true
	}

	return false
}

// SetSubAggregation gets a reference to the given SubSearchAggregationSpecification and assigns it to the SubAggregation field.
func (o *SearchAggregationSpecificationAllOf) SetSubAggregation(v SubSearchAggregationSpecification) {
	o.SubAggregation = &v
}

func (o SearchAggregationSpecificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubAggregation) {
		toSerialize["subAggregation"] = o.SubAggregation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchAggregationSpecificationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSearchAggregationSpecificationAllOf := _SearchAggregationSpecificationAllOf{}

	if err = json.Unmarshal(bytes, &varSearchAggregationSpecificationAllOf); err == nil {
		*o = SearchAggregationSpecificationAllOf(varSearchAggregationSpecificationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subAggregation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchAggregationSpecificationAllOf struct {
	value *SearchAggregationSpecificationAllOf
	isSet bool
}

func (v NullableSearchAggregationSpecificationAllOf) Get() *SearchAggregationSpecificationAllOf {
	return v.value
}

func (v *NullableSearchAggregationSpecificationAllOf) Set(val *SearchAggregationSpecificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAggregationSpecificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAggregationSpecificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAggregationSpecificationAllOf(val *SearchAggregationSpecificationAllOf) *NullableSearchAggregationSpecificationAllOf {
	return &NullableSearchAggregationSpecificationAllOf{value: val, isSet: true}
}

func (v NullableSearchAggregationSpecificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAggregationSpecificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


