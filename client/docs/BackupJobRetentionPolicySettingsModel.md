# BackupJobRetentionPolicySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ERetentionPolicyType**](ERetentionPolicyType.md) |  | 
**Quantity** | **int32** | Number of restore points or days to keep. | 

## Methods

### NewBackupJobRetentionPolicySettingsModel

`func NewBackupJobRetentionPolicySettingsModel(type_ ERetentionPolicyType, quantity int32, ) *BackupJobRetentionPolicySettingsModel`

NewBackupJobRetentionPolicySettingsModel instantiates a new BackupJobRetentionPolicySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobRetentionPolicySettingsModelWithDefaults

`func NewBackupJobRetentionPolicySettingsModelWithDefaults() *BackupJobRetentionPolicySettingsModel`

NewBackupJobRetentionPolicySettingsModelWithDefaults instantiates a new BackupJobRetentionPolicySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackupJobRetentionPolicySettingsModel) GetType() ERetentionPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupJobRetentionPolicySettingsModel) GetTypeOk() (*ERetentionPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupJobRetentionPolicySettingsModel) SetType(v ERetentionPolicyType)`

SetType sets Type field to given value.


### GetQuantity

`func (o *BackupJobRetentionPolicySettingsModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BackupJobRetentionPolicySettingsModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BackupJobRetentionPolicySettingsModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


