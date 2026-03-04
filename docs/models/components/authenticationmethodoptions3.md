# AuthenticationMethodOptions3

Enter API key directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptions3Manual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptions3("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AuthenticationMethodOptions3Manual` | manual                               |
| `AuthenticationMethodOptions3Secret` | secret                               |