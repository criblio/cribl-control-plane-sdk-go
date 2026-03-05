# CreateInputAuthenticationMethodEdgePrometheus

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationMethodEdgePrometheusManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationMethodEdgePrometheus("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateInputAuthenticationMethodEdgePrometheusManual`     | manual                                                    |
| `CreateInputAuthenticationMethodEdgePrometheusSecret`     | secret                                                    |
| `CreateInputAuthenticationMethodEdgePrometheusKubernetes` | kubernetes                                                |