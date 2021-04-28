# CredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the credentials record. | 
**Username** | **string** | User name. | 
**Description** | **string** | Description of the credentials record. | 
**Type** | [**ECredentialsType**](ECredentialsType.md) |  | 
**CreationTime** | **time.Time** | Date and time when the credentials were created. | 

## Methods

### NewCredentialsModel

`func NewCredentialsModel(id string, username string, description string, type_ ECredentialsType, creationTime time.Time, ) *CredentialsModel`

NewCredentialsModel instantiates a new CredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsModelWithDefaults

`func NewCredentialsModelWithDefaults() *CredentialsModel`

NewCredentialsModelWithDefaults instantiates a new CredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialsModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialsModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialsModel) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *CredentialsModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialsModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialsModel) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *CredentialsModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *CredentialsModel) GetType() ECredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsModel) GetTypeOk() (*ECredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsModel) SetType(v ECredentialsType)`

SetType sets Type field to given value.


### GetCreationTime

`func (o *CredentialsModel) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CredentialsModel) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CredentialsModel) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


