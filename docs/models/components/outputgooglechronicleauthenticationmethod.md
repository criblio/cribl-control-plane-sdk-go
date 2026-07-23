# OutputGoogleChronicleAuthenticationMethod

Authentication method

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputGoogleChronicleAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.OutputGoogleChronicleAuthenticationMethod("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `OutputGoogleChronicleAuthenticationMethodManual`               | manual                                                          |
| `OutputGoogleChronicleAuthenticationMethodSecret`               | secret                                                          |
| `OutputGoogleChronicleAuthenticationMethodServiceAccount`       | serviceAccount                                                  |
| `OutputGoogleChronicleAuthenticationMethodServiceAccountSecret` | serviceAccountSecret                                            |