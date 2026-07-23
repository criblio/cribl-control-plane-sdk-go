# CreateInputSystemByPackAuthenticationMethodEdgePrometheus

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthenticationMethodEdgePrometheusManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthenticationMethodEdgePrometheus("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CreateInputSystemByPackAuthenticationMethodEdgePrometheusManual`     | manual                                                                |
| `CreateInputSystemByPackAuthenticationMethodEdgePrometheusSecret`     | secret                                                                |
| `CreateInputSystemByPackAuthenticationMethodEdgePrometheusKubernetes` | kubernetes                                                            |