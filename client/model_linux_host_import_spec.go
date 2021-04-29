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

// LinuxHostImportSpec struct for LinuxHostImportSpec
type LinuxHostImportSpec struct {
	// Full DNS name or IP address of the server.
	Name string `json:"name"`
	// Description of the server.
	Description string `json:"description"`
	Type EManagedServerType `json:"type"`
	Credentials CredentialsImportModel `json:"credentials"`
	SshSettings *LinuxHostSSHSettingsModel `json:"sshSettings,omitempty"`
	// SSH key fingerprint used to verify the server identity.
	SshFingerprint string `json:"sshFingerprint"`
}

// NewLinuxHostImportSpec instantiates a new LinuxHostImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxHostImportSpec(name string, description string, type_ EManagedServerType, credentials CredentialsImportModel, sshFingerprint string) *LinuxHostImportSpec {
	this := LinuxHostImportSpec{}
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Credentials = credentials
	this.SshFingerprint = sshFingerprint
	return &this
}

// NewLinuxHostImportSpecWithDefaults instantiates a new LinuxHostImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxHostImportSpecWithDefaults() *LinuxHostImportSpec {
	this := LinuxHostImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *LinuxHostImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinuxHostImportSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinuxHostImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *LinuxHostImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LinuxHostImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LinuxHostImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *LinuxHostImportSpec) GetType() EManagedServerType {
	if o == nil {
		var ret EManagedServerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinuxHostImportSpec) GetTypeOk() (*EManagedServerType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinuxHostImportSpec) SetType(v EManagedServerType) {
	o.Type = v
}

// GetCredentials returns the Credentials field value
func (o *LinuxHostImportSpec) GetCredentials() CredentialsImportModel {
	if o == nil {
		var ret CredentialsImportModel
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *LinuxHostImportSpec) GetCredentialsOk() (*CredentialsImportModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *LinuxHostImportSpec) SetCredentials(v CredentialsImportModel) {
	o.Credentials = v
}

// GetSshSettings returns the SshSettings field value if set, zero value otherwise.
func (o *LinuxHostImportSpec) GetSshSettings() LinuxHostSSHSettingsModel {
	if o == nil || o.SshSettings == nil {
		var ret LinuxHostSSHSettingsModel
		return ret
	}
	return *o.SshSettings
}

// GetSshSettingsOk returns a tuple with the SshSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHostImportSpec) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool) {
	if o == nil || o.SshSettings == nil {
		return nil, false
	}
	return o.SshSettings, true
}

// HasSshSettings returns a boolean if a field has been set.
func (o *LinuxHostImportSpec) HasSshSettings() bool {
	if o != nil && o.SshSettings != nil {
		return true
	}

	return false
}

// SetSshSettings gets a reference to the given LinuxHostSSHSettingsModel and assigns it to the SshSettings field.
func (o *LinuxHostImportSpec) SetSshSettings(v LinuxHostSSHSettingsModel) {
	o.SshSettings = &v
}

// GetSshFingerprint returns the SshFingerprint field value
func (o *LinuxHostImportSpec) GetSshFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SshFingerprint
}

// GetSshFingerprintOk returns a tuple with the SshFingerprint field value
// and a boolean to check if the value has been set.
func (o *LinuxHostImportSpec) GetSshFingerprintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SshFingerprint, true
}

// SetSshFingerprint sets field value
func (o *LinuxHostImportSpec) SetSshFingerprint(v string) {
	o.SshFingerprint = v
}

func (o LinuxHostImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if o.SshSettings != nil {
		toSerialize["sshSettings"] = o.SshSettings
	}
	if true {
		toSerialize["sshFingerprint"] = o.SshFingerprint
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxHostImportSpec struct {
	value *LinuxHostImportSpec
	isSet bool
}

func (v NullableLinuxHostImportSpec) Get() *LinuxHostImportSpec {
	return v.value
}

func (v *NullableLinuxHostImportSpec) Set(val *LinuxHostImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxHostImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxHostImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxHostImportSpec(val *LinuxHostImportSpec) *NullableLinuxHostImportSpec {
	return &NullableLinuxHostImportSpec{value: val, isSet: true}
}

func (v NullableLinuxHostImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxHostImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

