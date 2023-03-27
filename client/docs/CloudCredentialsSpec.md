# CloudCredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 

## Methods

### NewCloudCredentialsSpec

`func NewCloudCredentialsSpec(type_ ECloudCredentialsType, ) *CloudCredentialsSpec`

NewCloudCredentialsSpec instantiates a new CloudCredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsSpecWithDefaults

`func NewCloudCredentialsSpecWithDefaults() *CloudCredentialsSpec`

NewCloudCredentialsSpecWithDefaults instantiates a new CloudCredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudCredentialsSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCredentialsSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCredentialsSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCredentialsSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CloudCredentialsSpec) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCredentialsSpec) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCredentialsSpec) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


