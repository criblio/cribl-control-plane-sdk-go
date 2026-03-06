# CreateOutputSystemByPackAuthenticationMethodCloudflareR2

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackAuthenticationMethodCloudflareR2Auto

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackAuthenticationMethodCloudflareR2("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CreateOutputSystemByPackAuthenticationMethodCloudflareR2Auto`   | auto                                                             |
| `CreateOutputSystemByPackAuthenticationMethodCloudflareR2Secret` | secret                                                           |