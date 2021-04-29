# GuestOsCredentialsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creds** | Pointer to [**CredentialsImportModel**](CredentialsImportModel.md) |  | [optional] 
**CredentialsPerMachine** | Pointer to [**[]GuestOsCredentialsPerMachineImportModel**](GuestOsCredentialsPerMachineImportModel.md) | Individual credentials for VMs. | [optional] 

## Methods

### NewGuestOsCredentialsImportModel

`func NewGuestOsCredentialsImportModel() *GuestOsCredentialsImportModel`

NewGuestOsCredentialsImportModel instantiates a new GuestOsCredentialsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestOsCredentialsImportModelWithDefaults

`func NewGuestOsCredentialsImportModelWithDefaults() *GuestOsCredentialsImportModel`

NewGuestOsCredentialsImportModelWithDefaults instantiates a new GuestOsCredentialsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreds

`func (o *GuestOsCredentialsImportModel) GetCreds() CredentialsImportModel`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *GuestOsCredentialsImportModel) GetCredsOk() (*CredentialsImportModel, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *GuestOsCredentialsImportModel) SetCreds(v CredentialsImportModel)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *GuestOsCredentialsImportModel) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetCredentialsPerMachine

`func (o *GuestOsCredentialsImportModel) GetCredentialsPerMachine() []GuestOsCredentialsPerMachineImportModel`

GetCredentialsPerMachine returns the CredentialsPerMachine field if non-nil, zero value otherwise.

### GetCredentialsPerMachineOk

`func (o *GuestOsCredentialsImportModel) GetCredentialsPerMachineOk() (*[]GuestOsCredentialsPerMachineImportModel, bool)`

GetCredentialsPerMachineOk returns a tuple with the CredentialsPerMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsPerMachine

`func (o *GuestOsCredentialsImportModel) SetCredentialsPerMachine(v []GuestOsCredentialsPerMachineImportModel)`

SetCredentialsPerMachine sets CredentialsPerMachine field to given value.

### HasCredentialsPerMachine

`func (o *GuestOsCredentialsImportModel) HasCredentialsPerMachine() bool`

HasCredentialsPerMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


