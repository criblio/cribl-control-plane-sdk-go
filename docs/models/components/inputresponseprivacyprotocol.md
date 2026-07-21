# InputResponsePrivacyProtocol

Privacy protocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponsePrivacyProtocolNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponsePrivacyProtocol("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `InputResponsePrivacyProtocolNone`    | none                                  |
| `InputResponsePrivacyProtocolDes`     | des                                   |
| `InputResponsePrivacyProtocolAes`     | aes                                   |
| `InputResponsePrivacyProtocolAes256b` | aes256b                               |
| `InputResponsePrivacyProtocolAes256r` | aes256r                               |