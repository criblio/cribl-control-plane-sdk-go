# RedisAuthTypeTextSecretAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeTextSecretAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeTextSecretAuthenticationMethod("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RedisAuthTypeTextSecretAuthenticationMethodNone`              | none                                                           |
| `RedisAuthTypeTextSecretAuthenticationMethodManual`            | manual                                                         |
| `RedisAuthTypeTextSecretAuthenticationMethodCredentialsSecret` | credentialsSecret                                              |
| `RedisAuthTypeTextSecretAuthenticationMethodTextSecret`        | textSecret                                                     |