# AzureBlobAuthTypeSecretAuthenticationMethod

Enter authentication data directly, or select a secret referencing your auth data

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AzureBlobAuthTypeSecretAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.AzureBlobAuthTypeSecretAuthenticationMethod("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `AzureBlobAuthTypeSecretAuthenticationMethodManual`       | manual                                                    |
| `AzureBlobAuthTypeSecretAuthenticationMethodSecret`       | secret                                                    |
| `AzureBlobAuthTypeSecretAuthenticationMethodClientSecret` | clientSecret                                              |
| `AzureBlobAuthTypeSecretAuthenticationMethodClientCert`   | clientCert                                                |