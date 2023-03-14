# CloudBrowserNewFolderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of a cloud credentials record requiered to connect to the object storage. | 
**ServiceType** | [**ECloudServiceType**](ECloudServiceType.md) |  | 
**NewFolderName** | **string** | Name of the new folder. | 

## Methods

### NewCloudBrowserNewFolderSpec

`func NewCloudBrowserNewFolderSpec(credentialsId string, serviceType ECloudServiceType, newFolderName string, ) *CloudBrowserNewFolderSpec`

NewCloudBrowserNewFolderSpec instantiates a new CloudBrowserNewFolderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBrowserNewFolderSpecWithDefaults

`func NewCloudBrowserNewFolderSpecWithDefaults() *CloudBrowserNewFolderSpec`

NewCloudBrowserNewFolderSpecWithDefaults instantiates a new CloudBrowserNewFolderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *CloudBrowserNewFolderSpec) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *CloudBrowserNewFolderSpec) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *CloudBrowserNewFolderSpec) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetServiceType

`func (o *CloudBrowserNewFolderSpec) GetServiceType() ECloudServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CloudBrowserNewFolderSpec) GetServiceTypeOk() (*ECloudServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CloudBrowserNewFolderSpec) SetServiceType(v ECloudServiceType)`

SetServiceType sets ServiceType field to given value.


### GetNewFolderName

`func (o *CloudBrowserNewFolderSpec) GetNewFolderName() string`

GetNewFolderName returns the NewFolderName field if non-nil, zero value otherwise.

### GetNewFolderNameOk

`func (o *CloudBrowserNewFolderSpec) GetNewFolderNameOk() (*string, bool)`

GetNewFolderNameOk returns a tuple with the NewFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFolderName

`func (o *CloudBrowserNewFolderSpec) SetNewFolderName(v string)`

SetNewFolderName sets NewFolderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


