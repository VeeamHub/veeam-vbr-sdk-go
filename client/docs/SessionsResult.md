# SessionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SessionModel**](SessionModel.md) | Array of sessions. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewSessionsResult

`func NewSessionsResult(data []SessionModel, pagination PaginationResult, ) *SessionsResult`

NewSessionsResult instantiates a new SessionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionsResultWithDefaults

`func NewSessionsResultWithDefaults() *SessionsResult`

NewSessionsResultWithDefaults instantiates a new SessionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SessionsResult) GetData() []SessionModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SessionsResult) GetDataOk() (*[]SessionModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SessionsResult) SetData(v []SessionModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *SessionsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SessionsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SessionsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


