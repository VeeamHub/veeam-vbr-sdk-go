# RestoreProxySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelection** | **bool** | If *true*, Veeam Backup &amp; Replication detects backup proxies that are connected to the source datastore and automatically assigns optimal proxy resources for processing VM data. | 
**ProxyIds** | Pointer to **[]string** | Array of backup proxy IDs. | [optional] 

## Methods

### NewRestoreProxySpec

`func NewRestoreProxySpec(autoSelection bool, ) *RestoreProxySpec`

NewRestoreProxySpec instantiates a new RestoreProxySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreProxySpecWithDefaults

`func NewRestoreProxySpecWithDefaults() *RestoreProxySpec`

NewRestoreProxySpecWithDefaults instantiates a new RestoreProxySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelection

`func (o *RestoreProxySpec) GetAutoSelection() bool`

GetAutoSelection returns the AutoSelection field if non-nil, zero value otherwise.

### GetAutoSelectionOk

`func (o *RestoreProxySpec) GetAutoSelectionOk() (*bool, bool)`

GetAutoSelectionOk returns a tuple with the AutoSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelection

`func (o *RestoreProxySpec) SetAutoSelection(v bool)`

SetAutoSelection sets AutoSelection field to given value.


### GetProxyIds

`func (o *RestoreProxySpec) GetProxyIds() []string`

GetProxyIds returns the ProxyIds field if non-nil, zero value otherwise.

### GetProxyIdsOk

`func (o *RestoreProxySpec) GetProxyIdsOk() (*[]string, bool)`

GetProxyIdsOk returns a tuple with the ProxyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyIds

`func (o *RestoreProxySpec) SetProxyIds(v []string)`

SetProxyIds sets ProxyIds field to given value.

### HasProxyIds

`func (o *RestoreProxySpec) HasProxyIds() bool`

HasProxyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


