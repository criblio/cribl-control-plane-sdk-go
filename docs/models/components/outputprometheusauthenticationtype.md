# OutputPrometheusAuthenticationType

Remote Write authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputPrometheusAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputPrometheusAuthenticationType("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `OutputPrometheusAuthenticationTypeNone`              | none                                                  |
| `OutputPrometheusAuthenticationTypeBasic`             | basic                                                 |
| `OutputPrometheusAuthenticationTypeCredentialsSecret` | credentialsSecret                                     |
| `OutputPrometheusAuthenticationTypeToken`             | token                                                 |
| `OutputPrometheusAuthenticationTypeTextSecret`        | textSecret                                            |
| `OutputPrometheusAuthenticationTypeAwsSigv4`          | aws_sigv4                                             |