# RedisAuthTypeManualAuthenticationMethod

Authentication method to use when connecting to Redis.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeManualAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeManualAuthenticationMethod("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `RedisAuthTypeManualAuthenticationMethodNone`              | none                                                       |
| `RedisAuthTypeManualAuthenticationMethodManual`            | manual                                                     |
| `RedisAuthTypeManualAuthenticationMethodCredentialsSecret` | credentialsSecret                                          |
| `RedisAuthTypeManualAuthenticationMethodTextSecret`        | textSecret                                                 |