# \CredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeCloudCertificate**](CredentialsApi.md#ChangeCloudCertificate) | **Post** /api/v1/cloudCredentials/{id}/changeCertificate | Change Certificate
[**ChangeCloudCredsSecretKey**](CredentialsApi.md#ChangeCloudCredsSecretKey) | **Post** /api/v1/cloudCredentials/{id}/changeSecretKey | Change Secret Key
[**ChangePasswordForCreds**](CredentialsApi.md#ChangePasswordForCreds) | **Post** /api/v1/credentials/{id}/changepassword | Change Password
[**ChangePrivateKeyForCreds**](CredentialsApi.md#ChangePrivateKeyForCreds) | **Post** /api/v1/credentials/{id}/changeprivatekey | Change Linux Private Key
[**ChangeRootPasswordForCreds**](CredentialsApi.md#ChangeRootPasswordForCreds) | **Post** /api/v1/credentials/{id}/changerootpassword | Change Linux Root Password
[**CreateCloudCreds**](CredentialsApi.md#CreateCloudCreds) | **Post** /api/v1/cloudCredentials | Add Cloud Credentials Record
[**CreateCloudCredsHelperAppliance**](CredentialsApi.md#CreateCloudCredsHelperAppliance) | **Post** /api/v1/cloudCredentials/{id}/helperAppliances | Add or Edit Helper Appliance
[**CreateCreds**](CredentialsApi.md#CreateCreds) | **Post** /api/v1/credentials | Add Credentials Record
[**DeleteCloudCreds**](CredentialsApi.md#DeleteCloudCreds) | **Delete** /api/v1/cloudCredentials/{id} | Remove Cloud Credentials Record
[**DeleteCloudCredsHelperAppliance**](CredentialsApi.md#DeleteCloudCredsHelperAppliance) | **Delete** /api/v1/cloudCredentials/{id}/helperAppliances/{applianceId} | Remove Helper Appliance
[**DeleteCreds**](CredentialsApi.md#DeleteCreds) | **Delete** /api/v1/credentials/{id} | Remove Credentials Record
[**FinishAppRegistrationByDeviceCode**](CredentialsApi.md#FinishAppRegistrationByDeviceCode) | **Post** /api/v1/cloudCredentials/appRegistration/{verificationCode} | Register Azure AD Application
[**GetAllCloudCreds**](CredentialsApi.md#GetAllCloudCreds) | **Get** /api/v1/cloudCredentials | Get All Cloud Credentials
[**GetAllCreds**](CredentialsApi.md#GetAllCreds) | **Get** /api/v1/credentials | Get All Credentials
[**GetAllCredsHelperAppliances**](CredentialsApi.md#GetAllCredsHelperAppliances) | **Get** /api/v1/cloudCredentials/{id}/helperAppliances | Get All Helper Appliances
[**GetCloudCreds**](CredentialsApi.md#GetCloudCreds) | **Get** /api/v1/cloudCredentials/{id} | Get Cloud Credentials Record
[**GetCloudCredsHelperAppliance**](CredentialsApi.md#GetCloudCredsHelperAppliance) | **Get** /api/v1/cloudCredentials/{id}/helperAppliances/{applianceId} | Get Helper Appliance
[**GetCreds**](CredentialsApi.md#GetCreds) | **Get** /api/v1/credentials/{id} | Get Credentials Record
[**RequestAppRegistrationByDeviceCode**](CredentialsApi.md#RequestAppRegistrationByDeviceCode) | **Post** /api/v1/cloudCredentials/appRegistration | Get Verification Code
[**UpdateCloudCreds**](CredentialsApi.md#UpdateCloudCreds) | **Put** /api/v1/cloudCredentials/{id} | Edit Cloud Credentials Record
[**UpdateCreds**](CredentialsApi.md#UpdateCreds) | **Put** /api/v1/credentials/{id} | Edit Credentials Record



## ChangeCloudCertificate

> CloudCredentialsModel ChangeCloudCertificate(ctx, id).XApiVersion(xApiVersion).CertificateUploadSpec(certificateUploadSpec).Execute()

Change Certificate



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Azure compute account.
    certificateUploadSpec := *client.NewCertificateUploadSpec("Certificate_example", client.ECertificateFileFormatType("pfx")) // CertificateUploadSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.ChangeCloudCertificate(context.Background(), id).XApiVersion(xApiVersion).CertificateUploadSpec(certificateUploadSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.ChangeCloudCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeCloudCertificate`: CloudCredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.ChangeCloudCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Azure compute account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeCloudCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **certificateUploadSpec** | [**CertificateUploadSpec**](CertificateUploadSpec.md) |  | 

### Return type

[**CloudCredentialsModel**](CloudCredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeCloudCredsSecretKey

> CloudCredentialsModel ChangeCloudCredsSecretKey(ctx, id).XApiVersion(xApiVersion).CloudCredentialsPasswordSpec(cloudCredentialsPasswordSpec).Execute()

Change Secret Key



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the cloud credentials record.
    cloudCredentialsPasswordSpec := *client.NewCloudCredentialsPasswordSpec("NewSecretKey_example") // CloudCredentialsPasswordSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.ChangeCloudCredsSecretKey(context.Background(), id).XApiVersion(xApiVersion).CloudCredentialsPasswordSpec(cloudCredentialsPasswordSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.ChangeCloudCredsSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeCloudCredsSecretKey`: CloudCredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.ChangeCloudCredsSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the cloud credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeCloudCredsSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **cloudCredentialsPasswordSpec** | [**CloudCredentialsPasswordSpec**](CloudCredentialsPasswordSpec.md) |  | 

### Return type

[**CloudCredentialsModel**](CloudCredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the credentials record.
    credentialsPasswordChangeSpec := *client.NewCredentialsPasswordChangeSpec("Password_example") // CredentialsPasswordChangeSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.ChangePasswordForCreds(context.Background(), id).XApiVersion(xApiVersion).CredentialsPasswordChangeSpec(credentialsPasswordChangeSpec).Execute()
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
**id** | **string** | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordForCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the credentials record.
    privateKeyChangeSpec := *client.NewPrivateKeyChangeSpec("PrivateKey_example") // PrivateKeyChangeSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.ChangePrivateKeyForCreds(context.Background(), id).XApiVersion(xApiVersion).PrivateKeyChangeSpec(privateKeyChangeSpec).Execute()
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
**id** | **string** | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePrivateKeyForCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the credentials record.
    credentialsPasswordChangeSpec := *client.NewCredentialsPasswordChangeSpec("Password_example") // CredentialsPasswordChangeSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.ChangeRootPasswordForCreds(context.Background(), id).XApiVersion(xApiVersion).CredentialsPasswordChangeSpec(credentialsPasswordChangeSpec).Execute()
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
**id** | **string** | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRootPasswordForCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

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


## CreateCloudCreds

> CloudCredentialsModel CreateCloudCreds(ctx).XApiVersion(xApiVersion).CloudCredentialsSpec(cloudCredentialsSpec).Execute()

Add Cloud Credentials Record



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
    cloudCredentialsSpec := *client.NewCloudCredentialsSpec(client.ECloudCredentialsType("AzureStorage")) // CloudCredentialsSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.CreateCloudCreds(context.Background()).XApiVersion(xApiVersion).CloudCredentialsSpec(cloudCredentialsSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.CreateCloudCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloudCreds`: CloudCredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.CreateCloudCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **cloudCredentialsSpec** | [**CloudCredentialsSpec**](CloudCredentialsSpec.md) |  | 

### Return type

[**CloudCredentialsModel**](CloudCredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudCredsHelperAppliance

> SessionModel CreateCloudCredsHelperAppliance(ctx, id).XApiVersion(xApiVersion).CloudHelperApplianceSpec(cloudHelperApplianceSpec).Execute()

Add or Edit Helper Appliance



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Microsoft Azure compute account.
    cloudHelperApplianceSpec := *client.NewCloudHelperApplianceSpec(client.ECloudCredentialsType("AzureStorage")) // CloudHelperApplianceSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.CreateCloudCredsHelperAppliance(context.Background(), id).XApiVersion(xApiVersion).CloudHelperApplianceSpec(cloudHelperApplianceSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.CreateCloudCredsHelperAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloudCredsHelperAppliance`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.CreateCloudCredsHelperAppliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Microsoft Azure compute account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudCredsHelperApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **cloudHelperApplianceSpec** | [**CloudHelperApplianceSpec**](CloudHelperApplianceSpec.md) |  | 

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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    credentialsSpec := *client.NewCredentialsSpec("Username_example", client.ECredentialsType("Standard")) // CredentialsSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.CreateCreds(context.Background()).XApiVersion(xApiVersion).CredentialsSpec(credentialsSpec).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
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


## DeleteCloudCreds

> map[string]interface{} DeleteCloudCreds(ctx, id).XApiVersion(xApiVersion).Execute()

Remove Cloud Credentials Record



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the cloud credentials record.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.DeleteCloudCreds(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.DeleteCloudCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCloudCreds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.DeleteCloudCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the cloud credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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


## DeleteCloudCredsHelperAppliance

> map[string]interface{} DeleteCloudCredsHelperAppliance(ctx, id, applianceId).XApiVersion(xApiVersion).Execute()

Remove Helper Appliance



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Microsoft Azure compute account.
    applianceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the helper appliance.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.DeleteCloudCredsHelperAppliance(context.Background(), id, applianceId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.DeleteCloudCredsHelperAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCloudCredsHelperAppliance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.DeleteCloudCredsHelperAppliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Microsoft Azure compute account. | 
**applianceId** | **string** | ID of the helper appliance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudCredsHelperApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]



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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the credentials record.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.DeleteCreds(context.Background(), id).XApiVersion(xApiVersion).Execute()
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
**id** | **string** | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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


## FinishAppRegistrationByDeviceCode

> CloudNativeApplicationModel FinishAppRegistrationByDeviceCode(ctx, verificationCode).XApiVersion(xApiVersion).Execute()

Register Azure AD Application



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
    verificationCode := "verificationCode_example" // string | Verification code. To obtain the code, use the [Get Verification Code](#tag/Credentials/operation/RequestAppRegistrationByDeviceCode) request.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.FinishAppRegistrationByDeviceCode(context.Background(), verificationCode).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.FinishAppRegistrationByDeviceCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinishAppRegistrationByDeviceCode`: CloudNativeApplicationModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.FinishAppRegistrationByDeviceCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verificationCode** | **string** | Verification code. To obtain the code, use the [Get Verification Code](#tag/Credentials/operation/RequestAppRegistrationByDeviceCode) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishAppRegistrationByDeviceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**CloudNativeApplicationModel**](CloudNativeApplicationModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCloudCreds

> CloudCredentialsResult GetAllCloudCreds(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).Execute()

Get All Cloud Credentials



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
    skip := int32(56) // int32 | Number of cloud credentials records to skip. (optional)
    limit := int32(56) // int32 | Maximum number of cloud credentials records to return. (optional)
    orderColumn := client.ECloudCredentialsFiltersOrderColumn("Name") // ECloudCredentialsFiltersOrderColumn | Sorts cloud credentials by one of the cloud credentials parameters. (optional)
    orderAsc := true // bool | Sorts cloud credentials in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters cloud credentials by the `nameFilter` pattern. The pattern can match any cloud credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := client.ECloudCredentialsType("AzureStorage") // ECloudCredentialsType |  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.GetAllCloudCreds(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.GetAllCloudCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCloudCreds`: CloudCredentialsResult
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.GetAllCloudCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCloudCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of cloud credentials records to skip. | 
 **limit** | **int32** | Maximum number of cloud credentials records to return. | 
 **orderColumn** | [**ECloudCredentialsFiltersOrderColumn**](ECloudCredentialsFiltersOrderColumn.md) | Sorts cloud credentials by one of the cloud credentials parameters. | 
 **orderAsc** | **bool** | Sorts cloud credentials in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters cloud credentials by the &#x60;nameFilter&#x60; pattern. The pattern can match any cloud credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **typeFilter** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 

### Return type

[**CloudCredentialsResult**](CloudCredentialsResult.md)

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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    skip := int32(56) // int32 | Number of credentials records to skip. (optional)
    limit := int32(56) // int32 | Maximum number of credentials records to return. (optional)
    orderColumn := client.ECredentialsFiltersOrderColumn("Username") // ECredentialsFiltersOrderColumn | Sorts credentials by one of the credentials parameters. (optional)
    orderAsc := true // bool | Sorts credentials in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters credentials by the `nameFilter` pattern. The pattern can match any credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.GetAllCreds(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
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


## GetAllCredsHelperAppliances

> CloudHelperApplianceResult GetAllCredsHelperAppliances(ctx, id).XApiVersion(xApiVersion).Execute()

Get All Helper Appliances



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Microsoft Azure compute account.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.GetAllCredsHelperAppliances(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.GetAllCredsHelperAppliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCredsHelperAppliances`: CloudHelperApplianceResult
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.GetAllCredsHelperAppliances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Microsoft Azure compute account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCredsHelperAppliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**CloudHelperApplianceResult**](CloudHelperApplianceResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudCreds

> CloudCredentialsModel GetCloudCreds(ctx, id).XApiVersion(xApiVersion).Execute()

Get Cloud Credentials Record



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the cloud credentials record.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.GetCloudCreds(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.GetCloudCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudCreds`: CloudCredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.GetCloudCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the cloud credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**CloudCredentialsModel**](CloudCredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudCredsHelperAppliance

> CloudHelperApplianceModel GetCloudCredsHelperAppliance(ctx, id, applianceId).XApiVersion(xApiVersion).Execute()

Get Helper Appliance



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Microsoft Azure compute account.
    applianceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the helper appliance.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.GetCloudCredsHelperAppliance(context.Background(), id, applianceId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.GetCloudCredsHelperAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudCredsHelperAppliance`: CloudHelperApplianceModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.GetCloudCredsHelperAppliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Microsoft Azure compute account. | 
**applianceId** | **string** | ID of the helper appliance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudCredsHelperApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]



### Return type

[**CloudHelperApplianceModel**](CloudHelperApplianceModel.md)

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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the credentials record.

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.GetCreds(context.Background(), id).XApiVersion(xApiVersion).Execute()
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
**id** | **string** | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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


## RequestAppRegistrationByDeviceCode

> CloudDeviceCodeModel RequestAppRegistrationByDeviceCode(ctx).XApiVersion(xApiVersion).CloudDeviceCodeSpec(cloudDeviceCodeSpec).Execute()

Get Verification Code



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
    cloudDeviceCodeSpec := *client.NewCloudDeviceCodeSpec(client.ECloudCredentialsType("AzureStorage")) // CloudDeviceCodeSpec |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.RequestAppRegistrationByDeviceCode(context.Background()).XApiVersion(xApiVersion).CloudDeviceCodeSpec(cloudDeviceCodeSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.RequestAppRegistrationByDeviceCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestAppRegistrationByDeviceCode`: CloudDeviceCodeModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.RequestAppRegistrationByDeviceCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAppRegistrationByDeviceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **cloudDeviceCodeSpec** | [**CloudDeviceCodeSpec**](CloudDeviceCodeSpec.md) |  | 

### Return type

[**CloudDeviceCodeModel**](CloudDeviceCodeModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudCreds

> CloudCredentialsModel UpdateCloudCreds(ctx, id).XApiVersion(xApiVersion).CloudCredentialsModel(cloudCredentialsModel).Execute()

Edit Cloud Credentials Record



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the cloud credentials record.
    cloudCredentialsModel := *client.NewCloudCredentialsModel("Id_example", client.ECloudCredentialsType("AzureStorage")) // CloudCredentialsModel |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.UpdateCloudCreds(context.Background(), id).XApiVersion(xApiVersion).CloudCredentialsModel(cloudCredentialsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApi.UpdateCloudCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudCreds`: CloudCredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApi.UpdateCloudCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the cloud credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **cloudCredentialsModel** | [**CloudCredentialsModel**](CloudCredentialsModel.md) |  | 

### Return type

[**CloudCredentialsModel**](CloudCredentialsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
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
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the credentials record.
    credentialsModel := *client.NewCredentialsModel("Id_example", "Username_example", "Description_example", client.ECredentialsType("Standard"), time.Now()) // CredentialsModel |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApi.UpdateCreds(context.Background(), id).XApiVersion(xApiVersion).CredentialsModel(credentialsModel).Execute()
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
**id** | **string** | ID of the credentials record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

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

