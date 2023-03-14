# GoogleCloudCredentialsModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access ID of a Google HMAC key. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewGoogleCloudCredentialsModelAllOf

`func NewGoogleCloudCredentialsModelAllOf(accessKey string, ) *GoogleCloudCredentialsModelAllOf`

NewGoogleCloudCredentialsModelAllOf instantiates a new GoogleCloudCredentialsModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudCredentialsModelAllOfWithDefaults

`func NewGoogleCloudCredentialsModelAllOfWithDefaults() *GoogleCloudCredentialsModelAllOf`

NewGoogleCloudCredentialsModelAllOfWithDefaults instantiates a new GoogleCloudCredentialsModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *GoogleCloudCredentialsModelAllOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GoogleCloudCredentialsModelAllOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GoogleCloudCredentialsModelAllOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetTag

`func (o *GoogleCloudCredentialsModelAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GoogleCloudCredentialsModelAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GoogleCloudCredentialsModelAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GoogleCloudCredentialsModelAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


