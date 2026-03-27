# GoogleAuthenticationMethodOptions

Choose Auto to use Google Application Default Credentials (ADC), Manual to enter Google service account credentials directly, or Secret to select or create a stored secret that references Google service account credentials.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.GoogleAuthenticationMethodOptionsAuto

// Open enum: custom values can be created with a direct type cast
custom := components.GoogleAuthenticationMethodOptions("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `GoogleAuthenticationMethodOptionsAuto`   | auto                                      |
| `GoogleAuthenticationMethodOptionsManual` | manual                                    |
| `GoogleAuthenticationMethodOptionsSecret` | secret                                    |