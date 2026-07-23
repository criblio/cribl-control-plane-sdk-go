# OutputGoogleBigqueryGoogleAuthenticationMethod

Choose Auto to use Google Application Default Credentials (ADC), or Secret to select or create a stored secret that references Google service account credentials

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputGoogleBigqueryGoogleAuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputGoogleBigqueryGoogleAuthenticationMethod("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `OutputGoogleBigqueryGoogleAuthenticationMethodAuto`   | auto                                                   |
| `OutputGoogleBigqueryGoogleAuthenticationMethodSecret` | secret                                                 |