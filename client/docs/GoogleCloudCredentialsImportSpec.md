# GoogleCloudCredentialsImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Tag** | **string** | Tag used to identify the cloud credentials record. | 
**AccessKey** | **string** | Access ID of a Google HMAC key. | 
**SecretKey** | **string** | Secret linked to the access ID. | 

## Methods

### NewGoogleCloudCredentialsImportSpec

`func NewGoogleCloudCredentialsImportSpec(type_ ECloudCredentialsType, tag string, accessKey string, secretKey string, ) *GoogleCloudCredentialsImportSpec`

NewGoogleCloudCredentialsImportSpec instantiates a new GoogleCloudCredentialsImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudCredentialsImportSpecWithDefaults

`func NewGoogleCloudCredentialsImportSpecWithDefaults() *GoogleCloudCredentialsImportSpec`

NewGoogleCloudCredentialsImportSpecWithDefaults instantiates a new GoogleCloudCredentialsImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GoogleCloudCredentialsImportSpec) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoogleCloudCredentialsImportSpec) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoogleCloudCredentialsImportSpec) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.


### GetDescription

`func (o *GoogleCloudCredentialsImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoogleCloudCredentialsImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoogleCloudCredentialsImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GoogleCloudCredentialsImportSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTag

`func (o *GoogleCloudCredentialsImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GoogleCloudCredentialsImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GoogleCloudCredentialsImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetAccessKey

`func (o *GoogleCloudCredentialsImportSpec) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GoogleCloudCredentialsImportSpec) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GoogleCloudCredentialsImportSpec) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *GoogleCloudCredentialsImportSpec) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *GoogleCloudCredentialsImportSpec) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *GoogleCloudCredentialsImportSpec) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


