# AzureLinuxHelperApplianceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | ID that Veeam Backup &amp; Replication assigned to the Azure subscription. | 
**Location** | Pointer to **string** | Storage account location where you want to configure the helper appliance. | [optional] 
**StorageAccount** | Pointer to **string** | Name of the Azure storage account whose resources are used to store the helper appliance. | [optional] 
**ResourceGroup** | Pointer to **string** | Resource group associated with the helper appliance. | [optional] 
**VirtualNetwork** | Pointer to **string** | Network to which the helper appliance is connected. | [optional] 
**Subnet** | Pointer to **string** | Subnet for the helper appliance. | [optional] 
**SSHPort** | Pointer to **int32** | Port over which Veeam Backup &amp; Replication communicates with the helper appliance. | [optional] 

## Methods

### NewAzureLinuxHelperApplianceSpec

`func NewAzureLinuxHelperApplianceSpec(subscriptionId string, ) *AzureLinuxHelperApplianceSpec`

NewAzureLinuxHelperApplianceSpec instantiates a new AzureLinuxHelperApplianceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureLinuxHelperApplianceSpecWithDefaults

`func NewAzureLinuxHelperApplianceSpecWithDefaults() *AzureLinuxHelperApplianceSpec`

NewAzureLinuxHelperApplianceSpecWithDefaults instantiates a new AzureLinuxHelperApplianceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureLinuxHelperApplianceSpec) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureLinuxHelperApplianceSpec) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureLinuxHelperApplianceSpec) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetLocation

`func (o *AzureLinuxHelperApplianceSpec) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureLinuxHelperApplianceSpec) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureLinuxHelperApplianceSpec) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureLinuxHelperApplianceSpec) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStorageAccount

`func (o *AzureLinuxHelperApplianceSpec) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureLinuxHelperApplianceSpec) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureLinuxHelperApplianceSpec) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AzureLinuxHelperApplianceSpec) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AzureLinuxHelperApplianceSpec) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureLinuxHelperApplianceSpec) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureLinuxHelperApplianceSpec) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureLinuxHelperApplianceSpec) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *AzureLinuxHelperApplianceSpec) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *AzureLinuxHelperApplianceSpec) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *AzureLinuxHelperApplianceSpec) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *AzureLinuxHelperApplianceSpec) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *AzureLinuxHelperApplianceSpec) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AzureLinuxHelperApplianceSpec) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AzureLinuxHelperApplianceSpec) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AzureLinuxHelperApplianceSpec) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSSHPort

`func (o *AzureLinuxHelperApplianceSpec) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *AzureLinuxHelperApplianceSpec) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *AzureLinuxHelperApplianceSpec) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *AzureLinuxHelperApplianceSpec) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


