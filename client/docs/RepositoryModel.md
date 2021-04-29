# RepositoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the backup repository. | 
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | Pointer to **string** | VMware vSphere tag assigned to the backup repository. | [optional] 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 

## Methods

### NewRepositoryModel

`func NewRepositoryModel(id string, name string, description string, type_ ERepositoryType, ) *RepositoryModel`

NewRepositoryModel instantiates a new RepositoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryModelWithDefaults

`func NewRepositoryModelWithDefaults() *RepositoryModel`

NewRepositoryModelWithDefaults instantiates a new RepositoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepositoryModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepositoryModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepositoryModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RepositoryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RepositoryModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *RepositoryModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RepositoryModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RepositoryModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RepositoryModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *RepositoryModel) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryModel) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryModel) SetType(v ERepositoryType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


