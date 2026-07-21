# InputResponseAuthenticationProtocol

Authentication protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseAuthenticationProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseAuthenticationProtocol("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `InputResponseAuthenticationProtocolNone`   | none                                        |
| `InputResponseAuthenticationProtocolMd5`    | md5                                         |
| `InputResponseAuthenticationProtocolSha`    | sha                                         |
| `InputResponseAuthenticationProtocolSha224` | sha224                                      |
| `InputResponseAuthenticationProtocolSha256` | sha256                                      |
| `InputResponseAuthenticationProtocolSha384` | sha384                                      |
| `InputResponseAuthenticationProtocolSha512` | sha512                                      |