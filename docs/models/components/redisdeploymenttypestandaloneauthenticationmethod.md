# RedisDeploymentTypeStandaloneAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeStandaloneAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeStandaloneAuthenticationMethod("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `RedisDeploymentTypeStandaloneAuthenticationMethodNone`              | none                                                                 |
| `RedisDeploymentTypeStandaloneAuthenticationMethodManual`            | manual                                                               |
| `RedisDeploymentTypeStandaloneAuthenticationMethodCredentialsSecret` | credentialsSecret                                                    |
| `RedisDeploymentTypeStandaloneAuthenticationMethodTextSecret`        | textSecret                                                           |