# CreateOutputAuthenticationTypePrometheus

Remote Write authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAuthenticationTypePrometheusNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAuthenticationTypePrometheus("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CreateOutputAuthenticationTypePrometheusNone`              | none                                                        |
| `CreateOutputAuthenticationTypePrometheusBasic`             | basic                                                       |
| `CreateOutputAuthenticationTypePrometheusCredentialsSecret` | credentialsSecret                                           |
| `CreateOutputAuthenticationTypePrometheusToken`             | token                                                       |
| `CreateOutputAuthenticationTypePrometheusTextSecret`        | textSecret                                                  |
| `CreateOutputAuthenticationTypePrometheusAwsSigv4`          | aws_sigv4                                                   |