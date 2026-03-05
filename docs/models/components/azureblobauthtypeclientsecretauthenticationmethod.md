# AzureBlobAuthTypeClientSecretAuthenticationMethod

Enter authentication data directly, or select a secret referencing your auth data

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AzureBlobAuthTypeClientSecretAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.AzureBlobAuthTypeClientSecretAuthenticationMethod("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AzureBlobAuthTypeClientSecretAuthenticationMethodManual`       | manual                                                          |
| `AzureBlobAuthTypeClientSecretAuthenticationMethodSecret`       | secret                                                          |
| `AzureBlobAuthTypeClientSecretAuthenticationMethodClientSecret` | clientSecret                                                    |
| `AzureBlobAuthTypeClientSecretAuthenticationMethodClientCert`   | clientCert                                                      |