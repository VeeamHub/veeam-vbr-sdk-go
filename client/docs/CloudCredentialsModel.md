# CloudCredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the cloud credentials record. | 
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 

## Methods

### NewCloudCredentialsModel

`func NewCloudCredentialsModel(id string, type_ ECloudCredentialsType, ) *CloudCredentialsModel`

NewCloudCredentialsModel instantiates a new CloudCredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsModelWithDefaults

`func NewCloudCredentialsModelWithDefaults() *CloudCredentialsModel`

NewCloudCredentialsModelWithDefaults instantiates a new CloudCredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudCredentialsModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCredentialsModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCredentialsModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *CloudCredentialsModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCredentialsModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCredentialsModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCredentialsModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CloudCredentialsModel) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCredentialsModel) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCredentialsModel) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


