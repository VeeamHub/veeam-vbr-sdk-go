# AzureLinuxHelperApplianceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Helper appliance ID. | 
**SubscriptionId** | **string** | ID that Veeam Backup &amp; Replication assigned to the Azure subscription. | 
**VmName** | Pointer to **string** | VM name of the helper appliance. | [optional] 
**Location** | Pointer to **string** | Storage account location where the helper appliance is configured. | [optional] 
**StorageAccount** | Pointer to **string** | Azure storage account whose resources are used to store disks of the helper appliance. | [optional] 
**ResourceGroup** | Pointer to **string** | Resource group associated with the helper appliance. | [optional] 
**VirtualNetwork** | Pointer to **string** | Network to which the helper appliance is connected. | [optional] 
**Subnet** | Pointer to **string** | Subnet for the helper appliance. | [optional] 
**SSHPort** | Pointer to **int32** | Port over which Veeam Backup &amp; Replication communicates with the helper appliance. | [optional] 

## Methods

### NewAzureLinuxHelperApplianceModel

`func NewAzureLinuxHelperApplianceModel(id string, subscriptionId string, ) *AzureLinuxHelperApplianceModel`

NewAzureLinuxHelperApplianceModel instantiates a new AzureLinuxHelperApplianceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureLinuxHelperApplianceModelWithDefaults

`func NewAzureLinuxHelperApplianceModelWithDefaults() *AzureLinuxHelperApplianceModel`

NewAzureLinuxHelperApplianceModelWithDefaults instantiates a new AzureLinuxHelperApplianceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureLinuxHelperApplianceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureLinuxHelperApplianceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureLinuxHelperApplianceModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubscriptionId

`func (o *AzureLinuxHelperApplianceModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureLinuxHelperApplianceModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureLinuxHelperApplianceModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetVmName

`func (o *AzureLinuxHelperApplianceModel) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *AzureLinuxHelperApplianceModel) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *AzureLinuxHelperApplianceModel) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *AzureLinuxHelperApplianceModel) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetLocation

`func (o *AzureLinuxHelperApplianceModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureLinuxHelperApplianceModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureLinuxHelperApplianceModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureLinuxHelperApplianceModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStorageAccount

`func (o *AzureLinuxHelperApplianceModel) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureLinuxHelperApplianceModel) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureLinuxHelperApplianceModel) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AzureLinuxHelperApplianceModel) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AzureLinuxHelperApplianceModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureLinuxHelperApplianceModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureLinuxHelperApplianceModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureLinuxHelperApplianceModel) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *AzureLinuxHelperApplianceModel) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *AzureLinuxHelperApplianceModel) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *AzureLinuxHelperApplianceModel) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *AzureLinuxHelperApplianceModel) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *AzureLinuxHelperApplianceModel) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AzureLinuxHelperApplianceModel) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AzureLinuxHelperApplianceModel) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AzureLinuxHelperApplianceModel) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSSHPort

`func (o *AzureLinuxHelperApplianceModel) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *AzureLinuxHelperApplianceModel) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *AzureLinuxHelperApplianceModel) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *AzureLinuxHelperApplianceModel) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


