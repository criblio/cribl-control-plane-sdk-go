# CreateInputAuthMethodsExtAuthenticationType

Authentication type

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthMethodsExtAuthenticationTypeToken

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthMethodsExtAuthenticationType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateInputAuthMethodsExtAuthenticationTypeToken`       | token                                                    |
| `CreateInputAuthMethodsExtAuthenticationTypeTokenSecret` | tokenSecret                                              |
| `CreateInputAuthMethodsExtAuthenticationTypeBasic`       | basic                                                    |
| `CreateInputAuthMethodsExtAuthenticationTypeBasicSecret` | basicSecret                                              |