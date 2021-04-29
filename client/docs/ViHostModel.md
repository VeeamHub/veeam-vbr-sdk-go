# ViHostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViHostType** | Pointer to [**EViHostType**](EViHostType.md) |  | [optional] 
**Port** | **int32** | Port used to communicate with the server. | 

## Methods

### NewViHostModel

`func NewViHostModel(port int32, ) *ViHostModel`

NewViHostModel instantiates a new ViHostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViHostModelWithDefaults

`func NewViHostModelWithDefaults() *ViHostModel`

NewViHostModelWithDefaults instantiates a new ViHostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViHostType

`func (o *ViHostModel) GetViHostType() EViHostType`

GetViHostType returns the ViHostType field if non-nil, zero value otherwise.

### GetViHostTypeOk

`func (o *ViHostModel) GetViHostTypeOk() (*EViHostType, bool)`

GetViHostTypeOk returns a tuple with the ViHostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViHostType

`func (o *ViHostModel) SetViHostType(v EViHostType)`

SetViHostType sets ViHostType field to given value.

### HasViHostType

`func (o *ViHostModel) HasViHostType() bool`

HasViHostType returns a boolean if a field has been set.

### GetPort

`func (o *ViHostModel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViHostModel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViHostModel) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


