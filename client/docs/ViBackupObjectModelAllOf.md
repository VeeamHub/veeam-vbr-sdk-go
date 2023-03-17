# ViBackupObjectModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **string** | ID of the virtual infrastructure object (mo-ref or ID, depending on the virtualization platform). | 
**ViType** | Pointer to [**EVmwareInventoryType**](EVmwareInventoryType.md) |  | [optional] 
**Path** | Pointer to **string** | Path to the object. | [optional] 

## Methods

### NewViBackupObjectModelAllOf

`func NewViBackupObjectModelAllOf(objectId string, ) *ViBackupObjectModelAllOf`

NewViBackupObjectModelAllOf instantiates a new ViBackupObjectModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViBackupObjectModelAllOfWithDefaults

`func NewViBackupObjectModelAllOfWithDefaults() *ViBackupObjectModelAllOf`

NewViBackupObjectModelAllOfWithDefaults instantiates a new ViBackupObjectModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *ViBackupObjectModelAllOf) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ViBackupObjectModelAllOf) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ViBackupObjectModelAllOf) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetViType

`func (o *ViBackupObjectModelAllOf) GetViType() EVmwareInventoryType`

GetViType returns the ViType field if non-nil, zero value otherwise.

### GetViTypeOk

`func (o *ViBackupObjectModelAllOf) GetViTypeOk() (*EVmwareInventoryType, bool)`

GetViTypeOk returns a tuple with the ViType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViType

`func (o *ViBackupObjectModelAllOf) SetViType(v EVmwareInventoryType)`

SetViType sets ViType field to given value.

### HasViType

`func (o *ViBackupObjectModelAllOf) HasViType() bool`

HasViType returns a boolean if a field has been set.

### GetPath

`func (o *ViBackupObjectModelAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ViBackupObjectModelAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ViBackupObjectModelAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ViBackupObjectModelAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


