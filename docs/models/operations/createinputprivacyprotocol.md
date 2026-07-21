# CreateInputPrivacyProtocol

Privacy protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputPrivacyProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputPrivacyProtocol("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CreateInputPrivacyProtocolNone`    | none                                |
| `CreateInputPrivacyProtocolDes`     | des                                 |
| `CreateInputPrivacyProtocolAes`     | aes                                 |
| `CreateInputPrivacyProtocolAes256b` | aes256b                             |
| `CreateInputPrivacyProtocolAes256r` | aes256r                             |