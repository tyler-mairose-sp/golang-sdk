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

// checks if the SearchAggregationSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAggregationSpecification{}

// SearchAggregationSpecification struct for SearchAggregationSpecification
type SearchAggregationSpecification struct {
	Nested *NestedAggregation `json:"nested,omitempty"`
	Metric *MetricAggregation `json:"metric,omitempty"`
	Filter *FilterAggregation `json:"filter,omitempty"`
	Bucket *BucketAggregation `json:"bucket,omitempty"`
	SubAggregation *SubSearchAggregationSpecification `json:"subAggregation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchAggregationSpecification SearchAggregationSpecification

// NewSearchAggregationSpecification instantiates a new SearchAggregationSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAggregationSpecification() *SearchAggregationSpecification {
	this := SearchAggregationSpecification{}
	return &this
}

// NewSearchAggregationSpecificationWithDefaults instantiates a new SearchAggregationSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAggregationSpecificationWithDefaults() *SearchAggregationSpecification {
	this := SearchAggregationSpecification{}
	return &this
}

// GetNested returns the Nested field value if set, zero value otherwise.
func (o *SearchAggregationSpecification) GetNested() NestedAggregation {
	if o == nil || isNil(o.Nested) {
		var ret NestedAggregation
		return ret
	}
	return *o.Nested
}

// GetNestedOk returns a tuple with the Nested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAggregationSpecification) GetNestedOk() (*NestedAggregation, bool) {
	if o == nil || isNil(o.Nested) {
		return nil, false
	}
	return o.Nested, true
}

// HasNested returns a boolean if a field has been set.
func (o *SearchAggregationSpecification) HasNested() bool {
	if o != nil && !isNil(o.Nested) {
		return true
	}

	return false
}

// SetNested gets a reference to the given NestedAggregation and assigns it to the Nested field.
func (o *SearchAggregationSpecification) SetNested(v NestedAggregation) {
	o.Nested = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SearchAggregationSpecification) GetMetric() MetricAggregation {
	if o == nil || isNil(o.Metric) {
		var ret MetricAggregation
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAggregationSpecification) GetMetricOk() (*MetricAggregation, bool) {
	if o == nil || isNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SearchAggregationSpecification) HasMetric() bool {
	if o != nil && !isNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given MetricAggregation and assigns it to the Metric field.
func (o *SearchAggregationSpecification) SetMetric(v MetricAggregation) {
	o.Metric = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SearchAggregationSpecification) GetFilter() FilterAggregation {
	if o == nil || isNil(o.Filter) {
		var ret FilterAggregation
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAggregationSpecification) GetFilterOk() (*FilterAggregation, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SearchAggregationSpecification) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given FilterAggregation and assigns it to the Filter field.
func (o *SearchAggregationSpecification) SetFilter(v FilterAggregation) {
	o.Filter = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *SearchAggregationSpecification) GetBucket() BucketAggregation {
	if o == nil || isNil(o.Bucket) {
		var ret BucketAggregation
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAggregationSpecification) GetBucketOk() (*BucketAggregation, bool) {
	if o == nil || isNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *SearchAggregationSpecification) HasBucket() bool {
	if o != nil && !isNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given BucketAggregation and assigns it to the Bucket field.
func (o *SearchAggregationSpecification) SetBucket(v BucketAggregation) {
	o.Bucket = &v
}

// GetSubAggregation returns the SubAggregation field value if set, zero value otherwise.
func (o *SearchAggregationSpecification) GetSubAggregation() SubSearchAggregationSpecification {
	if o == nil || isNil(o.SubAggregation) {
		var ret SubSearchAggregationSpecification
		return ret
	}
	return *o.SubAggregation
}

// GetSubAggregationOk returns a tuple with the SubAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAggregationSpecification) GetSubAggregationOk() (*SubSearchAggregationSpecification, bool) {
	if o == nil || isNil(o.SubAggregation) {
		return nil, false
	}
	return o.SubAggregation, true
}

// HasSubAggregation returns a boolean if a field has been set.
func (o *SearchAggregationSpecification) HasSubAggregation() bool {
	if o != nil && !isNil(o.SubAggregation) {
		return true
	}

	return false
}

// SetSubAggregation gets a reference to the given SubSearchAggregationSpecification and assigns it to the SubAggregation field.
func (o *SearchAggregationSpecification) SetSubAggregation(v SubSearchAggregationSpecification) {
	o.SubAggregation = &v
}

func (o SearchAggregationSpecification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAggregationSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Nested) {
		toSerialize["nested"] = o.Nested
	}
	if !isNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !isNil(o.SubAggregation) {
		toSerialize["subAggregation"] = o.SubAggregation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchAggregationSpecification) UnmarshalJSON(bytes []byte) (err error) {
	varSearchAggregationSpecification := _SearchAggregationSpecification{}

	if err = json.Unmarshal(bytes, &varSearchAggregationSpecification); err == nil {
		*o = SearchAggregationSpecification(varSearchAggregationSpecification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nested")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "subAggregation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchAggregationSpecification struct {
	value *SearchAggregationSpecification
	isSet bool
}

func (v NullableSearchAggregationSpecification) Get() *SearchAggregationSpecification {
	return v.value
}

func (v *NullableSearchAggregationSpecification) Set(val *SearchAggregationSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAggregationSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAggregationSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAggregationSpecification(val *SearchAggregationSpecification) *NullableSearchAggregationSpecification {
	return &NullableSearchAggregationSpecification{value: val, isSet: true}
}

func (v NullableSearchAggregationSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAggregationSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


