# OutputGoogleCloudObservabilityGoogleAuthenticationMethod

Choose Auto to use Google Application Default Credentials (ADC). Choose Secret to select or create a stored secret that references Google service account credentials.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputGoogleCloudObservabilityGoogleAuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputGoogleCloudObservabilityGoogleAuthenticationMethod("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `OutputGoogleCloudObservabilityGoogleAuthenticationMethodAuto`   | auto                                                             |
| `OutputGoogleCloudObservabilityGoogleAuthenticationMethodSecret` | secret                                                           |