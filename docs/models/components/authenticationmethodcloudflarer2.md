# AuthenticationMethodCloudflareR2

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodCloudflareR2Auto

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodCloudflareR2("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AuthenticationMethodCloudflareR2Auto`   | auto                                     |
| `AuthenticationMethodCloudflareR2Secret` | secret                                   |