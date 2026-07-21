# InputSnmpPrivacyProtocol

Privacy protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSnmpPrivacyProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputSnmpPrivacyProtocol("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `InputSnmpPrivacyProtocolNone`    | none                              |
| `InputSnmpPrivacyProtocolDes`     | des                               |
| `InputSnmpPrivacyProtocolAes`     | aes                               |
| `InputSnmpPrivacyProtocolAes256b` | aes256b                           |
| `InputSnmpPrivacyProtocolAes256r` | aes256r                           |