# AuthenticationMethodOptionsAPI

Enter API key directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsAPIManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsAPI("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AuthenticationMethodOptionsAPIManual` | manual                                 |
| `AuthenticationMethodOptionsAPISecret` | secret                                 |