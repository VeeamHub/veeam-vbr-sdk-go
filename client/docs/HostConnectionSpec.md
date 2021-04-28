# HostConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Full DNS name or IP address of the server. | 
**CredentialsId** | **string** | ID of a credentials record used to connect to the server. | 
**Type** | [**EManagedServerType**](EManagedServerType.md) |  | 
**Port** | Pointer to **int32** | Port used to communicate with the server. | [optional] 

## Methods

### NewHostConnectionSpec

`func NewHostConnectionSpec(serverName string, credentialsId string, type_ EManagedServerType, ) *HostConnectionSpec`

NewHostConnectionSpec instantiates a new HostConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostConnectionSpecWithDefaults

`func NewHostConnectionSpecWithDefaults() *HostConnectionSpec`

NewHostConnectionSpecWithDefaults instantiates a new HostConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *HostConnectionSpec) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *HostConnectionSpec) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *HostConnectionSpec) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetCredentialsId

`func (o *HostConnectionSpec) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *HostConnectionSpec) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *HostConnectionSpec) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetType

`func (o *HostConnectionSpec) GetType() EManagedServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostConnectionSpec) GetTypeOk() (*EManagedServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostConnectionSpec) SetType(v EManagedServerType)`

SetType sets Type field to given value.


### GetPort

`func (o *HostConnectionSpec) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HostConnectionSpec) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HostConnectionSpec) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HostConnectionSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


