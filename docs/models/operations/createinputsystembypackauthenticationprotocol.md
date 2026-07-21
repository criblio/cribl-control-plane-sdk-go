# CreateInputSystemByPackAuthenticationProtocol

Authentication protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthenticationProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthenticationProtocol("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `CreateInputSystemByPackAuthenticationProtocolNone`   | none                                                  |
| `CreateInputSystemByPackAuthenticationProtocolMd5`    | md5                                                   |
| `CreateInputSystemByPackAuthenticationProtocolSha`    | sha                                                   |
| `CreateInputSystemByPackAuthenticationProtocolSha224` | sha224                                                |
| `CreateInputSystemByPackAuthenticationProtocolSha256` | sha256                                                |
| `CreateInputSystemByPackAuthenticationProtocolSha384` | sha384                                                |
| `CreateInputSystemByPackAuthenticationProtocolSha512` | sha512                                                |