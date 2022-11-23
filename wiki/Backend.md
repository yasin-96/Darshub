# Microsoft Praph Go SDK

## Install the Microsoft Graph API

    go get github.com/microsoftgraph/msgraph-sdk-go
    go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
    go get github.com/microsoft/kiota-authentication-azure-go


## Authentication and authorization basics
- [Basics](https://learn.microsoft.com/en-us/graph/auth/auth-concepts)
- [Reigster your App](https://learn.microsoft.com/en-us/graph/auth-register-app-v2)
- [Get Access on behalf of a user](https://learn.microsoft.com/en-us/graph/auth-v2-user)
- [Get access without a user](https://learn.microsoft.com/en-us/graph/auth-v2-service)
- [Permissions](https://learn.microsoft.com/en-us/graph/permissions-reference)
- [Manage App access (CSP)](https://learn.microsoft.com/en-us/graph/auth-cloudsolutionprovider?tabs=azuread)

## Create a Microsoft Graph client

The following code examples show how to create an instance of a Microsoft Graph client with an <br>
authentication provider in Go. he authentication provider will handle acquiring access tokens for <br>
the application.

``` go
import (
    "context"
    "fmt"

    azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    a "github.com/microsoft/kiota-authentication-azure-go"
    msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
    ClientID: "CLIENT_ID",
    UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
        fmt.Println(message.Message)
        return nil
    },
})

if err != nil {
    fmt.Printf("Error creating credentials: %v\n", err)
    return
}

auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"User.Read"})

if err != nil {
    fmt.Printf("Error authentication provider: %v\n", err)
    return
}

adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
if err != nil {
    fmt.Printf("Error creating adapter: %v\n", err)
    return
`
client := msgraphsdk.NewGraphServiceClient(adapter)
```

## Customize the Microsoft Graph SDK service client

The Microsoft Graph SDK client configures a default set of middleware that allows the SDK to <br>
communicate with the Microsoft Graph endpoints. This default set is customizable, allowing you to <br>
change the behavior of the client.

```go
import (
    a "github.com/microsoft/kiota-authentication-azure-go"
    khttp "github.com/microsoft/kiota-http-go"
    msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
    core "github.com/microsoftgraph/msgraph-sdk-go-core"
)

// Auth provider
a   uth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(...)

// Get default middleware from SDK
defaultMiddleware := core.GetDefaultMiddlewaresWithOptions(msgraphsdk.GetDefaultClientOptions())

// Get instance of custom middleware
// Implement a custom middleware by implementing the Middleware interface
// https://github.com/microsoft/kiota-http-go/blob/main/middleware.go
allMiddleware := append(defaultMiddleware, mycustom.NewCustomHandler())

// Create an HTTP client with the middleware
httpClient := khttp.GetDefaultClient(allMiddleware...)

// Create the adapter
// Passing nil values causes the adapter to use default implementations
adapter, err :=
    msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
        auth, nil, nil, httpClient)

client := msgraphsdk.NewGraphServiceClient(adapter)
``` go