# RepositoryStateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the backup repository. | 
**Name** | **string** | Name of the backup repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Description** | **string** |  | 
**HostId** | Pointer to **string** | ID of the server that is used as a backup repository. | [optional] 
**HostName** | Pointer to **string** | Name of the server that is used as a backup repository. | [optional] 
**Path** | Pointer to **string** | Path to the folder where backup files are stored. | [optional] 
**CapacityGB** | **float64** | Repository capacity in GB. | 
**FreeGB** | **float64** | Repository free space in GB. | 
**UsedSpaceGB** | **float64** | Repository used space in GB. | 

## Methods

### NewRepositoryStateModel

`func NewRepositoryStateModel(id string, name string, type_ ERepositoryType, description string, capacityGB float64, freeGB float64, usedSpaceGB float64, ) *RepositoryStateModel`

NewRepositoryStateModel instantiates a new RepositoryStateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryStateModelWithDefaults

`func NewRepositoryStateModelWithDefaults() *RepositoryStateModel`

NewRepositoryStateModelWithDefaults instantiates a new RepositoryStateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepositoryStateModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepositoryStateModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepositoryStateModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RepositoryStateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryStateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryStateModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RepositoryStateModel) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryStateModel) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryStateModel) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RepositoryStateModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryStateModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryStateModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHostId

`func (o *RepositoryStateModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *RepositoryStateModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *RepositoryStateModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *RepositoryStateModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetHostName

`func (o *RepositoryStateModel) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *RepositoryStateModel) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *RepositoryStateModel) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *RepositoryStateModel) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPath

`func (o *RepositoryStateModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RepositoryStateModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RepositoryStateModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RepositoryStateModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCapacityGB

`func (o *RepositoryStateModel) GetCapacityGB() float64`

GetCapacityGB returns the CapacityGB field if non-nil, zero value otherwise.

### GetCapacityGBOk

`func (o *RepositoryStateModel) GetCapacityGBOk() (*float64, bool)`

GetCapacityGBOk returns a tuple with the CapacityGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityGB

`func (o *RepositoryStateModel) SetCapacityGB(v float64)`

SetCapacityGB sets CapacityGB field to given value.


### GetFreeGB

`func (o *RepositoryStateModel) GetFreeGB() float64`

GetFreeGB returns the FreeGB field if non-nil, zero value otherwise.

### GetFreeGBOk

`func (o *RepositoryStateModel) GetFreeGBOk() (*float64, bool)`

GetFreeGBOk returns a tuple with the FreeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeGB

`func (o *RepositoryStateModel) SetFreeGB(v float64)`

SetFreeGB sets FreeGB field to given value.


### GetUsedSpaceGB

`func (o *RepositoryStateModel) GetUsedSpaceGB() float64`

GetUsedSpaceGB returns the UsedSpaceGB field if non-nil, zero value otherwise.

### GetUsedSpaceGBOk

`func (o *RepositoryStateModel) GetUsedSpaceGBOk() (*float64, bool)`

GetUsedSpaceGBOk returns a tuple with the UsedSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceGB

`func (o *RepositoryStateModel) SetUsedSpaceGB(v float64)`

SetUsedSpaceGB sets UsedSpaceGB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


