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

// ManagedCluster Managed Cluster
type ManagedCluster struct {
	// ManagedCluster ID
	Id string `json:"id"`
	// ManagedCluster name
	Name *string `json:"name,omitempty"`
	// ManagedCluster pod
	Pod *string `json:"pod,omitempty"`
	// ManagedCluster org
	Org *string `json:"org,omitempty"`
	Type *ManagedClusterTypes `json:"type,omitempty"`
	// ManagedProcess configuration map
	Configuration *map[string]string `json:"configuration,omitempty"`
	KeyPair *ManagedClusterKeyPair `json:"keyPair,omitempty"`
	Attributes *ManagedClusterAttributes `json:"attributes,omitempty"`
	// ManagedCluster description
	Description *string `json:"description,omitempty"`
	Redis *ManagedClusterRedis `json:"redis,omitempty"`
	ClientType ManagedClientType `json:"clientType"`
	// CCG version used by the ManagedCluster
	CcgVersion string `json:"ccgVersion"`
	// boolean flag indiacting whether or not the cluster configuration is pinned
	PinnedConfig *bool `json:"pinnedConfig,omitempty"`
	LogConfiguration *ClientLogConfiguration `json:"logConfiguration,omitempty"`
	// Whether or not the cluster is operational or not
	Operational *bool `json:"operational,omitempty"`
	// Cluster status
	Status *string `json:"status,omitempty"`
	// Public key certificate
	PublicKeyCertificate *string `json:"publicKeyCertificate,omitempty"`
	// Public key thumbprint
	PublicKeyThumbprint *string `json:"publicKeyThumbprint,omitempty"`
	// Public key
	PublicKey *string `json:"publicKey,omitempty"`
	// Key describing any immediate cluster alerts
	AlertKey *string `json:"alertKey,omitempty"`
	// List of clients in a cluster
	ClientIds []string `json:"clientIds,omitempty"`
	// Number of services bound to a cluster
	ServiceCount *int32 `json:"serviceCount,omitempty"`
	// CC ID only used in calling CC, will be removed without notice when Migration to CEGS is finished
	CcId *string `json:"ccId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedCluster ManagedCluster

// NewManagedCluster instantiates a new ManagedCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedCluster(id string, clientType ManagedClientType, ccgVersion string) *ManagedCluster {
	this := ManagedCluster{}
	this.Id = id
	this.ClientType = clientType
	this.CcgVersion = ccgVersion
	var pinnedConfig bool = false
	this.PinnedConfig = &pinnedConfig
	var operational bool = false
	this.Operational = &operational
	var serviceCount int32 = 0
	this.ServiceCount = &serviceCount
	var ccId string = "0"
	this.CcId = &ccId
	return &this
}

// NewManagedClusterWithDefaults instantiates a new ManagedCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClusterWithDefaults() *ManagedCluster {
	this := ManagedCluster{}
	var pinnedConfig bool = false
	this.PinnedConfig = &pinnedConfig
	var operational bool = false
	this.Operational = &operational
	var serviceCount int32 = 0
	this.ServiceCount = &serviceCount
	var ccId string = "0"
	this.CcId = &ccId
	return &this
}

// GetId returns the Id field value
func (o *ManagedCluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ManagedCluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagedCluster) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagedCluster) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagedCluster) SetName(v string) {
	o.Name = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *ManagedCluster) GetPod() string {
	if o == nil || isNil(o.Pod) {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetPodOk() (*string, bool) {
	if o == nil || isNil(o.Pod) {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *ManagedCluster) HasPod() bool {
	if o != nil && !isNil(o.Pod) {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *ManagedCluster) SetPod(v string) {
	o.Pod = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *ManagedCluster) GetOrg() string {
	if o == nil || isNil(o.Org) {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetOrgOk() (*string, bool) {
	if o == nil || isNil(o.Org) {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *ManagedCluster) HasOrg() bool {
	if o != nil && !isNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *ManagedCluster) SetOrg(v string) {
	o.Org = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ManagedCluster) GetType() ManagedClusterTypes {
	if o == nil || isNil(o.Type) {
		var ret ManagedClusterTypes
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetTypeOk() (*ManagedClusterTypes, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ManagedCluster) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ManagedClusterTypes and assigns it to the Type field.
func (o *ManagedCluster) SetType(v ManagedClusterTypes) {
	o.Type = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ManagedCluster) GetConfiguration() map[string]string {
	if o == nil || isNil(o.Configuration) {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ManagedCluster) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *ManagedCluster) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

// GetKeyPair returns the KeyPair field value if set, zero value otherwise.
func (o *ManagedCluster) GetKeyPair() ManagedClusterKeyPair {
	if o == nil || isNil(o.KeyPair) {
		var ret ManagedClusterKeyPair
		return ret
	}
	return *o.KeyPair
}

// GetKeyPairOk returns a tuple with the KeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetKeyPairOk() (*ManagedClusterKeyPair, bool) {
	if o == nil || isNil(o.KeyPair) {
		return nil, false
	}
	return o.KeyPair, true
}

// HasKeyPair returns a boolean if a field has been set.
func (o *ManagedCluster) HasKeyPair() bool {
	if o != nil && !isNil(o.KeyPair) {
		return true
	}

	return false
}

// SetKeyPair gets a reference to the given ManagedClusterKeyPair and assigns it to the KeyPair field.
func (o *ManagedCluster) SetKeyPair(v ManagedClusterKeyPair) {
	o.KeyPair = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedCluster) GetAttributes() ManagedClusterAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedClusterAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetAttributesOk() (*ManagedClusterAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedCluster) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedClusterAttributes and assigns it to the Attributes field.
func (o *ManagedCluster) SetAttributes(v ManagedClusterAttributes) {
	o.Attributes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManagedCluster) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManagedCluster) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManagedCluster) SetDescription(v string) {
	o.Description = &v
}

// GetRedis returns the Redis field value if set, zero value otherwise.
func (o *ManagedCluster) GetRedis() ManagedClusterRedis {
	if o == nil || isNil(o.Redis) {
		var ret ManagedClusterRedis
		return ret
	}
	return *o.Redis
}

// GetRedisOk returns a tuple with the Redis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetRedisOk() (*ManagedClusterRedis, bool) {
	if o == nil || isNil(o.Redis) {
		return nil, false
	}
	return o.Redis, true
}

// HasRedis returns a boolean if a field has been set.
func (o *ManagedCluster) HasRedis() bool {
	if o != nil && !isNil(o.Redis) {
		return true
	}

	return false
}

// SetRedis gets a reference to the given ManagedClusterRedis and assigns it to the Redis field.
func (o *ManagedCluster) SetRedis(v ManagedClusterRedis) {
	o.Redis = &v
}

// GetClientType returns the ClientType field value
func (o *ManagedCluster) GetClientType() ManagedClientType {
	if o == nil {
		var ret ManagedClientType
		return ret
	}

	return o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetClientTypeOk() (*ManagedClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientType, true
}

// SetClientType sets field value
func (o *ManagedCluster) SetClientType(v ManagedClientType) {
	o.ClientType = v
}

// GetCcgVersion returns the CcgVersion field value
func (o *ManagedCluster) GetCcgVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CcgVersion
}

// GetCcgVersionOk returns a tuple with the CcgVersion field value
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetCcgVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcgVersion, true
}

// SetCcgVersion sets field value
func (o *ManagedCluster) SetCcgVersion(v string) {
	o.CcgVersion = v
}

// GetPinnedConfig returns the PinnedConfig field value if set, zero value otherwise.
func (o *ManagedCluster) GetPinnedConfig() bool {
	if o == nil || isNil(o.PinnedConfig) {
		var ret bool
		return ret
	}
	return *o.PinnedConfig
}

// GetPinnedConfigOk returns a tuple with the PinnedConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetPinnedConfigOk() (*bool, bool) {
	if o == nil || isNil(o.PinnedConfig) {
		return nil, false
	}
	return o.PinnedConfig, true
}

// HasPinnedConfig returns a boolean if a field has been set.
func (o *ManagedCluster) HasPinnedConfig() bool {
	if o != nil && !isNil(o.PinnedConfig) {
		return true
	}

	return false
}

// SetPinnedConfig gets a reference to the given bool and assigns it to the PinnedConfig field.
func (o *ManagedCluster) SetPinnedConfig(v bool) {
	o.PinnedConfig = &v
}

// GetLogConfiguration returns the LogConfiguration field value if set, zero value otherwise.
func (o *ManagedCluster) GetLogConfiguration() ClientLogConfiguration {
	if o == nil || isNil(o.LogConfiguration) {
		var ret ClientLogConfiguration
		return ret
	}
	return *o.LogConfiguration
}

// GetLogConfigurationOk returns a tuple with the LogConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetLogConfigurationOk() (*ClientLogConfiguration, bool) {
	if o == nil || isNil(o.LogConfiguration) {
		return nil, false
	}
	return o.LogConfiguration, true
}

// HasLogConfiguration returns a boolean if a field has been set.
func (o *ManagedCluster) HasLogConfiguration() bool {
	if o != nil && !isNil(o.LogConfiguration) {
		return true
	}

	return false
}

// SetLogConfiguration gets a reference to the given ClientLogConfiguration and assigns it to the LogConfiguration field.
func (o *ManagedCluster) SetLogConfiguration(v ClientLogConfiguration) {
	o.LogConfiguration = &v
}

// GetOperational returns the Operational field value if set, zero value otherwise.
func (o *ManagedCluster) GetOperational() bool {
	if o == nil || isNil(o.Operational) {
		var ret bool
		return ret
	}
	return *o.Operational
}

// GetOperationalOk returns a tuple with the Operational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetOperationalOk() (*bool, bool) {
	if o == nil || isNil(o.Operational) {
		return nil, false
	}
	return o.Operational, true
}

// HasOperational returns a boolean if a field has been set.
func (o *ManagedCluster) HasOperational() bool {
	if o != nil && !isNil(o.Operational) {
		return true
	}

	return false
}

// SetOperational gets a reference to the given bool and assigns it to the Operational field.
func (o *ManagedCluster) SetOperational(v bool) {
	o.Operational = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManagedCluster) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ManagedCluster) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ManagedCluster) SetStatus(v string) {
	o.Status = &v
}

// GetPublicKeyCertificate returns the PublicKeyCertificate field value if set, zero value otherwise.
func (o *ManagedCluster) GetPublicKeyCertificate() string {
	if o == nil || isNil(o.PublicKeyCertificate) {
		var ret string
		return ret
	}
	return *o.PublicKeyCertificate
}

// GetPublicKeyCertificateOk returns a tuple with the PublicKeyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetPublicKeyCertificateOk() (*string, bool) {
	if o == nil || isNil(o.PublicKeyCertificate) {
		return nil, false
	}
	return o.PublicKeyCertificate, true
}

// HasPublicKeyCertificate returns a boolean if a field has been set.
func (o *ManagedCluster) HasPublicKeyCertificate() bool {
	if o != nil && !isNil(o.PublicKeyCertificate) {
		return true
	}

	return false
}

// SetPublicKeyCertificate gets a reference to the given string and assigns it to the PublicKeyCertificate field.
func (o *ManagedCluster) SetPublicKeyCertificate(v string) {
	o.PublicKeyCertificate = &v
}

// GetPublicKeyThumbprint returns the PublicKeyThumbprint field value if set, zero value otherwise.
func (o *ManagedCluster) GetPublicKeyThumbprint() string {
	if o == nil || isNil(o.PublicKeyThumbprint) {
		var ret string
		return ret
	}
	return *o.PublicKeyThumbprint
}

// GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetPublicKeyThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.PublicKeyThumbprint) {
		return nil, false
	}
	return o.PublicKeyThumbprint, true
}

// HasPublicKeyThumbprint returns a boolean if a field has been set.
func (o *ManagedCluster) HasPublicKeyThumbprint() bool {
	if o != nil && !isNil(o.PublicKeyThumbprint) {
		return true
	}

	return false
}

// SetPublicKeyThumbprint gets a reference to the given string and assigns it to the PublicKeyThumbprint field.
func (o *ManagedCluster) SetPublicKeyThumbprint(v string) {
	o.PublicKeyThumbprint = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ManagedCluster) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ManagedCluster) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ManagedCluster) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetAlertKey returns the AlertKey field value if set, zero value otherwise.
func (o *ManagedCluster) GetAlertKey() string {
	if o == nil || isNil(o.AlertKey) {
		var ret string
		return ret
	}
	return *o.AlertKey
}

// GetAlertKeyOk returns a tuple with the AlertKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetAlertKeyOk() (*string, bool) {
	if o == nil || isNil(o.AlertKey) {
		return nil, false
	}
	return o.AlertKey, true
}

// HasAlertKey returns a boolean if a field has been set.
func (o *ManagedCluster) HasAlertKey() bool {
	if o != nil && !isNil(o.AlertKey) {
		return true
	}

	return false
}

// SetAlertKey gets a reference to the given string and assigns it to the AlertKey field.
func (o *ManagedCluster) SetAlertKey(v string) {
	o.AlertKey = &v
}

// GetClientIds returns the ClientIds field value if set, zero value otherwise.
func (o *ManagedCluster) GetClientIds() []string {
	if o == nil || isNil(o.ClientIds) {
		var ret []string
		return ret
	}
	return o.ClientIds
}

// GetClientIdsOk returns a tuple with the ClientIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetClientIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ClientIds) {
		return nil, false
	}
	return o.ClientIds, true
}

// HasClientIds returns a boolean if a field has been set.
func (o *ManagedCluster) HasClientIds() bool {
	if o != nil && !isNil(o.ClientIds) {
		return true
	}

	return false
}

// SetClientIds gets a reference to the given []string and assigns it to the ClientIds field.
func (o *ManagedCluster) SetClientIds(v []string) {
	o.ClientIds = v
}

// GetServiceCount returns the ServiceCount field value if set, zero value otherwise.
func (o *ManagedCluster) GetServiceCount() int32 {
	if o == nil || isNil(o.ServiceCount) {
		var ret int32
		return ret
	}
	return *o.ServiceCount
}

// GetServiceCountOk returns a tuple with the ServiceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetServiceCountOk() (*int32, bool) {
	if o == nil || isNil(o.ServiceCount) {
		return nil, false
	}
	return o.ServiceCount, true
}

// HasServiceCount returns a boolean if a field has been set.
func (o *ManagedCluster) HasServiceCount() bool {
	if o != nil && !isNil(o.ServiceCount) {
		return true
	}

	return false
}

// SetServiceCount gets a reference to the given int32 and assigns it to the ServiceCount field.
func (o *ManagedCluster) SetServiceCount(v int32) {
	o.ServiceCount = &v
}

// GetCcId returns the CcId field value if set, zero value otherwise.
func (o *ManagedCluster) GetCcId() string {
	if o == nil || isNil(o.CcId) {
		var ret string
		return ret
	}
	return *o.CcId
}

// GetCcIdOk returns a tuple with the CcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedCluster) GetCcIdOk() (*string, bool) {
	if o == nil || isNil(o.CcId) {
		return nil, false
	}
	return o.CcId, true
}

// HasCcId returns a boolean if a field has been set.
func (o *ManagedCluster) HasCcId() bool {
	if o != nil && !isNil(o.CcId) {
		return true
	}

	return false
}

// SetCcId gets a reference to the given string and assigns it to the CcId field.
func (o *ManagedCluster) SetCcId(v string) {
	o.CcId = &v
}

func (o ManagedCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Pod) {
		toSerialize["pod"] = o.Pod
	}
	if !isNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !isNil(o.KeyPair) {
		toSerialize["keyPair"] = o.KeyPair
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Redis) {
		toSerialize["redis"] = o.Redis
	}
	if true {
		toSerialize["clientType"] = o.ClientType
	}
	if true {
		toSerialize["ccgVersion"] = o.CcgVersion
	}
	if !isNil(o.PinnedConfig) {
		toSerialize["pinnedConfig"] = o.PinnedConfig
	}
	if !isNil(o.LogConfiguration) {
		toSerialize["logConfiguration"] = o.LogConfiguration
	}
	if !isNil(o.Operational) {
		toSerialize["operational"] = o.Operational
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.PublicKeyCertificate) {
		toSerialize["publicKeyCertificate"] = o.PublicKeyCertificate
	}
	if !isNil(o.PublicKeyThumbprint) {
		toSerialize["publicKeyThumbprint"] = o.PublicKeyThumbprint
	}
	if !isNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !isNil(o.AlertKey) {
		toSerialize["alertKey"] = o.AlertKey
	}
	if !isNil(o.ClientIds) {
		toSerialize["clientIds"] = o.ClientIds
	}
	if !isNil(o.ServiceCount) {
		toSerialize["serviceCount"] = o.ServiceCount
	}
	if !isNil(o.CcId) {
		toSerialize["ccId"] = o.CcId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ManagedCluster) UnmarshalJSON(bytes []byte) (err error) {
	varManagedCluster := _ManagedCluster{}

	if err = json.Unmarshal(bytes, &varManagedCluster); err == nil {
		*o = ManagedCluster(varManagedCluster)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pod")
		delete(additionalProperties, "org")
		delete(additionalProperties, "type")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "keyPair")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "description")
		delete(additionalProperties, "redis")
		delete(additionalProperties, "clientType")
		delete(additionalProperties, "ccgVersion")
		delete(additionalProperties, "pinnedConfig")
		delete(additionalProperties, "logConfiguration")
		delete(additionalProperties, "operational")
		delete(additionalProperties, "status")
		delete(additionalProperties, "publicKeyCertificate")
		delete(additionalProperties, "publicKeyThumbprint")
		delete(additionalProperties, "publicKey")
		delete(additionalProperties, "alertKey")
		delete(additionalProperties, "clientIds")
		delete(additionalProperties, "serviceCount")
		delete(additionalProperties, "ccId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedCluster struct {
	value *ManagedCluster
	isSet bool
}

func (v NullableManagedCluster) Get() *ManagedCluster {
	return v.value
}

func (v *NullableManagedCluster) Set(val *ManagedCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedCluster(val *ManagedCluster) *NullableManagedCluster {
	return &NullableManagedCluster{value: val, isSet: true}
}

func (v NullableManagedCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


