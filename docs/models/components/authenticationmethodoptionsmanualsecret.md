# AuthenticationMethodOptionsManualSecret

Enter client secret directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsManualSecretManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsManualSecret("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AuthenticationMethodOptionsManualSecretManual` | manual                                          |
| `AuthenticationMethodOptionsManualSecretSecret` | secret                                          |