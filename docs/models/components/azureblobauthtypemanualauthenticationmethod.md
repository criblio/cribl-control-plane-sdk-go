# AzureBlobAuthTypeManualAuthenticationMethod

Enter authentication data directly, or select a secret referencing your auth data

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AzureBlobAuthTypeManualAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.AzureBlobAuthTypeManualAuthenticationMethod("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `AzureBlobAuthTypeManualAuthenticationMethodManual`       | manual                                                    |
| `AzureBlobAuthTypeManualAuthenticationMethodSecret`       | secret                                                    |
| `AzureBlobAuthTypeManualAuthenticationMethodClientSecret` | clientSecret                                              |
| `AzureBlobAuthTypeManualAuthenticationMethodClientCert`   | clientCert                                                |