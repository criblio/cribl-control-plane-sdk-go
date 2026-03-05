# SignatureVersionOptionsMinIo

Signature version to use for signing MinIO requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptionsMinIoV2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptionsMinIo("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `SignatureVersionOptionsMinIoV2` | v2                               |
| `SignatureVersionOptionsMinIoV4` | v4                               |