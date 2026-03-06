# AuthenticationMethodOptionsSasl

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsSaslManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsSasl("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AuthenticationMethodOptionsSaslManual` | manual                                  |
| `AuthenticationMethodOptionsSaslSecret` | secret                                  |