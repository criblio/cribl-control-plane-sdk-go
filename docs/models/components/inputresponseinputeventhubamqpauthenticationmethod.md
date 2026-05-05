# InputResponseInputEventhubAmqpAuthenticationMethod

Enter connection string directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputEventhubAmqpAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputEventhubAmqpAuthenticationMethod("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `InputResponseInputEventhubAmqpAuthenticationMethodManual` | manual                                                     |
| `InputResponseInputEventhubAmqpAuthenticationMethodSecret` | secret                                                     |