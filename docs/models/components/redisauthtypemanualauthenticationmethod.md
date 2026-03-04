# RedisAuthTypeManualAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
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