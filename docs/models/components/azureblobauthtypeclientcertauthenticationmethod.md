# AzureBlobAuthTypeClientCertAuthenticationMethod

Enter authentication data directly, or select a secret referencing your auth data

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AzureBlobAuthTypeClientCertAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.AzureBlobAuthTypeClientCertAuthenticationMethod("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AzureBlobAuthTypeClientCertAuthenticationMethodManual`       | manual                                                        |
| `AzureBlobAuthTypeClientCertAuthenticationMethodSecret`       | secret                                                        |
| `AzureBlobAuthTypeClientCertAuthenticationMethodClientSecret` | clientSecret                                                  |
| `AzureBlobAuthTypeClientCertAuthenticationMethodClientCert`   | clientCert                                                    |