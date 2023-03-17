# \RestoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntireVmRestoreVmware**](RestoreApi.md#EntireVmRestoreVmware) | **Post** /api/v1/restore/vmRestore/vmware/ | Start Entire VM Restore
[**GetAllInstantViVMRecoveryMounts**](RestoreApi.md#GetAllInstantViVMRecoveryMounts) | **Get** /api/v1/restore/instantRecovery/vmware/vm | Get All VM Mounts
[**GetAllVmwareFcdInstantRecoveryMountModels**](RestoreApi.md#GetAllVmwareFcdInstantRecoveryMountModels) | **Get** /api/v1/restore/instantRecovery/vmware/fcd | Get All FCD Mounts
[**GetInstantViVMRecoveryMount**](RestoreApi.md#GetInstantViVMRecoveryMount) | **Get** /api/v1/restore/instantRecovery/vmware/vm/{mountId} | Get VM Mount
[**GetVmwareFcdInstantRecoveryMountModel**](RestoreApi.md#GetVmwareFcdInstantRecoveryMountModel) | **Get** /api/v1/restore/instantRecovery/vmware/fcd/{mountId} | Get FCD Mount
[**InstantRecoveryVmwareFcdDismountWithSession**](RestoreApi.md#InstantRecoveryVmwareFcdDismountWithSession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/dismount | Stop FCD Publishing
[**InstantRecoveryVmwareFcdMigrateWithSession**](RestoreApi.md#InstantRecoveryVmwareFcdMigrateWithSession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/migrate | Start FCD Migration
[**InstantRecoveryVmwareFcdMountWithSession**](RestoreApi.md#InstantRecoveryVmwareFcdMountWithSession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd | Start Instant FCD Recovery
[**InstantViVMRecoveryMigrate**](RestoreApi.md#InstantViVMRecoveryMigrate) | **Post** /api/v1/restore/instantRecovery/vmware/vm/{mountId}/migrate | Start VM Migration
[**InstantViVMRecoveryMount**](RestoreApi.md#InstantViVMRecoveryMount) | **Post** /api/v1/restore/instantRecovery/vmware/vm | Start Instant Recovery
[**InstantViVMRecoveryUnmount**](RestoreApi.md#InstantViVMRecoveryUnmount) | **Post** /api/v1/restore/instantRecovery/vmware/vm/{mountId}/unmount | Stop VM Publishing



## EntireVmRestoreVmware

> SessionModel EntireVmRestoreVmware(ctx).XApiVersion(xApiVersion).EntireViVMRestoreSpec(entireViVMRestoreSpec).Execute()

Start Entire VM Restore



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    entireViVMRestoreSpec := *client.NewEntireViVMRestoreSpec("ObjectRestorePointId_example", client.EEntireVMRestoreModeType("OriginalLocation")) // EntireViVMRestoreSpec | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.EntireVmRestoreVmware(context.Background()).XApiVersion(xApiVersion).EntireViVMRestoreSpec(entireViVMRestoreSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.EntireVmRestoreVmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntireVmRestoreVmware`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.EntireVmRestoreVmware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntireVmRestoreVmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **entireViVMRestoreSpec** | [**EntireViVMRestoreSpec**](EntireViVMRestoreSpec.md) |  | 

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


## GetAllInstantViVMRecoveryMounts

> InstantViVMRecoveryMountsResult GetAllInstantViVMRecoveryMounts(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).StateFilter(stateFilter).Execute()

Get All VM Mounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    skip := int32(56) // int32 | Number of mounts to skip. (optional)
    limit := int32(56) // int32 | Maximum number of mounts to return. (optional)
    orderColumn := client.EInstantViVMRecoveryMountsFiltersOrderColumn("state") // EInstantViVMRecoveryMountsFiltersOrderColumn | Sorts mounts by one of the mount parameters. (optional)
    orderAsc := true // bool | Sorts mounts in the ascending order by the `orderColumn` parameter. (optional)
    stateFilter := client.EInstantRecoveryMountState("Failed") // EInstantRecoveryMountState | Filters mounts by mount state. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.GetAllInstantViVMRecoveryMounts(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).StateFilter(stateFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.GetAllInstantViVMRecoveryMounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInstantViVMRecoveryMounts`: InstantViVMRecoveryMountsResult
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.GetAllInstantViVMRecoveryMounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInstantViVMRecoveryMountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of mounts to skip. | 
 **limit** | **int32** | Maximum number of mounts to return. | 
 **orderColumn** | [**EInstantViVMRecoveryMountsFiltersOrderColumn**](EInstantViVMRecoveryMountsFiltersOrderColumn.md) | Sorts mounts by one of the mount parameters. | 
 **orderAsc** | **bool** | Sorts mounts in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **stateFilter** | [**EInstantRecoveryMountState**](EInstantRecoveryMountState.md) | Filters mounts by mount state. | 

### Return type

[**InstantViVMRecoveryMountsResult**](InstantViVMRecoveryMountsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllVmwareFcdInstantRecoveryMountModels

> VmwareFcdInstantRecoveryMountsResult GetAllVmwareFcdInstantRecoveryMountModels(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).StateFilter(stateFilter).Execute()

Get All FCD Mounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    skip := int32(56) // int32 | Number of mounts to skip. (optional)
    limit := int32(56) // int32 | Maximum number of mounts to return. (optional)
    orderColumn := client.EVmwareFcdInstantRecoveryMountsFiltersOrderColumn("state") // EVmwareFcdInstantRecoveryMountsFiltersOrderColumn | Sorts mounts by one of the mount parameters. (optional)
    orderAsc := true // bool | Sorts mounts in the ascending order by the `orderColumn` parameter. (optional)
    stateFilter := client.EInstantRecoveryMountState("Failed") // EInstantRecoveryMountState | Filters mounts by mount state. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.GetAllVmwareFcdInstantRecoveryMountModels(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).StateFilter(stateFilter).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of mounts to skip. | 
 **limit** | **int32** | Maximum number of mounts to return. | 
 **orderColumn** | [**EVmwareFcdInstantRecoveryMountsFiltersOrderColumn**](EVmwareFcdInstantRecoveryMountsFiltersOrderColumn.md) | Sorts mounts by one of the mount parameters. | 
 **orderAsc** | **bool** | Sorts mounts in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **stateFilter** | [**EInstantRecoveryMountState**](EInstantRecoveryMountState.md) | Filters mounts by mount state. | 

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


## GetInstantViVMRecoveryMount

> InstantViVMRecoveryMount GetInstantViVMRecoveryMount(ctx, mountId).XApiVersion(xApiVersion).Execute()

Get VM Mount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    mountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mount ID.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.GetInstantViVMRecoveryMount(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.GetInstantViVMRecoveryMount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstantViVMRecoveryMount`: InstantViVMRecoveryMount
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.GetInstantViVMRecoveryMount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | **string** | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstantViVMRecoveryMountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**InstantViVMRecoveryMount**](InstantViVMRecoveryMount.md)

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

Get FCD Mount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    mountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mount ID.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.GetVmwareFcdInstantRecoveryMountModel(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
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
**mountId** | **string** | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmwareFcdInstantRecoveryMountModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    mountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mount ID.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.InstantRecoveryVmwareFcdDismountWithSession(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
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
**mountId** | **string** | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdDismountWithSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    mountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mount ID.
    vmwareFcdQuickMigrationSpec := *client.NewVmwareFcdQuickMigrationSpec(*client.NewVmwareObjectModel("HostName_example", "Name_example", client.EVmwareInventoryType("Unknown"))) // VmwareFcdQuickMigrationSpec | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.InstantRecoveryVmwareFcdMigrateWithSession(context.Background(), mountId).XApiVersion(xApiVersion).VmwareFcdQuickMigrationSpec(vmwareFcdQuickMigrationSpec).Execute()
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
**mountId** | **string** | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantRecoveryVmwareFcdMigrateWithSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    vmwareFcdInstantRecoverySpec := *client.NewVmwareFcdInstantRecoverySpec("ObjectRestorePointId_example", *client.NewVmwareObjectModel("HostName_example", "Name_example", client.EVmwareInventoryType("Unknown")), []client.VmwareFcdInstantRecoveryDiskSpec{*client.NewVmwareFcdInstantRecoveryDiskSpec("NameInBackup_example", "MountedDiskName_example", "RegisteredFcdName_example")}) // VmwareFcdInstantRecoverySpec | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.InstantRecoveryVmwareFcdMountWithSession(context.Background()).XApiVersion(xApiVersion).VmwareFcdInstantRecoverySpec(vmwareFcdInstantRecoverySpec).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
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


## InstantViVMRecoveryMigrate

> SessionModel InstantViVMRecoveryMigrate(ctx, mountId).XApiVersion(xApiVersion).ViVMQuickMigrationSpec(viVMQuickMigrationSpec).Execute()

Start VM Migration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    mountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mount ID.
    viVMQuickMigrationSpec := *client.NewViVMQuickMigrationSpec(*client.NewVmwareObjectModel("HostName_example", "Name_example", client.EVmwareInventoryType("Unknown"))) // ViVMQuickMigrationSpec | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.InstantViVMRecoveryMigrate(context.Background(), mountId).XApiVersion(xApiVersion).ViVMQuickMigrationSpec(viVMQuickMigrationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantViVMRecoveryMigrate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantViVMRecoveryMigrate`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantViVMRecoveryMigrate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | **string** | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantViVMRecoveryMigrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **viVMQuickMigrationSpec** | [**ViVMQuickMigrationSpec**](ViVMQuickMigrationSpec.md) |  | 

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


## InstantViVMRecoveryMount

> SessionModel InstantViVMRecoveryMount(ctx).XApiVersion(xApiVersion).InstantViVMRecoverySpec(instantViVMRecoverySpec).Execute()

Start Instant Recovery



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    instantViVMRecoverySpec := *client.NewInstantViVMRecoverySpec("ObjectRestorePointId_example", client.EInstantVMRecoveryModeType("OriginalLocation"), *client.NewSecureRestoreSpec(false)) // InstantViVMRecoverySpec | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.InstantViVMRecoveryMount(context.Background()).XApiVersion(xApiVersion).InstantViVMRecoverySpec(instantViVMRecoverySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantViVMRecoveryMount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantViVMRecoveryMount`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantViVMRecoveryMount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstantViVMRecoveryMountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **instantViVMRecoverySpec** | [**InstantViVMRecoverySpec**](InstantViVMRecoverySpec.md) |  | 

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


## InstantViVMRecoveryUnmount

> SessionModel InstantViVMRecoveryUnmount(ctx, mountId).XApiVersion(xApiVersion).Execute()

Stop VM Publishing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    mountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mount ID.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreApi.InstantViVMRecoveryUnmount(context.Background(), mountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.InstantViVMRecoveryUnmount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantViVMRecoveryUnmount`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RestoreApi.InstantViVMRecoveryUnmount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mountId** | **string** | Mount ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantViVMRecoveryUnmountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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

