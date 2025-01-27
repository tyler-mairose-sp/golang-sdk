# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the Source | [optional] [readonly] 
**Name** | **string** | Human-readable name of the source | 
**Description** | Pointer to **string** | Human-readable description of the source | [optional] 
**Owner** | [**SourceOwner**](SourceOwner.md) |  | 
**Cluster** | Pointer to [**SourceCluster**](SourceCluster.md) |  | [optional] 
**AccountCorrelationConfig** | Pointer to [**SourceAccountCorrelationConfig**](SourceAccountCorrelationConfig.md) |  | [optional] 
**AccountCorrelationRule** | Pointer to [**SourceAccountCorrelationRule**](SourceAccountCorrelationRule.md) |  | [optional] 
**ManagerCorrelationMapping** | Pointer to [**ManagerCorrelationMapping**](ManagerCorrelationMapping.md) |  | [optional] 
**ManagerCorrelationRule** | Pointer to [**SourceManagerCorrelationRule**](SourceManagerCorrelationRule.md) |  | [optional] 
**BeforeProvisioningRule** | Pointer to [**SourceBeforeProvisioningRule**](SourceBeforeProvisioningRule.md) |  | [optional] 
**Schemas** | Pointer to [**[]SourceSchemasInner**](SourceSchemasInner.md) | List of references to Schema objects | [optional] 
**PasswordPolicies** | Pointer to [**[]SourcePasswordPoliciesInner**](SourcePasswordPoliciesInner.md) | List of references to the associated PasswordPolicy objects. | [optional] 
**Features** | Pointer to [**[]SourceFeature**](SourceFeature.md) | Optional features that can be supported by a source. | [optional] 
**Type** | Pointer to **string** | Specifies the type of system being managed e.g. Active Directory, Workday, etc.. If you are creating a Delimited File source, you must set the &#x60;provisionasCsv&#x60; query parameter to &#x60;true&#x60;.  | [optional] 
**Connector** | **string** | Connector script name. | 
**ConnectorClass** | Pointer to **string** | The fully qualified name of the Java class that implements the connector interface. | [optional] 
**ConnectorAttributes** | Pointer to **map[string]interface{}** | Connector specific configuration; will differ from type to type. | [optional] 
**DeleteThreshold** | Pointer to **int32** | Number from 0 to 100 that specifies when to skip the delete phase. | [optional] 
**Authoritative** | Pointer to **bool** | When true indicates the source is referenced by an IdentityProfile. | [optional] 
**ManagementWorkgroup** | Pointer to [**SourceManagementWorkgroup**](SourceManagementWorkgroup.md) |  | [optional] 
**Healthy** | Pointer to **bool** | When true indicates a healthy source | [optional] 
**Status** | Pointer to **string** | A status identifier, giving specific information on why a source is healthy or not | [optional] 
**Since** | Pointer to **string** | Timestamp showing when a source health check was last performed | [optional] 
**ConnectorId** | Pointer to **string** | The id of connector | [optional] 
**ConnectorName** | Pointer to **string** | The name of the connector that was chosen on source creation | [optional] 
**ConnectionType** | Pointer to **string** | The type of connection (direct or file) | [optional] 
**ConnectorImplementstionId** | Pointer to **string** | The connector implementstion id | [optional] 

## Methods

### NewSource

`func NewSource(name string, owner SourceOwner, connector string, ) *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Source) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Source) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Source) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Source) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Source) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Source) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Source) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Source) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *Source) GetOwner() SourceOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Source) GetOwnerOk() (*SourceOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Source) SetOwner(v SourceOwner)`

SetOwner sets Owner field to given value.


### GetCluster

