# AzureStorageProxyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | ID that Veeam Backup &amp; Replication assigned to the Azure subscription. | 
**InstanceSize** | Pointer to **string** | Size of the appliance. | [optional] 
**ResourceGroup** | Pointer to **string** | Resource group associated with the proxy appliance. | [optional] 
**VirtualNetwork** | Pointer to **string** | Network to which the helper appliance is connected. | [optional] 
**Subnet** | Pointer to **string** | Subnet for the proxy appliance. | [optional] 
**RedirectorPort** | Pointer to **int32** | TCP port used to route requests between the proxy appliance and backup infrastructure components. | [optional] 

## Methods

### NewAzureStorageProxyModel

`func NewAzureStorageProxyModel(subscriptionId string, ) *AzureStorageProxyModel`

NewAzureStorageProxyModel instantiates a new AzureStorageProxyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageProxyModelWithDefaults

`func NewAzureStorageProxyModelWithDefaults() *AzureStorageProxyModel`

NewAzureStorageProxyModelWithDefaults instantiates a new AzureStorageProxyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureStorageProxyModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureStorageProxyModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureStorageProxyModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetInstanceSize

`func (o *AzureStorageProxyModel) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *AzureStorageProxyModel) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *AzureStorageProxyModel) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *AzureStorageProxyModel) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AzureStorageProxyModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureStorageProxyModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureStorageProxyModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureStorageProxyModel) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *AzureStorageProxyModel) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *AzureStorageProxyModel) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *AzureStorageProxyModel) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *AzureStorageProxyModel) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *AzureStorageProxyModel) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AzureStorageProxyModel) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AzureStorageProxyModel) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AzureStorageProxyModel) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetRedirectorPort

`func (o *AzureStorageProxyModel) GetRedirectorPort() int32`

GetRedirectorPort returns the RedirectorPort field if non-nil, zero value otherwise.

### GetRedirectorPortOk

`func (o *AzureStorageProxyModel) GetRedirectorPortOk() (*int32, bool)`

GetRedirectorPortOk returns a tuple with the RedirectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectorPort

`func (o *AzureStorageProxyModel) SetRedirectorPort(v int32)`

SetRedirectorPort sets RedirectorPort field to given value.

### HasRedirectorPort

`func (o *AzureStorageProxyModel) HasRedirectorPort() bool`

HasRedirectorPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


