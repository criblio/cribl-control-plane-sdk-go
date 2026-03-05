# RedisDeploymentTypeSentinelTLSFalseAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeSentinelTLSFalseAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeSentinelTLSFalseAuthenticationMethod("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `RedisDeploymentTypeSentinelTLSFalseAuthenticationMethodNone`              | none                                                                       |
| `RedisDeploymentTypeSentinelTLSFalseAuthenticationMethodManual`            | manual                                                                     |
| `RedisDeploymentTypeSentinelTLSFalseAuthenticationMethodCredentialsSecret` | credentialsSecret                                                          |
| `RedisDeploymentTypeSentinelTLSFalseAuthenticationMethodTextSecret`        | textSecret                                                                 |