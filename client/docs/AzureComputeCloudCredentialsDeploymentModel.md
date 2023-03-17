# AzureComputeCloudCredentialsDeploymentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | [**EAzureComputeCredentialsDeploymentType**](EAzureComputeCredentialsDeploymentType.md) |  | 
**Region** | Pointer to [**EAzureRegionType**](EAzureRegionType.md) |  | [optional] 

## Methods

### NewAzureComputeCloudCredentialsDeploymentModel

`func NewAzureComputeCloudCredentialsDeploymentModel(deploymentType EAzureComputeCredentialsDeploymentType, ) *AzureComputeCloudCredentialsDeploymentModel`

NewAzureComputeCloudCredentialsDeploymentModel instantiates a new AzureComputeCloudCredentialsDeploymentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsDeploymentModelWithDefaults

`func NewAzureComputeCloudCredentialsDeploymentModelWithDefaults() *AzureComputeCloudCredentialsDeploymentModel`

NewAzureComputeCloudCredentialsDeploymentModelWithDefaults instantiates a new AzureComputeCloudCredentialsDeploymentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentType

`func (o *AzureComputeCloudCredentialsDeploymentModel) GetDeploymentType() EAzureComputeCredentialsDeploymentType`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *AzureComputeCloudCredentialsDeploymentModel) GetDeploymentTypeOk() (*EAzureComputeCredentialsDeploymentType, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *AzureComputeCloudCredentialsDeploymentModel) SetDeploymentType(v EAzureComputeCredentialsDeploymentType)`

SetDeploymentType sets DeploymentType field to given value.


### GetRegion

`func (o *AzureComputeCloudCredentialsDeploymentModel) GetRegion() EAzureRegionType`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureComputeCloudCredentialsDeploymentModel) GetRegionOk() (*EAzureRegionType, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureComputeCloudCredentialsDeploymentModel) SetRegion(v EAzureRegionType)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AzureComputeCloudCredentialsDeploymentModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