`func (o *Source) GetCluster() SourceCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Source) GetClusterOk() (*SourceCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Source) SetCluster(v SourceCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Source) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetAccountCorrelationConfig

`func (o *Source) GetAccountCorrelationConfig() SourceAccountCorrelationConfig`

GetAccountCorrelationConfig returns the AccountCorrelationConfig field if non-nil, zero value otherwise.

### GetAccountCorrelationConfigOk

`func (o *Source) GetAccountCorrelationConfigOk() (*SourceAccountCorrelationConfig, bool)`

GetAccountCorrelationConfigOk returns a tuple with the AccountCorrelationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCorrelationConfig

`func (o *Source) SetAccountCorrelationConfig(v SourceAccountCorrelationConfig)`

SetAccountCorrelationConfig sets AccountCorrelationConfig field to given value.

### HasAccountCorrelationConfig

`func (o *Source) HasAccountCorrelationConfig() bool`

HasAccountCorrelationConfig returns a boolean if a field has been set.

### GetAccountCorrelationRule

`func (o *Source) GetAccountCorrelationRule() SourceAccountCorrelationRule`

GetAccountCorrelationRule returns the AccountCorrelationRule field if non-nil, zero value otherwise.

### GetAccountCorrelationRuleOk

`func (o *Source) GetAccountCorrelationRuleOk() (*SourceAccountCorrelationRule, bool)`

GetAccountCorrelationRuleOk returns a tuple with the AccountCorrelationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCorrelationRule

`func (o *Source) SetAccountCorrelationRule(v SourceAccountCorrelationRule)`

SetAccountCorrelationRule sets AccountCorrelationRule field to given value.

### HasAccountCorrelationRule

`func (o *Source) HasAccountCorrelationRule() bool`

HasAccountCorrelationRule returns a boolean if a field has been set.

### GetManagerCorrelationMapping

`func (o *Source) GetManagerCorrelationMapping() ManagerCorrelationMapping`

GetManagerCorrelationMapping returns the ManagerCorrelationMapping field if non-nil, zero value otherwise.

### GetManagerCorrelationMappingOk

`func (o *Source) GetManagerCorrelationMappingOk() (*ManagerCorrelationMapping, bool)`

GetManagerCorrelationMappingOk returns a tuple with the ManagerCorrelationMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerCorrelationMapping

`func (o *Source) SetManagerCorrelationMapping(v ManagerCorrelationMapping)`

SetManagerCorrelationMapping sets ManagerCorrelationMapping field to given value.

### HasManagerCorrelationMapping

`func (o *Source) HasManagerCorrelationMapping() bool`

HasManagerCorrelationMapping returns a boolean if a field has been set.

### GetManagerCorrelationRule

`func (o *Source) GetManagerCorrelationRule() SourceManagerCorrelationRule`

GetManagerCorrelationRule returns the ManagerCorrelationRule field if non-nil, zero value otherwise.

### GetManagerCorrelationRuleOk

`func (o *Source) GetManagerCorrelationRuleOk() (*SourceManagerCorrelationRule, bool)`

GetManagerCorrelationRuleOk returns a tuple with the ManagerCorrelationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerCorrelationRule

`func (o *Source) SetManagerCorrelationRule(v SourceManagerCorrelationRule)`

SetManagerCorrelationRule sets ManagerCorrelationRule field to given value.

### HasManagerCorrelationRule

`func (o *Source) HasManagerCorrelationRule() bool`

HasManagerCorrelationRule returns a boolean if a field has been set.

### GetBeforeProvisioningRule

`func (o *Source) GetBeforeProvisioningRule() SourceBeforeProvisioningRule`

GetBeforeProvisioningRule returns the BeforeProvisioningRule field if non-nil, zero value otherwise.

### GetBeforeProvisioningRuleOk

`func (o *Source) GetBeforeProvisioningRuleOk() (*SourceBeforeProvisioningRule, bool)`

GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeProvisioningRule

`func (o *Source) SetBeforeProvisioningRule(v SourceBeforeProvisioningRule)`

SetBeforeProvisioningRule sets BeforeProvisioningRule field to given value.

### HasBeforeProvisioningRule

`func (o *Source) HasBeforeProvisioningRule() bool`

HasBeforeProvisioningRule returns a boolean if a field has been set.

### GetSchemas

`func (o *Source) GetSchemas() []SourceSchemasInner`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Source) GetSchemasOk() (*[]SourceSchemasInner, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Source) SetSchemas(v []SourceSchemasInner)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *Source) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetPasswordPolicies

`func (o *Source) GetPasswordPolicies() []SourcePasswordPoliciesInner`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *Source) GetPasswordPoliciesOk() (*[]SourcePasswordPoliciesInner, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicies

`func (o *Source) SetPasswordPolicies(v []SourcePasswordPoliciesInner)`

SetPasswordPolicies sets PasswordPolicies field to given value.

### HasPasswordPolicies

`func (o *Source) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### GetFeatures

