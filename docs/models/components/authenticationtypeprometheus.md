# AuthenticationTypePrometheus

Remote Write authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypePrometheusNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypePrometheus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AuthenticationTypePrometheusNone`              | none                                            |
| `AuthenticationTypePrometheusBasic`             | basic                                           |
| `AuthenticationTypePrometheusCredentialsSecret` | credentialsSecret                               |
| `AuthenticationTypePrometheusToken`             | token                                           |
| `AuthenticationTypePrometheusTextSecret`        | textSecret                                      |
| `AuthenticationTypePrometheusAwsSigv4`          | aws_sigv4                                       |