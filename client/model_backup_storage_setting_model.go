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

// BackupStorageSettingModel Storage settings.
type BackupStorageSettingModel struct {
	// If *true*, Veeam Backup & Replication deduplicates VM data before storing it in the backup repository.
	EnableInlineDataDedup *bool `json:"enableInlineDataDedup,omitempty"`
	// If *true*, Veeam Backup & Replication excludes swap file blocks from processing.
	ExcludeSwapFileBlocks *bool `json:"excludeSwapFileBlocks,omitempty"`
	// If *true*, Veeam Backup & Replication does not copy deleted file blocks.
	ExcludeDeletedFileBlocks *bool `json:"excludeDeletedFileBlocks,omitempty"`
	CompressionLevel *ECompressionLevel `json:"compressionLevel,omitempty"`
	StorageOptimization *EStorageOptimization `json:"storageOptimization,omitempty"`
	Encryption *BackupStorageSettingsEncryptionModel `json:"encryption,omitempty"`
}

// NewBackupStorageSettingModel instantiates a new BackupStorageSettingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupStorageSettingModel() *BackupStorageSettingModel {
	this := BackupStorageSettingModel{}
	return &this
}

// NewBackupStorageSettingModelWithDefaults instantiates a new BackupStorageSettingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupStorageSettingModelWithDefaults() *BackupStorageSettingModel {
	this := BackupStorageSettingModel{}
	return &this
}

// GetEnableInlineDataDedup returns the EnableInlineDataDedup field value if set, zero value otherwise.
func (o *BackupStorageSettingModel) GetEnableInlineDataDedup() bool {
	if o == nil || o.EnableInlineDataDedup == nil {
		var ret bool
		return ret
	}
	return *o.EnableInlineDataDedup
}

// GetEnableInlineDataDedupOk returns a tuple with the EnableInlineDataDedup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingModel) GetEnableInlineDataDedupOk() (*bool, bool) {
	if o == nil || o.EnableInlineDataDedup == nil {
		return nil, false
	}
	return o.EnableInlineDataDedup, true
}

// HasEnableInlineDataDedup returns a boolean if a field has been set.
func (o *BackupStorageSettingModel) HasEnableInlineDataDedup() bool {
	if o != nil && o.EnableInlineDataDedup != nil {
		return true
	}

	return false
}

// SetEnableInlineDataDedup gets a reference to the given bool and assigns it to the EnableInlineDataDedup field.
func (o *BackupStorageSettingModel) SetEnableInlineDataDedup(v bool) {
	o.EnableInlineDataDedup = &v
}

// GetExcludeSwapFileBlocks returns the ExcludeSwapFileBlocks field value if set, zero value otherwise.
func (o *BackupStorageSettingModel) GetExcludeSwapFileBlocks() bool {
	if o == nil || o.ExcludeSwapFileBlocks == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeSwapFileBlocks
}

// GetExcludeSwapFileBlocksOk returns a tuple with the ExcludeSwapFileBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingModel) GetExcludeSwapFileBlocksOk() (*bool, bool) {
	if o == nil || o.ExcludeSwapFileBlocks == nil {
		return nil, false
	}
	return o.ExcludeSwapFileBlocks, true
}

// HasExcludeSwapFileBlocks returns a boolean if a field has been set.
func (o *BackupStorageSettingModel) HasExcludeSwapFileBlocks() bool {
	if o != nil && o.ExcludeSwapFileBlocks != nil {
		return true
	}

	return false
}

// SetExcludeSwapFileBlocks gets a reference to the given bool and assigns it to the ExcludeSwapFileBlocks field.
func (o *BackupStorageSettingModel) SetExcludeSwapFileBlocks(v bool) {
	o.ExcludeSwapFileBlocks = &v
}

// GetExcludeDeletedFileBlocks returns the ExcludeDeletedFileBlocks field value if set, zero value otherwise.
func (o *BackupStorageSettingModel) GetExcludeDeletedFileBlocks() bool {
	if o == nil || o.ExcludeDeletedFileBlocks == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeDeletedFileBlocks
}

// GetExcludeDeletedFileBlocksOk returns a tuple with the ExcludeDeletedFileBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingModel) GetExcludeDeletedFileBlocksOk() (*bool, bool) {
	if o == nil || o.ExcludeDeletedFileBlocks == nil {
		return nil, false
	}
	return o.ExcludeDeletedFileBlocks, true
}

