# CredentialsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CredentialsModel**](CredentialsModel.md) | Array of credentials. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewCredentialsResult

`func NewCredentialsResult(data []CredentialsModel, pagination PaginationResult, ) *CredentialsResult`

NewCredentialsResult instantiates a new CredentialsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResultWithDefaults

`func NewCredentialsResultWithDefaults() *CredentialsResult`

NewCredentialsResultWithDefaults instantiates a new CredentialsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CredentialsResult) GetData() []CredentialsModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CredentialsResult) GetDataOk() (*[]CredentialsModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CredentialsResult) SetData(v []CredentialsModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *CredentialsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CredentialsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CredentialsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


