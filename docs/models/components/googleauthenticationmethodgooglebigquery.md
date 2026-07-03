# GoogleAuthenticationMethodGoogleBigquery

Choose Auto to use Google Application Default Credentials (ADC), or Secret to select or create a stored secret that references Google service account credentials

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.GoogleAuthenticationMethodGoogleBigqueryAuto

// Open enum: custom values can be created with a direct type cast
custom := components.GoogleAuthenticationMethodGoogleBigquery("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `GoogleAuthenticationMethodGoogleBigqueryAuto`   | auto                                             |
| `GoogleAuthenticationMethodGoogleBigquerySecret` | secret                                           |