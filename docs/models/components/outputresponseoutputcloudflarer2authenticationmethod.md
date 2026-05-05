# OutputResponseOutputCloudflareR2AuthenticationMethod

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputCloudflareR2AuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputCloudflareR2AuthenticationMethod("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `OutputResponseOutputCloudflareR2AuthenticationMethodAuto`   | auto                                                         |
| `OutputResponseOutputCloudflareR2AuthenticationMethodSecret` | secret                                                       |