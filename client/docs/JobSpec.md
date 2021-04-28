# JobSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the job. | 
**Description** | **string** | Job description specified at the time of the job creation. | 
**Type** | [**EJobType**](EJobType.md) |  | 

## Methods

### NewJobSpec

`func NewJobSpec(name string, description string, type_ EJobType, ) *JobSpec`

NewJobSpec instantiates a new JobSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSpecWithDefaults

`func NewJobSpecWithDefaults() *JobSpec`

NewJobSpecWithDefaults instantiates a new JobSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *JobSpec) GetType() EJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobSpec) GetTypeOk() (*EJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobSpec) SetType(v EJobType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


