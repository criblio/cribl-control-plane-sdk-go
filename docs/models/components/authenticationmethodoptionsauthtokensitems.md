# AuthenticationMethodOptionsAuthTokensItems

Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsAuthTokensItemsManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsAuthTokensItems("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `AuthenticationMethodOptionsAuthTokensItemsManual` | manual                                             |
| `AuthenticationMethodOptionsAuthTokensItemsSecret` | secret                                             |