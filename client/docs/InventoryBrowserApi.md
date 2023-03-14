# \InventoryBrowserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllInventoryVmwareHosts**](InventoryBrowserApi.md#GetAllInventoryVmwareHosts) | **Get** /api/v1/inventory/vmware/hosts | Get All VMware vSphere Servers
[**GetVmwareHostObject**](InventoryBrowserApi.md#GetVmwareHostObject) | **Get** /api/v1/inventory/vmware/hosts/{name} | Get VMware vSphere Server Objects



## GetAllInventoryVmwareHosts

> ViRootsResult GetAllInventoryVmwareHosts(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()

Get All VMware vSphere Servers



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
    skip := int32(56) // int32 | Number of VMware vSphere servers to skip. (optional)
    limit := int32(56) // int32 | Maximum number of VMware vSphere servers to return. (optional)
    orderColumn := client.EViRootFiltersOrderColumn("Name") // EViRootFiltersOrderColumn | Sorts VMware vSphere servers by one of the VMware vSphere server parameters. (optional)
    orderAsc := true // bool | Sorts VMware vSphere servers in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters VMware vSphere servers by the `nameFilter` pattern. The pattern can match any VMware vSphere server parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryBrowserApi.GetAllInventoryVmwareHosts(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryBrowserApi.GetAllInventoryVmwareHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInventoryVmwareHosts`: ViRootsResult
    fmt.Fprintf(os.Stdout, "Response from `InventoryBrowserApi.GetAllInventoryVmwareHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInventoryVmwareHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of VMware vSphere servers to skip. | 
 **limit** | **int32** | Maximum number of VMware vSphere servers to return. | 
 **orderColumn** | [**EViRootFiltersOrderColumn**](EViRootFiltersOrderColumn.md) | Sorts VMware vSphere servers by one of the VMware vSphere server parameters. | 
 **orderAsc** | **bool** | Sorts VMware vSphere servers in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters VMware vSphere servers by the &#x60;nameFilter&#x60; pattern. The pattern can match any VMware vSphere server parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 

### Return type

[**ViRootsResult**](ViRootsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmwareHostObject

> VCenterInventoryResult GetVmwareHostObject(ctx, name).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).ObjectIdFilter(objectIdFilter).HierarchyTypeFilter(hierarchyTypeFilter).NameFilter(nameFilter).TypeFilter(typeFilter).ParentContainerNameFilter(parentContainerNameFilter).Execute()

Get VMware vSphere Server Objects



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
    name := "name_example" // string | Name of the VMware vSphere server.
    skip := int32(56) // int32 | Number of objects to skip. (optional)
    limit := int32(56) // int32 | Maximum number of objects to return. (optional)
    orderColumn := client.EvCentersInventoryFiltersOrderColumn("Name") // EvCentersInventoryFiltersOrderColumn | Sorts objects by one of the object parameters. (optional)
    orderAsc := true // bool | Sorts objects in the ascending order by the `orderColumn` parameter. (optional)
    objectIdFilter := "objectIdFilter_example" // string | Filters objects by object ID. (optional)
    hierarchyTypeFilter := client.EHierarchyType("HostsAndClusters") // EHierarchyType | Filters objects by hierarchy type. (optional)
    nameFilter := "nameFilter_example" // string | Filters objects by the `nameFilter` pattern. The pattern can match any object parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. (optional)
    typeFilter := client.EVmwareInventoryType("Unknown") // EVmwareInventoryType | Filters objects by virtual infrastructure type. (optional)
    parentContainerNameFilter := "parentContainerNameFilter_example" // string | Filters objects by name of the parent container. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryBrowserApi.GetVmwareHostObject(context.Background(), name).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).ObjectIdFilter(objectIdFilter).HierarchyTypeFilter(hierarchyTypeFilter).NameFilter(nameFilter).TypeFilter(typeFilter).ParentContainerNameFilter(parentContainerNameFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryBrowserApi.GetVmwareHostObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVmwareHostObject`: VCenterInventoryResult
    fmt.Fprintf(os.Stdout, "Response from `InventoryBrowserApi.GetVmwareHostObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the VMware vSphere server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmwareHostObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **skip** | **int32** | Number of objects to skip. | 
 **limit** | **int32** | Maximum number of objects to return. | 
 **orderColumn** | [**EvCentersInventoryFiltersOrderColumn**](EvCentersInventoryFiltersOrderColumn.md) | Sorts objects by one of the object parameters. | 
 **orderAsc** | **bool** | Sorts objects in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **objectIdFilter** | **string** | Filters objects by object ID. | 
 **hierarchyTypeFilter** | [**EHierarchyType**](EHierarchyType.md) | Filters objects by hierarchy type. | 
 **nameFilter** | **string** | Filters objects by the &#x60;nameFilter&#x60; pattern. The pattern can match any object parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | 
 **typeFilter** | [**EVmwareInventoryType**](EVmwareInventoryType.md) | Filters objects by virtual infrastructure type. | 
 **parentContainerNameFilter** | **string** | Filters objects by name of the parent container. | 

### Return type

[**VCenterInventoryResult**](VCenterInventoryResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

