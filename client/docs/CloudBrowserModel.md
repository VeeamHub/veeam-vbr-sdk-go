# CloudBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of the cloud credentials record. | 
**ServiceType** | [**ECloudServiceType**](ECloudServiceType.md) |  | 

## Methods

### NewCloudBrowserModel

`func NewCloudBrowserModel(credentialsId string, serviceType ECloudServiceType, ) *CloudBrowserModel`

NewCloudBrowserModel instantiates a new CloudBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBrowserModelWithDefaults

`func NewCloudBrowserModelWithDefaults() *CloudBrowserModel`

NewCloudBrowserModelWithDefaults instantiates a new CloudBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *CloudBrowserModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *CloudBrowserModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *CloudBrowserModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetServiceType

`func (o *CloudBrowserModel) GetServiceType() ECloudServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CloudBrowserModel) GetServiceTypeOk() (*ECloudServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CloudBrowserModel) SetServiceType(v ECloudServiceType)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


