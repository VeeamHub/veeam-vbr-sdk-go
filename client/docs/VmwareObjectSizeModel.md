# VmwareObjectSizeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**Size** | Pointer to **string** | Size used by the VMware vSphere object. | [optional] 

## Methods

### NewVmwareObjectSizeModel

`func NewVmwareObjectSizeModel(inventoryObject VmwareObjectModel, ) *VmwareObjectSizeModel`

NewVmwareObjectSizeModel instantiates a new VmwareObjectSizeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectSizeModelWithDefaults

`func NewVmwareObjectSizeModelWithDefaults() *VmwareObjectSizeModel`

NewVmwareObjectSizeModelWithDefaults instantiates a new VmwareObjectSizeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryObject

`func (o *VmwareObjectSizeModel) GetInventoryObject() VmwareObjectModel`

GetInventoryObject returns the InventoryObject field if non-nil, zero value otherwise.

### GetInventoryObjectOk

`func (o *VmwareObjectSizeModel) GetInventoryObjectOk() (*VmwareObjectModel, bool)`

GetInventoryObjectOk returns a tuple with the InventoryObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryObject

`func (o *VmwareObjectSizeModel) SetInventoryObject(v VmwareObjectModel)`

SetInventoryObject sets InventoryObject field to given value.


### GetSize

`func (o *VmwareObjectSizeModel) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VmwareObjectSizeModel) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VmwareObjectSizeModel) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VmwareObjectSizeModel) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


