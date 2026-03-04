# AuthenticationMethodOptionsSasl1

Enter password directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsSasl1Manual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsSasl1("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AuthenticationMethodOptionsSasl1Manual` | manual                                   |
| `AuthenticationMethodOptionsSasl1Secret` | secret                                   |