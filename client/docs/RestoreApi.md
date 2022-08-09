# \RestoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllVmwareFcdInstantRecoveryMountModels**](RestoreApi.md#GetAllVmwareFcdInstantRecoveryMountModels) | **Get** /api/v1/restore/instantRecovery/vmware/fcd/ | Get All FCD Mounts Information
[**GetVmwareFcdInstantRecoveryMountModel**](RestoreApi.md#GetVmwareFcdInstantRecoveryMountModel) | **Get** /api/v1/restore/instantRecovery/vmware/fcd/{mountId} | Get Mount Information
[**InstantRecoveryVmwareFcdDismountWithSession**](RestoreApi.md#InstantRecoveryVmwareFcdDismountWithSession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/dismount | Stop FCD Publishing
[**InstantRecoveryVmwareFcdMigrateWithSession**](RestoreApi.md#InstantRecoveryVmwareFcdMigrateWithSession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/migrate | Start FCD Migration
[**InstantRecoveryVmwareFcdMountWithSession**](RestoreApi.md#InstantRecoveryVmwareFcdMountWithSession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/ | Start Instant FCD Recovery



## GetAllVmwareFcdInstantRecoveryMountModels

> VmwareFcdInstantRecoveryMountsResult GetAllVmwareFcdInstantRecoveryMountModels(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).StateFilter(stateFilter).Execute()

Get All FCD Mounts Information



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    skip := int32(56) // int32 | Number of restore points to skip. (optional)
    limit := int32(56) // int32 | Maximum number of restore points to return. (optional)
    orderColumn := openapiclient.EVmwareFcdInstantRecoveryMountsFiltersOrderColumn("state") // EVmwareFcdInstantRecoveryMountsFiltersOrderColumn | Sorts restore points by one of the restore point parameters. (optional)
    orderAsc := true // bool | Sorts restore points in the ascending order by the `orderColumn` parameter. (optional)
    stateFilter := openapiclient.EInstantRecoveryMountState("Failed") // EInstantRecoveryMountState | Filters vm mounts by mount state. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.GetAllVmwareFcdInstantRecoveryMountModels(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).StateFilter(stateFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.GetAllVmwareFcdInstantRecoveryMountModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllVmwareFcdInstantRecoveryMountModels`: VmwareFcdInstantRecoveryMountsResult
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.GetAllVmwareFcdInstantRecoveryMountModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllVmwareFcdInstantRecoveryMountModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
 **skip** | **int32** | Number of restore points to skip. | 
 **limit** | **int32** | Maximum number of restore points to return. | 
 **orderColumn** | [**EVmwareFcdInstantRecoveryMountsFiltersOrderColumn**](EVmwareFcdInstantRecoveryMountsFiltersOrderColumn.md) | Sorts restore points by one of the restore point parameters. | 
 **orderAsc** | **bool** | Sorts restore points in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **stateFilter** | [**EInstantRecoveryMountState**](EInstantRecoveryMountState.md) | Filters vm mounts by mount state. | 

### Return type

[**VmwareFcdInstantRecoveryMountsResult**](VmwareFcdInstantRecoveryMountsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


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


## InstantRecoveryVmwareFcdDismountWithSession

> SessionModel InstantRecoveryVmwareFcdDismountWithSession(ctx, mountId).XApiVersion(xApiVersion).Execute()

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    mountId := TODO // string | Mount ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.InstantRecoveryVmwareFcdDismountWithSession(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantRecoveryVmwareFcdDismountWithSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantRecoveryVmwareFcdDismountWithSession`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantRecoveryVmwareFcdDismountWithSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | [**string**](.md) | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdDismountWithSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantRecoveryVmwareFcdMigrateWithSession

> SessionModel InstantRecoveryVmwareFcdMigrateWithSession(ctx, mountId).XApiVersion(xApiVersion).VmwareFcdQuickMigrationSpec(vmwareFcdQuickMigrationSpec).Execute()

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    mountId := TODO // string | Mount ID.
    vmwareFcdQuickMigrationSpec := *openapiclient.NewVmwareFcdQuickMigrationSpec(*openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown"))) // VmwareFcdQuickMigrationSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.InstantRecoveryVmwareFcdMigrateWithSession(context.Background(), mountId).XApiVersion(xApiVersion).VmwareFcdQuickMigrationSpec(vmwareFcdQuickMigrationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantRecoveryVmwareFcdMigrateWithSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantRecoveryVmwareFcdMigrateWithSession`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantRecoveryVmwareFcdMigrateWithSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | [**string**](.md) | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdMigrateWithSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]

 **vmwareFcdQuickMigrationSpec** | [**VmwareFcdQuickMigrationSpec**](VmwareFcdQuickMigrationSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantRecoveryVmwareFcdMountWithSession

> SessionModel InstantRecoveryVmwareFcdMountWithSession(ctx).XApiVersion(xApiVersion).VmwareFcdInstantRecoverySpec(vmwareFcdInstantRecoverySpec).Execute()

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    vmwareFcdInstantRecoverySpec := *openapiclient.NewVmwareFcdInstantRecoverySpec("ObjectRestorePointId_example", *openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown")), []openapiclient.VmwareFcdInstantRecoveryDiskSpec{*openapiclient.NewVmwareFcdInstantRecoveryDiskSpec("NameInBackup_example", "MountedDiskName_example", "RegisteredFcdName_example")}) // VmwareFcdInstantRecoverySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RestoreApi.InstantRecoveryVmwareFcdMountWithSession(context.Background()).XApiVersion(xApiVersion).VmwareFcdInstantRecoverySpec(vmwareFcdInstantRecoverySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantRecoveryVmwareFcdMountWithSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantRecoveryVmwareFcdMountWithSession`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantRecoveryVmwareFcdMountWithSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdMountWithSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
 **vmwareFcdInstantRecoverySpec** | [**VmwareFcdInstantRecoverySpec**](VmwareFcdInstantRecoverySpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

