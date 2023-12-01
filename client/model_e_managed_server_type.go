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

// EManagedServerType Type of the server.
type EManagedServerType string

// List of EManagedServerType
const (
	EMANAGEDSERVERTYPE_WINDOWS_HOST EManagedServerType = "WindowsHost"
	EMANAGEDSERVERTYPE_LINUX_HOST EManagedServerType = "LinuxHost"
	EMANAGEDSERVERTYPE_VI_HOST EManagedServerType = "ViHost"
)

// All allowed values of EManagedServerType enum
var AllowedEManagedServerTypeEnumValues = []EManagedServerType{
	"WindowsHost",
	"LinuxHost",
	"ViHost",
}

func (v *EManagedServerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EManagedServerType(value)
	for _, existing := range AllowedEManagedServerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EManagedServerType", value)
}

// NewEManagedServerTypeFromValue returns a pointer to a valid EManagedServerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEManagedServerTypeFromValue(v string) (*EManagedServerType, error) {
	ev := EManagedServerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EManagedServerType: valid values are %v", v, AllowedEManagedServerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EManagedServerType) IsValid() bool {
	for _, existing := range AllowedEManagedServerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EManagedServerType value
func (v EManagedServerType) Ptr() *EManagedServerType {
	return &v
}

type NullableEManagedServerType struct {
	value *EManagedServerType
	isSet bool
}

func (v NullableEManagedServerType) Get() *EManagedServerType {
	return v.value
}

func (v *NullableEManagedServerType) Set(val *EManagedServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableEManagedServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableEManagedServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEManagedServerType(val *EManagedServerType) *NullableEManagedServerType {
	return &NullableEManagedServerType{value: val, isSet: true}
}

func (v NullableEManagedServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEManagedServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
