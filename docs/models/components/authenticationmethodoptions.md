# AuthenticationMethodOptions

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptions("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AuthenticationMethodOptionsManual`             | manual                                          |
| `AuthenticationMethodOptionsSecret`             | secret                                          |
| `AuthenticationMethodOptionsClientSecret`       | clientSecret                                    |
| `AuthenticationMethodOptionsClientCert`         | clientCert                                      |
| `AuthenticationMethodOptionsClientAssertion`    | clientAssertion                                 |
| `AuthenticationMethodOptionsClientAssertionRPC` | clientAssertion_rpc                             |