`func (o *Source) GetFeatures() []SourceFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Source) GetFeaturesOk() (*[]SourceFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Source) SetFeatures(v []SourceFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Source) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetType

`func (o *Source) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Source) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Source) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Source) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConnector

`func (o *Source) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *Source) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *Source) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetConnectorClass

`func (o *Source) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *Source) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *Source) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.

### HasConnectorClass

`func (o *Source) HasConnectorClass() bool`

HasConnectorClass returns a boolean if a field has been set.

### GetConnectorAttributes

`func (o *Source) GetConnectorAttributes() map[string]interface{}`

GetConnectorAttributes returns the ConnectorAttributes field if non-nil, zero value otherwise.

### GetConnectorAttributesOk

`func (o *Source) GetConnectorAttributesOk() (*map[string]interface{}, bool)`

GetConnectorAttributesOk returns a tuple with the ConnectorAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorAttributes

`func (o *Source) SetConnectorAttributes(v map[string]interface{})`

SetConnectorAttributes sets ConnectorAttributes field to given value.

### HasConnectorAttributes

`func (o *Source) HasConnectorAttributes() bool`

HasConnectorAttributes returns a boolean if a field has been set.

### GetDeleteThreshold

`func (o *Source) GetDeleteThreshold() int32`

GetDeleteThreshold returns the DeleteThreshold field if non-nil, zero value otherwise.

### GetDeleteThresholdOk

`func (o *Source) GetDeleteThresholdOk() (*int32, bool)`

GetDeleteThresholdOk returns a tuple with the DeleteThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteThreshold

`func (o *Source) SetDeleteThreshold(v int32)`

SetDeleteThreshold sets DeleteThreshold field to given value.

### HasDeleteThreshold

`func (o *Source) HasDeleteThreshold() bool`

HasDeleteThreshold returns a boolean if a field has been set.

### GetAuthoritative

`func (o *Source) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *Source) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *Source) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *Source) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetManagementWorkgroup

`func (o *Source) GetManagementWorkgroup() SourceManagementWorkgroup`

GetManagementWorkgroup returns the ManagementWorkgroup field if non-nil, zero value otherwise.

### GetManagementWorkgroupOk

`func (o *Source) GetManagementWorkgroupOk() (*SourceManagementWorkgroup, bool)`

GetManagementWorkgroupOk returns a tuple with the ManagementWorkgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementWorkgroup

`func (o *Source) SetManagementWorkgroup(v SourceManagementWorkgroup)`

SetManagementWorkgroup sets ManagementWorkgroup field to given value.

### HasManagementWorkgroup

`func (o *Source) HasManagementWorkgroup() bool`

HasManagementWorkgroup returns a boolean if a field has been set.

### GetHealthy

`func (o *Source) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *Source) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *Source) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *Source) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetStatus

`func (o *Source) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Source) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Source) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Source) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSince

`func (o *Source) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *Source) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *Source) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *Source) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetConnectorId

`func (o *Source) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *Source) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *Source) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *Source) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetConnectorName

`func (o *Source) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *Source) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *Source) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.

### HasConnectorName

`func (o *Source) HasConnectorName() bool`

HasConnectorName returns a boolean if a field has been set.

### GetConnectionType

`func (o *Source) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *Source) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *Source) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *Source) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetConnectorImplementstionId

`func (o *Source) GetConnectorImplementstionId() string`

GetConnectorImplementstionId returns the ConnectorImplementstionId field if non-nil, zero value otherwise.

### GetConnectorImplementstionIdOk

`func (o *Source) GetConnectorImplementstionIdOk() (*string, bool)`

GetConnectorImplementstionIdOk returns a tuple with the ConnectorImplementstionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorImplementstionId

`func (o *Source) SetConnectorImplementstionId(v string)`

SetConnectorImplementstionId sets ConnectorImplementstionId field to given value.

### HasConnectorImplementstionId

`func (o *Source) HasConnectorImplementstionId() bool`

HasConnectorImplementstionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


