/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PaginationResult Pagination settings.
type PaginationResult struct {
	// Total number of results.
	Total int32 `json:"total"`
	// Number of returned results.
	Count int32 `json:"count"`
	// Number of skipped results.
	Skip *int32 `json:"skip,omitempty"`
	// Maximum number of results to return.
	Limit *int32 `json:"limit,omitempty"`
}

// NewPaginationResult instantiates a new PaginationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResult(total int32, count int32) *PaginationResult {
	this := PaginationResult{}
	this.Total = total
	this.Count = count
	return &this
}

// NewPaginationResultWithDefaults instantiates a new PaginationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResultWithDefaults() *PaginationResult {
	this := PaginationResult{}
	return &this
}

// GetTotal returns the Total field value
func (o *PaginationResult) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetTotalOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PaginationResult) SetTotal(v int32) {
	o.Total = v
}

// GetCount returns the Count field value
func (o *PaginationResult) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginationResult) SetCount(v int32) {
	o.Count = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PaginationResult) GetSkip() int32 {
	if o == nil || isNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetSkipOk() (*int32, bool) {
	if o == nil || isNil(o.Skip) {
    return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PaginationResult) HasSkip() bool {
	if o != nil && !isNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PaginationResult) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaginationResult) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaginationResult) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaginationResult) SetLimit(v int32) {
	o.Limit = &v
}

func (o PaginationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationResult struct {
	value *PaginationResult
	isSet bool
}

func (v NullablePaginationResult) Get() *PaginationResult {
	return v.value
}

func (v *NullablePaginationResult) Set(val *PaginationResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResult(val *PaginationResult) *NullablePaginationResult {
	return &NullablePaginationResult{value: val, isSet: true}
}

func (v NullablePaginationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

