# PerformanceExtentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the backup repository added as a performance extent. | 
**Name** | **string** | Name of the backup repository added as a performance extent. | 
**Status** | Pointer to [**ERepositoryExtentStatusType**](ERepositoryExtentStatusType.md) |  | [optional] 

## Methods

### NewPerformanceExtentModel

`func NewPerformanceExtentModel(id string, name string, ) *PerformanceExtentModel`

NewPerformanceExtentModel instantiates a new PerformanceExtentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceExtentModelWithDefaults

`func NewPerformanceExtentModelWithDefaults() *PerformanceExtentModel`

NewPerformanceExtentModelWithDefaults instantiates a new PerformanceExtentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PerformanceExtentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerformanceExtentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerformanceExtentModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PerformanceExtentModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PerformanceExtentModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PerformanceExtentModel) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PerformanceExtentModel) GetStatus() ERepositoryExtentStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PerformanceExtentModel) GetStatusOk() (*ERepositoryExtentStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PerformanceExtentModel) SetStatus(v ERepositoryExtentStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PerformanceExtentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


