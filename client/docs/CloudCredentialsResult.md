# CloudCredentialsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CloudCredentialsModel**](CloudCredentialsModel.md) | Array of cloud credentials. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewCloudCredentialsResult

`func NewCloudCredentialsResult(data []CloudCredentialsModel, pagination PaginationResult, ) *CloudCredentialsResult`

NewCloudCredentialsResult instantiates a new CloudCredentialsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsResultWithDefaults

`func NewCloudCredentialsResultWithDefaults() *CloudCredentialsResult`

NewCloudCredentialsResultWithDefaults instantiates a new CloudCredentialsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudCredentialsResult) GetData() []CloudCredentialsModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudCredentialsResult) GetDataOk() (*[]CloudCredentialsModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudCredentialsResult) SetData(v []CloudCredentialsModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *CloudCredentialsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CloudCredentialsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CloudCredentialsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


