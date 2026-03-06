# CreateInputSystemByPackAuthenticationTypeElastic

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthenticationTypeElasticNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthenticationTypeElastic("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CreateInputSystemByPackAuthenticationTypeElasticNone`              | none                                                                |
| `CreateInputSystemByPackAuthenticationTypeElasticBasic`             | basic                                                               |
| `CreateInputSystemByPackAuthenticationTypeElasticCredentialsSecret` | credentialsSecret                                                   |
| `CreateInputSystemByPackAuthenticationTypeElasticAuthTokens`        | authTokens                                                          |