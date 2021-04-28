# CredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | User name. | 
**Password** | Pointer to **string** | Password. | [optional] 
**Description** | Pointer to **string** | Description of the credentials record. | [optional] 
**Type** | [**ECredentialsType**](ECredentialsType.md) |  | 

## Methods

### NewCredentialsSpec

`func NewCredentialsSpec(username string, type_ ECredentialsType, ) *CredentialsSpec`

NewCredentialsSpec instantiates a new CredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsSpecWithDefaults

`func NewCredentialsSpecWithDefaults() *CredentialsSpec`

NewCredentialsSpecWithDefaults instantiates a new CredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CredentialsSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialsSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialsSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialsSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialsSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialsSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialsSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDescription

`func (o *CredentialsSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialsSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CredentialsSpec) GetType() ECredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsSpec) GetTypeOk() (*ECredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsSpec) SetType(v ECredentialsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


