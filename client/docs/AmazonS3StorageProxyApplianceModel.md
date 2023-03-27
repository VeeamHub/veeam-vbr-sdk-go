# AmazonS3StorageProxyApplianceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ec2InstanceType** | Pointer to **string** | EC2 instance type for the proxy appliance. | [optional] 
**Vpc** | Pointer to **string** | Amazon VPC where Veeam Backup &amp; Replication launches the target instance. | [optional] 
**Subnet** | Pointer to **string** | Subnet for the proxy appliance. | [optional] 
**SecurityGroup** | Pointer to **string** | Security group associated with the proxy appliance. | [optional] 
**RedirectorPort** | Pointer to **int32** | TCP port used to route requests between the proxy appliance and backup infrastructure components. | [optional] 

## Methods

### NewAmazonS3StorageProxyApplianceModel

`func NewAmazonS3StorageProxyApplianceModel() *AmazonS3StorageProxyApplianceModel`

NewAmazonS3StorageProxyApplianceModel instantiates a new AmazonS3StorageProxyApplianceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageProxyApplianceModelWithDefaults

`func NewAmazonS3StorageProxyApplianceModelWithDefaults() *AmazonS3StorageProxyApplianceModel`

NewAmazonS3StorageProxyApplianceModelWithDefaults instantiates a new AmazonS3StorageProxyApplianceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEc2InstanceType

`func (o *AmazonS3StorageProxyApplianceModel) GetEc2InstanceType() string`

GetEc2InstanceType returns the Ec2InstanceType field if non-nil, zero value otherwise.

### GetEc2InstanceTypeOk

`func (o *AmazonS3StorageProxyApplianceModel) GetEc2InstanceTypeOk() (*string, bool)`

GetEc2InstanceTypeOk returns a tuple with the Ec2InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2InstanceType

`func (o *AmazonS3StorageProxyApplianceModel) SetEc2InstanceType(v string)`

SetEc2InstanceType sets Ec2InstanceType field to given value.

### HasEc2InstanceType

`func (o *AmazonS3StorageProxyApplianceModel) HasEc2InstanceType() bool`

HasEc2InstanceType returns a boolean if a field has been set.

### GetVpc

`func (o *AmazonS3StorageProxyApplianceModel) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AmazonS3StorageProxyApplianceModel) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AmazonS3StorageProxyApplianceModel) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *AmazonS3StorageProxyApplianceModel) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetSubnet

`func (o *AmazonS3StorageProxyApplianceModel) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AmazonS3StorageProxyApplianceModel) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AmazonS3StorageProxyApplianceModel) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AmazonS3StorageProxyApplianceModel) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *AmazonS3StorageProxyApplianceModel) GetSecurityGroup() string`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *AmazonS3StorageProxyApplianceModel) GetSecurityGroupOk() (*string, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *AmazonS3StorageProxyApplianceModel) SetSecurityGroup(v string)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *AmazonS3StorageProxyApplianceModel) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetRedirectorPort

`func (o *AmazonS3StorageProxyApplianceModel) GetRedirectorPort() int32`

GetRedirectorPort returns the RedirectorPort field if non-nil, zero value otherwise.

### GetRedirectorPortOk

`func (o *AmazonS3StorageProxyApplianceModel) GetRedirectorPortOk() (*int32, bool)`

GetRedirectorPortOk returns a tuple with the RedirectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectorPort

`func (o *AmazonS3StorageProxyApplianceModel) SetRedirectorPort(v int32)`

SetRedirectorPort sets RedirectorPort field to given value.

### HasRedirectorPort

`func (o *AmazonS3StorageProxyApplianceModel) HasRedirectorPort() bool`

HasRedirectorPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


