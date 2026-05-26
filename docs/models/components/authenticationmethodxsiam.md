# AuthenticationMethodXsiam

Enter a token directly, or provide a secret referencing a token

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodXsiamToken

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodXsiam("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `AuthenticationMethodXsiamToken`  | token                             |
| `AuthenticationMethodXsiamSecret` | secret                            |