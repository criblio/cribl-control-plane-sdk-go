# RedisDeploymentTypeClusterTLSFalseAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterTLSFalseAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterTLSFalseAuthenticationMethod("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `RedisDeploymentTypeClusterTLSFalseAuthenticationMethodNone`              | none                                                                      |
| `RedisDeploymentTypeClusterTLSFalseAuthenticationMethodManual`            | manual                                                                    |
| `RedisDeploymentTypeClusterTLSFalseAuthenticationMethodCredentialsSecret` | credentialsSecret                                                         |
| `RedisDeploymentTypeClusterTLSFalseAuthenticationMethodTextSecret`        | textSecret                                                                |