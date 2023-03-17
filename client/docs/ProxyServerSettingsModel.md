# ProxyServerSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of the server. | 
**TransportMode** | Pointer to [**EBackupProxyTransportMode**](EBackupProxyTransportMode.md) |  | [optional] 
**FailoverToNetwork** | Pointer to **bool** | (For the Direct storage access and Virtual appliance transport modes) If *true*, Veeam Backup &amp; Replication failovers to the network transport mode in case the primary mode fails or is unavailable. | [optional] 
**HostToProxyEncryption** | Pointer to **bool** | (For the Network mode) If *true*, VM data is transferred over an encrypted TLS connection. | [optional] 
**ConnectedDatastores** | Pointer to [**ProxyDatastoreSettingsModel**](ProxyDatastoreSettingsModel.md) |  | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 

## Methods

### NewProxyServerSettingsModel

`func NewProxyServerSettingsModel(hostId string, ) *ProxyServerSettingsModel`

NewProxyServerSettingsModel instantiates a new ProxyServerSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyServerSettingsModelWithDefaults

`func NewProxyServerSettingsModelWithDefaults() *ProxyServerSettingsModel`

NewProxyServerSettingsModelWithDefaults instantiates a new ProxyServerSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *ProxyServerSettingsModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ProxyServerSettingsModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ProxyServerSettingsModel) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetTransportMode

`func (o *ProxyServerSettingsModel) GetTransportMode() EBackupProxyTransportMode`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *ProxyServerSettingsModel) GetTransportModeOk() (*EBackupProxyTransportMode, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *ProxyServerSettingsModel) SetTransportMode(v EBackupProxyTransportMode)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *ProxyServerSettingsModel) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetFailoverToNetwork

`func (o *ProxyServerSettingsModel) GetFailoverToNetwork() bool`

GetFailoverToNetwork returns the FailoverToNetwork field if non-nil, zero value otherwise.

### GetFailoverToNetworkOk

`func (o *ProxyServerSettingsModel) GetFailoverToNetworkOk() (*bool, bool)`

GetFailoverToNetworkOk returns a tuple with the FailoverToNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverToNetwork

`func (o *ProxyServerSettingsModel) SetFailoverToNetwork(v bool)`

SetFailoverToNetwork sets FailoverToNetwork field to given value.

### HasFailoverToNetwork

`func (o *ProxyServerSettingsModel) HasFailoverToNetwork() bool`

HasFailoverToNetwork returns a boolean if a field has been set.

### GetHostToProxyEncryption

`func (o *ProxyServerSettingsModel) GetHostToProxyEncryption() bool`

GetHostToProxyEncryption returns the HostToProxyEncryption field if non-nil, zero value otherwise.

### GetHostToProxyEncryptionOk

`func (o *ProxyServerSettingsModel) GetHostToProxyEncryptionOk() (*bool, bool)`

GetHostToProxyEncryptionOk returns a tuple with the HostToProxyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostToProxyEncryption

`func (o *ProxyServerSettingsModel) SetHostToProxyEncryption(v bool)`

SetHostToProxyEncryption sets HostToProxyEncryption field to given value.

### HasHostToProxyEncryption

`func (o *ProxyServerSettingsModel) HasHostToProxyEncryption() bool`

HasHostToProxyEncryption returns a boolean if a field has been set.

### GetConnectedDatastores

`func (o *ProxyServerSettingsModel) GetConnectedDatastores() ProxyDatastoreSettingsModel`

GetConnectedDatastores returns the ConnectedDatastores field if non-nil, zero value otherwise.

### GetConnectedDatastoresOk

`func (o *ProxyServerSettingsModel) GetConnectedDatastoresOk() (*ProxyDatastoreSettingsModel, bool)`

GetConnectedDatastoresOk returns a tuple with the ConnectedDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedDatastores

`func (o *ProxyServerSettingsModel) SetConnectedDatastores(v ProxyDatastoreSettingsModel)`

SetConnectedDatastores sets ConnectedDatastores field to given value.

### HasConnectedDatastores

`func (o *ProxyServerSettingsModel) HasConnectedDatastores() bool`

HasConnectedDatastores returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *ProxyServerSettingsModel) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *ProxyServerSettingsModel) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *ProxyServerSettingsModel) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *ProxyServerSettingsModel) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


