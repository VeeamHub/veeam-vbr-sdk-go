# Deprecation

> [!WARNING]
> The master branch is deprecated now. Please use the [v2](https://github.com/VeeamHub/veeam-vbr-sdk-go/tree/v2) branch instead.

# Veeam Backup & Replication SDK for Go

veeam-vbr-sdk-go is the unofficial Veeam Backup & Replication SDK for the Go programming language.

This SDK was generated using the [OpenAPI Generator](https://openapi-generator.tech/). As such, any API call in the [Veeam Backup & Replication v11 REST API](https://helpcenter.veeam.com/docs/backup/vbr_rest/overview.html?ver=110) is present in this SDK.

## 📗 Documentation

### Installation

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

```bash
go get github.com/veeamhub/veeam-vbr-sdk-go
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```bash
go get -u github.com/veeamhub/veeam-vbr-sdk-go
```

#### Dependencies

The SDK includes a `vendor` folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the Go module file `go.mod`.

### Example

This example shows a complete working Go file which will list the names of all Backup Jobs and demonstrate how to configure the timeout logic that will cancel the request if it takes too long. This example highlights how to login to a VBR server with trusted or self-signed certs, make a request, handle the error, process the response, and then logout.

```go
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/veeamhub/veeam-vbr-sdk-go/client"
)

// Retrieves and prints out all Backup Job names of the specified VBR server.
//
// The Variables below need to be set according to your environment.
func main() {
	// Setting variables
	host := "vbr.contoso.local:9419"  // default API port 9419
	username := "contoso\\jsmith"     // VBR username
	password = "password"             // VBR password

	// Setting constants
	const (
		apiVersion = "1.0-rev1"       // default API version (1.0-rev1)
		skipTls    = true             // skip TLS certificate verification
		timeout    = 30 * time.Second // 30 seconds
	)

	// Generating API client
	tlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTls,
			},
		},
	}
	config := client.NewConfiguration()
	config.HTTPClient = tlsClient
	config.Host = host
	config.HTTPClient.Timeout = timeout
	config.Scheme = "https"
	veeam := client.NewAPIClient(config)

	// Authenticating to VBR API
	login, r, err := veeam.LoginApi.CreateToken(context.Background()).XApiVersion(apiVersion).GrantType("password").Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	// Setting authorization
	// From this point, all calls will 'auth' as the context
	auth := context.WithValue(context.Background(), client.ContextAccessToken,
		login.AccessToken)

	// Retrieving all backup jobs
	jobs, r, err := veeam.JobsApi.GetAllJobs(auth).XApiVersion(apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	// Printing out job names
	// Working with the response payload
	fmt.Printf("Job Names:\n\n")
	for _, job := range jobs.Data {
		fmt.Println(job.Name)
	}

	// Logging out session
	_, r, err = veeam.LoginApi.Logout(auth).XApiVersion(apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}
}

```

## ✍ Contributions

We welcome contributions from the community! We encourage you to create [issues](https://github.com/VeeamHub/veeam-vbr-sdk-go/issues/new/choose) for Bugs & Feature Requests and submit Pull Requests. For more detailed information, refer to our [Contributing Guide](CONTRIBUTING.md).

## 🤝🏾 License

* [MIT License](LICENSE)

## 🤔 Questions

If you have any questions or something is unclear, please don't hesitate to [create an issue](https://github.com/VeeamHub/veeam-vbr-sdk-go/issues/new/choose) and let us know!
