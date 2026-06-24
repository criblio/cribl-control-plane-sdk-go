# OutputResponseGoogleAuthenticationMethod

Choose Auto to use Google Application Default Credentials (ADC). Choose Secret to select or create a stored secret that references Google service account credentials.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseGoogleAuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseGoogleAuthenticationMethod("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `OutputResponseGoogleAuthenticationMethodAuto`   | auto                                             |
| `OutputResponseGoogleAuthenticationMethodSecret` | secret                                           |