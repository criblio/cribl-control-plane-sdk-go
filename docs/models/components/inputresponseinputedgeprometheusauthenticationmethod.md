# InputResponseInputEdgePrometheusAuthenticationMethod

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputEdgePrometheusAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputEdgePrometheusAuthenticationMethod("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `InputResponseInputEdgePrometheusAuthenticationMethodManual`     | manual                                                           |
| `InputResponseInputEdgePrometheusAuthenticationMethodSecret`     | secret                                                           |
| `InputResponseInputEdgePrometheusAuthenticationMethodKubernetes` | kubernetes                                                       |