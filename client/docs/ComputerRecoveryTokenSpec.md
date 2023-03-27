# ComputerRecoveryTokenSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupIds** | **[]string** | Array of backup IDs whose data you want to restore with the recovery token. | 
**ExpirationDate** | **time.Time** | Date and time when the access token expires. | 

## Methods

### NewComputerRecoveryTokenSpec

`func NewComputerRecoveryTokenSpec(backupIds []string, expirationDate time.Time, ) *ComputerRecoveryTokenSpec`

NewComputerRecoveryTokenSpec instantiates a new ComputerRecoveryTokenSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerRecoveryTokenSpecWithDefaults

`func NewComputerRecoveryTokenSpecWithDefaults() *ComputerRecoveryTokenSpec`

NewComputerRecoveryTokenSpecWithDefaults instantiates a new ComputerRecoveryTokenSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupIds

`func (o *ComputerRecoveryTokenSpec) GetBackupIds() []string`

GetBackupIds returns the BackupIds field if non-nil, zero value otherwise.

### GetBackupIdsOk

`func (o *ComputerRecoveryTokenSpec) GetBackupIdsOk() (*[]string, bool)`

GetBackupIdsOk returns a tuple with the BackupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIds

`func (o *ComputerRecoveryTokenSpec) SetBackupIds(v []string)`

SetBackupIds sets BackupIds field to given value.


### GetExpirationDate

`func (o *ComputerRecoveryTokenSpec) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ComputerRecoveryTokenSpec) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ComputerRecoveryTokenSpec) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


