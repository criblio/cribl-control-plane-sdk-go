# CreateInputAuthAuthenticationMethod

Enter connection string directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthAuthenticationMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateInputAuthAuthenticationMethodManual` | manual                                      |
| `CreateInputAuthAuthenticationMethodSecret` | secret                                      |