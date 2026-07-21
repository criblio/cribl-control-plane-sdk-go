# CreateInputAuthenticationProtocol

Authentication protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationProtocol("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateInputAuthenticationProtocolNone`   | none                                      |
| `CreateInputAuthenticationProtocolMd5`    | md5                                       |
| `CreateInputAuthenticationProtocolSha`    | sha                                       |
| `CreateInputAuthenticationProtocolSha224` | sha224                                    |
| `CreateInputAuthenticationProtocolSha256` | sha256                                    |
| `CreateInputAuthenticationProtocolSha384` | sha384                                    |
| `CreateInputAuthenticationProtocolSha512` | sha512                                    |