// HasExcludeDeletedFileBlocks returns a boolean if a field has been set.
func (o *BackupStorageSettingModel) HasExcludeDeletedFileBlocks() bool {
	if o != nil && o.ExcludeDeletedFileBlocks != nil {
		return true
	}

	return false
}

// SetExcludeDeletedFileBlocks gets a reference to the given bool and assigns it to the ExcludeDeletedFileBlocks field.
func (o *BackupStorageSettingModel) SetExcludeDeletedFileBlocks(v bool) {
	o.ExcludeDeletedFileBlocks = &v
}

// GetCompressionLevel returns the CompressionLevel field value if set, zero value otherwise.
func (o *BackupStorageSettingModel) GetCompressionLevel() ECompressionLevel {
	if o == nil || o.CompressionLevel == nil {
		var ret ECompressionLevel
		return ret
	}
	return *o.CompressionLevel
}

// GetCompressionLevelOk returns a tuple with the CompressionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingModel) GetCompressionLevelOk() (*ECompressionLevel, bool) {
	if o == nil || o.CompressionLevel == nil {
		return nil, false
	}
	return o.CompressionLevel, true
}

// HasCompressionLevel returns a boolean if a field has been set.
func (o *BackupStorageSettingModel) HasCompressionLevel() bool {
	if o != nil && o.CompressionLevel != nil {
		return true
	}

	return false
}

// SetCompressionLevel gets a reference to the given ECompressionLevel and assigns it to the CompressionLevel field.
func (o *BackupStorageSettingModel) SetCompressionLevel(v ECompressionLevel) {
	o.CompressionLevel = &v
}

// GetStorageOptimization returns the StorageOptimization field value if set, zero value otherwise.
func (o *BackupStorageSettingModel) GetStorageOptimization() EStorageOptimization {
	if o == nil || o.StorageOptimization == nil {
		var ret EStorageOptimization
		return ret
	}
	return *o.StorageOptimization
}

// GetStorageOptimizationOk returns a tuple with the StorageOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingModel) GetStorageOptimizationOk() (*EStorageOptimization, bool) {
	if o == nil || o.StorageOptimization == nil {
		return nil, false
	}
	return o.StorageOptimization, true
}

// HasStorageOptimization returns a boolean if a field has been set.
func (o *BackupStorageSettingModel) HasStorageOptimization() bool {
	if o != nil && o.StorageOptimization != nil {
		return true
	}

	return false
}

// SetStorageOptimization gets a reference to the given EStorageOptimization and assigns it to the StorageOptimization field.
func (o *BackupStorageSettingModel) SetStorageOptimization(v EStorageOptimization) {
	o.StorageOptimization = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *BackupStorageSettingModel) GetEncryption() BackupStorageSettingsEncryptionModel {
	if o == nil || o.Encryption == nil {
		var ret BackupStorageSettingsEncryptionModel
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageSettingModel) GetEncryptionOk() (*BackupStorageSettingsEncryptionModel, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *BackupStorageSettingModel) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given BackupStorageSettingsEncryptionModel and assigns it to the Encryption field.
func (o *BackupStorageSettingModel) SetEncryption(v BackupStorageSettingsEncryptionModel) {
	o.Encryption = &v
}

func (o BackupStorageSettingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableInlineDataDedup != nil {
		toSerialize["enableInlineDataDedup"] = o.EnableInlineDataDedup
	}
	if o.ExcludeSwapFileBlocks != nil {
		toSerialize["excludeSwapFileBlocks"] = o.ExcludeSwapFileBlocks
	}
	if o.ExcludeDeletedFileBlocks != nil {
		toSerialize["excludeDeletedFileBlocks"] = o.ExcludeDeletedFileBlocks
	}
	if o.CompressionLevel != nil {
		toSerialize["compressionLevel"] = o.CompressionLevel
	}
	if o.StorageOptimization != nil {
		toSerialize["storageOptimization"] = o.StorageOptimization
	}
	if o.Encryption != nil {
		toSerialize["encryption"] = o.Encryption
	}
	return json.Marshal(toSerialize)
}

type NullableBackupStorageSettingModel struct {
	value *BackupStorageSettingModel
	isSet bool
}

func (v NullableBackupStorageSettingModel) Get() *BackupStorageSettingModel {
	return v.value
}

func (v *NullableBackupStorageSettingModel) Set(val *BackupStorageSettingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupStorageSettingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupStorageSettingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupStorageSettingModel(val *BackupStorageSettingModel) *NullableBackupStorageSettingModel {
	return &NullableBackupStorageSettingModel{value: val, isSet: true}
}

func (v NullableBackupStorageSettingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupStorageSettingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

