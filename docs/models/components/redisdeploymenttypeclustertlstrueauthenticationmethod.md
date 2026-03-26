# RedisDeploymentTypeClusterTLSTrueAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterTLSTrueAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterTLSTrueAuthenticationMethod("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `RedisDeploymentTypeClusterTLSTrueAuthenticationMethodNone`              | none                                                                     |
| `RedisDeploymentTypeClusterTLSTrueAuthenticationMethodManual`            | manual                                                                   |
| `RedisDeploymentTypeClusterTLSTrueAuthenticationMethodCredentialsSecret` | credentialsSecret                                                        |
| `RedisDeploymentTypeClusterTLSTrueAuthenticationMethodTextSecret`        | textSecret                                                               |