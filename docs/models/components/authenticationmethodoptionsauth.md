# AuthenticationMethodOptionsAuth

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsAuthManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsAuth("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AuthenticationMethodOptionsAuthManual`       | manual                                        |
| `AuthenticationMethodOptionsAuthSecret`       | secret                                        |
| `AuthenticationMethodOptionsAuthManualAPIKey` | manualAPIKey                                  |
| `AuthenticationMethodOptionsAuthTextSecret`   | textSecret                                    |