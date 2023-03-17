# SecureRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAntivirusScan** | **bool** | If *true*, Veeam Backup &amp; Replication scans machine data with antivirus software before restoring the machine to the production environment. | 
**VirusDetectionAction** | Pointer to [**EVirusDetectionAction**](EVirusDetectionAction.md) |  | [optional] 
**EnableEntireVolumeScan** | Pointer to **bool** | If *true*, the antivirus continues machine scan after the first malware is found. | [optional] 

## Methods

### NewSecureRestoreSpec

`func NewSecureRestoreSpec(enableAntivirusScan bool, ) *SecureRestoreSpec`

NewSecureRestoreSpec instantiates a new SecureRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureRestoreSpecWithDefaults

`func NewSecureRestoreSpecWithDefaults() *SecureRestoreSpec`

NewSecureRestoreSpecWithDefaults instantiates a new SecureRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAntivirusScan

`func (o *SecureRestoreSpec) GetEnableAntivirusScan() bool`

GetEnableAntivirusScan returns the EnableAntivirusScan field if non-nil, zero value otherwise.

### GetEnableAntivirusScanOk

`func (o *SecureRestoreSpec) GetEnableAntivirusScanOk() (*bool, bool)`

GetEnableAntivirusScanOk returns a tuple with the EnableAntivirusScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAntivirusScan

`func (o *SecureRestoreSpec) SetEnableAntivirusScan(v bool)`

SetEnableAntivirusScan sets EnableAntivirusScan field to given value.


### GetVirusDetectionAction

`func (o *SecureRestoreSpec) GetVirusDetectionAction() EVirusDetectionAction`

GetVirusDetectionAction returns the VirusDetectionAction field if non-nil, zero value otherwise.

### GetVirusDetectionActionOk

`func (o *SecureRestoreSpec) GetVirusDetectionActionOk() (*EVirusDetectionAction, bool)`

GetVirusDetectionActionOk returns a tuple with the VirusDetectionAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirusDetectionAction

`func (o *SecureRestoreSpec) SetVirusDetectionAction(v EVirusDetectionAction)`

SetVirusDetectionAction sets VirusDetectionAction field to given value.

### HasVirusDetectionAction

`func (o *SecureRestoreSpec) HasVirusDetectionAction() bool`

HasVirusDetectionAction returns a boolean if a field has been set.

### GetEnableEntireVolumeScan

`func (o *SecureRestoreSpec) GetEnableEntireVolumeScan() bool`

GetEnableEntireVolumeScan returns the EnableEntireVolumeScan field if non-nil, zero value otherwise.

### GetEnableEntireVolumeScanOk

`func (o *SecureRestoreSpec) GetEnableEntireVolumeScanOk() (*bool, bool)`

GetEnableEntireVolumeScanOk returns a tuple with the EnableEntireVolumeScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEntireVolumeScan

`func (o *SecureRestoreSpec) SetEnableEntireVolumeScan(v bool)`

SetEnableEntireVolumeScan sets EnableEntireVolumeScan field to given value.

### HasEnableEntireVolumeScan

`func (o *SecureRestoreSpec) HasEnableEntireVolumeScan() bool`

HasEnableEntireVolumeScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


