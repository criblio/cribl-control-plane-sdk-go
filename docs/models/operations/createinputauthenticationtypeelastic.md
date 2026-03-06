# CreateInputAuthenticationTypeElastic

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationTypeElasticNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationTypeElastic("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateInputAuthenticationTypeElasticNone`              | none                                                    |
| `CreateInputAuthenticationTypeElasticBasic`             | basic                                                   |
| `CreateInputAuthenticationTypeElasticCredentialsSecret` | credentialsSecret                                       |
| `CreateInputAuthenticationTypeElasticAuthTokens`        | authTokens                                              |