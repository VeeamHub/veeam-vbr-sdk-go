# ServerTimeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTime** | **time.Time** | Current date and time on the backup server. | 
**TimeZone** | Pointer to **string** | Time zone where the backup server is located. | [optional] 

## Methods

### NewServerTimeModel

`func NewServerTimeModel(serverTime time.Time, ) *ServerTimeModel`

NewServerTimeModel instantiates a new ServerTimeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTimeModelWithDefaults

`func NewServerTimeModelWithDefaults() *ServerTimeModel`

NewServerTimeModelWithDefaults instantiates a new ServerTimeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTime

`func (o *ServerTimeModel) GetServerTime() time.Time`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ServerTimeModel) GetServerTimeOk() (*time.Time, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ServerTimeModel) SetServerTime(v time.Time)`

SetServerTime sets ServerTime field to given value.


### GetTimeZone

`func (o *ServerTimeModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ServerTimeModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ServerTimeModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ServerTimeModel) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


