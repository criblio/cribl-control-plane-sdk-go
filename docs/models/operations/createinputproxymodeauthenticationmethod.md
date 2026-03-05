# CreateInputProxyModeAuthenticationMethod

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputProxyModeAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputProxyModeAuthenticationMethod("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateInputProxyModeAuthenticationMethodNone`   | none                                             |
| `CreateInputProxyModeAuthenticationMethodManual` | manual                                           |
| `CreateInputProxyModeAuthenticationMethodSecret` | secret                                           |