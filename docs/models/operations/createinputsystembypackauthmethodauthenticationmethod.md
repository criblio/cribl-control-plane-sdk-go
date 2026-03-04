# CreateInputSystemByPackAuthMethodAuthenticationMethod

How to authenticate incoming client connections

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthMethodAuthenticationMethodClientCert

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthMethodAuthenticationMethod("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CreateInputSystemByPackAuthMethodAuthenticationMethodClientCert` | clientCert                                                        |
| `CreateInputSystemByPackAuthMethodAuthenticationMethodKerberos`   | kerberos                                                          |