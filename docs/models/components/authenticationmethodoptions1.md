# AuthenticationMethodOptions1

Enter client secret directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptions1Manual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptions1("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AuthenticationMethodOptions1Manual` | manual                               |
| `AuthenticationMethodOptions1Secret` | secret                               |