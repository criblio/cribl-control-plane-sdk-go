# CreateOutputSystemByPackAuthenticationTypePrometheus

Remote Write authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackAuthenticationTypePrometheusNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackAuthenticationTypePrometheus("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `CreateOutputSystemByPackAuthenticationTypePrometheusNone`              | none                                                                    |
| `CreateOutputSystemByPackAuthenticationTypePrometheusBasic`             | basic                                                                   |
| `CreateOutputSystemByPackAuthenticationTypePrometheusCredentialsSecret` | credentialsSecret                                                       |
| `CreateOutputSystemByPackAuthenticationTypePrometheusToken`             | token                                                                   |
| `CreateOutputSystemByPackAuthenticationTypePrometheusTextSecret`        | textSecret                                                              |
| `CreateOutputSystemByPackAuthenticationTypePrometheusAwsSigv4`          | aws_sigv4                                                               |