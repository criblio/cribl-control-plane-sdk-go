# InputEventhubAmqpAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputEventhubAmqpAuthenticationMethodSecret

// Open enum: custom values can be created with a direct type cast
custom := components.InputEventhubAmqpAuthenticationMethod("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `InputEventhubAmqpAuthenticationMethodSecret`             | secret                                                    |
| `InputEventhubAmqpAuthenticationMethodClientSecret`       | clientSecret                                              |
| `InputEventhubAmqpAuthenticationMethodClientCert`         | clientCert                                                |
| `InputEventhubAmqpAuthenticationMethodClientAssertion`    | clientAssertion                                           |
| `InputEventhubAmqpAuthenticationMethodClientAssertionRPC` | clientAssertion_rpc                                       |