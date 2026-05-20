# AuthenticationMethodOptionsAuthManualManualAPIKey

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsAuthManualManualAPIKeyManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsAuthManualManualAPIKey("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AuthenticationMethodOptionsAuthManualManualAPIKeyManual`       | manual                                                          |
| `AuthenticationMethodOptionsAuthManualManualAPIKeySecret`       | secret                                                          |
| `AuthenticationMethodOptionsAuthManualManualAPIKeyManualAPIKey` | manualAPIKey                                                    |
| `AuthenticationMethodOptionsAuthManualManualAPIKeyTextSecret`   | textSecret                                                      |