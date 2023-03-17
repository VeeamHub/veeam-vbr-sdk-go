# AmazonCloudCredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | ID of the access key. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewAmazonCloudCredentialsModel

`func NewAmazonCloudCredentialsModel(accessKey string, ) *AmazonCloudCredentialsModel`

NewAmazonCloudCredentialsModel instantiates a new AmazonCloudCredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonCloudCredentialsModelWithDefaults

`func NewAmazonCloudCredentialsModelWithDefaults() *AmazonCloudCredentialsModel`

NewAmazonCloudCredentialsModelWithDefaults instantiates a new AmazonCloudCredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AmazonCloudCredentialsModel) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AmazonCloudCredentialsModel) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AmazonCloudCredentialsModel) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetTag

`func (o *AmazonCloudCredentialsModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AmazonCloudCredentialsModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AmazonCloudCredentialsModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AmazonCloudCredentialsModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


