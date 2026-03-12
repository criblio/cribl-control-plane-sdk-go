# CreateOutputAwsAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAwsAuthenticationMethodAuto

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAwsAuthenticationMethod("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateOutputAwsAuthenticationMethodAuto`    | auto                                         |
| `CreateOutputAwsAuthenticationMethodAutoRPC` | auto_rpc                                     |
| `CreateOutputAwsAuthenticationMethodManual`  | manual                                       |