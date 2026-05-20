# AuthenticationTypeElastic

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeElasticNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeElastic("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AuthenticationTypeElasticNone`              | none                                         |
| `AuthenticationTypeElasticBasic`             | basic                                        |
| `AuthenticationTypeElasticCredentialsSecret` | credentialsSecret                            |
| `AuthenticationTypeElasticAuthTokens`        | authTokens                                   |