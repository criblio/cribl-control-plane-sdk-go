# AuthenticationMethodEventhubAmqp

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodEventhubAmqpSecret

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodEventhubAmqp("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `AuthenticationMethodEventhubAmqpSecret`             | secret                                               |
| `AuthenticationMethodEventhubAmqpClientSecret`       | clientSecret                                         |
| `AuthenticationMethodEventhubAmqpClientCert`         | clientCert                                           |
| `AuthenticationMethodEventhubAmqpClientAssertion`    | clientAssertion                                      |
| `AuthenticationMethodEventhubAmqpClientAssertionRPC` | clientAssertion_rpc                                  |