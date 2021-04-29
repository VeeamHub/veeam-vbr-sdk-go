# \RestoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVmwareFcdInstantRecoveryMountModel**](RestoreApi.md#GetVmwareFcdInstantRecoveryMountModel) | **Get** /api/v1/restore/instantRecovery/vmware/fcd/{mountId} | Get Mount Information
[**InstantRecoveryVmwareFcdDismount**](RestoreApi.md#InstantRecoveryVmwareFcdDismount) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/dismount | Stop FCD Publishing
[**InstantRecoveryVmwareFcdMigrate**](RestoreApi.md#InstantRecoveryVmwareFcdMigrate) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/migrate | Start FCD Migration
[**InstantRecoveryVmwareFcdMount**](RestoreApi.md#InstantRecoveryVmwareFcdMount) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/ | Start Instant FCD Recovery



## GetVmwareFcdInstantRecoveryMountModel

> VmwareFcdInstantRecoveryMount GetVmwareFcdInstantRecoveryMountModel(ctx, mountId).XApiVersion(xApiVersion).Execute()

Get Mount Information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    mountId := TODO // string | Mount ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.GetVmwareFcdInstantRecoveryMountModel(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.GetVmwareFcdInstantRecoveryMountModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVmwareFcdInstantRecoveryMountModel`: VmwareFcdInstantRecoveryMount
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.GetVmwareFcdInstantRecoveryMountModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | [**string**](.md) | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmwareFcdInstantRecoveryMountModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


### Return type

[**VmwareFcdInstantRecoveryMount**](VmwareFcdInstantRecoveryMount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantRecoveryVmwareFcdDismount

> VmwareFcdInstantRecoveryMount InstantRecoveryVmwareFcdDismount(ctx, mountId).XApiVersion(xApiVersion).Execute()

Stop FCD Publishing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    mountId := TODO // string | Mount ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.InstantRecoveryVmwareFcdDismount(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantRecoveryVmwareFcdDismount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantRecoveryVmwareFcdDismount`: VmwareFcdInstantRecoveryMount
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantRecoveryVmwareFcdDismount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | [**string**](.md) | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdDismountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


### Return type

[**VmwareFcdInstantRecoveryMount**](VmwareFcdInstantRecoveryMount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantRecoveryVmwareFcdMigrate

> VmwareFcdInstantRecoveryMount InstantRecoveryVmwareFcdMigrate(ctx, mountId).XApiVersion(xApiVersion).VmwareFcdQuickMigrationSpec(vmwareFcdQuickMigrationSpec).Execute()

Start FCD Migration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    mountId := TODO // string | Mount ID.
    vmwareFcdQuickMigrationSpec := *openapiclient.NewVmwareFcdQuickMigrationSpec(*openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown"))) // VmwareFcdQuickMigrationSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.InstantRecoveryVmwareFcdMigrate(context.Background(), mountId).XApiVersion(xApiVersion).VmwareFcdQuickMigrationSpec(vmwareFcdQuickMigrationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantRecoveryVmwareFcdMigrate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantRecoveryVmwareFcdMigrate`: VmwareFcdInstantRecoveryMount
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantRecoveryVmwareFcdMigrate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | [**string**](.md) | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdMigrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

 **vmwareFcdQuickMigrationSpec** | [**VmwareFcdQuickMigrationSpec**](VmwareFcdQuickMigrationSpec.md) |  | 

### Return type

[**VmwareFcdInstantRecoveryMount**](VmwareFcdInstantRecoveryMount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantRecoveryVmwareFcdMount

> VmwareFcdInstantRecoveryMount InstantRecoveryVmwareFcdMount(ctx).XApiVersion(xApiVersion).VmwareFcdInstantRecoverySpec(vmwareFcdInstantRecoverySpec).Execute()

Start Instant FCD Recovery



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    vmwareFcdInstantRecoverySpec := *openapiclient.NewVmwareFcdInstantRecoverySpec("ObjectRestorePointId_example", *openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown")), []openapiclient.VmwareFcdInstantRecoveryDiskSpec{*openapiclient.NewVmwareFcdInstantRecoveryDiskSpec("NameInBackup_example", "MountedDiskName_example", "RegisteredFcdName_example")}) // VmwareFcdInstantRecoverySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.InstantRecoveryVmwareFcdMount(context.Background()).XApiVersion(xApiVersion).VmwareFcdInstantRecoverySpec(vmwareFcdInstantRecoverySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantRecoveryVmwareFcdMount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantRecoveryVmwareFcdMount`: VmwareFcdInstantRecoveryMount
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantRecoveryVmwareFcdMount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdMountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **vmwareFcdInstantRecoverySpec** | [**VmwareFcdInstantRecoverySpec**](VmwareFcdInstantRecoverySpec.md) |  | 

### Return type

[**VmwareFcdInstantRecoveryMount**](VmwareFcdInstantRecoveryMount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

