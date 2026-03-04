# InputElasticAuthenticationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputElasticAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputElasticAuthenticationType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `InputElasticAuthenticationTypeNone`              | none                                              |
| `InputElasticAuthenticationTypeBasic`             | basic                                             |
| `InputElasticAuthenticationTypeCredentialsSecret` | credentialsSecret                                 |
| `InputElasticAuthenticationTypeAuthTokens`        | authTokens                                        |