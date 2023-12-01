/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EBackupProxyImportType the model 'EBackupProxyImportType'
type EBackupProxyImportType string

// List of EBackupProxyImportType
const (
	EBACKUPPROXYIMPORTTYPE_VMWARE EBackupProxyImportType = "vmware"
)

// All allowed values of EBackupProxyImportType enum
var AllowedEBackupProxyImportTypeEnumValues = []EBackupProxyImportType{
	"vmware",
}

func (v *EBackupProxyImportType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EBackupProxyImportType(value)
	for _, existing := range AllowedEBackupProxyImportTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EBackupProxyImportType", value)
}

// NewEBackupProxyImportTypeFromValue returns a pointer to a valid EBackupProxyImportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEBackupProxyImportTypeFromValue(v string) (*EBackupProxyImportType, error) {
	ev := EBackupProxyImportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EBackupProxyImportType: valid values are %v", v, AllowedEBackupProxyImportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EBackupProxyImportType) IsValid() bool {
	for _, existing := range AllowedEBackupProxyImportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EBackupProxyImportType value
func (v EBackupProxyImportType) Ptr() *EBackupProxyImportType {
	return &v
}

type NullableEBackupProxyImportType struct {
	value *EBackupProxyImportType
	isSet bool
}

func (v NullableEBackupProxyImportType) Get() *EBackupProxyImportType {
	return v.value
}

func (v *NullableEBackupProxyImportType) Set(val *EBackupProxyImportType) {
	v.value = val
	v.isSet = true
}

func (v NullableEBackupProxyImportType) IsSet() bool {
	return v.isSet
}

func (v *NullableEBackupProxyImportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBackupProxyImportType(val *EBackupProxyImportType) *NullableEBackupProxyImportType {
	return &NullableEBackupProxyImportType{value: val, isSet: true}
}

func (v NullableEBackupProxyImportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBackupProxyImportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
