# JobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the job. | 
**Name** | **string** | Name of the job. | 
**Description** | **string** | Description of the job. | 
**Type** | [**EJobType**](EJobType.md) |  | 
**IsDisabled** | **bool** | If *true*, the job is disabled. | 

## Methods

### NewJobModel

`func NewJobModel(id string, name string, description string, type_ EJobType, isDisabled bool, ) *JobModel`

NewJobModel instantiates a new JobModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobModelWithDefaults

`func NewJobModelWithDefaults() *JobModel`

NewJobModelWithDefaults instantiates a new JobModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JobModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *JobModel) GetType() EJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobModel) GetTypeOk() (*EJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobModel) SetType(v EJobType)`

SetType sets Type field to given value.


### GetIsDisabled

`func (o *JobModel) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *JobModel) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *JobModel) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


