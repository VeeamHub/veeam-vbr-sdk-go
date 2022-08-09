# \CredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePasswordForCreds**](CredentialsApi.md#ChangePasswordForCreds) | **Post** /api/v1/credentials/{id}/changepassword | Change Password
[**ChangePrivateKeyForCreds**](CredentialsApi.md#ChangePrivateKeyForCreds) | **Post** /api/v1/credentials/{id}/changeprivatekey | Change Linux Private Key
[**ChangeRootPasswordForCreds**](CredentialsApi.md#ChangeRootPasswordForCreds) | **Post** /api/v1/credentials/{id}/changerootpassword | Change Linux Root Password
[**CreateCreds**](CredentialsApi.md#CreateCreds) | **Post** /api/v1/credentials | Add Credentials Record
[**DeleteCreds**](CredentialsApi.md#DeleteCreds) | **Delete** /api/v1/credentials/{id} | Remove Credentials Record
[**GetAllCreds**](CredentialsApi.md#GetAllCreds) | **Get** /api/v1/credentials | Get All Credentials
[**GetCreds**](CredentialsApi.md#GetCreds) | **Get** /api/v1/credentials/{id} | Get Credentials Record
[**UpdateCreds**](CredentialsApi.md#UpdateCreds) | **Put** /api/v1/credentials/{id} | Edit Credentials Record



## ChangePasswordForCreds

> map[string]interface{} ChangePasswordForCreds(ctx, id).XApiVersion(xApiVersion).CredentialsPasswordChangeSpec(credentialsPasswordChangeSpec).Execute()

Change Password



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
    id := TODO // string | ID of the credentials record.
    credentialsPasswordChangeSpec := *openapiclient.NewCredentialsPasswordChangeSpec("Password_example") // CredentialsPasswordChangeSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.ChangePasswordForCreds(context.Background(), id).XApiVersion(xApiVersion).CredentialsPasswordChangeSpec(credentialsPasswordChangeSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.ChangePasswordForCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangePasswordForCreds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.ChangePasswordForCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordForCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]

 **credentialsPasswordChangeSpec** | [**CredentialsPasswordChangeSpec**](CredentialsPasswordChangeSpec.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePrivateKeyForCreds

> map[string]interface{} ChangePrivateKeyForCreds(ctx, id).XApiVersion(xApiVersion).PrivateKeyChangeSpec(privateKeyChangeSpec).Execute()

Change Linux Private Key



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
    id := TODO // string | ID of the credentials record.
    privateKeyChangeSpec := *openapiclient.NewPrivateKeyChangeSpec("PrivateKey_example") // PrivateKeyChangeSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.ChangePrivateKeyForCreds(context.Background(), id).XApiVersion(xApiVersion).PrivateKeyChangeSpec(privateKeyChangeSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.ChangePrivateKeyForCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangePrivateKeyForCreds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.ChangePrivateKeyForCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePrivateKeyForCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]

 **privateKeyChangeSpec** | [**PrivateKeyChangeSpec**](PrivateKeyChangeSpec.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRootPasswordForCreds

> map[string]interface{} ChangeRootPasswordForCreds(ctx, id).XApiVersion(xApiVersion).CredentialsPasswordChangeSpec(credentialsPasswordChangeSpec).Execute()

Change Linux Root Password



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
    id := TODO // string | ID of the credentials record.
    credentialsPasswordChangeSpec := *openapiclient.NewCredentialsPasswordChangeSpec("Password_example") // CredentialsPasswordChangeSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.ChangeRootPasswordForCreds(context.Background(), id).XApiVersion(xApiVersion).CredentialsPasswordChangeSpec(credentialsPasswordChangeSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.ChangeRootPasswordForCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeRootPasswordForCreds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.ChangeRootPasswordForCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRootPasswordForCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]

 **credentialsPasswordChangeSpec** | [**CredentialsPasswordChangeSpec**](CredentialsPasswordChangeSpec.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreds

> CredentialsModel CreateCreds(ctx).XApiVersion(xApiVersion).CredentialsSpec(credentialsSpec).Execute()

Add Credentials Record



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
    credentialsSpec := *openapiclient.NewCredentialsSpec("Username_example", openapiclient.ECredentialsType("Standard")) // CredentialsSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.CreateCreds(context.Background()).XApiVersion(xApiVersion).CredentialsSpec(credentialsSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.CreateCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreds`: CredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.CreateCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
 **credentialsSpec** | [**CredentialsSpec**](CredentialsSpec.md) |  | 

### Return type

[**CredentialsModel**](CredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCreds

> map[string]interface{} DeleteCreds(ctx, id).XApiVersion(xApiVersion).Execute()

Remove Credentials Record



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
    id := TODO // string | ID of the credentials record.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.DeleteCreds(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.DeleteCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCreds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.DeleteCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCreds

> CredentialsResult GetAllCreds(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()

Get All Credentials



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
    skip := int32(56) // int32 | Number of credentials records to skip. (optional)
    limit := int32(56) // int32 | Maximum number of credentials records to return. (optional)
    orderColumn := openapiclient.ECredentialsFiltersOrderColumn("Username") // ECredentialsFiltersOrderColumn | Sorts credentials by one of the credentials parameters. (optional)
    orderAsc := true // bool | Sorts credentials in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters credentials by the `nameFilter` pattern. The pattern can match any credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.GetAllCreds(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.GetAllCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCreds`: CredentialsResult
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.GetAllCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
 **skip** | **int32** | Number of credentials records to skip. | 
 **limit** | **int32** | Maximum number of credentials records to return. | 
 **orderColumn** | [**ECredentialsFiltersOrderColumn**](ECredentialsFiltersOrderColumn.md) | Sorts credentials by one of the credentials parameters. | 
 **orderAsc** | **bool** | Sorts credentials in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters credentials by the &#x60;nameFilter&#x60; pattern. The pattern can match any credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 

### Return type

[**CredentialsResult**](CredentialsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreds

> CredentialsModel GetCreds(ctx, id).XApiVersion(xApiVersion).Execute()

Get Credentials Record



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
    id := TODO // string | ID of the credentials record.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.GetCreds(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.GetCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreds`: CredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.GetCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


### Return type

[**CredentialsModel**](CredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreds

> CredentialsModel UpdateCreds(ctx, id).XApiVersion(xApiVersion).CredentialsModel(credentialsModel).Execute()

Edit Credentials Record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    id := TODO // string | ID of the credentials record.
    credentialsModel := *openapiclient.NewCredentialsModel("Id_example", "Username_example", "Description_example", openapiclient.ECredentialsType("Standard"), time.Now()) // CredentialsModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialsApi.UpdateCreds(context.Background(), id).XApiVersion(xApiVersion).CredentialsModel(credentialsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.UpdateCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCreds`: CredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.UpdateCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]

 **credentialsModel** | [**CredentialsModel**](CredentialsModel.md) |  | 

### Return type

[**CredentialsModel**](CredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

