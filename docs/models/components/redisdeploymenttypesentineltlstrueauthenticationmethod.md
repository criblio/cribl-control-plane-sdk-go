# RedisDeploymentTypeSentinelTLSTrueAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeSentinelTLSTrueAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeSentinelTLSTrueAuthenticationMethod("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `RedisDeploymentTypeSentinelTLSTrueAuthenticationMethodNone`              | none                                                                      |
| `RedisDeploymentTypeSentinelTLSTrueAuthenticationMethodManual`            | manual                                                                    |
| `RedisDeploymentTypeSentinelTLSTrueAuthenticationMethodCredentialsSecret` | credentialsSecret                                                         |
| `RedisDeploymentTypeSentinelTLSTrueAuthenticationMethodTextSecret`        | textSecret                                                                |