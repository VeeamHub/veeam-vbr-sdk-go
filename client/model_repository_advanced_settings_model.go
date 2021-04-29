/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RepositoryAdvancedSettingsModel Advanced settings for the backup repository.
type RepositoryAdvancedSettingsModel struct {
	// If *true*, Veeam Backup & Replication aligns VM data saved to a backup file at a 4 KB block boundary.
	AlignDataBlocks *bool `json:"alignDataBlocks,omitempty"`
	// If *true*, Veeam Backup & Replication decompresses backup data blocks before storing to improve deduplication ratios.
	DecompressBeforeStoring *bool `json:"decompressBeforeStoring,omitempty"`
	// If *true*, the repository drive is rotated.
	RotatedDrives *bool `json:"rotatedDrives,omitempty"`
	// If *true*, Veeam Backup & Replication creates a separate backup file for every VM in the job.
	PerVmBackup *bool `json:"perVmBackup,omitempty"`
}

// NewRepositoryAdvancedSettingsModel instantiates a new RepositoryAdvancedSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryAdvancedSettingsModel() *RepositoryAdvancedSettingsModel {
	this := RepositoryAdvancedSettingsModel{}
	return &this
}

// NewRepositoryAdvancedSettingsModelWithDefaults instantiates a new RepositoryAdvancedSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryAdvancedSettingsModelWithDefaults() *RepositoryAdvancedSettingsModel {
	this := RepositoryAdvancedSettingsModel{}
	return &this
}

// GetAlignDataBlocks returns the AlignDataBlocks field value if set, zero value otherwise.
func (o *RepositoryAdvancedSettingsModel) GetAlignDataBlocks() bool {
	if o == nil || o.AlignDataBlocks == nil {
		var ret bool
		return ret
	}
	return *o.AlignDataBlocks
}

// GetAlignDataBlocksOk returns a tuple with the AlignDataBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAdvancedSettingsModel) GetAlignDataBlocksOk() (*bool, bool) {
	if o == nil || o.AlignDataBlocks == nil {
		return nil, false
	}
	return o.AlignDataBlocks, true
}

// HasAlignDataBlocks returns a boolean if a field has been set.
func (o *RepositoryAdvancedSettingsModel) HasAlignDataBlocks() bool {
	if o != nil && o.AlignDataBlocks != nil {
		return true
	}

	return false
}

// SetAlignDataBlocks gets a reference to the given bool and assigns it to the AlignDataBlocks field.
func (o *RepositoryAdvancedSettingsModel) SetAlignDataBlocks(v bool) {
	o.AlignDataBlocks = &v
}

// GetDecompressBeforeStoring returns the DecompressBeforeStoring field value if set, zero value otherwise.
func (o *RepositoryAdvancedSettingsModel) GetDecompressBeforeStoring() bool {
	if o == nil || o.DecompressBeforeStoring == nil {
		var ret bool
		return ret
	}
	return *o.DecompressBeforeStoring
}

// GetDecompressBeforeStoringOk returns a tuple with the DecompressBeforeStoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAdvancedSettingsModel) GetDecompressBeforeStoringOk() (*bool, bool) {
	if o == nil || o.DecompressBeforeStoring == nil {
		return nil, false
	}
	return o.DecompressBeforeStoring, true
}

// HasDecompressBeforeStoring returns a boolean if a field has been set.
func (o *RepositoryAdvancedSettingsModel) HasDecompressBeforeStoring() bool {
	if o != nil && o.DecompressBeforeStoring != nil {
		return true
	}

	return false
}

// SetDecompressBeforeStoring gets a reference to the given bool and assigns it to the DecompressBeforeStoring field.
func (o *RepositoryAdvancedSettingsModel) SetDecompressBeforeStoring(v bool) {
	o.DecompressBeforeStoring = &v
}

// GetRotatedDrives returns the RotatedDrives field value if set, zero value otherwise.
func (o *RepositoryAdvancedSettingsModel) GetRotatedDrives() bool {
	if o == nil || o.RotatedDrives == nil {
		var ret bool
		return ret
	}
	return *o.RotatedDrives
}

// GetRotatedDrivesOk returns a tuple with the RotatedDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAdvancedSettingsModel) GetRotatedDrivesOk() (*bool, bool) {
	if o == nil || o.RotatedDrives == nil {
		return nil, false
	}
	return o.RotatedDrives, true
}

// HasRotatedDrives returns a boolean if a field has been set.
func (o *RepositoryAdvancedSettingsModel) HasRotatedDrives() bool {
	if o != nil && o.RotatedDrives != nil {
		return true
	}

	return false
}

// SetRotatedDrives gets a reference to the given bool and assigns it to the RotatedDrives field.
func (o *RepositoryAdvancedSettingsModel) SetRotatedDrives(v bool) {
	o.RotatedDrives = &v
}

// GetPerVmBackup returns the PerVmBackup field value if set, zero value otherwise.
func (o *RepositoryAdvancedSettingsModel) GetPerVmBackup() bool {
	if o == nil || o.PerVmBackup == nil {
		var ret bool
		return ret
	}
	return *o.PerVmBackup
}

// GetPerVmBackupOk returns a tuple with the PerVmBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryAdvancedSettingsModel) GetPerVmBackupOk() (*bool, bool) {
	if o == nil || o.PerVmBackup == nil {
		return nil, false
	}
	return o.PerVmBackup, true
}

// HasPerVmBackup returns a boolean if a field has been set.
func (o *RepositoryAdvancedSettingsModel) HasPerVmBackup() bool {
	if o != nil && o.PerVmBackup != nil {
		return true
	}

	return false
}

// SetPerVmBackup gets a reference to the given bool and assigns it to the PerVmBackup field.
func (o *RepositoryAdvancedSettingsModel) SetPerVmBackup(v bool) {
	o.PerVmBackup = &v
}

func (o RepositoryAdvancedSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlignDataBlocks != nil {
		toSerialize["alignDataBlocks"] = o.AlignDataBlocks
	}
	if o.DecompressBeforeStoring != nil {
		toSerialize["decompressBeforeStoring"] = o.DecompressBeforeStoring
	}
	if o.RotatedDrives != nil {
		toSerialize["rotatedDrives"] = o.RotatedDrives
	}
	if o.PerVmBackup != nil {
		toSerialize["perVmBackup"] = o.PerVmBackup
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryAdvancedSettingsModel struct {
	value *RepositoryAdvancedSettingsModel
	isSet bool
}

func (v NullableRepositoryAdvancedSettingsModel) Get() *RepositoryAdvancedSettingsModel {
	return v.value
}

func (v *NullableRepositoryAdvancedSettingsModel) Set(val *RepositoryAdvancedSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryAdvancedSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryAdvancedSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryAdvancedSettingsModel(val *RepositoryAdvancedSettingsModel) *NullableRepositoryAdvancedSettingsModel {
	return &NullableRepositoryAdvancedSettingsModel{value: val, isSet: true}
}

func (v NullableRepositoryAdvancedSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryAdvancedSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

