# AuthenticationMethodGoogleChronicle

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodGoogleChronicleManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodGoogleChronicle("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `AuthenticationMethodGoogleChronicleManual`               | manual                                                    |
| `AuthenticationMethodGoogleChronicleSecret`               | secret                                                    |
| `AuthenticationMethodGoogleChronicleServiceAccount`       | serviceAccount                                            |
| `AuthenticationMethodGoogleChronicleServiceAccountSecret` | serviceAccountSecret                                      |