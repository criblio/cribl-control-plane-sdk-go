# CreateOutputAuthenticationMethodCloudflareR2

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAuthenticationMethodCloudflareR2Auto

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAuthenticationMethodCloudflareR2("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CreateOutputAuthenticationMethodCloudflareR2Auto`   | auto                                                 |
| `CreateOutputAuthenticationMethodCloudflareR2Secret` | secret                                               |