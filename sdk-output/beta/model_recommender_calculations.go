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

// RecommenderCalculations struct for RecommenderCalculations
type RecommenderCalculations struct {
	// The ID of the identity
	IdentityId *string `json:"identityId,omitempty"`
	// The entitlement ID
	EntitlementId *string `json:"entitlementId,omitempty"`
	// The actual recommendation
	Recommendation *string `json:"recommendation,omitempty"`
	// The overall weighted score
	OverallWeightedScore *float32 `json:"overallWeightedScore,omitempty"`
	// The weighted score of each individual feature
	FeatureWeightedScores *map[string]float32 `json:"featureWeightedScores,omitempty"`
	// The configured value against which the overallWeightedScore is compared
	Threshold *float32 `json:"threshold,omitempty"`
	// The values for your configured features
	IdentityAttributes *map[string]RecommenderCalculationsIdentityAttributesValue `json:"identityAttributes,omitempty"`
	FeatureValues *FeatureValueDto `json:"featureValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommenderCalculations RecommenderCalculations

// NewRecommenderCalculations instantiates a new RecommenderCalculations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommenderCalculations() *RecommenderCalculations {
	this := RecommenderCalculations{}
	return &this
}

// NewRecommenderCalculationsWithDefaults instantiates a new RecommenderCalculations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommenderCalculationsWithDefaults() *RecommenderCalculations {
	this := RecommenderCalculations{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *RecommenderCalculations) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetEntitlementId returns the EntitlementId field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetEntitlementId() string {
	if o == nil || isNil(o.EntitlementId) {
		var ret string
		return ret
	}
	return *o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetEntitlementIdOk() (*string, bool) {
	if o == nil || isNil(o.EntitlementId) {
		return nil, false
	}
	return o.EntitlementId, true
}

// HasEntitlementId returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasEntitlementId() bool {
	if o != nil && !isNil(o.EntitlementId) {
		return true
	}

	return false
}

// SetEntitlementId gets a reference to the given string and assigns it to the EntitlementId field.
func (o *RecommenderCalculations) SetEntitlementId(v string) {
	o.EntitlementId = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetRecommendation() string {
	if o == nil || isNil(o.Recommendation) {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetRecommendationOk() (*string, bool) {
	if o == nil || isNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasRecommendation() bool {
	if o != nil && !isNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *RecommenderCalculations) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetOverallWeightedScore returns the OverallWeightedScore field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetOverallWeightedScore() float32 {
	if o == nil || isNil(o.OverallWeightedScore) {
		var ret float32
		return ret
	}
	return *o.OverallWeightedScore
}

// GetOverallWeightedScoreOk returns a tuple with the OverallWeightedScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetOverallWeightedScoreOk() (*float32, bool) {
	if o == nil || isNil(o.OverallWeightedScore) {
		return nil, false
	}
	return o.OverallWeightedScore, true
}

// HasOverallWeightedScore returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasOverallWeightedScore() bool {
	if o != nil && !isNil(o.OverallWeightedScore) {
		return true
	}

	return false
}

// SetOverallWeightedScore gets a reference to the given float32 and assigns it to the OverallWeightedScore field.
func (o *RecommenderCalculations) SetOverallWeightedScore(v float32) {
	o.OverallWeightedScore = &v
}

// GetFeatureWeightedScores returns the FeatureWeightedScores field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetFeatureWeightedScores() map[string]float32 {
	if o == nil || isNil(o.FeatureWeightedScores) {
		var ret map[string]float32
		return ret
	}
	return *o.FeatureWeightedScores
}

// GetFeatureWeightedScoresOk returns a tuple with the FeatureWeightedScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetFeatureWeightedScoresOk() (*map[string]float32, bool) {
	if o == nil || isNil(o.FeatureWeightedScores) {
		return nil, false
	}
	return o.FeatureWeightedScores, true
}

// HasFeatureWeightedScores returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasFeatureWeightedScores() bool {
	if o != nil && !isNil(o.FeatureWeightedScores) {
		return true
	}

	return false
}

// SetFeatureWeightedScores gets a reference to the given map[string]float32 and assigns it to the FeatureWeightedScores field.
func (o *RecommenderCalculations) SetFeatureWeightedScores(v map[string]float32) {
	o.FeatureWeightedScores = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetThreshold() float32 {
	if o == nil || isNil(o.Threshold) {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetThresholdOk() (*float32, bool) {
	if o == nil || isNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasThreshold() bool {
	if o != nil && !isNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *RecommenderCalculations) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetIdentityAttributes returns the IdentityAttributes field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetIdentityAttributes() map[string]RecommenderCalculationsIdentityAttributesValue {
	if o == nil || isNil(o.IdentityAttributes) {
		var ret map[string]RecommenderCalculationsIdentityAttributesValue
		return ret
	}
	return *o.IdentityAttributes
}

// GetIdentityAttributesOk returns a tuple with the IdentityAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetIdentityAttributesOk() (*map[string]RecommenderCalculationsIdentityAttributesValue, bool) {
	if o == nil || isNil(o.IdentityAttributes) {
		return nil, false
	}
	return o.IdentityAttributes, true
}

// HasIdentityAttributes returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasIdentityAttributes() bool {
	if o != nil && !isNil(o.IdentityAttributes) {
		return true
	}

	return false
}

// SetIdentityAttributes gets a reference to the given map[string]RecommenderCalculationsIdentityAttributesValue and assigns it to the IdentityAttributes field.
func (o *RecommenderCalculations) SetIdentityAttributes(v map[string]RecommenderCalculationsIdentityAttributesValue) {
	o.IdentityAttributes = &v
}

// GetFeatureValues returns the FeatureValues field value if set, zero value otherwise.
func (o *RecommenderCalculations) GetFeatureValues() FeatureValueDto {
	if o == nil || isNil(o.FeatureValues) {
		var ret FeatureValueDto
		return ret
	}
	return *o.FeatureValues
}

// GetFeatureValuesOk returns a tuple with the FeatureValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculations) GetFeatureValuesOk() (*FeatureValueDto, bool) {
	if o == nil || isNil(o.FeatureValues) {
		return nil, false
	}
	return o.FeatureValues, true
}

// HasFeatureValues returns a boolean if a field has been set.
func (o *RecommenderCalculations) HasFeatureValues() bool {
	if o != nil && !isNil(o.FeatureValues) {
		return true
	}

	return false
}

// SetFeatureValues gets a reference to the given FeatureValueDto and assigns it to the FeatureValues field.
func (o *RecommenderCalculations) SetFeatureValues(v FeatureValueDto) {
	o.FeatureValues = &v
}

func (o RecommenderCalculations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.EntitlementId) {
		toSerialize["entitlementId"] = o.EntitlementId
	}
	if !isNil(o.Recommendation) {
		toSerialize["recommendation"] = o.Recommendation
	}
	if !isNil(o.OverallWeightedScore) {
		toSerialize["overallWeightedScore"] = o.OverallWeightedScore
	}
	if !isNil(o.FeatureWeightedScores) {
		toSerialize["featureWeightedScores"] = o.FeatureWeightedScores
	}
	if !isNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !isNil(o.IdentityAttributes) {
		toSerialize["identityAttributes"] = o.IdentityAttributes
	}
	if !isNil(o.FeatureValues) {
		toSerialize["featureValues"] = o.FeatureValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommenderCalculations) UnmarshalJSON(bytes []byte) (err error) {
	varRecommenderCalculations := _RecommenderCalculations{}

	if err = json.Unmarshal(bytes, &varRecommenderCalculations); err == nil {
		*o = RecommenderCalculations(varRecommenderCalculations)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "recommendation")
		delete(additionalProperties, "overallWeightedScore")
		delete(additionalProperties, "featureWeightedScores")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "identityAttributes")
		delete(additionalProperties, "featureValues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommenderCalculations struct {
	value *RecommenderCalculations
	isSet bool
}

func (v NullableRecommenderCalculations) Get() *RecommenderCalculations {
	return v.value
}

func (v *NullableRecommenderCalculations) Set(val *RecommenderCalculations) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommenderCalculations) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommenderCalculations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommenderCalculations(val *RecommenderCalculations) *NullableRecommenderCalculations {
	return &NullableRecommenderCalculations{value: val, isSet: true}
}

func (v NullableRecommenderCalculations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommenderCalculations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


