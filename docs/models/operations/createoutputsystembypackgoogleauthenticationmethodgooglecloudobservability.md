# CreateOutputSystemByPackGoogleAuthenticationMethodGoogleCloudObservability

Choose Auto to use Google Application Default Credentials (ADC). Choose Secret to select or create a stored secret that references Google service account credentials.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackGoogleAuthenticationMethodGoogleCloudObservabilityAuto

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackGoogleAuthenticationMethodGoogleCloudObservability("custom_value")
```


## Values

| Name                                                                               | Value                                                                              |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `CreateOutputSystemByPackGoogleAuthenticationMethodGoogleCloudObservabilityAuto`   | auto                                                                               |
| `CreateOutputSystemByPackGoogleAuthenticationMethodGoogleCloudObservabilitySecret` | secret                                                                             |