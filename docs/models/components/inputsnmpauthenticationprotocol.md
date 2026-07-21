# InputSnmpAuthenticationProtocol

Authentication protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSnmpAuthenticationProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputSnmpAuthenticationProtocol("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InputSnmpAuthenticationProtocolNone`   | none                                    |
| `InputSnmpAuthenticationProtocolMd5`    | md5                                     |
| `InputSnmpAuthenticationProtocolSha`    | sha                                     |
| `InputSnmpAuthenticationProtocolSha224` | sha224                                  |
| `InputSnmpAuthenticationProtocolSha256` | sha256                                  |
| `InputSnmpAuthenticationProtocolSha384` | sha384                                  |
| `InputSnmpAuthenticationProtocolSha512` | sha512                                  |