# AuthenticationMethodOptionsS3CollectorConf

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsS3CollectorConfAuto

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsS3CollectorConf("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `AuthenticationMethodOptionsS3CollectorConfAuto`   | auto                                               |
| `AuthenticationMethodOptionsS3CollectorConfManual` | manual                                             |
| `AuthenticationMethodOptionsS3CollectorConfSecret` | secret                                             |