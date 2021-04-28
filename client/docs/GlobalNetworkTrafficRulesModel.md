# GlobalNetworkTrafficRulesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseMultipleStreamsPerJob** | **bool** | If *true*, Veeam Backup &amp; Replication uses multiple TCP/IP transfer connection for every job session. | 
**UploadStreamsCount** | Pointer to **int32** | Number of TCP/IP connections per job. | [optional] 
**TrafficRules** | Pointer to [**[]TrafficRuleModel**](TrafficRuleModel.md) | Array of traffic rules. | [optional] 
**PreferredNetworks** | Pointer to [**PreferredNetworksModel**](PreferredNetworksModel.md) |  | [optional] 

## Methods

### NewGlobalNetworkTrafficRulesModel

`func NewGlobalNetworkTrafficRulesModel(useMultipleStreamsPerJob bool, ) *GlobalNetworkTrafficRulesModel`

NewGlobalNetworkTrafficRulesModel instantiates a new GlobalNetworkTrafficRulesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalNetworkTrafficRulesModelWithDefaults

`func NewGlobalNetworkTrafficRulesModelWithDefaults() *GlobalNetworkTrafficRulesModel`

NewGlobalNetworkTrafficRulesModelWithDefaults instantiates a new GlobalNetworkTrafficRulesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseMultipleStreamsPerJob

`func (o *GlobalNetworkTrafficRulesModel) GetUseMultipleStreamsPerJob() bool`

GetUseMultipleStreamsPerJob returns the UseMultipleStreamsPerJob field if non-nil, zero value otherwise.

### GetUseMultipleStreamsPerJobOk

`func (o *GlobalNetworkTrafficRulesModel) GetUseMultipleStreamsPerJobOk() (*bool, bool)`

GetUseMultipleStreamsPerJobOk returns a tuple with the UseMultipleStreamsPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMultipleStreamsPerJob

`func (o *GlobalNetworkTrafficRulesModel) SetUseMultipleStreamsPerJob(v bool)`

SetUseMultipleStreamsPerJob sets UseMultipleStreamsPerJob field to given value.


### GetUploadStreamsCount

`func (o *GlobalNetworkTrafficRulesModel) GetUploadStreamsCount() int32`

GetUploadStreamsCount returns the UploadStreamsCount field if non-nil, zero value otherwise.

### GetUploadStreamsCountOk

`func (o *GlobalNetworkTrafficRulesModel) GetUploadStreamsCountOk() (*int32, bool)`

GetUploadStreamsCountOk returns a tuple with the UploadStreamsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadStreamsCount

`func (o *GlobalNetworkTrafficRulesModel) SetUploadStreamsCount(v int32)`

SetUploadStreamsCount sets UploadStreamsCount field to given value.

### HasUploadStreamsCount

`func (o *GlobalNetworkTrafficRulesModel) HasUploadStreamsCount() bool`

HasUploadStreamsCount returns a boolean if a field has been set.

### GetTrafficRules

`func (o *GlobalNetworkTrafficRulesModel) GetTrafficRules() []TrafficRuleModel`

GetTrafficRules returns the TrafficRules field if non-nil, zero value otherwise.

### GetTrafficRulesOk

`func (o *GlobalNetworkTrafficRulesModel) GetTrafficRulesOk() (*[]TrafficRuleModel, bool)`

GetTrafficRulesOk returns a tuple with the TrafficRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRules

`func (o *GlobalNetworkTrafficRulesModel) SetTrafficRules(v []TrafficRuleModel)`

SetTrafficRules sets TrafficRules field to given value.

### HasTrafficRules

`func (o *GlobalNetworkTrafficRulesModel) HasTrafficRules() bool`

HasTrafficRules returns a boolean if a field has been set.

### GetPreferredNetworks

`func (o *GlobalNetworkTrafficRulesModel) GetPreferredNetworks() PreferredNetworksModel`

GetPreferredNetworks returns the PreferredNetworks field if non-nil, zero value otherwise.

### GetPreferredNetworksOk

`func (o *GlobalNetworkTrafficRulesModel) GetPreferredNetworksOk() (*PreferredNetworksModel, bool)`

GetPreferredNetworksOk returns a tuple with the PreferredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredNetworks

`func (o *GlobalNetworkTrafficRulesModel) SetPreferredNetworks(v PreferredNetworksModel)`

SetPreferredNetworks sets PreferredNetworks field to given value.

### HasPreferredNetworks

`func (o *GlobalNetworkTrafficRulesModel) HasPreferredNetworks() bool`

HasPreferredNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


