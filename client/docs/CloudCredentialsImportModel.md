# CloudCredentialsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User name, account name or access key. | 
**Tag** | **string** | Tag used to identify the cloud credentials record. | 
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 

## Methods

### NewCloudCredentialsImportModel

`func NewCloudCredentialsImportModel(name string, tag string, type_ ECloudCredentialsType, ) *CloudCredentialsImportModel`

NewCloudCredentialsImportModel instantiates a new CloudCredentialsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsImportModelWithDefaults

`func NewCloudCredentialsImportModelWithDefaults() *CloudCredentialsImportModel`

NewCloudCredentialsImportModelWithDefaults instantiates a new CloudCredentialsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudCredentialsImportModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentialsImportModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentialsImportModel) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *CloudCredentialsImportModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudCredentialsImportModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudCredentialsImportModel) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *CloudCredentialsImportModel) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCredentialsImportModel) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCredentialsImportModel) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


