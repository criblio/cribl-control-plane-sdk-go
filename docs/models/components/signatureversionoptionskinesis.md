# SignatureVersionOptionsKinesis

Signature version to use for signing Kinesis stream requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptionsKinesisV2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptionsKinesis("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `SignatureVersionOptionsKinesisV2` | v2                                 |
| `SignatureVersionOptionsKinesisV4` | v4                                 |