# GuestOsCredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredsId** | **string** | Credentials ID for Microsoft Windows VMs. | 
**CredsType** | [**ECredentialsType**](ECredentialsType.md) |  | 
**CredentialsPerMachine** | Pointer to [**[]GuestOsCredentialsPerMachineModel**](GuestOsCredentialsPerMachineModel.md) | Individual credentials for VMs. | [optional] 

## Methods

### NewGuestOsCredentialsModel

`func NewGuestOsCredentialsModel(credsId string, credsType ECredentialsType, ) *GuestOsCredentialsModel`

NewGuestOsCredentialsModel instantiates a new GuestOsCredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestOsCredentialsModelWithDefaults

`func NewGuestOsCredentialsModelWithDefaults() *GuestOsCredentialsModel`

NewGuestOsCredentialsModelWithDefaults instantiates a new GuestOsCredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredsId

`func (o *GuestOsCredentialsModel) GetCredsId() string`

GetCredsId returns the CredsId field if non-nil, zero value otherwise.

### GetCredsIdOk

`func (o *GuestOsCredentialsModel) GetCredsIdOk() (*string, bool)`

GetCredsIdOk returns a tuple with the CredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredsId

`func (o *GuestOsCredentialsModel) SetCredsId(v string)`

SetCredsId sets CredsId field to given value.


### GetCredsType

`func (o *GuestOsCredentialsModel) GetCredsType() ECredentialsType`

GetCredsType returns the CredsType field if non-nil, zero value otherwise.

### GetCredsTypeOk

`func (o *GuestOsCredentialsModel) GetCredsTypeOk() (*ECredentialsType, bool)`

GetCredsTypeOk returns a tuple with the CredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredsType

`func (o *GuestOsCredentialsModel) SetCredsType(v ECredentialsType)`

SetCredsType sets CredsType field to given value.


### GetCredentialsPerMachine

`func (o *GuestOsCredentialsModel) GetCredentialsPerMachine() []GuestOsCredentialsPerMachineModel`

GetCredentialsPerMachine returns the CredentialsPerMachine field if non-nil, zero value otherwise.

### GetCredentialsPerMachineOk

`func (o *GuestOsCredentialsModel) GetCredentialsPerMachineOk() (*[]GuestOsCredentialsPerMachineModel, bool)`

GetCredentialsPerMachineOk returns a tuple with the CredentialsPerMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsPerMachine

`func (o *GuestOsCredentialsModel) SetCredentialsPerMachine(v []GuestOsCredentialsPerMachineModel)`

SetCredentialsPerMachine sets CredentialsPerMachine field to given value.

### HasCredentialsPerMachine

`func (o *GuestOsCredentialsModel) HasCredentialsPerMachine() bool`

HasCredentialsPerMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


