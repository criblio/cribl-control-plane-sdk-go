# RedisAuthTypeCredentialsSecretAuthenticationMethod

Authentication method to use when connecting to Redis.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeCredentialsSecretAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeCredentialsSecretAuthenticationMethod("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `RedisAuthTypeCredentialsSecretAuthenticationMethodNone`              | none                                                                  |
| `RedisAuthTypeCredentialsSecretAuthenticationMethodManual`            | manual                                                                |
| `RedisAuthTypeCredentialsSecretAuthenticationMethodCredentialsSecret` | credentialsSecret                                                     |
| `RedisAuthTypeCredentialsSecretAuthenticationMethodTextSecret`        | textSecret                                                            |