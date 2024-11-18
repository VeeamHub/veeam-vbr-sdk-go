# Golang client for VBR REST API

[![Go Report Card](https://goreportcard.com/badge/github.com/VeeamHub/veeam-vbr-sdk-go/v2)](https://goreportcard.com/report/github.com/VeeamHub/veeam-vbr-sdk-go/v2)
[![GoDoc](https://godoc.org/github.com/VeeamHub/veeam-vbr-sdk-go/v2?status.svg)](https://godoc.org/github.com/VeeamHub/veeam-vbr-sdk-go/v2)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/license/mit/)

veeam-vbr-sdk-go is the unofficial Veeam Backup & Replication SDK for the Go programming language.
Client generation is based on https://github.com/oapi-codegen/oapi-codegen generator.
Due to specific of Golang, we had to make some changes in the original specification to make it work with the generator.
You can find changes description in the [Specification](#specification) section.

## 📁 Project layout
* `spec` - contains specification files. Both original `VBR REST API` specification and generated `openapi_spec.yaml` are placed here.
* `tools` - contains additional tools. Currently, it contains only `oapifixer` tool which is used to apply required changes to the original specification.
* `pkg/client` - contains generated client. It is not recommended to change anything in this directory. If you want to change something, please, change the specification and regenerate the client.

## 🔍 Specification
`openapi_spec.yaml` does not contain the whole original `VBR REST API` specification. We made several changes described below. 

### Changes in spec
For each schema which has both `OneOf` and `Properties` following changes were made:
* Properties were moved to separate schema with the same name as original schema but with `Base` prefix, e.g. `RepositoryModel` -> `BaseRepositoryModel`.
* `AllOf` was added to the original schema. It contains reference to the `Base` schema and `OneOf` from the original schema.
* All references to the original schema were replaced by reference to the `Base` schema.

### How to regenerate spec
Additional tool named `oapifixer` included in the project. You can find it in the `tools/oapifixer` directory. 
It applies all the changes described in the [Changes in spec](#changes-in-spec) section.
Tool integrated into the `Makefile` and can be used by the following command:
```bash
make convert
```
It will generate specification file which can be used to generate client. 
By default it expects that the original specification is placed in the `spec` directory and has the name `swagger.json`.
'Fixed' specification will be placed in the `spec` directory and will have the name `openapi_spec.yaml`.
To change the default behavior you can use the following command:
```bash
make convert vbr_spec=<path_to_original_vbr_specification> golang_spec=<path_to_result>
```

> It is possible to convert specification from JSON to YAML and vice versa. Just change the extension of the output files.


## 📝 How to generate code

To generate code just run the following command:
```bash
make generate
```
It will remove the previous version and generate the new one. The result of generation will be placed into the `pkg/client` directory.
The default value for specification is `./spec/openapi_spec.yaml`. To change it use the following command:
```bash
make generate golang_spec=<path_to_specification>
```

## 📗 Documentation

### Complete examples
You can find complete examples in the `pkg/client` directory, in the `example_test.go` file.
Examples made as a [Testable Examples](https://go.dev/blog/examples).

### Create client and authenticate
Create a new client via vbrapi.NewClient()

```golang
serverHost := "https://127.0.0.1:9398"
vbrClient := vbrclient.NewClientWithResponses(serverHost)
```

You can also provide your own `http.Client` to the constructor:
```golang
tlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
serverHost := "https://127.0.0.1:9398"
cl := vbrclient.NewClientWithResponses(serverHost, vbrclient.WithHTTPClient(tlsClient))
```

As a result you will get client which can make requests allowed to be made without authentication.
For example, you can get server info:
```golang
rsi, err := cl.GetServerInfoWithResponse(context.Background(), &vbrclient.GetServerInfoParams{
		XApiVersion: "1.1-rev0",
	})
if err != nil {
	panic(err)
}
log.Print(rsi.JSON200)
```

To be able to make authenticated requests you have to authenticate your client. 
Client uses bearer token authentication. So first of all you have to get the token:
```golang
user := "vbruser"
password := "vbrpassword"

rl, err := cl.CreateTokenWithFormdataBodyWithResponse(context.Background(), &vbrclient.CreateTokenParams{
		XApiVersion: "1.1-rev0",
	}, vbrclient.CreateTokenFormdataRequestBody{
		GrantType: "password",
		Username:  &user,
		Password:  &pass,
	})
if err != nil {
	panic(err)
}
```

Next, you need to create a new client with the bearer token provider:
```golang
bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(rl.JSON200.AccessToken)
if bearerTokenProviderErr != nil {
	panic(bearerTokenProviderErr)
}

serverHost := "https://127.0.0.1:9398"
authcl, err := vbrclient.NewClientWithResponses(serverHost, vbrclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
if err != nil {
	panic(err)
}
```

Now you can make authenticated requests:
```golang
nameFilter := "test"
rgr, err := authcl.GetAllRepositoriesWithResponse(context.Background(), &vbrclient.GetAllRepositoriesParams{
		XApiVersion: "1.1-rev0",
		NameFilter:  &nameFilter,
	})
if err != nil {
	panic(err)
}
for rmi := range rgr.JSON200.Data {
	rm := rgr.JSON200.Data[rmi]
	log.Info(rm)
}
```
### Making requests

There are 2 types of Client interfaces provided by the library
* Client which returns unparsed `*http.Response`
* Client which returns parsed model

So, if you're using client with parsed models you will have 2 functions for each `path` item. For example:
```golang
GetServerInfo(ctx context.Context, params *GetServerInfoParams, reqEditors ...RequestEditorFn) (*http.Response, error)
GetServerInfoWithResponse(ctx context.Context, params *GetServerInfoParams, reqEditors ...RequestEditorFn) (*GetServerInfoResponse, error)
```

GetServerInfoResponse contains parsed model:
```golang
type GetServerInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ServerInfoModel
	JSON401      *Error
	JSON403      *Error
	JSON500      *Error
}
```

As you can see there are several objects represents results for different HTTP codes.

Additionally, each response contains following functions to get HTTP Status and HTTP Code
```golang
// Status returns HTTPResponse.Status
func (r GetServerInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetServerInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
```

Please note that you should check status separately, the error will be returned only in case of failed request. If server returned an answer, no error will be returned.

## 🚦 Known Issues
* If you convert specification from `JSON` to `YAML` usiang `oapifixer` tool the order of sections in the specification will be changed(to alphabetical). It is not a problem for the generator but it makes it difficult to review changes in the specification.


## ✍ Contributions

We welcome contributions from the community! We encourage you to create [issues](https://github.com/VeeamHub/veeam-vbr-sdk-go/issues/new/choose) for Bugs & Feature Requests and submit Pull Requests. For more detailed information, refer to our [Contributing Guide](CONTRIBUTING.md).

## 🤝🏾 License

* [MIT License](LICENSE)

## 🤔 Questions

If you have any questions or something is unclear, please don't hesitate to [create an issue](https://github.com/VeeamHub/veeam-vbr-sdk-go/issues/new/choose) and let us know!
