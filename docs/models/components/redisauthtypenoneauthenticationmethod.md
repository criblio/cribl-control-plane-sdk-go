# RedisAuthTypeNoneAuthenticationMethod

Authentication method to use when connecting to Redis.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeNoneAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeNoneAuthenticationMethod("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `RedisAuthTypeNoneAuthenticationMethodNone`              | none                                                     |
| `RedisAuthTypeNoneAuthenticationMethodManual`            | manual                                                   |
| `RedisAuthTypeNoneAuthenticationMethodCredentialsSecret` | credentialsSecret                                        |
| `RedisAuthTypeNoneAuthenticationMethodTextSecret`        | textSecret                                               |