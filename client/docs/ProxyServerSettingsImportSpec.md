# ProxyServerSettingsImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** | Name of the server. | 
**HostTag** | Pointer to **string** | Tag assigned to the server. | [optional] 
**TransportMode** | Pointer to [**EBackupProxyTransportMode**](EBackupProxyTransportMode.md) |  | [optional] 
**FailoverToNetwork** | Pointer to **bool** |  | [optional] 
**HostToProxyEncryption** | Pointer to **bool** | [For the Network mode] If *true*, VM data is transferred over an encrypted TLS connection.  | [optional] 
**ConnectedDatastores** | Pointer to [**ProxyDatastoreSettingsModel**](ProxyDatastoreSettingsModel.md) |  | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 

## Methods

### NewProxyServerSettingsImportSpec

`func NewProxyServerSettingsImportSpec(hostName string, ) *ProxyServerSettingsImportSpec`

NewProxyServerSettingsImportSpec instantiates a new ProxyServerSettingsImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyServerSettingsImportSpecWithDefaults

`func NewProxyServerSettingsImportSpecWithDefaults() *ProxyServerSettingsImportSpec`

NewProxyServerSettingsImportSpecWithDefaults instantiates a new ProxyServerSettingsImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *ProxyServerSettingsImportSpec) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ProxyServerSettingsImportSpec) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ProxyServerSettingsImportSpec) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetHostTag

`func (o *ProxyServerSettingsImportSpec) GetHostTag() string`

GetHostTag returns the HostTag field if non-nil, zero value otherwise.

### GetHostTagOk

`func (o *ProxyServerSettingsImportSpec) GetHostTagOk() (*string, bool)`

GetHostTagOk returns a tuple with the HostTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTag

`func (o *ProxyServerSettingsImportSpec) SetHostTag(v string)`

SetHostTag sets HostTag field to given value.

### HasHostTag

`func (o *ProxyServerSettingsImportSpec) HasHostTag() bool`

HasHostTag returns a boolean if a field has been set.

### GetTransportMode

`func (o *ProxyServerSettingsImportSpec) GetTransportMode() EBackupProxyTransportMode`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *ProxyServerSettingsImportSpec) GetTransportModeOk() (*EBackupProxyTransportMode, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *ProxyServerSettingsImportSpec) SetTransportMode(v EBackupProxyTransportMode)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *ProxyServerSettingsImportSpec) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetFailoverToNetwork

`func (o *ProxyServerSettingsImportSpec) GetFailoverToNetwork() bool`

GetFailoverToNetwork returns the FailoverToNetwork field if non-nil, zero value otherwise.

### GetFailoverToNetworkOk

`func (o *ProxyServerSettingsImportSpec) GetFailoverToNetworkOk() (*bool, bool)`

GetFailoverToNetworkOk returns a tuple with the FailoverToNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverToNetwork

`func (o *ProxyServerSettingsImportSpec) SetFailoverToNetwork(v bool)`

SetFailoverToNetwork sets FailoverToNetwork field to given value.

### HasFailoverToNetwork

`func (o *ProxyServerSettingsImportSpec) HasFailoverToNetwork() bool`

HasFailoverToNetwork returns a boolean if a field has been set.

### GetHostToProxyEncryption

`func (o *ProxyServerSettingsImportSpec) GetHostToProxyEncryption() bool`

GetHostToProxyEncryption returns the HostToProxyEncryption field if non-nil, zero value otherwise.

### GetHostToProxyEncryptionOk

`func (o *ProxyServerSettingsImportSpec) GetHostToProxyEncryptionOk() (*bool, bool)`

GetHostToProxyEncryptionOk returns a tuple with the HostToProxyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostToProxyEncryption

`func (o *ProxyServerSettingsImportSpec) SetHostToProxyEncryption(v bool)`

SetHostToProxyEncryption sets HostToProxyEncryption field to given value.

### HasHostToProxyEncryption

`func (o *ProxyServerSettingsImportSpec) HasHostToProxyEncryption() bool`

HasHostToProxyEncryption returns a boolean if a field has been set.

### GetConnectedDatastores

`func (o *ProxyServerSettingsImportSpec) GetConnectedDatastores() ProxyDatastoreSettingsModel`

GetConnectedDatastores returns the ConnectedDatastores field if non-nil, zero value otherwise.

### GetConnectedDatastoresOk

`func (o *ProxyServerSettingsImportSpec) GetConnectedDatastoresOk() (*ProxyDatastoreSettingsModel, bool)`

GetConnectedDatastoresOk returns a tuple with the ConnectedDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedDatastores

`func (o *ProxyServerSettingsImportSpec) SetConnectedDatastores(v ProxyDatastoreSettingsModel)`

SetConnectedDatastores sets ConnectedDatastores field to given value.

### HasConnectedDatastores

`func (o *ProxyServerSettingsImportSpec) HasConnectedDatastores() bool`

HasConnectedDatastores returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *ProxyServerSettingsImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *ProxyServerSettingsImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *ProxyServerSettingsImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *ProxyServerSettingsImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


