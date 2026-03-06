# OutputCloudflareR2AuthenticationMethod

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputCloudflareR2AuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputCloudflareR2AuthenticationMethod("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `OutputCloudflareR2AuthenticationMethodAuto`   | auto                                           |
| `OutputCloudflareR2AuthenticationMethodSecret` | secret                                         |